import express, { Application } from 'express';
import cors from 'cors'

export class CoreApp
{
    private app: express.Application;

    constructor()
    {
        this.app = express();
        this.configureMiddlewares();
    }

    private configureMiddlewares()
    {
        this.app.use(cors({
            origin: '*',
            methods: ['GET', 'POST', 'PUT', 'DELETE', 'OPTIONS'],
            allowedHeaders: ['Content-Type', 'Authorization', 'X-Requested-With']
        }));
        this.app.use(express.json());
        this.app.use(express.urlencoded({ extended: true }));
    }

    public getApp(): Application
    {
        return this.app;
    }

    public addRoute(route: string, handler: (req: any, res: any) => void)
    {
        this.app.get(route, handler);
    }
}