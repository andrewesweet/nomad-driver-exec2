

package capabilities

import (
	"fmt"
	"slices"
	"strings"
)

var defaultCapabilities = []string{
	"CAP_CHOWN",
	"CAP_DAC_OVERRIDE", 
	"CAP_FOWNER",
	"CAP_FSETID",
	"CAP_KILL",
	"CAP_NET_BIND_SERVICE",
	"CAP_SETFCAP",
	"CAP_SETGID",
	"CAP_SETPCAP",
	"CAP_SETUID",
	"CAP_SYS_CHROOT",
}

var validCapabilities = map[string]bool{
	"CAP_CHOWN":            true,
	"CAP_DAC_OVERRIDE":     true,
	"CAP_DAC_READ_SEARCH":  true,
	"CAP_FOWNER":           true,
	"CAP_FSETID":           true,
	"CAP_KILL":             true,
	"CAP_SETGID":           true,
	"CAP_SETUID":           true,
	"CAP_SETPCAP":          true,
	"CAP_LINUX_IMMUTABLE":  true,
	"CAP_NET_BIND_SERVICE": true,
	"CAP_NET_BROADCAST":    true,
	"CAP_NET_ADMIN":        true,
	"CAP_NET_RAW":          true,
	"CAP_IPC_LOCK":         true,
	"CAP_IPC_OWNER":        true,
	"CAP_SYS_MODULE":       true,
	"CAP_SYS_RAWIO":        true,
	"CAP_SYS_CHROOT":       true,
	"CAP_SYS_PTRACE":       true,
	"CAP_SYS_PACCT":        true,
	"CAP_SYS_ADMIN":        true,
	"CAP_SYS_BOOT":         true,
	"CAP_SYS_NICE":         true,
	"CAP_SYS_RESOURCE":     true,
	"CAP_SYS_TIME":         true,
	"CAP_SYS_TTY_CONFIG":   true,
	"CAP_MKNOD":            true,
	"CAP_LEASE":            true,
	"CAP_AUDIT_WRITE":      true,
	"CAP_AUDIT_CONTROL":    true,
	"CAP_SETFCAP":          true,
	"CAP_MAC_OVERRIDE":     true,
	"CAP_MAC_ADMIN":        true,
	"CAP_SYSLOG":           true,
	"CAP_WAKE_ALARM":       true,
	"CAP_BLOCK_SUSPEND":    true,
	"CAP_AUDIT_READ":       true,
}



func normalizeCapability(cap string) string {
	cap = strings.ToUpper(cap)
	if !strings.HasPrefix(cap, "CAP_") {
		return ""
	}
	return cap
}

func ValidateCapabilities(caps []string) ([]string, error) {
	if len(caps) == 0 {
		return caps, nil
	}
	
	var normalized []string
	var invalid []string
	
	for _, cap := range caps {
		norm := normalizeCapability(cap)
		if norm == "" || !validCapabilities[norm] {
			invalid = append(invalid, cap)
		} else {
			normalized = append(normalized, norm)
		}
	}
	
	if len(invalid) > 0 {
		return nil, fmt.Errorf("%v", invalid)
	}
	
	return normalized, nil
}

func ResolveCapabilities(capAdd, capDrop []string) []string {
	result := make([]string, len(defaultCapabilities))
	copy(result, defaultCapabilities)
	
	// Add capabilities (assume already validated and normalized)
	for _, cap := range capAdd {
		if !slices.Contains(result, cap) {
			result = append(result, cap)
		}
	}
	
	// Drop capabilities (drop takes precedence, assume already validated and normalized)
	for _, cap := range capDrop {
		result = slices.DeleteFunc(result, func(c string) bool {
			return c == cap
		})
	}
	
	return result
}

func FormatCapabilitiesForSetpriv(caps []string) string {
	if len(caps) == 0 {
		return ""
	}
	
	formatted := make([]string, len(caps))
	for i, cap := range caps {
		formatted[i] = "+" + strings.ToLower(strings.TrimPrefix(cap, "CAP_"))
	}
	
	return strings.Join(formatted, ",")
}
