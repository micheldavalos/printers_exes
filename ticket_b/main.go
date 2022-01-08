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
        pterm.NewLettersFromStringWithStyle("<<", pterm.NewStyle(pterm.FgWhite)),
        pterm.NewLettersFromStringWithStyle("Ticket B", pterm.NewStyle(pterm.FgYellow)),
		pterm.NewLettersFromStringWithStyle(">>", pterm.NewStyle(pterm.FgWhite))).
        Render()

	go func ()  {
		for {
			time.Sleep(time.Second)
		}
	}()
	
	var input string
	fmt.Scanln(&input)
}