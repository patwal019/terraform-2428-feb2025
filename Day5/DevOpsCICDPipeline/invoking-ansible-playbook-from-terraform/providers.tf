terraform {
  required_providers {
    docker = {
      source = "tektutor/docker"
      version = "2.0.0"
    }
  }
}

provider "docker" {
  # Configuration options
}
