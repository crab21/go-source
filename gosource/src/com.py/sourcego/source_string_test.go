package sourcego

import (
	"testing"
	"unsafe"
)

var count =1000
func BenchmarkStringToByte(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		ss := "sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd"

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for s := 0; s < count; s++ {
				_ = []byte(ss)
			}
		}
	})

	b.Run("speed", func(b *testing.B) {
		ss := "sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd"

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for s := 0; s < count; s++ {
				_ = *(*[]byte)(unsafe.Pointer(&ss))
			}
		}
	})

	b.Run("normal", func(b *testing.B) {
		ss := []byte("sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd")

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for s := 0; s < count; s++ {
				_ = string(ss)
			}
		}
	})

	b.Run("speed", func(b *testing.B) {
		ss := []byte("sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd")

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for s := 0; s < count; s++ {
				_ = *(*string)(unsafe.Pointer(&ss))
			}
		}
	})

	b.Run("speedNormal", func(b *testing.B) {
		ss := []byte("sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd")
		h := (*[2]uintptr)(unsafe.Pointer(&ss))
		x := [3]uintptr{h[0], h[1], h[1]}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for s := 0; s < count; s++ {
				_ = *(*string)(unsafe.Pointer(&x))
				//fmt.Println(spp)
			}
		}

	})
}
