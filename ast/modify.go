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

	case *InfixExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Right, _ = Modify(node.Right, modifier).(Expression)

	case *PrefixExpression:
		node.Right, _ = Modify(node.Right, modifier).(Expression)

	case *IndexExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Index, _ = Modify(node.Index, modifier).(Expression)

	case *IfExpression:
		node.Condition, _ = Modify(node.Condition, modifier).(Expression)
		node.Consequence, _ = Modify(node.Consequence, modifier).(*BlockStatement)

		if node.Alternative != nil {
			node.Alternative, _ = Modify(node.Alternative, modifier).(*BlockStatement)
		}

	case *BlockStatement:
		for i, _ := range node.Statements {
			node.Statements[i], _ = Modify(node.Statements[i], modifier).(Statement)
		}

	case *ReturnStatement:
		node.ReturnValue, _ = Modify(node.ReturnValue, modifier).(Expression)

	case *LetStatement:
		node.Value, _ = Modify(node.Value, modifier).(Expression)

	}

	return modifier(node)
}
