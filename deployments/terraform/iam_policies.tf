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
        "Sid"     =   "SecretsManagerDbCredentialsAccess",
        "Effect"  =   "Allow",
        "Action"  =   [
          "secretsmanager:GetSecretValue",
          "secretsmanager:PutResourcePolicy",
          "secretsmanager:PutSecretValue",
          "secretsmanager:DeleteSecret",
          "secretsmanager:DescribeSecret",
          "secretsmanager:TagResource"
        ],
        "Resource"  = "arn:aws:secretsmanager:*:*:secret:rds-db-credentials/*"
      },
      {
        "Sid"     =   "RDSDataServiceAccess",
        "Effect"  =   "Allow",
        "Action"  =   [
          "dbqms:CreateFavoriteQuery",
          "dbqms:DescribeFavoriteQueries",
          "dbqms:UpdateFavoriteQuery",
          "dbqms:DeleteFavoriteQueries",
          "dbqms:GetQueryString",
          "dbqms:CreateQueryHistory",
          "dbqms:DescribeQueryHistory",
          "dbqms:UpdateQueryHistory",
          "dbqms:DeleteQueryHistory",
          "rds-data:ExecuteSql",
          "rds-data:ExecuteStatement",
          "rds-data:BatchExecuteStatement",
          "rds-data:BeginTransaction",
          "rds-data:CommitTransaction",
          "rds-data:RollbackTransaction",
          "secretsmanager:CreateSecret",
          "secretsmanager:ListSecrets",
          "secretsmanager:GetRandomPassword",
          "tag:GetResources",
          "ec2:DescribeNetworkInterfaces",
          "ec2:CreateNetworkInterface",
          "ec2:DeleteNetworkInterface",
          "ec2:DescribeInstances",
          "ec2:AttachNetworkInterface"
        ],
        "Resource"  =   "*"
      },
      {
        "Sid"     =   "AllowDescribeOnAllLogGroups",
        "Effect"  =   "Allow",
        "Action": [
          "logs:DescribeLogGroups"
        ],
        "Resource": [
          "*"
        ]
      },
      {
        "Sid"     =   "AllowDescribeOfAllLogStreamsOnDmsTasksLogGroup",
        "Effect"  =   "Allow",
        "Action"  =   [
          "logs:DescribeLogStreams"
        ],
        "Resource"  =   [
          "arn:aws:logs:*:*:log-group:dms-tasks-*",
          "arn:aws:logs:*:*:log-group:dms-serverless-replication-*"
        ]
      },
      {
        "Sid"     =   "AllowCreationOfDmsLogGroups",
        "Effect"  =   "Allow",
        "Action"  =   [
          "logs:CreateLogGroup"
        ],
        "Resource"  =   [
          "arn:aws:logs:*:*:log-group:dms-tasks-*",
          "arn:aws:logs:*:*:log-group:dms-serverless-replication-*:log-stream:"
        ]
      },
      {
        "Sid"     =   "AllowCreationOfDmsLogStream",
        "Effect"  =   "Allow",
        "Action"  =   [
          "logs:CreateLogStream"
        ],
        "Resource"  =   [
          "arn:aws:logs:*:*:log-group:dms-tasks-*:log-stream:dms-task-*",
          "arn:aws:logs:*:*:log-group:dms-serverless-replication-*:log-stream:dms-serverless-*"
        ]
      },
      {
        "Sid"     =   "AllowUploadOfLogEventsToDmsLogStream",
        "Effect"  =   "Allow",
        "Action"  =   [
          "logs:PutLogEvents"
        ],
        "Resource"  =   [
          "arn:aws:logs:*:*:log-group:dms-tasks-*:log-stream:dms-task-*",
          "arn:aws:logs:*:*:log-group:dms-serverless-replication-*:log-stream:dms-serverless-*"
        ]
      }
    ]
  })
}
