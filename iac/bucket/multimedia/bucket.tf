resource "minio_s3_bucket" "minio_bucket" {
  provider      = minio.local
  bucket        = "multimedia"
  acl           = "private"
  force_destroy = true
}
