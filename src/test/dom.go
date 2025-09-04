package test

import (
	"fmt"
	redsl_parser "redsl_parser/generated"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type DOMNode struct {
	Type     string
	Text     string
	lvl      int
	Children []*DOMNode
}

func (n *DOMNode) String() string {
	indent := strings.Repeat("  ", n.lvl)
	return fmt.Sprintf("%s%s: %s", indent, n.Type, n.Text)
}

func (n *DOMNode) PrintTree() string {
	myself := n.String()
	if len(n.Children) == 0 {
		return myself
	}
	subtree := myself
	for _, child := range n.Children {
		subtree += "\n" + child.PrintTree()
	}
	return subtree
}

var ruleNames = redsl_parser.NewReDSLParser(nil).GetRuleNames()

func buildDOM(node antlr.Tree, level int) *DOMNode {
	// Rekursiv durchreichen bei Kettenproduktion
	//if node.GetChildCount() == 1 {
	//	return buildDOM(node.GetChild(0), level)
	//}

	var text string
	var nodeType string

	switch n := node.(type) {
	case antlr.TerminalNode:
		text = n.GetSymbol().GetText()
		nodeType = "Terminal"
	case antlr.ParserRuleContext:
		ruleIndex := n.GetRuleIndex()
		if ruleIndex >= 0 && ruleIndex < len(ruleNames) {
			text = ruleNames[ruleIndex]
		} else {
			text = fmt.Sprintf("Rule[%d]", ruleIndex)
		}
		nodeType = "Rule"
	default:
		text = ""
		nodeType = fmt.Sprintf("%T", node)
	}

	dom := &DOMNode{
		Type: nodeType,
		Text: text,
		lvl:  level,
	}

	for i := 0; i < node.GetChildCount(); i++ {
		child := buildDOM(node.GetChild(i), level+1)
		if child != nil {
			dom.Children = append(dom.Children, child)
		}
	}

	return dom
}
