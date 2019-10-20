package problem0060

import (
	"fmt"
	"strings"
)


func nextPerm(p []int){ 
	for i := len(p) -1; i >= 0; i--{
		if i == 0 || p[i] < len(p) -i -1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result

}

func getPermutation(n int, k int) string {
	orig := []int{}
	for i:=1; i <= n; i++ {
		orig = append(orig, i)
	}
	j := 0
	res := []int{}
	for p:= make([]int, len(orig)); j < k; nextPerm(p) {
		res = getPerm(orig, p)
		j++
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), ""), "[]")
}

