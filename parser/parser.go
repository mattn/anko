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
	stmt_if                ast.Stmt
	stmt_for               ast.Stmt
	stmt_continue          ast.Stmt
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
const BREAK = 57377
const CONTINUE = 57378
const UNARY = 57379

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
	"BREAK",
	"CONTINUE",
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

//line parser.go.y:407

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 28,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	37, 28,
	53, 28,
	-2, 1,
	-1, 3,
	37, 28,
	53, 28,
	-2, 1,
	-1, 4,
	37, 28,
	53, 28,
	-2, 1,
	-1, 5,
	37, 28,
	53, 28,
	-2, 1,
	-1, 6,
	37, 28,
	53, 28,
	-2, 1,
	-1, 7,
	37, 28,
	53, 28,
	-2, 1,
	-1, 8,
	37, 28,
	53, 28,
	-2, 1,
	-1, 19,
	37, 29,
	53, 29,
	-2, 36,
	-1, 25,
	55, 31,
	-2, 28,
	-1, 32,
	37, 28,
	53, 28,
	-2, 1,
	-1, 58,
	52, 31,
	-2, 28,
	-1, 66,
	50, 1,
	-2, 28,
	-1, 71,
	52, 31,
	-2, 28,
	-1, 75,
	37, 29,
	-2, 36,
	-1, 84,
	37, 28,
	53, 28,
	-2, 31,
	-1, 86,
	50, 1,
	-2, 28,
	-1, 96,
	17, 0,
	18, 0,
	-2, 58,
	-1, 97,
	17, 0,
	18, 0,
	-2, 59,
	-1, 107,
	50, 1,
	-2, 28,
	-1, 108,
	37, 28,
	53, 28,
	-2, 31,
	-1, 118,
	37, 28,
	53, 28,
	-2, 31,
	-1, 119,
	37, 28,
	53, 28,
	-2, 31,
	-1, 127,
	52, 31,
	-2, 28,
	-1, 154,
	50, 1,
	-2, 28,
	-1, 155,
	50, 1,
	-2, 28,
	-1, 157,
	50, 1,
	-2, 28,
	-1, 166,
	50, 1,
	-2, 28,
	-1, 168,
	50, 1,
	-2, 28,
	-1, 169,
	50, 1,
	-2, 28,
	-1, 171,
	50, 1,
	-2, 28,
	-1, 179,
	50, 1,
	-2, 28,
	-1, 187,
	50, 1,
	-2, 28,
	-1, 191,
	50, 1,
	-2, 28,
	-1, 196,
	50, 1,
	-2, 28,
}

const yyNprod = 75
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 697

