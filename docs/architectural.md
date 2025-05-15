# 📘 Architectural Decisions

This document outlines the key architectural decisions and reasoning behind the Ethical E-Commerce Platform.

---

## 🏗 System Overview

We designed the platform as a **microservices-based architecture** to separate concerns, enable independent scaling, and showcase modern system design.  
The system consists of:

- Multiple backend microservices (Go, Node.js, Python)
- An API Gateway (NestJS)
- Separate frontends (Angular for admin, React for customer)
- Asynchronous messaging (RabbitMQ)
- Databases (PostgreSQL + MongoDB)
- Containerized deployment (Docker, Kubernetes)
- Integrated monitoring, logging, and CI/CD

---

## 🔑 Major Decisions

---

### 1️⃣ Monorepo vs Polyrepo

✅ **Decision:** Use a **monorepo** (single GitHub repository).

- Easier to manage as a solo developer.
- Single place to show all components in the portfolio.
- Allows shared configuration, documentation, and orchestration (e.g., docker-compose, CI/CD).

---

### 2️⃣ Backend Service Language Split

✅ **Decision:**  
- **Go** for Auth + Order services → performance, concurrency, small binaries.
- **Node.js** for Product service → fast prototyping, rich MongoDB ecosystem.
- **Python** for ML Recommender → strong ML library support (scikit-learn, TensorFlow).

**Reasoning:**  
This showcases cross-language microservices, highlights the developer’s ability to integrate diverse technologies, and aligns each language to its strengths.

---

### 3️⃣ API Gateway Choice

✅ **Decision:** Use **NestJS (Node.js)** as the API Gateway.

- Modular, scalable architecture.
- Native TypeScript support.
- Built-in decorators, interceptors, guards, rate limiting.
- Easy integration with Swagger/OpenAPI for documentation.

---

### 4️⃣ Database Design

✅ **Decision:**  
- **PostgreSQL** for relational data (users, orders, payments).
- **MongoDB** for flexible, evolving product catalog.

**Reasoning:**  
PostgreSQL provides ACID guarantees for critical transactions, while MongoDB offers flexibility for dynamic product attributes and embedded supplier info.

---

### 5️⃣ Frontend Framework Split

✅ **Decision:**  
- **Angular (Admin Dashboard)** → structured, large-scale forms, NgRx state management.
- **React (Customer Storefront)** → flexibility, modern UI/UX, React Query or Redux for state and data fetching.

**Reasoning:**  
Angular fits enterprise admin interfaces with strict patterns; React shines in consumer-facing, highly interactive applications.

---

### 6️⃣ Communication Patterns

✅ **Decision:**  
- **REST over HTTP** for synchronous service-to-service calls.
- **RabbitMQ** for asynchronous events (order confirmations, notifications).

**Reasoning:**  
REST is simple and widely supported; RabbitMQ enables decoupled, scalable background tasks.

---

### 7️⃣ Deployment & Orchestration

✅ **Decision:**  
- **Docker** for containerization.
- **Kubernetes (minikube/k3d)** for local orchestration.
- (Optional future) **Helm** for templating.
- **Terraform** (optional) for infra-as-code.

**Reasoning:**  
Demonstrates cloud-native skills, scalable deployment patterns, and readiness for production-scale orchestration.

---

### 8️⃣ Monitoring & Logging

✅ **Decision:**  
- **Prometheus + Grafana** for metrics and dashboards.
- **ELK Stack (Elasticsearch, Logstash, Kibana)** for centralized log aggregation.

**Reasoning:**  
These are industry-standard, open-source solutions that showcase advanced monitoring and troubleshooting setups.

---

### 9️⃣ CI/CD

✅ **Decision:** Use **GitHub Actions** for continuous integration and deployment.

- Free tier available.
- Native integration with GitHub repositories.
- Automates build, test, and deploy pipelines.
- Generates build/test badges for the portfolio README.

---

## 🚀 Future Considerations

- Introduce gRPC between backend services (e.g., for ML recommender).
- Implement autoscaling policies in Kubernetes.
- Add chaos testing for resilience (Chaos Mesh, Gremlin).
- Expand gamification and eco-impact tracking modules.

---

## 📎 Summary

This project balances technical depth with practical trade-offs, allowing the developer to demonstrate expertise in:

✅ Backend and frontend development  
✅ API design and service orchestration  
✅ DevOps practices and cloud-native deployment  
✅ System scalability, observability, and maintainability

