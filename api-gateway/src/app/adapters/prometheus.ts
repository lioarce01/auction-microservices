import client from 'prom-client';
import express from 'express'
import { registerPrometheusMetrics } from '../prometheus/metrics';

export class PrometheusAdapter
{
    constructor(private app: express.Application)
    {
        registerPrometheusMetrics(this.app);
        client.collectDefaultMetrics();
    }

    public async handleMetrics(req: any, res: any)
    {
        res.set('Content-Type', client.register.contentType);
        res.end(await client.register.metrics());
    }
}