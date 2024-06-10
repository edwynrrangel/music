variable "project_name" {
  description = "The project name"
  type        = string
  default     = "multimedia-app"
}

variable "cluster_name" {
  description = "The cluster name"
  type        = string
  default     = "grpc-service"
}

variable "image_pull_secret" {
  description = "The image pull secret"
  type        = string
  default     = "multimedia-app-docker-secrets"
}

variable "multimedia_app_secrets" {
  description = "The multimedia app secrets"
  type        = string
  default     = "multimedia-app-secrets"
}

variable "app_port" {
  description = "The app port"
  type        = number
  default     = 50051
}