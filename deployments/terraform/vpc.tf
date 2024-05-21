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

resource "aws_route_table" "wendover" {
  vpc_id = aws_vpc.wendover.id

  route {
    cidr_block = "10.0.0.0/16"
    gateway_id = "local"
  }

  tags = {
    Name = "Wendover"
  }
}

/*
Subnet A:
0000 1010 | 0000 0000 | 0000 0000 | 0000 0000
1111 1111 | 1111 1111 | 1111 1100 | 0000 0000

Subnet B:
0000 1010 | 0000 0000 | 0000 0100 | 0000 0000
1111 1111 | 1111 1111 | 1111 1100 | 0000 0000

Subnet C:
0000 1010 | 0000 0000 | 0000 1000 | 0000 0000
1111 1111 | 1111 1111 | 1111 1100 | 0000 0000

Subnet D:
0000 1010 | 0000 0000 | 0000 1100 | 0000 0000
1111 1111 | 1111 1111 | 1111 1100 | 0000 0000

Subnet E:
0000 1010 | 0000 0000 | 0001 0000 | 0000 0000
1111 1111 | 1111 1111 | 1111 1100 | 0000 0000
*/
resource "aws_subnet" "wendover_a" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.0.0/22"
  availability_zone   = "${var.region}a"

  tags = {
    Name = "Wendover-A"
  }
}

resource "aws_route_table_association" "wendover_a" {
  subnet_id       = aws_subnet.wendover_a.id
  route_table_id  = aws_route_table.wendover.id
}

resource "aws_subnet" "wendover_b" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.4.0/22"
  availability_zone   = "${var.region}b"

  tags = {
    Name    = "Wendover-B"
  }
}

resource "aws_route_table_association" "wendover_b" {
  subnet_id       = aws_subnet.wendover_b.id
  route_table_id  = aws_route_table.wendover.id
}

resource "aws_subnet" "wendover_c" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.8.0/22"
  availability_zone   = "${var.region}c"

  tags = {
    Name    = "Wendover-C"
  }
}

resource "aws_route_table_association" "wendover_c" {
  subnet_id       = aws_subnet.wendover_c.id
  route_table_id  = aws_route_table.wendover.id
}

resource "aws_subnet" "wendover_d" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.12.0/22"
  availability_zone   = "${var.region}d"

  tags = {
    Name    = "Wendover-D"
  }
}

resource "aws_route_table_association" "wendover_d" {
  subnet_id       = aws_subnet.wendover_d.id
  route_table_id  = aws_route_table.wendover.id
}

resource "aws_subnet" "wendover_e" {
  vpc_id              = aws_vpc.wendover.id
  cidr_block          = "10.0.16.0/22"

  tags = {
    Name    = "Wendover-E"
  }
}

resource "aws_eip" "wendover_e" {
  domain      = "vpc"

  depends_on  = [aws_internet_gateway.wendover]
}

resource "aws_nat_gateway" "wendover_e" {
  subnet_id     = aws_subnet.wendover_e.id
  allocation_id = aws_eip.wendover_e.id

  tags = {
    Name = "Wendover-E"
  }

  depends_on  = [aws_internet_gateway.wendover]
}

resource "aws_route_table" "wendover_e" {
  vpc_id = aws_vpc.wendover.id

  route {
    cidr_block      = aws_subnet.wendover_e.cidr_block
    nat_gateway_id  = aws_nat_gateway.wendover_e.id
  }

  tags = {
    Name = "Wendover-E"
  }
}

resource "aws_route_table_association" "wendover_e" {
  subnet_id       = aws_subnet.wendover_e.id
  route_table_id  = aws_route_table.wendover_e.id
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
