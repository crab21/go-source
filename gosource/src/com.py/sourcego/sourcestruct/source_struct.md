#### method values
    参考----[go官方文档](https://golang.org/ref/spec#Method_values)

##### go方法集规约
###### 从值得角度来看
```cgo
Values	Methods Receivers
T	    (t T)
*T	    (t T) and (t *T)
```

###### 从接收者角度来看
```cgo
Values	Methods Receivers
(t T)	T and *T
(t *T)	*T
```