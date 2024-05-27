resource "kubernetes_service" "banking_service_minio_service" {
  metadata {
    name      = "${var.service_name}-service"
    namespace = var.namespace
    labels = var.labels
  }
  spec {
    selector = var.selector_labels
    type = "LoadBalancer"
    port {
      name        = "minio"
      protocol = "TCP" 
      port        = 9000
      target_port = 9000
    }
    port {
      name        = "minio-ui"
      protocol = "TCP" 
      port        = 9001
      target_port = 9001
    }
  }
}