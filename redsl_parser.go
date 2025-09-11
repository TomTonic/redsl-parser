// Code generated from ReDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package redsl_parser // ReDSLParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ReDSLParser struct {
	*antlr.BaseParser
}

var ReDSLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func redslparserParserInit() {
	staticData := &ReDSLParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'file'", "", "", "'package'", "", "", "", "'document'", "'version-info'",
		"'glossary'", "'local'", "'global'", "'['", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "']'", "", "'='", "", "",
		"','", "", "", "", "", "';'", "", "", "", "", "", "", "", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "FILE_KEYWORD", "BLOCK_START", "BLOCK_END", "PACKAGE_KEYWORD", "ID_STR",
		"WS", "REQ_DEF", "DOCUMENT_KEYWORD", "VERSION_INFO_KEYWORD", "GLOSSARY_KEYWORD",
		"LOCAL_KEYWORD", "GLOBAL_KEYWORD", "PARAM_START", "TEXT_START", "DEDUCT_START",
		"REQ_DEF_WS", "TEXT_CLOSE", "TEXT_COMMENT", "TEXT_WS", "TEXT_LINE_BREAK",
		"TEXT_NEXT_PARA", "TEXT_EXAMPLE_MARKER", "TEXT_RATIO_MARKER", "TEXT_REF_MARKER",
		"TEXT_RE_ID_REF", "TEXT_URI", "TEXT_TERM_REF", "TEXT_START_MATH", "TEXT_ESC_SEQ",
		"TEXT_CONTENT", "PARAM_CLOSE", "PARAM_ID", "PARAM_EQUALS", "PARAM_STRING",
		"PARAM_ID_LIST_START", "PARAM_ID_LIST_SEP", "PARAM_ID_LIST_END", "PARAM_BOOL",
		"PARAM_NUMBER", "PARAM_WS", "DEDUCT_CLOSE", "DEDUCT_COMMENT", "DEDUCT_RE_ID_REF",
		"DEDUCT_AND", "DEDUCT_OR", "DEDUCT_NOT", "DEDUCT_TRUE", "DEDUCT_FALSE",
		"DEDUCT_LPAREN", "DEDUCT_RPAREN", "DEDUCT_WS", "MATH_CLOSE", "MATH_ESC_SEQ",
		"MATH_CONTENT",
	}
	staticData.RuleNames = []string{
		"parse", "fileDecl", "packageDecl", "requirementDecl", "requirementDeduction",
		"logicalExpression", "logicalAtom", "logicalUnaryExpression", "logicalBinaryExpression",
		"exampleDecl", "rationaleDecl", "referenceDecl", "runningText", "parameterBlock",
		"paramDecl", "stringParamDecl", "idParamDecl", "numericParamDecl", "boolParamDecl",
		"idListParamDecl", "mathModeExpression", "documentDecl", "packageRef",
		"fileRef",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 237, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 4, 0, 50, 8, 0, 11, 0, 12,
		0, 51, 1, 0, 1, 0, 4, 0, 56, 8, 0, 11, 0, 12, 0, 57, 3, 0, 60, 8, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 69, 8, 1, 10, 1, 12, 1, 72,
		9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 81, 8, 3, 1, 3, 1,
		3, 3, 3, 85, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 93, 8, 3,
		10, 3, 12, 3, 96, 9, 3, 3, 3, 98, 8, 3, 1, 3, 3, 3, 101, 8, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 113, 8, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 123, 8, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 4, 12,
		147, 8, 12, 11, 12, 12, 12, 148, 1, 13, 1, 13, 5, 13, 153, 8, 13, 10, 13,
		12, 13, 156, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3,
		14, 165, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 3, 18, 182, 8, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 190, 8, 19, 10, 19, 12, 19,
		193, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 5, 20, 199, 8, 20, 10, 20, 12,
		20, 202, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 3, 21, 213, 8, 21, 5, 21, 215, 8, 21, 10, 21, 12, 21, 218, 9, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 226, 8, 22, 10, 22, 12,
		22, 229, 9, 22, 1, 22, 3, 22, 232, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 0,
		0, 24, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 0, 4, 2, 0, 43, 43, 47, 48, 1, 0, 44, 45, 1, 0,
		53, 54, 1, 0, 11, 12, 255, 0, 59, 1, 0, 0, 0, 2, 61, 1, 0, 0, 0, 4, 75,
		1, 0, 0, 0, 6, 78, 1, 0, 0, 0, 8, 104, 1, 0, 0, 0, 10, 112, 1, 0, 0, 0,
		12, 114, 1, 0, 0, 0, 14, 122, 1, 0, 0, 0, 16, 124, 1, 0, 0, 0, 18, 128,
		1, 0, 0, 0, 20, 131, 1, 0, 0, 0, 22, 134, 1, 0, 0, 0, 24, 146, 1, 0, 0,
		0, 26, 150, 1, 0, 0, 0, 28, 164, 1, 0, 0, 0, 30, 166, 1, 0, 0, 0, 32, 170,
		1, 0, 0, 0, 34, 174, 1, 0, 0, 0, 36, 178, 1, 0, 0, 0, 38, 183, 1, 0, 0,
		0, 40, 196, 1, 0, 0, 0, 42, 205, 1, 0, 0, 0, 44, 221, 1, 0, 0, 0, 46, 233,
		1, 0, 0, 0, 48, 50, 3, 2, 1, 0, 49, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0,
		51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 60, 1, 0, 0, 0, 53, 56, 3,
		4, 2, 0, 54, 56, 3, 6, 3, 0, 55, 53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56,
		57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0,
		0, 59, 49, 1, 0, 0, 0, 59, 55, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 1, 1,
		0, 0, 0, 61, 62, 5, 1, 0, 0, 62, 63, 5, 5, 0, 0, 63, 70, 5, 2, 0, 0, 64,
		69, 3, 4, 2, 0, 65, 69, 3, 6, 3, 0, 66, 69, 3, 8, 4, 0, 67, 69, 3, 42,
		21, 0, 68, 64, 1, 0, 0, 0, 68, 65, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68,
		67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0,
		0, 71, 73, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74, 5, 3, 0, 0, 74, 3, 1,
		0, 0, 0, 75, 76, 5, 4, 0, 0, 76, 77, 5, 5, 0, 0, 77, 5, 1, 0, 0, 0, 78,
		80, 5, 7, 0, 0, 79, 81, 3, 26, 13, 0, 80, 79, 1, 0, 0, 0, 80, 81, 1, 0,
		0, 0, 81, 82, 1, 0, 0, 0, 82, 84, 5, 14, 0, 0, 83, 85, 5, 21, 0, 0, 84,
		83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 97, 1, 0, 0, 0, 86, 94, 3, 24,
		12, 0, 87, 93, 3, 18, 9, 0, 88, 93, 3, 20, 10, 0, 89, 93, 3, 22, 11, 0,
		90, 91, 5, 21, 0, 0, 91, 93, 3, 24, 12, 0, 92, 87, 1, 0, 0, 0, 92, 88,
		1, 0, 0, 0, 92, 89, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1,
		0, 0, 0, 97, 86, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 100, 1, 0, 0, 0, 99,
		101, 5, 21, 0, 0, 100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102,
		1, 0, 0, 0, 102, 103, 5, 17, 0, 0, 103, 7, 1, 0, 0, 0, 104, 105, 5, 7,
		0, 0, 105, 106, 5, 15, 0, 0, 106, 107, 3, 10, 5, 0, 107, 108, 5, 41, 0,
		0, 108, 9, 1, 0, 0, 0, 109, 113, 3, 12, 6, 0, 110, 113, 3, 14, 7, 0, 111,
		113, 3, 16, 8, 0, 112, 109, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 111,
		1, 0, 0, 0, 113, 11, 1, 0, 0, 0, 114, 115, 7, 0, 0, 0, 115, 13, 1, 0, 0,
		0, 116, 117, 5, 46, 0, 0, 117, 123, 3, 10, 5, 0, 118, 119, 5, 49, 0, 0,
		119, 120, 3, 10, 5, 0, 120, 121, 5, 50, 0, 0, 121, 123, 1, 0, 0, 0, 122,
		116, 1, 0, 0, 0, 122, 118, 1, 0, 0, 0, 123, 15, 1, 0, 0, 0, 124, 125, 5,
		43, 0, 0, 125, 126, 7, 1, 0, 0, 126, 127, 3, 10, 5, 0, 127, 17, 1, 0, 0,
		0, 128, 129, 5, 22, 0, 0, 129, 130, 3, 24, 12, 0, 130, 19, 1, 0, 0, 0,
		131, 132, 5, 23, 0, 0, 132, 133, 3, 24, 12, 0, 133, 21, 1, 0, 0, 0, 134,
		135, 5, 24, 0, 0, 135, 136, 3, 24, 12, 0, 136, 23, 1, 0, 0, 0, 137, 147,
		5, 30, 0, 0, 138, 147, 5, 29, 0, 0, 139, 147, 5, 25, 0, 0, 140, 147, 5,
		27, 0, 0, 141, 147, 5, 18, 0, 0, 142, 147, 5, 26, 0, 0, 143, 147, 5, 19,
		0, 0, 144, 147, 5, 20, 0, 0, 145, 147, 3, 40, 20, 0, 146, 137, 1, 0, 0,
		0, 146, 138, 1, 0, 0, 0, 146, 139, 1, 0, 0, 0, 146, 140, 1, 0, 0, 0, 146,
		141, 1, 0, 0, 0, 146, 142, 1, 0, 0, 0, 146, 143, 1, 0, 0, 0, 146, 144,
		1, 0, 0, 0, 146, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 146, 1, 0,
		0, 0, 148, 149, 1, 0, 0, 0, 149, 25, 1, 0, 0, 0, 150, 154, 5, 13, 0, 0,
		151, 153, 3, 28, 14, 0, 152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154,
		152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 154,
		1, 0, 0, 0, 157, 158, 5, 31, 0, 0, 158, 27, 1, 0, 0, 0, 159, 165, 3, 30,
		15, 0, 160, 165, 3, 32, 16, 0, 161, 165, 3, 34, 17, 0, 162, 165, 3, 36,
		18, 0, 163, 165, 3, 38, 19, 0, 164, 159, 1, 0, 0, 0, 164, 160, 1, 0, 0,
		0, 164, 161, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 163, 1, 0, 0, 0, 165,
		29, 1, 0, 0, 0, 166, 167, 5, 32, 0, 0, 167, 168, 5, 33, 0, 0, 168, 169,
		5, 34, 0, 0, 169, 31, 1, 0, 0, 0, 170, 171, 5, 32, 0, 0, 171, 172, 5, 33,
		0, 0, 172, 173, 5, 32, 0, 0, 173, 33, 1, 0, 0, 0, 174, 175, 5, 32, 0, 0,
		175, 176, 5, 33, 0, 0, 176, 177, 5, 39, 0, 0, 177, 35, 1, 0, 0, 0, 178,
		181, 5, 32, 0, 0, 179, 180, 5, 33, 0, 0, 180, 182, 5, 38, 0, 0, 181, 179,
		1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 37, 1, 0, 0, 0, 183, 184, 5, 32,
		0, 0, 184, 185, 5, 33, 0, 0, 185, 186, 5, 35, 0, 0, 186, 191, 5, 32, 0,
		0, 187, 188, 5, 36, 0, 0, 188, 190, 5, 32, 0, 0, 189, 187, 1, 0, 0, 0,
		190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192,
		194, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 195, 5, 37, 0, 0, 195, 39,
		1, 0, 0, 0, 196, 200, 5, 28, 0, 0, 197, 199, 7, 2, 0, 0, 198, 197, 1, 0,
		0, 0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 203, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 204, 5, 52, 0, 0, 204,
		41, 1, 0, 0, 0, 205, 206, 5, 8, 0, 0, 206, 207, 5, 5, 0, 0, 207, 216, 5,
		2, 0, 0, 208, 215, 3, 44, 22, 0, 209, 215, 5, 9, 0, 0, 210, 212, 5, 10,
		0, 0, 211, 213, 7, 3, 0, 0, 212, 211, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0,
		213, 215, 1, 0, 0, 0, 214, 208, 1, 0, 0, 0, 214, 209, 1, 0, 0, 0, 214,
		210, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217,
		1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 220, 5, 3,
		0, 0, 220, 43, 1, 0, 0, 0, 221, 222, 5, 4, 0, 0, 222, 231, 5, 5, 0, 0,
		223, 227, 5, 2, 0, 0, 224, 226, 3, 46, 23, 0, 225, 224, 1, 0, 0, 0, 226,
		229, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 230,
		1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 230, 232, 5, 3, 0, 0, 231, 223, 1, 0,
		0, 0, 231, 232, 1, 0, 0, 0, 232, 45, 1, 0, 0, 0, 233, 234, 5, 1, 0, 0,
		234, 235, 5, 5, 0, 0, 235, 47, 1, 0, 0, 0, 26, 51, 55, 57, 59, 68, 70,
		80, 84, 92, 94, 97, 100, 112, 122, 146, 148, 154, 164, 181, 191, 200, 212,
		214, 216, 227, 231,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ReDSLParserInit initializes any static state used to implement ReDSLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewReDSLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ReDSLParserInit() {
	staticData := &ReDSLParserParserStaticData
	staticData.once.Do(redslparserParserInit)
}

// NewReDSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewReDSLParser(input antlr.TokenStream) *ReDSLParser {
	ReDSLParserInit()
	this := new(ReDSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ReDSLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ReDSLParser.g4"

	return this
}

// ReDSLParser tokens.
const (
	ReDSLParserEOF                  = antlr.TokenEOF
	ReDSLParserFILE_KEYWORD         = 1
	ReDSLParserBLOCK_START          = 2
	ReDSLParserBLOCK_END            = 3
	ReDSLParserPACKAGE_KEYWORD      = 4
	ReDSLParserID_STR               = 5
	ReDSLParserWS                   = 6
	ReDSLParserREQ_DEF              = 7
	ReDSLParserDOCUMENT_KEYWORD     = 8
	ReDSLParserVERSION_INFO_KEYWORD = 9
	ReDSLParserGLOSSARY_KEYWORD     = 10
	ReDSLParserLOCAL_KEYWORD        = 11
	ReDSLParserGLOBAL_KEYWORD       = 12
	ReDSLParserPARAM_START          = 13
	ReDSLParserTEXT_START           = 14
	ReDSLParserDEDUCT_START         = 15
	ReDSLParserREQ_DEF_WS           = 16
	ReDSLParserTEXT_CLOSE           = 17
	ReDSLParserTEXT_COMMENT         = 18
	ReDSLParserTEXT_WS              = 19
	ReDSLParserTEXT_LINE_BREAK      = 20
	ReDSLParserTEXT_NEXT_PARA       = 21
	ReDSLParserTEXT_EXAMPLE_MARKER  = 22
	ReDSLParserTEXT_RATIO_MARKER    = 23
	ReDSLParserTEXT_REF_MARKER      = 24
	ReDSLParserTEXT_RE_ID_REF       = 25
	ReDSLParserTEXT_URI             = 26
	ReDSLParserTEXT_TERM_REF        = 27
	ReDSLParserTEXT_START_MATH      = 28
	ReDSLParserTEXT_ESC_SEQ         = 29
	ReDSLParserTEXT_CONTENT         = 30
	ReDSLParserPARAM_CLOSE          = 31
	ReDSLParserPARAM_ID             = 32
	ReDSLParserPARAM_EQUALS         = 33
	ReDSLParserPARAM_STRING         = 34
	ReDSLParserPARAM_ID_LIST_START  = 35
	ReDSLParserPARAM_ID_LIST_SEP    = 36
	ReDSLParserPARAM_ID_LIST_END    = 37
	ReDSLParserPARAM_BOOL           = 38
	ReDSLParserPARAM_NUMBER         = 39
	ReDSLParserPARAM_WS             = 40
	ReDSLParserDEDUCT_CLOSE         = 41
	ReDSLParserDEDUCT_COMMENT       = 42
	ReDSLParserDEDUCT_RE_ID_REF     = 43
	ReDSLParserDEDUCT_AND           = 44
	ReDSLParserDEDUCT_OR            = 45
	ReDSLParserDEDUCT_NOT           = 46
	ReDSLParserDEDUCT_TRUE          = 47
	ReDSLParserDEDUCT_FALSE         = 48
	ReDSLParserDEDUCT_LPAREN        = 49
	ReDSLParserDEDUCT_RPAREN        = 50
	ReDSLParserDEDUCT_WS            = 51
	ReDSLParserMATH_CLOSE           = 52
	ReDSLParserMATH_ESC_SEQ         = 53
	ReDSLParserMATH_CONTENT         = 54
)

// ReDSLParser rules.
const (
	ReDSLParserRULE_parse                   = 0
	ReDSLParserRULE_fileDecl                = 1
	ReDSLParserRULE_packageDecl             = 2
	ReDSLParserRULE_requirementDecl         = 3
	ReDSLParserRULE_requirementDeduction    = 4
	ReDSLParserRULE_logicalExpression       = 5
	ReDSLParserRULE_logicalAtom             = 6
	ReDSLParserRULE_logicalUnaryExpression  = 7
	ReDSLParserRULE_logicalBinaryExpression = 8
	ReDSLParserRULE_exampleDecl             = 9
	ReDSLParserRULE_rationaleDecl           = 10
	ReDSLParserRULE_referenceDecl           = 11
	ReDSLParserRULE_runningText             = 12
	ReDSLParserRULE_parameterBlock          = 13
	ReDSLParserRULE_paramDecl               = 14
	ReDSLParserRULE_stringParamDecl         = 15
	ReDSLParserRULE_idParamDecl             = 16
	ReDSLParserRULE_numericParamDecl        = 17
	ReDSLParserRULE_boolParamDecl           = 18
	ReDSLParserRULE_idListParamDecl         = 19
	ReDSLParserRULE_mathModeExpression      = 20
	ReDSLParserRULE_documentDecl            = 21
	ReDSLParserRULE_packageRef              = 22
	ReDSLParserRULE_fileRef                 = 23
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFileDecl() []IFileDeclContext
	FileDecl(i int) IFileDeclContext
	AllPackageDecl() []IPackageDeclContext
	PackageDecl(i int) IPackageDeclContext
	AllRequirementDecl() []IRequirementDeclContext
	RequirementDecl(i int) IRequirementDeclContext

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) AllFileDecl() []IFileDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFileDeclContext); ok {
			len++
		}
	}

	tst := make([]IFileDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFileDeclContext); ok {
			tst[i] = t.(IFileDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) FileDecl(i int) IFileDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFileDeclContext)
}

