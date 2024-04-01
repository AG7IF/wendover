data "aws_iam_policy_document" "wendover_lambda_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "wendover_lambda_role" {
  statement {
    sid     = "DBSecretsManagerAccess"
    effect  = "Allow"

    actions = [
      "secretsmanager:GetSecretValue",
      "secretsmanager:PutResourcePolicy",
      "secretsmanager:PutSecretValue",
      "secretsmanager:DeleteSecret",
      "secretsmanager:DescribeSecret",
      "secretsmanager:TagResource"
    ]

    resources = ["arn:aws:secretsmanager:${var.region}:${var.aws_account_id}:secret:rds!db*"]
  }

  statement {
    sid     = "ParameterStoreAccess"
    effect  = "Allow"

    actions = [
      "ssm:GetParameter",
      "ssm:GetParameterHistory",
      "ssm:GetParameters",
      "ssm:GetParametersByPath"
    ]

    resources = ["arn:aws:ssm:${var.region}:${var.aws_account_id}:parameter/wendover/*"]
  }

  statement {
    sid     = "RDSDataServiceAccess"
    effect  = "Allow"

    actions = [
      "rds-data:ExecuteSql",
      "rds-data:ExecuteStatement",
      "rds-data:BatchExecuteStatement",
      "rds-data:BeginTransaction",
      "rds-data:CommitTransaction",
      "rds-data:RollbackTransaction",
      "tag:GetResources"
    ]

    resources = [aws_db_instance.wendover.arn]
  }

  statement {
    sid     = "AccessRDSSubnets"
    effect  = "Allow"

    actions = [
      "ec2:DescribeNetworkInterfaces",
      "ec2:CreateNetworkInterface",
      "ec2:DeleteNetworkInterface",
      "ec2:DescribeInstances",
      "ec2:AttachNetworkInterface"
    ]

    resources = [aws_vpc.wendover.arn]
  }

  statement {
    sid     = "CloudWatchAccess"
    effect  = "Allow"

    actions = [
      "logs:DescribeLogGroups",
      "logs:DescribeLogStreams",
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]

    resources = ["arn:aws:logs:${var.region}:${var.aws_account_id}:log-group:/aws/lambda/*"]
  }
}

resource "aws_iam_policy" "wendover_lambda_role" {
  name = "WendoverLambdaRolePolicy"
  policy = data.aws_iam_policy_document.wendover_lambda_role.json
}

resource "aws_iam_role" "wendover_lambda_role" {
  name                  =   "WendoverLambdaRole"
  assume_role_policy    =   data.aws_iam_policy_document.wendover_lambda_assume_role.json
  managed_policy_arns   =   [aws_iam_policy.wendover_lambda_role.arn]
}

resource "aws_cloudwatch_log_group" "wendover_lambda" {
  name                =   "/aws/lambda/wendover"
  retention_in_days   =   14
}
