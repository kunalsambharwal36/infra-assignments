# Kubernetes Config Service

## Overview

This project implements a **Configuration Service** in Go and deploys it on a local Kubernetes cluster using **Helm** and **Terraform**.

The objective is to demonstrate infrastructure engineering best practices including Infrastructure as Code (IaC), Kubernetes deployment, local reproducibility, operational readiness, and deployment automation.

The application exposes REST APIs for storing and retrieving configuration records from PostgreSQL.

---

# Architecture

```
                +-------------------------+
                |     Client / curl       |
                +-----------+-------------+
                            |
                            |
                    Kubernetes Service
                            |
                            |
                    Config Service (Go)
                            |
                            |
                    PostgreSQL Database
                            |
                      Persistent Volume

Terraform
      │
      └──────────────► Helm Chart
                              │
                              ▼
                     Kubernetes Resources
```

---

# Technology Stack

| Component | Technology |
|------------|------------|
| Language | Go 1.22 |
| Database | PostgreSQL |
| Container | Docker |
| Orchestration | Kubernetes |
| Package Manager | Helm |
| Infrastructure as Code | Terraform |
| Local Cluster | Kind / Minikube |
| Validation | Helm, Terraform, Go Tests |

---

# Repository Structure

```
submission/kunal-sambharwal
├── cmd/                    # Application entry point
├── internal/
│   ├── domain/             # Models
│   ├── handler/            # HTTP handlers
│   ├── repository/         # Database layer
│   └── service/            # Business logic
├── charts/
│   └── config-service/     # Helm chart
├── infra/
│   └── terraform/          # Terraform configuration
├── scripts/
│   └── smoke-test.sh
├── Dockerfile
├── Makefile
├── README.md
├── go.mod
└── go.sum
```

---

# API Endpoints

## Health Check

```
GET /ping
```

Response

```
pong
```

Used for Kubernetes liveness and readiness validation.

---

## Get Configuration

```
GET /configs/{id}
```

Example

```
GET /configs/cfg_1
```

Response

```json
{
  "id": "cfg_1",
  "host": "localhost",
  "port": 8080,
  "app_name": "config-service",
  "log_level": "INFO"
}
```

Returns **404** if the configuration does not exist.

---

## Create / Update Configuration

```
POST /configs
```

Example

```json
{
  "id":"cfg_1",
  "host":"localhost",
  "port":8080,
  "app_name":"config-service",
  "log_level":"INFO"
}
```

The API performs an **upsert**, meaning an existing configuration is updated if the ID already exists.

---

# Database

The application uses PostgreSQL.

The database schema is initialized using SQL migrations packaged with the Helm chart.

The `configs` table stores:

- id (Primary Key)
- host
- port
- app_name
- log_level

---

# Infrastructure

Infrastructure provisioning is performed using Terraform.

Terraform is responsible for:

- Helm release deployment
- Kubernetes resources
- Namespace creation
- PostgreSQL deployment
- Config Service deployment

Application packaging is managed using Helm.

---

# Configuration Management

Application configuration is supplied through Kubernetes ConfigMaps.

Sensitive values (database credentials) are stored using Kubernetes Secrets.

For a production deployment, these would be replaced with a dedicated secrets manager such as:

- HashiCorp Vault
- AWS Secrets Manager
- External Secrets Operator

---

# Deployment

## Prerequisites

- Docker
- kubectl
- Helm
- Terraform
- Kind or Minikube

---

## Deploy

Initialize Terraform

```bash
make init
```

Validate

```bash
make validate
```

Deploy

```bash
make deploy
```

---

# Verification

Check Pods

```bash
kubectl get pods -n config-service
```

Check Services

```bash
kubectl get svc -n config-service
```

Port Forward

```bash
make port-forward
```

Health Check

```bash
curl http://localhost:8080/ping
```

---

# Validation

Go

```bash
go build ./...
go test ./...
go vet ./...
gofmt -l .
```

Terraform

```bash
terraform fmt -check -recursive
terraform validate
```

Helm

```bash
helm lint charts/config-service
```

---

# Smoke Test

Run

```bash
make test
```

or

```bash
./scripts/smoke-test.sh
```

The smoke test validates:

- Application startup
- Health endpoint
- Configuration creation
- Configuration retrieval

---

# Operational Readiness

The deployment includes:

- Kubernetes readiness probes
- Kubernetes liveness probes
- ConfigMaps
- Secrets
- Persistent Volume Claim
- Structured application startup logging

If PostgreSQL is unavailable, the application logs the connection failure and Kubernetes restart policies ensure retry behaviour.

---

# Design Decisions

- Helm was selected to simplify application packaging.
- Terraform provides Infrastructure as Code and reproducible deployments.
- Repository, Service and Handler layers are separated to keep the code modular and maintainable.
- PostgreSQL is deployed inside Kubernetes to provide a fully local reproducible environment.

---

# Production Improvements

For a production deployment the following enhancements would be implemented:

- Remote Terraform Backend
- External Secret Management
- TLS Ingress
- Horizontal Pod Autoscaler
- Prometheus Metrics
- Grafana Dashboards
- CI/CD Pipeline for image publishing
- PostgreSQL High Availability
- Automated Database Migrations
- Vulnerability Scanning during CI

---

# Cleanup

Destroy infrastructure

```bash
make destroy
```

or

```bash
kubectl delete namespace config-service
```

---

# AI Usage

AI tools (ChatGPT) were used to assist with brainstorming, documentation, and validating implementation ideas. All infrastructure, Go application code, Kubernetes manifests, Terraform configuration, and deployment workflow were reviewed, tested, and verified locally before submission.