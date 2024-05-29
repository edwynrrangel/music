variable "minio_server" {
  description = "The MinIO server endpoint"
  type        = string
  sensitive   = true
}

variable "minio_user" {
  description = "The user for MinIO"
  type        = string
  sensitive   = true
}

variable "minio_password" {
  description = "The password for MinIO"
  type        = string
  sensitive   = true
}

variable "multimedia_readonly_user_secret" {
    description = "The secret of the app readonly user"
    type        = string
    sensitive   = true
}