package sourcego

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

func BenchmarkSliceAppendMake(b *testing.B) {
	b.Run("make", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			des := []string{"1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4",}
			for a := 0; a < 1000; a++ {
				SliceAppendMake(des)
			}
		}
	})
	b.Run("normal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			des := []string{"1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4", "1", "2", "3", "4",}
			for a := 0; a < 1000; a++ {

				SliceAppendNormal(des)
			}
		}
	})

	b.Run("for-range", func(b *testing.B) {
		a := []int{1, 2, 3, 4, 5, 6,}
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

func Benchmark_MakeAndCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMakeCopy = make([]T, N)
		copy(yForMakeCopy, xForMakeCopy)
	}
}


//more efficiency
func Benchmark_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForAppend = append(xForAppend[:0:0], xForAppend...)
	}
}