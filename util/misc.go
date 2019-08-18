package util

import "strings"

func ContainString(s []string, i string) bool {
	for _, a := range s {
		if strings.ToLower(a) == i {
			return true
		}
	}
	return false
}
