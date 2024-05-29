resource "kubernetes_service" "minio_service" {
  metadata {
    name      = "${var.service_name}-service"
    namespace = var.namespace
    labels = var.labels
  }
  spec {
    selector = var.selector_labels
    type = "LoadBalancer"
    port {
      name     = "minio"
      protocol = "TCP" 
      port     = var.minio_port
    }
    port {
      name     = "minio-ui"
      protocol = "TCP" 
      port     = 9001
    }
  }
}