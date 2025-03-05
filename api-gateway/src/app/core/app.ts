import express, { Application } from 'express';

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