resource "aws_vpc" "wendover" {
  cidr_block            = "10.0.0.0/16"
  enable_dns_support    = true
  enable_dns_hostnames  = true

  tags = {
    Name    = "Wendover"
    Service = "wendover"
  }
}

resource "aws_subnet" "wendover_a" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.1.0/24"
  availability_zone   = "${var.region}a"

  tags = {
    Name    = "Wendover-A"
    Service = "wendover"
  }
}

resource "aws_subnet" "wendover_b" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.2.0/24"
  availability_zone   = "${var.region}b"

  tags = {
    Name    = "Wendover-B"
    Service = "wendover"
  }
}

resource "aws_subnet" "wendover_c" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.3.0/24"
  availability_zone   = "${var.region}c"

  tags = {
    Name    = "Wendover-C"
    Service = "wendover"
  }
}

resource "aws_subnet" "wendover_d" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.4.0/24"
  availability_zone   = "${var.region}d"

  tags = {
    Name    = "Wendover-D"
    Service = "wendover"
  }
}

resource "aws_db_subnet_group" "wendover" {
  name        = "wendover-db"
  subnet_ids  = [
    aws_subnet.wendover_a.id,
    aws_subnet.wendover_b.id,
    aws_subnet.wendover_c.id,
    aws_subnet.wendover_d.id
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
