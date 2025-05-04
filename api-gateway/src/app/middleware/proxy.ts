import http from 'http';
import { ServiceDiscovery } from '../interfaces/service-discovery';

export const proxyMiddleware = (serviceDiscovery: ServiceDiscovery) =>
{
    return async (req: any, res: any) =>
    {
        console.log(`Received ${req.method} request to proxy for service: ${req.params.serviceName}`);
        console.log("🔹 Headers received:", req.headers);
        console.log("Request body:", req.body);

        const serviceName = req.params.serviceName;

        try {
            console.log(`Discovering service: ${serviceName}`);
            const service = await serviceDiscovery.discoverService(serviceName);

            if (!service) {
                console.log(`Service ${serviceName} not found`);
                return res.status(404).json({ error: `Service ${serviceName} not found` });
            }

            console.log(`Discovered service: ${JSON.stringify(service)}`);

            const options = {
                hostname: service.host,
                port: service.port,
                path: req.originalUrl.replace(`/service/${serviceName}`, ''),
                method: req.method,
                headers: req.headers,
            };

            console.log(`Proxying request to: ${options.hostname}:${options.port}${options.path}`);

            const proxyReq = http.request({
                ...options,
                timeout: 30000
            },
                (proxyRes) =>
                {
                    console.log(`Received response from proxied service with status: ${proxyRes.statusCode}`);
                    res.writeHead(proxyRes.statusCode || 500, proxyRes.headers);
                    proxyRes.pipe(res);
                });

            proxyReq.on('timeout', () =>
            {
                console.error('Proxy request timed out');
                proxyReq.destroy();
            });

            if (req.method === 'POST' || req.method === 'PUT') {
                if (req.body && Object.keys(req.body).length > 0) {
                    const bodyData = JSON.stringify(req.body);
                    proxyReq.setHeader('Content-Length', Buffer.byteLength(bodyData));
                    proxyReq.write(bodyData);
                    proxyReq.end();
                    console.log('Body sent to destination service');
                } else {
                    req.pipe(proxyReq);
                }
            } else {
                proxyReq.end();
            }

            proxyReq.on('error', (err: any) =>
            {
                console.error(`Proxy error: ${err.message}`);
                res.status(500).json({ error: 'Proxy request failed' });
            });

            console.log("Proxy request initiated");
        } catch (error) {
            console.error(`Service discovery error: ${(error as Error).message}`);
            res.status(500).json({ error: 'Service discovery failed' });
        }
    };
};