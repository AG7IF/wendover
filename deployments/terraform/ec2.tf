resource "aws_instance" "wendover_a" {
  name          = "wendover-a"
  instance_type = "t2.micro"
  ami           = var.ami_id
}


/*
resource "aws_eip" "wendover_a" {}
*/