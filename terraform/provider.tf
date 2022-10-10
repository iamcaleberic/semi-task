provider "kubernetes" {
  # Configuration options
  config_path = "~/.kube/config"
  config_context = "semi-task"
}


provider "helm" {
  # Configuration options
  kubernetes {
    config_path = "~/.kube/config"
    config_context = "semi-task"
  }
}

