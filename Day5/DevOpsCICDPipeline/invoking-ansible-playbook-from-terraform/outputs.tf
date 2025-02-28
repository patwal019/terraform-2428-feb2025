output "image_id" {
   value = data.docker_image.tektutor_image.id
}

output "container_id" {
   value = docker_container.container.id
}
