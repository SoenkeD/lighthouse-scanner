package lib

import (
	"fmt"
	"lighthouse-scanner/lib"
	"strings"
)

type CliCmdGroup struct {
	Id  string
	Cmd []*CliCmd
}

func (group *CliCmdGroup) Init() {
	for _, cmd := range group.Cmd {
		cmd.InitFlagSet()
	}
}

func (group *CliCmdGroup) Input(args []string, cfg lib.ConfigController) error {

	// list available commands
	var cmdList []string
	for _, cmd := range group.Cmd {
		cmdList = append(cmdList, cmd.Id)
	}
	if displayHelp(args, strings.Join(cmdList, ", "), 0) {
		return nil
	}

	if len(args) < 1 {
		return fmt.Errorf("to few arguments provided")
	}

	for _, cmd := range group.Cmd {
		if cmd.Id == args[0] {
			return cmd.Input(args[1:], cfg)
		}
	}

	return fmt.Errorf("cmd=%s was not found", args[0])
}
