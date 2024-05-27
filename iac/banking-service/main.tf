resource "kubernetes_namespace" "banking_service_namespace" {
  metadata {
    name = var.namespace
  }
}

module "mongodb" {
  source = "./mongodb"
  depends_on = [ kubernetes_namespace.banking_service_namespace ]
  namespace = kubernetes_namespace.banking_service_namespace.metadata[0].name
  service_account_name = kubernetes_service_account.banking_service_pvc_sa.metadata[0].name
}

module "minio" {
  source = "./minio"
  depends_on = [ kubernetes_namespace.banking_service_namespace ]
  namespace = kubernetes_namespace.banking_service_namespace.metadata[0].name
}