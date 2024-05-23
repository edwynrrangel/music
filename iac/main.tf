terraform {
  required_version = ">= 0.12"
  backend "local" {
    path = "/home/egui/terraform/grpc/infra/terraform.tfstate"
  }
}

module "banking_service" {
  source = "./banking-service"
}

module "grpc_service" {
  source = "./grpc-service"
}