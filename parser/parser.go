//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:23
type yySymType struct {
	yys                    int
	stmt_func              ast.Stmt
	stmt_if                ast.Stmt
	stmt_if_else           ast.Stmt
	stmt_for               ast.Stmt
	stmt_try_catch_finally ast.Stmt
	stmts                  []ast.Stmt
	stmt                   ast.Stmt
	teim                   ast.Expr
	expr                   ast.Expr
	tok                    Token
	idents                 []string
	exprs                  []ast.Expr
	pair                   ast.Expr
	pairs                  []ast.Expr
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VARARG = 57350
const FUNC = 57351
const RETURN = 57352
const THROW = 57353
const IF = 57354
const ELSE = 57355
const FOR = 57356
const IN = 57357
const EQ = 57358
const NE = 57359
const GE = 57360
const LE = 57361
const OR = 57362
const AND = 57363
const NEW = 57364
const TRUE = 57365
const FALSE = 57366
const NIL = 57367
const MODULE = 57368
const TRY = 57369
const CATCH = 57370
const FINALLY = 57371
const UNARY = 57372

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"FUNC",
	"RETURN",
	"THROW",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQ",
	"NE",
	"GE",
	"LE",
	"OR",
	"AND",
	"NEW",
	"TRUE",
	"FALSE",
	"NIL",
	"MODULE",
	"TRY",
	"CATCH",
	"FINALLY",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:434

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 26,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	36, 26,
	42, 26,
	-2, 1,
	-1, 3,
	36, 26,
	42, 26,
	-2, 1,
	-1, 4,
	36, 26,
	42, 26,
	-2, 1,
	-1, 5,
	36, 26,
	42, 26,
	-2, 1,
	-1, 6,
	36, 26,
	42, 26,
	-2, 1,
	-1, 7,
	36, 26,
	42, 26,
	-2, 1,
	-1, 18,
	42, 27,
	-2, 33,
	-1, 63,
	38, 1,
	-2, 26,
	-1, 94,
	38, 1,
	-2, 26,
	-1, 128,
	38, 1,
	-2, 26,
	-1, 130,
	38, 1,
	-2, 26,
	-1, 131,
	38, 1,
	-2, 26,
	-1, 134,
	38, 1,
	-2, 26,
	-1, 137,
	38, 1,
	-2, 26,
	-1, 141,
	38, 1,
	-2, 26,
	-1, 143,
	38, 1,
	-2, 26,
	-1, 157,
	38, 1,
	-2, 26,
	-1, 158,
	38, 1,
	-2, 26,
	-1, 162,
	38, 1,
	-2, 26,
	-1, 168,
	38, 1,
	-2, 26,
}

const yyNprod = 65
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 406

