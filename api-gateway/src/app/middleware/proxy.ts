import http from 'http';
import { ServiceDiscovery } from '../interfaces/service-discovery';

export const proxyMiddleware = (serviceDiscovery: ServiceDiscovery) =>
{
    return async (req: any, res: any) =>
    {
        console.log("🔹 Headers que se envían al servicio:", req.headers);

        const serviceName = req.params.serviceName;

        try {
            const service = await serviceDiscovery.discoverService(serviceName);

            if (!service) {
                return res.status(404).json({ error: `Service ${serviceName} not found` });
            }

            const options = {
                hostname: service.host,
                port: service.port,
                path: req.originalUrl.replace(`/service/${serviceName}`, ''),
                method: req.method,
                headers: req.headers,
            };

            console.log("🔹 Headers que se envían al servicio:", options.headers);

            const proxyReq = http.request(options, (proxyRes) =>
            {
                res.writeHead(proxyRes.statusCode || 500, proxyRes.headers);
                proxyRes.pipe(res);
            });

            req.pipe(proxyReq);
            proxyReq.on('error', (err: any) =>
            {
                console.error(`Proxy error: ${err.message}`);
                res.status(500).json({ error: 'Proxy request failed' });
            });
        } catch (error) {
            console.error(`Service discovery error: ${(error as Error).message}`);
            res.status(500).json({ error: 'Service discovery failed' });
        }
    };
};