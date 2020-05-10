
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

### mallocgc过程分析

### slice 创建方式的区别

### slice make注意事项
