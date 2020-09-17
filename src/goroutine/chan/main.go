package main

import(
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func recv(c chan int){
	time.Sleep(time.Millisecond)
	for{
		
		if ret, ok := <-c; ok{
			fmt.Println("recv :", ret)
		} else {
			break
		}
		//wg.Done()
		
	}
}
func send(c chan int){
	for i :=0 ; i < 5 ; i++{
		//wg.Add(1)
		c <- i
		fmt.Println("send ok:", i)
	} 
	//close(c)
}
func main(){
	// ch := make(chan int)
	// go send(ch)
	
	// go recv(ch)
	ch1 := make(chan int)
    ch2 := make(chan int)
    // 开启goroutine将0~100的数发送到ch1中
    go func() {
        for i := 0; i < 100; i++ {
            ch1 <- i
        }
        close(ch1)
    }()
    // 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
    go func() {
        for {
            i, ok := <-ch1 // 通道关闭后再取值ok=false
            if !ok {
                break
            }
            ch2 <- i * i
        }
        close(ch2)
    }()
    // 在主goroutine中从ch2中接收值打印
    for i := range ch2 { // 通道关闭后会退出for range循环
        fmt.Println(i)
    }
	//wg.Wait()
	time.Sleep(2*time.Millisecond)
	defer func(){
		fmt.Println("main is over")
	}()
}