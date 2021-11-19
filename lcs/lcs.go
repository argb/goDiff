package lcs

import (
	"fmt"
	"strings"
)


type Operation byte

const (
	DEL Operation= iota + 1
	ADD
	MOV
)

var OpMap = map[Operation]string{
	DEL: "DEL",
	ADD: "ADD",
	MOV: "MOV",
}

var colors = map[Operation]string{
	ADD: "\033[32m",
	DEL: "\033[31m",
	MOV:   "\033[39m",
}

func Lsc(src, target []string) []string {
	m := len(src) // x 轴坐标，列数
	n := len(target) // y 轴坐标， 行数

	var lcs [] string

	matrix := make([][]int,n+1)

	for j := 0; j <= n; j ++ {
		line := make([]int, m+1)
		line[0] = 0

		matrix[j] = line
	}

	for i := 0; i <= m; i++ {
		matrix[0][i] = 0
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i ++ {
			if src[i-1] == target[j-1] {
				matrix[j][i] = matrix[j-1][i-1] + 1
			}else{
				matrix[j][i] = max(matrix[j-1][i], matrix[j][i-1])
			}
		}
	}

	//visibleDpTable(src, target, matrix)

	i := m
	j := n
	for i>=1 && j>=1 {
		if src[i-1] == target[j-1] {
			lcs = append(lcs, src[i-1])
			i --
			j --
		}else {
			if matrix[j-1][i] > matrix[j][i-1] {
				j --
			}else {
				i --
			}
		}
	}

	return reverse(lcs)
}

func visibleDpTable(src, target []string, matrix [][]int)  {
	m := len(src)
	n := len(target)
	for i:=0; i <= m; i ++ {
		if i== 0 {
			fmt.Printf("      ")
		}else {
			fmt.Printf("%s ", src[i-1])
		}
	}
	fmt.Println()
	for i:=0; i <= m; i ++ {
		if i== 0 {
			fmt.Printf("  - -")
		}else {
			fmt.Printf(" -")
		}
	}
	fmt.Println()
	for j := 0; j <= n; j ++ {
		if j == 0 {
			fmt.Printf("  | ")
		}else {
			fmt.Printf("%s | ", target[j-1])
		}
		for i:=0; i <= m; i ++ {
			fmt.Printf("%d ", matrix[j][i])
		}
		fmt.Println()
	}
}

func reverse(lcs []string) []string {
	l := len(lcs)
	for i:=0; i< l/2; i ++ {
		lcs[i], lcs[l-i-1] = lcs[l-i-1], lcs[i]
	}

	return lcs
}

func max(x, y int) int {
	if x >= y {
		return x
	}else {
		return y
	}
}

func Show(lcs [] string) {
	fmt.Println()
	fmt.Println("A Longest common subsequence:")
	for _, item := range lcs {
		fmt.Printf("%s \n", item)
	}
}

func ShortestEditScript(src, target, lcs []string) []Operation {
	var scripts []Operation
	var j int
	var k int
	for i :=0; i<len(lcs); i ++ {
		if lcs[i] == src[j] && lcs[i] == target[k] {
			scripts = append(scripts, MOV)
			j++
			k++
		}else{
			for j<len(src) {
				if lcs[i] != src[j] {
					scripts = append(scripts, DEL)
					j ++
				}else {
					break
				}
			}
			for k<len(target) {
				if lcs[i] != target[k] {
					scripts = append(scripts, ADD)
					k++
				}else {
					break
				}
			}

			scripts = append(scripts, MOV)
			j++
			k++
		}
	}
	for j < len(src) {
		scripts = append(scripts, DEL)
		j++
	}
	for k < len(target) {
		scripts = append(scripts, ADD)
		k++
	}

	return scripts
}

func readableOp(op Operation) string {
	script, ok := OpMap[op]
	if ok {
		return script
	}else {
		return "unknown"
	}
}

func GetReadableScripts(ops []Operation) []string {
	l := len(ops)
	scripts := make([]string, l)

	for i:=0; i< l; i++ {
		scripts[i] = readableOp(ops[i])
	}

	return scripts
}

func Diff(src, dst []string) {
	lcs := Lsc(src, dst)
	sec := ShortestEditScript(src, dst, lcs)
	var result [] string
	//var op Operation
	l := len(sec)
	var j, k int
	for i:=0; i<l; i++ {
		op := sec[i]
		switch op {
		case MOV:
			modification := colors[op]+"  " +src[j]
			result = append(result, modification)
			j ++
			k ++
		case DEL:
			modification := colors[op]+"- " +src[j]
			result = append(result, modification)
			j++
		case ADD:
			modification := colors[op]+"+ " + dst[k]
			result = append(result, modification)
			k++
		}
	}


	//fmt.Println(GetReadableScripts(sec))
	fmt.Println(strings.Join(result, "\n"))

	// 清除命令行所控制字符效果
	fmt.Println("\033[0m")

	//fmt.Println("cleaning test")
}