data "aws_iam_policy_document" "wendover_ecs_execution_role_trust" {
  statement {
    sid     = ""

    effect  = "Allow"
    principals {
      identifiers = ["ecs-tasks.amazonaws.com"]
      type        = "Service"
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "wendover_ecs_execution_role" {
  name                = "WendoverECSExecutionRole"
  assume_role_policy  = data.aws_iam_policy_document.wendover_ecs_execution_role_trust.json
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"]
}

data "aws_iam_policy_document" "wendover_ecs_task_role" {
  statement {
    sid     = "SecretsManagerAccess"
    effect  = "Allow"
    actions = [
      "secretsmanager:GetSecretValue",
      "kms:Decrypt"
    ]
    resources = [
      aws_db_instance.wendover.master_user_secret[0].kms_key_id,
      aws_db_instance.wendover.master_user_secret[0].secret_arn
    ]
  }

  statement {
    sid     = "ParamStoreAccess"
    effect  = "Allow"
    actions = [
      "ssm:GetParameters",
    ]
    resources = [
      aws_ssm_parameter.wendover-aws-cognito-iss.arn,
      aws_ssm_parameter.wendover-aws-cognito-userpool_id.arn,
      aws_ssm_parameter.wendover-aws-log_group_name.arn,
      aws_ssm_parameter.wendover-aws-log_stream_name.arn,
      aws_ssm_parameter.wendover-config-directory.arn,
      aws_ssm_parameter.wendover-database-host.arn,
      aws_ssm_parameter.wendover-database-port.arn,
      aws_ssm_parameter.wendover-database-name.arn,
      aws_ssm_parameter.wendover-database-ssl.arn,
      aws_ssm_parameter.wendover-database-migration-source.arn,
      aws_ssm_parameter.wendover-server-root_path.arn
    ]
  }

  statement {
    sid     = "CloudWatchAccess"
    effect  = "Allow"
    actions = [
      "logs:PutLogEvents"
    ]
    resources = [
      aws_cloudwatch_log_stream.wendover_api.arn
    ]
  }
}

resource "aws_iam_policy" "wendover_ecs_task_role" {
  name    = "WendoverECSTaskRole"
  policy  = data.aws_iam_policy_document.wendover_ecs_task_role.json
}

resource "aws_iam_role" "wendover_ecs_task_role" {
  name                = "WendoverECSTaskRole"
  assume_role_policy  = data.aws_iam_policy_document.wendover_ecs_execution_role_trust.json
  inline_policy {
    name   = "WendoverECSTaskPolicy"
    policy = data.aws_iam_policy_document.wendover_ecs_task_role.json
  }
}