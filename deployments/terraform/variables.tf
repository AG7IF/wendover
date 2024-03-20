variable "region" {
  description = "AWS region"
  default     = "us-west-2"
}

variable "terraform_cloud_role_arn" {
  description = "The ARN to use when granting permissions to the Terraform Cloud role"
}

variable "aws_principal_arn" {
  description = "Value to use as the principal for IAM users for KMS keys"
  default = ""
}

variable "amplify_repository" {
  description = "Git repository used with AWS Amplify"
  default = "https://github.com/ag7if/wendover"
}

variable "amplify_repository_access_token" {
  description = "The repository access token that Amplify will use to deploy the website"
  default = ""
}

variable "wendover_domain" {
  description = "The domain name for weblair"
  default = ""
}
