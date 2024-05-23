resource "kubernetes_namespace" "grpc_service_namespace" {
  metadata {
    name = var.namespace
  }
}
