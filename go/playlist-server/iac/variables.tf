variable "project_name" {
  description = "The project name"
  type        = string
  default     = "playlist-server"
}

variable "labels" {
  description = "The labels to apply to the resources"
  type        = map(string)
  default     = {
    cluster = "grpc-service"
    "app.kubernetes.io/managed-by" = "Terraform"
  }
}

variable "selector_labels" {
  description = "The labels to apply to the selector"
  type        = map(string)
  default     = {
    "app.kubernetes.io/name" = "playlist-server"
    "app.kubernetes.io/instance" = "playlist-server-0"
  }
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

