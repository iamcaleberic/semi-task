resource "random_password" "mongo_root_password" {
  length  = 10
  special = false
}

resource "random_password" "mongo_semi_task_user_pass" {
  length  = 10
  special = false
}

resource "random_string" "mongo_root_username" {
  length           = 5
  special          = false
}

resource "kubernetes_namespace" "mongo" {
  metadata {
    name = "mongo"
  }
}

resource "kubernetes_namespace" "semi_task" {
  metadata {
    name = "semi-task"
  }
}

locals {
  mongodb_semi_task_db = "semi-task"
  mongodb_semi_task_user = "app"
  mongodb_port           = 27017
  mongodb_uri            = base64encode("mongodb://${local.mongodb_semi_task_user}:${random_password.mongo_semi_task_user_pass.result}@mongodb.${kubernetes_namespace.mongo.metadata.0.name}:${local.mongodb_port}")
}

resource "helm_release" "mongodb" {
  name      = "mongodb"
  chart     = "../charts/mongo"
  namespace = kubernetes_namespace.mongo.metadata.0.name

  set {
    name  = "credentials.MONGO_INITDB_ROOT_USERNAME"
    value = base64encode(random_password.mongo_root_password.result)
  }
  set {
    name  = "credentials.MONGO_INITDB_ROOT_PASSWORD"
    value = base64encode(random_password.mongo_root_password.result)
  }

  set {
    name  = "credentials.MONGO_INITDB_USERNAME"
    value = base64encode(local.mongodb_semi_task_user)
  }

  set {
    name  = "credentials.MONGO_INITDB_PASSWORD"
    value = base64encode(random_password.mongo_semi_task_user_pass.result)
  }

  set {
    name  = "credentials.MONGO_INITDB_DATABASE"
    value = base64encode(local.mongodb_semi_task_db)
  }

}

resource "helm_release" "semi_task" {
  name      = "semi-task"
  chart     = "../charts/semi-task"
  namespace = kubernetes_namespace.semi_task.metadata.0.name

  set {
    name  = "credentials.MONGODB_DATABASE"
    value = base64encode(local.mongodb_semi_task_db)
  }
  set {
    name  = "credentials.MONGODB_URI"
    value = local.mongodb_uri
  }

  depends_on = [
    helm_release.mongodb
  ]
}
