data "docker_image" "ubuntu_image" {
  name = "ubuntu:latest"
}

import {
  id =  "4083dd9d00687c51e7792f2155009a0b71b5ad9c95dc9bbcf1fd33c5b78898ba"
  to = docker_container.ubuntu_container2
}

resource "docker_container" "ubuntu_container2" {
  name = "container_2"
  hostname = "container_2"
  image = "tektutor/ubuntu-ansible-node:1.0"
}

