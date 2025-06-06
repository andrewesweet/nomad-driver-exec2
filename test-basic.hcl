job "test-basic" {
  datacenters = ["dc1"]
  type        = "batch"

  group "test" {
    task "basic-test" {
      driver = "exec2"

      config {
        command = "setpriv"
        args    = ["--dump"]
      }
    }
  }
}
