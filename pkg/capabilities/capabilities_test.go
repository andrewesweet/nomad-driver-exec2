
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
			name:     "add valid capabilities",
			capAdd:   []string{"CAP_SYS_ADMIN", "CAP_NET_ADMIN"},
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
		{"CAP_CHOWN", "CAP_CHOWN"},
		{"cap_chown", "CAP_CHOWN"},
		{"CAP_SYS_ADMIN", "CAP_SYS_ADMIN"},
		{"cap_sys_admin", "CAP_SYS_ADMIN"},
		{"CAP_NET_BIND_SERVICE", "CAP_NET_BIND_SERVICE"},
		{"chown", ""},
		{"sys_admin", ""},
		{"net_bind_service", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := normalizeCapability(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestValidateCapabilities(t *testing.T) {
	tests := []struct {
		name        string
		caps        []string
		expected    []string
		expectError bool
		errorContains string
	}{
		{
			name:        "empty capabilities",
			caps:        []string{},
			expected:    []string{},
			expectError: false,
		},
		{
			name:        "valid capabilities uppercase",
			caps:        []string{"CAP_SYS_TIME", "CAP_NET_ADMIN"},
			expected:    []string{"CAP_SYS_TIME", "CAP_NET_ADMIN"},
			expectError: false,
		},
		{
			name:        "valid capabilities lowercase",
			caps:        []string{"cap_sys_time", "cap_net_admin"},
			expected:    []string{"CAP_SYS_TIME", "CAP_NET_ADMIN"},
			expectError: false,
		},
		{
			name:        "mixed case valid capabilities",
			caps:        []string{"Cap_Sys_Time", "CAP_net_admin"},
			expected:    []string{"CAP_SYS_TIME", "CAP_NET_ADMIN"},
			expectError: false,
		},
		{
			name:        "capabilities without CAP_ prefix",
			caps:        []string{"sys_time", "net_admin"},
			expected:    nil,
			expectError: true,
			errorContains: "sys_time",
		},
		{
			name:        "unknown capabilities",
			caps:        []string{"CAP_UNKNOWN", "CAP_INVALID"},
			expected:    nil,
			expectError: true,
			errorContains: "CAP_UNKNOWN",
		},
		{
			name:        "mixed valid and invalid capabilities",
			caps:        []string{"CAP_SYS_TIME", "invalid_cap", "CAP_UNKNOWN"},
			expected:    nil,
			expectError: true,
			errorContains: "invalid_cap",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateCapabilities(tt.caps)
			if tt.expectError {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errorContains)
				require.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.ElementsMatch(t, tt.expected, result)
			}
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
