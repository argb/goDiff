package main

import (
	"diff/lcs"
	"diff/myers"
	"fmt"
	"strings"
	"time"
)

func main() {
	originStr := "Recent work improves upon the basic O(N2 ) time arbitrary alphabet algorithm by being sensitive to other problem\nsize parameters."
	modifiedStr := "bcdexfgz"

	originTokens := strings.Split(originStr, "")
	modifiedTokens := strings.Split(modifiedStr, "")

	//myersExample(originTokens, modifiedTokens)

	//lcsExample(originTokens, modifiedTokens)

	// diff with myers algorithm
	startTime := time.Now()
	scripts := myers.GetReadableScripts(myers.ShortestEditScript(originTokens, modifiedTokens))
	fmt.Println(scripts)
	d := time.Now().Sub(startTime)
	fmt.Printf("time spent(myers): %d microseconds(%d milliseconds)", d.Microseconds(), d.Milliseconds())

	// diff with the algorithm based on lcs
	startTime = time.Now()
	scripts = lcs.ShortestEditScript(originTokens, modifiedTokens, lcs.Lsc(originTokens, modifiedTokens))
	fmt.Println(scripts)
	d = time.Now().Sub(startTime)
	fmt.Printf("time spent(lcs): %d microseconds(%d milliseconds)", d.Microseconds(), d.Milliseconds())
}

func myersExample(originTokens, modifiedTokens []string) {
	//myers.Diff(originTokens, modifiedTokens)
	myers.ShowLCS(originTokens, modifiedTokens)
}

func lcsExample(src, target []string) {
	lcsArr := lcs.Lsc(src, target)
	lcs.Show(lcsArr)
}
