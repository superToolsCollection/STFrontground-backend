package util

import "testing"

/**
* @Author: super
* @Date: 2021-01-26 21:54
* @Description:
**/

func TestGenerateVerificationCode(t *testing.T) {
	t.Log(GenerateVerificationCode())
}

func BenchmarkGenerateVerificationCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateVerificationCode()
	}
}
