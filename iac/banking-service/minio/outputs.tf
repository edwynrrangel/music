output "minio_url" {
  value = "${kubernetes_stateful_set.minio.metadata[0].name}.${var.namespace}.svc.cluster.local:${var.minio_port}"
}

output "minio_port" {
  value = module.minio.minio_port
}

output "minio_ip" {
  value =  kubernetes_stateful_set.minio.status[0].load_balancer[0].ingress[0].ip
}