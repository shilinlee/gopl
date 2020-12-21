// 包word提供文字游戏相关的工具函数
package word

import "unicode"

// IsPalindrome 判断一个字符串是否为回文字符串
func IsPalindrome(s string) bool{
	var letters []rune
	for _, r := range s{
		if unicode.IsLetter(r){
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i:= range letters{
		if letters[i] != letters[len(letters)-1-i]{
			return false
		}
	}
	return true
}

