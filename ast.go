package main

import (
	"errors"
	"strconv"
)

type AST struct {
	lex   *Lexer
	index int
	nodes []ASTNode
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
		lex:   l,
		index: 0,
		nodes: []ASTNode{},
	}
}

func (ast *AST) tk(offset int) *Token {
	i := ast.index + offset
	if i >= 0 && i < len(ast.lex.Tokens) {
		return ast.lex.Tokens[i]
	}
	return EmptyToken
}

func (ast *AST) parse() {
	//auto root = NODE_P(BlockNode);
	//checkNode(root);
	//return root;
}

func newASTError(text string) error {
	t := ast.tk(0)
	return errors.New("[" + strconv.Itoa(t.Row) + "," + strconv.Itoa(t.Col) + "] " + text)
}
