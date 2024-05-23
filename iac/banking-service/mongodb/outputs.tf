output "mongodb_url" {
  value = "${helm_release.mongodb_release.name}.${var.namespace}.svc.cluster.local"
}
