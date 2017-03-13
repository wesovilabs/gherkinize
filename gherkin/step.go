package gherkin

import (
	"container/list"
)

type Feature struct {
	GherkinToken
	Scenarios	*list.List
	Successful	bool
}

func (gherkinToken *GherkinToken) ToFeature() Feature {
	return Feature{*gherkinToken,list.New(), true}
}



type Scenario struct {
	GherkinToken
	Steps	*list.List
}

func (gherkinToken *GherkinToken) ToScenario() Scenario {
	return Scenario{*gherkinToken,list.New()}
}

type Step struct {
	GherkinToken
}

func (gherkinToken *GherkinToken) ToStep() Step {
	return Step{*gherkinToken}
}

