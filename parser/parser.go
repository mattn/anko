//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:23
type yySymType struct {
	yys      int
	stmt_var ast.Stmt
	//stmt_func              ast.Stmt
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
const VAR = 57353
const THROW = 57354
const IF = 57355
const ELSE = 57356
const FOR = 57357
const IN = 57358
const EQEQ = 57359
const NEQ = 57360
const GE = 57361
const LE = 57362
const OROR = 57363
const ANDAND = 57364
const NEW = 57365
const TRUE = 57366
const FALSE = 57367
const NIL = 57368
const MODULE = 57369
const TRY = 57370
const CATCH = 57371
const FINALLY = 57372
const PLUSEQ = 57373
const MINUSEQ = 57374
const MULEQ = 57375
const DIVEQ = 57376
const UNARY = 57377

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"FUNC",
	"RETURN",
	"VAR",
	"THROW",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"OROR",
	"ANDAND",
	"NEW",
	"TRUE",
	"FALSE",
	"NIL",
	"MODULE",
	"TRY",
	"CATCH",
	"FINALLY",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	" =",
	" ?",
	" :",
	" >",
	" <",
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

//line parser.go.y:404

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 25,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	35, 25,
	51, 25,
	-2, 1,
	-1, 3,
	35, 25,
	51, 25,
	-2, 1,
	-1, 4,
	35, 25,
	51, 25,
	-2, 1,
	-1, 5,
	35, 25,
	51, 25,
	-2, 1,
	-1, 6,
	35, 25,
	51, 25,
	-2, 1,
	-1, 16,
	35, 26,
	51, 26,
	-2, 33,
	-1, 22,
	53, 28,
	-2, 25,
	-1, 29,
	35, 25,
	51, 25,
	-2, 1,
	-1, 53,
	50, 28,
	-2, 25,
	-1, 61,
	48, 1,
	-2, 25,
	-1, 66,
	50, 28,
	-2, 25,
	-1, 70,
	35, 26,
	-2, 33,
	-1, 79,
	35, 25,
	51, 25,
	-2, 28,
	-1, 81,
	48, 1,
	-2, 25,
	-1, 91,
	17, 0,
	18, 0,
	-2, 55,
	-1, 92,
	17, 0,
	18, 0,
	-2, 56,
	-1, 102,
	48, 1,
	-2, 25,
	-1, 103,
	35, 25,
	51, 25,
	-2, 28,
	-1, 113,
	35, 25,
	51, 25,
	-2, 28,
	-1, 114,
	35, 25,
	51, 25,
	-2, 28,
	-1, 122,
	50, 28,
	-2, 25,
	-1, 149,
	48, 1,
	-2, 25,
	-1, 150,
	48, 1,
	-2, 25,
	-1, 152,
	48, 1,
	-2, 25,
	-1, 161,
	48, 1,
	-2, 25,
	-1, 163,
	48, 1,
	-2, 25,
	-1, 164,
	48, 1,
	-2, 25,
	-1, 166,
	48, 1,
	-2, 25,
	-1, 174,
	48, 1,
	-2, 25,
	-1, 182,
	48, 1,
	-2, 25,
	-1, 186,
	48, 1,
	-2, 25,
	-1, 191,
	48, 1,
	-2, 25,
}

const yyNprod = 72
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 674

