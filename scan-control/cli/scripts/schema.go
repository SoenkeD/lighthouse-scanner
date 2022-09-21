package scripts

import lib "lighthouse-scanner/cli"

func InitSchema() *lib.CliTopGroup {
	ret := lib.CliTopGroup{
		Id: "cli",
		Groups: []*lib.CliCmdGroup{
			InitLoad(),
			InitAnalyzer(),
		},
	}
	return &ret
}
