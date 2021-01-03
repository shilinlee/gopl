// go test -v  -bench=.
package word

import (
	"testing"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: POanama")
	}
}
