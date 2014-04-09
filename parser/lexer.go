// Package parser implements parser for anko.
package parser

import (
	"errors"
	"fmt"
	"github.com/mattn/anko/ast"
)

const (
	EOF        = -1   // End of file.
	EOL        = '\n' // End of line.
)

// Token provide structure of identities or literals.
type Token struct {
	tok int
	lit string
	pos ast.Position
}

// Error provides a convenient interface for handling runtime error.
// It can be Error inteface with type cast which can call Pos().
type Error struct {
	Message  string
	Pos      ast.Position
	Filename string
	Fatal    bool
}

// Error return the error message.
func (e *Error) Error() string {
	return e.Message
}

// Scanner store informations for lexer.
type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

// opName is correction of operation names.
var opName = map[string]int{
	"func":     FUNC,
	"return":   RETURN,
	"var":      VAR,
	"throw":    THROW,
	"if":       IF,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"in":       IN,
	"else":     ELSE,
	"new":      NEW,
	"true":     TRUE,
	"false":    FALSE,
	"nil":      NIL,
	"module":   MODULE,
	"try":      TRY,
	"catch":    CATCH,
	"finally":  FINALLY,
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
}

// Init reset code to scan.
func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

// Scan analyses token, and decide identify or literals.
func (s *Scanner) Scan() (tok int, lit string, pos ast.Position, err error) {
retry:
	s.skipBlank()
	pos = s.pos()
	switch ch := s.peek(); {
	case isLetter(ch):
		lit, err = s.scanIdentifier()
		if err != nil {
			return
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
			return
		}
	case ch == '"':
		tok = STRING
		lit, err = s.scanString('"')
		if err != nil {
			return
		}
	case ch == '\'':
		tok = STRING
		lit, err = s.scanString('\'')
		if err != nil {
			return
		}
	case ch == '`':
		tok = STRING
		lit, err = s.scanRawString()
		if err != nil {
			return
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
			switch s.peek() {
			case '=':
				tok = NEQ
				lit = "!="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '=':
			s.next()
			switch s.peek() {
			case '=':
				tok = EQEQ
				lit = "=="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '+':
			s.next()
			switch s.peek() {
			case '+':
				tok = PLUSPLUS
				lit = "++"
			case '=':
				tok = PLUSEQ
				lit = "+="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '-':
			s.next()
			switch s.peek() {
			case '-':
				tok = MINUSMINUS
				lit = "--"
			case '=':
				tok = MINUSEQ
				lit = "-="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '*':
			s.next()
			switch s.peek() {
			case '*':
				tok = POW
				lit = "**"
			case '=':
				tok = MULEQ
				lit = "*="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '/':
			s.next()
			switch s.peek() {
			case '=':
				tok = DIVEQ
				lit = "/="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '>':
			s.next()
			switch s.peek() {
			case '=':
				tok = GE
				lit = ">="
			case '>':
				tok = SHIFTRIGHT
				lit = ">>"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '<':
			s.next()
			switch s.peek() {
			case '=':
				tok = LE
				lit = "<="
			case '<':
				tok = SHIFTLEFT
				lit = "<<"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '|':
			s.next()
			switch s.peek() {
			case '|':
				tok = OROR
				lit = "||"
			case '=':
				tok = OREQ
				lit = "|="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '&':
			s.next()
			switch s.peek() {
			case '&':
				tok = ANDAND
				lit = "&&"
			case '=':
				tok = ANDEQ
				lit = "&="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '.':
			s.next()
			if s.peek() == '.' {
				s.next()
				if s.peek() == '.' {
					tok = VARARG
				} else {
					err = fmt.Errorf(`syntax error "%s"`, "..")
					return
				}
			} else {
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '(', ')', ':', ';', '%', '?', '{', '}', ',', '[', ']', '\n', '^':
			tok = int(ch)
			lit = string(ch)
		default:
			err = fmt.Errorf(`syntax error "%s"`, ch)
			return
		}
		s.next()
	}
	return
}

// isLetter return true if the rune is a letter for identity.
func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit return true if the rune is a number.
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// isHex return true if the rune is a hex digits.
func isHex(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

// isEOL return true if the rune is at end-of-line or end-of-file.
func isEOL(ch rune) bool {
	return ch == '\n' || ch == -1
}

// isBrank return true if the rune is empty character..
func isBrank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// peek return current rune in the code.
func (s *Scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	} else {
		return -1
	}
}

// next move offset to next.
func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

// current return the current offset.
func (s *Scanner) current() int {
	return s.offset
}

// offset set the offset value.
func (s *Scanner) set(o int) {
	s.offset = o
}

// back move back offset once to top.
func (s *Scanner) back() {
	s.offset--
}

// reachEOF return true if offset is at end-of-file.
func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

// pos return the position of current.
func (s *Scanner) pos() ast.Position {
	return ast.Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

// skipBlank move position into non-black character.
func (s *Scanner) skipBlank() {
	for isBrank(s.peek()) {
		s.next()
	}
}

// scanIdentifier return identifier begining at current position.
func (s *Scanner) scanIdentifier() (string, error) {
	var ret []rune
	for {
		if !isLetter(s.peek()) && !isDigit(s.peek()) {
			break
		}
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

// scanIdentifier return number begining at current position.
func (s *Scanner) scanNumber() (string, error) {
	var ret []rune
	ch := s.peek()
	ret = append(ret, ch)
	s.next()
	if ch == '0' && s.peek() == 'x' {
		ret = append(ret, s.peek())
		s.next()
		for isHex(s.peek()) {
			ret = append(ret, s.peek())
			s.next()
		}
	} else {
		for isDigit(s.peek()) || s.peek() == '.' || s.peek() == 'e' {
			ret = append(ret, s.peek())
			s.next()
		}
		if isLetter(s.peek()) {
			return "", errors.New("identifier starts immediately after numeric literal")
		}
	}
	return string(ret), nil
}

// scanIdentifier return raw-string begining at current position.
func (s *Scanner) scanRawString() (string, error) {
	var ret []rune
	for {
		s.next()
		if s.peek() == EOF {
			return "", errors.New("unexpected EOF")
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

// scanIdentifier return string begining at current position. This handle backslash escaping.
func (s *Scanner) scanString(l rune) (string, error) {
	var ret []rune
eos:
	for {
		s.next()
		switch s.peek() {
		case EOL:
			return "", errors.New("unexpected EOL")
		case EOF:
			return "", errors.New("unexpected EOF")
		case l:
			s.next()
			break eos
		case '\\':
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
			ret = append(ret, s.peek())
			continue
		default:
			ret = append(ret, s.peek())
		}
	}
	return string(ret), nil
}

// Lexer provide inteface to parse codes.
type Lexer struct {
	s     *Scanner
	lit   string
	pos   ast.Position
	e     error
	stmts []ast.Stmt
}

// Lex scan the token and literals.
func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos, err := l.s.Scan()
	if err != nil {
		l.e = &Error{Message: fmt.Sprintf("%s", err.Error()), Pos: pos, Fatal: true}
	}
	if tok == EOF {
		return 0
	}
	lval.tok = Token{tok: tok, lit: lit, pos: pos}
	l.lit = lit
	l.pos = pos
	return tok
}

// Error set parse error.
func (l *Lexer) Error(e string) {
	l.e = &Error{Message: e, Pos: l.pos, Fatal: false}
}

// Parser provide way to parse the code.
func Parse(s *Scanner) ([]ast.Stmt, error) {
	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		return nil, l.e
	}
	return l.stmts, l.e
}
