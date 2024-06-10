output "grpc_server_url" {
  value = "${helm_release.grpc_multimedia_release.metadata[0].name}.${data.terraform_remote_state.infra.outputs.grpc_service_namespace}.svc.cluster.local:${var.app_port}"
  
}