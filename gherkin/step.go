package gherkin

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Feature struct {
	GherkinToken
	Scenarios	*list.List
}

func (gherkinToken *GherkinToken) ToFeature() Feature {
	return Feature{*gherkinToken,list.New()}
}

func (this *Feature) ToString() string {
	var s string
	if this.LineNumber == 0 {
		fmt.Println("error")
		os.Exit(0)
	}
	s = ": " + this.Text + " at LINE:" + strconv.Itoa(this.LineNumber)
	for e := this.Scenarios.Front(); e != nil; e = e.Next() {
		s += "\n" + e.Value.(Scenario).Tag + " \t" + e.Value.(Scenario).Text
		for f := e.Value.(Scenario).Steps.Front(); f != nil; f = f.Next() {
			s += "\n   " + f.Value.(Step).Tag + " \t\t" + f.Value.(Step).Text

		}

	}
	return tMap[this.Kind] + s
}

type Scenario struct {
	GherkinToken
	Steps	*list.List
}

func (gherkinToken *GherkinToken) ToScenario() Scenario {
	return Scenario{*gherkinToken,list.New()}
}

type StepType int

type Step struct {
	GherkinToken
}

func (gherkinToken *GherkinToken) ToStep() Step {
	return Step{*gherkinToken}
}

