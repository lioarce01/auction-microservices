import { Metrics } from "../Domain/Repositories/Metrics";

class MetricService
{
    private metricRepository: Metrics;

    constructor(metricRepository: Metrics)
    {
        this.metricRepository = metricRepository;
    }

    async handleRequest()
    {
        const end = this.metricRepository.startRequestTimer();
        end();
        await this.metricRepository.incrementRequestCount();
    }
}

export default MetricService;
