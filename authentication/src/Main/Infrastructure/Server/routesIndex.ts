import { FastifyInstance } from 'fastify';
import setupContainer from '@Shared/DI/DIContainer';
import userRoutes from '@User/Infrastructure/HTTP/Routes/UserRoutes';
import client from 'prom-client';


export default async function routes(fastify: FastifyInstance)
{
  setupContainer();
  // register health route
  fastify.get('/health', async (req, res) =>
  {
    res.send('OK');
  });

  // // register metrics route
  // fastify.get('/metrics', async (req, reply) =>
  // {
  //   try {
  //     reply.header('Content-Type', client.register.contentType);
  //     reply.send(await client.register.metrics());
  //   } catch (err) {
  //     reply.status(500).send(err);
  //   }
  // });

  // register user routes
  fastify.register(userRoutes, { prefix: '/users' });
}
