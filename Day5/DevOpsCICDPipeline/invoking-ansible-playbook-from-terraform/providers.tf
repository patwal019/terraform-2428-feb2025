terraform {
  required_providers {
    docker = {
      source = "tektutor/docker"
      version = "2.0.0"
    }
    local = {
      source = "tektutor/localfile"
    }
  }
}

provider "docker" {
  # Configuration options
}

provider "localfile" {
  # Configuration options
}
