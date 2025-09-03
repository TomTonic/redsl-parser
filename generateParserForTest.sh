#!/bin/bash
SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"
SOURCEPATH="$SCRIPTPATH/src/antlr4/"
TARGETPATH="$SCRIPTPATH/src/generated/"
rm -rf "$TARGETPATH/*"
cd "$SOURCEPATH"
java -Xmx500M -cp "antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool -o "$TARGETPATH" -package redsl_parser -Dlanguage=Go ReDSLParser.g4 ReDSLLexer.g4