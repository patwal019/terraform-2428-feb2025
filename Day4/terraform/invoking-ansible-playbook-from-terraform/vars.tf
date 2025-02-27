variable "container_name" {
  description = "Name of the docker container"
  type        = string
}

variable "image_name" {
  description = "Name of the docker image"
  type        = string
  default     = "tektutor/ubuntu-ansible-node:1.0"
}
