package tools

import (
	"testing"
)

func TestLeftRuneMax(t *testing.T) {

	testCase := []struct {
		test   string
		len    int
		result string
	}{
		{"1234567890Ы", 10, "1234567890"},
		{"1234567890Ы", 10, "1234567890"},
		{"1234567", 10, "1234567"},
	}

	for _, v := range testCase {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			res := LeftRuneMax(v.test, v.len)
			if res != v.result {
				t.Error("[E] !=", v.test)
				t.Fail()
			}
		})
	}
	t.Parallel()
}

func BenchmarkLeftRuneMaxB(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = LeftRuneMax("1234567890Ы", 10)
	}
}
