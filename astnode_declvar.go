package main

type ASTNodeDeclVar struct {
	Token *Token
	Name  *ASTNodeIdentifier
	Type  *ASTNodeType
}

func (ast *AST) parseDeclVar() (*ASTNodeDeclVar, error) {
	n := &ASTNodeDeclVar{}

	ast.index += 1

	if n.Name, err = ast.parseIdentifier(); err != nil {
		return nil, err
	}

	t := ast.tk(0)
	if !t.isOperatorV(":") && !t.isOperatorV("=") {
		return nil, newASTError("cannot determine the type of variable `" + n.Name.Token.Value + "`")
	}

	if t.isOperatorV(":") {
		ast.index += 1
		if n.Type, err = ast.parseType(); err != nil {
			return nil, err
		}
	}

	return n, nil
}
