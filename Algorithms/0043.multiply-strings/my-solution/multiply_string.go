package problem0043

import "strings"
import "fmt"
import "strconv"

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers) - 1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func multiply1(num1 string, num2 string) string {
	if num1 == "0" || num2 =="0" {
		return "0"
	}

	b1 := []byte(num1)
	b2 := []byte(num2)
	tmp := make([]int, len(num1) + len(num2))
	for i:=0; i < len(b1); i++ {
		for j:=0; j < len(b2); j++ {
			tmp[i+j+1] += int(b1[i]-'0') * int(b2[j]-'0')
		}
	}

	for i:= len(tmp) - 1; i > 0; i-- {
		tmp[i-1] += tmp[i] / 10
		tmp[i] = tmp[i] % 10
	}
	if tmp[0] == 0 {
		tmp = tmp[1:]
	}
	res := make([]byte, len(tmp))
	for i:= 0; i < len(tmp); i++ {
		res[i] = '0' + byte(tmp[i])
	}
	return string(res)
}

func stringReverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b) - 1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func multiply(num1 string, num2 string) string {
	/*
	:type num1: str
	:type num2: str
	:rtype: str
	*/
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1, num2 = stringReverse(num1), stringReverse(num2)
	tmp  := make([]int, len(num1)+len(num2))
	for i:=0; i<len(num1); i++ {
		for j:=0; j<len(num2);j++{	
			tmp[i+j] += (int(num1[i])-48) * (int(num2[j])-48)
			//fmt.Println(tmp[i+j])
			tmp[i+j+1] += tmp[i+j] / 10
			tmp[i+j] %= 10
		}
	}

	// skip leading 0s
	i := len(tmp) - 1
	for i > 0 && tmp[i] == 0 {
		i -= 1
	}
	fmt.Println(tmp)
	var res string
	res = ""
	for i >= 0 {
		res += strconv.Itoa(tmp[i])
		i--
	}
	fmt.Println(res)
	return res
}
