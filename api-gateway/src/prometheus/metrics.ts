import { Express } from 'express';
import client from 'prom-client';

export function registerPrometheusMetrics(app: Express)
{
    client.collectDefaultMetrics();
    app.get('/metrics', async (req, res) =>
    {
        res.set('Content-Type', client.register.contentType);
        res.end(await client.register.metrics());
    });
}