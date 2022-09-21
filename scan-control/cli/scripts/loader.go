package scripts

import (
	"encoding/json"
	cliLib "lighthouse-scanner/cli"
	"lighthouse-scanner/lib"
	"log"
	"os"
	"sync"
	"time"
)

func runAnalysis(url string, nodePath string, lightHouseClient string, startDir string, chromeBin string) (resp lib.LightHouseResp, err error) {

	rawResp, err := lib.RunCmd(nodePath, []string{lightHouseClient, url, chromeBin}, startDir+"/..")
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResp, &resp)
	if err != nil {
		return
	}

	return
}

func initLoaderStart() *cliLib.CliCmd {

	startCmd := cliLib.CliCmd{
		Id:   "start",
		Args: []*cliLib.CliArg{},
		Run: func(cmd *cliLib.CliCmd, cfg lib.ConfigController) error {

			// get context dir
			startDir, pwdErr := os.Getwd()
			if pwdErr != nil {
				log.Fatalln(pwdErr)
			}

			res := lib.InitLightHouseResults(cfg.OutputPath)

			urlQ := make(chan string, cfg.UrlChannelBufferSize)
			defer close(urlQ)

			var wg sync.WaitGroup
			wg.Add(len(cfg.Urls))
			for idx := 0; idx < cfg.AgentCount; idx++ {
				go func() {
					for url := range urlQ {

						log.Println("starting analyzing url=" + url)
						timeStart := time.Now()

						resp, err := runAnalysis(url, cfg.NodePath, cfg.ClientPath, startDir, cfg.ChromeBin)
						if err != nil {
							log.Println("analysis for url="+url+" failed", err)
							wg.Done()
							continue
						}

						timeEnd := time.Now()
						result, err := res.AddResp(url, timeEnd.Sub(timeStart), resp)
						if err != nil {
							log.Println("failed to dump result to file", err)
							wg.Done()
							continue
						}

						log.Println("finished analyzing of url=" + url + " with guuid=" + result.Id)
						wg.Done()
					}
				}()
			}

			for _, url := range cfg.Urls {
				urlQ <- url
			}

			wg.Wait()

			for _, result := range res.GetResults() {
				log.Println(result)
			}

			return nil
		},
	}
	return &startCmd
}

func InitLoad() *cliLib.CliCmdGroup {
	group := cliLib.CliCmdGroup{
		Id: "load",
		Cmd: []*cliLib.CliCmd{
			initLoaderStart(),
		},
	}
	return &group
}
