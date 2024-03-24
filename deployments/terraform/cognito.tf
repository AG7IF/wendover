resource "aws_cognito_user_pool" "wendover" {
  name = "wendover"

  deletion_protection = "ACTIVE"

  username_configuration {
    case_sensitive = false
  }

  password_policy {
    minimum_length                    =   16
    temporary_password_validity_days  =    7
  }

  mfa_configuration = "ON"
  software_token_mfa_configuration {
    enabled = true
  }

  account_recovery_setting {
    recovery_mechanism {
      name      =  "verified_email"
      priority  =  1
    }
  }

  admin_create_user_config {
    allow_admin_create_user_only = true
  }

  email_configuration {
    email_sending_account   =   "DEVELOPER"
    source_arn              =   var.cognito_from_email_arn
    from_email_address      =   var.cognito_from_email
    reply_to_email_address  =   var.cognito_reply_to_email
  }
}

resource "aws_cognito_user_pool_client" "wendover_web" {
  user_pool_id = aws_cognito_user_pool.wendover.id
  name = "wendover-web"

  generate_secret = true

  explicit_auth_flows     =   ["ALLOW_USER_SRP_AUTH", "ALLOW_REFRESH_TOKEN_AUTH"]
  auth_session_validity   =    3 # minutes
  refresh_token_validity  =   30 # days
  access_token_validity   =    1 # 1 hour
  id_token_validity       =    1 # hour
  enable_token_revocation =   true

  read_attributes   =   ["email"]
  write_attributes  =   ["email"]
}
