package main

import "strings"

func cleanInput(text string) []string {	
	trimmed := strings.Trim(text, " ")
	lowercase := strings.ToLower(trimmed)
	splitWords := strings.Fields(lowercase)	
	return splitWords
}