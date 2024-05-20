# AWS Configuration
variable "region" {
  description = "AWS region"
  default     = "us-west-2"
}

# Parameters
variable "api_root_path" {
  description = "Root path for the API endpoint"
  default     = "/api/v1"
}

# DNS configuration
variable "web_dns_zone_id" {
  description = "The ID of the Route53 zone to which DNS CNAME records should be added"
}

variable "web_domain" {
  description = "The domain to which the wendover web app is deployed"
  default = ""
}

# Email configuration
variable "cognito_reply_to_email" {
  description = "The reply-to email address used for Cognito SES configuration"
  default = ""
}

# KMS configuration
variable "aws_kms_principal_arn" {
  description = "ARN for the role that will own the KMS Key"
  default = ""
}

variable "terraform_cloud_role_arn" {
  description = "ARN for Terraform Cloud Role, used to allow TFC to manage the KMS key"
  default = ""
}
