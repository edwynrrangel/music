variable "namespace" {
  description = "The namespace to deploy resources"
  type        = string
}

variable "service_name" {
  description = "The release name of the application"
  type        = string
  default     = "minio"
}

variable "labels" {
  description = "The labels to apply to the resources"
  type        = map(string)
  default     = {
    cluster = "banking-service"
    "app.kubernetes.io/managed-by" = "Terraform"
  }
}

variable "selector_labels" {
  description = "The labels to apply to the selector"
  type        = map(string)
  default     = {
    "app.kubernetes.io/name" = "minio"
    "app.kubernetes.io/instance" = "minio-intance-1"
  }
}

variable "minio_secret_name" {
  description = "The name of the secret to store Minio credentials"
  type        = string
  default     = "minio-secrets"
}
