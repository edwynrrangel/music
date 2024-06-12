resource "kubernetes_service" "multimedia_client_service" {
  metadata {
    name      = "${var.project_name}-service"
    namespace = data.terraform_remote_state.infra.outputs.grpc_service_namespace
    labels = var.labels
  }
  spec {
    selector = var.selector_labels
    type = "LoadBalancer"
    port {
      name     = "http"
      protocol = "TCP" 
      port     = var.app_port
    }
  }
}