data "docker_image" "ubuntu_image" {
  name = "ubuntu:latest"
}

resource "docker_container" "ubuntu_container1" {
  name  = "ubuntu_container1"
  hostname = "ubuntu_container1"
  image = data.docker_image.ubuntu_image.name

  must_run = true
  command = [
     "tail",
     "-f",
     "/dev/null"
  ]
}

resource "docker_container" "ubuntu_container2" {
  name  = "ubuntu_container2"
  hostname = "ubuntu_container2"
  image = data.docker_image.ubuntu_image.name

  must_run = true
  command = [
     "tail",
     "-f",
     "/dev/null"
  ]
}
