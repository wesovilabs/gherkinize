package gherkin

import (
	"fmt"
	"container/list"
	"github.com/wesovilabs/gherkinize/util"
	"github.com/wesovilabs/gherkinize/config"
)



func (feature *Feature) validate_max_steps(max_steps int) {
	var next *list.Element
	for scenario := feature.Scenarios.Front(); scenario != nil; scenario = next {
		next = scenario.Next()
		if scenario.Value.(Scenario).Steps.Len() > max_steps {
			feature.Successful = false
			util.Print_error("Line: %d", scenario.Value.(Scenario).LineNumber)
			util.Print_error("Scenario: %s", scenario.Value.(Scenario).Text)
			util.Print_error("Number of steps: %d", scenario.Value.(Scenario).Steps.Len())
			util.Print_error("Max. number of steps: %d", max_steps)
			fmt.Println()
		}
	}
}

func (feature *Feature) validate_steps_length(steps_length int){
	var next *list.Element
	for scenario := feature.Scenarios.Front(); scenario != nil; scenario = next {
		next = scenario.Next()
		var nextStep *list.Element
		for step := scenario.Value.(Scenario).Steps.Front(); step != nil; step = nextStep {
			nextStep = step.Next()
			if len(step.Value.(Step).Text) > steps_length {
				feature.Successful = false
				util.Print_error("Line: %d", step.Value.(Step).LineNumber)
				util.Print_error("Scenario: %s", scenario.Value.(Scenario).Text)
				util.Print_error("Step: %s", step.Value.(Step).Text)
				util.Print_error("Length: %d", len(step.Value.(Step).Text))
				util.Print_error("Max. length for steps: %d", steps_length)
				fmt.Println()
			}
		}
	}
}

func (feature *Feature) validate_empty_feature() {
	if(feature.Scenarios.Len() <= 0){
		feature.Successful = false
		util.Print_error("Missing scenarios.")
		fmt.Println()
	}
}

func (feature *Feature) validate_empty_scenarios() {
	if(feature.Scenarios.Len() > 0){
		var next *list.Element
		for scenario := feature.Scenarios.Front(); scenario != nil; scenario = next {
			next = scenario.Next()
			if(scenario.Value.(Scenario).Steps.Len() <=0){
				feature.Successful = false
				util.Print_error("Line: %d", scenario.Value.(Scenario).LineNumber)
				util.Print_error("Scenario: %s", scenario.Value.(Scenario).Text)
				util.Print_error("Missing steps.")
				fmt.Print()
				fmt.Print()
			}
		}
	}
}

func (feature *Feature) validate_strict(){

	if(feature.Scenarios.Len() <= 0){
		feature.Successful = false
	} else {

		var next *list.Element
		for scenario := feature.Scenarios.Front(); scenario != nil; scenario = next {
			if(scenario.Value.(Scenario).Steps.Len() >0) {
				stage := 0
				next = scenario.Next()
				var nextStep *list.Element
				for step := scenario.Value.(Scenario).Steps.Front(); step != nil; step = nextStep {
					nextStep = step.Next()
					switch step.Value.(Step).Kind {
					case TOKEN_AND:
						if (stage == 0) {
							feature.Successful = false
							displayError(step.Value.(Step), scenario.Value.(Scenario).Text, "The first keyword ina  scenary must be Given.")
						}
					case TOKEN_GIVEN:
						if (stage != 0) {
							feature.Successful = false
							displayError(step.Value.(Step), scenario.Value.(Scenario).Text, "Given keyword can only be used once at the beggining of a sceneario.")
						}
						stage = 1
					case TOKEN_WHEN:
						if (stage != 1 && stage != 3) {
							feature.Successful = false
							displayError(step.Value.(Step), scenario.Value.(Scenario).Text, "When keyword can only be used after Given or Then statements.")
						}
						stage = 2
					case TOKEN_THEN:
						if (stage != 2) {
							feature.Successful = false
							displayError(step.Value.(Step), scenario.Value.(Scenario).Text, "Then keyword can only be used after When statements.")
						}
						stage = 3
					}
				}
				if (stage < 3) {
					feature.Successful = false
					displayScenarioError(scenario.Value.(Scenario), scenario.Value.(Scenario).Text, "The scenario must end with a Then statement.")
				}
			} else {
				break
			}
		}

	}
}

func displayError(step Step, scenario string, message string){
	util.Print_error("Line: %d", step.LineNumber)
	util.Print_error("Scenario: %s", scenario)
	util.Print_error("Step: %s", step.Text)
	util.Print_error(message)
	fmt.Println()
}

func displayScenarioError(scenario Scenario, text string, message string){
	util.Print_error("Line: %d", scenario.LineNumber)
	util.Print_error("Scenario: %s", text)
	util.Print_error(message)
	fmt.Println()
}

func (feature *Feature) Validate(config config.Config) {

	fmt.Println()
	util.Print_subtitle(feature.Text)
	fmt.Println()
	feature.validate_max_steps(config.Errors.MaxStepsPerScenario)
	feature.validate_steps_length(config.Errors.MaxLenStep)
	if (!config.Errors.EmptyFeature) {
		feature.validate_empty_feature()
	}
	if (!config.Errors.EmptyScenario) {
		feature.validate_empty_scenarios()
	}
	if (config.Errors.Strict) {
		feature.validate_strict()
	}
	if(feature.Successful){
		util.Print_success("Feature is OK.")
	}
}
