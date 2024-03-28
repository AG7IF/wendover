resource "aws_kms_key" "wendover" {
  description = "Wendover secrets key"
  policy = <<-EOT
  {
      "Version": "2012-10-17",
      "Id": "wendover-secrets-1",
      "Statement": [
          {
              "Sid": "Enable IAM User Permissions",
              "Effect": "Allow",
              "Principal": {
                  "AWS": "${var.aws_principal_arn}"
              },
              "Action": "kms:*",
              "Resource": "*"
          },
          {
              "Sid": "Allow Terraform Cloud access to this key",
              "Effect": "Allow",
              "Principal": {
                  "AWS": "${var.terraform_cloud_role_arn}"
              },
              "Action": "kms:*",
              "Resource": "*"
          }
      ]
  }
EOT

}

resource "aws_kms_alias" "wendover" {
  name          = "alias/wendover"
  target_key_id = aws_kms_key.wendover.key_id
}

resource "aws_db_subnet_group" "wendover" {
  name        =   "wendover-db"
  subnet_ids  =   [aws_subnet.wendover_db_aza.id, aws_subnet.wendover_db_azb.id]
}

resource "aws_security_group" "wendover-db" {
  name    =   "wendover-db"
  vpc_id  =   aws_vpc.wendover.id

  ingress {
    from_port   =   5432
    to_port     =   5432
    protocol    =   "tcp"
    cidr_blocks =   ["0.0.0.0/0"]
  }

  ingress {
    from_port   =   5432
    to_port     =   5432
    protocol    =   "tcp"
    cidr_blocks =   ["0.0.0.0/0"]
  }
}

resource "aws_db_parameter_group" "wendover" {
  name = "wendover"
  family = "postgres15"
}

resource "aws_db_instance" "wendover" {
  identifier                            =   "wendover-db"
  instance_class                        =   "db.t3.micro"
  allocated_storage                     =   20
  parameter_group_name                  =   aws_db_parameter_group.wendover.name
  blue_green_update {
    enabled = true
  }

  engine                                =   "postgres"
  engine_version                        =   "15"

  username                              =   "postgres"
  manage_master_user_password           =   true
  master_user_secret_kms_key_id         =   aws_kms_key.wendover.id
  iam_database_authentication_enabled   =   true

  db_name                               =   "wendover"
  storage_encrypted                     =   true

  db_subnet_group_name                  =   aws_db_subnet_group.wendover.name
  vpc_security_group_ids                =   [aws_security_group.wendover-db.id]

  skip_final_snapshot                   =   true
  apply_immediately                     =   true
}


resource "aws_s3_bucket" "wendover_db_migration" {
  bucket = "wendover-migrations"

}

resource "aws_iam_role" "wendover_db_migration" {
  name = "WendoverDBMigrationRole"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
  managed_policy_arns = [aws_iam_policy.wendover_lambda_role_policy.arn]
}

resource "aws_lambda_function" "wendover_db_migration" {
  function_name   =   "wendover-migrate-db"
  package_type    =   "Image"
  image_uri       =   "712249788489.dkr.ecr.us-west-2.amazonaws.com/ag7if:wendsrv-latest"
  role            =   aws_iam_role.wendover_db_migration.arn
  image_config {
    entry_point =   ["./wendsrv-migrate-lambda"]
  }
}