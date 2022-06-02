terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.16.0"
    }
  }
}

provider "aws" {
  region = "eu-central-1"
}

# EC2 Instance
resource "aws_instance" "test-instance-flugel" {
  ami               = data.aws_ami.ubuntu.id
  instance_type     = "t2.micro"
  availability_zone = "eu-central-1a"

  tags = {
    Name  = "Flugel"
    Owner = "InfraTeam"
  }
}

data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"] # Canonical

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}


# S3 bucket
resource "aws_s3_bucket" "test-bucket-flugel" {
  bucket = "test-bucket-flugel"

  tags = {
    Name  = "Flugel"
    Owner = "InfraTeam"
  }
}

