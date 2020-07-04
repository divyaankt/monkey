package ast

import "monkey/token"

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

//If statement is let x = 5*5
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token //the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
