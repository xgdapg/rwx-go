package main

type ASTNodeBlock struct {
	outerBlock *ASTNodeBlock
	Nodes      []ASTNode
}

func (ast *AST) newBlock() *ASTNodeBlock {
	n := &ASTNodeBlock{
		outerBlock: ast.currBlock(),
		Nodes:      []ASTNode{},
	}
	ast.blockStack = append(ast.blockStack, n)
	return n
}

func (ast *AST) popBlock() {
	ast.blockStack = ast.blockStack[0 : len(ast.blockStack)-1]
}

func (ast *AST) currBlock() *ASTNodeBlock {
	l := len(ast.blockStack)
	if l == 0 {
		return nil
	}
	return ast.blockStack[l-1]
}

func (ast *AST) parseBlock() (*ASTNodeBlock, error) {
	n := ast.newBlock()
	defer ast.popBlock()

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
		if t.isAssignOperator() {
			if nn, err := ast.parseExpr(); err == nil {

			}
		}
		return nil, newASTError("unmatched statement, got `" + t.Value + "`")
	}
	return n, nil
}
