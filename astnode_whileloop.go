package main

type ASTNodeWhile struct {
	condExpr *ASTNodeExpr
	block    *ASTNodeBlock
}

func (ast *AST) parseWhile() (*ASTNodeWhile, error) {
	n := &ASTNodeWhile{}

	ast.index += 1 //while

	if n.condExpr, err = ast.parseExpr(); err != nil {
		return nil, err
	}

	if !ast.tk(0).isOperatorV("{") {
		return nil, newASTError("expect `{`, got `" + ast.tk(0).Value + "`")
	}

	if n.block, err = ast.parseBlock(); err != nil {
		return nil, err
	}

	return n, nil
}
