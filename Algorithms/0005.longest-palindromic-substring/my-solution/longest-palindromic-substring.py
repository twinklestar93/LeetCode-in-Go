class Solution:
    def longestPalindrome(self, s: str) -> str:
        begin, maxLen = 0, 1
        for i in range(len(s)):
            if len(s) - i < maxLen/2:
                break
            #循环结束后，s[b:e+1]为一相同的字符串
            b, e = i, i
            while e < len(s)-1 and s[e+1] == s[e]:
                e+=1

            #下一个回文的正中间段的首字符只会是s[e+1]
            i = e+1
            while e < len(s)-1 and b > 0 and s[e+1]==s[b-1]:
                    e+=1
                    b-=1

            newLen = e + 1 - b
            if newLen > maxLen:
                begin = b
                maxLen = newLen

        return s[begin: begin++maxLen]
                


    def isPalindrome(self, s: str) -> bool:
        if len(s) == 1 or len(s) == 0:
            return True
        else:
            if s[0] == s[-1]:
                return self.isPalindrome(s[1:-1])
            else:
                return False

if __name__ == '__main__':
    sol = Solution()
    res = []
    res.append(sol.isPalindrome("bab"))
    res.append(sol.isPalindrome("gghh"))
    res.append(sol.isPalindrome("a"))
    res.append(sol.isPalindrome("cbbc"))
    print(res)
    print(sol.longestPalindrome("abbcccddcccbba")=="abbcccddcccbba")
    print(sol.longestPalindrome("a"))

