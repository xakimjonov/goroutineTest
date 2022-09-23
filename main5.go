package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main(){
// 	channel := make(chan int)
// 	go getRandomNumber(channel)
// 	for randomNumber := range channel{

// 		fmt.Println("randomNumber", randomNumber)
// 	}
// }

// func getRandomNumber(channel chan int){
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 1 ; i <= 5; i++ {
// 		number := rand.Intn(10)
// 		time.Sleep(time.Second * 1)
// 		channel		<- number
// 	}
// 	close(channel)
// }