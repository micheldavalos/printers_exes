package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main()  {
	cmd := "START"
	arg1 := "\"\""
	
	a := "C:\\Users\\Michel Davalos\\Desktop\\Network ticket_a.exe.lnk"

	fmt.Println("Corriendo Ticket A")
	ticket_a := exec.Command("cmd", cmd, arg1, a) 
	err := ticket_a.Run()
	if err != nil {
		log.Fatal(err)
	}

	var input string
	fmt.Scanln(&input) 
}