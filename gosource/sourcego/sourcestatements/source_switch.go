package sourcestatements

import "fmt"

//case表达式布尔值可以重复，非布尔值不可重复，编译器报错，但不影响运行
func SwitchDeafultBool() {
	switch false {
	case false:
		fmt.Println("111")
	case false:
		fmt.Println("222")
	}
}

//非布尔类型不可以重复
func SwitchDeafultNowBool() {
	a := 2
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	//case 3:   //duplicate case 3 in switch
	//	fmt.Println("3-1")
	default:
		fmt.Println("4")

	}
}

func SwitchDeafultTrue() {
	switch { // ==> switch true{}
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("aa")

	}
}

func SwitchDefaultTrue_1() {
	switch False()
	{
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}
func False() bool{
	return true
}