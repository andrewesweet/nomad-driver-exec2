// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build e2e

package e2e

import (
	"context"
	"regexp"
	"strings"
	"testing"

	"github.com/shoenig/test/must"
)

func TestBreakout_CDK_CapDacReadSearch(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-cap-dac-read-search")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=cap-dac-read-search",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-cap-dac-read-search")

	logs := logs2(t, ctx, "cdk-cap-dac-read-search", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK cap-dac-read-search exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK cap-dac-read-search test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK cap-dac-read-search exploit contained successfully")
}

func TestBreakout_CDK_DockerSockPwn(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-sock-pwn")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-sock-pwn",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-sock-pwn")

	logs := logs2(t, ctx, "cdk-docker-sock-pwn", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-sock-pwn exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-sock-pwn test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-sock-pwn exploit contained successfully")
}

func TestBreakout_CDK_MountDisk(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-mount-disk")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=mount-disk",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-mount-disk")

	logs := logs2(t, ctx, "cdk-mount-disk", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK mount-disk exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK mount-disk test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK mount-disk exploit contained successfully")
}

func TestBreakout_CDK_MountProcfs(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-mount-procfs")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=mount-procfs",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-mount-procfs")

	logs := logs2(t, ctx, "cdk-mount-procfs", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK mount-procfs exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK mount-procfs test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK mount-procfs exploit contained successfully")
}

func TestBreakout_CDK_MountCgroup(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-mount-cgroup")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=mount-cgroup",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-mount-cgroup")

	logs := logs2(t, ctx, "cdk-mount-cgroup", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK mount-cgroup exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK mount-cgroup test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK mount-cgroup exploit contained successfully")
}

func TestBreakout_CDK_WriteEtcPasswd(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-write-etc-passwd")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=write-etc-passwd",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-write-etc-passwd")

	logs := logs2(t, ctx, "cdk-write-etc-passwd", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK write-etc-passwd exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK write-etc-passwd test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK write-etc-passwd exploit contained successfully")
}

func TestBreakout_CDK_WriteEtcShadow(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-write-etc-shadow")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=write-etc-shadow",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-write-etc-shadow")

	logs := logs2(t, ctx, "cdk-write-etc-shadow", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK write-etc-shadow exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK write-etc-shadow test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK write-etc-shadow exploit contained successfully")
}

func TestBreakout_CDK_WriteEtcCrontab(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-write-etc-crontab")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=write-etc-crontab",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-write-etc-crontab")

	logs := logs2(t, ctx, "cdk-write-etc-crontab", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK write-etc-crontab exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK write-etc-crontab test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK write-etc-crontab exploit contained successfully")
}

func TestBreakout_CDK_WriteSSHAuthorizedKeys(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-write-ssh-authorized-keys")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=write-ssh-authorized-keys",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-write-ssh-authorized-keys")

	logs := logs2(t, ctx, "cdk-write-ssh-authorized-keys", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK write-ssh-authorized-keys exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK write-ssh-authorized-keys test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK write-ssh-authorized-keys exploit contained successfully")
}

func TestBreakout_CDK_ReverseShell(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-reverse-shell")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=reverse-shell",
		"-var=args=127.0.0.1:4444",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-reverse-shell")

	logs := logs2(t, ctx, "cdk-reverse-shell", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK reverse-shell exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK reverse-shell test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK reverse-shell exploit contained successfully")
}

func TestBreakout_CDK_K8sEscape(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-k8s-escape")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=k8s-escape",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-k8s-escape")

	logs := logs2(t, ctx, "cdk-k8s-escape", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK k8s-escape exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK k8s-escape test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK k8s-escape exploit contained successfully")
}

func TestBreakout_CDK_K8sServiceAccountToken(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-k8s-service-account-token")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=k8s-service-account-token",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-k8s-service-account-token")

	logs := logs2(t, ctx, "cdk-k8s-service-account-token", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK k8s-service-account-token exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK k8s-service-account-token test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK k8s-service-account-token exploit contained successfully")
}

func TestBreakout_CDK_K8sConfigmapToken(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-k8s-configmap-token")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=k8s-configmap-token",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-k8s-configmap-token")

	logs := logs2(t, ctx, "cdk-k8s-configmap-token", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK k8s-configmap-token exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK k8s-configmap-token test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK k8s-configmap-token exploit contained successfully")
}

func TestBreakout_CDK_K8sSecretToken(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-k8s-secret-token")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=k8s-secret-token",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-k8s-secret-token")

	logs := logs2(t, ctx, "cdk-k8s-secret-token", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK k8s-secret-token exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK k8s-secret-token test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK k8s-secret-token exploit contained successfully")
}

func TestBreakout_CDK_K8sShadowApiserver(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-k8s-shadow-apiserver")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=k8s-shadow-apiserver",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-k8s-shadow-apiserver")

	logs := logs2(t, ctx, "cdk-k8s-shadow-apiserver", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK k8s-shadow-apiserver exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK k8s-shadow-apiserver test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK k8s-shadow-apiserver exploit contained successfully")
}

func TestBreakout_CDK_DockerApiExpose(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-api-expose")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-api-expose",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-api-expose")

	logs := logs2(t, ctx, "cdk-docker-api-expose", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-api-expose exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-api-expose test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-api-expose exploit contained successfully")
}

func TestBreakout_CDK_DockerSockDeploy(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-sock-deploy")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-sock-deploy",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-sock-deploy")

	logs := logs2(t, ctx, "cdk-docker-sock-deploy", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-sock-deploy exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-sock-deploy test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-sock-deploy exploit contained successfully")
}

func TestBreakout_CDK_DockerSockCheck(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-sock-check")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-sock-check",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-sock-check")

	logs := logs2(t, ctx, "cdk-docker-sock-check", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-sock-check exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-sock-check test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-sock-check exploit contained successfully")
}

func TestBreakout_CDK_DockerHistoryCheck(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-history-check")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-history-check",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-history-check")

	logs := logs2(t, ctx, "cdk-docker-history-check", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-history-check exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-history-check test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-history-check exploit contained successfully")
}

func TestBreakout_CDK_DockerRegistryCheck(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-registry-check")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-registry-check",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-registry-check")

	logs := logs2(t, ctx, "cdk-docker-registry-check", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-registry-check exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-registry-check test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-registry-check exploit contained successfully")
}

func TestBreakout_CDK_DockerRemoteApiCheck(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-remote-api-check")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-remote-api-check",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-remote-api-check")

	logs := logs2(t, ctx, "cdk-docker-remote-api-check", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-remote-api-check exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-remote-api-check test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-remote-api-check exploit contained successfully")
}

func TestBreakout_CDK_DockerEscapeCheck(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-escape-check")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-escape-check",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-escape-check")

	logs := logs2(t, ctx, "cdk-docker-escape-check", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-escape-check exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-escape-check test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-escape-check exploit contained successfully")
}

func TestBreakout_CDK_DockerEscapePrivileged(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-escape-privileged")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-escape-privileged",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-escape-privileged")

	logs := logs2(t, ctx, "cdk-docker-escape-privileged", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-escape-privileged exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-escape-privileged test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-escape-privileged exploit contained successfully")
}

func TestBreakout_CDK_DockerEscapeCapSysAdmin(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-escape-cap-sys-admin")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-escape-cap-sys-admin",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-escape-cap-sys-admin")

	logs := logs2(t, ctx, "cdk-docker-escape-cap-sys-admin", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-escape-cap-sys-admin exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-escape-cap-sys-admin test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-escape-cap-sys-admin exploit contained successfully")
}

func TestBreakout_CDK_DockerEscapeCapSysPtrace(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-escape-cap-sys-ptrace")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-escape-cap-sys-ptrace",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-escape-cap-sys-ptrace")

	logs := logs2(t, ctx, "cdk-docker-escape-cap-sys-ptrace", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-escape-cap-sys-ptrace exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-escape-cap-sys-ptrace test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-escape-cap-sys-ptrace exploit contained successfully")
}

func TestBreakout_CDK_DockerEscapeCapDacReadSearch(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-escape-cap-dac-read-search")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-escape-cap-dac-read-search",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-escape-cap-dac-read-search")

	logs := logs2(t, ctx, "cdk-docker-escape-cap-dac-read-search", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-escape-cap-dac-read-search exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-escape-cap-dac-read-search test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-escape-cap-dac-read-search exploit contained successfully")
}

func TestBreakout_CDK_DockerEscapeCapDacOverride(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-escape-cap-dac-override")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-escape-cap-dac-override",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-escape-cap-dac-override")

	logs := logs2(t, ctx, "cdk-docker-escape-cap-dac-override", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-escape-cap-dac-override exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-escape-cap-dac-override test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-escape-cap-dac-override exploit contained successfully")
}

func TestBreakout_CDK_DockerEscapeCapSysModule(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "cdk-docker-escape-cap-sys-module")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=docker-escape-cap-sys-module",
		"./jobs/breakout/cdk.hcl")
	wait(t, ctx, "cdk-docker-escape-cap-sys-module")

	logs := logs2(t, ctx, "cdk-docker-escape-cap-sys-module", "exploit")
	
	if strings.Contains(logs, "exploit success") || strings.Contains(logs, "breakout successful") {
		t.Fatal("SECURITY BREACH: CDK docker-escape-cap-sys-module exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: CDK docker-escape-cap-sys-module test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: CDK docker-escape-cap-sys-module exploit contained successfully")
}

func TestBreakout_Deepce_Docker(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "deepce-docker")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=DOCKER",
		"./jobs/breakout/deepce.hcl")
	wait(t, ctx, "deepce-docker")

	logs := logs2(t, ctx, "deepce-docker", "exploit")
	
	if strings.Contains(logs, "Exploit Successful") || strings.Contains(logs, "Container Escaped") {
		t.Fatal("SECURITY BREACH: deepce DOCKER exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: deepce DOCKER test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: deepce DOCKER exploit contained successfully")
}

func TestBreakout_Deepce_Privileged(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "deepce-privileged")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=PRIVILEGED",
		"./jobs/breakout/deepce.hcl")
	wait(t, ctx, "deepce-privileged")

	logs := logs2(t, ctx, "deepce-privileged", "exploit")
	
	if strings.Contains(logs, "Exploit Successful") || strings.Contains(logs, "Container Escaped") {
		t.Fatal("SECURITY BREACH: deepce PRIVILEGED exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: deepce PRIVILEGED test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: deepce PRIVILEGED exploit contained successfully")
}

func TestBreakout_Deepce_Sock(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "deepce-sock")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=SOCK",
		"./jobs/breakout/deepce.hcl")
	wait(t, ctx, "deepce-sock")

	logs := logs2(t, ctx, "deepce-sock", "exploit")
	
	if strings.Contains(logs, "Exploit Successful") || strings.Contains(logs, "Container Escaped") {
		t.Fatal("SECURITY BREACH: deepce SOCK exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: deepce SOCK test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: deepce SOCK exploit contained successfully")
}

func TestBreakout_Deepce_CVE(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "deepce-cve")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=CVE",
		"./jobs/breakout/deepce.hcl")
	wait(t, ctx, "deepce-cve")

	logs := logs2(t, ctx, "deepce-cve", "exploit")
	
	if strings.Contains(logs, "Exploit Successful") || strings.Contains(logs, "Container Escaped") {
		t.Fatal("SECURITY BREACH: deepce CVE exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: deepce CVE test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: deepce CVE exploit contained successfully")
}

func TestBreakout_Deepce_SysModule(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "deepce-sys-module")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=SYS_MODULE",
		"./jobs/breakout/deepce.hcl")
	wait(t, ctx, "deepce-sys-module")

	logs := logs2(t, ctx, "deepce-sys-module", "exploit")
	
	if strings.Contains(logs, "Exploit Successful") || strings.Contains(logs, "Container Escaped") {
		t.Fatal("SECURITY BREACH: deepce SYS_MODULE exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: deepce SYS_MODULE test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: deepce SYS_MODULE exploit contained successfully")
}

func TestBreakout_Deepce_General(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "deepce-general")()

	_ = run(t, ctx, "nomad", "job", "run", 
		"-var=exploit=GENERAL",
		"./jobs/breakout/deepce.hcl")
	wait(t, ctx, "deepce-general")

	logs := logs2(t, ctx, "deepce-general", "exploit")
	
	if strings.Contains(logs, "Exploit Successful") || strings.Contains(logs, "Container Escaped") {
		t.Fatal("SECURITY BREACH: deepce GENERAL exploit succeeded - sandbox failed to contain attack")
	}
	
	if strings.Contains(logs, "error") && !strings.Contains(logs, "permission denied") && !strings.Contains(logs, "operation not permitted") {
		t.Fatalf("EXECUTION ERROR: deepce GENERAL test failed to execute properly: %s", logs)
	}
	
	t.Log("SANDBOX EFFECTIVE: deepce GENERAL exploit contained successfully")
}
