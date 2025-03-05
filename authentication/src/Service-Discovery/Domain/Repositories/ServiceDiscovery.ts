export interface ServiceDiscovery
{
    registerService(): Promise<void>;
}