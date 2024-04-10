resource "aws_vpc" "wendover" {
  cidr_block            = "10.0.0.0/16"
  enable_dns_support    = true
  enable_dns_hostnames  = true

  tags = {
    Name    = "Wendover"
    Service = "wendover"
  }
}

resource "aws_internet_gateway" "wendover" {
  vpc_id = aws_vpc.wendover.id

  tags = {
    Name = "main"
  }
}

resource "aws_route_table" "wendover" {
  vpc_id = aws_vpc.wendover.id

  route {
    cidr_block = "10.0.0.0/16"
    gateway_id = "local"
  }
}
/*
Subnet A:
0000 0010 | 0000 0000 | 0000 0000 | 0000 0000
1111 1111 | 1111 1111 | 1111 1100 | 0000 0000

Subnet B:
*/

resource "aws_subnet" "wendover_a" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.4.0/22"
  availability_zone   = "${var.region}a"

  tags = {
    Name    = "Wendover-A"
    Service = "wendover"
  }
}

resource "aws_subnet" "wendover_b" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.8.0/22"
  availability_zone   = "${var.region}b"

  tags = {
    Name    = "Wendover-B"
    Service = "wendover"
  }
}

resource "aws_db_subnet_group" "wendover" {
  name        = "wendover-db"
  subnet_ids  = [
    aws_subnet.wendover_a.id,
    aws_subnet.wendover_b.id
  ]
}

resource "aws_security_group" "wendover_api" {
  name    = "wendover-api"
  vpc_id  = aws_vpc.wendover.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks =   ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks =   ["0.0.0.0/0"]
  }

  egress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "wendover_db" {
  name    = "wendover-db"
  vpc_id  = aws_vpc.wendover.id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_acm_certificate" "wendover_vpn" {
  domain_name         =   var.web_full_domain
  validation_method   =   "DNS"
}

resource "aws_route53_record" "wendover_vpn_validation" {
  for_each = {
    for dvo in aws_acm_certificate.wendover_vpn.domain_validation_options : dvo.domain_name => {
      name    =   dvo.resource_record_name
      record  =   dvo.resource_record_value
      type    =   dvo.resource_record_type
    }
  }

  allow_overwrite =   true
  name            =   each.value.name
  records         =   [each.value.record]
  ttl             =   60
  type            =   each.value.type
  zone_id         =   var.web_dns_zone_id
}

/*
resource "aws_acm_certificate_validation" "wendover_vpn_validation" {
  certificate_arn         =   aws_acm_certificate.wendover.arn
  validation_record_fqdns =   [ for record in aws_route53_record.wendover_validation : record.fqdn ]
}
*/

resource "aws_ec2_client_vpn_endpoint" "wendover" {
  description             = "wendover-a"
  server_certificate_arn  = aws_acm_certificate.wendover_vpn.arn
  client_cidr_block       = "10.0.252.0/22"
  vpn_port = "1194"

  connection_log_options {
    enabled = true
    cloudwatch_log_group  = aws_cloudwatch_log_group.wendover.name
    cloudwatch_log_stream = aws_cloudwatch_log_stream.wendover_vpn.name
  }

}