package main

type ASTNodeUnaryOperator struct {
	tk   *Token
	expr ASTNodeExpr
}

func (ast *AST) newUnaryOperator(t *Token) *ASTNodeUnaryOperator {
	return &ASTNodeUnaryOperator{
		tk:   t,
		expr: nil,
	}
}

type ASTNodeBinaryOperator struct {
	tk    *Token
	left  ASTNodeExpr
	right ASTNodeExpr
}

func (ast *AST) newBinaryOperator(t *Token) *ASTNodeBinaryOperator {
	return &ASTNodeBinaryOperator{
		tk:    t,
		left:  nil,
		right: nil,
	}
}

type ASTNodeFnCall struct {
	tk   *Token
	expr ASTNodeExpr
	args *ASTNodeFnCallArgs
}

func (ast *AST) newFnCall(t *Token) *ASTNodeFnCall {
	// t.Type =
	return &ASTNodeFnCall{
		tk:   t,
		expr: nil,
		args: nil,
	}
}

type ASTNodeSubscript struct {
	tk    *Token
	expr  ASTNodeExpr
	index ASTNodeExpr
}

func (ast *AST) newSubscript(t *Token) *ASTNodeSubscript {
	// t.Type =
	return &ASTNodeSubscript{
		tk:    t,
		expr:  nil,
		index: nil,
	}
}
