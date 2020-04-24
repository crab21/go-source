package sourcego

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"runtime"
	"strconv"
	"unsafe"
)

func StringToByte(source string) []byte {
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
	s := fnv.New64a()
	for i := 0; i < 100000; i++ {
		s.Reset()
		s.Write([]byte(strconv.Itoa(i)))
		a := hex.EncodeToString(s.Sum(nil))

		as, error := strconv.ParseUint(a, 16, 64)
		if error != nil {
			return
		}
		sp[as] = as
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
	fmt.Printf("HeapAlloc = %v HeapIdel= %v HeapSys = %v  HeapReleased = %v\n", m.HeapAlloc/1024, m.HeapIdle/1024, m.HeapSys/1024, m.HeapReleased/1024)
}

func StringToBytes() {
	spp := ""
	s := []byte(spp)  // 0x0042 00066 (source_string.go:64)      CALL    runtime.stringtoslicebyte(SB)
	sp := []byte(spp) // CALL    runtime.stringtoslicebyte(SB)
	fmt.Println(cap(sp))
	s1 := append(s, 'a')
	s2 := append(s, 'b')
	fmt.Println(s1, "==========", s2)
	fmt.Println(cap(s), "==========", cap(s2), "==", cap(s1))
}

/**
type string struct{
	data
	len
}

type slice struct{
	data
	len
	cap
}
*T1--->uintptr---->*T2
T1的尺寸不能小于T2
当由小尺寸往大尺寸转时，不能直接转换，需要借助中间辅助量来完成
 */
func StringToByteSlice() {
	sa:= "abc"

	//error done:
	sa_slice:=*(*[]byte)(unsafe.Pointer(&sa))
	//[97 98 99]
	fmt.Println(sa_slice)
	//len:  3  cap:  4294967297
	fmt.Println("len: ",len(sa_slice), " cap: ",cap(sa_slice))

	//right done:
	strEx :=StringEx{sa,len(sa)}
	strEx_slice :=*(*[]byte)(unsafe.Pointer(&strEx))
	//[97 98 99]
	fmt.Println(strEx_slice)
	//len:  3  cap:  3
	fmt.Println("len: ",len(strEx_slice), " cap: ",cap(strEx_slice))

}
type StringEx struct {
	string
	cap int
}