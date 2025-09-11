#!/bin/bash
SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"
SOURCEPATH="$SCRIPTPATH/antlr4/"
# unfortunately, the generated code for the go parser has to be in the root directory
# of the GitHub project for the go package name to be correct
TARGETPATH="$SCRIPTPATH/"
cd "$TARGETPATH"
rm -f *.interp *.tokens
rm -f redsl_lexer.go redsl_parser.go
rm -f redslparser_base_listener.go redslparser_listener.go
cd "$SOURCEPATH"
java -Xmx500M -cp "antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool -o "$TARGETPATH" -package redsl_parser -Dlanguage=Go ReDSLParser.g4 ReDSLLexer.g4
cd "$TARGETPATH"
rm -f *.interp *.tokens