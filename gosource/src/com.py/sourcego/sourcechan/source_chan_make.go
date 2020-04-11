package sourcechan

import (
	"fmt"
	"unsafe"
)

func MakeChanProcess() {

	intChan := make(chan string, 1)
	fmt.Println("before:======", uintptr(unsafe.Pointer(&intChan)), "")
	isChanClod(intChan)
	isChanClosed(intChan)
	close(intChan)
	spin()

	var spmap map[string]string
	fmt.Println("before map:======", uintptr(unsafe.Pointer(&spmap)), "")

	isMap(spmap)
}

func isMap(spmap interface{}) {
	fmt.Println("now    map:======", uintptr(unsafe.Pointer(&spmap)), "")
}
func spin() {
	fmt.Println(unsafe.Sizeof(uint64(0)))
}

/**

type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
*/
func isChanClosed(ch chan string) {
	//取到地址
	cptr := *(*uintptr)(unsafe.Pointer(&ch))
	fmt.Println("chanclosed:", cptr)

	//qcount 和dataqsiz 偏移量大小
	cptr += unsafe.Sizeof(uint(0)) * 2

	//buf指针大小
	cptr += unsafe.Sizeof(unsafe.Pointer(uintptr(0)))

	//elemsize偏移大小
	cptr += unsafe.Sizeof(uint16(0))

	//根据偏移量的大小，反推目前closed字段的大小
	p := *(*uint32)(unsafe.Pointer(cptr)) > 0
	fmt.Println(p)
}

func isChanClod(ch interface{}) bool {
	//chp:=ch.(chan int)
	//up := uintptr(unsafe.Pointer(&chp))
	//fmt.Println("turn over:======",up)
	//fmt.Println(&ch,"=====")
	u := uintptr(unsafe.Pointer(&ch))
	fmt.Println("turn now==== ", u)

	cptr := *(*uintptr)(unsafe.Pointer(
		unsafe.Pointer(u + unsafe.Sizeof(uint(0))),
	))

	fmt.Println("turn now=--  ", cptr)
	// this function will return true if chan.closed > 0
	// see hchan on https://github.com/golang/go/blob/master/src/runtime/chan.go
	// type hchan struct {
	// qcount   uint           // total data in the queue
	// dataqsiz uint           // size of the circular queue
	// buf      unsafe.Pointer // points to an array of dataqsiz elements
	// elemsize uint16
	// closed   uint32
	// **

	cptr += unsafe.Sizeof(uint(0)) * 2
	cptr += unsafe.Sizeof(unsafe.Pointer(uintptr(0)))
	cptr += unsafe.Sizeof(uint16(0))
	p := *(*uint32)(unsafe.Pointer(cptr)) > 0
	fmt.Println(p)

	return p
}
