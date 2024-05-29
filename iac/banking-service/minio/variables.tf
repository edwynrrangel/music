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
}

variable "selector_labels" {
  description = "The labels to apply to the selector"
  type        = map(string)
  default     = {
    "app.kubernetes.io/name" = "minio"
    "app.kubernetes.io/instance" = "minio-intance-0"
  }
}

variable "minio_secret_name" {
  description = "The name of the secret to store Minio credentials"
  type        = string
  default     = "minio-secrets"
}

variable "minio_port" {
  description = "The port to expose Minio"
  type        = number
  default     = 9000
}
