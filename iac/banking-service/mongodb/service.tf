resource "kubernetes_service" "mongodb_service" {
  metadata {
    name      = "${var.service_name}-service"
    namespace = var.namespace
    labels = var.labels
  }
  spec {
    selector = var.selector_labels
    type = "LoadBalancer"
    port {
      name        = "mongodb"
      protocol    = "TCP" 
      port        = var.mongodb_port
    }
  }
}