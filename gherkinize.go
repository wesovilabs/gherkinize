package main

import (
	"fmt"
	"github.com/wesovilabs/gherkinize/path"
	"os"
	"github.com/wesovilabs/gherkinize/gherkin"
	"log"
	"container/list"
)

func main() {
	validateFiles("./testdata/scenarios/")
}

func validateFiles(scenarios_path string) error {
	gherkin.InitializeTokenMap()
	list_files := path.ReadDirectory(scenarios_path)
	tokens := list.New()

	for _, filePath := range list_files {
		log.Print(filePath)
		file, err := os.Open(filePath)
		if err != nil {
			log.Print("Error:", err)
			return nil
		}
		scanner := gherkin.NewGherkinScanner(file)
		lineNumber :=1
		var feature gherkin.Feature
		for {
			token := scanner.Scan(lineNumber)
			switch token.Kind {
				case gherkin.EOF:
					log.Print("EOFile")
					fmt.Println(tokens)
					var next *list.Element
					for e := tokens.Front(); e != nil; e = next {
						next = e.Next()
						log.Println(e.Value)
						tokens.Remove(e)
					}
					fmt.Println(feature.ToString())
					return nil
				case gherkin.NEW_LINE:
					lineNumber+=1
				case gherkin.TOKEN_FEATURE:
					feature = token.ToFeature()
				case gherkin.TOKEN_SCENARIO:
					scenario := token.ToScenario()
					feature.Scenarios.PushBack(scenario)
				case gherkin.TOKEN_GIVEN:
					step := token.ToStep()
					feature.Scenarios.Back().Value.(gherkin.Scenario).Steps.PushBack(step)
				case gherkin.TOKEN_WHEN:
					step := token.ToStep()
					feature.Scenarios.Back().Value.(gherkin.Scenario).Steps.PushBack(step)
				case gherkin.TOKEN_THEN:
					step := token.ToStep()
					feature.Scenarios.Back().Value.(gherkin.Scenario).Steps.PushBack(step)
			}
			if(token.IsKeyword()){
				tokens.PushBack(token)
				log.Print(token.ToString())
				lineNumber++
			}
		}


	}

	return nil
}