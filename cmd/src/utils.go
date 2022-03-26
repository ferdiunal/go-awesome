package src

import (
	"strings"
)

func SpaceCleaner(str string) string {
	return strings.Join(strings.Fields(str), " ")
}
