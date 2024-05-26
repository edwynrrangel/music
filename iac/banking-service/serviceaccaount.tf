resource "kubernetes_service_account" "banking_service_pvc_sa" {
  metadata {
    name      = "pvc-access-sa"
    namespace = var.namespace
  }
}
