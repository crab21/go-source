
* [分析Slice grow过程](#分析slice-grow过程)
   * [简单过程分析](#简单过程分析)
   * [竟态变量介绍](#竟态变量介绍)
      * [raceenabled](#raceenabled)
      * [msanenabled](#msanenabled)
   * [内存模型](#内存模型)
      * [内存计算](#内存计算)
      * [内存分配方式](#内存分配方式)
   * [内存溢出](#内存溢出)
      * [判断](#判断)
      * [处理方式](#处理方式)
   * [正常处理结果](#正常处理结果)
      * [值](#值)
      * [len](#len)
      * [cap](#cap)
   * [grow 总结](#grow-总结)
      * [处理逻辑](#处理逻辑)
      * [错误处理方式](#错误处理方式)

# 分析Slice grow过程

## 简单过程分析

>[见Slice grow函数扩容分析](https://github.com/crab21/go-source/blob/master/gosource/src/com.py/sourcego/sourceslice/source_slice_append.md#slice-grow%E5%87%BD%E6%95%B0%E6%89%A9%E5%AE%B9%E5%88%86%E6%9E%90)

## 竟态变量介绍

### raceenabled

### msanenabled

## 内存模型
### 内存计算
### 内存分配方式

## 内存溢出
### 判断
### 处理方式

## 正常处理结果
### 值
### len
### cap

## grow 总结

### 处理逻辑
### 错误处理方式

