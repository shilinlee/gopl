// 包word提供文字游戏相关的工具函数
package word


// IsPalindrome 判断一个字符串是否为回文字符串
func IsPalindrome(s string) bool{
	for i:= range s{
		if s[i] != s[len(s)-1-i]{
			return false
		}
	}
	return true
}