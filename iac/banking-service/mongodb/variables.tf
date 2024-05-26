variable "namespace" {
  description = "The namespace to deploy resources"
  type        = string
}

variable "service_account_name" {
  description = "The name of the service account to use"
  type        = string
}

variable "release_name" {
  description = "The release name of the application"
  type        = string
  default     = "mongodb"
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