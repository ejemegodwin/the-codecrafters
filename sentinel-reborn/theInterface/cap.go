package theInterface

import (
	"strings"
)

func ParseModifier(t string) string {
	t = strings.TrimSpace(t)
	index := strings.Index(t, "(")
	if strings.Contains(t, "(up)") {
		t = strings.ToUpper(t)
	}
	if strings.Contains(t, "(low)") {
		t = strings.ToLower(t)
	}
	if strings.Contains(t, "(cap)") {
		t = strings.Title(t)
	}
	return t[:index]
}