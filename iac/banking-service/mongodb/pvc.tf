resource "kubernetes_persistent_volume_claim" "banking_service_mongodb_pvc" {
  metadata {
    name      = "${var.namespace}-${var.release_name}-pvc"
    namespace = var.namespace
  }
  spec {
    access_modes = ["ReadWriteOnce"]
    resources {
      requests = {
        storage = "1Gi"
      }
    }
    storage_class_name = "standard"
  }
}