variable "namespace" {
  description = "The namespace to deploy resources"
  type        = string
  default     = "banking-service"
}

variable "labels" {
  description = "The labels to apply to the resources"
  type        = map(string)
  default     = {
    cluster = "banking-service"
    "app.kubernetes.io/managed-by" = "Terraform"
  }
}