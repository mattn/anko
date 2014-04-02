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
	stmt_var               ast.Stmt
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

//line parser.go.y:396

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 28,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	35, 28,
	51, 28,
	-2, 1,
	-1, 3,
	35, 28,
	51, 28,
	-2, 1,
	-1, 4,
	35, 28,
	51, 28,
	-2, 1,
	-1, 5,
	35, 28,
	51, 28,
	-2, 1,
	-1, 6,
	35, 28,
	51, 28,
	-2, 1,
	-1, 7,
	35, 28,
	51, 28,
	-2, 1,
	-1, 18,
	35, 29,
	51, 29,
	-2, 36,
	-1, 24,
	53, 31,
	-2, 28,
	-1, 30,
	35, 28,
	51, 28,
	-2, 1,
	-1, 55,
	50, 31,
	-2, 28,
	-1, 66,
	48, 1,
	-2, 28,
	-1, 71,
	50, 31,
	-2, 28,
	-1, 75,
	35, 29,
	-2, 36,
	-1, 82,
	35, 28,
	51, 28,
	-2, 31,
	-1, 84,
	48, 1,
	-2, 28,
	-1, 94,
	17, 0,
	18, 0,
	-2, 56,
	-1, 95,
	17, 0,
	18, 0,
	-2, 57,
	-1, 105,
	48, 1,
	-2, 28,
	-1, 106,
	35, 28,
	51, 28,
	-2, 31,
	-1, 119,
	35, 28,
	51, 28,
	-2, 31,
	-1, 120,
	35, 28,
	51, 28,
	-2, 31,
	-1, 125,
	50, 31,
	-2, 28,
	-1, 154,
	48, 1,
	-2, 28,
	-1, 156,
	48, 1,
	-2, 28,
	-1, 157,
	48, 1,
	-2, 28,
	-1, 161,
	48, 1,
	-2, 28,
	-1, 164,
	48, 1,
	-2, 28,
	-1, 168,
	48, 1,
	-2, 28,
	-1, 169,
	48, 1,
	-2, 28,
	-1, 171,
	48, 1,
	-2, 28,
	-1, 186,
	48, 1,
	-2, 28,
	-1, 189,
	48, 1,
	-2, 28,
	-1, 194,
	48, 1,
	-2, 28,
}

const yyNprod = 73
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 701

