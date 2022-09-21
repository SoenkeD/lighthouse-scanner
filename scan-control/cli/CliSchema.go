package lib

import (
	"fmt"
	"lighthouse-scanner/lib"
	"strings"
)

type CliSchema struct {
	Top []*CliTopGroup
}

func (schema *CliSchema) Init() {
	for _, group := range schema.Top {
		group.Init()
	}
}

func (schema *CliSchema) Input(args []string, cfg lib.ConfigController) (any, error) {

	// list available commands
	var cmdList []string
	for _, group := range schema.Top {
		cmdList = append(cmdList, group.Id)
	}
	if displayHelp(args, strings.Join(cmdList, ", "), 0) {
		return "", nil
	}

	for _, group := range schema.Top {
		if group.Id == args[0] {
			return group.Input(args[1:], cfg), nil
		}
	}
	return "", fmt.Errorf("top group=%s was not found", args[0])
}
