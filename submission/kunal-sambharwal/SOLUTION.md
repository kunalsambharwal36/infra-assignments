# Kubernetes Config Service – Solution Documentation

## Solution Overview

This project implements a **Kubernetes-based Configuration Service** written in Go. The application exposes REST APIs for creating, updating, and retrieving configuration records stored in PostgreSQL.

The solution is designed with a production-oriented mindset by separating application logic, infrastructure provisioning, deployment automation, configuration management, and operational concerns.

The application is containerized using Docker, packaged using Helm, and deployed to a local Kubernetes cluster. Infrastructure provisioning is managed through Terraform.

---

# Architecture

```
                           Client
                              │
                              ▼
                    Kubernetes Service
                              │
                              ▼
                    Config Service (Go)
                              │
                              ▼
                        PostgreSQL
                              │
                    Persistent Volume Claim

Terraform
      │
      ▼
Helm Release
      │
      ▼
Kubernetes Resources
```

---

# Technology Stack

| Component | Technology |
|------------|------------|
| Language | Go 1.22 |
| Database | PostgreSQL 16 |
| Container | Docker |
| Container Orchestration | Kubernetes |
| Package Manager | Helm |
| Infrastructure as Code | Terraform |
| Automation | Makefile |
| Validation | Smoke Test Script |

---

# Repository Structure

```
infra-assignments/

├── submission/
│   └── kunal-sambharwal/
│       ├── cmd/
│       ├── internal/
│       │   ├── domain/
│       │   ├── handler/
│       │   ├── repository/
│       │   └── service/
│       ├── charts/
│       │   └── config-service/
│       ├── infra/
│       │   └── terraform/
│       ├── scripts/
│       ├── Dockerfile
│       ├── Makefile
│       ├── README.md
│       ├── go.mod
│       └── go.sum
│
├── INFRA_ASSIGNMENT.md
├── README.md
└── .github/
```

---

# Infrastructure Design

Infrastructure provisioning is managed using **Terraform**, while application packaging and deployment are handled through **Helm**.

Terraform is responsible for:

- Initializing Kubernetes infrastructure
- Deploying the Helm release
- Managing the infrastructure lifecycle

Helm is responsible for rendering and deploying Kubernetes resources, including:

- Namespace
- Deployments
- Services
- ConfigMaps
- Secrets
- PersistentVolumeClaim
- PostgreSQL
- Config Service

This separation keeps infrastructure lifecycle management independent from application packaging.

---

# Kubernetes Deployment

The application is deployed using a Helm chart.

Resources deployed include:

- Namespace
- Config Service Deployment
- Config Service Service
- PostgreSQL Deployment
- PostgreSQL Service
- ConfigMap
- Secret
- PersistentVolumeClaim

---

# Database Design

The application uses PostgreSQL.

The database contains a single table:

**configs**

Columns:

- id (Primary Key)
- host
- port
- app_name
- log_level

The schema is automatically initialized using SQL migration scripts packaged with the Helm chart.

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

Used for Kubernetes liveness and readiness probes.

---

## Create / Update Configuration

```
POST /configs
```

Example Request

```json
{
  "id": "cfg1",
  "host": "localhost",
  "port": 8080,
  "app_name": "config-service",
  "log_level": "INFO"
}
```

Behavior

- Creates a new configuration if the ID does not exist.
- Updates the existing configuration when the ID already exists (upsert).

---

## Retrieve Configuration

```
GET /configs/{id}
```

Returns the configuration associated with the supplied identifier.

Returns **HTTP 404** if the configuration does not exist.

---

# Configuration Management

Application configuration is externalized using Kubernetes ConfigMaps.

Examples include:

- Database Host
- Database Port
- Database Name

This keeps configuration separate from application code.

---

# Secret Management

Sensitive information is stored using Kubernetes Secrets.

Secrets include:

- PostgreSQL Username
- PostgreSQL Password

