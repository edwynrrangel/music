resource "minio_iam_user" "app_readonly_user" {
  provider      = minio.local
  name          = "multimedia-readonly-user"
  force_destroy = true
  secret        = var.multimedia_readonly_user_secret
}