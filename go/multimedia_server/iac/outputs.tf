output "grpc_server_url" {
  value = "${kubernetes_deployment.multimedia_server.metadata[0].name}.${data.terraform_remote_state.infra.outputs.grpc_service_namespace}.svc.cluster.local:${var.app_port}"
}