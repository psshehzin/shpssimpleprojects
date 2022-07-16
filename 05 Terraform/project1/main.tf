provider "aws" {
  region = "ap-south-1"
  access_key = ""
  secret_key = ""
}
# terraform init : to install providers
# terraform plan : to see what all changes will be applied
# terraform apply : to apply changes  (--auto-approve can be given and we will not have to manually give yes)
# terraform destroy to destroy images
# the resoursce name will be of format: resource "<provider>_<resource_type" "name" the name will be used  by terraform only
# terraform state list  will list all resources deployed
# terraform state show resource --> will give a lot of info
# if we declare output blocks in terraform and get the outputs from there
# output name {
#    value = resoursename.customname.attribute 
# }
# we can run a terraform refresh to get all details and pint output
# terraform apply -target aws_instance.servername will just install awsinstance
#variables in terraform
# variable "name"{   "all beow fields are optional"
#   decription =
#   default =
#   type =
# }
# referncing var  var.variable name
# if we dont provide any value it will auto prompt for that value while running
# terraform apply -var "varname=value" //to give value at the time of apply
# atually how we do it is we create a file called terraform.tfvars name should be this and specify value ex var1 = "value"
# if custom file name is used then terraform apply -var-file custom.tfvars
# terraform vars can have types and if types are used it will put error if a different type of value is provided
# variable as a list of objects below
# var1 = [(name= "hi",vwa= "test"),(name= "there", vwa="cute")]
# in main file var.var1[0].name
resource "aws_instance" "test1" {
  ami           = "ami-052cef05d01020f1d"
  instance_type = "t2.micro"
  tags = {
    Name = "Terraformmade"
  }

 }
resource "aws_vpc" "main" {
  cidr_block       = "10.0.0.0/16"

  tags = {
    Name = "terraformvpc"
  }
}
resource "aws_subnet" "main" {
  vpc_id     = aws_vpc.main.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "terraform subnet"
  }
}