package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"unicode/utf8"
)

const (
	EOL byte = 255 // -1
	EOF byte = 255 // -1
)

type Lexer struct {
	Tokens []*Token

	file      string
	lines     []string
	row       int
	col       int
	tokenList []*Token
}

func NewLexer(path string) (*Lexer, error) {
	l := &Lexer{
		Tokens:    []*Token{},
		file:      path,
		row:       0,
		col:       0,
		tokenList: []*Token{},
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if !utf8.Valid(bytes) {
		return nil, errors.New("UTF-8 required")
	}

	l.lines = getLines(bytes)

	if err = l.parse(); err != nil {
		return nil, err
	}

	l.makeArray()

	return l, nil
}

func (this *Lexer) newError(err string) error {
	return errors.New("[" + strconv.Itoa(this.row+1) + "," + strconv.Itoa(this.col+1) + "]" + err)
}

func (this *Lexer) parse() error {
	lines := len(this.lines)
	for this.row < lines {
		c := this.getLineChar(0)
		if c == EOL {
			this.row++
			this.col = 0
			continue
		}

		s := []byte{}
		nc := this.getLineChar(1)

		if c == '"' {
			if err := this.fetchString(); err != nil {
				return err
			}
			continue
		}
		if c == '/' && nc == '/' {
			if err := this.fetchLineComment(); err != nil {
				return err
			}
			continue
		}
		if c == '/' && nc == '*' {
			if err := this.fetchBlockComment(); err != nil {
				return err
			}
			continue
		}
		if c == '*' && nc == '/' {
			return this.newError("unmatched block comment")
		}
		if c >= '0' && c <= '9' {
			if err := this.fetchNumber(); err != nil {
				return err
			}
			continue
		}
		if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
			this.movePos(1)
			continue
		}
		if c == '-' && nc == '>' || c == '&' && nc == '&' || c == '|' && nc == '|' {
			s = append(s, c, nc)
			this.addToken(string(s), KOperator, TUnknown)
			continue
		}
		if c == '+' || c == '-' || c == '*' || c == '/' || c == '%' || c == '!' || c == '=' || c == '>' || c == '<' {
			s = append(s, c)
			if nc == '=' {
				s = append(s, nc)
			}
			this.addToken(string(s), KOperator, TUnknown)
			continue
		}
		if c == '(' || c == ')' || c == '[' || c == ']' || c == '{' || c == '}' || c == ':' || c == ',' || c == ';' || c == '.' || c == '&' {
			s = append(s, c)
			this.addToken(string(s), KOperator, TUnknown)
			continue
		}
		i := 0
		for c == '_' || c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || int8(c) < 0 && c != EOL {
			s = append(s, c)
			i++
			c = this.getLineChar(i)
		}
		kind := KIdentifier
		if _, ok := keywords[string(s)]; ok {
			kind = KKeyword
		}
		this.addToken(string(s), kind, TUnknown)
	}
	return nil
}

func (this *Lexer) addToken(v string, k TokenKind, t TokenType) {
	token := NewToken(k, t, v, this.row, this.col)
	this.tokenList = append(this.tokenList, token)
	this.movePos(len(v))

	if k == KKeyword {
		switch v {
		case "var":
			token.Type = TDeclareVar
		case "const":
			token.Type = TDeclareConst
		case "if":
			token.Type = TIf
		case "else":
			token.Type = TElse
		case "while":
			token.Type = TWhile
		case "break":
			token.Type = TBreak
		case "continue":
			token.Type = TContinue
		case "true", "false":
			token.Kind = KLiteral
			token.Type = TBoolean
		}

	}
	if k == KOperator {
		switch v {
		case "+":
			token.Type = TPlus
		case "-":
			token.Type = TMinus
		case "*":
			token.Type = TMulti
		case "/":
			token.Type = TDivide
		case "%":
			token.Type = TModulus

		case "==":
			token.Type = TEqual
		case "<":
			token.Type = TLessThan
		case "<=":
			token.Type = TLessEqual
		case "!=":
			token.Type = TNotEqual
		case ">":
			token.Type = TGreaterThan
		case ">=":
			token.Type = TGreaterEqual

		case "&&":
			token.Type = TLogicAnd
		case "||":
			token.Type = TLogicOr
		case "!":
			token.Type = TLogicNot

		case "&":
			token.Type = TBinOpAnd
		case "|":
			token.Type = TBinOpOr
		case "^":
			token.Type = TBinOpXor
		case "~":
			token.Type = TBinOpNot
		case "<<":
			token.Type = TBinOpLShift
		case ">>":
			token.Type = TBinOpRShift

		case ".":
			token.Type = TDot
		case ",":
			token.Type = TComma
		case ";":
			token.Type = TSemicolon
		case ":":
			token.Type = TColon

		case "(":
			token.Type = TLParen
		case ")":
			token.Type = TRParen
		case "[":
			token.Type = TLBracket
		case "]":
			token.Type = TRBracket
		case "{":
			token.Type = TLBrace
		case "}":
			token.Type = TRBrace

		case "=":
			token.Type = TAssign
		}
	}
}

