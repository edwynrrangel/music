resource "helm_release" "mongodb_release" {
  name             = var.release_name
  chart            = "edwynrrangel/banking-service"
  namespace        = var.namespace
  version          = "0.1.0"
  cleanup_on_fail  = true
  lint             = true
  timeout          = 60
  wait             = true

  set {
    name  = "preInstallCheck.enabled"
    value = "true"
  }

  set {
    name  = "preInstallCheck.pvcName"
    value = kubernetes_persistent_volume_claim.banking_service_mongodb_pvc.metadata[0].name
  }

  set {
    name  = "preInstallCheck.serviceAccountName"
    value = var.service_account_name
  }

  set {
    name  = "global.enabled"
    value = "true"
  }

  set {
    name  = "global.projectName"
    value = var.release_name
  }

  set {
    name  = "global.clusterName"
    value = var.namespace
  }

  set {
    name  = "global.image.repository"
    value = "mongo"
  }

  set {
    name  = "global.image.tag"
    value = "7.0.10-rc0-jammy"
  }

  set {
    name  = "global.image.pullPolicy"
    value = "Always"
  }

  set {
    name  = "global.namespace"
    value = var.namespace
  }

  set {
    name  = "app.command[0]"
    value = "mongod"
  }

  set {
    name  = "app.command[1]"
    value = "--bind_ip_all"
  }

  set {
    name  = "app.command[2]"
    value = "--tlsMode"
  }

  set {
    name  = "app.command[3]"
    value = "requireTLS"
  }

  set {
    name  = "app.command[4]"
    value = "--tlsCertificateKeyFile"
  }

  set {
    name  = "app.command[5]"
    value = "/etc/ssl/mongo/mongodb.pem"
  }

  set {
    name  = "app.command[6]"
    value = "--tlsCAFile"
  }

  set {
    name  = "app.command[7]"
    value = "/etc/ssl/mongo/ca.pem"
  }

  set {
    name  = "app.command[8]"
    value = "--tlsAllowConnectionsWithoutCertificates"
  }

  set {
    name  = "app.replicaCount"
    value = "1"
  }

  set {
    name  = "app.serviceAccount.enabled"
    value = "false"
  }
  
  set {
    name  = "app.env[0].name"
    value = ""
  }

  set {
    name  = "app.secrets[0].secretKey"
    value = "MONGO_INITDB_ROOT_USERNAME"
  }

  set {
    name  = "app.secrets[0].secretName"
    value = var.db_secret_name
  }

  set {
    name  = "app.secrets[1].secretKey"
    value = "MONGO_INITDB_ROOT_PASSWORD"
  }

  set {
    name  = "app.secrets[1].secretName"
    value = var.db_secret_name
  }

  set {
    name  = "app.service.type"
    value = "NodePort"
  }

  set {
    name  = "app.service.ports[0].name"
    value = "mongodb"
  }

  set {
    name  = "app.service.ports[0].port"
    value = "27017"
  }

  set {
    name  = "app.service.ports[0].targetPort"
    value = "27017"
  }

  set {
    name  = "app.service.ports[0].nodePort"
    value = "30017"
  }

  set {
    name  = "app.service.ports[0].protocol"
    value = "TCP"
  }

  set {
    name  = "app.volumes[0].name"
    value = kubernetes_persistent_volume_claim.banking_service_mongodb_pvc.metadata[0].name
  }

  set {
    name  = "app.volumes[0].persistentVolumeClaim.claimName"
    value = kubernetes_persistent_volume_claim.banking_service_mongodb_pvc.metadata[0].name
  }

  set {
    name  = "app.volumeMounts[0].name"
    value = kubernetes_persistent_volume_claim.banking_service_mongodb_pvc.metadata[0].name
  }

  set {
    name  = "app.volumeMounts[0].mountPath"
    value = "/data/db"
  }

  set {
    name  = "app.volumes[1].name"
    value = "${var.tls_secret_name}-vm"
  }

  set {
    name  = "app.volumes[1].secret.secretName"
    value = var.tls_secret_name
  }

  set {
    name  = "app.volumeMounts[1].name"
    value = "${var.tls_secret_name}-vm"
  }

  set {
    name  = "app.volumeMounts[1].mountPath"
    value = "/etc/ssl/mongo"
  }

  set {
    name  = "app.volumeMounts[1].readOnly"
    value = "true"
  }
}
