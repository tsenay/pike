data "aws_vpcs" "main" {
  tags = {
    Name = "test"
  }
}

output "aws_vpcs" {
  value = data.aws_vpcs.main
}


data "aws_caller_identity" "me" {}

output "caller_id" {
  value = data.aws_caller_identity.me
}

data "aws_subnet_ids" "public" {
  vpc_id = "vpc-0c33dc8cd64f408c4"
}

output "subnet_ids" {
  value = data.aws_subnet_ids.public
}

data "aws_vpc" "this" {
  id = "vpc-0c33dc8cd64f408c4"
}
