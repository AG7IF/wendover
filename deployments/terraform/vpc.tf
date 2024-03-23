resource "aws_vpc" "wendover" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "wendover_db_aza" {
  vpc_id = aws_vpc.wendover.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "${var.region}a"
}

resource "aws_subnet" "wendover_db_azb" {
  vpc_id = aws_vpc.wendover.id
  cidr_block = "10.0.2.0/24"
  availability_zone = "${var.region}b"
}
