import axios from 'axios';
import { Agent } from 'http';


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