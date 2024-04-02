resource "aws_vpc" "wendover" {
  cidr_block            =   "10.0.0.0/16"
  enable_dns_support    =   true
  enable_dns_hostnames  =   true

  tags = {
    Name    = "Wendover"
    Service = "wendover"
  }
}

resource "aws_subnet" "wendover_db_aza" {
  vpc_id              =   aws_vpc.wendover.id
  cidr_block          =   "10.0.1.0/24"
  availability_zone   =   "${var.region}a"

  tags = {
    Name    = "WendoverDB-A"
    Service = "wendover"
  }
}

resource "aws_subnet" "wendover_db_azb" {
  vpc_id              =   aws_vpc.wendover.id
  cidr_block          =   "10.0.2.0/24"
  availability_zone   =   "${var.region}b"

  tags = {
    Name    = "WendoverDB-B"
    Service = "wendover"
  }
}

resource "aws_db_subnet_group" "wendover" {
  name        =   "wendover-db"
  subnet_ids  =   [aws_subnet.wendover_db_aza.id, aws_subnet.wendover_db_azb.id]
}

resource "aws_security_group" "wendover_db" {
  name    =   "wendover-db"
  vpc_id  =   aws_vpc.wendover.id

  ingress {
    from_port   =   5432
    to_port     =   5432
    protocol    =   "tcp"
    cidr_blocks =   ["0.0.0.0/0"]
  }

  egress {
    from_port   =   5432
    to_port     =   5432
    protocol    =   "tcp"
    cidr_blocks =   ["0.0.0.0/0"]
  }
}
