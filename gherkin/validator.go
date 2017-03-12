package gherkin

import (
	"container/list"
	"fmt"
	"github.com/wesovilabs/gherkinize/util"
	"github.com/wesovilabs/gherkinize/config"
)



func (feature *Feature) validate_max_steps(max_steps int) {
	var next *list.Element
	for scenario := feature.Scenarios.Front(); scenario != nil; scenario = next {
		next = scenario.Next()
		if scenario.Value.(Scenario).Steps.Len() > max_steps {

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
				util.Print_error("Line: %d", scenario.Value.(Scenario).LineNumber)
				util.Print_error("Scenario: %s", scenario.Value.(Scenario).Text)
				util.Print_error("A valid scenario must contains steps.")
			}
		}
	}
}

func (feature *Feature) validate_lines_between_steps(){
	var next *list.Element
	for scenario := feature.Scenarios.Front(); scenario != nil; scenario = next {
		next = scenario.Next()
		var nextStep *list.Element
		for step := scenario.Value.(Scenario).Steps.Front(); step != nil; step = nextStep {
			nextStep = step.Next()
			if(step.Value.(Step).Kind == NEW_LINE){
				util.Print_warning("Line %d", scenario.Value.(Scenario).LineNumber)
				util.Print_warning("Scenario %s", scenario.Value.(Scenario).Text)
				util.Print_warning("No empty lines are allowed.")
				fmt.Println()
			}
		}
	}
}

func (feature *Feature) validate_strict(){

	if(feature.Scenarios.Len() <= 0){
		util.Print_error("Missing scenarios.")
	} else {
		stage := 0
		var next *list.Element
		for scenario := feature.Scenarios.Front(); scenario != nil; scenario = next {
			next = scenario.Next()
			var nextStep *list.Element
			for step := scenario.Value.(Scenario).Steps.Front(); step != nil; step = nextStep {
				nextStep = step.Next()
				switch step.Value.(Step).Kind {
					case TOKEN_AND:
					if(stage==0){
						displayError(step.Value.(Step),scenario.Value.(Scenario).Text,"The first keyword ina  scenary must be Given.")
					}
					case TOKEN_GIVEN:
						if(stage!=0){
							displayError(step.Value.(Step),scenario.Value.(Scenario).Text,"Given keyword can only be used once at the beggining of a sceneario.")
						}
						stage = 1
					case TOKEN_WHEN:
						if(stage!=1 && stage!=3){
							displayError(step.Value.(Step),scenario.Value.(Scenario).Text,"When keyword can only be used after Given or Then statements.")
						}
						stage = 2
					case TOKEN_THEN:
						if(stage!=2){
							displayError(step.Value.(Step),scenario.Value.(Scenario).Text,"Then keyword can only be used after Then statements.")
						}
						stage = 3
				}
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

func (feature *Feature) Validate(config config.Config) {

	fmt.Println()
	util.Print_subtitle(feature.Text)
	fmt.Println()
	feature.validate_max_steps(config.Errors.MaxStepsPerScenario)
	feature.validate_steps_length(config.Errors.MaxLenStep)
	if (!config.Warnings.AllowedEmptyLinesBetweenSteps) {
		feature.validate_lines_between_steps()
	}
	if (!config.Errors.EmptyFeature) {
		feature.validate_empty_feature()
	}
	if (!config.Errors.EmptyScenario) {
		feature.validate_empty_scenarios()
	}
	if (!config.Errors.Strict) {
		feature.validate_strict()

	}
}
