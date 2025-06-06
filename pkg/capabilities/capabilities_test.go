
package capabilities

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResolveCapabilities(t *testing.T) {
	tests := []struct {
		name     string
		capAdd   []string
		capDrop  []string
		expected []string
	}{
		{
			name:     "default capabilities only",
			capAdd:   nil,
			capDrop:  nil,
			expected: defaultCapabilities,
		},
		{
			name:     "add capability",
			capAdd:   []string{"CAP_SYS_ADMIN"},
			capDrop:  nil,
			expected: append(defaultCapabilities, "CAP_SYS_ADMIN"),
		},
		{
			name:     "drop capability",
			capAdd:   nil,
			capDrop:  []string{"CAP_CHOWN"},
			expected: []string{"CAP_DAC_OVERRIDE", "CAP_FOWNER", "CAP_FSETID", "CAP_KILL", "CAP_NET_BIND_SERVICE", "CAP_SETFCAP", "CAP_SETGID", "CAP_SETPCAP", "CAP_SETUID", "CAP_SYS_CHROOT"},
		},
		{
			name:     "drop takes precedence over add",
			capAdd:   []string{"CAP_CHOWN"},
			capDrop:  []string{"CAP_CHOWN"},
			expected: []string{"CAP_DAC_OVERRIDE", "CAP_FOWNER", "CAP_FSETID", "CAP_KILL", "CAP_NET_BIND_SERVICE", "CAP_SETFCAP", "CAP_SETGID", "CAP_SETPCAP", "CAP_SETUID", "CAP_SYS_CHROOT"},
		},
		{
			name:     "normalize capability names",
			capAdd:   []string{"sys_admin", "CAP_NET_ADMIN"},
			capDrop:  nil,
			expected: append(defaultCapabilities, "CAP_SYS_ADMIN", "CAP_NET_ADMIN"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ResolveCapabilities(tt.capAdd, tt.capDrop)
			require.ElementsMatch(t, tt.expected, result)
		})
	}
}

func TestNormalizeCapability(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"chown", "CAP_CHOWN"},
		{"CAP_CHOWN", "CAP_CHOWN"},
		{"sys_admin", "CAP_SYS_ADMIN"},
		{"CAP_SYS_ADMIN", "CAP_SYS_ADMIN"},
		{"net_bind_service", "CAP_NET_BIND_SERVICE"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := normalizeCapability(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatCapabilitiesForSetpriv(t *testing.T) {
	tests := []struct {
		name     string
		caps     []string
		expected string
	}{
		{
			name:     "empty capabilities",
			caps:     []string{},
			expected: "",
		},
		{
			name:     "single capability",
			caps:     []string{"CAP_CHOWN"},
			expected: "+chown",
		},
		{
			name:     "multiple capabilities",
			caps:     []string{"CAP_CHOWN", "CAP_FOWNER"},
			expected: "+chown,+fowner",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatCapabilitiesForSetpriv(tt.caps)
			require.Equal(t, tt.expected, result)
		})
	}
}
