# Channel

> Don't communicate by sharing memory; share memory by communicating.不要通过共享内存来通信，应该通过通信来共享内存。

channel目的是相当与一个先进先出的队列，channel中的元素是按照发送顺序排序的，元素的发送和接收使用操作符`<-`，默认容量是1。

    var a int
    Ch := make(chan int)
    Ch <- 1  // 1传入Ch
    a <- Ch  // Ch中的数据传给a

channel是阻塞的，下游从channel取走数据，上游才能传入一个数据，否则代码会一直阻塞等待直到channel被清空，需要一个生成方，一个消费方。