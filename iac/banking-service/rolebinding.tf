resource "kubernetes_role_binding" "banking_service_pvc_reader_binding" {
  metadata {
    name      = "pvc-reader-binding"
    namespace = var.namespace
  }

  subject {
    kind      = "ServiceAccount"
    name      = kubernetes_service_account.banking_service_pvc_sa.metadata[0].name
    namespace = var.namespace
  }

  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
    name      = kubernetes_role.banking_service_pvc_reader_role.metadata[0].name
  }
}
