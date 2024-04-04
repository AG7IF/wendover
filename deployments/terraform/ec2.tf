resource "aws_key_pair" "wendover" {
  key_name    = "wendover"
  public_key  = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIIdNXc6FMBxuBeDvMGKLeVrJHiIWmCOV2SRgF44x/I1z"
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