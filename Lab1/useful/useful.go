package useful

import "strings"

func CutBlankSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
