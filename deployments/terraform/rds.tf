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

resource "aws_db_instance" "wendover" {
  allocated_storage                     =   20
  ca_cert_identifier                    =   "rds-ca-ecc384-g1"
  db_name                               =   "wendover"
  engine                                =   "postgres"
  engine_version                        =   "15.4"
  identifier                            =   "wendover-db"
  instance_class                        =   "db.t3.micro"
  manage_master_user_password           =   true
  master_user_secret_kms_key_id         =   aws_kms_key.wendover.id
  username                              =   "postgres"
  storage_encrypted                     =   true
  iam_database_authentication_enabled   =   true

  tags = {
    Service = "wendover"
  }
}


/*
resource "aws_db_subnet_group" "wendover" {
  name       = "wendover-dbsg"
  subnet_ids = [aws_subnet.wendover_db_aza.id, aws_subnet.wendover_db_azb.id]

  tags = {
    Service = "wendover"
  }
}
*/