package main

type ASTNodeBlock struct {
	Nodes []ASTNode
}

func (ast *AST) parseBlock() (*ASTNodeBlock, error) {
	n := &ASTNodeBlock{
		Nodes: []ASTNode{},
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
				n.Nodes = append(n.Nodes, nn)
			} else {
				return nil, err
			}
			continue
		}
		if t.isKeywordV("const") {
			if nn, err := ast.parseDeclConst(); err == nil {
				n.Nodes = append(n.Nodes, nn)
			} else {
				return nil, err
			}
			continue
		}
		if t.isKeywordV("if") {
			if nn, err := ast.parseIf(); err == nil {
				n.Nodes = append(n.Nodes, nn)
			} else {
				return nil, err
			}
			continue
		}
		if t.isKeywordV("while") {
			if nn, err := ast.parseWhile(); err == nil {
				n.Nodes = append(n.Nodes, nn)
			} else {
				return nil, err
			}
			continue
		}
		return nil, newASTError("unmatched statement, got `" + t.Value + "`")
	}
	return n, nil
}
