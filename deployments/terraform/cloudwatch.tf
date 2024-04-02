resource "aws_cloudwatch_log_group" "wendover" {
  name = "wendover"
}

resource "aws_cloudwatch_log_stream" "wendover_web" {
  name            = "wendover-web"
  log_group_name  = aws_cloudwatch_log_group.wendover.name
}

resource "aws_cloudwatch_log_stream" "wendover_api" {
  name            = "wendover-api"
  log_group_name  = aws_cloudwatch_log_group.wendover.name
}
