package sourcereflect

import (
	"fmt"
	"reflect"
)

type S struct {x int}
func (s S) M() int {return s.x}
type ST struct {*S}

func ReflectStruct() {
	t := ST{&S{1}}
	f := t.M
	g := reflect.ValueOf(&t).Elem().MethodByName("M").Interface().(func() int)
	t.S = &S{2}
	println(f(), g())
}
func ReflectInfo() {

	var s *interface{}  = new(interface{})
	*s = "wanga"
	fmt.Println(s)

	//反射取值
	value := reflect.ValueOf(s)
	//判空
	result := value.IsNil()
	fmt.Println(result) //false

	//取元素
	ele :=value.Elem()
	fmt.Println(ele.Interface()) //wanga
	fmt.Println(ele.Kind()) 	//interface


	var p *interface{} = nil
	spinfo :=reflect.ValueOf(p)
	isnil :=spinfo.IsNil()
	fmt.Println(isnil) //true

	ele_value :=spinfo.Elem()
	fmt.Println(ele_value) //<invalid reflect.Value>
	fmt.Println("================================")
	fmt.Println()

	var np interface{}   // 若为:var np *interface{}  或 var np *string    指针类型  则下面的valueOf不会为 <invalid reflect.Value>
	fmt.Println(np)
	nospinfo :=reflect.ValueOf(np)
	fmt.Println("interface .... nil ",nospinfo) //interface .... nil  <invalid reflect.Value>
	kn :=nospinfo.Kind()
	fmt.Println("kind=====>",kn)  //kind=====> invalid

	nospinfoKind := nospinfo.Kind()
	fmt.Println("noinfokind: ",nospinfoKind) //noinfokind:  invalid

	nele_value :=nospinfo.Elem() //painc ×
	fmt.Println(nele_value)

	nnil :=nospinfo.IsNil()
	fmt.Println(nnil)

	fmt.Println("================================")
	fmt.Println()

	var str_p *string = new(string)   //如果不new，下一行赋值会painc
	*str_p = "ssss"
	fmt.Println("value: ",str_p)  //value:  <nil>
	strInfo := reflect.ValueOf(str_p)
	strNil := strInfo.IsNil()
	fmt.Println(strNil)	 //true

	strEle := strInfo.Elem()
	fmt.Println(strEle)	 //ssss


}

func PointerInterface(sp interface{}) {
	fmt.Println("spinfo")

	sp_result := reflect.ValueOf(sp)
	fmt.Println(sp_result)

	info := sp_result.IsNil()
	fmt.Println(info)
}
