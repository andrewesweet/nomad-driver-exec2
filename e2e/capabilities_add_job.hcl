job "test-capabilities-add" {
  datacenters = ["dc1"]
  type        = "batch"

  group "test" {
    task "capabilities-test" {
      driver = "exec2"

      config {
        command  = "setpriv"
        args     = ["--dump"]
        cap_add  = ["CAP_SYS_TIME"]
      }
    }
  }
}
