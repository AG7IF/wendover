resource "aws_kms_key" "wendover" {
  description = "Wendover secrets key"
  policy = jsonencode({
      "Version"   = "2012-10-17",
      "Id"        = "wendover-secrets-1",
      "Statement" = [
          {
              "Sid"       = "Enable IAM User Permissions",
              "Effect"    = "Allow",
              "Principal" = {
                  "AWS"   = var.aws_kms_principal_arn
              },
              "Action"    = "kms:*",
              "Resource"  = "*"
          },
          {
              "Sid"       = "Allow Terraform Cloud access to this key",
              "Effect"    = "Allow",
              "Principal" = {
                  "AWS"   = var.terraform_cloud_role_arn
              },
              "Action"    = "kms:*",
              "Resource"  = "*"
          }
      ]
  })
}

resource "aws_kms_alias" "wendover" {
  name          = "alias/wendover"
  target_key_id = aws_kms_key.wendover.key_id
}

resource "aws_db_parameter_group" "wendover" {
  name = "wendover"
  family = "postgres15"
}

resource "aws_db_instance" "wendover" {
  identifier                            = "wendover-db"
  instance_class                        = "db.t3.micro"
  availability_zone                     = "${var.region}a"
  allocated_storage                     = 20
  parameter_group_name                  = aws_db_parameter_group.wendover.name
  blue_green_update {
    enabled = true
  }

  engine                                = "postgres"
  engine_version                        = "15"

  username                              = "postgres"
  manage_master_user_password           = true
  master_user_secret_kms_key_id         = aws_kms_key.wendover.id

  db_name                               = "wendover"
  storage_encrypted                     = true

  db_subnet_group_name                  = aws_db_subnet_group.wendover.name
  vpc_security_group_ids                = [aws_security_group.wendover_db.id]

  skip_final_snapshot                   = true
  apply_immediately                     = true
}

resource "aws_s3_bucket" "wendover_db_migration" {
  bucket = "wendover-migrations"
}
