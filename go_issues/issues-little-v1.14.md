## v1.14 isseus整理
### [39607](https://github.com/golang/go/issues/39607) 
#### reflect.DeepEqual()问题：
> go version: v1.14

[reflect.DeepEqual介绍](https://golang.google.cn/pkg/reflect/#DeepEqual)
```cgo

Array  ---> elements.
Struct  ---> fields.
Interface --->concrete values.
Func ---> values are deeply equal if both are nil,otherwise they are not deeply equal.
Map ---> same len;same map object;both nil .
Pointer ---> equal using Go’s == operator.
```


### [39806](https://github.com/golang/go/issues/39806)
#### strings.split问题：
>go version: v1.14

[split Api](https://godoc.org/strings#Split)
```cgo
摘选：
If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
```


### [38173](https://github.com/golang/go/issues/38173)
#### encoding/json: marshal result of string type struct field with ",string" option change 
> go version: v1.14
```cgo
当处理string json.Marshal()时，编码格式问题：
如：”aaa\tbbb” 用Marshal后：
期待：		“\“aaa\\tbbb\””
实际输出： 	“\“aaa\tbbb\””
```


