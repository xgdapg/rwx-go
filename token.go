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
	TMulit      // *
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

func (this *Token) Next(offset int) *Token {
	if this.Lex != nil {
		i := this.Index + offset
		if i >= 0 && i < len(this.Lex.Tokens) {
			return this.Lex.Tokens[i]
		}
	}
	return EmptyToken
}

func (this *Token) IsSameTo(t *Token) bool {
	return this.Kind == t.Kind && this.Type == t.Type && this.Value == t.Value
}

func (this *Token) isIdentifier() bool {
	return this.Kind == KIdentifier
}

func (this *Token) isLiteral() bool {
	return this.Kind == KLiteral
}

func (this *Token) isLiteralType(t TokenType) bool {
	return this.isLiteral() && this.Type == t
}

func (this *Token) isOperator() bool {
	return this.Kind == KOperator
}

func (this *Token) isOperatorT(t TokenType) bool {
	return this.isOperator() && this.Type == t
}

func (this *Token) isOperatorV(v string) bool {
	return this.isOperator() && this.Value == v
}

func (this *Token) isKeyword() bool {
	return this.Kind == KKeyword
}

func (this *Token) isKeywordT(t TokenType) bool {
	return this.isKeyword() && this.Type == t
}

func (this *Token) isKeywordV(v string) bool {
	return this.isKeyword() && this.Value == v
}

func (this *Token) isComment() bool {
	return this.Kind == KComment
}
