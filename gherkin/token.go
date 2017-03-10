package gherkin

import (
	"fmt"
	"os"
	"strconv"
)

type GherkinToken struct {
	Kind 		GherkinTokenType
	Tag		string
	Text 		string
	LineNumber	int
}

func newGherkinToken(kind GherkinTokenType, lineNumber int) *GherkinToken {
	return &GherkinToken{kind, tMap[kind], "", lineNumber}
}

var tokenMap map[string]GherkinTokenType
var tMap map[GherkinTokenType]string

func InitializeTokenMap() {
	tokenMap = make(map[string]GherkinTokenType)
	tokenMap["Feature:"] = TOKEN_FEATURE
	tokenMap["Scenario:"] = TOKEN_SCENARIO
	tokenMap["Given"] = TOKEN_GIVEN
	tokenMap["When"] = TOKEN_WHEN
	tokenMap["Then"] = TOKEN_THEN
	tokenMap["And"] = TOKEN_AND
	tokenMap["Eof"] = EOF
	tokenMap["\n"] = NEW_LINE
	tokenMap[" "] = AVOID

	tMap = make(map[GherkinTokenType]string)
	tMap[TOKEN_FEATURE] = "TOKEN_FEATURE"
	tMap[TOKEN_SCENARIO] = "TOKEN_SCENARIO"
	tMap[TOKEN_GIVEN] = "TOKEN_GIVEN"
	tMap[TOKEN_WHEN] = "TOKEN_WHEN"
	tMap[TOKEN_THEN] = "TOKEN_THEN"
	tMap[TOKEN_AND] = "TOKEN_AND"
	tMap[EOF] = "TOKEN_EOF"
	tMap[NEW_LINE] = "NEW_LINE"
	tMap[AVOID] = "AVOID"

}

type GherkinTokenType int

const (
	// Special tokens
	ILLEGAL GherkinTokenType = iota
	EOF
	keyword_begin
	TOKEN_FEATURE
	TOKEN_SCENARIO
	TOKEN_GIVEN
	TOKEN_WHEN
	TOKEN_THEN
	TOKEN_AND
	keyword_end
	NEW_LINE
	AVOID
)

func (this *GherkinToken) ToString() string {
	var s string
	if this.LineNumber == 0 {
		fmt.Println("error")
		os.Exit(0)
	}

	s = ": " + this.Text + " at LINE:" + strconv.Itoa(this.LineNumber)
	return tMap[this.Kind] + s
}

func (token GherkinToken) IsKeyword() bool {
	return token.Kind < keyword_end  && token.Kind > keyword_begin
}



