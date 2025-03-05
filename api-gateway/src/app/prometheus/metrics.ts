import express from 'express';
import client from 'prom-client';

export function registerPrometheusMetrics(app: express.Application)
{
    // Configuración de Prometheus
    app.use('/metrics', async (req, res) =>
    {
        res.set('Content-Type', client.register.contentType);
        res.end(await client.register.metrics());
    });
}