package main

import (
	"fmt"
	"sync"
)

// 管道测试
// 基于管道实现生产者消费者模型

// 管道是引用类型，且必须有类型
// 管道需要先进行分配内存，即make后才能使用
// 管道的长度是不可变的
// 其使用价值在于一边放，一边取

// GoodChannel 商品管道
var GoodChannel chan int = make(chan int, 10)

// ProduceChannel 生产者阻塞管道
var ProduceChannel chan int = make(chan int, 1)

// ConsumptionChannel 消费者阻塞管道
var ConsumptionChannel chan bool = make(chan bool, 1)

// sync.WaitGroup提供了两个方法来实现主线程等待所有goroutine的功能
// (wg *WaitGroup) Add(delta int) 给计数器加delta
// (wg *WaitGroup) Done() 给计数器减1
// (wg *WaitGroup) Wait() 一直阻塞到计数器为0
var wg sync.WaitGroup

// Produce 生产者生产num个商品
func Produce(num int) {
	// 方法结束后，计数器减一
	defer wg.Done()
	for true {
		// 这里从channel中获取编号，channel为空，则阻塞
		// 意味着有另外一个生产者在进行操作
		proNum, ok := <-ProduceChannel
		// 如果channel被关闭，就退出任务
		if !ok {
			break
		}
		// 如果达到要生产的数量，那么关闭管道并退出
		// 关闭管道相当于向其他goroutine发信号
		if proNum > num {
			close(ProduceChannel)
			break
		}
		// 如果不是上面的两种情况，说明正常的生产商品，入商品通道
		fmt.Println("produce    ", proNum)
		GoodChannel <- proNum
		// 将要生产的下一个编号放入生产者channel
		ProduceChannel <- proNum + 1
	}
}

// Consumption 消费者消费一个商品
func Consumption(num int) {
	// 方法结束后，计数器减一
	defer wg.Done()
	for true {
		// 消费者通道中取出一个元素代表获取锁
		_, ok := <-ConsumptionChannel
		// 如果通道关闭，则说明完成了任务，所有goroutine退出
		if !ok {
			break
		}

		// 从商品channel中获取一个商品
		// 这里不需要先进行判断，因为能进入商品channel的商品都是符合生产者任务的
		goodNum := <-GoodChannel
		fmt.Println("Consumption", goodNum)

		// 当拿到对应序号的商品后
		// 关闭channel，说明所有消费者的任务都完成了
		if goodNum == num {
			close(ConsumptionChannel)
			break
		}
		// 如果没有完成所有任务，那么释放锁
		ConsumptionChannel <- true
	}
}

func main() {
	// 设置生产者的启动商品编号
	ProduceChannel <- 1
	// 设置消费者的锁
	ConsumptionChannel <- true

	// 启动若干个生产者和消费者协程
	for i := 0; i < 4; i++ {
		// 这里控制主线程等待所有goroutine完成，类似Java中的回环屏障
		wg.Add(1)
		go Produce(1000)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Consumption(1000)
	}
	// 回环屏障的阻塞
	wg.Wait()
}
