

package capabilities

import (
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



func normalizeCapability(cap string) string {
	cap = strings.ToUpper(cap)
	if !strings.HasPrefix(cap, "CAP_") {
		cap = "CAP_" + cap
	}
	return cap
}

func ResolveCapabilities(capAdd, capDrop []string) []string {
	result := make([]string, len(defaultCapabilities))
	copy(result, defaultCapabilities)
	
	// Normalize and add capabilities
	for _, cap := range capAdd {
		normalized := normalizeCapability(cap)
		if !slices.Contains(result, normalized) {
			result = append(result, normalized)
		}
	}
	
	// Normalize and drop capabilities (drop takes precedence)
	for _, cap := range capDrop {
		normalized := normalizeCapability(cap)
		result = slices.DeleteFunc(result, func(c string) bool {
			return c == normalized
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
