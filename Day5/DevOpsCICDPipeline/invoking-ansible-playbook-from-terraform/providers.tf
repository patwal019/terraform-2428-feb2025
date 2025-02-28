terraform {
  required_providers {
    docker = {
      source = "tektutor/docker"
      version = "2.0.0"
    }
    local = {
      source = "hashicorp/local"
      version = "2.5.2"
    }
  }
}

provider "docker" {
  # Configuration options
}

provider "local" {
  # Configuration options
}
