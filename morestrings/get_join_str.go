package morestrings

import (
	"strings"
)

func GetStr() string {
	s := []string{"Peter", "Pan", "zar"}
	return strings.Join(s, ", ")
}
