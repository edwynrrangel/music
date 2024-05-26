resource "kubernetes_role" "banking_service_pvc_reader_role" {
  metadata {
    name      = "pvc-reader-role"
    namespace = var.namespace
  }

  rule {
    api_groups = [""]
    resources  = ["persistentvolumeclaims"]
    verbs      = ["get", "list"]
  }
}
