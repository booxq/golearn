package main

import(
	"fmt"
	"time"
	"sync"
	"strconv"
)

var wg sync.WaitGroup
var sum int

func say(s string){
	defer wg.Done()
	for i:= 0; i < 10; i++{
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s, i)
		sum++
	}
}
func say2(s string){
	defer wg.Done()
	for i:= 0; i < 10; i++{
		//time.Sleep(10 * time.Millisecond)
		fmt.Println(s, i)
		sum++
	}
}
func main(){
	for i := 0; i < 5; i++{
		wg.Add(1)
		go say2("helo"+ strconv.Itoa(i) )
	}
	fmt.Println(sum)
	wg.Wait()
	fmt.Println(sum)


	// 合起来写
    go func() {
        i := 0
        for {
            i++
            fmt.Printf("new goroutine: i = %d\n", i)
            time.Sleep(time.Second)
        }
    }()
    i := 0
    for {
        i++
        fmt.Printf("main goroutine: i = %d\n", i)
        time.Sleep(time.Second)
        if i == 2 {
            break
        }
	}
	defer func() {
		fmt.Println("main over")
	}()
}