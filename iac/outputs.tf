output "banking_service_mongodb_url" {
    value = module.banking_service.mongodb_url
}

output "grpc_service_namespace" {
    value = module.grpc_service.namespace
}

output "minio_bucket_multimedia_url" {
    value = module.bucket.multimedia_bucket_url
}