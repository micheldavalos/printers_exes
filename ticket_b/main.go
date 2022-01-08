package main

import (
	"fmt"
	"time"
	"github.com/pterm/pterm"
)

func main()  {
	// fmt.Println("Impresora Ticket B")
	// Print a large text with differently colored letters.
    pterm.DefaultBigText.WithLetters(
        pterm.NewLettersFromStringWithStyle("I:", pterm.NewStyle(pterm.FgLightGreen)),
        pterm.NewLettersFromStringWithStyle("Ticket B", pterm.NewStyle(pterm.FgLightYellow))).
        Render()

	go func ()  {
		for {
			time.Sleep(time.Second)
		}
	}()
	
	var input string
	fmt.Scanln(&input)
}