var yyAct = []int{

	1, 8, 70, 28, 29, 30, 31, 32, 33, 103,
	9, 103, 54, 57, 104, 133, 126, 103, 53, 107,
	67, 66, 120, 68, 103, 72, 115, 110, 53, 103,
	106, 52, 65, 154, 105, 64, 68, 53, 76, 77,
	78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 91, 68, 74, 42, 43, 45, 47,
	49, 51, 135, 98, 100, 129, 68, 102, 141, 59,
	140, 96, 109, 92, 95, 60, 61, 170, 168, 34,
	167, 164, 163, 35, 36, 101, 44, 46, 48, 50,
	156, 155, 152, 150, 147, 112, 146, 144, 125, 119,
	162, 118, 132, 158, 60, 121, 113, 157, 122, 123,
	143, 68, 137, 134, 130, 128, 94, 63, 166, 159,
	99, 153, 127, 116, 71, 148, 114, 97, 93, 136,
	124, 138, 139, 75, 73, 142, 62, 58, 145, 69,
	7, 6, 149, 5, 151, 4, 3, 2, 0, 0,
	42, 43, 45, 47, 49, 51, 0, 0, 160, 161,
	0, 0, 0, 165, 37, 38, 39, 40, 41, 169,
	0, 131, 0, 34, 0, 0, 0, 35, 36, 0,
	44, 46, 48, 50, 42, 43, 45, 47, 49, 51,
	0, 0, 0, 0, 0, 0, 0, 0, 37, 38,
	39, 40, 41, 0, 0, 0, 0, 34, 117, 0,
	0, 35, 36, 0, 44, 46, 48, 50, 42, 43,
	45, 47, 49, 51, 0, 0, 0, 0, 0, 0,
	0, 0, 37, 38, 39, 40, 41, 0, 0, 0,
	0, 34, 0, 0, 0, 35, 36, 111, 44, 46,
	48, 50, 42, 43, 45, 47, 49, 51, 0, 0,
	0, 0, 0, 0, 0, 0, 37, 38, 39, 40,
	41, 0, 0, 0, 0, 34, 108, 0, 0, 35,
	36, 0, 44, 46, 48, 50, 42, 43, 45, 47,
	49, 51, 0, 0, 0, 0, 0, 0, 0, 0,
	37, 38, 39, 40, 41, 0, 0, 0, 0, 34,
	0, 0, 0, 35, 36, 0, 44, 46, 48, 50,
	42, 43, 45, 47, 49, 51, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 39, 40, 41, 0,
	0, 0, 0, 34, 0, 0, 0, 35, 36, 0,
	44, 46, 48, 50, 18, 17, 20, 0, 0, 13,
	10, 11, 14, 0, 15, 55, 17, 20, 0, 0,
	56, 0, 24, 25, 26, 27, 12, 16, 0, 0,
	0, 19, 0, 24, 25, 26, 27, 22, 0, 23,
	0, 0, 19, 0, 21, 0, 0, 0, 22, 0,
	23, 0, 0, 0, 0, 21,
}
var yyPact = []int{

	350, -1000, 350, 350, 350, 350, 350, 350, 270, -5,
	361, 361, 133, 65, 37, 132, 80, -1000, -4, 361,
	-1000, 361, 118, 361, 130, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 361, 129, 361, 361, 361, 361,
	361, 361, 361, 361, 361, 361, 361, 361, 361, 361,
	361, 361, 361, 124, 270, -4, 36, 270, 79, 35,
	123, 361, 105, 350, 361, 361, 40, -31, 270, -8,
	-1000, -22, 236, 33, -13, -1000, 202, 304, 304, 40,
	40, 40, 270, 270, 270, 270, 270, 270, 270, 270,
	270, 270, -33, -1000, 350, 122, -14, 115, 168, 361,
	61, -18, 270, 361, -1000, 118, -1000, 361, -1000, 361,
	-1000, -1000, 60, -24, 114, 78, 25, 77, 134, 74,
	-1000, 270, -1000, 270, -25, -1000, 76, 22, 350, 75,
	350, 350, 31, -1000, 350, 73, 59, 350, 58, 56,
	121, 350, 55, 350, -1000, 54, 108, -1000, -7, 53,
	-1000, 52, -1000, 70, 66, 90, -1000, 350, 350, 63,
	44, 43, 350, -1000, 89, 42, 41, -1000, 350, 39,
	-1000,
}
var yyPgo = []int{

	0, 0, 147, 146, 145, 143, 141, 140, 1, 20,
	2, 139, 10,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 2, 2, 3, 3, 5, 4, 6, 7, 7,
	7, 7, 10, 11, 11, 11, 12, 12, 12, 9,
	9, 9, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 2, 1, 3,
	2, 2, 5, 8, 9, 11, 7, 7, 15, 12,
	11, 8, 3, 0, 1, 3, 0, 1, 3, 0,
	1, 3, 1, 1, 2, 1, 7, 4, 3, 4,
	3, 8, 3, 4, 3, 5, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -12,
	10, 11, 26, 9, 12, 14, 27, 5, 4, 31,
	6, 44, 37, 39, 22, 23, 24, 25, -1, -1,
	-1, -1, -1, -1, 39, 43, 44, 30, 31, 32,
	33, 34, 16, 17, 46, 18, 47, 19, 48, 20,
	49, 21, 36, 42, -8, 4, 9, -8, 4, 4,
	39, 39, 4, 37, 39, 36, -8, -9, -8, -11,
	-10, 6, -8, 4, -9, 4, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -9, 4, 37, 39, -12, 4, -8, 15,
	-1, -9, -8, 42, 45, 42, 38, 41, 40, 39,
	40, 45, -1, -12, 4, 40, 8, 40, -8, 38,
	40, -8, -10, -8, -9, 38, 40, 8, 37, 40,
	37, 37, 28, 40, 37, 40, -1, 37, -1, -1,
	39, 37, -1, 37, 38, -1, 38, 38, 4, -1,
	38, -1, 38, 13, 40, 38, 38, 37, 37, 29,
	-1, -1, 37, 38, 38, -1, 29, 38, 37, -1,
	38,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, 8, 0,
	0, 0, 0, 0, 0, 0, 0, 32, -2, 0,
	35, 29, 23, 0, 0, 46, 47, 48, 2, 3,
	4, 5, 6, 7, 29, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 29, 0, 10, 33, 0, 11, 0, 0,
	26, 0, 0, -2, 29, 0, 34, 0, 30, 0,
	24, 0, 0, 0, 0, 38, 0, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 61, 62,
	63, 64, 9, 28, -2, 26, 0, 27, 0, 0,
	0, 0, 60, 0, 40, 0, 42, 0, 44, 29,
	37, 43, 0, 0, 27, 0, 0, 0, 0, 0,
	39, 31, 25, 22, 0, 12, 0, 0, -2, 0,
	-2, -2, 0, 45, -2, 0, 0, -2, 0, 0,
	0, -2, 0, -2, 36, 0, 16, 17, 0, 0,
	13, 0, 41, 0, 0, 21, 14, -2, -2, 0,
	0, 0, -2, 15, 20, 0, 0, 19, -2, 0,
	18,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 34, 49, 3,
	39, 40, 32, 30, 42, 31, 43, 33, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 41, 3,
	47, 36, 46, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 44, 3, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 37, 48, 38,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 35,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:49
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:56
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:63
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:70
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:77
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if_else}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:84
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:91
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:99
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 9:
		//line parser.go.y:106
		{
			yyVAL.stmt = &ast.LetStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.stmt.SetPos(l.pos)
			}
		}
	case 10:
		//line parser.go.y:113
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:120
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 12:
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 13:
		//line parser.go.y:132
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:136
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 15:
		//line parser.go.y:141
		{
			yyVAL.stmt_if_else = &ast.IfStmt{If: yyS[yypt-8].expr, Then: yyS[yypt-5].stmts, Else: yyS[yypt-1].stmts}
		}
	case 16:
		//line parser.go.y:146
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 17:
		//line parser.go.y:151
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:156
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:160
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:164
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:168
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:173
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 23:
		//line parser.go.y:178
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 24:
		//line parser.go.y:182
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 25:
		//line parser.go.y:186
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 26:
		//line parser.go.y:191
		{
			yyVAL.idents = []string{}
		}
	case 27:
		//line parser.go.y:195
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 28:
		//line parser.go.y:199
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 29:
		//line parser.go.y:204
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 30:
		//line parser.go.y:208
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 31:
		//line parser.go.y:212
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 32:
		//line parser.go.y:217
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 33:
		//line parser.go.y:221
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 34:
		//line parser.go.y:225
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 35:
		//line parser.go.y:229
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 36:
		//line parser.go.y:233
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 37:
		//line parser.go.y:240
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 38:
		//line parser.go.y:247
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 39:
		//line parser.go.y:254
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 40:
		//line parser.go.y:261
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 41:
		//line parser.go.y:268
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 42:
		//line parser.go.y:275
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 43:
		//line parser.go.y:286
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 44:
		//line parser.go.y:293
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 45:
		//line parser.go.y:300
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 46:
		//line parser.go.y:307
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 47:
		//line parser.go.y:314
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 48:
		//line parser.go.y:318
		{
			yyVAL.expr = &ast.ConstExpr{Value: nil}
		}
	case 49:
		//line parser.go.y:322
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 50:
		//line parser.go.y:329
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 51:
		//line parser.go.y:336
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 52:
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:350
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:357
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:364
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:378
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:385
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 59:
		//line parser.go.y:392
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 60:
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 61:
		//line parser.go.y:406
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 62:
		//line parser.go.y:413
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:420
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 64:
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
