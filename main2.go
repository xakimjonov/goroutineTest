package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		display("How are you")
// 		wg.Done()
// 	}()

// 	go func() {
// 		display("I'm ok")
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	fmt.Println("completed")
// }

// func display(input string) {
// 	for i := 1; i <= 100; i++ {
// 		fmt.Println(i, input)
// 		time.Sleep(1 * time.Second)

// 	}

// }
