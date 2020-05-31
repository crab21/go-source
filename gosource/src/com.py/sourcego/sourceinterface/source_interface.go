package sourceinterface

import (
	"fmt"
	"reflect"
)

func InterfaceNil() {
	fmt.Println("isNil test...........")
	var a = returnNil()
	fmt.Println(a == nil)
	of := reflect.ValueOf(a)
	fmt.Println("value: ", of)
	if of.IsNil() {
		fmt.Println("is nil")
	}

	var mmp map[string]string = nil
	mof := reflect.ValueOf(mmp)
	fmt.Println("mof: ", mof)
	if mof.IsNil() {
		fmt.Println("mof is nil")
	}

}

func InterfaceIsValid() {

	fmt.Println("IsValid test............")
	var value_a *interface{} = nil
	of := reflect.ValueOf(value_a)
	if of.IsValid() {
		fmt.Println(of, " is valid") //<nil>  is valid   [如果30L为：var value_a interface{} = nil,则会引发painc]s
	}else{
		fmt.Println(of," is not valid")
	}

	zero := of.IsZero()
	fmt.Println("zero: ",zero) //zero:  true
	fmt.Println("")


	var value_b int = 0
	valueOf := reflect.ValueOf(value_b)
	valid := valueOf.IsValid()

	if valid{
		fmt.Println(value_b," is valid") //0  is valid
	}else{
		fmt.Println(value_b,"is not valid")
	}

	isZero := valueOf.IsZero()
	fmt.Println("iszero: ",isZero) //iszero:  true

	kind := valueOf.Kind()
	fmt.Println("kind: ",kind)
}

func returnNil() *interface{} {
	return nil
}
