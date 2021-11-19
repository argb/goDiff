package myers

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

func Diff(src, dst []string) {
	sec := ShortestEditScript(src, dst)
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

	fmt.Println("cleaning test")
}

/**
每一步的可行解
 */
func getTrace(src, dst []string) []map[int]int {
	var m = len(src)
	var n = len(dst)

	steps := m + n

	trace := make([]map[int]int, 0, steps)

	v0 := make(map[int]int)
	i := 0
	j := 0

	for i<m && j<n {
		if src[i] == dst[j] {
			i++
			j++
		}else {
			break
		}
	}
	v0[0] = i
	trace = append(trace, v0)
	if i == m && j == n {
		return trace
	}

loop:
	for d := 1; d <= steps; d++ {
		lastV := trace[d-1]
		v := make(map[int]int)

		for k := -d; k <= d; k+=2 {
			var x int
			switch k {
			case -d:
				x = lastV[k+1]
			case d:
				x = lastV[k-1] + 1
			default:
				if lastV[k-1] >= lastV[k+1] {
					x = lastV[k-1] + 1
				}else {
					x = lastV[k+1]
				}
			}

			y := x - k
			for x<m && y<n {
				if src[x] == dst[y] {
					x ++
					y ++
				}else {
					break
				}
			}
			v[k] = x

			if x == m && y == n {
				trace = append(trace, v)
				break loop
			}
		}
		trace = append(trace, v)
	}
	printTrace(trace)
	return trace
}

// ShortestEditScript /**
/**
最终可行解
 */
func ShortestEditScript(src, dst []string) []Operation {
	trace := getTrace(src, dst)
	var (
		prevK int
		m = len(src)
		n = len(dst)
	)
	x := m
	y := n

	var operations []Operation

	for d := len(trace) -1; d > 0; d -- {
		k := x - y
		lastV := trace[d-1]
		switch k {
		case -d: // 向下
			prevK = k+1
		case d:
			prevK = k-1
		default:
			// 向右
			if lastV[k-1] >= lastV[k+1] {
				prevK = k - 1
			}else {
				prevK = k + 1
			}
		}
		lastX := lastV[prevK]
		lastY := lastX - prevK

		// handle diagonal
		for x>lastX && y>lastY {
			operations = append(operations, MOV)
			x--
			y--
		}
		if x == lastX {
			operations = append(operations, ADD)
		}else {
			operations = append(operations, DEL)
		}
		x, y = lastX, lastY
	}
	fmt.Printf("x:%d, y:%d\n", x, y)
	// d0, 坐滑梯
	for x>0 && y>0 {
		operations = append(operations, MOV)
		x --
		y --
	}

	return reverse(operations)
}

func reverse(ops []Operation) []Operation {
	l := len(ops)
	middle := l/2

	for i:=0; i < middle; i++ {
		ops[i], ops[l - i - 1] = ops[l - i -1 ] , ops[i]
	}

	return ops
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

func Lcs(src []string, scripts []Operation) []string {
	var lcs [] string
	srcIndex := 0
	for _, script := range scripts { // for range 对类型的处理行为有点怪异，直觉上来说我这里期望的Operation自定义类型，但结果是int
		script := Operation(script)
		if script == DEL {
			srcIndex ++
		}else if script == MOV { // 因为是公共子序列，所以从任何一个原始序列获取都是一样的
			lcs = append(lcs, src[srcIndex])
			srcIndex ++
		}
	}

	return lcs
}

func printTrace(trace []map[int]int) {
	for d := 0; d < len(trace); d++ {
		fmt.Printf("d = %d:\n", d)
		v := trace[d]
		for k := -d; k <= d; k += 2 {
			x := v[k]
			y := x - k
			fmt.Printf("  k = %2d: (%d, %d)\n", k, x, y)
		}
	}
}

func ShowLCS(src, target []string) {
	ses := ShortestEditScript(src, target)
	//fmt.Println(ses)
	lcs := Lcs(src, ses)
	fmt.Println()
	fmt.Println("A Longest common subsequence:")
	for _, item := range lcs {
		fmt.Printf("%s \n", item)
	}
}