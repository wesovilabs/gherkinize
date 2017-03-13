package main

import (
	"github.com/wesovilabs/gherkinize/path"
	"os"
	"github.com/wesovilabs/gherkinize/gherkin"
	"container/list"
	"github.com/wesovilabs/gherkinize/util"
	"fmt"
	"github.com/fatih/color"
	"github.com/wesovilabs/gherkinize/config"
	"github.com/urfave/cli"
	"sort"
)

func main() {
	app := cli.NewApp()
	app.Name = "Gherkinize"
	app.Usage = "Find the issues in your Gherkin features."
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Iván Corrales Solera",
			Email: "developer@wesovilabs.com",
		},
	}
	app.Copyright = "(c) 2017 Wesovilabs"
	app.EnableBashCompletion = true
	var configurationFile string
	var scenariosPath string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load toml configuration from `FILE`",
			Destination: &configurationFile,
		},
		cli.StringFlag{
			Name:  "input, i",
			Usage: "Scenarios directory path",
			Destination: &scenariosPath,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "validate",
			Aliases: []string{"v"},
			Usage:   "validate the scenarios",
			Action:  func(c *cli.Context) error {
				if( scenariosPath == ""){
					return cli.NewExitError("Please specify the scenarios path by using option -i", 86)
				}
				if(configurationFile == ""){
					configurationFile = "./config/gherkin-rules.toml"
				}
				validateFiles(scenariosPath,configurationFile)
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)

}

func showInvalidStructureMessage(feature string, lineNumber int){
	fmt.Println()
	util.Print_subtitle(feature)
	fmt.Println()
	util.Print_error("Line: %d", lineNumber)
	util.Print_error("Missing scenario keyword.")
	color.White("[ USAGE ]")
	util.Print_message("\tScenario:")
	util.Print_message("\t   Given ...")
	util.Print_message("\t   When ...")
	util.Print_message("\t   Then ...")
}

func validateFiles(scenarios_path string, config_path string) error {
	gherkin.InitializeTokenMap()
	list_files := path.ReadDirectory(scenarios_path)
	configFile, err := os.Open(config_path)
	if err != nil {
		return nil
	}
	config := config.GetConfig(configFile)
	for _, filePath := range list_files {

		file, err := os.Open(filePath)
		if err != nil {
			return nil
		}



		defer file.Close()
		scanner := gherkin.NewGherkinScanner(file)
		lineNumber :=1
		tokens := list.New()
		fmt.Println()
		fmt.Println()
		util.Print_title("-------------------------------------------------- ")
		util.Print_title("File " + filePath)
		util.Print_title("-------------------------------------------------- ")
		var feature = &gherkin.Feature{}
		var invalidFeature = false
		for {
			token := scanner.Scan(lineNumber)
			switch token.Kind {
				case gherkin.EOF:
					feature.Validate(*config)
				case gherkin.NEW_LINE:
					if(feature.LineNumber == 0) {
						util.Print_error("Missing feature tag.")
						invalidFeature = true
						break
					} else {
						if (feature.Scenarios.Back() != nil){
							step := token.ToStep()
							feature.Scenarios.Back().Value.(gherkin.Scenario).Steps.PushBack(step)
						}
						lineNumber+=1
					}
				case gherkin.TOKEN_FEATURE:
					*feature = token.ToFeature()
				case gherkin.TOKEN_SCENARIO:
					if(feature.LineNumber == 0) {
						break
					}
					scenario := token.ToScenario()
					feature.Scenarios.PushBack(scenario)
				case gherkin.TOKEN_GIVEN:
					step := token.ToStep()
					if(feature.Scenarios.Back() == nil){
						showInvalidStructureMessage(feature.Text, token.LineNumber)
						invalidFeature = true
						break
					} else {
						feature.Scenarios.Back().Value.(gherkin.Scenario).Steps.PushBack(step)
					}
				case gherkin.TOKEN_WHEN:
					step := token.ToStep()
					if(feature.Scenarios.Back() == nil){
						showInvalidStructureMessage(feature.Text, token.LineNumber)
						invalidFeature = true
						break
					} else {
						feature.Scenarios.Back().Value.(gherkin.Scenario).Steps.PushBack(step)
					}
				case gherkin.TOKEN_THEN:
					step := token.ToStep()
					if(feature.Scenarios.Back() == nil){
						showInvalidStructureMessage(feature.Text, token.LineNumber)
						invalidFeature = true
						break
					} else {
						feature.Scenarios.Back().Value.(gherkin.Scenario).Steps.PushBack(step)
					}
				case gherkin.TOKEN_AND:
					step := token.ToStep()
					if(feature.Scenarios.Back() == nil){
						showInvalidStructureMessage(feature.Text, token.LineNumber)
						invalidFeature = true
						break
					} else {
						feature.Scenarios.Back().Value.(gherkin.Scenario).Steps.PushBack(step)
					}
			}
			if(invalidFeature){
				break
			}
			if(token.IsKeyword()){
				tokens.PushBack(token)
				lineNumber++
			}
			if(token.Kind == gherkin.EOF) {
				break
			}
		}


	}

	return nil
}