func (s *ParseContext) AllPackageDecl() []IPackageDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPackageDeclContext); ok {
			len++
		}
	}

	tst := make([]IPackageDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPackageDeclContext); ok {
			tst[i] = t.(IPackageDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) PackageDecl(i int) IPackageDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageDeclContext)
}

func (s *ParseContext) AllRequirementDecl() []IRequirementDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRequirementDeclContext); ok {
			len++
		}
	}

	tst := make([]IRequirementDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRequirementDeclContext); ok {
			tst[i] = t.(IRequirementDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) RequirementDecl(i int) IRequirementDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequirementDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequirementDeclContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *ReDSLParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ReDSLParserRULE_parse)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case ReDSLParserFILE_KEYWORD:
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ReDSLParserFILE_KEYWORD {
			{
				p.SetState(48)
				p.FileDecl()
			}

			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ReDSLParserPACKAGE_KEYWORD, ReDSLParserREQ_DEF:
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ReDSLParserPACKAGE_KEYWORD || _la == ReDSLParserREQ_DEF {
			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case ReDSLParserPACKAGE_KEYWORD:
				{
					p.SetState(53)
					p.PackageDecl()
				}

			case ReDSLParserREQ_DEF:
				{
					p.SetState(54)
					p.RequirementDecl()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ReDSLParserEOF:

	default:
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFileDeclContext is an interface to support dynamic dispatch.
type IFileDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILE_KEYWORD() antlr.TerminalNode
	ID_STR() antlr.TerminalNode
	BLOCK_START() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
	AllPackageDecl() []IPackageDeclContext
	PackageDecl(i int) IPackageDeclContext
	AllRequirementDecl() []IRequirementDeclContext
	RequirementDecl(i int) IRequirementDeclContext
	AllRequirementDeduction() []IRequirementDeductionContext
	RequirementDeduction(i int) IRequirementDeductionContext
	AllDocumentDecl() []IDocumentDeclContext
	DocumentDecl(i int) IDocumentDeclContext

	// IsFileDeclContext differentiates from other interfaces.
	IsFileDeclContext()
}

type FileDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileDeclContext() *FileDeclContext {
	var p = new(FileDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_fileDecl
	return p
}

func InitEmptyFileDeclContext(p *FileDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_fileDecl
}

func (*FileDeclContext) IsFileDeclContext() {}

func NewFileDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileDeclContext {
	var p = new(FileDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_fileDecl

	return p
}

func (s *FileDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FileDeclContext) FILE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(ReDSLParserFILE_KEYWORD, 0)
}

func (s *FileDeclContext) ID_STR() antlr.TerminalNode {
	return s.GetToken(ReDSLParserID_STR, 0)
}

func (s *FileDeclContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(ReDSLParserBLOCK_START, 0)
}

func (s *FileDeclContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(ReDSLParserBLOCK_END, 0)
}

func (s *FileDeclContext) AllPackageDecl() []IPackageDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPackageDeclContext); ok {
			len++
		}
	}

	tst := make([]IPackageDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPackageDeclContext); ok {
			tst[i] = t.(IPackageDeclContext)
			i++
		}
	}

	return tst
}

