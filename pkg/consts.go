package pkg

import "strings"

type ValidationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var letters = "abcdefghijklmnopqrstuvwxyz"
var LETTERS = []rune(letters + strings.ToUpper(letters))
