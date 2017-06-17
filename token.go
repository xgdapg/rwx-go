package main

type TokenKind int
type TokenType int

const (
	KUnknown TokenKind = iota
	KKeyword
	KLiteral
	KIdentifier
	KOperator
	KComment
	KBlock
)

const (
	TUnknown TokenType = iota

	//Literal
	TInteger
	TFloat
	TString
	TBoolean

	//Identifier
	TVarType
	TVarName

	//Comment
	TLineComment
	TBlockComment

	//Keyword
	TDefineFunction
	TDefineStruct
	TDefineInterface

	TDeclareVar
	TDeclareConst

	TIf
	TElse

	TWhile
	TBreak
	TContinue

	//Operator
	TPlus       // +
	TMinus      // -
	TMulti      // *
	TDivide     // /
	TModulus    // %
	TUnaryMinus // -

	TAssign       // =
	TEqual        // ==
	TLessThan     // <
	TLessEqual    // <=
	TNotEqual     // !=
	TGreaterThan  // >
	TGreaterEqual // >=

	TLogicAnd // &&
	TLogicOr  // ||
	TLogicNot // !

	TBinOpAnd    // &
	TBinOpOr     // |
	TBinOpNot    // ~
	TBinOpXor    // ^
	TBinOpLShift // <<
	TBinOpRShift // >>

	TRef // &

	TTilde      // ~
	TDot        // .
	TComma      //
	TSemicolon  // ;
	TColon      // :
	TRArrow     // ->
	TLArrow     // <-
	TFatArrow   // =>
	TPound      // #
	TDollar     // $
	TQuestion   // ?
	TUnderscore // _

	TLParen   // (
	TRParen   // )
	TLBracket // [
	TRBracket // ]
	TLBrace   // {
	TRBrace   // }

	TFnCall     //
	TFnCallArgs //
	TFnDefArgs  //

	TSubscript //
	TTuple     //

	TField //

	TType //

	//tTypeBinding
)

type Token struct {
	Kind  TokenKind
	Type  TokenType
	Value string
	Row   int
	Col   int
	Index int
	Lex   *Lexer
}

func NewToken(k TokenKind, t TokenType, v string, r, c int) *Token {
	return &Token{
		Kind:  k,
		Type:  t,
		Value: v,
		Row:   r,
		Col:   c,
		Index: 0,
		Lex:   nil,
	}
}

var EmptyToken *Token = NewToken(KUnknown, TUnknown, "#", 0, 0)

func (t *Token) Next(offset int) *Token {
	if t.Lex != nil {
		i := t.Index + offset
		if i >= 0 && i < len(t.Lex.Tokens) {
			return t.Lex.Tokens[i]
		}
	}
	return EmptyToken
}

func (t *Token) isSameTo(tk *Token) bool {
	return t.Kind == tk.Kind && t.Type == tk.Type && t.Value == tk.Value
}

func (t *Token) isIdentifier() bool {
	return t.Kind == KIdentifier
}

func (t *Token) isLiteral() bool {
	return t.Kind == KLiteral
}

func (t *Token) isLiteralType(tt TokenType) bool {
	return t.isLiteral() && t.Type == tt
}

func (t *Token) isOperator() bool {
	return t.Kind == KOperator
}

func (t *Token) isOperatorT(tt TokenType) bool {
	return t.isOperator() && t.Type == tt
}

func (t *Token) isOperatorV(v string) bool {
	return t.isOperator() && t.Value == v
}

func (t *Token) isKeyword() bool {
	return t.Kind == KKeyword
}

func (t *Token) isKeywordT(tt TokenType) bool {
	return t.isKeyword() && t.Type == tt
}

func (t *Token) isKeywordV(v string) bool {
	return t.isKeyword() && t.Value == v
}

func (t *Token) isComment() bool {
	return t.Kind == KComment
}

func (t *Token) isAssignOperator() bool {
	if !t.isOperator() {
		return false
	}
	switch t.Value {
	case "=", "+=", "-=", "*=", "/=", "%=":
		return true
	}
	return false
}

func (t *Token) isBinaryOperator() bool {
	if !t.isOperator() {
		return false
	}
	// if t.isAssignOperator() {
	// 	return true
	// }
	switch t.Value {
	case "+", "-", "*", "/", "%", "==", "!=", "<", "<=", ">", ">=", "&&", "||", "&", "|", "^", "<<", ">>", ".":
		return true
	}
	return false
}

func (t *Token) isUnaryOperator() bool {
	if !t.isOperator() {
		return false
	}
	switch t.Value {
	case "!", "~", "-", "&":
		return true
	}
	return false
}

func (t *Token) getPriority(unary bool) (priority int) {
	if unary && t.isUnaryOperator() {
		goto UNARY
	}
	priority = 1
	// if t.isAssignOperator() {
	// 	return
	// }
	priority = 2
	if t.Value == "|" {
		return
	}
	priority = 3
	if t.Value == "&" {
		return
	}
	priority = 4
	if t.Value == "==" ||
		t.Value == "!=" ||
		t.Value == "<" ||
		t.Value == "<=" ||
		t.Value == ">" ||
		t.Value == ">=" {
		return
	}
	priority = 5
	if t.Value == "+" ||
		t.Value == "-" {
		return
	}
	priority = 6
	if t.Value == "*" ||
		t.Value == "/" ||
		t.Value == "%" {
		return
	}
UNARY:
	priority = 7
	if t.isUnaryOperator() {
		return
	}
	priority = 8
	if t.Value == "." ||
		t.Value == "fn" ||
		t.Value == "[]" {
		return
	}

	return 0
}
