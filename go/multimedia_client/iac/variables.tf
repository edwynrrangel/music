variable "project_name" {
  description = "The project name"
  type        = string
  default     = "multimedia-client"
}

variable "cluster_name" {
  description = "The cluster name"
  type        = string
  default     = "grpc-service"
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
    "app.kubernetes.io/name" = "multimedia-client"
    "app.kubernetes.io/instance" = "multimedia-client-0"
  }
}

variable "app_port" {
  description = "The port the application listens on"
  type        = number
  default     = 3000
}