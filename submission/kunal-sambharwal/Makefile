.PHONY: help init fmt validate plan deploy destroy lint template status logs port-forward test clean

help:
	@echo "Available Commands:"
	@echo " make init          - Initialize Terraform"
	@echo " make fmt           - Format Terraform files"
	@echo " make validate      - Validate Terraform"
	@echo " make plan          - Terraform Plan"
	@echo " make deploy        - Deploy Infrastructure"
	@echo " make destroy       - Destroy Infrastructure"
	@echo " make lint          - Helm Lint"
	@echo " make template      - Render Helm Templates"
	@echo " make status        - Show Kubernetes Resources"
	@echo " make logs          - Show Application Logs"
	@echo " make port-forward  - Port Forward Application"
	@echo " make test          - Run Smoke Test"
	@echo " make clean         - Cleanup Namespace"

init:
	cd infra/terraform && terraform init

fmt:
	cd infra/terraform && terraform fmt

validate:
	cd infra/terraform && terraform validate

plan:
	cd infra/terraform && terraform plan

deploy:
	cd infra/terraform && terraform apply -auto-approve

destroy:
	cd infra/terraform && terraform destroy -auto-approve

lint:
	helm lint charts/config-service

template:
	helm template config-service charts/config-service

status:
	kubectl get all -n config-service

logs:
	kubectl logs deployment/config-service -n config-service

port-forward:
	kubectl port-forward svc/config-service 8080:8080 -n config-service

test:
	./scripts/smoke-test.sh

clean:
	kubectl delete namespace config-service --ignore-not-found=true