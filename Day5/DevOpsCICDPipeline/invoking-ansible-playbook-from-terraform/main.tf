resource "docker_container" "container" {
  container_name  = var.container_name
  host_name = var.container_name 
  image_name = var.image.name
}

resource "file_localfile" "myfile" {
  content = "some ip"
  filename = "./ip.txt"
}
