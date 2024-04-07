resource "aws_key_pair" "wendover" {
  key_name    = "wendover"
  public_key  = var.public_key
}

resource "aws_instance" "wendover_a" {
  instance_type = "t2.micro"
  ami           = var.ami_id
  key_name      = aws_key_pair.wendover.key_name

  tags = {
    Name = "wendover-a"
  }
}

/*
resource "aws_eip" "wendover_a" {}
*/