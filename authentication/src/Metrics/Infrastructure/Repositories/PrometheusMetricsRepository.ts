import client from 'prom-client';
import { Metrics } from 'src/Metrics/Domain/Repositories/Metrics';

class PrometheusMetricRepository implements Metrics
{
    private requestCount: client.Counter<string>;
    private requestDuration: client.Histogram<string>;

    constructor()
    {
        this.requestCount = new client.Counter({
            name: 'request_count',
            help: 'Total number of requests',
        });

        this.requestDuration = new client.Histogram({
            name: 'request_duration_seconds',
            help: 'Duration of requests in seconds',
            buckets: [0.1, 0.5, 1, 2.5, 5, 10],
        });
    }
    startRequestTimer(): () => void
    {
        const end = this.requestDuration.startTimer();
        return () =>
        {
            end();
        };
    }

    async incrementRequestCount(): Promise<void>
    {
        this.requestCount.inc();
    }

    async observeRequestDuration(duration: number): Promise<void>
    {
        this.requestDuration.observe(duration);
    }

}

export default PrometheusMetricRepository;