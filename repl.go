package main

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	slicedString := strings.Fields(text)
	return slicedString
}
