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
        // Simula la lógica de procesamiento...
        end(); // Registra la duración de la solicitud
        await this.metricRepository.incrementRequestCount();
    }
}

export default MetricService;
