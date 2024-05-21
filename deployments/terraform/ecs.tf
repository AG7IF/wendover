resource "aws_ecr_repository" "wendover" {
  name = "wendover"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  encryption_configuration {
    encryption_type = "AES256"
  }
}

resource "aws_ecs_cluster" "wendover" {
  name = "wendover"
}

resource "aws_ecs_service" "wendover_api" {
  name            = "wendover-api"
  launch_type     = "FARGATE"
  cluster         = aws_ecs_cluster.wendover.id
  task_definition = aws_ecs_task_definition.wendover_api.arn
  desired_count   = 1

  network_configuration {
    subnets  = [
      aws_subnet.wendover_private_c.id
    ]
    security_groups = [
      aws_security_group.wendover_api.id
    ]
  }

  depends_on = []
}

resource "aws_ecs_task_definition" "wendover_api"{
  family                    = "wendover-api"
  requires_compatibilities  = ["FARGATE"]
  cpu                       = 1024
  memory                    = 2048
  network_mode              = "awsvpc"

  execution_role_arn        = aws_iam_role.wendover_ecs_execution_role.arn
  task_role_arn             = aws_iam_role.wendover_ecs_task_role.arn

  container_definitions     = jsonencode([
    {
      name          = "wendsrv"
      image         = "${aws_ecr_repository.wendover.repository_url}:latest"
      essential     = true
      cpu           = 1024
      memory        = 2048
      interactive   = true
      portMappings  = [
        {
          containerPort = 22
          hostPort      = 22
        },
        {
          containerPort = 80
          hostPort      = 80
        }
      ]

      healthcheck = {
        command     = ["CMD-SHELL", "curl -f http://localhost/${var.api_root_path}/healthcheck || exit 1"]
        startPeriod = 300
      }

      /*
      secrets = [
        {
          name      = "WENDOVER_AWS_COGNITO_ISS"
          valueFrom = aws_ssm_parameter.wendover-aws-cognito-iss.arn
        },
        {
          name      = "WENDOVER_AWS_COGNITO_USERPOOL_ID"
          valueFrom = aws_ssm_parameter.wendover-aws-cognito-userpool_id.arn
        },
        {
          name      = "WENDOVER_AWS_LOG_GROUP_NAME"
          valueFrom = aws_ssm_parameter.wendover-aws-log_group_name.arn
        },
        {
          name      = "WENDOVER_AWS_LOG_STREAM_NAME"
          valueFrom = aws_ssm_parameter.wendover-aws-log_stream_name.arn
        },
        {
          name      = "WENDOVER_DATABASE_HOST"
          valueFrom = aws_ssm_parameter.wendover-database-host.arn
        },
        {
          name      = "WENDOVER_DATABASE_PORT"
          valueFrom = aws_ssm_parameter.wendover-database-port.arn
        },
        {
          name      = "WENDOVER_DATABASE_NAME"
          valueFrom = aws_ssm_parameter.wendover-database-name.arn
        },
        {
          name      = "WENDOVER_DATABASE_CREDENTIALS"
          valueFrom = aws_db_instance.wendover.master_user_secret[0].secret_arn
        },
        {
          name      = "WENDOVER_DATABASE_SSL"
          valueFrom = aws_ssm_parameter.wendover-database-ssl.arn
        },
        {
          name      = "WENDOVER_DATABASE_MIGRATION_SOURCE"
          valueFrom = aws_ssm_parameter.wendover-database-migration-source.arn
        },
        {
          name      = "WENDOVER_SERVER_ROOT_PATH"
          valueFrom = aws_ssm_parameter.wendover-server-root_path.arn
        },
      ]
      */

      logConfiguration = {
        logDriver = "awslogs"
        options   = {
          awslogs-group         = aws_cloudwatch_log_group.wendover.name
          awslogs-region        = var.region
          awslogs-stream-prefix = aws_cloudwatch_log_stream.wendover_api.name
        }
      }
    }
  ])
}
