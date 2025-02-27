data "docker_image" "tektutor_image" {
  name = var.image_name
}

resource "docker_container" "container" {
  name  = var.container_name
  hostname = var.container_name 
  image = data.docker_image.tektutor_image.name
  must_run = true

  ports {
    internal = "22"
    external = "2001"
  }

  ports {
     internal = "80"
     external = "8001"
  }

  lifecyle {
     precondition {
        condition = data.docker_image == "tektutor/ubuntu-ansible-node:1.0"
	error_message = "The selected docker image is not supported"
     }

     postcondition {
        condition = name == hostname
	error_message = "As per best practices container name and hostname must be same"
     }
  }
}

resource "local_file" "invoke_ansible_playbook" {
  content = "some ip"
  filename = "./ip.txt"

  connection {
     type = "ssh"
     user = "root"
     port = "2001"
     host = "localhost"
  }

  provisioner "remote-exec" {
  	inline = [
		"hostname",
		"hostname -i",
	]
  }
}
