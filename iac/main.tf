terraform {
  required_version = ">= 0.12"
  backend "local" {
    path = "/home/egui/Dev/grpc/iac/terraform.tfstate"
  }
}

module "banking_service" {
  source = "./banking-service"
}
