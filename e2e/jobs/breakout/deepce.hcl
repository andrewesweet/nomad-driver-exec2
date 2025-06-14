# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "exploit" {
  type        = string
  description = "deepce exploit type to run"
}

job "deepce-${lower(var.exploit)}" {
  type = "batch"

  constraint {
    attribute = "${attr.kernel.name}"
    value     = "linux"
  }

  group "group" {
    reschedule {
      attempts  = 0
      unlimited = false
    }

    restart {
      attempts = 0
      mode     = "fail"
    }

    task "exploit" {
      driver = "exec2"

      template {
        destination = "local/deepce.sh"
        perms       = "755"
        data        = <<EOH
{{ with file "../../breakout-tools/deepce.sh" }}{{ . }}{{ end }}
        EOH
      }

      template {
        destination = "local/run_exploit.sh"
        perms       = "755"
        data        = <<EOH
#!/bin/bash
set -e

echo "Starting deepce exploit: $${var.exploit}"

# Run the deepce exploit
./deepce.sh -e $${var.exploit} 2>&1 || {
    echo "deepce exploit $${var.exploit} failed with exit code $$?"
    exit 1
}

echo "deepce exploit $${var.exploit} completed"
        EOH
      }

      config {
        command = "/bin/bash"
        args    = ["local/run_exploit.sh"]
      }

      resources {
        cpu    = 500
        memory = 256
      }
    }
  }
}
