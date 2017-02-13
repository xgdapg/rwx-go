package main

type ASTNodeReturn struct {
	expr *ASTNodeExpr
}

func (ast *AST) parseReturn() (*ASTNodeReturn, error) {
	n := &ASTNodeReturn{}

	ast.index += 1 //return

	if n.condExpr, err = ast.parseExpr(); err != nil {
		return nil, err
	}

	return n, nil
}
