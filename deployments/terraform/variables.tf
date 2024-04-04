# AWS Configuration
variable "aws_account_id" {
  description =   "AWS Account ID"
  default     =   ""
}

variable "region" {
  description = "AWS region"
  default     = "us-west-2"
}

# EC2 Configuration
variable "ami_id" {
  description = "The AMI ID to use for EC2 instances"
  default     = "ami-0395649fbe870727e"
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

variable "web_full_domain" {
  description = "The full domain—including subdomain—for the wendover web app"
  default = ""
}

# Email configuration
variable "cognito_from_email_arn" {
  description = "ARN for the SES verified email identity that cognito will use for sending verification emails"
  default = ""
}

variable "cognito_from_email" {
  description = "The email address used for Cognito SES configuration"
  default = ""
}

variable "cognito_reply_to_email" {
  description = "The reply-to email address used for Cognito SES configuration"
  default = ""
}

# KMS configuration
variable "aws_principal_arn" {
  description = "ARN for the role that will own the KMS Key"
  default = ""
}

variable "terraform_cloud_role_arn" {
  description = "ARN for Terraform Cloud Role, used to allow TFC to manage the KMS key"
  default = ""
}
