resource "kubernetes_deployment" "multimedia_client" {
  timeouts {
    create = "2m"
    delete = "5m"
  }

  metadata {
    name      = var.project_name
    namespace = data.terraform_remote_state.infra.outputs.grpc_service_namespace
    labels = merge(
      {
        service = var.project_name
      },
      var.labels,
      var.selector_labels
    )
  }

  spec {
    replicas = 1

    selector {
      match_labels = var.selector_labels
    }

    template {
      metadata {
        namespace = data.terraform_remote_state.infra.outputs.grpc_service_namespace
        labels    = merge(
          {
            project = var.project_name
          },
          var.labels,
          var.selector_labels
        )
      }

      spec {
        container {
          name              = var.project_name
          image             = "edwynrangel/multimedia-client:latest"
          image_pull_policy = "Always"

          port {
            container_port = var.app_port
            protocol       = "TCP"
          }

          env {
            name  = "PORT"
            value = var.app_port
          }
          env {
            name  = "CONTENT_SERVER_URI"
            value = data.terraform_remote_state.server.outputs.content_server_url
          }

          security_context {
            run_as_user = 1000
            run_as_group = 1000
          }

          resources {
            requests = {
              cpu    = "75m"
              memory = "128Mi"
            }
            limits = {
              cpu    = "100m"
              memory = "256Mi"
            }
          }
        }
      }
    }
  }
}
