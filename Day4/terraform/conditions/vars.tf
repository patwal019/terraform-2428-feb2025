variable "container_name" {
  description = "Name of the docker container"
  type        = string

  validation {
     condition = length(var.container_name) >= 5
     error_message = "The container name must be of length 5 or greater."
  }
}

variable "image_name" {
  description = "Name of the docker image"
  type        = string
  default     = "tektutor/ubuntu-ansible-node:1.0"
}