func (s *FileDeclContext) PackageDecl(i int) IPackageDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageDeclContext)
}

func (s *FileDeclContext) AllRequirementDecl() []IRequirementDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRequirementDeclContext); ok {
			len++
		}
	}

	tst := make([]IRequirementDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRequirementDeclContext); ok {
			tst[i] = t.(IRequirementDeclContext)
			i++
		}
	}

	return tst
}

func (s *FileDeclContext) RequirementDecl(i int) IRequirementDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequirementDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequirementDeclContext)
}

func (s *FileDeclContext) AllRequirementDeduction() []IRequirementDeductionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRequirementDeductionContext); ok {
			len++
		}
	}

	tst := make([]IRequirementDeductionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRequirementDeductionContext); ok {
			tst[i] = t.(IRequirementDeductionContext)
			i++
		}
	}

	return tst
}

func (s *FileDeclContext) RequirementDeduction(i int) IRequirementDeductionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequirementDeductionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequirementDeductionContext)
}

func (s *FileDeclContext) AllDocumentDecl() []IDocumentDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDocumentDeclContext); ok {
			len++
		}
	}

	tst := make([]IDocumentDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDocumentDeclContext); ok {
			tst[i] = t.(IDocumentDeclContext)
			i++
		}
	}

	return tst
}

func (s *FileDeclContext) DocumentDecl(i int) IDocumentDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocumentDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocumentDeclContext)
}

func (s *FileDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterFileDecl(s)
	}
}

func (s *FileDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitFileDecl(s)
	}
}

func (p *ReDSLParser) FileDecl() (localctx IFileDeclContext) {
	localctx = NewFileDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ReDSLParserRULE_fileDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(ReDSLParserFILE_KEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Match(ReDSLParserID_STR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(ReDSLParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&400) != 0 {
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(64)
				p.PackageDecl()
			}

		case 2:
			{
				p.SetState(65)
				p.RequirementDecl()
			}

		case 3:
			{
				p.SetState(66)
				p.RequirementDeduction()
			}

		case 4:
			{
				p.SetState(67)
				p.DocumentDecl()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(73)
		p.Match(ReDSLParserBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageDeclContext is an interface to support dynamic dispatch.
type IPackageDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE_KEYWORD() antlr.TerminalNode
	ID_STR() antlr.TerminalNode

	// IsPackageDeclContext differentiates from other interfaces.
	IsPackageDeclContext()
}

type PackageDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageDeclContext() *PackageDeclContext {
	var p = new(PackageDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_packageDecl
	return p
}

func InitEmptyPackageDeclContext(p *PackageDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_packageDecl
}

func (*PackageDeclContext) IsPackageDeclContext() {}

func NewPackageDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageDeclContext {
	var p = new(PackageDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_packageDecl

	return p
}

func (s *PackageDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageDeclContext) PACKAGE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPACKAGE_KEYWORD, 0)
}

func (s *PackageDeclContext) ID_STR() antlr.TerminalNode {
	return s.GetToken(ReDSLParserID_STR, 0)
}

func (s *PackageDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterPackageDecl(s)
	}
}

func (s *PackageDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitPackageDecl(s)
	}
}

func (p *ReDSLParser) PackageDecl() (localctx IPackageDeclContext) {
	localctx = NewPackageDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ReDSLParserRULE_packageDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(ReDSLParserPACKAGE_KEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(ReDSLParserID_STR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRequirementDeclContext is an interface to support dynamic dispatch.
type IRequirementDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REQ_DEF() antlr.TerminalNode
	TEXT_START() antlr.TerminalNode
	TEXT_CLOSE() antlr.TerminalNode
	ParameterBlock() IParameterBlockContext
	AllTEXT_NEXT_PARA() []antlr.TerminalNode
	TEXT_NEXT_PARA(i int) antlr.TerminalNode
	AllRunningText() []IRunningTextContext
	RunningText(i int) IRunningTextContext
	AllExampleDecl() []IExampleDeclContext
	ExampleDecl(i int) IExampleDeclContext
	AllRationaleDecl() []IRationaleDeclContext
	RationaleDecl(i int) IRationaleDeclContext
	AllReferenceDecl() []IReferenceDeclContext
	ReferenceDecl(i int) IReferenceDeclContext

	// IsRequirementDeclContext differentiates from other interfaces.
	IsRequirementDeclContext()
}

type RequirementDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequirementDeclContext() *RequirementDeclContext {
	var p = new(RequirementDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_requirementDecl
	return p
}

func InitEmptyRequirementDeclContext(p *RequirementDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_requirementDecl
}

func (*RequirementDeclContext) IsRequirementDeclContext() {}

func NewRequirementDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequirementDeclContext {
	var p = new(RequirementDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_requirementDecl

	return p
}

func (s *RequirementDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *RequirementDeclContext) REQ_DEF() antlr.TerminalNode {
	return s.GetToken(ReDSLParserREQ_DEF, 0)
}

func (s *RequirementDeclContext) TEXT_START() antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_START, 0)
}

func (s *RequirementDeclContext) TEXT_CLOSE() antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_CLOSE, 0)
}

func (s *RequirementDeclContext) ParameterBlock() IParameterBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterBlockContext)
}

