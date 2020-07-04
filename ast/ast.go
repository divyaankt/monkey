package ast

//Node of the AST
type Node interface {
	TokenLiteral() string //Provides the literal value of the token of the Node
}

//Each Node can either be an expression or a statement
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

//This Node is the Root Node of every AST our parser produces
//Every valid program is a series of statements they are contained in the Program.Statements
//which is just a slice of AST nodes that implement the Statement interface
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
