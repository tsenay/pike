resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "s3:ListBucket",

          "cloudfront:ListTagsForResource",
          "cloudfront:GetDistribution",
          "ec2:DescribeAccountAttributes",
          "cloudfront:CreateDistribution",
          "cloudfront:TagResource",
          "cloudfront:UntagResource",
          "cloudfront:DeleteDistribution",
          "cloudfront:UpdateDistribution",

          "cloudfront:CreatePublicKey",
          "cloudfront:GetPublicKey",
          "cloudfront:DeletePublicKey",
          "cloudfront:UpdatePublicKey",

          #          aws_cloudfront_field_level_encryption_profile
          "cloudfront:CreateFieldLevelEncryptionProfile",
          "cloudfront:GetFieldLevelEncryptionProfile",
          "cloudfront:DeleteFieldLevelEncryptionProfile",
          "cloudfront:UpdateFieldLevelEncryptionProfile",
          #config
          "cloudfront:CreateFieldLevelEncryptionConfig",
          "cloudfront:GetFieldLevelEncryptionConfig",
          "cloudfront:DeleteFieldLevelEncryptionConfig",
          "cloudfront:UpdateFieldLevelEncryptionConfig",

          #          keygroup
          "cloudfront:CreateKeyGroup",
          "cloudfront:GetKeyGroup",
          "cloudfront:DeleteKeyGroup",
          "cloudfront:UpdateKeyGroup",

          "cloudfront:CreateMonitoringSubscription",
          "cloudfront:GetMonitoringSubscription",
          "cloudfront:DeleteMonitoringSubscription"
        ]
        "Resource" : "*"
      }
    ]
  })
  tags = {
    pike      = "permissions"
    createdby = "JamesWoolfenden"
  }
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  # checkov:skip=CKV_AWS_40: By design
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}
