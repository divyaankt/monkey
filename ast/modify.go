package ast

type ModifierFunc func(Node) Node

func Modify(node Node, modifier ModifierFunc) Node {
	//Recursively walk down the children of any given ast.Node
	//here certain ast.Nodes do not and won’t have their own case branch, e.g. *ast.IntegerLiteral
	//That’s because we can’t traverse their children, even if we wanted to, because they don’t have any
	//if they have children, as is the case with *ast.Program, we call ast.Modify with each child,
	//which again could result in calls to ast.Modify with the children of the child, and so on
	switch node := node.(type) {

	case *Program:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, modifier).(Statement)
		}

	case *ExpressionStatement:
		node.Expression, _ = Modify(node.Expression, modifier).(Expression)
	}

	return modifier(node)
}