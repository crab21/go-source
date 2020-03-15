package sourcego

import "unsafe"

func StringToByte(source string)[]byte{
	return *(*[]byte)(unsafe.Pointer(&source))
}
