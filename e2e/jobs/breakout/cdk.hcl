# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "exploit" {
  type        = string
  description = "CDK exploit technique to run"
}

variable "args" {
  type        = string
  default     = ""
  description = "Additional arguments for the exploit"
}

job "cdk-${var.exploit}" {
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
        destination = "local/cdk"
        perms       = "755"
        data        = <<EOH
{{ with file "../../breakout-tools/cdk" }}{{ . | base64Decode }}{{ end }}
        EOH
      }

      template {
        destination = "local/run_exploit.sh"
        perms       = "755"
        data        = <<EOH
#!/bin/bash
set -e

echo "Starting CDK exploit: ${var.exploit}"
echo "Arguments: ${var.args}"

# Run the CDK exploit
if [ -n "${var.args}" ]; then
    ./cdk run ${var.exploit} ${var.args} 2>&1 || {
        echo "CDK exploit ${var.exploit} failed with exit code $?"
        exit 1
    }
else
    ./cdk run ${var.exploit} 2>&1 || {
        echo "CDK exploit ${var.exploit} failed with exit code $?"
        exit 1
    }
fi

echo "CDK exploit ${var.exploit} completed"
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
