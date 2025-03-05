import axios from 'axios';
import { Agent } from 'http';

interface ConsulService
{
    Address: string;
    Port: number;
}

export async function registerConsul(
    serviceName: string,
    servicePort: number,
    config: { host?: string; port?: number; token?: string; agent?: Agent } = {}
): Promise<void>
{
    const { host, port, token, agent } = config;

    try {
        const response = await axios.put(
            `http://${host}:${port}/v1/agent/service/register`,
            {
                Name: serviceName,
                Port: servicePort,
                Check: {
                    interval: '10s',
                    http: `http://localhost:${servicePort}/health`,
                },
            },
            {
                headers: {
                    'X-Consul-Token': token || '',
                },
                httpAgent: agent,
            }
        );

        if (response.status !== 200) {
            throw new Error(`Consul registration failed: ${response.status}`);
        }
    } catch (error) {
        console.error('Consul registration error:', error);
        throw error;
    }
}

export async function getService(
    serviceName: string,
    config: { host?: string; port?: number; token?: string; agent?: Agent } = {}
): Promise<ConsulService | null>
{
    const { host, port, token, agent } = config;

    try {
        const response = await axios.get(
            `http://${host}:${port}/v1/catalog/service/${serviceName}`,
            {
                headers: {
                    'X-Consul-Token': token || '',
                },
                httpAgent: agent,
            }
        );

        if (response.data.length === 0) return null;

        return {
            Address: response.data[0].ServiceAddress || response.data[0].Address,
            Port: response.data[0].ServicePort,
        };
    } catch (error) {
        console.error('Consul service discovery error:', error);
        throw error;
    }
}