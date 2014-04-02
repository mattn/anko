//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:24
type yySymType struct {
	yys                    int
	stmt_var               ast.Stmt
	stmt_if                ast.Stmt
	stmt_for               ast.Stmt
	stmt_break             ast.Stmt
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

//line parser.go.y:414

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 29,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	37, 29,
	53, 29,
	-2, 1,
	-1, 3,
	37, 29,
	53, 29,
	-2, 1,
	-1, 4,
	37, 29,
	53, 29,
	-2, 1,
	-1, 5,
	37, 29,
	53, 29,
	-2, 1,
	-1, 6,
	37, 29,
	53, 29,
	-2, 1,
	-1, 7,
	37, 29,
	53, 29,
	-2, 1,
	-1, 8,
	37, 29,
	53, 29,
	-2, 1,
	-1, 20,
	37, 30,
	53, 30,
	-2, 37,
	-1, 26,
	55, 32,
	-2, 29,
	-1, 33,
	37, 29,
	53, 29,
	-2, 1,
	-1, 59,
	52, 32,
	-2, 29,
	-1, 67,
	50, 1,
	-2, 29,
	-1, 72,
	52, 32,
	-2, 29,
	-1, 76,
	37, 30,
	-2, 37,
	-1, 85,
	37, 29,
	53, 29,
	-2, 32,
	-1, 87,
	50, 1,
	-2, 29,
	-1, 97,
	17, 0,
	18, 0,
	-2, 59,
	-1, 98,
	17, 0,
	18, 0,
	-2, 60,
	-1, 108,
	50, 1,
	-2, 29,
	-1, 109,
	37, 29,
	53, 29,
	-2, 32,
	-1, 119,
	37, 29,
	53, 29,
	-2, 32,
	-1, 120,
	37, 29,
	53, 29,
	-2, 32,
	-1, 128,
	52, 32,
	-2, 29,
	-1, 155,
	50, 1,
	-2, 29,
	-1, 156,
	50, 1,
	-2, 29,
	-1, 158,
	50, 1,
	-2, 29,
	-1, 167,
	50, 1,
	-2, 29,
	-1, 169,
	50, 1,
	-2, 29,
	-1, 170,
	50, 1,
	-2, 29,
	-1, 172,
	50, 1,
	-2, 29,
	-1, 180,
	50, 1,
	-2, 29,
	-1, 188,
	50, 1,
	-2, 29,
	-1, 192,
	50, 1,
	-2, 29,
	-1, 197,
	50, 1,
	-2, 29,
}

const yyNprod = 76
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 698

