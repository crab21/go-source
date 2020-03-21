package sourcego

import (
	"fmt"
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
