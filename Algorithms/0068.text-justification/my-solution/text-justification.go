package problem0068

import (
	"strings"
)

func bool2int(can bool) int {
	if can {
		return 1
	}else {
		return 0
	}
}

func addSpaces(i int, spaceCnt int, maxWidth int, is_last bool, s string) string {
	spaces := 0
	if i < spaceCnt {
		if is_last {
			spaces = 1
		}else {
			spaces = (maxWidth  / spaceCnt) + bool2int(i < maxWidth % spaceCnt)
		}
	}
	return s + strings.Repeat(" ", spaces)
}

func connect(words []string, maxWidth int, begin int, end int, length int, is_last bool) string {
	var s string
	n := end - begin
	for i:=0; i<n; i++ {
		s += words[begin+i]
		s = addSpaces(i, n-1, maxWidth-length, is_last, s)
	}
	if len(s) < maxWidth {
		s += strings.Repeat(" ", maxWidth-len(s))
	}
	return s
}

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	n := len(words)
	begin, length := 0, 0

	for i:=0; i < n; i++ {
		if (length + len(words[i]) + (i-begin) > maxWidth) {
			res = append(res, connect(words, maxWidth, begin, i, length, false))
			begin, length = i, 0
		}
		length += len(words[i])
	}
	res = append(res, connect(words, maxWidth, begin, n, length, true))
	return res
}
