package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

/*
Rather than creating a new object.Boolean every time
we encounter a true or false, let's reference them instead of creating
new ones
*/
var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

//Eval function helps the REPL to evaluates the AST.
//It is a tree-walking interpreter.
//Eval is ...
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	//Recurisvely evaluate each Statement of the Program
	case *ast.Program:
		return evalStatements(node.Statements)

	//Evaluate an Expression Statement
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	//Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return nativeBooltoBooleanObject(node.Value)
	}

	return nil
}

func nativeBooltoBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
