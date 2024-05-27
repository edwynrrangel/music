resource "minio_s3_bucket" "minio_bucket" {
  provider      = minio.local
  bucket        = "multimedia"
  acl           = "private"
  force_destroy = true
}

output "bucket_url" {
  description = "The URL of the bucket multimedia"
  value = minio_s3_bucket.minio_bucket.bucket_domain_name
}