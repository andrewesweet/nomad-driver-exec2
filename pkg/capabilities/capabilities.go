package capabilities

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/hashicorp/go-set/v2"
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
	if runtime.GOOS != "linux" {
		return nil
	}
	
	caps := set.From(defaultCapabilities)
	
	for _, cap := range capAdd {
		caps.Insert(normalizeCapability(cap))
	}
	
	for _, cap := range capDrop {
		caps.Remove(normalizeCapability(cap))
	}
	
	return caps.Slice()
}

func normalizeCapability(cap string) string {
	cap = strings.ToUpper(cap)
	if !strings.HasPrefix(cap, "CAP_") {
		cap = "CAP_" + cap
	}
	return cap
}

func FormatCapabilitiesForSetpriv(caps []string) string {
	if len(caps) == 0 {
		return ""
	}
	
	formatted := make([]string, len(caps))
	for i, cap := range caps {
		formatted[i] = fmt.Sprintf("+%s", strings.ToLower(cap))
	}
	
	return strings.Join(formatted, ",")
}
