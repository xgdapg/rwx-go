package main

import (
	"errors"
	"strconv"
)

type AST struct {
	lex        *Lexer
	index      int
	blockStack []*ASTNodeBlock
}

type ASTNode interface {
	token() *Token
	typeName() string

	parse() *ASTNode
	eval() *ASTNode
	print()
}

func NewAST(l *Lexer) *AST {
	return &AST{
		lex:        l,
		index:      0,
		blockStack: []*ASTNodeBlock{},
	}
}

func (ast *AST) tk(offset int) *Token {
	i := ast.index + offset
	if i >= 0 && i < len(ast.lex.Tokens) {
		return ast.lex.Tokens[i]
	}
	return EmptyToken
}

func (ast *AST) parse() (ASTNode, error) {
	return ast.parseBlock()
}

func newASTError(text string) error {
	t := ast.tk(0)
	return errors.New("[" + strconv.Itoa(t.Row) + "," + strconv.Itoa(t.Col) + "] " + text)
}
