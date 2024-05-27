resource "minio_iam_group" "readonly_group" {
  provider      = minio.local
  name          = "multimedia-readonly-group"
  force_destroy = true
}

resource "minio_iam_group_policy_attachment" "readonly_group_attachment" {
  provider = minio.local
  group_name = minio_iam_group.readonly_group.name
  policy_name = minio_iam_policy.readonly_policy.name
}

resource "minio_iam_group_user_attachment" "readonly_group_user_attachment" {
  provider = minio.local
  group_name = minio_iam_group.readonly_group.name
  user_name = minio_iam_user.app_readonly_user.name
}