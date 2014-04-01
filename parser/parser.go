//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:22
type yySymType struct {
	yys                    int
	stmt_func              ast.Stmt
	stmt_if                ast.Stmt
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

//line parser.go.y:432

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
	-1, 17,
	42, 27,
	-2, 33,
	-1, 62,
	38, 1,
	-2, 26,
	-1, 73,
	38, 1,
	-2, 26,
	-1, 95,
	38, 1,
	-2, 26,
	-1, 133,
	38, 1,
	-2, 26,
	-1, 135,
	38, 1,
	-2, 26,
	-1, 136,
	38, 1,
	-2, 26,
	-1, 140,
	38, 1,
	-2, 26,
	-1, 143,
	38, 1,
	-2, 26,
	-1, 147,
	38, 1,
	-2, 26,
	-1, 148,
	38, 1,
	-2, 26,
	-1, 150,
	38, 1,
	-2, 26,
	-1, 165,
	38, 1,
	-2, 26,
	-1, 168,
	38, 1,
	-2, 26,
	-1, 173,
	38, 1,
	-2, 26,
}

const yyNprod = 65
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 478

var yyAct = []int{

	1, 7, 69, 27, 28, 29, 31, 32, 8, 104,
	104, 53, 56, 105, 138, 131, 104, 52, 108, 161,
	65, 66, 67, 123, 71, 104, 118, 113, 52, 104,
	107, 147, 112, 146, 106, 67, 141, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
	90, 91, 92, 67, 134, 75, 41, 42, 44, 46,
	48, 50, 99, 101, 51, 67, 103, 175, 97, 64,
	52, 110, 63, 93, 111, 96, 58, 59, 60, 33,
	172, 169, 164, 34, 35, 102, 43, 45, 47, 49,
	163, 162, 160, 158, 154, 153, 115, 151, 54, 16,
	19, 130, 121, 55, 128, 116, 124, 122, 173, 125,
	126, 59, 67, 168, 129, 74, 23, 24, 25, 26,
	165, 150, 148, 143, 140, 18, 135, 133, 95, 62,
	171, 21, 127, 22, 142, 166, 144, 145, 20, 137,
	73, 149, 100, 132, 152, 119, 70, 155, 156, 157,
	117, 159, 98, 94, 76, 72, 61, 57, 68, 6,
	5, 4, 3, 2, 0, 0, 167, 0, 0, 170,
	0, 0, 0, 0, 174, 41, 42, 44, 46, 48,
	50, 0, 0, 0, 0, 0, 0, 0, 0, 36,
	37, 38, 39, 40, 0, 0, 0, 0, 33, 139,
	0, 0, 34, 35, 0, 43, 45, 47, 49, 41,
	42, 44, 46, 48, 50, 0, 0, 0, 0, 0,
	0, 0, 0, 36, 37, 38, 39, 40, 0, 0,
	136, 0, 33, 0, 0, 0, 34, 35, 0, 43,
	45, 47, 49, 41, 42, 44, 46, 48, 50, 0,
	0, 0, 0, 0, 0, 0, 0, 36, 37, 38,
	39, 40, 0, 0, 0, 0, 33, 120, 0, 0,
	34, 35, 0, 43, 45, 47, 49, 41, 42, 44,
	46, 48, 50, 0, 0, 0, 0, 0, 0, 0,
	0, 36, 37, 38, 39, 40, 0, 0, 0, 0,
	33, 0, 0, 0, 34, 35, 114, 43, 45, 47,
	49, 41, 42, 44, 46, 48, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 36, 37, 38, 39, 40,
	0, 0, 0, 0, 33, 109, 0, 0, 34, 35,
	0, 43, 45, 47, 49, 41, 42, 44, 46, 48,
	50, 0, 0, 0, 0, 0, 0, 0, 0, 36,
	37, 38, 39, 40, 0, 0, 0, 0, 33, 0,
	0, 0, 34, 35, 0, 43, 45, 47, 49, 41,
	42, 44, 46, 48, 50, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 38, 39, 40, 0, 0,
	0, 0, 33, 0, 0, 0, 34, 35, 0, 43,
	45, 47, 49, 17, 16, 19, 0, 0, 12, 9,
	10, 13, 30, 14, 0, 0, 0, 0, 0, 0,
	0, 23, 24, 25, 26, 11, 15, 17, 16, 19,
	18, 0, 12, 9, 10, 13, 21, 14, 22, 0,
	0, 0, 0, 20, 0, 23, 24, 25, 26, 11,
	15, 0, 0, 0, 18, 0, 0, 0, 0, 0,
	21, 0, 22, 0, 0, 0, 0, 20,
}
var yyPact = []int{

	433, -1000, 433, 433, 409, 433, 433, 329, 28, 94,
	94, 153, 72, 39, 152, 92, -1000, 33, 94, -1000,
	94, 140, 94, 151, -1000, -1000, -1000, -1000, -1000, -1000,
	103, -1000, -1000, 94, 150, 94, 94, 94, 94, 94,
	94, 94, 94, 94, 94, 94, 94, 94, 94, 94,
	94, 94, 149, 329, 33, 38, 329, 91, 36, 148,
	94, 127, 433, 94, 94, 40, -32, 329, -8, -1000,
	-23, 295, 32, 433, -7, -13, -1000, 261, 363, 363,
	40, 40, 40, 329, 329, 329, 329, 329, 329, 329,
	329, 329, 329, -33, -1000, 433, 146, -14, 137, 227,
	94, 69, -17, 329, 94, -1000, 140, -1000, 94, -1000,
	94, 66, 94, -1000, -1000, 63, -25, 135, 90, 14,
	89, 193, 111, -1000, 329, -1000, 329, -26, -1000, 159,
	-1000, 87, -4, 433, 86, 433, 433, -6, -1000, 85,
	433, 84, 59, 433, 57, 56, 143, 433, 433, 55,
	433, -1000, 54, -1000, -1000, -21, 53, 52, -1000, 44,
	-1000, 83, 106, -1000, -1000, 433, 76, 43, 433, 101,
	42, 71, -1000, 433, 29, -1000,
}
var yyPgo = []int{

	0, 0, 163, 162, 161, 160, 159, 1, 21, 2,
	158, 8,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	2, 2, 3, 3, 4, 4, 4, 5, 6, 6,
	6, 6, 9, 10, 10, 10, 11, 11, 11, 8,
	8, 8, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 1, 3, 2,
	2, 5, 8, 9, 7, 5, 9, 7, 15, 12,
	11, 8, 3, 0, 1, 3, 0, 1, 3, 0,
	1, 3, 1, 1, 2, 1, 7, 4, 3, 4,
	3, 8, 3, 4, 3, 5, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -11, 10,
	11, 26, 9, 12, 14, 27, 5, 4, 31, 6,
	44, 37, 39, 22, 23, 24, 25, -1, -1, -1,
	13, -1, -1, 39, 43, 44, 30, 31, 32, 33,
	34, 16, 17, 46, 18, 47, 19, 48, 20, 49,
	21, 36, 42, -7, 4, 9, -7, 4, 4, 39,
	39, 4, 37, 39, 36, -7, -8, -7, -10, -9,
	6, -7, 4, 37, 12, -8, 4, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -8, 4, 37, 39, -11, 4, -7,
	15, -1, -8, -7, 42, 45, 42, 38, 41, 40,
	39, -1, 39, 40, 45, -1, -11, 4, 40, 8,
	40, -7, 38, 40, -7, -9, -7, -8, 38, -7,
	38, 40, 8, 37, 40, 37, 37, 28, 40, 40,
	37, 40, -1, 37, -1, -1, 39, 37, 37, -1,
	37, 38, -1, 38, 38, 4, -1, -1, 38, -1,
	38, 40, 38, 38, 38, 37, 29, -1, 37, 38,
	-1, 29, 38, 37, -1, 38,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 32, -2, 0, 35,
	29, 23, 0, 0, 46, 47, 48, 2, 3, 4,
	0, 5, 6, 29, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 29, 0, 9, 33, 0, 10, 0, 0, 26,
	0, 0, -2, 29, 0, 34, 0, 30, 0, 24,
	0, 0, 0, -2, 0, 0, 38, 0, 49, 50,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 61,
	62, 63, 64, 8, 28, -2, 26, 0, 27, 0,
	0, 0, 0, 60, 0, 40, 0, 42, 0, 44,
	29, 0, 0, 37, 43, 0, 0, 27, 0, 0,
	0, 0, 0, 39, 31, 25, 22, 0, 15, 0,
	11, 0, 0, -2, 0, -2, -2, 0, 45, 0,
	-2, 0, 0, -2, 0, 0, 0, -2, -2, 0,
	-2, 36, 0, 14, 17, 0, 0, 0, 12, 0,
	41, 0, 21, 16, 13, -2, 0, 0, -2, 20,
	0, 0, 19, -2, 0, 18,
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
		//line parser.go.y:47
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:54
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:61
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:68
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:75
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:82
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:90
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 8:
		//line parser.go.y:97
		{
			yyVAL.stmt = &ast.LetStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.stmt.SetPos(l.pos)
			}
		}
	case 9:
		//line parser.go.y:104
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 10:
		//line parser.go.y:111
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:118
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 12:
		//line parser.go.y:123
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 13:
		//line parser.go.y:127
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 14:
		//line parser.go.y:132
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 15:
		//line parser.go.y:136
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 16:
		//line parser.go.y:144
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 17:
		//line parser.go.y:149
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:154
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:158
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:162
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:166
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:171
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 23:
		//line parser.go.y:176
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 24:
		//line parser.go.y:180
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 25:
		//line parser.go.y:184
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 26:
		//line parser.go.y:189
		{
			yyVAL.idents = []string{}
		}
	case 27:
		//line parser.go.y:193
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 28:
		//line parser.go.y:197
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 29:
		//line parser.go.y:202
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 30:
		//line parser.go.y:206
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 31:
		//line parser.go.y:210
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 32:
		//line parser.go.y:215
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 33:
		//line parser.go.y:219
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 34:
		//line parser.go.y:223
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 35:
		//line parser.go.y:227
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 36:
		//line parser.go.y:231
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 37:
		//line parser.go.y:238
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 38:
		//line parser.go.y:245
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 39:
		//line parser.go.y:252
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 40:
		//line parser.go.y:259
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 41:
		//line parser.go.y:266
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 42:
		//line parser.go.y:273
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
		//line parser.go.y:284
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 44:
		//line parser.go.y:291
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 45:
		//line parser.go.y:298
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 46:
		//line parser.go.y:305
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 47:
		//line parser.go.y:312
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 48:
		//line parser.go.y:316
		{
			yyVAL.expr = &ast.ConstExpr{Value: nil}
		}
	case 49:
		//line parser.go.y:320
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 50:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 51:
		//line parser.go.y:334
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 52:
		//line parser.go.y:341
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:348
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:355
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:362
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:376
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 59:
		//line parser.go.y:390
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 60:
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 61:
		//line parser.go.y:404
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 62:
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:418
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 64:
		//line parser.go.y:425
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
