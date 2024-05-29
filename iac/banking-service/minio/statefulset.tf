resource "kubernetes_stateful_set" "minio" {
  depends_on = [ kubernetes_service.minio_service ]
  timeouts {
    create = "2m"
    delete = "5m"
  }
  metadata {
    name      = var.service_name
    namespace = var.namespace
    labels = merge(
      {
        service = var.service_name
      },
      var.labels,
      var.selector_labels
    )
  }

  spec {
    service_name = var.service_name
    replicas     = 1
    selector {
      match_labels = var.selector_labels
    }
    template {
      metadata {
        labels = var.selector_labels
        namespace = var.namespace
      }
      spec {
        container {
          name  = var.service_name
          image = "minio/minio:latest"
          image_pull_policy = "Always"
          args = [ "server", "/data", "--console-address", ":9001"]
          env {
              name = "MINIO_ROOT_USER"
              value_from {
                  secret_key_ref {
                      name = var.minio_secret_name
                      key = "MINIO_ROOT_USER"
                  }
              }
          }
          env {
            name = "MINIO_ROOT_PASSWORD"
            value_from {
              secret_key_ref {
                name = var.minio_secret_name
                key = "MINIO_ROOT_PASSWORD"
              }
            }
          }
          port {
            container_port = var.minio_port
            protocol = "TCP"
          }
          port {
            container_port = 9001
            protocol = "TCP"
          }
          
          volume_mount {
            name = kubernetes_persistent_volume_claim.banking_service_minio_pvc.metadata[0].name
            mount_path = "/data"
          }
        }
        volume {
          name = kubernetes_persistent_volume_claim.banking_service_minio_pvc.metadata[0].name
          persistent_volume_claim {
            claim_name = kubernetes_persistent_volume_claim.banking_service_minio_pvc.metadata[0].name
          }
        }
      }
    }
  }
}