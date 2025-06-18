# Ethical E-Commerce Platform 🌱

A microservices-based e-commerce system focused on sustainability and eco-friendly products.

## 🏗 Project Structure

- `/services` → Backend microservices (Go, Node.js)
- `/gateway` → API Gateway (NestJS)
- `/frontend` → Angular (Admin), React (Storefront)
- `/ml` → Python ML recommender
- `/infra` → Kubernetes, Helm, Terraform
- `/docs` → ERD, architecture, design records

## 🚀 Tech Stack

| Layer             | Tools                                  |
|-------------------|---------------------------------------|
| Backend          | Go, Node.js                            |
| Databases        | PostgreSQL, MongoDB                    |
| Search           | Elasticsearch                          |
| Messaging        | RabbitMQ                               |
| Frontend         | React, Angular                         |
| Orchestration    | Docker, Kubernetes (minikube)          |
| CI/CD           | GitHub Actions                         |
| Monitoring       | Prometheus + Grafana                   |
| Logging         | ELK Stack (Elasticsearch, Logstash, Kibana) |
| Infra-as-Code    | Terraform (optional)                   |

## 📖 Documentation

- [Architecture](./docs/architecture.md)
- [ER Diagram](./docs/erd.md)
- [Deployment](./docs/deployment.md)
- [Testing](./docs/testing.md)
- [Security](./docs/security.md)

## 📦 Getting Started

```bash
git clone https://github.com/juthakor/EcoMerce.git
cd EcoMerce
docker-compose up --build
```

## 🔒 Secrets & Environment Variables

For local development, copy `.env.example` to `.env` and fill in your own values.

```bash
cp .env.example .env
```

### Setup MongoDB schema validation
```bash
make product-setup-validator
```