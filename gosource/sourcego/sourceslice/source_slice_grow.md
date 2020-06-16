* [分析Slice grow过程](#分析slice-grow过程)
   * [简单过程分析](#简单过程分析)
   * [竟态变量介绍](#竟态变量介绍)
      * [raceenabled](#raceenabled)
   * [内存模型](#内存模型)
      * [内存计算](#内存计算)
      * [内存分配方式](#内存分配方式)
   * [内存溢出](#内存溢出)
      * [判断及处理方式](#判断及处理方式)
   * [grow 总结](#grow-总结)
      * [处理逻辑](#处理逻辑)
      * [错误处理方式](#错误处理方式)


# 分析Slice grow过程

## 简单过程分析

>[见Slice grow函数扩容分析](https://github.com/crab21/go-source/blob/master/gosource/src/com.py/sourcego/sourceslice/source_slice_append.md#slice-grow%E5%87%BD%E6%95%B0%E6%89%A9%E5%AE%B9%E5%88%86%E6%9E%90)

## 竟态变量介绍

### raceenabled
>[来源](https://github.com/golang/go/blob/master/src/runtime/race0.go)
主要是函数使用与否：
![](https://raw.githubusercontent.com/crab21/Images/master//blog/20200514215318.png)

### 
>[来源](https://github.com/golang/go/blob/master/src/runtime/msan0.go)

主要是函数使用与否：
![](https://raw.githubusercontent.com/crab21/Images/master//blog/20200514215133.png)

## 内存模型
### 内存计算

```cgo
func growslice(et *_type, old slice, cap int) slice {
    ......
    //容量判断
    if cap < old.cap {
    		panic(errorString("growslice: cap out of range"))
    	}
                  
     // old slice size is zero
    if et.size == 0 {
    	// append should not create a slice with nil pointer but non-zero len.
    	// We assume that append doesn't need to preserve old.array in this case.
    	return slice{unsafe.Pointer(&zerobase), old.len, cap}
    }
    
    newcap := old.cap
    doublecap := newcap + newcap    //计算两次old cap大小
    if cap > doublecap {            // cap容量足够，则无需处理
    	newcap = cap
    } else {
    	if old.len < 1024 {           // 比1024小，则增加一倍
    		newcap = doublecap
    	} else {
    		// Check 0 < newcap to detect overflow
    		// and prevent an infinite loop.
    		for 0 < newcap && newcap < cap {                //超过1024，增加1/4容量
    			newcap += newcap / 4
    		}
    		// Set newcap to the requested cap when
    		// the newcap calculation overflowed.
    		if newcap <= 0 {                             //大小overflow时
    			newcap = cap
    		}
    	}
	}              
    ......

}

```
### 内存分配方式

```cgo
func growslice(et *_type, old slice, cap int) slice {
    ......
    ......

    var overflow bool
	var lenmem, newlenmem, capmem uintptr
	// Specialize for common values of et.size.
	// For 1 we don't need any division/multiplication.
	// For sys.PtrSize, compiler will optimize division/multiplication into a shift by a constant.
	// For powers of 2, use a variable shift.
	switch {
	case et.size == 1:                                            //直接分配
		lenmem = uintptr(old.len)
		newlenmem = uintptr(cap)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem)
	case et.size == sys.PtrSize:                                  //指针计算分配
		lenmem = uintptr(old.len) * sys.PtrSize
		newlenmem = uintptr(cap) * sys.PtrSize
		capmem = roundupsize(uintptr(newcap) * sys.PtrSize)
		overflow = uintptr(newcap) > maxAlloc/sys.PtrSize
		newcap = int(capmem / sys.PtrSize)
	case isPowerOfTwo(et.size):                                     //2的N次方分配
		var shift uintptr
		if sys.PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
		}
		lenmem = uintptr(old.len) << shift
		newlenmem = uintptr(cap) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
	default:
		lenmem = uintptr(old.len) * et.size
		newlenmem = uintptr(cap) * et.size
		capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.size)
	}
    ......
}
```
## 内存溢出
### 判断及处理方式

```cgo
func growslice(et *_type, old slice, cap int) slice {
    ......
    ......
    //容量超过上限
    if overflow || capmem > maxAlloc {
		panic(errorString("growslice: cap out of range"))
	}

	var p unsafe.Pointer
	if et.ptrdata == 0 {
		p = mallocgc(capmem, nil, false)
		// The append() that calls growslice is going to overwrite from old.len to cap (which will be the new length).
		// Only clear the part that will not be overwritten.
		memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
	} else {
		// Note: can't use rawmem (which avoids zeroing of memory), because then GC can scan uninitialized memory.
		p = mallocgc(capmem, et, true)
		if lenmem > 0 && writeBarrier.enabled {
			// Only shade the pointers in old.array since we know the destination slice p
			// only contains nil pointers because it has been cleared during alloc.
			bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(old.array), lenmem)
		}
	}
    ......
    ......
}
```

## grow 总结

### 处理逻辑

```cgo
*   先预判大小
*   计算大小值
*   计算分配内存大小
*   mallocgc内存
```
### 错误处理方式
*   严重错误直接panic。
*   返回空值。

