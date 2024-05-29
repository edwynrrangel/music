variable "namespace" {
  description = "The namespace to deploy resources"
  type        = string
}

variable "service_name" {
  description = "The name of the service"
  type        = string
  default     = "mongodb"
}

variable "labels" {
  description = "The labels to apply to the resources"
  type        = map(string)
}

variable "selector_labels" {
  description = "The labels to apply to the selector"
  type        = map(string)
  default     = {
    "app.kubernetes.io/name" = "mongodb"
    "app.kubernetes.io/instance" = "mongodb-intance-0"
  }
}

variable "db_secret_name" {
  description = "The name of the secrets to use for the database"
  type        = string
  default     = "mongodb-db-secrets"
}

variable "tls_secret_name" {
  description = "The name of the secrets to use for TLS"
  type        = string
  default     = "mongodb-tls-secrets"
}

variable "mongodb_port" {
  description = "The port to expose MongoDB"
  type        = number
  default     = 27017
}