package sourceerror

import "fmt"

func ErrorOrPainc() {
	defer func() {
		fmt.Println("ok")
	}()

	defer func() {
		fmt.Println("no....")
	}()

	defer func() {
		fmt.Println("yes")
	}()
	panic("oh no")
}
