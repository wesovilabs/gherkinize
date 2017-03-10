package gherkin

import (
	"bufio"
	"bytes"
	"strings"
	"strconv"
	"io"
)

const (
	MaxScanTokenSize = 64 * 1024
	startBufSize = 4096 // Size of initial allocation for buffer.
)

type GherkinScanner struct {
	token		GherkinToken
	r 		*bufio.Reader
	Split 		bufio.SplitFunc
}




func NewGherkinScanner(r io.Reader) *GherkinScanner {

	return &GherkinScanner{
		r: bufio.NewReader(r),
	}
}

func  GherkinSplitFun(data []byte, atEOF bool)  (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanLines(data, atEOF)
	if err == nil && token != nil {
		_, err = strconv.ParseInt(string(token), 10, 32)
	}
	return
}


func (s *GherkinScanner) Scan(lineNumber int) (gherkinToken GherkinToken) {

	character := s.read()
	if isWhiteSpace(character) {
		return *newGherkinToken(AVOID,lineNumber)
	}
	if isLetter(character) {
		s.unread()
		gherkinToken =  s.scanIdent(lineNumber)
		gherkinToken.Text=s.scanGherkinTokenText()
		return gherkinToken
	}
	switch character {
		case '\n':
			return *newGherkinToken(NEW_LINE,lineNumber)
		case eof:
			return *newGherkinToken(EOF,lineNumber)
	}
	return *newGherkinToken(ILLEGAL,lineNumber)
}

func (s *GherkinScanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *GherkinScanner) scanGherkinTokenText() string {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); isEndOfLine(ch) {
			return buf.String()
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *GherkinScanner) scanIdent(lineNumber int) GherkinToken {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			return *newGherkinToken(EOF,lineNumber)
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' && ch != ':' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	switch strings.ToUpper(buf.String()) {
		case "FEATURE:":
			return *newGherkinToken(TOKEN_FEATURE,lineNumber)
		case "SCENARIO:":
			return *newGherkinToken(TOKEN_SCENARIO,lineNumber)
		case "GIVEN":
			return *newGherkinToken(TOKEN_GIVEN,lineNumber)
		case "WHEN":
			return *newGherkinToken(TOKEN_WHEN,lineNumber)
		case "THEN":
			return *newGherkinToken(TOKEN_THEN,lineNumber)
		case "AND":
			return *newGherkinToken(TOKEN_AND,lineNumber)
	}
	// Otherwise return as a regular identifier.
	return *newGherkinToken(ILLEGAL,lineNumber)
}

func (s *GherkinScanner) unread() { _ = s.r.UnreadRune() }

func isEndOfLine(ch rune) bool { return ch == '\n' }
// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhiteSpace(ch rune) bool { return ch == ' ' }
// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }
// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

var eof = rune(0)
