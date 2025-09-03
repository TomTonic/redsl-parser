package test

import (
	"os"
	redsl_parser "redsl_parser/generated"
	"testing"

	"github.com/antlr4-go/antlr/v4" // Adjust import path if needed
)

type SyntaxErrorListener struct {
	antlr.DefaultErrorListener
	ErrorCount int
}

func (l *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	l.ErrorCount++
}

func TestValidReDSLDocument(t *testing.T) {
	input, err := os.ReadFile("data/t1.redsl")
	if err != nil {
		t.Fatal(err)
	}
	inputStream := antlr.NewInputStream(string(input))
	lexer := redsl_parser.NewReDSLLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := redsl_parser.NewReDSLParser(stream)

	listener := &SyntaxErrorListener{}
	//parser.RemoveErrorListeners() // optional: remove default console error listener
	parser.AddErrorListener(listener)

	// Try to parse the document
	tree := parser.Parse()
	if listener.ErrorCount != 0 {
		t.Errorf("expected valid ReDSL document, got %d syntax errors", listener.ErrorCount)
	}
	if tree == nil {
		t.Errorf("parse tree is nil")
	}
}
