resource "kubernetes_stateful_set" "mongodb" {
  depends_on = [ kubernetes_service.mongodb_service ]
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
          image = "mongo:7.0.10-rc0-jammy"

          image_pull_policy = "Always"
          args = [ "mongod", "--bind_ip_all", "--tlsMode", "requireTLS", "--tlsCertificateKeyFile", "/etc/ssl/mongo/mongodb.pem", "--tlsCAFile", "/etc/ssl/mongo/ca.pem", "--tlsAllowConnectionsWithoutCertificates", "--auth"]
          env {
              name = "MONGO_INITDB_ROOT_USERNAME"
              value_from {
                  secret_key_ref {
                      name = var.db_secret_name
                      key = "MONGO_INITDB_ROOT_USERNAME"
                  }
              }
          }
          env {
            name = "MONGO_INITDB_ROOT_PASSWORD"
            value_from {
              secret_key_ref {
                name = var.db_secret_name
                key = "MONGO_INITDB_ROOT_PASSWORD"
              }
            }
          }
          port {
            container_port = var.mongodb_port
            protocol = "TCP"
          }
          volume_mount {
            name = kubernetes_persistent_volume_claim.banking_service_mongodb_pvc.metadata[0].name
            mount_path = "/data/db"
          }
          volume_mount {
            name = "${var.tls_secret_name}-vm"
            mount_path = "/etc/ssl/mongo"
            read_only = true
          }
        }
        volume {
          name = kubernetes_persistent_volume_claim.banking_service_mongodb_pvc.metadata[0].name
          persistent_volume_claim {
            claim_name = kubernetes_persistent_volume_claim.banking_service_mongodb_pvc.metadata[0].name
          }
        }
        volume {
          name = "${var.tls_secret_name}-vm"
          secret {
            secret_name = var.tls_secret_name
          }
        }
      }
    }
  }
}