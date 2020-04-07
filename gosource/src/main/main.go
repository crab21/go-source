package main

import (
	"fmt"
)

/*func smain() {
	for a:=0;a<1000 ;a++  {
		a := []string{"a", "b", "c"}
		var bs []string
		sourcego.SliceCopyOfApi(bs,a)

	}
}*/
type ByteSize float64
const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)//sourcehttp.WebsocketServerStart()
//sourcechan.MakeChanProcess()
func main() {

	s := ""
	x := []byte(s)              // len(s) == 1
	fmt.Println(cap([]byte(s))) // 32
	fmt.Println(cap(x))         // 8
	//fmt.Println(x)
}