func (s *RequirementDeclContext) AllTEXT_NEXT_PARA() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserTEXT_NEXT_PARA)
}

func (s *RequirementDeclContext) TEXT_NEXT_PARA(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_NEXT_PARA, i)
}

func (s *RequirementDeclContext) AllRunningText() []IRunningTextContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRunningTextContext); ok {
			len++
		}
	}

	tst := make([]IRunningTextContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRunningTextContext); ok {
			tst[i] = t.(IRunningTextContext)
			i++
		}
	}

	return tst
}

func (s *RequirementDeclContext) RunningText(i int) IRunningTextContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunningTextContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunningTextContext)
}

func (s *RequirementDeclContext) AllExampleDecl() []IExampleDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExampleDeclContext); ok {
			len++
		}
	}

	tst := make([]IExampleDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExampleDeclContext); ok {
			tst[i] = t.(IExampleDeclContext)
			i++
		}
	}

	return tst
}

func (s *RequirementDeclContext) ExampleDecl(i int) IExampleDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExampleDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExampleDeclContext)
}

func (s *RequirementDeclContext) AllRationaleDecl() []IRationaleDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRationaleDeclContext); ok {
			len++
		}
	}

	tst := make([]IRationaleDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRationaleDeclContext); ok {
			tst[i] = t.(IRationaleDeclContext)
			i++
		}
	}

	return tst
}

func (s *RequirementDeclContext) RationaleDecl(i int) IRationaleDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRationaleDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRationaleDeclContext)
}

func (s *RequirementDeclContext) AllReferenceDecl() []IReferenceDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReferenceDeclContext); ok {
			len++
		}
	}

	tst := make([]IReferenceDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReferenceDeclContext); ok {
			tst[i] = t.(IReferenceDeclContext)
			i++
		}
	}

	return tst
}

func (s *RequirementDeclContext) ReferenceDecl(i int) IReferenceDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceDeclContext)
}

func (s *RequirementDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequirementDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequirementDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterRequirementDecl(s)
	}
}

func (s *RequirementDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitRequirementDecl(s)
	}
}

func (p *ReDSLParser) RequirementDecl() (localctx IRequirementDeclContext) {
	localctx = NewRequirementDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ReDSLParserRULE_requirementDecl)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(ReDSLParserREQ_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ReDSLParserPARAM_START {
		{
			p.SetState(79)
			p.ParameterBlock()
		}

	}
	{
		p.SetState(82)
		p.Match(ReDSLParserTEXT_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(83)
			p.Match(ReDSLParserTEXT_NEXT_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2115764224) != 0 {
		{
			p.SetState(86)
			p.RunningText()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(92)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case ReDSLParserTEXT_EXAMPLE_MARKER:
					{
						p.SetState(87)
						p.ExampleDecl()
					}

				case ReDSLParserTEXT_RATIO_MARKER:
					{
						p.SetState(88)
						p.RationaleDecl()
					}

				case ReDSLParserTEXT_REF_MARKER:
					{
						p.SetState(89)
						p.ReferenceDecl()
					}

				case ReDSLParserTEXT_NEXT_PARA:
					{
						p.SetState(90)
						p.Match(ReDSLParserTEXT_NEXT_PARA)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(91)
						p.RunningText()
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

			}
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ReDSLParserTEXT_NEXT_PARA {
		{
			p.SetState(99)
			p.Match(ReDSLParserTEXT_NEXT_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(102)
		p.Match(ReDSLParserTEXT_CLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRequirementDeductionContext is an interface to support dynamic dispatch.
type IRequirementDeductionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REQ_DEF() antlr.TerminalNode
	DEDUCT_START() antlr.TerminalNode
	LogicalExpression() ILogicalExpressionContext
	DEDUCT_CLOSE() antlr.TerminalNode

	// IsRequirementDeductionContext differentiates from other interfaces.
	IsRequirementDeductionContext()
}

type RequirementDeductionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequirementDeductionContext() *RequirementDeductionContext {
	var p = new(RequirementDeductionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_requirementDeduction
	return p
}

func InitEmptyRequirementDeductionContext(p *RequirementDeductionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_requirementDeduction
}

func (*RequirementDeductionContext) IsRequirementDeductionContext() {}

func NewRequirementDeductionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequirementDeductionContext {
	var p = new(RequirementDeductionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_requirementDeduction

	return p
}

func (s *RequirementDeductionContext) GetParser() antlr.Parser { return s.parser }

func (s *RequirementDeductionContext) REQ_DEF() antlr.TerminalNode {
	return s.GetToken(ReDSLParserREQ_DEF, 0)
}

func (s *RequirementDeductionContext) DEDUCT_START() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_START, 0)
}

func (s *RequirementDeductionContext) LogicalExpression() ILogicalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalExpressionContext)
}

func (s *RequirementDeductionContext) DEDUCT_CLOSE() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_CLOSE, 0)
}

func (s *RequirementDeductionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequirementDeductionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequirementDeductionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterRequirementDeduction(s)
	}
}

func (s *RequirementDeductionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitRequirementDeduction(s)
	}
}

func (p *ReDSLParser) RequirementDeduction() (localctx IRequirementDeductionContext) {
	localctx = NewRequirementDeductionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ReDSLParserRULE_requirementDeduction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(ReDSLParserREQ_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(ReDSLParserDEDUCT_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.LogicalExpression()
	}
	{
		p.SetState(107)
		p.Match(ReDSLParserDEDUCT_CLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalExpressionContext is an interface to support dynamic dispatch.
type ILogicalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalAtom() ILogicalAtomContext
	LogicalUnaryExpression() ILogicalUnaryExpressionContext
	LogicalBinaryExpression() ILogicalBinaryExpressionContext

	// IsLogicalExpressionContext differentiates from other interfaces.
	IsLogicalExpressionContext()
}

type LogicalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalExpressionContext() *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_logicalExpression
	return p
}

func InitEmptyLogicalExpressionContext(p *LogicalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_logicalExpression
}

func (*LogicalExpressionContext) IsLogicalExpressionContext() {}

func NewLogicalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_logicalExpression

	return p
}

func (s *LogicalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalExpressionContext) LogicalAtom() ILogicalAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAtomContext)
}

func (s *LogicalExpressionContext) LogicalUnaryExpression() ILogicalUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalUnaryExpressionContext)
}

func (s *LogicalExpressionContext) LogicalBinaryExpression() ILogicalBinaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalBinaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalBinaryExpressionContext)
}

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

func (p *ReDSLParser) LogicalExpression() (localctx ILogicalExpressionContext) {
	localctx = NewLogicalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ReDSLParserRULE_logicalExpression)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.LogicalAtom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.LogicalUnaryExpression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.LogicalBinaryExpression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalAtomContext is an interface to support dynamic dispatch.
type ILogicalAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEDUCT_RE_ID_REF() antlr.TerminalNode
	DEDUCT_TRUE() antlr.TerminalNode
	DEDUCT_FALSE() antlr.TerminalNode

	// IsLogicalAtomContext differentiates from other interfaces.
	IsLogicalAtomContext()
}

type LogicalAtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAtomContext() *LogicalAtomContext {
	var p = new(LogicalAtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_logicalAtom
	return p
}

func InitEmptyLogicalAtomContext(p *LogicalAtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_logicalAtom
}

func (*LogicalAtomContext) IsLogicalAtomContext() {}

func NewLogicalAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAtomContext {
	var p = new(LogicalAtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_logicalAtom

	return p
}

func (s *LogicalAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAtomContext) DEDUCT_RE_ID_REF() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_RE_ID_REF, 0)
}

func (s *LogicalAtomContext) DEDUCT_TRUE() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_TRUE, 0)
}

func (s *LogicalAtomContext) DEDUCT_FALSE() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_FALSE, 0)
}

func (s *LogicalAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterLogicalAtom(s)
	}
}

func (s *LogicalAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitLogicalAtom(s)
	}
}

func (p *ReDSLParser) LogicalAtom() (localctx ILogicalAtomContext) {
	localctx = NewLogicalAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ReDSLParserRULE_logicalAtom)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&431008558088192) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalUnaryExpressionContext is an interface to support dynamic dispatch.
type ILogicalUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEDUCT_NOT() antlr.TerminalNode
	LogicalExpression() ILogicalExpressionContext
	DEDUCT_LPAREN() antlr.TerminalNode
	DEDUCT_RPAREN() antlr.TerminalNode

	// IsLogicalUnaryExpressionContext differentiates from other interfaces.
	IsLogicalUnaryExpressionContext()
}

type LogicalUnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalUnaryExpressionContext() *LogicalUnaryExpressionContext {
	var p = new(LogicalUnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_logicalUnaryExpression
	return p
}

func InitEmptyLogicalUnaryExpressionContext(p *LogicalUnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_logicalUnaryExpression
}

func (*LogicalUnaryExpressionContext) IsLogicalUnaryExpressionContext() {}

func NewLogicalUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalUnaryExpressionContext {
	var p = new(LogicalUnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_logicalUnaryExpression

	return p
}

func (s *LogicalUnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalUnaryExpressionContext) DEDUCT_NOT() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_NOT, 0)
}

func (s *LogicalUnaryExpressionContext) LogicalExpression() ILogicalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalExpressionContext)
}

func (s *LogicalUnaryExpressionContext) DEDUCT_LPAREN() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_LPAREN, 0)
}

func (s *LogicalUnaryExpressionContext) DEDUCT_RPAREN() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_RPAREN, 0)
}

func (s *LogicalUnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalUnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalUnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterLogicalUnaryExpression(s)
	}
}

func (s *LogicalUnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitLogicalUnaryExpression(s)
	}
}

func (p *ReDSLParser) LogicalUnaryExpression() (localctx ILogicalUnaryExpressionContext) {
	localctx = NewLogicalUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ReDSLParserRULE_logicalUnaryExpression)
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ReDSLParserDEDUCT_NOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(ReDSLParserDEDUCT_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.LogicalExpression()
		}

	case ReDSLParserDEDUCT_LPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(ReDSLParserDEDUCT_LPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.LogicalExpression()
		}
		{
			p.SetState(120)
			p.Match(ReDSLParserDEDUCT_RPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalBinaryExpressionContext is an interface to support dynamic dispatch.
type ILogicalBinaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEDUCT_RE_ID_REF() antlr.TerminalNode
	LogicalExpression() ILogicalExpressionContext
	DEDUCT_AND() antlr.TerminalNode
	DEDUCT_OR() antlr.TerminalNode

	// IsLogicalBinaryExpressionContext differentiates from other interfaces.
	IsLogicalBinaryExpressionContext()
}

type LogicalBinaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalBinaryExpressionContext() *LogicalBinaryExpressionContext {
	var p = new(LogicalBinaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_logicalBinaryExpression
	return p
}

func InitEmptyLogicalBinaryExpressionContext(p *LogicalBinaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_logicalBinaryExpression
}

func (*LogicalBinaryExpressionContext) IsLogicalBinaryExpressionContext() {}

func NewLogicalBinaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalBinaryExpressionContext {
	var p = new(LogicalBinaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_logicalBinaryExpression

	return p
}

func (s *LogicalBinaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalBinaryExpressionContext) DEDUCT_RE_ID_REF() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_RE_ID_REF, 0)
}

func (s *LogicalBinaryExpressionContext) LogicalExpression() ILogicalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalExpressionContext)
}

func (s *LogicalBinaryExpressionContext) DEDUCT_AND() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_AND, 0)
}

func (s *LogicalBinaryExpressionContext) DEDUCT_OR() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDEDUCT_OR, 0)
}

func (s *LogicalBinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalBinaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalBinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterLogicalBinaryExpression(s)
	}
}

func (s *LogicalBinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitLogicalBinaryExpression(s)
	}
}

func (p *ReDSLParser) LogicalBinaryExpression() (localctx ILogicalBinaryExpressionContext) {
	localctx = NewLogicalBinaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ReDSLParserRULE_logicalBinaryExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(ReDSLParserDEDUCT_RE_ID_REF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ReDSLParserDEDUCT_AND || _la == ReDSLParserDEDUCT_OR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(126)
		p.LogicalExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExampleDeclContext is an interface to support dynamic dispatch.
type IExampleDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_EXAMPLE_MARKER() antlr.TerminalNode
	RunningText() IRunningTextContext

	// IsExampleDeclContext differentiates from other interfaces.
	IsExampleDeclContext()
}

type ExampleDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExampleDeclContext() *ExampleDeclContext {
	var p = new(ExampleDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_exampleDecl
	return p
}

func InitEmptyExampleDeclContext(p *ExampleDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_exampleDecl
}

func (*ExampleDeclContext) IsExampleDeclContext() {}

func NewExampleDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExampleDeclContext {
	var p = new(ExampleDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_exampleDecl

	return p
}

func (s *ExampleDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ExampleDeclContext) TEXT_EXAMPLE_MARKER() antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_EXAMPLE_MARKER, 0)
}

func (s *ExampleDeclContext) RunningText() IRunningTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunningTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunningTextContext)
}

func (s *ExampleDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExampleDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExampleDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterExampleDecl(s)
	}
}

func (s *ExampleDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitExampleDecl(s)
	}
}

func (p *ReDSLParser) ExampleDecl() (localctx IExampleDeclContext) {
	localctx = NewExampleDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ReDSLParserRULE_exampleDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(ReDSLParserTEXT_EXAMPLE_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.RunningText()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRationaleDeclContext is an interface to support dynamic dispatch.
type IRationaleDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_RATIO_MARKER() antlr.TerminalNode
	RunningText() IRunningTextContext

	// IsRationaleDeclContext differentiates from other interfaces.
	IsRationaleDeclContext()
}

type RationaleDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRationaleDeclContext() *RationaleDeclContext {
	var p = new(RationaleDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_rationaleDecl
	return p
}

func InitEmptyRationaleDeclContext(p *RationaleDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_rationaleDecl
}

func (*RationaleDeclContext) IsRationaleDeclContext() {}

func NewRationaleDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RationaleDeclContext {
	var p = new(RationaleDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_rationaleDecl

	return p
}

func (s *RationaleDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *RationaleDeclContext) TEXT_RATIO_MARKER() antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_RATIO_MARKER, 0)
}

func (s *RationaleDeclContext) RunningText() IRunningTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunningTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunningTextContext)
}

func (s *RationaleDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RationaleDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RationaleDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterRationaleDecl(s)
	}
}

func (s *RationaleDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitRationaleDecl(s)
	}
}

func (p *ReDSLParser) RationaleDecl() (localctx IRationaleDeclContext) {
	localctx = NewRationaleDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ReDSLParserRULE_rationaleDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(ReDSLParserTEXT_RATIO_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.RunningText()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReferenceDeclContext is an interface to support dynamic dispatch.
type IReferenceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_REF_MARKER() antlr.TerminalNode
	RunningText() IRunningTextContext

	// IsReferenceDeclContext differentiates from other interfaces.
	IsReferenceDeclContext()
}

type ReferenceDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceDeclContext() *ReferenceDeclContext {
	var p = new(ReferenceDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_referenceDecl
	return p
}

func InitEmptyReferenceDeclContext(p *ReferenceDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_referenceDecl
}

func (*ReferenceDeclContext) IsReferenceDeclContext() {}

func NewReferenceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceDeclContext {
	var p = new(ReferenceDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_referenceDecl

	return p
}

func (s *ReferenceDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceDeclContext) TEXT_REF_MARKER() antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_REF_MARKER, 0)
}

func (s *ReferenceDeclContext) RunningText() IRunningTextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunningTextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunningTextContext)
}

func (s *ReferenceDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterReferenceDecl(s)
	}
}

func (s *ReferenceDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitReferenceDecl(s)
	}
}

func (p *ReDSLParser) ReferenceDecl() (localctx IReferenceDeclContext) {
	localctx = NewReferenceDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ReDSLParserRULE_referenceDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(ReDSLParserTEXT_REF_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.RunningText()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRunningTextContext is an interface to support dynamic dispatch.
type IRunningTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT_CONTENT() []antlr.TerminalNode
	TEXT_CONTENT(i int) antlr.TerminalNode
	AllTEXT_ESC_SEQ() []antlr.TerminalNode
	TEXT_ESC_SEQ(i int) antlr.TerminalNode
	AllTEXT_RE_ID_REF() []antlr.TerminalNode
	TEXT_RE_ID_REF(i int) antlr.TerminalNode
	AllTEXT_TERM_REF() []antlr.TerminalNode
	TEXT_TERM_REF(i int) antlr.TerminalNode
	AllTEXT_COMMENT() []antlr.TerminalNode
	TEXT_COMMENT(i int) antlr.TerminalNode
	AllTEXT_URI() []antlr.TerminalNode
	TEXT_URI(i int) antlr.TerminalNode
	AllTEXT_WS() []antlr.TerminalNode
	TEXT_WS(i int) antlr.TerminalNode
	AllTEXT_LINE_BREAK() []antlr.TerminalNode
	TEXT_LINE_BREAK(i int) antlr.TerminalNode
	AllMathModeExpression() []IMathModeExpressionContext
	MathModeExpression(i int) IMathModeExpressionContext

	// IsRunningTextContext differentiates from other interfaces.
	IsRunningTextContext()
}

type RunningTextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunningTextContext() *RunningTextContext {
	var p = new(RunningTextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_runningText
	return p
}

func InitEmptyRunningTextContext(p *RunningTextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_runningText
}

func (*RunningTextContext) IsRunningTextContext() {}

func NewRunningTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunningTextContext {
	var p = new(RunningTextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_runningText

	return p
}

func (s *RunningTextContext) GetParser() antlr.Parser { return s.parser }

func (s *RunningTextContext) AllTEXT_CONTENT() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserTEXT_CONTENT)
}

func (s *RunningTextContext) TEXT_CONTENT(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_CONTENT, i)
}

