resource "aws_route53domains_registered_domain" "wendover" {
  domain_name = var.wendover_domain
}

resource "aws_route53_zone" "wendover" {
  name = "ag7if.net"
}