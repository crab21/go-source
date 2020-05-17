

## json使用及转换问题

### marshal 

#### TOP_1: [omitempty关键字](https://golang.org/pkg/encoding/json/#Marshal)
原文参考：
![](https://raw.githubusercontent.com/crab21/Images/master//blog/20200517165131.png)
```cgo
omitempty作用域:
* an empty value
* defined as false
* 0
* a nil pointer
* a nil interface value
* and any empty array, slice, map, or string
```

### unmarshal

```cgo
func Unmarshal(data []byte, v interface{}) error
```

#### TOP_1: [抛出异常InvalidUnmarshalError](https://golang.org/pkg/encoding/json/#Unmarshal)
原文参考：
![](https://raw.githubusercontent.com/crab21/Images/master//blog/20200517164731.png)
```cgo
v 为空        或          v不是pointer类型
```

## [测试用例](https://github.com/crab21/go-source/blob/master/gosource/src/com.py/sourcego/sourceencoding/source_json.go)