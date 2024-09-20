package parser

import (
	"github.com/SanduCondorache/Interpreter/ast"
	"github.com/SanduCondorache/Interpreter/internals/lexer"
	"github.com/SanduCondorache/Interpreter/internals/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
