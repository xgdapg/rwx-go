package main

type ASTNodeIf struct {
	condExpr  ASTNodeExpr
	ifBlock   *ASTNodeBlock
	elseBlock *ASTNodeBlock
}

func (ast *AST) parseIf() (*ASTNodeIf, error) {
	n := &ASTNodeIf{}

	ast.index += 1 //if

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
			n.elseBlock = ast.newBlock()
			defer ast.popBlock()
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
