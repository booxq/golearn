package main

import(
	"fmt"
	"time"
)

func main(){
	// 
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	//for{
		t2 := <-timer1.C
		fmt.Printf("t2:%v\n", t2)
	//}
	t3 := <-time.After(2*time.Second)
	fmt.Printf("t3:%v\n",t3)
	
	ticker := time.NewTicker(10*time.Millisecond)
	// timer10 := time.NewTimer(10000*time.Millisecond)
	for{
		fmt.Println(<-ticker.C)
		go main()
	}
}