func (s *RunningTextContext) AllTEXT_ESC_SEQ() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserTEXT_ESC_SEQ)
}

func (s *RunningTextContext) TEXT_ESC_SEQ(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_ESC_SEQ, i)
}

func (s *RunningTextContext) AllTEXT_RE_ID_REF() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserTEXT_RE_ID_REF)
}

func (s *RunningTextContext) TEXT_RE_ID_REF(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_RE_ID_REF, i)
}

func (s *RunningTextContext) AllTEXT_TERM_REF() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserTEXT_TERM_REF)
}

func (s *RunningTextContext) TEXT_TERM_REF(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_TERM_REF, i)
}

func (s *RunningTextContext) AllTEXT_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserTEXT_COMMENT)
}

func (s *RunningTextContext) TEXT_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_COMMENT, i)
}

func (s *RunningTextContext) AllTEXT_URI() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserTEXT_URI)
}

func (s *RunningTextContext) TEXT_URI(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_URI, i)
}

func (s *RunningTextContext) AllTEXT_WS() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserTEXT_WS)
}

func (s *RunningTextContext) TEXT_WS(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_WS, i)
}

func (s *RunningTextContext) AllTEXT_LINE_BREAK() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserTEXT_LINE_BREAK)
}

func (s *RunningTextContext) TEXT_LINE_BREAK(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_LINE_BREAK, i)
}

func (s *RunningTextContext) AllMathModeExpression() []IMathModeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMathModeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMathModeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMathModeExpressionContext); ok {
			tst[i] = t.(IMathModeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RunningTextContext) MathModeExpression(i int) IMathModeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathModeExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathModeExpressionContext)
}

func (s *RunningTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunningTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RunningTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterRunningText(s)
	}
}

func (s *RunningTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitRunningText(s)
	}
}

func (p *ReDSLParser) RunningText() (localctx IRunningTextContext) {
	localctx = NewRunningTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ReDSLParserRULE_runningText)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2115764224) != 0) {
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ReDSLParserTEXT_CONTENT:
			{
				p.SetState(137)
				p.Match(ReDSLParserTEXT_CONTENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ReDSLParserTEXT_ESC_SEQ:
			{
				p.SetState(138)
				p.Match(ReDSLParserTEXT_ESC_SEQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ReDSLParserTEXT_RE_ID_REF:
			{
				p.SetState(139)
				p.Match(ReDSLParserTEXT_RE_ID_REF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ReDSLParserTEXT_TERM_REF:
			{
				p.SetState(140)
				p.Match(ReDSLParserTEXT_TERM_REF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ReDSLParserTEXT_COMMENT:
			{
				p.SetState(141)
				p.Match(ReDSLParserTEXT_COMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ReDSLParserTEXT_URI:
			{
				p.SetState(142)
				p.Match(ReDSLParserTEXT_URI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ReDSLParserTEXT_WS:
			{
				p.SetState(143)
				p.Match(ReDSLParserTEXT_WS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ReDSLParserTEXT_LINE_BREAK:
			{
				p.SetState(144)
				p.Match(ReDSLParserTEXT_LINE_BREAK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ReDSLParserTEXT_START_MATH:
			{
				p.SetState(145)
				p.MathModeExpression()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterBlockContext is an interface to support dynamic dispatch.
type IParameterBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_START() antlr.TerminalNode
	PARAM_CLOSE() antlr.TerminalNode
	AllParamDecl() []IParamDeclContext
	ParamDecl(i int) IParamDeclContext

	// IsParameterBlockContext differentiates from other interfaces.
	IsParameterBlockContext()
}

type ParameterBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterBlockContext() *ParameterBlockContext {
	var p = new(ParameterBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_parameterBlock
	return p
}

func InitEmptyParameterBlockContext(p *ParameterBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_parameterBlock
}

func (*ParameterBlockContext) IsParameterBlockContext() {}

func NewParameterBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterBlockContext {
	var p = new(ParameterBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_parameterBlock

	return p
}

func (s *ParameterBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterBlockContext) PARAM_START() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_START, 0)
}

func (s *ParameterBlockContext) PARAM_CLOSE() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_CLOSE, 0)
}

func (s *ParameterBlockContext) AllParamDecl() []IParamDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamDeclContext); ok {
			len++
		}
	}

	tst := make([]IParamDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamDeclContext); ok {
			tst[i] = t.(IParamDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParameterBlockContext) ParamDecl(i int) IParamDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamDeclContext)
}

func (s *ParameterBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterParameterBlock(s)
	}
}

func (s *ParameterBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitParameterBlock(s)
	}
}

func (p *ReDSLParser) ParameterBlock() (localctx IParameterBlockContext) {
	localctx = NewParameterBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ReDSLParserRULE_parameterBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(ReDSLParserPARAM_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ReDSLParserPARAM_ID {
		{
			p.SetState(151)
			p.ParamDecl()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(ReDSLParserPARAM_CLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamDeclContext is an interface to support dynamic dispatch.
type IParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringParamDecl() IStringParamDeclContext
	IdParamDecl() IIdParamDeclContext
	NumericParamDecl() INumericParamDeclContext
	BoolParamDecl() IBoolParamDeclContext
	IdListParamDecl() IIdListParamDeclContext

	// IsParamDeclContext differentiates from other interfaces.
	IsParamDeclContext()
}

type ParamDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamDeclContext() *ParamDeclContext {
	var p = new(ParamDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_paramDecl
	return p
}

func InitEmptyParamDeclContext(p *ParamDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_paramDecl
}

func (*ParamDeclContext) IsParamDeclContext() {}

func NewParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclContext {
	var p = new(ParamDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_paramDecl

	return p
}

func (s *ParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclContext) StringParamDecl() IStringParamDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringParamDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringParamDeclContext)
}

func (s *ParamDeclContext) IdParamDecl() IIdParamDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdParamDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdParamDeclContext)
}

func (s *ParamDeclContext) NumericParamDecl() INumericParamDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericParamDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericParamDeclContext)
}

func (s *ParamDeclContext) BoolParamDecl() IBoolParamDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolParamDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolParamDeclContext)
}

func (s *ParamDeclContext) IdListParamDecl() IIdListParamDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdListParamDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdListParamDeclContext)
}

func (s *ParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterParamDecl(s)
	}
}

func (s *ParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitParamDecl(s)
	}
}

func (p *ReDSLParser) ParamDecl() (localctx IParamDeclContext) {
	localctx = NewParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ReDSLParserRULE_paramDecl)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(159)
			p.StringParamDecl()
		}

	case 2:
		{
			p.SetState(160)
			p.IdParamDecl()
		}

	case 3:
		{
			p.SetState(161)
			p.NumericParamDecl()
		}

	case 4:
		{
			p.SetState(162)
			p.BoolParamDecl()
		}

	case 5:
		{
			p.SetState(163)
			p.IdListParamDecl()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringParamDeclContext is an interface to support dynamic dispatch.
type IStringParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_ID() antlr.TerminalNode
	PARAM_EQUALS() antlr.TerminalNode
	PARAM_STRING() antlr.TerminalNode

	// IsStringParamDeclContext differentiates from other interfaces.
	IsStringParamDeclContext()
}

type StringParamDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringParamDeclContext() *StringParamDeclContext {
	var p = new(StringParamDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_stringParamDecl
	return p
}

func InitEmptyStringParamDeclContext(p *StringParamDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_stringParamDecl
}

func (*StringParamDeclContext) IsStringParamDeclContext() {}

func NewStringParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringParamDeclContext {
	var p = new(StringParamDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_stringParamDecl

	return p
}

func (s *StringParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StringParamDeclContext) PARAM_ID() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_ID, 0)
}

func (s *StringParamDeclContext) PARAM_EQUALS() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_EQUALS, 0)
}

func (s *StringParamDeclContext) PARAM_STRING() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_STRING, 0)
}

func (s *StringParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterStringParamDecl(s)
	}
}

func (s *StringParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitStringParamDecl(s)
	}
}

