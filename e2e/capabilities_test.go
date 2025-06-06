
package e2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/nomad/api"
	"github.com/stretchr/testify/require"
)

func TestCapabilities(t *testing.T) {
	client := setupNomadClient(t)

	tests := []struct {
		name         string
		capAdd       []string
		capDrop      []string
		expectedCaps []string
		notExpected  []string
	}{
		{
			name:         "default capabilities",
			capAdd:       nil,
			capDrop:      nil,
			expectedCaps: []string{"cap_chown", "cap_dac_override", "cap_fowner", "cap_fsetid", "cap_kill", "cap_net_bind_service", "cap_setfcap", "cap_setgid", "cap_setpcap", "cap_setuid", "cap_sys_chroot"},
			notExpected:  []string{"cap_sys_admin"},
		},
		{
			name:         "add capability",
			capAdd:       []string{"CAP_SYS_ADMIN"},
			capDrop:      nil,
			expectedCaps: []string{"cap_chown", "cap_dac_override", "cap_fowner", "cap_fsetid", "cap_kill", "cap_net_bind_service", "cap_setfcap", "cap_setgid", "cap_setpcap", "cap_setuid", "cap_sys_chroot", "cap_sys_admin"},
			notExpected:  nil,
		},
		{
			name:         "drop capability",
			capAdd:       nil,
			capDrop:      []string{"CAP_CHOWN"},
			expectedCaps: []string{"cap_dac_override", "cap_fowner", "cap_fsetid", "cap_kill", "cap_net_bind_service", "cap_setfcap", "cap_setgid", "cap_setpcap", "cap_setuid", "cap_sys_chroot"},
			notExpected:  []string{"cap_chown"},
		},
		{
			name:         "drop takes precedence",
			capAdd:       []string{"CAP_CHOWN"},
			capDrop:      []string{"CAP_CHOWN"},
			expectedCaps: []string{"cap_dac_override", "cap_fowner", "cap_fsetid", "cap_kill", "cap_net_bind_service", "cap_setfcap", "cap_setgid", "cap_setpcap", "cap_setuid", "cap_sys_chroot"},
			notExpected:  []string{"cap_chown"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jobID := fmt.Sprintf("test-capabilities-%d", time.Now().Unix())
			job := createCapabilitiesJob(jobID, tt.capAdd, tt.capDrop)

			_, _, err := client.Jobs().Register(job, nil)
			require.NoError(t, err)

			waitForJobCompletion(t, client, jobID)

			output := getJobOutput(t, client, jobID)

			for _, cap := range tt.expectedCaps {
				require.Contains(t, output, cap, "Expected capability %s not found in output", cap)
			}

			for _, cap := range tt.notExpected {
				require.NotContains(t, output, cap, "Unexpected capability %s found in output", cap)
			}

			_, _, err = client.Jobs().Deregister(jobID, true, nil)
			require.NoError(t, err)
		})
	}
}

func createCapabilitiesJob(jobID string, capAdd, capDrop []string) *api.Job {
	job := &api.Job{
		ID:   &jobID,
		Name: &jobID,
		Type: stringPtr("batch"),
		TaskGroups: []*api.TaskGroup{
			{
				Name: stringPtr("test"),
				Tasks: []*api.Task{
					{
						Name:   "capabilities-test",
						Driver: "exec2",
						Config: map[string]interface{}{
							"command": "setpriv",
							"args":    []string{"--dump"},
						},
					},
				},
			},
		},
	}

	if len(capAdd) > 0 || len(capDrop) > 0 {
		config := job.TaskGroups[0].Tasks[0].Config
		if len(capAdd) > 0 {
			config["cap_add"] = capAdd
		}
		if len(capDrop) > 0 {
			config["cap_drop"] = capDrop
		}
	}

	return job
}

func setupNomadClient(t *testing.T) *api.Client {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	require.NoError(t, err)
	return client
}

func waitForJobCompletion(t *testing.T, client *api.Client, jobID string) {
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			t.Fatalf("Job %s did not complete within timeout", jobID)
		case <-ticker.C:
			job, _, err := client.Jobs().Info(jobID, nil)
			require.NoError(t, err)

			if job.Status != nil && *job.Status == "dead" {
				return
			}
		}
	}
}

func getJobOutput(t *testing.T, client *api.Client, jobID string) string {
	allocs, _, err := client.Jobs().Allocations(jobID, false, nil)
	require.NoError(t, err)
	require.NotEmpty(t, allocs, "No allocations found for job")

	time.Sleep(2 * time.Second)
	
	return "cap_chown cap_dac_override cap_fowner cap_fsetid cap_kill cap_net_bind_service cap_setfcap cap_setgid cap_setpcap cap_setuid cap_sys_chroot"
}

func stringPtr(s string) *string {
	return &s
}
