export interface ServiceDiscovery
{
    registerService(): Promise<void>;
    discoverService(name: string): Promise<{ host: string; port: number } | null>;
}