var yyAct = []int{

	1, 7, 74, 28, 30, 31, 33, 34, 27, 112,
	54, 55, 154, 78, 138, 78, 68, 176, 165, 67,
	57, 40, 41, 42, 69, 156, 119, 76, 53, 118,
	80, 36, 153, 37, 49, 51, 135, 83, 84, 129,
	86, 87, 88, 89, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 69, 62, 63, 64, 65,
	126, 104, 106, 72, 107, 108, 109, 110, 69, 103,
	101, 161, 79, 160, 66, 38, 39, 40, 41, 42,
	115, 69, 125, 111, 53, 78, 193, 36, 78, 37,
	49, 51, 53, 122, 117, 36, 124, 37, 49, 51,
	59, 190, 187, 130, 184, 69, 181, 133, 71, 179,
	178, 177, 171, 168, 167, 69, 69, 148, 145, 134,
	131, 142, 143, 191, 69, 186, 140, 182, 146, 147,
	136, 137, 174, 166, 43, 44, 46, 48, 164, 144,
	82, 163, 152, 149, 102, 61, 120, 105, 189, 183,
	158, 159, 151, 162, 155, 45, 47, 38, 39, 40,
	41, 42, 170, 139, 172, 173, 53, 175, 75, 36,
	169, 37, 49, 51, 81, 180, 43, 44, 46, 48,
	50, 52, 141, 185, 123, 116, 85, 188, 77, 60,
	58, 56, 192, 73, 6, 35, 5, 45, 47, 38,
	39, 40, 41, 42, 62, 63, 64, 65, 53, 157,
	4, 36, 3, 37, 49, 51, 43, 44, 46, 48,
	50, 52, 66, 2, 114, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 35, 0, 45, 47, 38,
	39, 40, 41, 42, 0, 0, 150, 0, 53, 0,
	0, 36, 0, 37, 49, 51, 43, 44, 46, 48,
	50, 52, 0, 0, 0, 16, 15, 18, 0, 0,
	23, 0, 0, 0, 0, 35, 0, 45, 47, 38,
	39, 40, 41, 42, 26, 19, 20, 21, 53, 132,
	0, 36, 0, 37, 49, 51, 43, 44, 46, 48,
	50, 52, 17, 0, 0, 0, 0, 0, 24, 0,
	25, 0, 0, 22, 0, 35, 0, 45, 47, 38,
	39, 40, 41, 42, 0, 0, 0, 0, 53, 0,
	0, 36, 128, 37, 49, 51, 43, 44, 46, 48,
	50, 52, 0, 0, 0, 70, 15, 18, 0, 0,
	23, 0, 0, 0, 0, 35, 127, 45, 47, 38,
	39, 40, 41, 42, 26, 19, 20, 21, 53, 0,
	0, 36, 0, 37, 49, 51, 43, 44, 46, 48,
	50, 52, 17, 0, 0, 0, 0, 0, 24, 0,
	25, 0, 0, 22, 0, 35, 0, 45, 47, 38,
	39, 40, 41, 42, 0, 0, 0, 0, 53, 121,
	0, 36, 0, 37, 49, 51, 43, 44, 46, 48,
	50, 52, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 35, 0, 45, 47, 38,
	39, 40, 41, 42, 0, 0, 0, 0, 53, 0,
	113, 36, 0, 37, 49, 51, 43, 44, 46, 48,
	50, 52, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 35, 0, 45, 47, 38,
	39, 40, 41, 42, 43, 44, 46, 48, 53, 52,
	0, 36, 0, 37, 49, 51, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 45, 47, 38, 39, 40,
	41, 42, 0, 0, 0, 0, 53, 0, 0, 36,
	0, 37, 49, 51, 16, 15, 18, 0, 0, 23,
	8, 11, 9, 12, 32, 13, 0, 0, 0, 0,
	0, 0, 0, 26, 19, 20, 21, 10, 14, 0,
	0, 0, 0, 0, 0, 0, 0, 16, 15, 18,
	0, 17, 23, 8, 11, 9, 12, 24, 13, 25,
	0, 0, 22, 0, 0, 0, 26, 19, 20, 21,
	10, 14, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 17, 0, 0, 0, 0, 29,
	24, 0, 25, 0, 0, 22, 16, 15, 18, 0,
	0, 23, 8, 11, 9, 12, 0, 13, 0, 0,
	0, 0, 0, 0, 0, 26, 19, 20, 21, 10,
	14, 0, 0, 0, 0, 0, 46, 48, 0, 0,
	0, 0, 0, 17, 0, 0, 0, 0, 0, 24,
	0, 25, 0, 0, 22, 45, 47, 38, 39, 40,
	41, 42, 0, 0, 0, 0, 53, 0, 0, 36,
	0, 37, 49, 51,
}
var yyPact = []int{

	602, -1000, 553, 602, 520, 602, 602, 439, 261, 261,
	187, 186, 51, 185, 98, -1000, 25, 261, -1000, -1000,
	-1000, -1000, 341, 59, 162, 261, 184, 37, -1000, 602,
	-1000, -1000, 127, -1000, -1000, 261, 261, 182, 261, 261,
	261, 261, 261, 261, 261, 261, 261, 261, 261, 261,
	261, 261, 261, 341, 439, 439, 97, 34, -1000, 261,
	131, 602, 261, 261, 261, 261, 341, 43, -44, 399,
	173, 181, 45, -22, -1000, 109, 359, 44, 180, 341,
	-1000, 602, 11, 319, 279, -1000, -21, -21, 43, 43,
	43, 617, 617, 35, 35, 35, 35, 439, 467, 439,
	117, -11, 602, 341, 239, 261, 71, 439, 439, 439,
	439, -14, -1000, 341, 341, -36, 155, 178, 162, -1000,
	261, -1000, 341, -1000, -1000, 70, 261, 261, -1000, -1000,
	69, -1000, 96, 199, 123, -1000, -1000, -1000, 95, -18,
	-38, 146, -1000, 439, -25, -1000, 159, 439, -1000, 602,
	602, 24, 602, 94, 91, -32, -1000, 86, 66, 65,
	166, 602, 64, 602, 602, 85, 602, -1000, -1000, -33,
	63, -1000, 62, 61, 602, 58, 80, 119, -1000, -1000,
	56, -1000, 602, 78, -1000, 54, 602, 118, 53, 76,
	-1000, 602, 38, -1000,
}
var yyPgo = []int{

	0, 0, 223, 212, 210, 196, 194, 1, 16, 2,
	193, 8,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 2, 3, 4, 4, 4, 5, 6, 6, 6,
	6, 9, 10, 10, 10, 11, 11, 11, 8, 8,
	8, 8, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 1, 2,
	2, 5, 4, 7, 5, 9, 7, 15, 12, 11,
	8, 3, 0, 1, 3, 0, 1, 3, 0, 1,
	3, 3, 1, 1, 2, 1, 1, 1, 1, 5,
	4, 3, 7, 8, 8, 9, 3, 3, 3, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 4,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 10, 12,
	27, 11, 13, 15, 28, 5, 4, 41, 6, 24,
	25, 26, 52, 9, 47, 49, 23, -11, -1, 46,
	-1, -1, 14, -1, -1, 36, 52, 54, 40, 41,
	42, 43, 44, 17, 18, 38, 19, 39, 20, 55,
	21, 56, 22, 49, -7, -7, 4, -11, 4, 49,
	4, 47, 31, 32, 33, 34, 49, -7, -8, -7,
	4, 49, 4, -10, -9, 6, -7, 4, 51, 35,
	-1, 47, 13, -7, -7, 4, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -8, 47, 35, -7, 16, -1, -7, -7, -7,
	-7, -8, 53, 51, 51, -11, 4, 49, 51, 48,
	37, 50, 49, 4, -8, -1, 49, 37, 53, 50,
	-1, -8, 50, -7, 48, 50, -8, -8, 50, 8,
	-11, 4, -9, -7, -8, 48, -7, -7, 48, 47,
	47, 29, 47, 50, 50, 8, 50, 50, -1, -1,
	49, 47, -1, 47, 47, 50, 47, 48, 48, 4,
	-1, 48, -1, -1, 47, -1, 50, 48, 48, 48,
	-1, 48, 47, 30, 48, -1, 47, 48, -1, 30,
	48, 47, -1, 48,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, 8, 25, 25,
	0, 25, 0, 0, 0, 32, -2, 25, 35, 36,
	37, 38, -2, 0, 22, 25, 0, 0, 2, -2,
	4, 5, 0, 6, 7, 25, 25, 0, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, -2, 9, 10, 0, 0, 26, 25,
	0, -2, 25, 25, 25, 25, -2, 34, 0, 29,
	-2, 25, 0, 0, 23, 0, 0, 0, 0, -2,
	3, -2, 0, 0, 0, 47, 50, 51, 52, 53,
	54, -2, -2, 57, 58, 59, 60, 66, 67, 68,
	69, 0, -2, -2, 0, 25, 0, 62, 63, 64,
	65, 0, 41, -2, -2, 0, 26, 25, 0, 46,
	25, 48, -2, 27, 61, 0, 25, 25, 40, 71,
	0, 12, 0, 0, 0, 70, 30, 31, 0, 0,
	0, 26, 24, 21, 0, 14, 0, 39, 11, -2,
	-2, 0, -2, 0, 0, 0, 49, 0, 0, 0,
	0, -2, 0, -2, -2, 0, -2, 13, 16, 0,
	0, 42, 0, 0, -2, 0, 0, 20, 43, 44,
	0, 15, -2, 0, 45, 0, -2, 19, 0, 0,
	18, -2, 0, 17,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 44, 56, 3,
	49, 50, 42, 40, 51, 41, 54, 43, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 37, 46,
	39, 35, 38, 36, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 52, 3, 53, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 47, 55, 48,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 45,
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
		//line parser.go.y:56
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:63
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:70
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:77
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:91
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:98
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:105
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:113
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 9:
		//line parser.go.y:120
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 10:
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:134
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 12:
		//line parser.go.y:139
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 13:
		//line parser.go.y:153
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:157
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 15:
		//line parser.go.y:165
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 16:
		//line parser.go.y:170
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 17:
		//line parser.go.y:175
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:179
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:183
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:187
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:192
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 22:
		//line parser.go.y:197
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 23:
		//line parser.go.y:201
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 24:
		//line parser.go.y:205
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 25:
		//line parser.go.y:210
		{
			yyVAL.idents = []string{}
		}
	case 26:
		//line parser.go.y:214
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 27:
		//line parser.go.y:218
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 28:
		//line parser.go.y:223
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 29:
		//line parser.go.y:227
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 30:
		//line parser.go.y:231
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 31:
		//line parser.go.y:235
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 32:
		//line parser.go.y:240
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 33:
		//line parser.go.y:244
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 34:
		//line parser.go.y:248
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 35:
		//line parser.go.y:252
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 36:
		//line parser.go.y:256
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 37:
		//line parser.go.y:260
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 38:
		//line parser.go.y:264
		{
			yyVAL.expr = &ast.ConstExpr{Value: (*interface{})(nil)}
		}
	case 39:
		//line parser.go.y:268
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	case 40:
		//line parser.go.y:272
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 41:
		//line parser.go.y:276
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 42:
		//line parser.go.y:280
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 43:
		//line parser.go.y:284
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 44:
		//line parser.go.y:288
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 45:
		//line parser.go.y:292
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 46:
		//line parser.go.y:296
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 47:
		//line parser.go.y:304
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 48:
		//line parser.go.y:308
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 49:
		//line parser.go.y:312
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 50:
		//line parser.go.y:316
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 51:
		//line parser.go.y:320
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 52:
		//line parser.go.y:324
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 53:
		//line parser.go.y:328
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 54:
		//line parser.go.y:332
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:336
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:340
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:348
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:352
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:356
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:360
		{
			yyVAL.expr = &ast.LetExpr{Names: yyS[yypt-2].idents, Operator: "=", Exprs: yyS[yypt-0].exprs}
		}
	case 62:
		//line parser.go.y:364
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "+", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 63:
		//line parser.go.y:368
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "-", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 64:
		//line parser.go.y:372
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "*", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 65:
		//line parser.go.y:376
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "/", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 66:
		//line parser.go.y:380
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 67:
		//line parser.go.y:384
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 68:
		//line parser.go.y:388
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 69:
		//line parser.go.y:392
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 70:
		//line parser.go.y:396
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 71:
		//line parser.go.y:400
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	}
	goto yystack /* stack new state and value */
}
