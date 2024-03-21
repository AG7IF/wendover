resource "aws_amplify_app" "wendover" {
  name         = "wendover"
  repository   = var.amplify_repository

  environment_variables = {
    AMPLIFY_DIFF_DEPLOY       = "false"
    AMPLIFY_MONOREPO_APP_ROOT = "web"
  }

  custom_rule {
    source = "/<*>"
    status = "404-200"
    target = "/index.html"
  }

  build_spec = <<-EOT
    version: 1
    applications:
      - appRoot: web
        frontend:
          phases:
            preBuild:
              commands:
                - yarn install
            build:
              commands:
                - yarn generate
          artifacts:
            baseDirectory: '.output/public'
            files:
              - '**/*'
          cache:
            paths:
              - node_modules/**/*
EOT
}

resource "aws_amplify_branch" "main" {
  app_id      = aws_amplify_app.wendover.id
  branch_name = "main"
}

resource "aws_amplify_domain_association" "wendover" {
  app_id      = aws_amplify_app.wendover.id
  domain_name = aws_route53domains_registered_domain.wendover.domain_name

  sub_domain {
    branch_name = aws_amplify_branch.main.branch_name
    prefix      = "encampment"
  }
}