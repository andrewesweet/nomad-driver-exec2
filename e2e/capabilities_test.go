

package shim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCapabilities_Default(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "capabilities-default")()
	
	_ = run(t, ctx, "nomad", "job", "run", "./jobs/capabilities-default.hcl")
	
	statusOutput := run(t, ctx, "nomad", "job", "status", "capabilities-default")
	alloc := allocFromJobStatus(t, statusOutput)
	
	output := logs(t, ctx, alloc)
	require.Contains(t, output, "cap_chown")
	require.Contains(t, output, "cap_dac_override")
	require.Contains(t, output, "cap_fowner")
	require.NotContains(t, output, "cap_sys_time")
}

func TestCapabilities_Add(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "capabilities-add")()
	
	_ = run(t, ctx, "nomad", "job", "run", "./jobs/capabilities-add.hcl")
	
	statusOutput := run(t, ctx, "nomad", "job", "status", "capabilities-add")
	alloc := allocFromJobStatus(t, statusOutput)
	
	output := logs(t, ctx, alloc)
	require.Contains(t, output, "cap_chown")
	require.Contains(t, output, "cap_sys_time")
}

func TestCapabilities_Drop(t *testing.T) {
	ctx := setup(t)
	defer purge(t, ctx, "capabilities-drop")()
	
	_ = run(t, ctx, "nomad", "job", "run", "./jobs/capabilities-drop.hcl")
	
	statusOutput := run(t, ctx, "nomad", "job", "status", "capabilities-drop")
	alloc := allocFromJobStatus(t, statusOutput)
	
	output := logs(t, ctx, alloc)
	require.NotContains(t, output, "cap_chown")
	require.Contains(t, output, "cap_dac_override")
}
