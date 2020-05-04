package sourceslice

import "testing"

func TestSliceInte_1(t *testing.T) {
	SliceInte_1()
}

/**
  0x0032 00050 (source_interception.go:42)        PCDATA  $0, $1
  0x0032 00050 (source_interception.go:42)        PCDATA  $1, $0
  0x0032 00050 (source_interception.go:42)        LEAQ    type.[3]int(SB), AX //leaq 把地址type.[3]int(SB)放到寄存器AX中
  0x0039 00057 (source_interception.go:42)        PCDATA  $0, $0
  0x0039 00057 (source_interception.go:42)        MOVQ    AX, (SP)
  0x003d 00061 (source_interception.go:42)        CALL    runtime.newobject(SB)
  0x0042 00066 (source_interception.go:42)        PCDATA  $0, $1
  0x0042 00066 (source_interception.go:42)        MOVQ    8(SP), AX
  0x0047 00071 (source_interception.go:42)        PCDATA  $1, $1
  0x0047 00071 (source_interception.go:42)        MOVQ    AX, ""..autotmp_10+104(SP)
  0x004c 00076 (source_interception.go:42)        PCDATA  $0, $0
  0x004c 00076 (source_interception.go:42)        MOVQ    $1, (AX) //AX添加第一个元素，移位操作
  0x0053 00083 (source_interception.go:42)        PCDATA  $0, $1
  0x0053 00083 (source_interception.go:42)        MOVQ    ""..autotmp_10+104(SP), AX
  0x0058 00088 (source_interception.go:42)        TESTB   AL, (AX)
  0x005a 00090 (source_interception.go:42)        PCDATA  $0, $0
  0x005a 00090 (source_interception.go:42)        MOVQ    $2, 8(AX)  //AX添加第一个元素，移位操作
  0x0062 00098 (source_interception.go:42)        PCDATA  $0, $1
  0x0062 00098 (source_interception.go:42)        MOVQ    ""..autotmp_10+104(SP), AX
  0x0067 00103 (source_interception.go:42)        TESTB   AL, (AX)
  0x0069 00105 (source_interception.go:42)        PCDATA  $0, $0
  0x0069 00105 (source_interception.go:42)        MOVQ    $3, 16(AX)   //第三个元素  AX寄存器地址，移位操作
  0x0071 00113 (source_interception.go:42)        PCDATA  $0, $1
  0x0071 00113 (source_interception.go:42)        PCDATA  $1, $0
  0x0071 00113 (source_interception.go:42)        MOVQ    ""..autotmp_10+104(SP), AX
  0x0076 00118 (source_interception.go:42)        TESTB   AL, (AX)
  0x0078 00120 (source_interception.go:42)        JMP     122
  0x007a 00122 (source_interception.go:42)        MOVQ    AX, "".number+232(SP)
  0x0082 00130 (source_interception.go:42)        MOVQ    $3, "".number+240(SP)
  0x008e 00142 (source_interception.go:42)        MOVQ    $3, "".number+248(SP)
  0x009a 00154 (source_interception.go:43)        JMP     156
  0x009c 00156 (source_interception.go:43)        PCDATA  $0, $-1
  0x009c 00156 (source_interception.go:43)        PCDATA  $1, $-1
  0x009c 00156 (source_interception.go:43)        JMP     158
  0x009e 00158 (source_interception.go:43)        PCDATA  $0, $1
  0x009e 00158 (source_interception.go:43)        PCDATA  $1, $0
// todo 当43行改为：result :=number[1:2]，则寄存器需要移动8byte						todo 当43行改为：result :=number[2:2]，则寄存器需要移动16byte
// 0x009e 00158 (source_interception.go:43)        ADDQ    $8, AX 					0x009e 00158 (source_interception.go:43)        ADDQ    $16, AX
  0x009e 00158 (source_interception.go:43)        MOVQ    AX, "".result+208(SP)
  0x00a6 00166 (source_interception.go:43)        MOVQ    $2, "".result+216(SP)
  0x00b2 00178 (source_interception.go:43)        MOVQ    $3, "".result+224(SP)

 */
func TestSliceInteAB(t *testing.T) {
	SliceInteAB()
}
/**
  0x0124 00292 (source_interception.go:50)        JMP     294
  0x0126 00294 (source_interception.go:50)        PCDATA  $0, $-1
  0x0126 00294 (source_interception.go:50)        PCDATA  $1, $-1
  0x0126 00294 (source_interception.go:50)        JMP     296
  0x0128 00296 (source_interception.go:50)        JMP     298
  0x012a 00298 (source_interception.go:50)        PCDATA  $0, $1
  0x012a 00298 (source_interception.go:50)        PCDATA  $1, $2
  0x012a 00298 (source_interception.go:50)        ADDQ    $8, AX
  0x012e 00302 (source_interception.go:50)        MOVQ    AX, "".result+480(SP)
  0x0136 00310 (source_interception.go:50)        MOVQ    $1, "".result+488(SP)  $1指当前长度
  0x0142 00322 (source_interception.go:50)        MOVQ    $5, "".result+496(SP)  $5指当前容量
  ......
  0x02f4 00756 (source_interception.go:54)        MOVQ    "".result+488(SP), AX
  0x02fc 00764 (source_interception.go:54)        MOVQ    AX, "".length+88(SP)
  ......
   0x051d 01309 (source_interception.go:57)        MOVQ    "".result+496(SP), AX
   0x0525 01317 (source_interception.go:57)        MOVQ    AX, "".caption+96(SP)


 */
func TestSliceInteABC(t *testing.T) {
	SliceInteABC()
}

func TestSliceAndArray(t *testing.T) {
	SliceAndArray()
}