func (p *ReDSLParser) StringParamDecl() (localctx IStringParamDeclContext) {
	localctx = NewStringParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ReDSLParserRULE_stringParamDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(ReDSLParserPARAM_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(ReDSLParserPARAM_EQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(ReDSLParserPARAM_STRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdParamDeclContext is an interface to support dynamic dispatch.
type IIdParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPARAM_ID() []antlr.TerminalNode
	PARAM_ID(i int) antlr.TerminalNode
	PARAM_EQUALS() antlr.TerminalNode

	// IsIdParamDeclContext differentiates from other interfaces.
	IsIdParamDeclContext()
}

type IdParamDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdParamDeclContext() *IdParamDeclContext {
	var p = new(IdParamDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_idParamDecl
	return p
}

func InitEmptyIdParamDeclContext(p *IdParamDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_idParamDecl
}

func (*IdParamDeclContext) IsIdParamDeclContext() {}

func NewIdParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdParamDeclContext {
	var p = new(IdParamDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_idParamDecl

	return p
}

func (s *IdParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *IdParamDeclContext) AllPARAM_ID() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserPARAM_ID)
}

func (s *IdParamDeclContext) PARAM_ID(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_ID, i)
}

func (s *IdParamDeclContext) PARAM_EQUALS() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_EQUALS, 0)
}

func (s *IdParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterIdParamDecl(s)
	}
}

func (s *IdParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitIdParamDecl(s)
	}
}

func (p *ReDSLParser) IdParamDecl() (localctx IIdParamDeclContext) {
	localctx = NewIdParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ReDSLParserRULE_idParamDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(ReDSLParserPARAM_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(ReDSLParserPARAM_EQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Match(ReDSLParserPARAM_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericParamDeclContext is an interface to support dynamic dispatch.
type INumericParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_ID() antlr.TerminalNode
	PARAM_EQUALS() antlr.TerminalNode
	PARAM_NUMBER() antlr.TerminalNode

	// IsNumericParamDeclContext differentiates from other interfaces.
	IsNumericParamDeclContext()
}

type NumericParamDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericParamDeclContext() *NumericParamDeclContext {
	var p = new(NumericParamDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_numericParamDecl
	return p
}

func InitEmptyNumericParamDeclContext(p *NumericParamDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_numericParamDecl
}

func (*NumericParamDeclContext) IsNumericParamDeclContext() {}

func NewNumericParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericParamDeclContext {
	var p = new(NumericParamDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_numericParamDecl

	return p
}

func (s *NumericParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericParamDeclContext) PARAM_ID() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_ID, 0)
}

func (s *NumericParamDeclContext) PARAM_EQUALS() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_EQUALS, 0)
}

func (s *NumericParamDeclContext) PARAM_NUMBER() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_NUMBER, 0)
}

func (s *NumericParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterNumericParamDecl(s)
	}
}

func (s *NumericParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitNumericParamDecl(s)
	}
}

func (p *ReDSLParser) NumericParamDecl() (localctx INumericParamDeclContext) {
	localctx = NewNumericParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ReDSLParserRULE_numericParamDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(ReDSLParserPARAM_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(ReDSLParserPARAM_EQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(ReDSLParserPARAM_NUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolParamDeclContext is an interface to support dynamic dispatch.
type IBoolParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_ID() antlr.TerminalNode
	PARAM_EQUALS() antlr.TerminalNode
	PARAM_BOOL() antlr.TerminalNode

	// IsBoolParamDeclContext differentiates from other interfaces.
	IsBoolParamDeclContext()
}

type BoolParamDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolParamDeclContext() *BoolParamDeclContext {
	var p = new(BoolParamDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_boolParamDecl
	return p
}

func InitEmptyBoolParamDeclContext(p *BoolParamDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_boolParamDecl
}

func (*BoolParamDeclContext) IsBoolParamDeclContext() {}

func NewBoolParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolParamDeclContext {
	var p = new(BoolParamDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_boolParamDecl

	return p
}

func (s *BoolParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolParamDeclContext) PARAM_ID() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_ID, 0)
}

func (s *BoolParamDeclContext) PARAM_EQUALS() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_EQUALS, 0)
}

func (s *BoolParamDeclContext) PARAM_BOOL() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_BOOL, 0)
}

func (s *BoolParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterBoolParamDecl(s)
	}
}

func (s *BoolParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitBoolParamDecl(s)
	}
}

func (p *ReDSLParser) BoolParamDecl() (localctx IBoolParamDeclContext) {
	localctx = NewBoolParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ReDSLParserRULE_boolParamDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(ReDSLParserPARAM_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ReDSLParserPARAM_EQUALS {
		{
			p.SetState(179)
			p.Match(ReDSLParserPARAM_EQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(ReDSLParserPARAM_BOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdListParamDeclContext is an interface to support dynamic dispatch.
type IIdListParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPARAM_ID() []antlr.TerminalNode
	PARAM_ID(i int) antlr.TerminalNode
	PARAM_EQUALS() antlr.TerminalNode
	PARAM_ID_LIST_START() antlr.TerminalNode
	PARAM_ID_LIST_END() antlr.TerminalNode
	AllPARAM_ID_LIST_SEP() []antlr.TerminalNode
	PARAM_ID_LIST_SEP(i int) antlr.TerminalNode

	// IsIdListParamDeclContext differentiates from other interfaces.
	IsIdListParamDeclContext()
}

type IdListParamDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdListParamDeclContext() *IdListParamDeclContext {
	var p = new(IdListParamDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_idListParamDecl
	return p
}

func InitEmptyIdListParamDeclContext(p *IdListParamDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_idListParamDecl
}

func (*IdListParamDeclContext) IsIdListParamDeclContext() {}

func NewIdListParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdListParamDeclContext {
	var p = new(IdListParamDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_idListParamDecl

	return p
}

func (s *IdListParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *IdListParamDeclContext) AllPARAM_ID() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserPARAM_ID)
}

func (s *IdListParamDeclContext) PARAM_ID(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_ID, i)
}

func (s *IdListParamDeclContext) PARAM_EQUALS() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_EQUALS, 0)
}

func (s *IdListParamDeclContext) PARAM_ID_LIST_START() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_ID_LIST_START, 0)
}

func (s *IdListParamDeclContext) PARAM_ID_LIST_END() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_ID_LIST_END, 0)
}

func (s *IdListParamDeclContext) AllPARAM_ID_LIST_SEP() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserPARAM_ID_LIST_SEP)
}

func (s *IdListParamDeclContext) PARAM_ID_LIST_SEP(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserPARAM_ID_LIST_SEP, i)
}

func (s *IdListParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdListParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterIdListParamDecl(s)
	}
}

func (s *IdListParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitIdListParamDecl(s)
	}
}

