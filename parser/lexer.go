package parser

import (
	"errors"
	"fmt"
	"github.com/mattn/anko/ast"
)

const (
	EOF        = -1
	ParseError = 0
)

type Token struct {
	tok int
	lit string
	pos ast.Position
}

type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

var opName = map[string]int{
	"var":    VAR,
	"func":   FUNC,
	"return": RETURN,
	"throw":  THROW,
	"if":     IF,
	"for":    FOR,
	"in":     IN,
	"else":   ELSE,
	"new":    NEW,
	"true":   TRUE,
	"false":  FALSE,
	"nil":    NIL,
	"module": MODULE,
}

func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

func (s *Scanner) Scan() (tok int, lit string, pos ast.Position) {
	var err error
retry:
	s.skipBlank()
	pos = s.pos()
	switch ch := s.peek(); {
	case isLetter(ch):
		lit, err = s.scanIdentifier()
		if err != nil {
			tok = ParseError
		}
		if name, ok := opName[lit]; ok {
			tok = name
		} else {
			tok = IDENT
		}
	case isDigit(ch):
		tok = NUMBER
		lit, err = s.scanNumber()
		if err != nil {
			tok = ParseError
		}
	case ch == '"':
		tok = STRING
		lit, err = s.scanString()
		if err != nil {
			tok = ParseError
		}
	case ch == '`':
		tok = STRING
		lit, err = s.scanRawString()
		if err != nil {
			tok = ParseError
		}
	default:
		switch ch {
		case -1:
			tok = EOF
		case '#':
			for !isEOL(s.peek()) {
				s.next()
			}
			s.next()
			goto retry
		case '!':
			s.next()
			if s.peek() == '=' {
				tok = NE
			} else {
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '=':
			s.next()
			if s.peek() == '=' {
				tok = EQ
			} else {
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '>':
			s.next()
			tok = int(ch)
			if s.peek() == '=' {
				tok = GE
			} else {
				s.back()
			}
		case '<':
			s.next()
			tok = int(ch)
			if s.peek() == '=' {
				tok = LE
			} else {
				s.back()
			}
		case '|':
			s.next()
			tok = int(ch)
			if s.peek() == '|' {
				tok = OR
			} else {
				s.back()
			}
		case '&':
			s.next()
			tok = int(ch)
			if s.peek() == '&' {
				tok = AND
			} else {
				s.back()
			}
		case '.':
			tok = int(ch)
			lit = string(ch)
			s.next()
			if s.peek() == '.' {
				s.next()
				tok = ParseError
				if s.peek() == '.' {
					tok = VARARG
				} else {
					s.back()
				}
			} else {
				s.back()
			}
		case '(', ')', ':', ';', '+', '-', '*', '/', '%', '{', '}', ',', '[', ']', '\n':
			tok = int(ch)
			lit = string(ch)
		default:
			tok = ParseError
		}
		s.next()
	}
	return
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isEOL(ch rune) bool {
	return ch == '\n' || ch == -1
}

func isBrank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (s *Scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	} else {
		return -1
	}
}

func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

func (s *Scanner) back() {
	s.offset--
}

func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

func (s *Scanner) pos() ast.Position {
	return ast.Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *Scanner) skipBlank() {
	for isBrank(s.peek()) {
		s.next()
	}
}

func (s *Scanner) scanIdentifier() (string, error) {
	var ret []rune
	for {
		if s.peek() == ':' {
			s.next()
			if s.peek() != ':' {
				s.back()
				s.back()
				break
			}
			s.next()
			ret = append(ret, ':')
			ret = append(ret, ':')
			continue
		} else if !isLetter(s.peek()) && !isDigit(s.peek()) {
			break
		}
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

func (s *Scanner) scanNumber() (string, error) {
	var ret []rune
	for isDigit(s.peek()) || s.peek() == '.' || s.peek() == 'e' {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

func (s *Scanner) scanRawString() (string, error) {
	var ret []rune
	for {
		s.next()
		if s.peek() == EOF {
			return "", errors.New("Parser Error")
			break
		}
		if s.peek() == '`' {
			s.next()
			break
		}
		ret = append(ret, s.peek())
	}
	return string(ret), nil
}

func (s *Scanner) scanString() (string, error) {
	var ret []rune
	for {
		s.next()
		if s.peek() == EOF {
			return "", errors.New("Parser Error")
			break
		}
		if s.peek() == '\\' {
			s.next()
			switch s.peek() {
			case 'b':
				ret = append(ret, '\b')
				continue
			case 'f':
				ret = append(ret, '\f')
				continue
			case 'r':
				ret = append(ret, '\r')
				continue
			case 'n':
				ret = append(ret, '\n')
				continue
			case 't':
				ret = append(ret, '\t')
				continue
			}
			return "", errors.New("Parser Error")
		}
		if s.peek() == '"' {
			s.next()
			break
		}
		ret = append(ret, s.peek())
	}
	return string(ret), nil
}

type Lexer struct {
	s     *Scanner
	lit   string
	pos   ast.Position
	e     error
	stmts []ast.Stmt
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	if tok == EOF {
		return 0
	}
	lval.tok = Token{tok: tok, lit: lit, pos: pos}
	l.lit = lit
	l.pos = pos
	return tok
}

func (l *Lexer) Error(e string) {
	l.e = fmt.Errorf("Line %d, Column %d: %q %s", l.pos.Line, l.pos.Column, l.lit, e)
}

func Parse(s *Scanner) ([]ast.Stmt, error) {
	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		return nil, l.e
	}
	return l.stmts, nil
}
