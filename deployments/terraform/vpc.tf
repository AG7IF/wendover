resource "aws_vpc" "wendover" {
  cidr_block            = "10.0.0.0/16"
  enable_dns_support    = true
  enable_dns_hostnames  = true

  tags = {
    Name = "Wendover"
  }
}

resource "aws_internet_gateway" "wendover" {
  vpc_id = aws_vpc.wendover.id

  tags = {
    Name = "Wendover"
  }
}

resource "aws_route_table" "wendover_public" {
  vpc_id = aws_vpc.wendover.id

  route {
    cidr_block  = "0.0.0.0/0"
    gateway_id  = aws_internet_gateway.wendover.id
  }

  tags = {
    Name = "WendoverPublic"
  }
}

/*
Subnet PublicC: 10.0.0.0/22
0000 1010 | 0000 0000 | 0000 0000 | 0000 0000
1111 1111 | 1111 1111 | 1111 1100 | 0000 0000

Subnet PrivateC: 10.0.4.0/22
0000 1010 | 0000 0000 | 0000 0100 | 0000 0000
1111 1111 | 1111 1111 | 1111 1100 | 0000 0000
*/
resource "aws_subnet" "wendover_public_c" {
  vpc_id                  = aws_vpc.wendover.id
  cidr_block              = "10.0.0.0/22"
  availability_zone       = "${var.region}c"
  map_public_ip_on_launch = true

  tags = {
    Name = "WendoverPublic-C"
  }
}

resource "aws_route_table_association" "wendover_public_c" {
  route_table_id  = aws_route_table.wendover_public.id
  subnet_id       = aws_subnet.wendover_public_c.id
}

resource "aws_subnet" "wendover_private_c" {
  vpc_id                  = aws_vpc.wendover.id
  cidr_block              = "10.0.4.0/22"
  availability_zone       = "${var.region}c"
  map_public_ip_on_launch = false

  tags = {
    Name = "WendoverPrivate-C"
  }
}

resource "aws_eip" "wendover_private_c" {
  tags = {
    Name = "WendoverPrivate-C"
  }
}

resource "aws_nat_gateway" "wendover_private_c" {
  allocation_id = aws_eip.wendover_private_c.id
  subnet_id     = aws_subnet.wendover_private_c.id

  tags = {
    Name = "WendoverPrivate-C"
  }

  depends_on = [aws_internet_gateway.wendover]
}

resource "aws_route_table" "wendover_private_c" {
  vpc_id = aws_vpc.wendover.id

  route {
    cidr_block           = "0.0.0.0/0"
    nat_gateway_id       = aws_nat_gateway.wendover_private_c.id
  }

  tags = {
    Name = "WendoverPrivate-C"
  }
}

resource "aws_route_table_association" "wendover_private_c" {
  route_table_id  = aws_route_table.wendover_private_c.id
  subnet_id       = aws_subnet.wendover_private_c.id
}

resource "aws_db_subnet_group" "wendover" {
  name = "wendover"

  subnet_ids = [
    aws_subnet.wendover_private_c.id
  ]
}

resource "aws_security_group" "wendover_api" {
  name    = "wendover-api"
  vpc_id  = aws_vpc.wendover.id

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
    from_port         = 0
    to_port           = 0
    protocol          = "-1"
    cidr_blocks       = ["0.0.0.0/0"]
    ipv6_cidr_blocks  = ["::/0"]
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
    from_port         = 0
    to_port           = 0
    protocol          = "-1"
    cidr_blocks       = ["0.0.0.0/0"]
    ipv6_cidr_blocks  = ["::/0"]
  }

}
