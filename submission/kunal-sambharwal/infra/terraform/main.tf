resource "kubernetes_namespace" "config_service" {
  metadata {
    name = var.namespace
  }
}

resource "helm_release" "config_service" {

  name = var.release_name

  chart = var.chart_path

  namespace = kubernetes_namespace.config_service.metadata[0].name

  dependency_update = true

  create_namespace = false

  depends_on = [
    kubernetes_namespace.config_service
  ]
}