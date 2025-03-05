import * as dotenv from 'dotenv';

dotenv.config();

export const getConfig = () => ({
    port: process.env.PORT || 3000,
    serviceName: 'api-gateway',
});