var yyAct = []int{

	1, 8, 77, 29, 31, 32, 33, 35, 36, 73,
	28, 56, 58, 55, 118, 182, 38, 106, 39, 51,
	53, 72, 82, 60, 152, 81, 74, 162, 79, 137,
	81, 83, 122, 81, 62, 121, 159, 129, 81, 86,
	87, 155, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 74, 67, 68,
	69, 70, 168, 142, 167, 104, 110, 112, 132, 113,
	114, 115, 116, 74, 108, 125, 71, 107, 120, 63,
	63, 117, 64, 196, 74, 128, 40, 41, 42, 43,
	44, 193, 127, 190, 85, 55, 185, 184, 38, 183,
	39, 51, 53, 181, 179, 175, 133, 174, 74, 172,
	151, 148, 141, 140, 194, 189, 134, 186, 135, 171,
	169, 74, 74, 164, 145, 146, 161, 74, 84, 143,
	144, 149, 150, 156, 154, 147, 42, 43, 44, 105,
	66, 123, 192, 55, 187, 158, 38, 111, 39, 51,
	53, 153, 138, 78, 176, 163, 136, 165, 166, 126,
	109, 88, 170, 80, 65, 173, 61, 59, 76, 177,
	178, 7, 180, 6, 5, 45, 46, 48, 50, 52,
	54, 4, 3, 2, 18, 17, 20, 188, 0, 57,
	191, 0, 0, 0, 37, 195, 47, 49, 40, 41,
	42, 43, 44, 27, 21, 22, 23, 55, 160, 0,
	38, 0, 39, 51, 53, 45, 46, 48, 50, 52,
	54, 19, 0, 0, 0, 0, 0, 25, 0, 26,
	0, 0, 24, 0, 37, 0, 47, 49, 40, 41,
	42, 43, 44, 0, 0, 157, 0, 55, 0, 0,
	38, 0, 39, 51, 53, 45, 46, 48, 50, 52,
	54, 0, 0, 0, 75, 17, 20, 0, 0, 57,
	0, 0, 0, 0, 37, 0, 47, 49, 40, 41,
	42, 43, 44, 27, 21, 22, 23, 55, 139, 0,
	38, 0, 39, 51, 53, 45, 46, 48, 50, 52,
	54, 19, 0, 0, 0, 0, 0, 25, 0, 26,
	0, 0, 24, 0, 37, 0, 47, 49, 40, 41,
	42, 43, 44, 67, 68, 69, 70, 55, 0, 0,
	38, 131, 39, 51, 53, 45, 46, 48, 50, 52,
	54, 71, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 37, 130, 47, 49, 40, 41,
	42, 43, 44, 0, 0, 0, 0, 55, 0, 0,
	38, 0, 39, 51, 53, 45, 46, 48, 50, 52,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 37, 0, 47, 49, 40, 41,
	42, 43, 44, 0, 0, 0, 0, 55, 124, 0,
	38, 0, 39, 51, 53, 45, 46, 48, 50, 52,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 37, 0, 47, 49, 40, 41,
	42, 43, 44, 0, 0, 0, 0, 55, 0, 119,
	38, 0, 39, 51, 53, 45, 46, 48, 50, 52,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 37, 0, 47, 49, 40, 41,
	42, 43, 44, 45, 46, 48, 50, 55, 54, 0,
	38, 0, 39, 51, 53, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 49, 40, 41, 42, 43,
	44, 45, 46, 48, 50, 55, 0, 0, 38, 0,
	39, 51, 53, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 47, 49, 40, 41, 42, 43, 44, 0,
	0, 0, 0, 55, 0, 0, 38, 0, 39, 51,
	53, 18, 17, 20, 0, 0, 13, 9, 12, 10,
	14, 34, 15, 0, 0, 0, 0, 0, 0, 0,
	27, 21, 22, 23, 11, 16, 0, 0, 0, 0,
	0, 0, 0, 0, 18, 17, 20, 0, 19, 13,
	9, 12, 10, 14, 25, 15, 26, 0, 0, 24,
	0, 0, 0, 27, 21, 22, 23, 11, 16, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 19, 0, 0, 0, 0, 30, 25, 0, 26,
	0, 0, 24, 18, 17, 20, 0, 0, 13, 9,
	12, 10, 14, 0, 15, 0, 0, 0, 0, 0,
	0, 0, 27, 21, 22, 23, 11, 16, 0, 0,
	0, 0, 0, 48, 50, 0, 0, 0, 0, 0,
	19, 0, 0, 0, 0, 0, 25, 0, 26, 0,
	0, 24, 47, 49, 40, 41, 42, 43, 44, 0,
	0, 0, 0, 55, 0, 0, 38, 0, 39, 51,
	53,
}
var yyPact = []int{

	629, -1000, 580, 629, 629, 547, 629, 629, 438, 180,
	180, 163, 162, 30, 33, 160, 93, -1000, 292, 180,
	-1000, -1000, -1000, -1000, 260, 147, 180, 159, -13, -1000,
	629, -1000, -1000, -1000, 81, -1000, -1000, 180, 180, 157,
	180, 180, 180, 180, 180, 180, 180, 180, 180, 180,
	180, 180, 180, 180, 180, 260, 438, 31, 438, 92,
	-18, -1000, 28, 156, 180, 131, 629, 180, 180, 180,
	180, 260, -36, -39, 398, 27, -16, -1000, 104, 358,
	26, 155, 260, -1000, 629, -12, 318, 278, -1000, 94,
	94, -36, -36, -36, 644, 644, 46, 46, 46, 46,
	438, 466, 438, 494, 18, 629, 260, 152, -21, 144,
	238, 180, 64, 438, 438, 438, 438, 13, -1000, 260,
	260, 147, -1000, 180, -1000, 260, -1000, -1000, 63, 180,
	180, -1000, -1000, 62, -1000, -26, 143, 87, -9, 86,
	198, 116, -1000, -1000, -1000, -1000, 438, -14, -1000, 158,
	438, -1000, 79, -23, 629, 76, 629, 629, 15, -1000,
	73, 629, 72, 61, 629, 59, 57, 150, 629, 629,
	56, 629, -1000, 55, -1000, -1000, -35, 51, 49, -1000,
	48, -1000, 70, 114, -1000, -1000, 629, 68, 45, 629,
	112, 43, 67, -1000, 629, 35, -1000,
}
var yyPgo = []int{

	0, 0, 183, 182, 181, 174, 173, 171, 1, 9,
	2, 168, 10,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 2, 2, 3, 4, 4, 5, 5, 5, 6,
	7, 7, 7, 7, 10, 11, 11, 11, 12, 12,
	12, 9, 9, 9, 9, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 1,
	2, 2, 5, 4, 8, 9, 7, 5, 9, 7,
	15, 12, 11, 8, 3, 0, 1, 3, 0, 1,
	3, 0, 1, 3, 3, 1, 1, 2, 1, 1,
	1, 1, 5, 4, 3, 7, 8, 3, 3, 3,
	5, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, 10,
	12, 27, 11, 9, 13, 15, 28, 5, 4, 41,
	6, 24, 25, 26, 52, 47, 49, 23, -12, -1,
	46, -1, -1, -1, 14, -1, -1, 36, 52, 54,
	40, 41, 42, 43, 44, 17, 18, 38, 19, 39,
	20, 55, 21, 56, 22, 49, -8, 9, -8, 4,
	-12, 4, 4, 49, 49, 4, 47, 31, 32, 33,
	34, 49, -8, -9, -8, 4, -11, -10, 6, -8,
	4, 51, 35, -1, 47, 13, -8, -8, 4, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -9, 47, 35, 49, -12, 4,
	-8, 16, -1, -8, -8, -8, -8, -9, 53, 51,
	51, 51, 48, 37, 50, 49, 4, -9, -1, 49,
	37, 53, 50, -1, -9, -12, 4, 50, 8, 50,
	-8, 48, 50, -9, -9, -10, -8, -9, 48, -8,
	-8, 48, 50, 8, 47, 50, 47, 47, 29, 50,
	50, 47, 50, -1, 47, -1, -1, 49, 47, 47,
	-1, 47, 48, -1, 48, 48, 4, -1, -1, 48,
	-1, 48, 50, 48, 48, 48, 47, 30, -1, 47,
	48, -1, 30, 48, 47, -1, 48,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, 9, 28,
	28, 0, 28, 0, 0, 0, 0, 35, -2, 28,
	38, 39, 40, 41, -2, 25, 28, 0, 0, 2,
	-2, 4, 5, 6, 0, 7, 8, 28, 28, 0,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, -2, 10, 0, 11, 0,
	0, 29, 0, 28, 28, 0, -2, 28, 28, 28,
	28, -2, 37, 0, 32, -2, 0, 26, 0, 0,
	0, 0, -2, 3, -2, 0, 0, 0, 48, 51,
	52, 53, 54, 55, -2, -2, 58, 59, 60, 61,
	67, 68, 69, 70, 0, -2, -2, 28, 0, 29,
	0, 28, 0, 63, 64, 65, 66, 0, 44, -2,
	-2, 0, 47, 28, 49, -2, 30, 62, 0, 28,
	28, 43, 72, 0, 13, 0, 29, 0, 0, 0,
	0, 0, 71, 33, 34, 27, 24, 0, 17, 0,
	42, 12, 0, 0, -2, 0, -2, -2, 0, 50,
	0, -2, 0, 0, -2, 0, 0, 0, -2, -2,
	0, -2, 45, 0, 16, 19, 0, 0, 0, 14,
	0, 46, 0, 23, 18, 15, -2, 0, 0, -2,
	22, 0, 0, 21, -2, 0, 20,
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
		//line parser.go.y:84
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:91
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:98
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:105
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 9:
		//line parser.go.y:113
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 10:
		//line parser.go.y:120
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 12:
		//line parser.go.y:134
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 13:
		//line parser.go.y:139
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 14:
		//line parser.go.y:144
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 15:
		//line parser.go.y:148
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 16:
		//line parser.go.y:153
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 17:
		//line parser.go.y:157
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 18:
		//line parser.go.y:165
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 19:
		//line parser.go.y:170
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:175
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:179
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:183
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:187
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:192
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 25:
		//line parser.go.y:197
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 26:
		//line parser.go.y:201
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 27:
		//line parser.go.y:205
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 28:
		//line parser.go.y:210
		{
			yyVAL.idents = []string{}
		}
	case 29:
		//line parser.go.y:214
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 30:
		//line parser.go.y:218
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 31:
		//line parser.go.y:223
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 32:
		//line parser.go.y:227
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 33:
		//line parser.go.y:231
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 34:
		//line parser.go.y:235
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 35:
		//line parser.go.y:240
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 36:
		//line parser.go.y:244
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 37:
		//line parser.go.y:248
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 38:
		//line parser.go.y:252
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 39:
		//line parser.go.y:256
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 40:
		//line parser.go.y:260
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 41:
		//line parser.go.y:264
		{
			yyVAL.expr = &ast.ConstExpr{Value: nil}
		}
	case 42:
		//line parser.go.y:268
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	case 43:
		//line parser.go.y:272
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 44:
		//line parser.go.y:276
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 45:
		//line parser.go.y:280
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 46:
		//line parser.go.y:284
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 47:
		//line parser.go.y:288
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 48:
		//line parser.go.y:296
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 49:
		//line parser.go.y:300
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 50:
		//line parser.go.y:304
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 51:
		//line parser.go.y:308
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 52:
		//line parser.go.y:312
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 53:
		//line parser.go.y:316
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 54:
		//line parser.go.y:320
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:324
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:328
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:332
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:336
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:340
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:348
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:352
		{
			yyVAL.expr = &ast.LetExpr{Names: yyS[yypt-2].idents, Operator: "=", Exprs: yyS[yypt-0].exprs}
		}
	case 63:
		//line parser.go.y:356
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "+", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 64:
		//line parser.go.y:360
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "-", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 65:
		//line parser.go.y:364
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "*", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 66:
		//line parser.go.y:368
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "/", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 67:
		//line parser.go.y:372
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 68:
		//line parser.go.y:376
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 69:
		//line parser.go.y:380
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 70:
		//line parser.go.y:384
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 71:
		//line parser.go.y:388
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 72:
		//line parser.go.y:392
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	}
	goto yystack /* stack new state and value */
}
