job "chintya-echo" {
  datacenters = ["dc1"]
  type = "service"

  group "web" {
    count = 1

    network {
      port "chintya-http" {
        to = 1323
      }
    }

    task "chintya-echo" {
      driver = "docker"

      config {
        image = "chintyas/first-demo"
        ports = ["chintya-http"]
      }

      resources {
        cpu    = 100
        memory = 128
      }
    }

    service {
      name = "chintya-echo"
      port = "chintya-http"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.chintya-http.rule=Host(\"chintyaecho.cupang.efishery.ai\")",
      ]
      check {
        port        = "chintya-http"
        type        = "tcp"
        interval    = "15s"
        timeout     = "14s"
      }
    }

  }
}