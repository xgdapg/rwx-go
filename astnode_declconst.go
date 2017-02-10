package main

type ASTNodeDeclConst struct {
	Token *Token
	Name  *ASTNodeIdentifier
	Type  *ASTNodeType
}

func (ast *AST) parseDeclConst() (*ASTNodeDeclConst, error) {
	n := &ASTNodeDeclConst{}

	ast.index += 1

	if n.Name, err = ast.parseIdentifier(); err != nil {
		return nil, err
	}

	if ast.tk(0).isOperatorV(":") {
		ast.index += 1
		if n.Type, err = ast.parseType(); err != nil {
			return nil, err
		}
	}

	if !ast.tk(0).isOperatorV("=") {
		return nil, newASTError("constant value required")
	}

	return n, nil
}