For production deployments these should be replaced with a dedicated secret management solution such as:

- AWS Secrets Manager
- HashiCorp Vault
- External Secrets Operator

---

# Health Checks

The application exposes:

```
GET /ping
```

This endpoint is used for:

- Liveness Probe
- Readiness Probe

Kubernetes automatically restarts unhealthy containers when liveness checks fail.

---

# Deployment Automation

Infrastructure initialization

```bash
make init
```

Terraform validation

```bash
make validate
```

Deployment

```bash
make deploy
```

Smoke Test

```bash
make test
```

Cleanup

```bash
make destroy
```

---

# Validation

The following validations were successfully performed.

## Application

```bash
go build ./...
go test ./... -count=1 -race
go vet ./...
gofmt -l .
```

## Docker

```bash
docker build -t config-service:pr .
```

## Terraform

```bash
terraform fmt -check -recursive
terraform init -backend=false
terraform validate
```

## Helm

```bash
helm lint charts/config-service
```

## Kubernetes

```bash
kubectl get all -n config-service
```

Application verification included:

- Health endpoint
- Create configuration
- Retrieve configuration
- PostgreSQL connectivity

---

# Observability

Current implementation includes:

- Application startup logs
- Kubernetes liveness probe
- Kubernetes readiness probe
- Smoke test script

Future improvements:

- Prometheus metrics
- Grafana dashboards
- Structured JSON logging
- Distributed tracing (OpenTelemetry)

---

# Operational Readiness

The application is designed with basic operational considerations.

- Configuration is externalized through ConfigMaps.
- Sensitive information is stored in Kubernetes Secrets.
- Persistent data is stored using PersistentVolumeClaims.
- Infrastructure deployment is repeatable using Terraform.
- Application deployment is repeatable using Helm.
- Health is verified using Kubernetes readiness and liveness probes.

If PostgreSQL is unavailable during startup, the application logs the connection failure and Kubernetes restart policies allow the pod to retry until the database becomes available.

---

# Production Improvements

For a production deployment the following improvements would be implemented:

- Amazon RDS instead of in-cluster PostgreSQL
- Remote Terraform Backend
- AWS Secrets Manager
- External Secrets Operator
- TLS Ingress
- Horizontal Pod Autoscaler
- Network Policies
- Pod Security Standards
- Prometheus Metrics
- Grafana Dashboards
- Centralized Logging
- Automated Database Migrations
- CI/CD deployment pipeline
- Image vulnerability scanning

---

# Testing

Smoke Test

```bash
make test
```

Infrastructure validation

```bash
terraform validate
```

Helm validation

```bash
helm lint charts/config-service
```

Deployment validation

```bash
kubectl get all -n config-service
```

---

# Assumptions

- Docker is installed.
- Kubernetes cluster (Kind or Minikube) is available.
- Helm is installed.
- Terraform is installed.
- kubectl is configured.

---

# Known Limitations

This implementation is intended for local development.

Current limitations include:

- PostgreSQL runs inside the Kubernetes cluster.
- No authentication or authorization layer.
- No TLS termination.
- No remote Terraform backend.
- No automatic deployment pipeline.
- No Prometheus metrics endpoint.

---

# Responsible AI Usage

AI tools (ChatGPT) were used to assist with:

- Reviewing documentation
- Brainstorming infrastructure design
- Improving project documentation
- Validating implementation ideas

All Go application code, Terraform configuration, Helm templates, Kubernetes manifests, deployment automation, and validation steps were manually reviewed, tested, and verified before submission.

---

# Conclusion

This project demonstrates:

- Infrastructure as Code using Terraform
- Kubernetes deployment using Helm
- Docker containerization
- PostgreSQL integration
- Configuration and Secret management
- Health checks
- Deployment automation
- Operational readiness
- Local reproducibility
- Smoke testing

The solution is fully reproducible in a local Kubernetes environment and follows common cloud-native engineering practices while remaining simple and maintainable.