package testcode

import (
	"os"
	redsl_parser "redsl_parser"
	"testing"

	"github.com/antlr4-go/antlr/v4" // Adjust import path if needed
)

type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorCount int
}

func (l *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any,
	line, column int, msg string, e antlr.RecognitionException) {
	l.ErrorCount++
}

func TestParseReDSLFilesAndCompareDOM(t *testing.T) {
	files, err := os.ReadDir("../testdata")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && len(name) > 6 && name[len(name)-6:] == ".redsl" {
			t.Run(name, func(t *testing.T) {
				input, err := os.ReadFile("../testdata/" + name)
				if err != nil {
					t.Fatalf("failed to read file: %v", err)
				}
				inputStream := antlr.NewInputStream(string(input))
				lexer := redsl_parser.NewReDSLLexer(inputStream)
				stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
				parser := redsl_parser.NewReDSLParser(stream)

				listener := &SyntaxErrorListener{}
				lexer.RemoveErrorListeners()
				lexer.AddErrorListener(listener)
				parser.RemoveErrorListeners()
				parser.AddErrorListener(listener)
				parser.SetErrorHandler(antlr.NewDefaultErrorStrategy())

				tree := parser.Parse()
				if tree == nil {
					t.Errorf("parse tree is nil for %s", name)
					t.Fail()
					return
				}

				if !shouldBeValid(name) {
					if listener.ErrorCount == 0 {
						dom := buildDOM(tree, 0)
						if dom == nil {
							t.Errorf("expected syntax errors for %s, got none. DOM is nil.", name)
							t.Fail()
							return
						}
						domStr := dom.PrintTree()
						t.Errorf("expected syntax errors for %s, got none. The DOM is:\n%s", name, domStr)
					}
				} else {
					if listener.ErrorCount != 0 {
						t.Errorf("expected no syntax errors for %s, got %d", name, listener.ErrorCount)
						t.Fail()
						return
					}
					dom := buildDOM(tree, 0)
					if dom == nil {
						t.Errorf("DOM is nil for %s", name)
						t.Fail()
						return
					}
					domStr := dom.PrintTree()
					domFile := "../testdata/" + name[:len(name)-6] + ".dom"
					expectedDOM, err := os.ReadFile(domFile)
					if err != nil {
						t.Errorf("failed to read DOM file: %v, should be:\n%s", err, domStr)
						t.Fail()
						return
					}
					if domStr != string(expectedDOM) {
						t.Errorf("DOM string does not match expected for %s:\nexpected: %q\ngot:      %q", name, expectedDOM, domStr)
						t.Fail()
						return
					}
				}
			})
		}
	}
}

func shouldBeValid(filename string) bool {
	expected_suffix := "_valid.redsl"
	actual_suffix := filename[len(filename)-len(expected_suffix):]
	return actual_suffix == expected_suffix
}
