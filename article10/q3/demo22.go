package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v...%d \n", i, len(ch1))
			ch1 <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// 接收方。
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")
}

/*
老师，根据您的提供的源码有三个问题需请教。
第一问题:第一次携程调度应该发生主携程中的elem, ok := <-ch1 这个代码处，这时候应该在chan有等待的协成，再第一向chan1<-i传值得时候，根据您的描述"当发送操作在执行的时候发现空的通道中，正好有等待的接收,那么会把元素直接复制给对方"。照这么说应该在这次就跳转到主协成中，并打印出接收到的数据了。但是实际是先发送i=3的时候才做第一次携程调度，请问这是为什么？
第二问题：缓存区的大小不是设置的是2么，为什么length当发送了i=3的时候才会阻塞发生调度呢，正常不是应该i=2的时候么
第三个问题：当for循环结束了以后 就是在chan关闭之前，为什么又能调度到主协成让他接收呢。不应该到这个协成调用结束么？
作者回复: 所以说不要用“协程”这个概念，因为“协程（coroutine）”指的是程序在同一个线程内的自行调度，是应用程序本身完全可控的。而 goroutine 的调度是 Go 语言的运行时系统发起的。

你不要揣测 Go 语言的调度器会怎样调度。你首先要知道哪些代码点是调度的时机（注意，到了调度时机也不一定发生调度，只是时机而已）。你还要知道如果想让多个 goroutine 按照你拟定的流程执行就需要用到 Channel 以及各种同步工具。

你说的“跳转到”只能在 coroutine 场景下才能这么说。在 goroutine 的场景下，没有“跳转”这么一说。

其一，你在上面的 for 语句中启用了一个 goroutine，你怎么就能断定后面的代码一定会先于这个 go 函数执行？不要做这种假设。因为连 goroutine 的调度都是并发的。

其二，两个 goroutine 一个 channel，一个 goroutine 发，一个 goroutine 取。这个 ch1 什么时候满、什么时候空，你基本上是确定不了的。因为两个 for 循环 在迭代的过程中都可能因被调用而换下 CPU。

其三，你要知道，几乎任何函数调用都存在调度时机，更何况是像 fmt.Println 这种需要 I/O 的重型操作。所以，为什么你那前一个 for 循环结束之后就不能被调度了呢？

以上是我通过你的文字表达猜测并回答的，并不一定完全匹配你要问的问题。还有问题的话再问我。

我觉得你对“并发”和“调度”这两个概念不清楚。我建议你好好看看专栏里讲 goroutine 的那几篇文章。有必要的话，买我的《Go 并发编程实战》第二版从头学一下。


*/
