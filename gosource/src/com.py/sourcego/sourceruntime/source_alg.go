package sourceruntime

import "fmt"

func c(a, b T) {
	defer func() {
		if sp:=recover();sp != nil {
			fmt.Println(sp)
			print(0)
		}
	}()
	if a != b {
		print(1)
	} else {
		print(2)
	}
}

type T [2]interface{}

/**
the question is : compare A and B by =? ptr or value?
 */
func AlgMemequals() {
	f := func(){
		fmt.Println("ok")
	}
	///usr/local/go/src/runtime/alg.go -->func efaceeq(t *_type, x, y unsafe.Pointer) bool { ...... }
	//runtime.f64equal
	c(T{1.0, f}, T{2.0, f})
	c(T{f, 1}, T{f, 2}) //runtime error: comparing uncomparable type func()

	//runtime.memequal64
	c(T{1, f}, T{2, f})
	c(T{f, 1}, T{f, 2}) //runtime error: comparing uncomparable type func()
}

