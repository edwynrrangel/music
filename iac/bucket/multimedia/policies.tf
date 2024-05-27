resource "minio_iam_policy" "readonly_policy" {
  provider = minio.local
  name     = "multimedia-readonly-policy"
  policy   = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:GetBucketLocation",
        "s3:ListBucket"
      ],
      "Resource": "arn:aws:s3:::multimedia"
    },
    {
      "Effect": "Allow",
      "Action": [
        "s3:GetObject"
      ],
      "Resource": "arn:aws:s3:::multimedia/*"
    }
  ]
}
EOF
}
