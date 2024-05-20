resource "kubernetes_role" "banking_service_mongodb_pvc_reader_role" {
  metadata {
    name      = "${var.namespace}-${var.release_name}-pvc-reader-role"
    namespace = var.namespace
  }

  rule {
    api_groups = [""]
    resources  = ["persistentvolumeclaims"]
    verbs      = ["get", "list"]
  }
}
