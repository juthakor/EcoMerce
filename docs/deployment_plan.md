# 🚀 Deployment Scope – Ethical E-Commerce Platform

This document explains what parts of the system are live-deployed for demo purposes, what parts are implemented but not deployed, and the reasoning behind those decisions.

---

## ✅ What’s Deployed (Live Demo)

| Layer        | Details                                      |
|--------------|----------------------------------------------|
| Frontend     | React Storefront → Deployed on Netlify       |
| Admin Frontend | Angular Dashboard → Deployed on Vercel    |
| Backend      | Auth Service (Go), Product Service (Node.js) → Deployed on Railway / Render |
| Database     | PostgreSQL (Railway), MongoDB Atlas (Free Tier) |
| Core API Flow| Register, login, product browsing, mock checkout |

---

## ⚠ What’s Not Deployed (Implemented but Offline)

| Component              | Reason for Not Deploying                |
|------------------------|-----------------------------------------|
| Order Service (Go)     | Not critical for demo; logic mocked     |
| ML Recommender (Python)| Not needed in MVP; resource-heavy       |
| RabbitMQ Queue         | Asynchronous flow skipped for simplicity|
| Kubernetes (K8s)       | Used locally only (minikube); not cloud |
| Monitoring (Prometheus, Grafana) | Present in dev only; not live in demo |
| Logging (ELK Stack)    | Local logging only; centralized stack skipped |
| Terraform / Helm       | Infra-as-code present but not applied   |

---

## 🧠 Why This Scope?

- **Free-tier optimized**: All live services run on Netlify, Vercel, Render, Railway, or Atlas Free Tier
- **Focus on user journey**: Frontend + basic API flow works as expected
- **Minimal infrastructure overhead**: Reduces setup complexity for reviewer or recruiter
- **Security & performance concerns**: Avoid exposing unfinished or risky services

---

## 🛡️ How We Handle Undeployed Components

| Component        | Handling Strategy                         |
|------------------|--------------------------------------------|
| Missing Services | Gateway routes return mock or stub data    |
| Async Features   | Converted to direct sync call or skipped   |
| Secrets          | All live secrets injected via environment  |
| Documentation    | Clearly marked in README + `/docs/`        |

---

## 📝 Notes for Reviewers

This project is designed for portfolio demonstration, not full-scale production.  
It balances **technical depth** with **practical deployability**.  
Services not deployed are still fully implemented and documented, with testing and integration in local dev environment.

For deeper technical documentation, see:

- [Architecture Overview](./architecture.md)
- [Security Overview](./security.md)
- [Testing Plan](./testing.md)
