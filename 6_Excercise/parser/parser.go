package parser

import (
	"github.com/pioda/KP_Go/6_Excercise/lexer"
	"github.com/pioda/KP_Go/6_Excercise/node"
)

type Parser struct{
	Lex lexer.Lexer
	rootNode node.Node
	token string
}

func (p *Parser) Parse() node.Node{
	p.rootNode = nil
	p.expression()
	return p.rootNode
}

// 	<expression> ::= <term> { <or> <term> }
func (p *Parser) expression() {
	p.term()
	if p.token == "|" {
		lhs := p.rootNode
		p.term()
		rhs := p.rootNode
		p.rootNode = node.Or{
			LeftExpr:  lhs,
			RightExpr: rhs,
		}
	}
}

// 	<term> ::= <factor> { <and> <factor> }
func (p *Parser) term() {
	p.factor()
	if p.token == "&" {
		lhs := p.rootNode
		p.factor()
		rhs := p.rootNode
		p.rootNode = node.And{
			LeftExpr:  lhs,
			RightExpr: rhs,
		}
	}
}
// 	<factor> ::= <var> | <not> <factor> | (<expression>)
func (p *Parser) factor(){
	p.token = p.lexer.NextToken()
	if p.token == "" {
		return
	}
	if p.token == "!"{
		p.rootNode = &node.Not{Expr : p.rootNode}
	}
	else if p.token == "("{
		p.expression()
		p.token = p.lexer.NextToken()
	}
	else {
		p.rootNode = node.Var{Name:p.token}
		p.token = p.lexer.NextToken()
	}
}