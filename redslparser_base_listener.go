// Code generated from ReDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package redsl_parser // ReDSLParser
import "github.com/antlr4-go/antlr/v4"

// BaseReDSLParserListener is a complete listener for a parse tree produced by ReDSLParser.
type BaseReDSLParserListener struct{}

var _ ReDSLParserListener = &BaseReDSLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseReDSLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseReDSLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseReDSLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseReDSLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseReDSLParserListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseReDSLParserListener) ExitParse(ctx *ParseContext) {}

// EnterFileDecl is called when production fileDecl is entered.
func (s *BaseReDSLParserListener) EnterFileDecl(ctx *FileDeclContext) {}

// ExitFileDecl is called when production fileDecl is exited.
func (s *BaseReDSLParserListener) ExitFileDecl(ctx *FileDeclContext) {}

// EnterPackageDecl is called when production packageDecl is entered.
func (s *BaseReDSLParserListener) EnterPackageDecl(ctx *PackageDeclContext) {}

// ExitPackageDecl is called when production packageDecl is exited.
func (s *BaseReDSLParserListener) ExitPackageDecl(ctx *PackageDeclContext) {}

// EnterRequirementDecl is called when production requirementDecl is entered.
func (s *BaseReDSLParserListener) EnterRequirementDecl(ctx *RequirementDeclContext) {}

// ExitRequirementDecl is called when production requirementDecl is exited.
func (s *BaseReDSLParserListener) ExitRequirementDecl(ctx *RequirementDeclContext) {}

// EnterRequirementDeduction is called when production requirementDeduction is entered.
func (s *BaseReDSLParserListener) EnterRequirementDeduction(ctx *RequirementDeductionContext) {}

