
* [pool简介](#pool简介)
* [使用场景](#使用场景)
* [注意事项](#注意事项)
* [源码分析](#源码分析)
      * [初始化函数--&gt;GC回收分析：](#初始化函数--gc回收分析)
      * [线程安全吗？可以在多个goroutine中同时使用吗？](#线程安全吗可以在多个goroutine中同时使用吗)
   * [那么也产生了一个新的问题：](#那么也产生了一个新的问题)

## pool简介
>a pool is a set of temporary objects that may be individually saved and retrieved.--[go package](https://golang.org/pkg/sync/#pool)

>线程安全：A Pool is safe for use by multiple goroutines simultaneously.

## 使用场景
   
   * 临时对象存取
   
>Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector.

## 注意事项
   * 默认两分钟触发一次GC，会不定期回收pool中的临时变量
   * So:
        * 不适合存放长期使用的对象：socket长连接之类的持久对象

## 源码分析

#### 初始化函数-->GC回收分析：
```cgo
func init() {
	runtime_registerPoolCleanup(poolCleanup)
}
```

从上面可以看出来注册了runtime的清空函数：
```cgo

func poolCleanup() {
	// This function is called with the world stopped, at the beginning of a garbage collection.
	// It must not allocate and probably should not call any runtime functions.

	// Because the world is stopped, no pool user can be in a
	// pinned section (in effect, this has all Ps pinned).

	// Drop victim caches from all pools.
    //清空old缓存
	for _, p := range oldPools {
		p.victim = nil
		p.victimSize = 0
	}

	// Move primary cache to victim cache.
	for _, p := range allPools {
		p.victim = p.local
		p.victimSize = p.localSize
		p.local = nil
		p.localSize = 0
	}

	// The pools with non-empty primary caches now have non-empty
	// victim caches and no pools have primary caches.
    //把目前现存的搬运到old上面去
	oldPools, allPools = allPools, nil
}
```
>至于这个oldPools和allPools是两个指针数组：[]*pool

再往下：
allPools何时被赋值？
sure,在put时候被赋值

```cgo



func (p *Pool) pinSlow() (*poolLocal, int) {
	....
    ....
	if p.local == nil {
		allPools = append(allPools, p)
	}
	....
    ....
}
```

#### 线程安全吗？可以在多个goroutine中同时使用吗？
>A Pool is safe for use by multiple goroutines simultaneously.--[go package](https://golang.org/pkg/sync/#pool)


涉及到另外一个知识点：调度模型[【PGM】](https://medium.com/@riteeksrivastava/a-complete-journey-with-goroutines-8472630c7f5c)

>P和M绑定关系：
```cgo
// pin pins the current goroutine to P, disables preemption and
// returns poolLocal pool for the P and the P's id.
// Caller must call runtime_procUnpin() when done with the pool.
func (p *Pool) pin() (*poolLocal, int) {
	pid := runtime_procPin()
	// In pinSlow we store to local and then to localSize, here we load in opposite order.
	// Since we've disabled preemption, GC cannot happen in between.
	// Thus here we must observe local at least as large localSize.
	// We can observe a newer/larger local, it is fine (we must observe its zero-initialized-ness).
	s := atomic.LoadUintptr(&p.localSize) // load-acquire
	l := p.local                          // load-consume
	if uintptr(pid) < s {
		return indexLocal(l, pid), pid
	}
	return p.pinSlow()
}
```

>***pool在put时根据P绑定M值，把相应的G中值，放到相应的地址中.
pool在get时P会绑定M.***


因此，对于pool来说是线程安全。

### 那么也产生了一个新的问题：
* goroutine调度模型？高效解决的中心思想？
* 线程安全的产生？定义？解决方法？

[go调度模型](https://docs.google.com/presentation/d/1YyJLB8ihCB0T4FS01W38okL0FwQAdtRQpquNF-1mU3A/edit?usp=sharing)