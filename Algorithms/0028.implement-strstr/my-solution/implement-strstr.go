package problem0028



func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
	    return 0
	}
	return KMP(haystack, needle)
}

func KMP(text string, pattern string) int {
	prefix := getPrefix(pattern)
	textLen := len(text)
	j := -1
	for i:=0; i<textLen;i++ {
	    for j > -1 && pattern[j+1] != text[i] {
		j = prefix[i]
	    }
	    if pattern[j+1] == text[i] {
		j++;
	    }
	    if j == (len(pattern) - 1) {
		    return i - j
	    }
	}
	return -1
}

func getPrefix(pattern string) []int {
	prefix := make([]int, len(pattern))
	j := -1
	for i:=1; i < len(pattern); i++ {
	    for j > -1 && pattern[j+1] != pattern[j] {
	    	j = prefix[j]
	    }
	    if pattern[j+1] == pattern[i] {
		j++;
	    }
	    prefix[i] = j
	}
	return prefix
}
