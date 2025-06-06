
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

func ResolveCapabilities(capAdd, capDrop []string) []string {
	result := make([]string, len(defaultCapabilities))
	copy(result, defaultCapabilities)
	
	for _, cap := range capAdd {
		if !slices.Contains(result, cap) {
			result = append(result, cap)
		}
	}
	
	for _, cap := range capDrop {
		result = slices.DeleteFunc(result, func(c string) bool {
			return c == cap
		})
	}
	
	return result
}

func FormatCapabilitiesForSetpriv(caps []string) string {
	if len(caps) == 0 {
		return "-all"
	}
	
	formatted := make([]string, len(caps))
	for i, cap := range caps {
		formatted[i] = "+" + strings.ToLower(strings.TrimPrefix(cap, "CAP_"))
	}
	
	return strings.Join(formatted, ",")
}
