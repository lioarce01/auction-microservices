import Consul from 'consul';

export function registerConsul(serviceName: string, port: number)
{
    const consul = new Consul({
        host: process.env.CONSUL_HOST || 'consul',
        port: Number(process.env.CONSUL_PORT) || 8500
    });

    const details = {
        name: serviceName,
        address: process.env.HOST_IP || 'api-gateway',
        port,
        check: {
            name: `${serviceName} check`,
            http: `http://${process.env.HOST_IP || 'api-gateway'}:${port}/health`,
            interval: '10s',
            timeout: '5s'
        }
    };

    consul.agent.service.register(details)
        .then(() =>
        {
            console.log(`${serviceName} registered with Consul successfully.`);
        })
        .catch((err: any) =>
        {
            console.error(`Error registering ${serviceName} with Consul:`, err);
        });
}

export async function getService(serviceName: string)
{
    const consul = new Consul({
        host: process.env.CONSUL_HOST || 'consul',
        port: Number(process.env.CONSUL_PORT) || 8500
    });

    try {
        const services = await consul.agent.service.list();
        return services[serviceName];
    } catch (err) {
        console.error(`Error fetching service ${serviceName} from Consul:`, err);
        return null;
    }
}