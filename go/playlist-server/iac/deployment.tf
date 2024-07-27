resource "kubernetes_deployment" "multimedia_server" {
  timeouts {
    create = "2m"
    delete = "5m"
  }

  metadata {
    name      = var.project_name
    namespace = data.terraform_remote_state.infra.outputs.grpc_service_namespace
    labels = merge(
      {
        project = var.project_name
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
            service = var.project_name
          },
          var.labels,
          var.selector_labels
        )
      }

      spec {
        container {
          name  = var.project_name
          image = "edwynrangel/multimedia-server:latest"

          port {
            container_port = var.app_port
            protocol = "TCP"
          }

          env {
            name  = "PORT"
            value = "${var.app_port}"
          }
          env {
            name  = "MONGO_HOST"
            value = data.terraform_remote_state.infra.outputs.banking_service_mongodb_url
          }
          env {
            name  = "MONGO_PORT"
            value = data.terraform_remote_state.infra.outputs.banking_service_mongodb_port
          }
          env {
            name  = "MONGO_DB_NAME"
            value = "multimedia"
          }
          env {
            name  = "MONGO_TLS"
            value = "true"
          }
          env {
            name  = "MONGO_CONTENT_COLLECTION_NAME"
            value = "contents"
          }
          env {
            name  = "MONGO_PLAYLIST_COLLECTION_NAME"
            value = "playlists"
          }
          env {
            name  = "MINIO_ENDPOINT"
            value = "${data.terraform_remote_state.infra.outputs.banking_service_minio_url}:${data.terraform_remote_state.infra.outputs.banking_service_minio_port}"
          }
          env {
            name  = "MINIO_SSL"
            value = "false"
          }
          env {
            name  = "BUCKET_TYPE"
            value = "minio"
          }
          env {
            name = "MONGO_USERNAME"
            value_from {
              secret_key_ref {
                name = var.multimedia_app_secrets
                key  = "MONGO_USERNAME"
              }
            }
          }
          env {
            name = "MONGO_PASSWORD"
            value_from {
              secret_key_ref {
                name = var.multimedia_app_secrets
                key  = "MONGO_PASSWORD"
              }
            }
          }
          env {
            name = "MONGO_TLS_CA"
            value_from {
              secret_key_ref {
                name = var.multimedia_app_secrets
                key  = "MONGO_TLS_CA"
              }
            }
          }
          env {
            name = "MINIO_READONLY_ACCESS_KEY"
            value_from {
              secret_key_ref {
                name = var.multimedia_app_secrets
                key  = "MINIO_READONLY_ACCESS_KEY"
              }
            }
          }
          env {
            name = "MINIO_READONLY_SECRET_KEY"
            value_from {
              secret_key_ref {
                name = var.multimedia_app_secrets
                key  = "MINIO_READONLY_SECRET_KEY"
              }
            }
          }
          
          security_context {
            run_as_user = 1000
            run_as_group = 1000
          }

          resources {
            requests = {
              cpu    = "75m"
              memory = "200Mi"
            }

            limits = {
              cpu    = "100m"
              memory = "384Mi"
            }
          }
        }
      }
    }
  }
}
