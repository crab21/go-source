### G-P-M调度模型
>参考---[runtime/proc.go](https://golang.org/src/runtime/proc.go)

* G(goroutine):用户轻量级线程，每个goroutine对象中的sched保存中上下文信息。

* P(processor):即G和M的调度对象，用来调度G和M的关系，可以用GOMAXPROCS()来设置，默认为核心数。

* M(machine)：对内核级线程的封装，数量==真实CPU数量。


Design doc at:[设计文档](https://golang.org/s/go11sched)

