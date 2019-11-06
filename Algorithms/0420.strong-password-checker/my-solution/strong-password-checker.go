package problem0420

import  (
	"unicode"
)

func max(x, y int) int{
	if x < y {
		return y
	} 
	return x
}

func min(x, y int) int{
	if x < y {
		return  x
	}
	return y
}

func strongPasswordChecker(s string) int {
	missing_type_cnt := 3
	for _, c := range(s) {
		if (c >= 'a' && c <= 'z') {
			missing_type_cnt  -= 1
		}
		if (c >= 'A' && c <= 'Z') {
			missing_type_cnt  -= 1
		}
		if (!unicode.IsDigit(c)) {
			missing_type_cnt  -= 1
		}
	}
	
	total_change_cnt := 0
	one_change_cnt, two_change_cnt, three_change_cnt := 0, 0, 0
	i := 2
	for i < len(s) {
		if s[i] == s[i-1] && s[i-1] == s[i-2] {
			length := 2
			for  i < len(s) && s[i] == s[i-1] {
				length+=1
				i+=1
			}
			total_change_cnt += length / 3
			if length % 3 == 0 {
				one_change_cnt += 1
			} else if  length % 3 == 1 {
				two_change_cnt += 1
			} else {
				three_change_cnt += 1
			}
		} else {
			i+=1
		}
	}
	if len(s) < 6 {
		return max(missing_type_cnt,  6-len(s))
	} else if len(s) <= 20 {
		return max(missing_type_cnt, total_change_cnt)
	}
	delete_cnt := len(s) - 20

	total_change_cnt -= min(delete_cnt, one_change_cnt) / 1
	total_change_cnt -= min(max(delete_cnt  - one_change_cnt, 0), two_change_cnt * 2) / 2
	total_change_cnt -= min(max(delete_cnt - one_change_cnt - 2 * two_change_cnt, 0), three_change_cnt * 3) / 3

	return delete_cnt + max(missing_type_cnt, total_change_cnt)
}
