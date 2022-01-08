package main

import (
	"fmt"
	"time"
	"github.com/pterm/pterm"
)

func main()  {
	// fmt.Println("Impresora Ticket A")
	// Print a large text with differently colored letters.
    pterm.DefaultBigText.WithLetters(
        pterm.NewLettersFromStringWithStyle("<<", pterm.NewStyle(pterm.FgWhite)),
        pterm.NewLettersFromStringWithStyle("Ticket A", pterm.NewStyle(pterm.FgCyan)),
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