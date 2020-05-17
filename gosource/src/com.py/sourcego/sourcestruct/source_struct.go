package sourcestruct

import (
	"encoding/json"
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct {
	Age int
}

func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	stu.Age = 2
	return
}

/**go方法集规约
Values	Methods Receivers
(t T)	T and *T
(t *T)	*T
*/
func SSMain() {
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
	fmt.Println(peo)
}

type Usr struct {
	Names string `json:",omitempty"`
}
type UsrName struct {
	Age  int  `json:"age"`
	Number int
	User *Usr `json:"sp,omitempty"`
}

func Omitempty() {
	op := UsrName{Age:10}
	result,error := json.Marshal(op)
	fmt.Println(string(result))
	fmt.Println(error)
}

func JsonMarshalUnmarshal(){
	var username UsrName
	var userNamePointer *UsrName
	usernameType := fmt.Sprintf("%T",username)

	userNamePointerType := fmt.Sprintf("%T",userNamePointer)

	fmt.Println("init value: ",username,"  ",usernameType, "  ",&username)
	fmt.Println("init value: ",userNamePointer,"  ",userNamePointerType,"   ",userNamePointer)
}