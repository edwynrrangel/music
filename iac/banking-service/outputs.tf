output "namespace" {
  value = var.namespace
}

output "mongodb_url" {
  value = module.mongodb.mongodb_url
}

output "mongodb_port" {
  value = module.mongodb.mongodb_port
}

output "mongodb_ip" {
  value = module.mongodb.mongodb_ip
}

output "minio_url" {
  value = module.minio.minio_url
}

output "minio_port" {
  value = module.minio.minio_port
}

output "minio_ip" {
  value = module.minio.minio_ip
}