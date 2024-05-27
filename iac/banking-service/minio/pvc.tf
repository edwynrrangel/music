resource "kubernetes_persistent_volume_claim" "banking_service_minio_pvc" {
  metadata {
    name      = "${var.service_name}-pvc"
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