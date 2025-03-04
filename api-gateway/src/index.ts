import express, { Request, Response } from 'express';
import { registerPrometheusMetrics } from './prometheus/metrics';
import { registerConsul, getService } from './service-discovery/consul';
import http from 'http';
import client from 'prom-client';

const app = express();
const PORT = process.env.PORT || 3000;
const SERVICE_NAME = 'api-gateway';

registerPrometheusMetrics(app);
registerConsul(SERVICE_NAME, Number(PORT));

// Health Check
app.get('/health', (req, res) => res.send('OK'));

// Prometheus Metrics
client.collectDefaultMetrics();
app.get('/metrics', async (req: Request, res: Response) =>
{
    res.set('Content-Type', client.register.contentType);
    res.end(await client.register.metrics());
});

// Dynamic Service Discovery and Proxy
app.use('/service/:serviceName/*', async (req: Request, res: Response) =>
{
    const serviceName = req.params.serviceName;
    const service = await getService(serviceName);

    if (!service) {
        return res.status(404).json({ error: `Service ${serviceName} not found` });
    }

    const options = {
        hostname: service.Address,
        port: service.Port,
        path: req.originalUrl.replace(`/service/${serviceName}`, ''),
        method: req.method,
        headers: req.headers,
    };

    const proxyReq = http.request(options, (proxyRes) =>
    {
        res.writeHead(proxyRes.statusCode || 500, proxyRes.headers);
        proxyRes.pipe(res, { end: true });
    });

    req.pipe(proxyReq, { end: true });
    proxyReq.on('error', (err) =>
    {
        console.error(`Proxy error: ${err.message}`);
        res.status(500).json({ error: 'Proxy request failed' });
    });
});

app.listen(PORT, () =>
{
    console.log(`API Gateway listening on port ${PORT}`);
});