// ExitRequirementDeduction is called when production requirementDeduction is exited.
func (s *BaseReDSLParserListener) ExitRequirementDeduction(ctx *RequirementDeductionContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseReDSLParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseReDSLParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterLogicalAtom is called when production logicalAtom is entered.
func (s *BaseReDSLParserListener) EnterLogicalAtom(ctx *LogicalAtomContext) {}

// ExitLogicalAtom is called when production logicalAtom is exited.
func (s *BaseReDSLParserListener) ExitLogicalAtom(ctx *LogicalAtomContext) {}

// EnterLogicalUnaryExpression is called when production logicalUnaryExpression is entered.
func (s *BaseReDSLParserListener) EnterLogicalUnaryExpression(ctx *LogicalUnaryExpressionContext) {}

// ExitLogicalUnaryExpression is called when production logicalUnaryExpression is exited.
func (s *BaseReDSLParserListener) ExitLogicalUnaryExpression(ctx *LogicalUnaryExpressionContext) {}

// EnterLogicalBinaryExpression is called when production logicalBinaryExpression is entered.
func (s *BaseReDSLParserListener) EnterLogicalBinaryExpression(ctx *LogicalBinaryExpressionContext) {}

// ExitLogicalBinaryExpression is called when production logicalBinaryExpression is exited.
func (s *BaseReDSLParserListener) ExitLogicalBinaryExpression(ctx *LogicalBinaryExpressionContext) {}

// EnterExampleDecl is called when production exampleDecl is entered.
func (s *BaseReDSLParserListener) EnterExampleDecl(ctx *ExampleDeclContext) {}

// ExitExampleDecl is called when production exampleDecl is exited.
func (s *BaseReDSLParserListener) ExitExampleDecl(ctx *ExampleDeclContext) {}

// EnterRationaleDecl is called when production rationaleDecl is entered.
func (s *BaseReDSLParserListener) EnterRationaleDecl(ctx *RationaleDeclContext) {}

// ExitRationaleDecl is called when production rationaleDecl is exited.
func (s *BaseReDSLParserListener) ExitRationaleDecl(ctx *RationaleDeclContext) {}

// EnterReferenceDecl is called when production referenceDecl is entered.
func (s *BaseReDSLParserListener) EnterReferenceDecl(ctx *ReferenceDeclContext) {}

// ExitReferenceDecl is called when production referenceDecl is exited.
func (s *BaseReDSLParserListener) ExitReferenceDecl(ctx *ReferenceDeclContext) {}

// EnterRunningText is called when production runningText is entered.
func (s *BaseReDSLParserListener) EnterRunningText(ctx *RunningTextContext) {}

// ExitRunningText is called when production runningText is exited.
func (s *BaseReDSLParserListener) ExitRunningText(ctx *RunningTextContext) {}

// EnterParameterBlock is called when production parameterBlock is entered.
func (s *BaseReDSLParserListener) EnterParameterBlock(ctx *ParameterBlockContext) {}

// ExitParameterBlock is called when production parameterBlock is exited.
func (s *BaseReDSLParserListener) ExitParameterBlock(ctx *ParameterBlockContext) {}

// EnterParamDecl is called when production paramDecl is entered.
func (s *BaseReDSLParserListener) EnterParamDecl(ctx *ParamDeclContext) {}

// ExitParamDecl is called when production paramDecl is exited.
func (s *BaseReDSLParserListener) ExitParamDecl(ctx *ParamDeclContext) {}

// EnterStringParamDecl is called when production stringParamDecl is entered.
func (s *BaseReDSLParserListener) EnterStringParamDecl(ctx *StringParamDeclContext) {}

// ExitStringParamDecl is called when production stringParamDecl is exited.
func (s *BaseReDSLParserListener) ExitStringParamDecl(ctx *StringParamDeclContext) {}

// EnterIdParamDecl is called when production idParamDecl is entered.
func (s *BaseReDSLParserListener) EnterIdParamDecl(ctx *IdParamDeclContext) {}

// ExitIdParamDecl is called when production idParamDecl is exited.
func (s *BaseReDSLParserListener) ExitIdParamDecl(ctx *IdParamDeclContext) {}

// EnterNumericParamDecl is called when production numericParamDecl is entered.
func (s *BaseReDSLParserListener) EnterNumericParamDecl(ctx *NumericParamDeclContext) {}

// ExitNumericParamDecl is called when production numericParamDecl is exited.
func (s *BaseReDSLParserListener) ExitNumericParamDecl(ctx *NumericParamDeclContext) {}

// EnterBoolParamDecl is called when production boolParamDecl is entered.
func (s *BaseReDSLParserListener) EnterBoolParamDecl(ctx *BoolParamDeclContext) {}

// ExitBoolParamDecl is called when production boolParamDecl is exited.
func (s *BaseReDSLParserListener) ExitBoolParamDecl(ctx *BoolParamDeclContext) {}

// EnterIdListParamDecl is called when production idListParamDecl is entered.
func (s *BaseReDSLParserListener) EnterIdListParamDecl(ctx *IdListParamDeclContext) {}

// ExitIdListParamDecl is called when production idListParamDecl is exited.
func (s *BaseReDSLParserListener) ExitIdListParamDecl(ctx *IdListParamDeclContext) {}

// EnterMathModeExpression is called when production mathModeExpression is entered.
func (s *BaseReDSLParserListener) EnterMathModeExpression(ctx *MathModeExpressionContext) {}

// ExitMathModeExpression is called when production mathModeExpression is exited.
func (s *BaseReDSLParserListener) ExitMathModeExpression(ctx *MathModeExpressionContext) {}

// EnterDocumentDecl is called when production documentDecl is entered.
func (s *BaseReDSLParserListener) EnterDocumentDecl(ctx *DocumentDeclContext) {}

// ExitDocumentDecl is called when production documentDecl is exited.
func (s *BaseReDSLParserListener) ExitDocumentDecl(ctx *DocumentDeclContext) {}

// EnterPackageRef is called when production packageRef is entered.
func (s *BaseReDSLParserListener) EnterPackageRef(ctx *PackageRefContext) {}

// ExitPackageRef is called when production packageRef is exited.
func (s *BaseReDSLParserListener) ExitPackageRef(ctx *PackageRefContext) {}

// EnterFileRef is called when production fileRef is entered.
func (s *BaseReDSLParserListener) EnterFileRef(ctx *FileRefContext) {}

// ExitFileRef is called when production fileRef is exited.
func (s *BaseReDSLParserListener) ExitFileRef(ctx *FileRefContext) {}
