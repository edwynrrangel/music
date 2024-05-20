terraform {
  required_version = ">= 0.12"
}

module "banking_service" {
  source = "./banking-service"
}
