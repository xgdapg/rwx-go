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

	if !ast.tk(0).isOperatorV("{") {
		return nil, newASTError("expect `{`, got `" + ast.tk(0).Value + "`")
	}

	if n.ifBlock, err = ast.parseBlock(); err != nil {
		return nil, err
	}

	if ast.tk(0).isKeywordV("else") {
		ast.index += 1 //else
		if ast.tk(0).isKeywordV("if") {
			n.elseBlock = &ASTNodeBlock{
				Nodes: []ASTNode{},
			}
			bn, err := ast.parseIf()
			if err != nil {
				return nil, err
			}
			n.elseBlock.Nodes = append(n.elseBlock.Nodes, bn)
		} else if ast.tk(0).isOperatorV("{") {
			if n.elseBlock, err = ast.parseBlock(); err != nil {
				return nil, err
			}
		} else {
			return nil, newASTError("expect `{` or `if`, got `" + ast.tk(0).Value + "`")
		}
	}

	return n, nil
}
