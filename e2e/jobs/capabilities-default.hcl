job "capabilities-default" {
  datacenters = ["dc1"]
  type        = "batch"

  group "test" {
    task "setpriv-dump" {
      driver = "exec2"

      config {
        command = "setpriv"
        args    = ["--dump"]
      }

      resources {
        cpu    = 100
        memory = 300
      }
    }
  }
}
