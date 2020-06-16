package sourceerror

import "testing"

func TestErrorOrPainc(t *testing.T) {
	ErrorOrPainc()
}

func BenchmarkErrorOrPainc(b *testing.B) {
	ErrorOrPainc()
}