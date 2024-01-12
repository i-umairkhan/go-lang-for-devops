package main

import "fmt"
import "time"

type Order struct {
	ID int
	Status string
}

func processOrder(orderID int,statusChannel chan Order){
	time.Sleep(time.Second * 2)
	statusChannel <- Order{ID: orderID, Status: "Completed"}
}

func main(){
	// Channels 

	// used to communicated between goroutines(thread) 

	// can be sync or async

	statusChannel := make(chan Order)

	for i := 1; i <= 5; i++ {
		go processOrder(i, statusChannel) 
	}

	for i := 1; i <= 5; i++ {
		processedOrder:= <- statusChannel
		fmt.Printf("Order %d processed: %s\n", processedOrder.ID, processedOrder.Status)
	}

}
