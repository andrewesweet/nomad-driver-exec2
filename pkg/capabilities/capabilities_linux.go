// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0



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
		if norm == "" {
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
