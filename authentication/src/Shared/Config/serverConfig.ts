import { config } from 'dotenv';

config();

const APIConfig = {
  PORT: Number(process.env.PORT),
  VERSION: 'v1',
  serviceName: 'authentication',
};

export default APIConfig;
