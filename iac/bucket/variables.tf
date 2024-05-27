variable "minio_server" {
  description = "The MinIO server endpoint"
  default     = "10.99.177.160:9000"
}

variable "minio_user" {
  description = "The user for MinIO"
}

variable "minio_password" {
  description = "The password for MinIO"
}