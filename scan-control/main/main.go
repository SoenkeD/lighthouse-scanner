package main

import (
	cliLib "lighthouse-scanner/cli"
	"lighthouse-scanner/cli/scripts"
	"lighthouse-scanner/lib"
	"log"
	"os"
)

func schema() *cliLib.CliSchema {
	cliOpt := cliLib.CliSchema{
		Top: []*cliLib.CliTopGroup{
			scripts.InitSchema(),
		},
	}
	cliOpt.Init()
	return &cliOpt
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Verify that a subcommand has been provided
	log.Println(os.Args)
	if len(os.Args) < 2 {
		log.Println("list or count subcommand is required")
		os.Exit(1)
	}

	cfg, cfgErr := lib.ReadConfig("../config/config.json")
	if cfgErr != nil {
		log.Fatalln(cfgErr)
	}

	cliOpt := schema()
	resp, err := cliOpt.Input(os.Args[1:], cfg)
	if err != nil {
		log.Println("some error occurred", err)
		os.Exit(1)
	}
	log.Println("success", resp)
	os.Exit(0)
}
