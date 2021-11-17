package main

import (
	"diff/diff"
	"strings"
)

func main() {
	originStr := "abcdefg"
	modifiedStr := "33axcdeg999"

	originTokens := strings.Split(originStr, "")
	modifiedTokens := strings.Split(modifiedStr, "")
	diff.Diff(originTokens, modifiedTokens)

}
