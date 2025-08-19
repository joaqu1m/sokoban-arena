package utils

import "strings"

func NormalizeName(name string) string {
	if !strings.Contains(name, ":") {
		name = "default:" + name
	}
	return name
}
