package sourceslice


### 扩容分析
```
func AppendLearn_1() {
	AppendError()

}

```

>result:

#### 当前容量够用时
```
        0x0032 00050 (source_slice.go:130)      PCDATA  $0, $1
        0x0032 00050 (source_slice.go:130)      PCDATA  $1, $0
        0x0032 00050 (source_slice.go:130)      LEAQ    type.[5]int(SB), AX
        0x0039 00057 (source_slice.go:130)      PCDATA  $0, $0
        0x0039 00057 (source_slice.go:130)      MOVQ    AX, (SP)
        0x003d 00061 (source_slice.go:130)      CALL    runtime.newobject(SB)
        0x0042 00066 (source_slice.go:130)      PCDATA  $0, $1
        0x0042 00066 (source_slice.go:130)      MOVQ    8(SP), AX
        0x0047 00071 (source_slice.go:130)      PCDATA  $1, $1
        0x0047 00071 (source_slice.go:130)      MOVQ    AX, "".&arr1+232(SP)
        0x004f 00079 (source_slice.go:130)      MOVQ    ""..stmp_5(SB), CX
        0x0056 00086 (source_slice.go:130)      MOVQ    CX, (AX)
        0~~~~x0059 00089 (source_slice.go:130)      MOVUPS  ""..stmp_5+8(SB), X0
        0x0060 00096 (source_slice.go:130)      MOVUPS  X0, 8(AX)
        0x0064 00100 (source_slice.go:130)      MOVUPS  ""..stmp_5+24(SB), X0    //movups将4个不精准的单精度值传送到内存
        0x006b 00107 (source_slice.go:130)      PCDATA  $0, $0
        0x006b 00107 (source_slice.go:130)      MOVUPS  X0, 24(AX)
        0x006f 00111 (source_slice.go:131)      PCDATA  $0, $1
        0x006f 00111 (source_slice.go:131)      MOVQ    "".&arr1+232(SP), AX
        0x0077 00119 (source_slice.go:131)      TESTB   AL, (AX)
        0x0079 00121 (source_slice.go:131)      JMP     123
        0x007b 00123 (source_slice.go:131)      PCDATA  $0, $-1
        0x007b 00123 (source_slice.go:131)      PCDATA  $1, $-1
        0x007b 00123 (source_slice.go:131)      JMP     125
        0x007d 00125 (source_slice.go:131)      PCDATA  $0, $2
        0x007d 00125 (source_slice.go:131)      PCDATA  $1, $1
        0x007d 00125 (source_slice.go:131)      LEAQ    8(AX), CX
        //MOVQ搬运8字节  664+8=672
        0x0081 00129 (source_slice.go:131)      MOVQ    CX, "".slice1+664(SP)
        0x0089 00137 (source_slice.go:131)      MOVQ    $1, "".slice1+672(SP)
        0x0095 00149 (source_slice.go:131)      MOVQ    $4, "".slice1+680(SP)
  		0x00a1 00161 (source_slice.go:132)      JMP     163
        0x00a3 00163 (source_slice.go:132)      MOVQ    $6, 16(AX)
        0x00ab 00171 (source_slice.go:132)      MOVQ    $7, 24(AX)
        0x00b3 00179 (source_slice.go:132)      PCDATA  $0, $3
        0x00b3 00179 (source_slice.go:132)      MOVQ    $8, 32(AX)
        0x00bb 00187 (source_slice.go:132)      PCDATA  $0, $0
        0x00bb 00187 (source_slice.go:132)      PCDATA  $1, $2
        0x00bb 00187 (source_slice.go:132)      MOVQ    CX, "".slice1+664(SP)
        0x00c3 00195 (source_slice.go:132)      MOVQ    $4, "".slice1+672(SP)
        0x00cf 00207 (source_slice.go:132)      MOVQ    $4, "".slice1+680(SP)
```
#### 当前容量不够用时
```
    0x0591 01425 (source_slice.go:136)      PCDATA  $0, $1
        0x0591 01425 (source_slice.go:136)      MOVQ    "".&arr2+224(SP), AX
        0x0599 01433 (source_slice.go:136)      TESTB   AL, (AX)
        0x059b 01435 (source_slice.go:136)      JMP     1437
        0x059d 01437 (source_slice.go:136)      PCDATA  $0, $-1
        0x059d 01437 (source_slice.go:136)      PCDATA  $1, $-1
        0x059d 01437 (source_slice.go:136)      JMP     1439
        0x059f 01439 (source_slice.go:136)      PCDATA  $0, $1
        0x059f 01439 (source_slice.go:136)      PCDATA  $1, $30
        0x059f 01439 (source_slice.go:136)      ADDQ    $8, AX
        0x05a3 01443 (source_slice.go:136)      MOVQ    AX, "".slice2+640(SP)
        0x05ab 01451 (source_slice.go:136)      MOVQ    $2, "".slice2+648(SP)
        0x05b7 01463 (source_slice.go:136)      MOVQ    $4, "".slice2+656(SP)
  		0x05c3 01475 (source_slice.go:137)      JMP     1477
      0x05c5 01477 (source_slice.go:137)      PCDATA  $0, $2
      0x05c5 01477 (source_slice.go:137)      LEAQ    type.int(SB), CX
      0x05cc 01484 (source_slice.go:137)      PCDATA  $0, $1
      0x05cc 01484 (source_slice.go:137)      MOVQ    CX, (SP)
      0x05d0 01488 (source_slice.go:137)      PCDATA  $0, $0
      0x05d0 01488 (source_slice.go:137)      MOVQ    AX, 8(SP)
      0x05d5 01493 (source_slice.go:137)      MOVQ    $2, 16(SP) //起始添加的指针位置
      0x05de 01502 (source_slice.go:137)      MOVQ    $4, 24(SP) //现有的最大容量指针位置
      0x05e7 01511 (source_slice.go:137)      MOVQ    $5, 32(SP) //添加完成后的指针位置位置
// 调用切片扩容？？依据呢？？
//todo 分析扩容原因和关键点
//==> MOVQ将参数$5，移动到32（SP）寄存器上(理解有偏差，后期查阅后继续修正
      0x05f0 01520 (source_slice.go:137)      CALL    runtime.growslice(SB)              =============*********=============
      0x05f5 01525 (source_slice.go:137)      PCDATA  $0, $1
      0x05f5 01525 (source_slice.go:137)      MOVQ    40(SP), AX
      0x05fa 01530 (source_slice.go:137)      MOVQ    48(SP), CX
      0x05ff 01535 (source_slice.go:137)      MOVQ    56(SP), DX
      0x0604 01540 (source_slice.go:137)      ADDQ    $3, CX
      0x0608 01544 (source_slice.go:137)      JMP     1546
      0x060a 01546 (source_slice.go:137)      MOVQ    $6, 16(AX)
      0x0612 01554 (source_slice.go:137)      MOVQ    $7, 24(AX)
      0x061a 01562 (source_slice.go:137)      MOVQ    $8, 32(AX)
      0x0622 01570 (source_slice.go:137)      PCDATA  $0, $0
      0x0622 01570 (source_slice.go:137)      PCDATA  $1, $31
      0x0622 01570 (source_slice.go:137)      MOVQ    AX, "".slice2+640(SP)
      0x062a 01578 (source_slice.go:137)      MOVQ    CX, "".slice2+648(SP)
      0x0632 01586 (source_slice.go:137)      MOVQ    DX, "".slice2+656(SP)
 ```
