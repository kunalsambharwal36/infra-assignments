output "namespace" {
  value = kubernetes_namespace.config_service.metadata[0].name
}

output "helm_release" {
  value = helm_release.config_service.name
}