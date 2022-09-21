package lib

type ConfigController struct {
	NodePath             string
	OutputPath           string
	ClientPath           string
	ChromeBin            string
	AgentCount           int
	UrlChannelBufferSize int
	Urls                 []string
}

func ReadConfig(cfgPath string) (conf ConfigController, err error) {
	conf, err = FileParse[ConfigController](cfgPath)
	return
}
