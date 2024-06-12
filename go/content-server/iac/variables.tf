variable "project_name" {
  description = "The project name"
  type        = string
  default     = "content-server"
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
    "app.kubernetes.io/name" = "multimedia-server"
    "app.kubernetes.io/instance" = "multimedia-server-0"
  }
}

variable "multimedia_server_secrets" {
  description = "The multimedia server secrets"
  type        = string
  default     = "multimedia-server-secrets"
}

variable "app_port" {
  description = "The app port"
  type        = number
  default     = 50051
}

