package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Impresora Ticket A")
	
	go func ()  {
		for {
			time.Sleep(time.Second)
		}
	}()
	
	var input string
	fmt.Scanln(&input)
}