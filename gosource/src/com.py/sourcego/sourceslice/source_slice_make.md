
- [slice make过程](#slice-make--)
  * [make大致流程分析](#----)
  * [len、size、obj.size判断](#len-size-objsize--)
  * [mallocgc过程分析](#mallocgc----)
  * [slice 创建方式的区别](#slice--------)
  * [slice make注意事项](#slice-make----)

## slice make过程

### make大致流程分析
```cgo
1、计算预分配内存大小，mem和overflow（溢出标志）。
2、判断len、cap、size关系。
3、mallocgc分配内存。
```

### len、size、obj.size判断
 
>以下任何一个条件成立，则会抛出异常信息
>
>panicmakeslicelen()   or  panicmakeslicecap()
 ```cgo
overflow || mem > maxAlloc || len < 0 || len > cap
溢出位 || 内存大于最大分配内存 || 长度小于0 || 长度比容量大
```
 
 [Slice创建之make源码](https://github.com/golang/go/blob/master/src/runtime/slice.go#L83)：
 
 ```cgo
func makeslice(et *_type, len, cap int) unsafe.Pointer {
    //计算内存和是否溢出位  按容量
	mem, overflow := math.MulUintptr(et.size, uintptr(cap))

    //溢出  || 内存超出最大 || 长度不足 || 长度比容量大
	if overflow || mem > maxAlloc || len < 0 || len > cap {
		// NOTE: Produce a 'len out of range' error instead of a
		// 'cap out of range' error when someone does make([]T, bignumber).
		// 'cap out of range' is true too, but since the cap is only being
		// supplied implicitly, saying len is clearer.
		// See golang.org/issue/4085.
    
        //再次计算  按长度
		mem, overflow := math.MulUintptr(et.size, uintptr(len))
        //长度异常....
		if overflow || mem > maxAlloc || len < 0 {
			panicmakeslicelen()
		}
        //容量异常....
		panicmakeslicecap()
	}
    
    //分配内存
	return mallocgc(mem, et, true)
}
```
### mallocgc过程分析

### slice 创建方式的区别

### slice make注意事项
