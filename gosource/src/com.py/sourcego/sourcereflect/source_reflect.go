package sourcereflect

import "reflect"

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
