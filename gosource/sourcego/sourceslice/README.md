* [slice总结：](#slice总结)
   * [描述](#描述)
   * [slice和array区别](#slice和array区别)
      * [结构体定义](#结构体定义)
      * [创建方式](#创建方式)
      * [扩容？](#扩容)
      * [slice append](#slice-append)
      * [slice截取](#slice截取)
   * [slice复制](#slice复制)
   * [标准库之--slice](#标准库之--slice)
   * [other:](#other)
      * [逃逸分析](#逃逸分析)
      * [边界检查分析](#边界检查分析)
   * [参考资料：](#参考资料)
      * [官方文档:](#官方文档)
    
## slice总结：
### 描述
[Language Specification之slice定义](https://golang.org/ref/spec#Slice_types)

![](https://raw.githubusercontent.com/crab21/Images/master//blog/20200506110013.png)

### slice和array区别
#### 结构体定义
```cgo
type array struct{
    data
    len
}

type Slice struct{
    data
    len
    cap
}
```

#### 创建方式
> [slice创建方式](https://github.com/crab21/go-source/blob/master/gosource/sourcego/sourceslice/source_slice_make.md)
```cgo
make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m
```

常见错误：
```cgo
s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
var s []int    //直接使用 but:  s==nil
s:=make([]int,0,0)  // s!=nil
```

> array创建方式
```cgo
var s = [N]T{...}    // N一定为常数，N>=后面元素的个数
var s = [3]int{1,2,3}
```

#### 扩容？
[Issues# 6](https://github.com/crab21/go-source/issues/6)

[Slice grow](https://github.com/crab21/go-source/blob/master/gosource/sourcego/sourceslice/source_slice_grow.md)

#### slice append 
>[Append方式中的问题](https://github.com/crab21/go-source/blob/master/gosource/sourcego/sourceslice/source_slice_append.md)

推荐方式如下：
![](https://raw.githubusercontent.com/crab21/Images/master//blog/20200727230057.png)

#### slice截取
```cgo
var num = make([]int,100)
interce := num[A:B]  // A<=B interce属性==>len:B-A cap: oldcap-A 
interceThree := num[A:B:C] // A<=B B<=C 注：C为截止的下标元素；interceThree属性===> len: B-A cap:C-A
```
### slice复制
![](https://raw.githubusercontent.com/crab21/Images/master//blog/20200506113250.png)

### 标准库之--slice


### other:
#### 逃逸分析

1、[Language Mechanics On Escape Analysis](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html)

2、[Go 语言机制之逃逸分析](https://studygolang.com/articles/12444?fr=sidebar)

#### 边界检查分析 

1、[go101边界消除](https://gfw.go101.org/article/bounds-check-elimination.html)

2、[bounds check in go](https://www.ardanlabs.com/blog/2018/04/bounds-check-elimination-in-go.html)

3、[官方文档](https://docs.google.com/document/d/1vdAEAjYdzjnPA9WDOQ1e4e05cYVMpqSxJYZT33Cqw2g/edit#)

### 参考资料：
#### 官方文档:
1、[slice定义](https://golang.org/ref/spec#Slice_types)

2、[make slice](https://golang.org/ref/spec#Making_slices_maps_and_channels)

3、[Append and Copy slice](https://golang.org/ref/spec#Appending_and_copying_slices)

4、[Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices)