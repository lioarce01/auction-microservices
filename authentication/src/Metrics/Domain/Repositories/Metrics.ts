export interface Metrics
{
    incrementRequestCount(): Promise<void>;
    observeRequestDuration(duration: number): Promise<void>;
    startRequestTimer(): () => void;
}
