# string

### 结构
#### string结构
```cgo
 type string struct{
    len
    data
 }
```
#### slice结构
```cgo
type slice struct{
	data
	len
	cap
}
```

### 用法
#### string2byte  ⚔
* 不可以直接转换，错误示例：
```cgo
_ = *(*[]byte)(unsafe.Pointer(&ss)) 
```

* 推荐用法：✔️

```cgo
type StringEx struct {
	string
	cap int
}
strEx := StringEx{ss,len(ss)} 	
_ =*(*[]byte)(unsafe.Pointer(&strEx))
```
#### byte2string
* 可以直接转换
![](https://raw.githubusercontent.com/crab21/Images/master//blog/20200523150035.png)