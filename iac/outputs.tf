output "banking_service_mongodb_url" {
    value = module.banking_service.mongodb_url
}

output "banking_service_mongodb_port" {
    value = module.banking_service.mongodb_port
}

output "banking_service_mongodb_ip" {
    value = module.banking_service.mongodb_ip
}

output "banking_service_minio_url" {
    value = module.banking_service.minio_url
}

output "banking_service_minio_port" {
    value = module.banking_service.minio_port
}

output "banking_service_minio_ip" {
    value = module.banking_service.minio_ip
}

output "grpc_service_namespace" {
    value = module.grpc_service.namespace
}
