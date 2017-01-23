package main

type ASTNodeBlock struct {
	nodes []ASTNode
}

func (ast *AST) parseBlock() (*ASTNodeBlock, error) {
	n := &ASTNodeBlock{
		nodes: []ASTNode{},
	}

	if ast.tk(0).isOperatorV("{") {
		ast.index += 1
	}

	for ast.tk(0) != EmptyToken {
		t := ast.tk(0)
		if t.isOperatorV("}") {
			ast.index += 1
			break
		}
		if t.isOperatorV(";") {
			ast.index += 1
			continue
		}
		if t.isKeywordV("var") {
			if nn, err := ast.parseDeclVar(); err == nil {
				n.nodes = append(n.nodes, nn)
			} else {
				return nil, err
			}
		}
		return nil, newASTError("unmatched statement, got `" + t.Value + "`")
	}
	return n, nil
}
