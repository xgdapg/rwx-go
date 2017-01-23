package main

type ASTNodeDeclVar struct {
	token   *Token
	varName *ASTNodeIdentifier
	varType *ASTNodeType
}

func (ast *AST) parseDeclVar(inLoop bool) (*ASTNodeDeclVar, error) {
	n := &ASTNodeDeclVar{}

	ast.index += 1

	if n.varName, err = ast.parseIdentifier(); err != nil {
		return nil, err
	}

	t := ast.tk(0)
	if !t.isOperatorV(":") && !(!inLoop && t.isOperatorV("=")) && !(inLoop && t.isKeywordV("in")) {
		newASTError("cannot determine the type of variable `" + n.varName.token.Value + "`")
	}

	if t.isOperatorV(":") {
		ast.index += 1
		if n.varType, err = ast.parseType(); err != nil {
			return nil, err
		}
	}

	return n, nil
}
