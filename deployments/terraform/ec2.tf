resource "aws_key_pair" "wendover" {
  key_name    = "wendover"
  public_key  = var.public_key
}

resource "aws_network_interface" "wendover_a" {
  ami       = var.ami_id
  subnet_id = aws_subnet.wendover_a.id

  tags = {
    Name = "wendover-a"
  }
}

resource "aws_instance" "wendover_a" {
  instance_type = "t2.micro"
  ami           = var.ami_id
  key_name      = aws_key_pair.wendover.key_name

  network_interface {
    network_interface_id = aws_network_interface.wendover_a.id
    device_index         = 0
  }

  tags = {
    Name = "wendover-a"
  }
}

/*
resource "aws_eip" "wendover_a" {}
*/