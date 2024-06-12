output "content_server_url" {
  value = "${kubernetes_service.content_server_service.metadata[0].name}.${data.terraform_remote_state.infra.outputs.grpc_service_namespace}.svc.cluster.local:${var.app_port}"
}