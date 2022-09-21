package lib

import (
	"fmt"
	"lighthouse-scanner/lib"
	"strings"
)

type CliTopGroup struct {
	Id     string
	Groups []*CliCmdGroup
}

func (top *CliTopGroup) Init() {
	for _, group := range top.Groups {
		group.Init()
	}
}
func (top *CliTopGroup) Input(args []string, cfg lib.ConfigController) error {

	// list available commands
	var cmdList []string
	for _, groups := range top.Groups {
		cmdList = append(cmdList, groups.Id)
	}
	if displayHelp(args, strings.Join(cmdList, ", "), 0) {
		return nil
	}

	if len(args) < 2 {
		return fmt.Errorf("to few arguments provided")
	}

	for _, group := range top.Groups {
		if group.Id == args[0] {
			return group.Input(args[1:], cfg)
		}
	}

	return fmt.Errorf("cmd group=%s was not found", args[0])
}