func (this *Lexer) movePos(offset int) {
	this.col += offset
	lines := len(this.lines)
	for this.row < lines {
		s := this.lines[this.row]
		sl := len(s)
		if this.col >= sl {
			this.col -= sl
			this.row++
			continue
		}
		break
	}
}

func (this *Lexer) getChar(offset int) byte {
	r := this.row
	c := this.col + offset
	lines := len(this.lines)
	for r < lines {
		s := this.lines[r]
		sl := len(s)
		if c >= sl {
			c -= sl
			r++
			continue
		}
		return s[c]
	}
	return EOF
}

func (this *Lexer) getLineChar(offset int) byte {
	c := this.col + offset
	s := this.lines[this.row]
	if c < len(s) {
		return s[c]
	}
	return EOL
}

func (this *Lexer) fetchString() error {
	this.movePos(1) // skip the first "
	c := this.getChar(0)
	i := 0
	escapeCnt := 0
	s := []byte{}
	for c != EOF {
		if c == '"' {
			this.addToken(string(s), KLiteral, TString)
			this.movePos(1 + escapeCnt)
			return nil
		} else if c == '\\' {
			i++
			nc := this.getChar(i)
			switch nc {
			case '\\':
				s = append(s, '\\')
			case 't':
				s = append(s, '\t')
			case 'r':
				s = append(s, '\r')
			case 'n':
				s = append(s, '\n')
			case '"':
				s = append(s, '"')
			default:
				return this.newError("unknown escape \\" + string(nc))
			}
			escapeCnt++
		} else {
			s = append(s, c)
		}
		i++
		c = this.getChar(i)
	}
	return this.newError("unclosed string")
}

func (this *Lexer) fetchLineComment() error {
	i := 0
	c := this.getLineChar(i)
	s := []byte{}
	for c != EOL && c != '\r' && c != '\n' {
		s = append(s, c)
		i++
		c = this.getLineChar(i)
	}
	this.addToken(string(s), KComment, TLineComment)
	this.row++
	this.col = 0
	return nil
}

func (this *Lexer) fetchBlockComment() error {
	clv := 0
	i := 0
	c := this.getLineChar(i)
	s := []byte{}
	for c != EOF {
		s = append(s, c)
		i++
		nc := this.getChar(i)
		if c == '/' && nc == '*' {
			s = append(s, nc)
			i++
			clv++
		}
		if c == '*' && nc == '/' {
			s = append(s, nc)
			i++
			clv--
		}
		if clv == 0 {
			this.addToken(string(s), KComment, TBlockComment)
			return nil
		}
		c = this.getChar(i)
	}
	return this.newError("unclosed block comment")
}

func (this *Lexer) fetchNumber() error {
	hasDot := false
	i := 0
	c := this.getLineChar(i)
	s := []byte{}
	for c != EOF {
		if c == '.' {
			if hasDot {
				break
			}
			nc := this.getChar(i + 1)
			if !(nc >= '0' && nc <= '9') {
				break
			}
			hasDot = true
			s = append(s, c)
		} else if c >= '0' && c <= '9' {
			s = append(s, c)
		} else {
			break
		}
		i++
		c = this.getChar(i)
	}
	t := TInteger
	if hasDot {
		t = TFloat
	}
	this.addToken(string(s), KLiteral, t)
	return nil
}

func (this *Lexer) makeArray() {
	for _, t := range this.tokenList {
		if t.isComment() {
			continue
		}
		t.Index = len(this.Tokens)
		t.Lex = this
		this.Tokens = append(this.Tokens, t)
	}
}
