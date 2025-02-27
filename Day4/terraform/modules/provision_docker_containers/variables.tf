variable docker_image_name {
   description = "Name of the docker image"
   default     = "tektutor/ubuntu-ansible-node:1.0"
}

variable "container_count" {
    description = "Number of containers you wish to create"
    type = number
     
    validation {
       condition = var.container_count > 1 && var.container_count <=5
       error_message = "The container count should be in the range 0 to 5"
    }
}
