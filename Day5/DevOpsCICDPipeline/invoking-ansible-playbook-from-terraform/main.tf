resource "docker_container" "container" {
  container_name  = var.container_name
  host_name = var.container_name 
  image_name = var.image.name
}

resource "local_file" "invoke_ansible_playbook" {
  content = "some ip"
  filename = "./ip.txt"

  provisioner "local-exec" {
     command = "ansible-playbook -i ./hosts install-nginx-playbook.yml"
  }
}
