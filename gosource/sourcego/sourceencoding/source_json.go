package sourceencoding

import (
	"encoding/json"
	"fmt"
)

type UsrName struct {
	Age  int  `json:"age"`
	Number int
	User *Usr `json:"sp,omitempty"`
}
type Usr struct {
	Names string `json:",omitempty"`
}


func JsonMarshalUnmarshal(){
	var username UsrName
	var userNamePointer *UsrName

	usernameType := fmt.Sprintf("%T",username)
	userNamePointerType := fmt.Sprintf("%T",userNamePointer)

	//init value:  {0 0 <nil>}    sourceencoding.UsrName    &{0 0 <nil>}
	fmt.Println("init value: ",username,"  ",usernameType, "  ",&username)
	//init value:  <nil>    *sourceencoding.UsrName     <nil>
	fmt.Println("init value: ",userNamePointer,"  ",userNamePointerType,"   ",userNamePointer)
}

type UsrNameOm struct {
	Age  int  `json:"age"`
	Number int
	UserAges []int      `json:",omitempty"`
	User *Usr           `json:"sp,omitempty"`
	Users []Usr         `json:"usrs,omitempty"`
	UsersPointer []*Usr `json:"usrsPointer,omitempty"`
	// Usera Usr `json:"spa,omitempty"`  //非pointer类型  omitempty不会生效
}

func Omitempty() {
	op := UsrNameOm{Age: 10}
	result,error := json.Marshal(op)
	//{"age":10,"Number":0}
	fmt.Println(string(result))
	fmt.Println(error)
}