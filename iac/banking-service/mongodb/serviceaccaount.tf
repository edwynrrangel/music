resource "kubernetes_service_account" "banking_service_mongodb_pvc_access_sa" {
  metadata {
    name      = "${var.release_name}-pvc-access-sa"
    namespace = var.namespace
  }
}
