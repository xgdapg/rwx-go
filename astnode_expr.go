package main

import (
	"errors"
)

func (ast *AST) parseBasicExpr() (ASTNodeExpr, error) {
	t := ast.tk(0)
	if t.isLiteral() {
		return ast.parseLiteral()
	}
	if t.isIdentifier() {
		return ast.parseIdentifier()
	}
	if t.isOperatorV("(") {
		ast.index += 1 //(
		n, err := ast.parseExpr()
		if err != nil {
			return nil, err
		}
		t = ast.tk(1)
		if !t.isOperatorV(")") {
			return nil, newASTError("expect `)`, got `" + t.Value + "`")
		}
		ast.index += 1 //)
		return n, nil
	}
	if t.isOperatorV("{") {
		return ast.parseBlock()
	}
	if t.isKeywordV("if") {
		return ast.parseIf()
	}
	if t.isKeywordV("fn") {
		return ast.parseUnnamedFunc()
	}
	return nil, nil
}

var errExprNotFound = errors.New("no expression be found")

func (ast *AST) parseExpr() (ASTNodeExpr, error) {
	sOperator, sExpression := 1, 2
	prev := sOperator
	list := []ASTNodeExpr{}
	for ast.tk(0) != EmptyToken {
		t := ast.tk(0)
		if prev == sOperator && t.isUnaryOperator() {
			list = append(list, ast.newUnaryOperator(t))
			prev = sOperator
			continue
		}
		if prev == sExpression && t.isBinaryOperator() {
			list = append(list, ast.newBinaryOperator(t))
			prev = sOperator
			continue
		}
		if prev == sExpression && t.isOperatorV("(") {
			list = append(list, ast.newFnCall(t))
			ast.index += 1 //(
			n, err := ast.parseFnCallArgs()
			if err != nil {
				return nil, err
			}
			if !ast.tk(0).isOperatorV(")") {
				return nil, newASTError("expect `)`, got `" + ast.tk(0).Value + "`")
			}
			ast.index += 1 //)
			prev = sOperator
			continue
		}
		if prev == sExpression && t.isOperatorV("[") {
			list = append(list, ast.newSubscript(t))
			ast.index += 1 //[
			n, err := ast.parseExpr()
			if err != nil {
				return nil, err
			}
			list = append(list, n)
			if !ast.tk(0).isOperatorV("]") {
				return nil, newASTError("expect `]`, got `" + ast.tk(0).Value + "`")
			}
			ast.index += 1 //]
			prev = sOperator
			continue
		}
		if prev == sOperator {
			n, err := ast.parseBasicExpr()
			if err != nil {
				return nil, err
			}
			list = append(list, n)
			continue
		}
		break
	}
	if len(list) == 0 {
		return nil, errExprNotFound
	}
	return ast.buildTree(list)
}

func (ast *AST) buildTree(list []ASTNodeExpr) ASTNodeExpr {
	opi := -1
	priority := 99
	for i, n := range list {
		t := n.token()
		if !t.isOperator() {
			continue
		}
		switch n.(type) {
		case *ASTNodeUnaryOperator:
			if p := t.getPriority(true); p < priority {
				opi = i
				priority = p
				continue
			}
		case *ASTNodeBinaryOperator, *ASTNodeFnCall, *ASTNodeSubscript:
			if p := t.getPriority(false); p <= priority {
				opi = i
				priority = p
				continue
			}
		}
	}
	if opi != -1 {
		n := list[opi]
		switch nn := n.(type) {
		case *ASTNodeUnaryOperator:
			nn.expr = ast.buildTree(list[opi+1:])
		case *ASTNodeBinaryOperator:
			nn.left = ast.buildTree(list[:opi])
			nn.right = ast.buildTree(list[opi+1:])
		case *ASTNodeFnCall:
			nn.expr = ast.buildTree(list[:opi])
			nn.args = ast.buildTree(list[opi+1:])
		case *ASTNodeSubscript:
			nn.expr = ast.buildTree(list[:opi])
			nn.index = ast.buildTree(list[opi+1:])
		}
		return n
	}
	return list[0]
}
