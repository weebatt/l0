# L0 Order Service
## Features

- Order, delivery, payment, and items management
- Kafka consumer for order messages
- In-memory cache with cache warm-up from Postgres
- Prometheus metrics and Grafana dashboards
- Graceful shutdown and validation

## Quick Start

### 1. Clone and build

```sh
git clone <your-repo-url>
cd l0
docker compose -f docker-compose.prod.yaml up --build
```

### 3. Metrics

- Prometheus: [http://localhost:9090](http://localhost:9090)
- Grafana: [http://localhost:3000](http://localhost:3000)
- Service metrics endpoint: [http://localhost:8081/metrics](http://localhost:8081/metrics)

### 4. Health Check

- [http://localhost:8081/health](http://localhost:8081/health)

### 5. Example API Request

```sh
curl "http://localhost:8081/order?order_uid=<uuid>"
```

## Resources

- [The service connects to a message broker (Kafka) and processes messages online and The HTTP server returns correct data in JSON format](resources/kafka_and_response.mov)
- [The interface displays the data in a clear way after entering the ID and pressing the button and the cache really speeds up data retrieval](resources/metrics.mov)
