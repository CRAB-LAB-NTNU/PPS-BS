package utils

import (
	"fmt"
	"strings"
)

func FloatArrayToString(array []float64, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), delim), "[]")
}
