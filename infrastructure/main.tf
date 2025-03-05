resource "kubernetes_manifest" "server" {
  manifest = yamldecode(file("${path.root}/server-deployment.yml"))
}

resource "kubernetes_manifest" "service" {
  manifest = yamldecode(file("${path.root}/server-service.yml"))
}