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

resource "aws_eip" "wendover_a" {
  domain      = "vpc"

  depends_on  = [aws_internet_gateway.wendover]
}

resource "aws_nat_gateway" "wendover_a" {
  subnet_id     = aws_subnet.wendover_a.id
  allocation_id = aws_eip.wendover_a.id

  depends_on    = [aws_internet_gateway.wendover]
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

resource "aws_eip" "wendover_b" {
  domain      = "vpc"

  depends_on  = [aws_internet_gateway.wendover]
}

resource "aws_nat_gateway" "wendover_b" {
  subnet_id     = aws_subnet.wendover_b.id
  allocation_id = aws_eip.wendover_b.id

  depends_on    = [aws_internet_gateway.wendover]
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
