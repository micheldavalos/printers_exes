package main

import (
	"fmt"
	"time"
	"github.com/pterm/pterm"
)

func main()  {
	// fmt.Println("Impresora Ticket C")
	pterm.DefaultBigText.WithLetters(
        pterm.NewLettersFromStringWithStyle("<<", pterm.NewStyle(pterm.FgWhite)),
        pterm.NewLettersFromStringWithStyle("Ticket C", pterm.NewStyle(pterm.FgLightRed)),
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