func (p *ReDSLParser) IdListParamDecl() (localctx IIdListParamDeclContext) {
	localctx = NewIdListParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ReDSLParserRULE_idListParamDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(ReDSLParserPARAM_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Match(ReDSLParserPARAM_EQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(185)
		p.Match(ReDSLParserPARAM_ID_LIST_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(ReDSLParserPARAM_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ReDSLParserPARAM_ID_LIST_SEP {
		{
			p.SetState(187)
			p.Match(ReDSLParserPARAM_ID_LIST_SEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(ReDSLParserPARAM_ID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		p.Match(ReDSLParserPARAM_ID_LIST_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMathModeExpressionContext is an interface to support dynamic dispatch.
type IMathModeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_START_MATH() antlr.TerminalNode
	MATH_CLOSE() antlr.TerminalNode
	AllMATH_ESC_SEQ() []antlr.TerminalNode
	MATH_ESC_SEQ(i int) antlr.TerminalNode
	AllMATH_CONTENT() []antlr.TerminalNode
	MATH_CONTENT(i int) antlr.TerminalNode

	// IsMathModeExpressionContext differentiates from other interfaces.
	IsMathModeExpressionContext()
}

type MathModeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathModeExpressionContext() *MathModeExpressionContext {
	var p = new(MathModeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_mathModeExpression
	return p
}

func InitEmptyMathModeExpressionContext(p *MathModeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_mathModeExpression
}

func (*MathModeExpressionContext) IsMathModeExpressionContext() {}

func NewMathModeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathModeExpressionContext {
	var p = new(MathModeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_mathModeExpression

	return p
}

func (s *MathModeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MathModeExpressionContext) TEXT_START_MATH() antlr.TerminalNode {
	return s.GetToken(ReDSLParserTEXT_START_MATH, 0)
}

func (s *MathModeExpressionContext) MATH_CLOSE() antlr.TerminalNode {
	return s.GetToken(ReDSLParserMATH_CLOSE, 0)
}

func (s *MathModeExpressionContext) AllMATH_ESC_SEQ() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserMATH_ESC_SEQ)
}

func (s *MathModeExpressionContext) MATH_ESC_SEQ(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserMATH_ESC_SEQ, i)
}

func (s *MathModeExpressionContext) AllMATH_CONTENT() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserMATH_CONTENT)
}

func (s *MathModeExpressionContext) MATH_CONTENT(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserMATH_CONTENT, i)
}

func (s *MathModeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathModeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathModeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterMathModeExpression(s)
	}
}

func (s *MathModeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitMathModeExpression(s)
	}
}

func (p *ReDSLParser) MathModeExpression() (localctx IMathModeExpressionContext) {
	localctx = NewMathModeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ReDSLParserRULE_mathModeExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(ReDSLParserTEXT_START_MATH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ReDSLParserMATH_ESC_SEQ || _la == ReDSLParserMATH_CONTENT {
		{
			p.SetState(197)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ReDSLParserMATH_ESC_SEQ || _la == ReDSLParserMATH_CONTENT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(ReDSLParserMATH_CLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDocumentDeclContext is an interface to support dynamic dispatch.
type IDocumentDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOCUMENT_KEYWORD() antlr.TerminalNode
	ID_STR() antlr.TerminalNode
	BLOCK_START() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
	AllPackageRef() []IPackageRefContext
	PackageRef(i int) IPackageRefContext
	AllVERSION_INFO_KEYWORD() []antlr.TerminalNode
	VERSION_INFO_KEYWORD(i int) antlr.TerminalNode
	AllGLOSSARY_KEYWORD() []antlr.TerminalNode
	GLOSSARY_KEYWORD(i int) antlr.TerminalNode
	AllLOCAL_KEYWORD() []antlr.TerminalNode
	LOCAL_KEYWORD(i int) antlr.TerminalNode
	AllGLOBAL_KEYWORD() []antlr.TerminalNode
	GLOBAL_KEYWORD(i int) antlr.TerminalNode

	// IsDocumentDeclContext differentiates from other interfaces.
	IsDocumentDeclContext()
}

type DocumentDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentDeclContext() *DocumentDeclContext {
	var p = new(DocumentDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_documentDecl
	return p
}

func InitEmptyDocumentDeclContext(p *DocumentDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_documentDecl
}

func (*DocumentDeclContext) IsDocumentDeclContext() {}

func NewDocumentDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentDeclContext {
	var p = new(DocumentDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_documentDecl

	return p
}

func (s *DocumentDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentDeclContext) DOCUMENT_KEYWORD() antlr.TerminalNode {
	return s.GetToken(ReDSLParserDOCUMENT_KEYWORD, 0)
}

func (s *DocumentDeclContext) ID_STR() antlr.TerminalNode {
	return s.GetToken(ReDSLParserID_STR, 0)
}

func (s *DocumentDeclContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(ReDSLParserBLOCK_START, 0)
}

func (s *DocumentDeclContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(ReDSLParserBLOCK_END, 0)
}

func (s *DocumentDeclContext) AllPackageRef() []IPackageRefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPackageRefContext); ok {
			len++
		}
	}

	tst := make([]IPackageRefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPackageRefContext); ok {
			tst[i] = t.(IPackageRefContext)
			i++
		}
	}

	return tst
}

func (s *DocumentDeclContext) PackageRef(i int) IPackageRefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageRefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageRefContext)
}

func (s *DocumentDeclContext) AllVERSION_INFO_KEYWORD() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserVERSION_INFO_KEYWORD)
}

func (s *DocumentDeclContext) VERSION_INFO_KEYWORD(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserVERSION_INFO_KEYWORD, i)
}

func (s *DocumentDeclContext) AllGLOSSARY_KEYWORD() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserGLOSSARY_KEYWORD)
}

func (s *DocumentDeclContext) GLOSSARY_KEYWORD(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserGLOSSARY_KEYWORD, i)
}

func (s *DocumentDeclContext) AllLOCAL_KEYWORD() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserLOCAL_KEYWORD)
}

func (s *DocumentDeclContext) LOCAL_KEYWORD(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserLOCAL_KEYWORD, i)
}

func (s *DocumentDeclContext) AllGLOBAL_KEYWORD() []antlr.TerminalNode {
	return s.GetTokens(ReDSLParserGLOBAL_KEYWORD)
}

func (s *DocumentDeclContext) GLOBAL_KEYWORD(i int) antlr.TerminalNode {
	return s.GetToken(ReDSLParserGLOBAL_KEYWORD, i)
}

func (s *DocumentDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterDocumentDecl(s)
	}
}

func (s *DocumentDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitDocumentDecl(s)
	}
}

func (p *ReDSLParser) DocumentDecl() (localctx IDocumentDeclContext) {
	localctx = NewDocumentDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ReDSLParserRULE_documentDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(ReDSLParserDOCUMENT_KEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(ReDSLParserID_STR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Match(ReDSLParserBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1552) != 0 {
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ReDSLParserPACKAGE_KEYWORD:
			{
				p.SetState(208)
				p.PackageRef()
			}

		case ReDSLParserVERSION_INFO_KEYWORD:
			{
				p.SetState(209)
				p.Match(ReDSLParserVERSION_INFO_KEYWORD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ReDSLParserGLOSSARY_KEYWORD:
			{
				p.SetState(210)
				p.Match(ReDSLParserGLOSSARY_KEYWORD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == ReDSLParserLOCAL_KEYWORD || _la == ReDSLParserGLOBAL_KEYWORD {
				{
					p.SetState(211)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ReDSLParserLOCAL_KEYWORD || _la == ReDSLParserGLOBAL_KEYWORD) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(219)
		p.Match(ReDSLParserBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageRefContext is an interface to support dynamic dispatch.
type IPackageRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE_KEYWORD() antlr.TerminalNode
	ID_STR() antlr.TerminalNode
	BLOCK_START() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
	AllFileRef() []IFileRefContext
	FileRef(i int) IFileRefContext

	// IsPackageRefContext differentiates from other interfaces.
	IsPackageRefContext()
}

type PackageRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageRefContext() *PackageRefContext {
	var p = new(PackageRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_packageRef
	return p
}

func InitEmptyPackageRefContext(p *PackageRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_packageRef
}

func (*PackageRefContext) IsPackageRefContext() {}

func NewPackageRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageRefContext {
	var p = new(PackageRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_packageRef

	return p
}

func (s *PackageRefContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageRefContext) PACKAGE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(ReDSLParserPACKAGE_KEYWORD, 0)
}

func (s *PackageRefContext) ID_STR() antlr.TerminalNode {
	return s.GetToken(ReDSLParserID_STR, 0)
}

func (s *PackageRefContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(ReDSLParserBLOCK_START, 0)
}

func (s *PackageRefContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(ReDSLParserBLOCK_END, 0)
}

func (s *PackageRefContext) AllFileRef() []IFileRefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFileRefContext); ok {
			len++
		}
	}

	tst := make([]IFileRefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFileRefContext); ok {
			tst[i] = t.(IFileRefContext)
			i++
		}
	}

	return tst
}

func (s *PackageRefContext) FileRef(i int) IFileRefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileRefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFileRefContext)
}

func (s *PackageRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterPackageRef(s)
	}
}

func (s *PackageRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitPackageRef(s)
	}
}

func (p *ReDSLParser) PackageRef() (localctx IPackageRefContext) {
	localctx = NewPackageRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ReDSLParserRULE_packageRef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(ReDSLParserPACKAGE_KEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(ReDSLParserID_STR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ReDSLParserBLOCK_START {
		{
			p.SetState(223)
			p.Match(ReDSLParserBLOCK_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ReDSLParserFILE_KEYWORD {
			{
				p.SetState(224)
				p.FileRef()
			}

			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(230)
			p.Match(ReDSLParserBLOCK_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFileRefContext is an interface to support dynamic dispatch.
type IFileRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILE_KEYWORD() antlr.TerminalNode
	ID_STR() antlr.TerminalNode

	// IsFileRefContext differentiates from other interfaces.
	IsFileRefContext()
}

type FileRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileRefContext() *FileRefContext {
	var p = new(FileRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_fileRef
	return p
}

func InitEmptyFileRefContext(p *FileRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ReDSLParserRULE_fileRef
}

func (*FileRefContext) IsFileRefContext() {}

func NewFileRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileRefContext {
	var p = new(FileRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReDSLParserRULE_fileRef

	return p
}

func (s *FileRefContext) GetParser() antlr.Parser { return s.parser }

func (s *FileRefContext) FILE_KEYWORD() antlr.TerminalNode {
	return s.GetToken(ReDSLParserFILE_KEYWORD, 0)
}

func (s *FileRefContext) ID_STR() antlr.TerminalNode {
	return s.GetToken(ReDSLParserID_STR, 0)
}

func (s *FileRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.EnterFileRef(s)
	}
}

func (s *FileRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReDSLParserListener); ok {
		listenerT.ExitFileRef(s)
	}
}

func (p *ReDSLParser) FileRef() (localctx IFileRefContext) {
	localctx = NewFileRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ReDSLParserRULE_fileRef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(ReDSLParserFILE_KEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.Match(ReDSLParserID_STR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
