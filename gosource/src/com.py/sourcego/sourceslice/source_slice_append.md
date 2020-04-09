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
      0x05d5 01493 (source_slice.go:137)      MOVQ    $2, 16(SP)
      0x05de 01502 (source_slice.go:137)      MOVQ    $4, 24(SP)
      0x05e7 01511 (source_slice.go:137)      MOVQ    $5, 32(SP)
// 调用切片扩容？？依据呢？？
//todo 分析扩容原因和关键点
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
