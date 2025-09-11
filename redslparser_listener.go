// Code generated from ReDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package redsl_parser // ReDSLParser
import "github.com/antlr4-go/antlr/v4"

// ReDSLParserListener is a complete listener for a parse tree produced by ReDSLParser.
type ReDSLParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterFileDecl is called when entering the fileDecl production.
	EnterFileDecl(c *FileDeclContext)

	// EnterPackageDecl is called when entering the packageDecl production.
	EnterPackageDecl(c *PackageDeclContext)

	// EnterRequirementDecl is called when entering the requirementDecl production.
	EnterRequirementDecl(c *RequirementDeclContext)

	// EnterRequirementDeduction is called when entering the requirementDeduction production.
	EnterRequirementDeduction(c *RequirementDeductionContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterLogicalAtom is called when entering the logicalAtom production.
	EnterLogicalAtom(c *LogicalAtomContext)

	// EnterLogicalUnaryExpression is called when entering the logicalUnaryExpression production.
	EnterLogicalUnaryExpression(c *LogicalUnaryExpressionContext)

	// EnterLogicalBinaryExpression is called when entering the logicalBinaryExpression production.
	EnterLogicalBinaryExpression(c *LogicalBinaryExpressionContext)

	// EnterExampleDecl is called when entering the exampleDecl production.
	EnterExampleDecl(c *ExampleDeclContext)

	// EnterRationaleDecl is called when entering the rationaleDecl production.
	EnterRationaleDecl(c *RationaleDeclContext)

	// EnterReferenceDecl is called when entering the referenceDecl production.
	EnterReferenceDecl(c *ReferenceDeclContext)

	// EnterRunningText is called when entering the runningText production.
	EnterRunningText(c *RunningTextContext)

	// EnterParameterBlock is called when entering the parameterBlock production.
	EnterParameterBlock(c *ParameterBlockContext)

	// EnterParamDecl is called when entering the paramDecl production.
	EnterParamDecl(c *ParamDeclContext)

	// EnterStringParamDecl is called when entering the stringParamDecl production.
	EnterStringParamDecl(c *StringParamDeclContext)

	// EnterIdParamDecl is called when entering the idParamDecl production.
	EnterIdParamDecl(c *IdParamDeclContext)

	// EnterNumericParamDecl is called when entering the numericParamDecl production.
	EnterNumericParamDecl(c *NumericParamDeclContext)

	// EnterBoolParamDecl is called when entering the boolParamDecl production.
	EnterBoolParamDecl(c *BoolParamDeclContext)

	// EnterIdListParamDecl is called when entering the idListParamDecl production.
	EnterIdListParamDecl(c *IdListParamDeclContext)

	// EnterMathModeExpression is called when entering the mathModeExpression production.
	EnterMathModeExpression(c *MathModeExpressionContext)

	// EnterDocumentDecl is called when entering the documentDecl production.
	EnterDocumentDecl(c *DocumentDeclContext)

	// EnterPackageRef is called when entering the packageRef production.
	EnterPackageRef(c *PackageRefContext)

	// EnterFileRef is called when entering the fileRef production.
	EnterFileRef(c *FileRefContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitFileDecl is called when exiting the fileDecl production.
	ExitFileDecl(c *FileDeclContext)

	// ExitPackageDecl is called when exiting the packageDecl production.
	ExitPackageDecl(c *PackageDeclContext)

	// ExitRequirementDecl is called when exiting the requirementDecl production.
	ExitRequirementDecl(c *RequirementDeclContext)

	// ExitRequirementDeduction is called when exiting the requirementDeduction production.
	ExitRequirementDeduction(c *RequirementDeductionContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitLogicalAtom is called when exiting the logicalAtom production.
	ExitLogicalAtom(c *LogicalAtomContext)

	// ExitLogicalUnaryExpression is called when exiting the logicalUnaryExpression production.
	ExitLogicalUnaryExpression(c *LogicalUnaryExpressionContext)

	// ExitLogicalBinaryExpression is called when exiting the logicalBinaryExpression production.
	ExitLogicalBinaryExpression(c *LogicalBinaryExpressionContext)

	// ExitExampleDecl is called when exiting the exampleDecl production.
	ExitExampleDecl(c *ExampleDeclContext)

	// ExitRationaleDecl is called when exiting the rationaleDecl production.
	ExitRationaleDecl(c *RationaleDeclContext)

	// ExitReferenceDecl is called when exiting the referenceDecl production.
	ExitReferenceDecl(c *ReferenceDeclContext)

	// ExitRunningText is called when exiting the runningText production.
	ExitRunningText(c *RunningTextContext)

	// ExitParameterBlock is called when exiting the parameterBlock production.
	ExitParameterBlock(c *ParameterBlockContext)

	// ExitParamDecl is called when exiting the paramDecl production.
	ExitParamDecl(c *ParamDeclContext)

	// ExitStringParamDecl is called when exiting the stringParamDecl production.
	ExitStringParamDecl(c *StringParamDeclContext)

	// ExitIdParamDecl is called when exiting the idParamDecl production.
	ExitIdParamDecl(c *IdParamDeclContext)

	// ExitNumericParamDecl is called when exiting the numericParamDecl production.
	ExitNumericParamDecl(c *NumericParamDeclContext)

	// ExitBoolParamDecl is called when exiting the boolParamDecl production.
	ExitBoolParamDecl(c *BoolParamDeclContext)

	// ExitIdListParamDecl is called when exiting the idListParamDecl production.
	ExitIdListParamDecl(c *IdListParamDeclContext)

	// ExitMathModeExpression is called when exiting the mathModeExpression production.
	ExitMathModeExpression(c *MathModeExpressionContext)

	// ExitDocumentDecl is called when exiting the documentDecl production.
	ExitDocumentDecl(c *DocumentDeclContext)

	// ExitPackageRef is called when exiting the packageRef production.
	ExitPackageRef(c *PackageRefContext)

	// ExitFileRef is called when exiting the fileRef production.
	ExitFileRef(c *FileRefContext)
}
