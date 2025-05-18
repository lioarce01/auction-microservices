# Auction Microservices Platform

A microservices-based auction platform designed for scalability, modularity, and real-time operations.

---

## Overview

This project implements an auction platform using a microservices architecture. Each service is responsible for a specific domain, ensuring separation of concerns and ease of maintenance.

---

## Architecture

The system comprises the following microservices:

* **Authentication Service**: Handles user registration, login, and JWT-based authentication.
* **Auctions Service**: Manages auction creation, bidding, and retrieval.
* **Payments Service**: Processes payments and handles transaction records.
* **API Gateway**: Serves as a single entry point, routing requests to appropriate services.
* **Prometheus**: Monitors services and collects metrics.

All services communicate over HTTP and are containerized using Docker.

---

## Technologies Used

* **Languages**: TypeScript (54.1%), Go (40.6%), JavaScript (3.4%)
* **Containerization**: Docker & Docker Compose
* **Monitoring**: Prometheus
* **Authentication**: JWT

---

## Getting Started

### Prerequisites

* Docker and Docker Compose installed on your machine.

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/lioarce01/auction-microservices.git
   cd auction-microservices
   ```

2. **Set up environment variables**:

   Copy the example environment file and modify as needed:

   ```bash
   cp .env.example .env
   ```

3. **Build and run the services**:

   ```bash
   docker-compose up --build
   ```

   This command will build and start all the microservices along with Prometheus for monitoring.

---

## Usage

Once all services are up and running:

* **API Gateway**: Accessible at `http://localhost:PORT` (replace `PORT` with the configured port).
* **Prometheus Dashboard**: Accessible at `http://localhost:9090`.

You can interact with the API Gateway to perform actions such as user registration, login, creating auctions, placing bids, and processing payments.

---

## Project Structure

```
auction-microservices/
├── api-gateway/        # Handles routing to microservices
├── auctions/           # Auction service
├── authentication/     # Authentication service
├── payments/           # Payments service
├── db-init/            # Database initialization scripts
├── prometheus/         # Prometheus configuration
├── .env.example        # Example environment variables
├── docker-compose.yml  # Docker Compose configuration
└── ...
```

---
