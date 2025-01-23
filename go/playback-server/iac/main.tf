terraform {
  required_version = ">= 0.12"
  backend "local" {
    path = "/home/egui/terraform/music/apps/go/playback-server/terraform.tfstate"
  }
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

data "terraform_remote_state" "infra" {
  backend = "local"
  config  = {
    path = "/home/egui/terraform/music/infra/terraform.tfstate"
  }
}

data "terraform_remote_state" "content_server" {
  backend = "local"
  config  = {
    path = "/home/egui/terraform/music/apps/go/content-server/terraform.tfstate"
  }
}