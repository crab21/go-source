package sourcestruct

import "fmt"

type People interface {
	Speak(string) string
}

type Stduent struct{
	Age int
}

func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	stu.Age=2
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
