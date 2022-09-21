package lib

import (
	"log"
)

func displayHelp(args []string, helpTxt string, maxIdx int) bool {
	for idx, arg := range args {
		if arg == "--help" && idx <= maxIdx {
			log.Println("available: " + helpTxt)
			return true
		}
	}
	return false
}
