package sourceslice

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSliceCopyOfApi(b *testing.B) {
	b.Run("api", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for a := 0; a < 1000; a++ {
				a := []string{"a", "b", "c"}
				var bs []string = make([]string, len(a), len(a))
				SliceCopyOfApi(bs, a)
			}
		}
	})
}

func BenchmarkSliceCopyOfNormal(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for a := 0; a < 1000; a++ {
				a := []string{"a", "b", "c"}
				var bs []string
				SliceCopyOfNormal(bs, a)
			}
		}
	})
}

//go test -bench BenchmarkSliceAppendMake -run=^$ -benchtime 1s -count 4 -benchmem -cover
/**
BenchmarkSliceAppendMake/make-8                     3606            333504 ns/op         1152004 B/op       2000 allocs/op
BenchmarkSliceAppendMake/make-8                     3586            334474 ns/op         1152004 B/op       2000 allocs/op
BenchmarkSliceAppendMake/make-8                     3590            334853 ns/op         1152004 B/op       2000 allocs/op
BenchmarkSliceAppendMake/make-8                     3598            331386 ns/op         1152004 B/op       2000 allocs/op
BenchmarkSliceAppendMake/normal-8                   8679            136884 ns/op          384001 B/op       1000 allocs/op
BenchmarkSliceAppendMake/normal-8                   8703            137581 ns/op          384001 B/op       1000 allocs/op
BenchmarkSliceAppendMake/normal-8                   8727            137954 ns/op          384001 B/op       1000 allocs/op
BenchmarkSliceAppendMake/normal-8                   8728            137728 ns/op          384001 B/op       1000 allocs/op
*/
func BenchmarkSliceAppendMake(b *testing.B) {
	b.Run("make", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			des := []string{"1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4"}
			for a := 0; a < 1000; a++ {
				SliceAppendMake(des)
			}
		}
	})
	
	b.Run("normal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			des := []string{"1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4"}
			for a := 0; a < 1000; a++ {

				SliceAppendNormal(des)
			}
		}
	})

	b.Run("for-range", func(b *testing.B) {
		a := []int{1, 2, 3, 4, 5, 6}
		myMap := make(map[int]*int)
		for k, i := range a {
			fmt.Println(k)
			myMap[k] = &i

		}
		for k, v := range myMap {

			fmt.Println(k, "    ", *v)
		}
	})

}

func TestMapInit(t *testing.T) {
	MapInit()
}

var x = strings.Repeat("hello", 100) + " world!"
var y = strings.Repeat("hello", 99) + " world!"

func BenchmarkNumSameBytes_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumSameBytes_1(x, y)
	}
}

func BenchmarkNumSameBytes_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumSameBytes_2(x, y)
	}
}

func TestSliceCopyOfSelf(t *testing.T) {
	SliceCopyOfSelf()
}

type Name struct {
	Age int
}

const N = 1024 * 1024

type T = float64

var xForMakeCopy = make([]T, N)
var xForAppend = make([]T, N)
var yForMakeCopy []T
var yForAppend []T

//go test -bench Benchmark_MakeAndCopy -run=^$ -benchtime 1s -count 4 -benchmem -cover
/**
Benchmark_MakeAndCopy/copy-8                1071           1019899 ns/op         8388612 B/op          1 allocs/op
Benchmark_MakeAndCopy/copy-8                1069           1012537 ns/op         8388611 B/op          1 allocs/op
Benchmark_MakeAndCopy/copy-8                1089           1012071 ns/op         8388610 B/op          1 allocs/op
Benchmark_MakeAndCopy/copy-8                1093           1014036 ns/op         8388611 B/op          1 allocs/op
Benchmark_MakeAndCopy/append-8              2178            573667 ns/op         8388610 B/op          1 allocs/op
Benchmark_MakeAndCopy/append-8              1891            555216 ns/op         8388611 B/op          1 allocs/op
Benchmark_MakeAndCopy/append-8              1974            555206 ns/op         8388610 B/op          1 allocs/op
Benchmark_MakeAndCopy/append-8              1910            553872 ns/op         8388610 B/op          1 allocs/op

*/
func Benchmark_MakeAndCopy(b *testing.B) {
	b.Run("copy", func(bc *testing.B) {
		for i := 0; i < bc.N; i++ {
			for i := 0; i < b.N; i++ {
				yForMakeCopy = make([]T, N)
				copy(yForMakeCopy, xForMakeCopy)
			}
		}
	})

	b.Run("append", func(bc *testing.B) {
		for i := 0; i < bc.N; i++ {
			for i := 0; i < b.N; i++ {
				yForAppend = append(xForAppend[:0:0], xForAppend...)
			}
		}
	})
}

//more efficiency
func Benchmark_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForAppend = append(xForAppend[:0:0], xForAppend...)
	}
}

func TestEscapeSliceSize(t *testing.T) {
	EscapeSliceSize()
}

func TestAppendSliceError(t *testing.T) {
	AppendSliceError()
}


func TestSliceAppendSlice(t *testing.T) {
	SliceAppendSlice()
}