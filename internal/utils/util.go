package utils

import "strings"

func StrPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func StringToBooleanMapper(s string) bool {
	//not handling nil cases for now to keep it simple
	return strings.TrimSpace(strings.ToLower(s)) == "yes"
}

func CheckIfTrailExists(s string) bool {
	return strings.TrimSpace(strings.ToLower(s)) != "no"
}
