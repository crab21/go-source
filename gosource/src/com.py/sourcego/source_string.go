package sourcego

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"runtime"
	"strconv"
	"unsafe"
)

func StringToByte(source string)[]byte{
	return *(*[]byte)(unsafe.Pointer(&source))
}

/**
HeapAlloc = 148 HeapIdel= 64784 HeapSys = 65312  HeapReleased = 64752
HeapAlloc = 70069 HeapIdel= 60056 HeapSys = 130656  HeapReleased = 54400
8
1000000

 */
func IntHashInfo() {
	var sp map[uint64]uint64 = make(map[uint64]uint64)
	printMemStats()
	s:=fnv.New64a()
	for i:=0;i< 100000 ;i++  {
		s.Reset()
		s.Write([]byte(strconv.Itoa(i)))
		a :=hex.EncodeToString(s.Sum(nil))


		as,error :=strconv.ParseUint(a,16,64)
		if error!= nil {
			return
		}
		sp[as]=as
		//fmt.Println(as)

	}
	printMemStats()
	fmt.Println(unsafe.Sizeof(sp))
	fmt.Println(len(sp))
}

/**
HeapSys：程序向应用程序申请的内存

HeapAlloc：堆上目前分配的内存

HeapIdle：堆上目前没有使用的内存

HeapReleased：回收到操作系统的内存
 */
func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %v HeapIdel= %v HeapSys = %v  HeapReleased = %v\n", m.HeapAlloc/1024, m.HeapIdle/1024, m.HeapSys/1024,  m.HeapReleased/1024)
}