package scripts

import (
	"encoding/json"
	cliLib "lighthouse-scanner/cli"
	"lighthouse-scanner/lib"
	"log"
)

func initAnalyzerStart() *cliLib.CliCmd {

	startCmd := cliLib.CliCmd{
		Id:   "start",
		Args: []*cliLib.CliArg{},
		Run: func(cmd *cliLib.CliCmd, cfg lib.ConfigController) error {

			res := lib.InitLightHouseResults(cfg.OutputPath)
			ids, err := res.LoadAvailableResults()
			if err != nil {
				log.Fatalln(err)
			}

			stats := lib.InitLightHouseStats()
			for _, id := range ids {
				result, resultErr := res.LoadResult(id)
				if resultErr != nil {
					log.Fatalln(resultErr)
				}

				stats.AddResult(result)
			}

			jsonBytes, err := json.MarshalIndent(stats, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}

			log.Println(string(jsonBytes))

			return nil
		},
	}
	return &startCmd
}

func InitAnalyzer() *cliLib.CliCmdGroup {
	group := cliLib.CliCmdGroup{
		Id: "analyze",
		Cmd: []*cliLib.CliCmd{
			initAnalyzerStart(),
		},
	}
	return &group
}
