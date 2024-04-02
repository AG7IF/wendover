resource "aws_ssm_parameter" "wendover-aws-cognito-iss" {
  name  = "/wendover/aws/cognito/iss"
  type  = "String"
  value = "https://cognito-idp.${var.region}.amazonaws.com/${aws_cognito_user_pool.wendover.id}"
}

resource "aws_ssm_parameter" "wendover-aws-cognito-userpool_id" {
  name  = "/wendover/aws/cognito/userpool_id"
  type  = "String"
  value = aws_cognito_user_pool.wendover.id
}

resource "aws_ssm_parameter" "wendover-aws-log_group_name" {
  name  = "/wendover/aws/log_group_name"
  type  = "String"
  value = aws_cloudwatch_log_group.wendover.name
}

resource "aws_ssm_parameter" "wendover-aws-log_stream_name" {
  name  = "/wendover/aws/log_stream_name"
  type  = "String"
  value = aws_cloudwatch_log_stream.wendover_api.name
}

resource "aws_ssm_parameter" "wendover-config-directory" {
  name  = "/wendover/config/directory"
  type  = "String"
  value = "DOZFAC"
}

resource "aws_ssm_parameter" "wendover-database-host" {
  name  = "/wendover/database/host"
  type  = "String"
  value = aws_db_instance.wendover.address
}

resource "aws_ssm_parameter" "wendover-database-port" {
  name  = "/wendover/database/port"
  type  = "String"
  value = aws_db_instance.wendover.port
}

resource "aws_ssm_parameter" "wendover-database-name" {
  name  = "/wendover/database/name"
  type  = "String"
  value = aws_db_instance.wendover.db_name
}

resource "aws_ssm_parameter" "wendover-database-ssl" {
  name  = "/wendover/database/ssl"
  type  = "String"
  value = "true"
}

resource "aws_ssm_parameter" "wendover-database-migration-source" {
  name  = "/wendover/database/migration/source"
  type  = "String"
  value = "s3://${aws_s3_bucket.wendover_db_migration.bucket}"
}

resource "aws_ssm_parameter" "wendover-server-run_address" {
  name  = "/wendover/server/run_address"
  type  = "String"
  value = var.api_run_address
}

resource "aws_ssm_parameter" "wendover-server-root_path" {
  name  = "/wendover/server/root_path"
  type  = "String"
  value = var.api_root_path
}
