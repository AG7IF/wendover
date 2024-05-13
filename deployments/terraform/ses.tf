resource "aws_ses_domain_identity" "wendover" {
  domain = var.web_domain
}

resource "aws_ses_email_identity" "wendover" {
  email       = var.cognito_from_email

  depends_on  = [aws_ses_domain_identity.wendover]
}