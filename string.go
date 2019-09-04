package go_utils

import "regexp"

func StringRemoveAllNum(string2 string) string {
	re, _ := regexp.Compile("[0-9]")
	return re.ReplaceAllString(string2, "")
}
