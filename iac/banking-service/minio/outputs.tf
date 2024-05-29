output "minio_url" {
  value = "${kubernetes_service.minio_service.metadata[0].name}.${var.namespace}"
}

output "minio_port" {
  value = var.minio_port
}

output "minio_ip" {
  value = kubernetes_service.minio_service.status[0].load_balancer[0].ingress[0].ip
}