var yyAct = []int{

	1, 9, 80, 32, 34, 35, 36, 37, 39, 40,
	118, 182, 60, 61, 31, 59, 74, 171, 42, 162,
	43, 55, 57, 73, 159, 46, 47, 48, 75, 109,
	63, 82, 59, 125, 86, 42, 124, 43, 55, 57,
	160, 84, 141, 89, 90, 84, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 101, 102, 103, 104, 105,
	106, 75, 144, 84, 85, 135, 132, 110, 112, 78,
	113, 114, 115, 116, 75, 167, 107, 166, 197, 128,
	84, 44, 45, 46, 47, 48, 123, 75, 131, 117,
	59, 65, 121, 42, 199, 43, 55, 57, 68, 69,
	70, 71, 130, 68, 69, 70, 71, 196, 193, 136,
	190, 75, 187, 139, 185, 88, 77, 192, 72, 184,
	120, 75, 75, 72, 183, 177, 137, 148, 149, 174,
	75, 173, 154, 151, 152, 153, 142, 143, 146, 140,
	52, 54, 188, 180, 172, 150, 170, 169, 158, 155,
	108, 87, 67, 126, 195, 189, 164, 165, 157, 168,
	111, 51, 53, 44, 45, 46, 47, 48, 176, 161,
	178, 179, 59, 181, 145, 42, 81, 43, 55, 57,
	175, 186, 49, 50, 52, 54, 56, 58, 147, 191,
	129, 122, 91, 194, 83, 66, 64, 62, 198, 79,
	8, 4, 3, 41, 7, 51, 53, 44, 45, 46,
	47, 48, 6, 5, 2, 0, 59, 163, 0, 42,
	0, 43, 55, 57, 49, 50, 52, 54, 56, 58,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 41, 0, 51, 53, 44,
	45, 46, 47, 48, 0, 0, 156, 0, 59, 0,
	0, 42, 0, 43, 55, 57, 49, 50, 52, 54,
	56, 58, 0, 0, 0, 0, 0, 20, 19, 22,
	0, 0, 27, 0, 0, 0, 0, 41, 0, 51,
	53, 44, 45, 46, 47, 48, 30, 23, 24, 25,
	59, 138, 0, 42, 0, 43, 55, 57, 49, 50,
	52, 54, 56, 58, 0, 0, 21, 0, 0, 0,
	0, 0, 28, 0, 29, 0, 0, 26, 0, 41,
	0, 51, 53, 44, 45, 46, 47, 48, 0, 0,
	0, 0, 59, 0, 0, 42, 134, 43, 55, 57,
	49, 50, 52, 54, 56, 58, 0, 0, 0, 0,
	0, 76, 19, 22, 0, 0, 27, 0, 0, 0,
	0, 41, 133, 51, 53, 44, 45, 46, 47, 48,
	30, 23, 24, 25, 59, 0, 0, 42, 0, 43,
	55, 57, 49, 50, 52, 54, 56, 58, 0, 0,
	21, 0, 0, 0, 0, 0, 28, 0, 29, 0,
	0, 26, 0, 41, 0, 51, 53, 44, 45, 46,
	47, 48, 0, 0, 0, 0, 59, 127, 0, 42,
	0, 43, 55, 57, 49, 50, 52, 54, 56, 58,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 41, 0, 51, 53, 44,
	45, 46, 47, 48, 0, 0, 0, 0, 59, 0,
	119, 42, 0, 43, 55, 57, 49, 50, 52, 54,
	56, 58, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 41, 0, 51,
	53, 44, 45, 46, 47, 48, 0, 0, 0, 0,
	59, 0, 0, 42, 0, 43, 55, 57, 20, 19,
	22, 0, 0, 27, 10, 15, 11, 16, 38, 17,
	0, 0, 0, 0, 0, 0, 0, 30, 23, 24,
	25, 12, 18, 0, 0, 0, 0, 0, 0, 13,
	14, 0, 0, 20, 19, 22, 0, 21, 27, 10,
	15, 11, 16, 28, 17, 29, 0, 0, 26, 0,
	0, 0, 30, 23, 24, 25, 12, 18, 0, 0,
	0, 0, 0, 0, 13, 14, 0, 0, 0, 0,
	0, 0, 21, 0, 0, 0, 0, 33, 28, 0,
	29, 0, 0, 26, 20, 19, 22, 0, 0, 27,
	10, 15, 11, 16, 0, 17, 0, 0, 0, 0,
	0, 0, 0, 30, 23, 24, 25, 12, 18, 0,
	0, 0, 0, 0, 0, 13, 14, 49, 50, 52,
	54, 0, 58, 21, 0, 0, 0, 0, 0, 28,
	0, 29, 0, 0, 26, 0, 49, 50, 52, 54,
	51, 53, 44, 45, 46, 47, 48, 0, 0, 0,
	0, 59, 0, 0, 42, 0, 43, 55, 57, 51,
	53, 44, 45, 46, 47, 48, 0, 0, 0, 0,
	59, 0, 0, 42, 0, 43, 55, 57,
}
var yyPact = []int{

	600, -1000, 549, 600, 600, 600, 514, 600, 600, 459,
	273, 273, 193, -1000, -1000, 192, 40, 191, 103, -1000,
	72, 273, -1000, -1000, -1000, -1000, 357, 65, 170, 273,
	190, 27, -1000, 600, -1000, -1000, -1000, -1000, 102, -1000,
	-1000, 273, 273, 188, 273, 273, 273, 273, 273, 273,
	273, 273, 273, 273, 273, 273, 273, 273, 273, 357,
	459, 459, 101, -8, -1000, 273, 144, 600, 273, 273,
	273, 273, 357, -36, -45, 417, 67, 187, 35, -17,
	-1000, 114, 375, 28, 186, 357, -1000, 600, 15, 333,
	291, -1000, -19, -19, -36, -36, -36, 121, 121, 39,
	39, 39, 39, 459, 620, 459, 639, 13, 600, 357,
	249, 273, 89, 459, 459, 459, 459, -10, -1000, 357,
	357, 10, 166, 184, 170, -1000, 273, -1000, 357, -1000,
	-1000, 83, 273, 273, -1000, -1000, 82, -1000, 100, 207,
	129, -1000, -1000, -1000, 99, -28, -12, 161, -1000, 459,
	-33, -1000, 165, 459, -1000, 600, 600, 26, 600, 98,
	97, -35, -1000, 95, 81, 79, 176, 600, 75, 600,
	600, 94, 600, -1000, -1000, -41, 74, -1000, 69, 64,
	600, 62, 93, 125, -1000, -1000, 60, -1000, 600, 68,
	-1000, 58, 600, 124, 57, 29, -1000, 600, 44, -1000,
}
var yyPgo = []int{

	0, 0, 214, 213, 212, 204, 202, 201, 200, 1,
	16, 2, 199, 14,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 3, 4, 4, 4, 5, 6,
	7, 8, 8, 8, 8, 11, 12, 12, 12, 13,
	13, 13, 10, 10, 10, 10, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 5, 4, 7, 5, 9, 7, 1,
	1, 15, 12, 11, 8, 3, 0, 1, 3, 0,
	1, 3, 0, 1, 3, 3, 1, 1, 2, 1,
	1, 1, 1, 5, 4, 3, 7, 8, 8, 9,
	3, 3, 3, 5, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, -6, -7, -3, -4, -5, -8, -9,
	10, 12, 27, 35, 36, 11, 13, 15, 28, 5,
	4, 43, 6, 24, 25, 26, 54, 9, 49, 51,
	23, -13, -1, 48, -1, -1, -1, -1, 14, -1,
	-1, 38, 54, 56, 42, 43, 44, 45, 46, 17,
	18, 40, 19, 41, 20, 57, 21, 58, 22, 51,
	-9, -9, 4, -13, 4, 51, 4, 49, 31, 32,
	33, 34, 51, -9, -10, -9, 4, 51, 4, -12,
	-11, 6, -9, 4, 53, 37, -1, 49, 13, -9,
	-9, 4, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -10, 49, 37,
	-9, 16, -1, -9, -9, -9, -9, -10, 55, 53,
	53, -13, 4, 51, 53, 50, 39, 52, 51, 4,
	-10, -1, 51, 39, 55, 52, -1, -10, 52, -9,
	50, 52, -10, -10, 52, 8, -13, 4, -11, -9,
	-10, 50, -9, -9, 50, 49, 49, 29, 49, 52,
	52, 8, 52, 52, -1, -1, 51, 49, -1, 49,
	49, 52, 49, 50, 50, 4, -1, 50, -1, -1,
	49, -1, 52, 50, 50, 50, -1, 50, 49, 30,
	50, -1, 49, 50, -1, 30, 50, 49, -1, 50,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, 10,
	29, 29, 0, 19, 20, 29, 0, 0, 0, 36,
	-2, 29, 39, 40, 41, 42, -2, 0, 26, 29,
	0, 0, 2, -2, 4, 5, 6, 7, 0, 8,
	9, 29, 29, 0, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, -2,
	11, 12, 0, 0, 30, 29, 0, -2, 29, 29,
	29, 29, -2, 38, 0, 33, -2, 29, 0, 0,
	27, 0, 0, 0, 0, -2, 3, -2, 0, 0,
	0, 51, 54, 55, 56, 57, 58, -2, -2, 61,
	62, 63, 64, 70, 71, 72, 73, 0, -2, -2,
	0, 29, 0, 66, 67, 68, 69, 0, 45, -2,
	-2, 0, 30, 29, 0, 50, 29, 52, -2, 31,
	65, 0, 29, 29, 44, 75, 0, 14, 0, 0,
	0, 74, 34, 35, 0, 0, 0, 30, 28, 25,
	0, 16, 0, 43, 13, -2, -2, 0, -2, 0,
	0, 0, 53, 0, 0, 0, 0, -2, 0, -2,
	-2, 0, -2, 15, 18, 0, 0, 46, 0, 0,
	-2, 0, 0, 24, 47, 48, 0, 17, -2, 0,
	49, 0, -2, 23, 0, 0, 22, -2, 0, 21,
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
		//line parser.go.y:58
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:65
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:72
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:79
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_break}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:86
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_continue}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:93
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:100
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:107
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 9:
		//line parser.go.y:114
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 10:
		//line parser.go.y:122
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:129
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 12:
		//line parser.go.y:136
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 13:
		//line parser.go.y:143
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:148
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 15:
		//line parser.go.y:153
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 16:
		//line parser.go.y:157
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 17:
		//line parser.go.y:165
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 18:
		//line parser.go.y:170
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:175
		{
			yyVAL.stmt_break = &ast.BreakStmt{}
		}
	case 20:
		//line parser.go.y:180
		{
			yyVAL.stmt_continue = &ast.ContinueStmt{}
		}
	case 21:
		//line parser.go.y:185
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:189
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:193
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:197
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:202
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:207
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 27:
		//line parser.go.y:211
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 28:
		//line parser.go.y:215
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 29:
		//line parser.go.y:220
		{
			yyVAL.idents = []string{}
		}
	case 30:
		//line parser.go.y:224
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 31:
		//line parser.go.y:228
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 32:
		//line parser.go.y:233
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 33:
		//line parser.go.y:237
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 34:
		//line parser.go.y:241
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 35:
		//line parser.go.y:245
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 36:
		//line parser.go.y:250
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 37:
		//line parser.go.y:254
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 38:
		//line parser.go.y:258
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 39:
		//line parser.go.y:262
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 40:
		//line parser.go.y:266
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 41:
		//line parser.go.y:270
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 42:
		//line parser.go.y:274
		{
			yyVAL.expr = &ast.ConstExpr{Value: (*interface{})(nil)}
		}
	case 43:
		//line parser.go.y:278
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	case 44:
		//line parser.go.y:282
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 45:
		//line parser.go.y:286
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 46:
		//line parser.go.y:290
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 47:
		//line parser.go.y:294
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 48:
		//line parser.go.y:298
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 49:
		//line parser.go.y:302
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 50:
		//line parser.go.y:306
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 51:
		//line parser.go.y:314
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 52:
		//line parser.go.y:318
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 53:
		//line parser.go.y:322
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 54:
		//line parser.go.y:326
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:330
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:334
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:338
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:342
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:346
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:350
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:354
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:358
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 63:
		//line parser.go.y:362
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 64:
		//line parser.go.y:366
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 65:
		//line parser.go.y:370
		{
			yyVAL.expr = &ast.LetExpr{Names: yyS[yypt-2].idents, Operator: "=", Exprs: yyS[yypt-0].exprs}
		}
	case 66:
		//line parser.go.y:374
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "+", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 67:
		//line parser.go.y:378
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "-", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 68:
		//line parser.go.y:382
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "*", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 69:
		//line parser.go.y:386
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "/", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 70:
		//line parser.go.y:390
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 71:
		//line parser.go.y:394
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 72:
		//line parser.go.y:398
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 73:
		//line parser.go.y:402
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 74:
		//line parser.go.y:406
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 75:
		//line parser.go.y:410
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	}
	goto yystack /* stack new state and value */
}