var yyAct = []int{

	1, 9, 79, 31, 33, 34, 35, 36, 38, 39,
	117, 181, 59, 60, 30, 170, 73, 159, 83, 143,
	83, 161, 72, 108, 45, 46, 47, 74, 84, 62,
	81, 58, 124, 85, 41, 123, 42, 54, 56, 83,
	158, 140, 88, 89, 83, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 104, 105,
	74, 166, 134, 165, 77, 131, 109, 111, 127, 112,
	113, 114, 115, 74, 122, 106, 64, 198, 195, 192,
	43, 44, 45, 46, 47, 189, 74, 130, 116, 58,
	186, 120, 41, 184, 42, 54, 56, 58, 183, 182,
	41, 129, 42, 54, 56, 176, 173, 172, 135, 153,
	74, 76, 138, 150, 139, 196, 87, 191, 187, 179,
	74, 74, 171, 169, 168, 136, 147, 148, 157, 74,
	154, 107, 66, 151, 152, 141, 142, 145, 125, 51,
	53, 194, 188, 156, 149, 110, 160, 144, 67, 68,
	69, 70, 86, 80, 174, 163, 164, 146, 167, 128,
	50, 52, 43, 44, 45, 46, 47, 175, 71, 177,
	178, 58, 180, 121, 41, 90, 42, 54, 56, 82,
	185, 48, 49, 51, 53, 55, 57, 65, 190, 63,
	61, 78, 193, 8, 4, 7, 6, 197, 5, 2,
	0, 0, 40, 0, 50, 52, 43, 44, 45, 46,
	47, 67, 68, 69, 70, 58, 162, 0, 41, 0,
	42, 54, 56, 48, 49, 51, 53, 55, 57, 0,
	0, 71, 0, 119, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 40, 0, 50, 52, 43, 44,
	45, 46, 47, 0, 0, 155, 0, 58, 0, 0,
	41, 0, 42, 54, 56, 48, 49, 51, 53, 55,
	57, 0, 0, 0, 0, 0, 19, 18, 21, 0,
	0, 26, 0, 0, 0, 0, 40, 0, 50, 52,
	43, 44, 45, 46, 47, 29, 22, 23, 24, 58,
	137, 0, 41, 0, 42, 54, 56, 48, 49, 51,
	53, 55, 57, 0, 0, 20, 0, 0, 0, 0,
	0, 27, 0, 28, 0, 0, 25, 0, 40, 0,
	50, 52, 43, 44, 45, 46, 47, 0, 0, 0,
	0, 58, 0, 0, 41, 133, 42, 54, 56, 48,
	49, 51, 53, 55, 57, 0, 0, 0, 0, 0,
	75, 18, 21, 0, 0, 26, 0, 0, 0, 0,
	40, 132, 50, 52, 43, 44, 45, 46, 47, 29,
	22, 23, 24, 58, 0, 0, 41, 0, 42, 54,
	56, 48, 49, 51, 53, 55, 57, 0, 0, 20,
	0, 0, 0, 0, 0, 27, 0, 28, 0, 0,
	25, 0, 40, 0, 50, 52, 43, 44, 45, 46,
	47, 0, 0, 0, 0, 58, 126, 0, 41, 0,
	42, 54, 56, 48, 49, 51, 53, 55, 57, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 40, 0, 50, 52, 43, 44,
	45, 46, 47, 0, 0, 0, 0, 58, 0, 118,
	41, 0, 42, 54, 56, 48, 49, 51, 53, 55,
	57, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 40, 0, 50, 52,
	43, 44, 45, 46, 47, 0, 0, 0, 0, 58,
	0, 0, 41, 0, 42, 54, 56, 19, 18, 21,
	0, 0, 26, 10, 14, 11, 15, 37, 16, 0,
	0, 0, 0, 0, 0, 0, 29, 22, 23, 24,
	12, 17, 0, 0, 0, 0, 0, 0, 3, 13,
	0, 0, 19, 18, 21, 0, 20, 26, 10, 14,
	11, 15, 27, 16, 28, 0, 0, 25, 0, 0,
	0, 29, 22, 23, 24, 12, 17, 0, 0, 0,
	0, 0, 0, 3, 13, 0, 0, 0, 0, 0,
	0, 20, 0, 0, 0, 0, 32, 27, 0, 28,
	0, 0, 25, 19, 18, 21, 0, 0, 26, 10,
	14, 11, 15, 0, 16, 0, 0, 0, 0, 0,
	0, 0, 29, 22, 23, 24, 12, 17, 0, 0,
	0, 0, 0, 0, 3, 13, 48, 49, 51, 53,
	0, 57, 20, 0, 0, 0, 0, 0, 27, 0,
	28, 0, 0, 25, 0, 48, 49, 51, 53, 50,
	52, 43, 44, 45, 46, 47, 0, 0, 0, 0,
	58, 0, 0, 41, 0, 42, 54, 56, 50, 52,
	43, 44, 45, 46, 47, 0, 0, 0, 0, 58,
	0, 0, 41, 0, 42, 54, 56,
}
var yyPact = []int{

	599, -1000, 548, 599, 599, 599, 513, 599, 599, 458,
	272, 272, 186, -1000, 185, 25, 183, 83, -1000, 117,
	272, -1000, -1000, -1000, -1000, 356, 60, 147, 272, 175,
	-9, -1000, 599, -1000, -1000, -1000, -1000, 103, -1000, -1000,
	272, 272, 171, 272, 272, 272, 272, 272, 272, 272,
	272, 272, 272, 272, 272, 272, 272, 272, 356, 458,
	458, 82, -14, -1000, 272, 129, 599, 272, 272, 272,
	272, 356, 46, -45, 416, 180, 169, 23, -18, -1000,
	99, 374, 17, 155, 356, -1000, 599, 14, 332, 290,
	-1000, -20, -20, 46, 46, 46, 120, 120, 38, 38,
	38, 38, 458, 619, 458, 638, 10, 599, 356, 248,
	272, 64, 458, 458, 458, 458, -11, -1000, 356, 356,
	-33, 139, 153, 147, -1000, 272, -1000, 356, -1000, -1000,
	63, 272, 272, -1000, -1000, 59, -1000, 81, 206, 114,
	-1000, -1000, -1000, 79, -12, -35, 138, -1000, 458, -31,
	-1000, 164, 458, -1000, 599, 599, 12, 599, 75, 74,
	-37, -1000, 73, 57, 56, 150, 599, 55, 599, 599,
	70, 599, -1000, -1000, -41, 49, -1000, 48, 43, 599,
	40, 69, 112, -1000, -1000, 35, -1000, 599, 68, -1000,
	29, 599, 111, 28, 66, -1000, 599, 27, -1000,
}
var yyPgo = []int{

	0, 0, 199, 198, 196, 195, 194, 193, 1, 16,
	2, 191, 14,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 3, 4, 4, 4, 5, 6,
	7, 7, 7, 7, 10, 11, 11, 11, 12, 12,
	12, 9, 9, 9, 9, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 5, 4, 7, 5, 9, 7, 1,
	15, 12, 11, 8, 3, 0, 1, 3, 0, 1,
	3, 0, 1, 3, 3, 1, 1, 2, 1, 1,
	1, 1, 5, 4, 3, 7, 8, 8, 9, 3,
	3, 3, 5, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, 35, -6, -3, -4, -5, -7, -8,
	10, 12, 27, 36, 11, 13, 15, 28, 5, 4,
	43, 6, 24, 25, 26, 54, 9, 49, 51, 23,
	-12, -1, 48, -1, -1, -1, -1, 14, -1, -1,
	38, 54, 56, 42, 43, 44, 45, 46, 17, 18,
	40, 19, 41, 20, 57, 21, 58, 22, 51, -8,
	-8, 4, -12, 4, 51, 4, 49, 31, 32, 33,
	34, 51, -8, -9, -8, 4, 51, 4, -11, -10,
	6, -8, 4, 53, 37, -1, 49, 13, -8, -8,
	4, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -9, 49, 37, -8,
	16, -1, -8, -8, -8, -8, -9, 55, 53, 53,
	-12, 4, 51, 53, 50, 39, 52, 51, 4, -9,
	-1, 51, 39, 55, 52, -1, -9, 52, -8, 50,
	52, -9, -9, 52, 8, -12, 4, -10, -8, -9,
	50, -8, -8, 50, 49, 49, 29, 49, 52, 52,
	8, 52, 52, -1, -1, 51, 49, -1, 49, 49,
	52, 49, 50, 50, 4, -1, 50, -1, -1, 49,
	-1, 52, 50, 50, 50, -1, 50, 49, 30, 50,
	-1, 49, 50, -1, 30, 50, 49, -1, 50,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, 10,
	28, 28, 0, 19, 28, 0, 0, 0, 35, -2,
	28, 38, 39, 40, 41, -2, 0, 25, 28, 0,
	0, 2, -2, 4, 5, 6, 7, 0, 8, 9,
	28, 28, 0, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, -2, 11,
	12, 0, 0, 29, 28, 0, -2, 28, 28, 28,
	28, -2, 37, 0, 32, -2, 28, 0, 0, 26,
	0, 0, 0, 0, -2, 3, -2, 0, 0, 0,
	50, 53, 54, 55, 56, 57, -2, -2, 60, 61,
	62, 63, 69, 70, 71, 72, 0, -2, -2, 0,
	28, 0, 65, 66, 67, 68, 0, 44, -2, -2,
	0, 29, 28, 0, 49, 28, 51, -2, 30, 64,
	0, 28, 28, 43, 74, 0, 14, 0, 0, 0,
	73, 33, 34, 0, 0, 0, 29, 27, 24, 0,
	16, 0, 42, 13, -2, -2, 0, -2, 0, 0,
	0, 52, 0, 0, 0, 0, -2, 0, -2, -2,
	0, -2, 15, 18, 0, 0, 45, 0, 0, -2,
	0, 0, 23, 46, 47, 0, 17, -2, 0, 48,
	0, -2, 22, 0, 0, 21, -2, 0, 20,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 46, 58, 3,
	51, 52, 44, 42, 53, 43, 56, 45, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 39, 48,
	41, 37, 40, 38, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 54, 3, 55, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 49, 57, 50,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 47,
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
			yyVAL.stmts = append([]ast.Stmt{&ast.BreakStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:84
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_continue}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:91
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:98
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:105
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 9:
		//line parser.go.y:112
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 10:
		//line parser.go.y:120
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 12:
		//line parser.go.y:134
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 13:
		//line parser.go.y:141
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:146
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 15:
		//line parser.go.y:151
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 16:
		//line parser.go.y:155
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 17:
		//line parser.go.y:163
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 18:
		//line parser.go.y:168
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:173
		{
			yyVAL.stmt_continue = &ast.ContinueStmt{}
		}
	case 20:
		//line parser.go.y:178
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:182
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:186
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:190
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:195
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 25:
		//line parser.go.y:200
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 26:
		//line parser.go.y:204
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 27:
		//line parser.go.y:208
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 28:
		//line parser.go.y:213
		{
			yyVAL.idents = []string{}
		}
	case 29:
		//line parser.go.y:217
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 30:
		//line parser.go.y:221
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 31:
		//line parser.go.y:226
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 32:
		//line parser.go.y:230
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 33:
		//line parser.go.y:234
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 34:
		//line parser.go.y:238
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 35:
		//line parser.go.y:243
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 36:
		//line parser.go.y:247
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 37:
		//line parser.go.y:251
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 38:
		//line parser.go.y:255
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 39:
		//line parser.go.y:259
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 40:
		//line parser.go.y:263
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 41:
		//line parser.go.y:267
		{
			yyVAL.expr = &ast.ConstExpr{Value: (*interface{})(nil)}
		}
	case 42:
		//line parser.go.y:271
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	case 43:
		//line parser.go.y:275
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 44:
		//line parser.go.y:279
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 45:
		//line parser.go.y:283
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 46:
		//line parser.go.y:287
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 47:
		//line parser.go.y:291
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 48:
		//line parser.go.y:295
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 49:
		//line parser.go.y:299
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 50:
		//line parser.go.y:307
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 51:
		//line parser.go.y:311
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 52:
		//line parser.go.y:315
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 53:
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 54:
		//line parser.go.y:323
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:335
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:351
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:355
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 63:
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 64:
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.LetExpr{Names: yyS[yypt-2].idents, Operator: "=", Exprs: yyS[yypt-0].exprs}
		}
	case 65:
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "+", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 66:
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "-", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 67:
		//line parser.go.y:375
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "*", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 68:
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "/", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 69:
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 70:
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 71:
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 72:
		//line parser.go.y:395
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 73:
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 74:
		//line parser.go.y:403
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	}
	goto yystack /* stack new state and value */
}
