package sourcehttp

import (
	"fmt"
	"strconv"
)

type sp func(sp string) int
type spInterface interface {
	sp(sp int)int
}

type Name struct {
	Age int
}
func (s Name)sp(sp int)int {
	s.Age+=sp
	return s.Age
}
func spInit() {
	spa := func(sp string) int {
		atoi, err := strconv.Atoi(sp)
		if err != nil {
			return 0
		}
		return atoi
	}
	spKill(spa,"1")
}


func spKill(g sp,sps string) {
	fmt.Println(g(sps))
	calc(1,2,2)
}

func calc(a,_ int,b int) bool {
	return false
}
