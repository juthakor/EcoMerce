# 🧪 Testing Strategy – Ethical E-Commerce Platform

This document outlines the testing approach used across the system, covering:

- Unit testing
- Integration testing
- End-to-End (E2E) testing
- Load & stress testing
- CI/CD integration
- Known limitations

---

## ✅ What’s Tested (Implemented)

### 1. Unit Testing

| Area            | Tool                | Status   |
|-----------------|---------------------|----------|
| Go services     | `testing` pkg       | ✅ Done (auth) |
| Node.js services| Jest                | ✅ Done (product) |
| React frontend  | React Testing Library | ✅ Basic snapshot/component |
| Angular admin   | Karma + Jasmine     | ✅ Scaffolded |

All logic is modularized and tested with local mocks or test doubles. Auth uses `httptest` and mock JWT.

---

### 2. Integration Testing

| Scope              | Tool              | Status |
|--------------------|-------------------|--------|
| Auth + DB (Postgres) | pgx + Docker test DB | ✅ Local only |
| Product + Mongo     | Jest + Supertest | ✅ Partial |
| API Gateway routing | E2E smoke test    | ✅ Basic check |

> Note: Integration tests are local-only due to DB reliance and free-tier hosting limitations.

---

### 3. End-to-End (E2E) Testing

| Scope          | Tool      | Status |
|----------------|-----------|--------|
| User flow (React) | Cypress  | ✅ Sample test exists |
| Admin flow (Angular) | ❌ Planned | Needs auth token + backend integration |

E2E tests run in dev environment against full local system via Docker Compose.

---

## 🧪 Load & Stress Testing

### Tool: [k6](https://k6.io/) + [Vegeta](https://github.com/tsenart/vegeta)

| Endpoint          | Scenario            | Tool  | Status |
|-------------------|---------------------|-------|--------|
| `GET /products`   | 100 RPS for 30s     | k6    | ✅ Local test only |
| `POST /login`     | 50 RPS burst attack | Vegeta| ✅ Simulated |
| `POST /orders`    | Simulated flood     | ❌ Skipped (no real service) |

### 📄 Results:

- p95 latency under 200ms (local)
- 0% error under 200 RPS for product endpoint
- Auth login survived burst test but needs rate-limit middleware (planned)

> See test script in: `/tests/load/`  
> See raw results in: `/docs/load_test_results.md`

---

## ⚠ Limitations

- Load & integration tests **not runnable on public deploy**, only local (due to DB and resource constraints)
- Monitoring & tracing not active in demo (no latency chart)
- Admin E2E tests require backend token + test data mocking

---

## 🚀 Test Workflow

| Layer     | Run Command                       |
|-----------|----------------------------------|
| Go        | `go test ./...`                  |
| Node.js   | `npm run test`                   |
| React     | `npm test`                       |
| Cypress   | `npx cypress open`               |
| k6        | `k6 run load_test.js`            |

### 💡 CI/CD Integration
- GitHub Actions runs unit tests on each push
- Future: Add test coverage badge, Trivy scan, and ESLint validation

---

## 📂 File Structure

```plaintext
/tests
  /unit
    auth_test.go
    product.test.js
  /integration
    auth_db_test.go
    product_mongo.test.js
  /e2e
    cypress/
  /load
    k6_product.js
/docs
  load_test_results.md
  testing.md
