output "mongodb_url" {
  value = "${kubernetes_service.mongodb_service.metadata[0].name}.${var.namespace}"
}

output "mongodb_port" {
  value = var.mongodb_port
}

output "mongodb_ip" {
  value = kubernetes_service.mongodb_service.status[0].load_balancer[0].ingress[0].ip
}