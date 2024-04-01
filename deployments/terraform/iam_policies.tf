data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_policy" "wendover_lambda_role_policy" {
  name = "WendoverLambdaRolePolicy"
  policy = jsonencode({
    "Version"     =   "2012-10-17",
    "Statement"   =   [
      {
        "Sid" = "DBSecretsManagerAccess",
        "Effect" = "Allow",
        "Action" = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:PutResourcePolicy",
          "secretsmanager:PutSecretValue",
          "secretsmanager:DeleteSecret",
          "secretsmanager:DescribeSecret",
          "secretsmanager:TagResource"
        ],
        "Resource"  = "arn:aws:secretsmanager:${var.region}:${var.aws_account_id}:secret:rds!db*"
      },
      {
        "Sid"     = "ParameterStoreAccess",
        "Effect"  = "Allow",
        "Action"  = [
          "ssm:GetParameter",
          "ssm:GetParameterHistory",
          "ssm:GetParameters",
          "ssm:GetParametersByPath"
        ],
        "Resource" = "arn:aws:ssm:${var.region}:${var.aws_account_id}:parameter/wendover/*"

      },
      {
        "Sid"     = "RDSDataServiceAccess",
        "Effect"  = "Allow",
        "Action"  = [
          "rds-data:ExecuteSql",
          "rds-data:ExecuteStatement",
          "rds-data:BatchExecuteStatement",
          "rds-data:BeginTransaction",
          "rds-data:CommitTransaction",
          "rds-data:RollbackTransaction",
          "tag:GetResources"
        ]
        "Resource" = aws_db_instance.wendover.arn
      },
      {
        "Sid"     = "AccessRDSSubnets"
        "Effect"  = "Allow",
        "Action"  = [
          "ec2:DescribeNetworkInterfaces",
          "ec2:CreateNetworkInterface",
          "ec2:DeleteNetworkInterface",
          "ec2:DescribeInstances",
          "ec2:AttachNetworkInterface"
        ],
        "Resource" = aws_vpc.wendover.arn
      },
      {
        "Sid"     = "CloudWatchAccess",
        "Effect"  = "Allow",
        "Action"  = [
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        "Resource" = "arn:aws:logs:${var.region}:${var.aws_account_id}:log-group:/aws/lambda/*"
      }
    ]
  })
}
