package sourceerror

import "testing"

func TestErrorOrPainc(t *testing.T) {
	//yes
	//no....
	//ok
	//error:..... oh no
	ErrorOrPainc()
}

func BenchmarkErrorOrPainc(b *testing.B) {
	ErrorOrPainc()
}