package main

type ASTNodeAssign struct {
	lvalue ASTNode
	rvalue *ASTNodeExpr
}

func (ast *AST) parseAssign() (*ASTNodeAssign, error) {
	n := &ASTNodeAssign{}

	ast.index += 1 //=

	block := ast.currBlock()
	if len(block.Nodes) == 0 {
		return nil, newASTError("lvalue not found")
	}

	n.lvalue = block.Nodes[len(block.Nodes)-1]
	if n.rvalue, err = ast.parseExpr(); err != nil {
		return nil, err
	}
	block.Nodes[len(block.Nodes)-1] = n

	return n, nil
}