#### slice append扩容说明
>修正：根据builtin.go源码说明，append根据大小，自动选择是否扩容，再根据上面汇编结果，runtime.growslice，判断当容量不足时，会选择自动扩容。

### plan9源码分析扩容情况

[go-plan9-slice源码分析](https://plan9.io/sources/contrib/ericvh/go-plan9/src/pkg/runtime/slice.c)
#### 分析其中一种截取情况：slice[a:b:c]  条件： a<b b<=c  len:=b-a cap:=c-a
```cgo
// sliceslice(old []any, lb int, hb int, width int) (ary []any);
void
runtime·sliceslice(Slice old, uint32 lb, uint32 hb, uint32 width, Slice ret)
{
	if(hb > old.cap || lb > hb) {
		if(debug) {
			prints("runtime·sliceslice: old=");
			runtime·printslice(old);
			prints("; lb=");
			runtime·printint(lb);
			prints("; hb=");
			runtime·printint(hb);
			prints("; width=");
			runtime·printint(width);
			prints("\n");

			prints("oldarray: nel=");
			runtime·printint(old.len);
			prints("; cap=");
			runtime·printint(old.cap);
			prints("\n");
		}
        //如果容量超出或者是最大下标超出，则抛出异常
		throwslice(lb, hb, old.cap);
	}

	// new array is inside old array
	ret.len = hb - lb;
	ret.cap = old.cap - lb;
	ret.array = old.array + lb*width;
    ...
    ...
}
```

#### slice扩容再度分析：

[slice源码分析扩容机制](https://docs.google.com/presentation/d/1Ldeya8FtmOuwD1nsouVJ2YDDXXlsTph1QCr59j2hyTc/edit?usp=sharing)

#### slice grow函数扩容分析：


```cgo
//为s添加x...数组。
func Append(s Value, x ...Value) Value {
	//确定类型
    s.mustBe(Slice)
    //取长度比较，扩容决定
	s, i0, i1 := grow(s, len(x))
	for i, j := i0, 0; i < i1; i, j = i+1, j+1 {
		s.Index(i).Set(x[j])
	}
	return s
}
```

```cgo
// grow grows the slice s so that it can hold extra more values, allocating
// more capacity if needed. It also returns the old and new slice lengths.
func grow(s Value, extra int) (Value, int, int) {
	i0 := s.Len()
	i1 := i0 + extra
	if i1 < i0 {
		panic("reflect.Append: slice overflow")
	}
	m := s.Cap()
    //算上添加的新元素，没有超过当前的最大容量值，不需要扩容。
	if i1 <= m {
		return s.Slice(0, i1), i0, i1
	}
	if m == 0 {
		m = extra
	} else {
		for m < i1 {
			if i0 < 1024 {
				m += m
			} else {
				m += m / 4
			}
		}
	}
	t := MakeSlice(s.Type(), i1, m)
    //额。。。。很致命的一个copy，导致slice地址的改变。
	Copy(t, s)
	return t, i0, i1
}
```

#### slice Append Slice（致命copy）
```cgo
    
// AppendSlice appends a slice t to a slice s and returns the resulting slice.
// The slices s and t must have the same element type.
func AppendSlice(s, t Value) Value {
	s.mustBe(Slice)
	t.mustBe(Slice)
	typesMustMatch("reflect.AppendSlice", s.Type().Elem(), t.Type().Elem())
	s, i0, i1 := grow(s, t.Len())
    //值拷贝，并追加其后，意味着t中值得改变，不会影响s中的值
	Copy(s.Slice(i0, i1), t)
	return s
}
```


#### slice and array diff:
* 结构不同：
```cgo
    type array struct{
        data 
        len
    }
    type slice struct{
        data
        len
        cap
    }
```

* 扩容方式：
    array: 不可动态扩容。
    slice:  根据容量判断是否需要动态扩容
* 截取方式：
    slice & array: num[a:b:c]/num[a:b]
* 截取分析：
    [源码分析](https://github.com/golang/go/blob/master/src/reflect/value.go)
    [plan 9源码分析](https://plan9.io/sources/contrib/ericvh/go-plan9/src/pkg/runtime/slice.c)
    
#### 注意事项：
 ```cgo
1、致命copy函数
2、扩容的机制   >1024??
3、截取时候容量的变化。
```