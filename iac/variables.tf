variable "minio_user" {
  description = "The user for MinIO"
  type        = string
  sensitive   = false
}

variable "minio_password" {
  description = "The password for MinIO"
  type        = string
  sensitive   = true
}