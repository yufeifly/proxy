package httputils

import (
	"strings"
)

// BoolValue transforms a form value in different formats into a boolean type.
func BoolValue(k string) bool {
	s := strings.ToLower(strings.TrimSpace(k))
	return !(s == "" || s == "0" || s == "no" || s == "false" || s == "none")
}
