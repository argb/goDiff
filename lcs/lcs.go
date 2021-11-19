package lcs

import (
	"fmt"
)

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

func ShortestEditScript(src, target, lcs []string) []string {
	var scripts []string
	var j int
	var k int
	for i :=0; i<len(lcs); i ++ {
		if lcs[i] == src[j] && lcs[i] == target[k] {
			scripts = append(scripts, "MOV")
			j++
			k++
		}else{
			for j<len(src) {
				if lcs[i] != src[j] {
					scripts = append(scripts, "DEL")
					j ++
				}else {
					break
				}
			}
			for k<len(target) {
				if lcs[i] != target[k] {
					scripts = append(scripts, "ADD")
					k++
				}else {
					break
				}
			}

			scripts = append(scripts, "MOV")
			j++
			k++
		}
	}
	for j < len(src) {
		scripts = append(scripts, "DEL")
		j++
	}
	for k < len(target) {
		scripts = append(scripts, "ADD")
		k++
	}

	return scripts
}