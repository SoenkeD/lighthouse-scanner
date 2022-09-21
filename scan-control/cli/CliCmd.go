package lib

import (
	"flag"
	"lighthouse-scanner/lib"
)

type CliCmd struct {
	Id      string
	Args    []*CliArg
	Run     func(cmd *CliCmd, cfg lib.ConfigController) error
	flagSet *flag.FlagSet
}

func (cmd *CliCmd) InitFlagSet() {
	cmd.flagSet = flag.NewFlagSet(cmd.Id, flag.ExitOnError)
	for _, arg := range cmd.Args {
		arg.InitVal()
		if arg.SetOnly {
			cmd.flagSet.BoolVar(arg.SetValue, arg.Id, arg.Default == "true", arg.Description)
		} else {
			cmd.flagSet.StringVar(arg.Value, arg.Id, arg.Default, arg.Description)
		}
	}
}

func (cmd *CliCmd) GetValue(valId string) string {
	for _, cmd := range cmd.Args {
		if cmd.Id == valId {
			return *cmd.Value
		}
	}
	return ""
}

func (cmd *CliCmd) GetSetValue(valId string) bool {
	for _, cmd := range cmd.Args {
		if cmd.Id == valId {
			return *cmd.SetValue
		}
	}
	return false
}

func (cmd *CliCmd) Input(args []string, cfg lib.ConfigController) error {
	parseErr := cmd.flagSet.Parse(args)
	if parseErr != nil {
		return parseErr
	}

	return cmd.Run(cmd, cfg)
}
