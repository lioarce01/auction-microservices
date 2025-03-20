import { CoreApp } from './app/core/app';
import { ConsulAdapter } from './app/adapters/consul';
import { PrometheusAdapter } from './app/adapters/prometheus';
import { proxyMiddleware } from './app/middleware/proxy';
import { getConfig } from './config/env';
import { authMiddleware } from './app/middleware/authMiddleware';

async function startServer()
{
    const config = getConfig();
    const core = new CoreApp();
    config.port = parseInt(config.port as string, 10);

    // Health Check
    core.addRoute('/health', (req, res) => res.send('OK'));

    // Prometheus
    const prometheus = new PrometheusAdapter(core.getApp());
    core.addRoute('/metrics', prometheus.handleMetrics.bind(prometheus));

    // Consul
    const serviceDiscovery = new ConsulAdapter(
        config.serviceName,
        config.port,
        {
            host: process.env.CONSUL_HOST,
            port: parseInt(process.env.CONSUL_PORT || '8500')
        }
    );

    try {
        await serviceDiscovery.registerService();

        core.getApp().use(authMiddleware)

        // Proxy Middleware
        core.getApp().use(
            '/service/:serviceName/*',
            proxyMiddleware(serviceDiscovery)
        );

        // Iniciar servidor
        core.getApp().listen(config.port, () =>
        {
            console.log(`API Gateway running on port ${config.port}`);
        });
    } catch (error) {
        console.error('Failed to start server:', error);
        process.exit(1);
    }
}

startServer();