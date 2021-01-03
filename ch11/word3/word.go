// 包word提供文字游戏相关的工具函数
package word

import "unicode"

// IsPalindrome 判断一个字符串是否为回文字符串
func IsPalindrome(s string) bool {
	var letters = make([]rune, 0, len(s)) // 提前分配空间
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
