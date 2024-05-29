terraform {
  required_providers {
    minio = {
      source  = "aminueza/minio"
      version = "2.2.1"
    }
  }
}

module "multimedia" {
  source    = "./multimedia"
  providers = {
    minio.local = minio.local
  }
  multimedia_readonly_user_secret = var.multimedia_readonly_user_secret
}