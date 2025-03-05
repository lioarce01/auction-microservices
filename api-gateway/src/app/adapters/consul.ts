import Consul from 'consul';
import { ServiceDiscovery } from '../interfaces/service-discovery';

export class ConsulAdapter implements ServiceDiscovery
{
    private client: Consul;
    private serviceDetails: {
        name: string;
        address: string;
        port: number;
        check: {
            name: string;
            http: string;
            interval: string;
            timeout: string;
        };
    };

    constructor(
        serviceName: string,
        port: number,
        private config: {
            host?: string;
            port?: number;
        } = {}
    )
    {
        this.client = new Consul({
            host: this.config.host || 'consul',
            port: this.config.port || 8500
        });

        this.serviceDetails = {
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
    }

    async registerService(): Promise<void>
    {
        try {
            await this.client.agent.service.register(this.serviceDetails);
            console.log(`${this.serviceDetails.name} registered with Consul successfully.`);
        } catch (err) {
            console.error(`Error registering service with Consul:`, err);
            throw err;
        }
    }

    async discoverService(name: string): Promise<{ host: string; port: number } | null>
    {
        try {
            const services = await this.client.agent.service.list();
            const service = services[name];

            if (!service) return null;

            return {
                host: service.Address,
                port: service.Port
            };
        } catch (err) {
            console.error(`Error discovering service ${name}:`, err);
            return null;
        }
    }
}