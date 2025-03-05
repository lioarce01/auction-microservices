import { ServiceDiscovery } from '@Service-Discovery/Domain/Repositories/ServiceDiscovery';
import Consul from 'consul';
import { config } from 'dotenv';

config()

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
            address: process.env.HOST_IP || 'authentication',
            port,
            check: {
                name: `${serviceName} check`,
                http: `http://${process.env.HOST_IP || 'authentication'}:${port}/api/v1/health`,
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
}