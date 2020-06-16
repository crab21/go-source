package sourcepool

import (
	"fmt"
	"sync"
)

func PoolBuild() {
	sp:=&sync.Pool{
		//若没有New函数，则Get返回值为nil。
		New: func() interface{} {
			return 0
		},
	}

	a:=sp.Get().(int)
	a1:=sp.Get().(int)
	sp.Put(1)
	b :=sp.Get().(int)
	fmt.Println(a," ",a1,"  ",b)
}
