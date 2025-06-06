

package shim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCapabilities_Default(t *testing.T) {
	defer purge()
	setup()
	run("nomad job run ./jobs/capabilities-default.hcl")
	wait(`Status\s+running`)
	
	output := logs("capabilities-default")
	require.Contains(t, output, "cap_chown")
	require.Contains(t, output, "cap_dac_override")
	require.Contains(t, output, "cap_fowner")
	require.NotContains(t, output, "cap_sys_time")
}

func TestCapabilities_Add(t *testing.T) {
	defer purge()
	setup()
	run("nomad job run ./jobs/capabilities-add.hcl")
	wait(`Status\s+running`)
	
	output := logs("capabilities-add")
	require.Contains(t, output, "cap_chown")
	require.Contains(t, output, "cap_sys_time")
}

func TestCapabilities_Drop(t *testing.T) {
	defer purge()
	setup()
	run("nomad job run ./jobs/capabilities-drop.hcl")
	wait(`Status\s+running`)
	
	output := logs("capabilities-drop")
	require.NotContains(t, output, "cap_chown")
	require.Contains(t, output, "cap_dac_override")
}
