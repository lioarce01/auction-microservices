import 'reflect-metadata';
import Fastify from 'fastify';

import routes from './routesIndex';
import AuthPlugin from '@Auth/Plugins/AuthPlugin';

import fastifyHelmet from '@fastify/helmet';

import setupContainer from '@Shared/DI/DIContainer';

import { setupSwagger } from '@Shared/Config/swaggerConfig';

import { errorHandler } from '@Main/Infrastructure/Errors/ErrorHandler';

import APIConfig from '@Shared/Config/serverConfig';

import client from 'prom-client';
import { config } from 'dotenv';
import { ConsulAdapter } from '@Service-Discovery/Infrastructure/Repositories/ConsulRepository';

config()


setupContainer();

const app = Fastify({ logger: true });

setupSwagger(app)

app.register(fastifyHelmet)
app.register(AuthPlugin);
app.register(routes, { prefix: `/api/${APIConfig.VERSION}` });
app.setErrorHandler(errorHandler);

const collectDefaultMetrics = client.collectDefaultMetrics;
collectDefaultMetrics({ register: client.register });

app.get('/metrics', async (req, reply) =>
{
  try {
    reply.header('Content-Type', client.register.contentType);
    reply.send(await client.register.metrics());
  } catch (err) {
    reply.status(500).send(err);
  }
});

const serviceDiscovery = new ConsulAdapter(
  APIConfig.serviceName,
  APIConfig.PORT,
  {
    host: process.env.CONSUL_HOST,
    port: parseInt(process.env.CONSUL_PORT || '8500')
  }
)

const start = async () =>
{
  try {

    await serviceDiscovery.registerService();

    await app.listen({ port: APIConfig.PORT, host: '0.0.0.0' });
    console.log(`Server running on port ${APIConfig.PORT}`);
  }
  catch (err) {
    app.log.error(err);
    process.exit(1);
  }
};

start();
