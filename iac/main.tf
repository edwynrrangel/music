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

# module "bucket" {
#   source                   = "./bucket"
#   minio_server             = "${module.banking_service.minio_url}:${module.banking_service.minio_port}"
#   minio_user               = var.minio_user
#   minio_password           = var.minio_password
#   app_readonly_user_secret = var.app_readonly_user_secret
# }