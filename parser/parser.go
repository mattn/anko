//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
	"unsafe"
)

//line parser.go.y:23
type yySymType struct {
	yys                    int
	stmt_var               ast.Stmt
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
const BREAK = 57377
const CONTINUE = 57378
const PLUSPLUS = 57379
const MINUSMINUS = 57380
const POW = 57381
const UNARY = 57382

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
	"PLUSPLUS",
	"MINUSMINUS",
	"POW",
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

//line parser.go.y:413

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 27,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	40, 27,
	56, 27,
	-2, 1,
	-1, 3,
	40, 27,
	56, 27,
	-2, 1,
	-1, 4,
	40, 27,
	56, 27,
	-2, 1,
	-1, 5,
	40, 27,
	56, 27,
	-2, 1,
	-1, 6,
	40, 27,
	56, 27,
	-2, 1,
	-1, 7,
	40, 27,
	56, 27,
	-2, 1,
	-1, 8,
	40, 27,
	56, 27,
	-2, 1,
	-1, 18,
	40, 28,
	56, 28,
	-2, 35,
	-1, 24,
	58, 30,
	-2, 27,
	-1, 31,
	40, 27,
	56, 27,
	-2, 1,
	-1, 58,
	55, 30,
	-2, 27,
	-1, 66,
	53, 1,
	-2, 27,
	-1, 73,
	55, 30,
	-2, 27,
	-1, 77,
	40, 28,
	-2, 35,
	-1, 86,
	40, 27,
	56, 27,
	-2, 30,
	-1, 88,
	53, 1,
	-2, 27,
	-1, 99,
	17, 0,
	18, 0,
	-2, 58,
	-1, 100,
	17, 0,
	18, 0,
	-2, 59,
	-1, 110,
	53, 1,
	-2, 27,
	-1, 111,
	40, 27,
	56, 27,
	-2, 30,
	-1, 121,
	40, 27,
	56, 27,
	-2, 30,
	-1, 122,
	40, 27,
	56, 27,
	-2, 30,
	-1, 130,
	55, 30,
	-2, 27,
	-1, 157,
	53, 1,
	-2, 27,
	-1, 158,
	53, 1,
	-2, 27,
	-1, 160,
	53, 1,
	-2, 27,
	-1, 169,
	53, 1,
	-2, 27,
	-1, 171,
	53, 1,
	-2, 27,
	-1, 172,
	53, 1,
	-2, 27,
	-1, 174,
	53, 1,
	-2, 27,
	-1, 182,
	53, 1,
	-2, 27,
	-1, 190,
	53, 1,
	-2, 27,
	-1, 194,
	53, 1,
	-2, 27,
	-1, 199,
	53, 1,
	-2, 27,
}

const yyNprod = 77
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 794

var yyAct = []int{

	1, 9, 29, 30, 32, 33, 34, 35, 37, 38,
	120, 81, 59, 60, 184, 47, 62, 75, 162, 85,
	173, 74, 111, 44, 45, 46, 76, 146, 85, 83,
	58, 127, 87, 40, 126, 41, 54, 56, 85, 164,
	161, 90, 91, 143, 93, 94, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 104, 105, 106, 107, 108,
	76, 169, 86, 168, 79, 137, 112, 114, 134, 115,
	116, 117, 118, 130, 125, 76, 109, 47, 85, 64,
	201, 123, 198, 42, 43, 44, 45, 46, 76, 133,
	195, 119, 58, 192, 189, 40, 187, 41, 54, 56,
	186, 185, 179, 176, 132, 175, 156, 153, 142, 199,
	194, 138, 197, 76, 78, 141, 190, 89, 182, 174,
	172, 171, 160, 76, 76, 51, 53, 157, 148, 139,
	151, 110, 76, 66, 128, 191, 154, 155, 150, 144,
	145, 159, 113, 163, 147, 47, 82, 177, 152, 50,
	52, 42, 43, 44, 45, 46, 88, 149, 166, 167,
	58, 170, 131, 40, 80, 41, 54, 56, 124, 92,
	178, 84, 180, 181, 65, 183, 63, 48, 49, 51,
	53, 55, 57, 188, 61, 8, 7, 67, 68, 69,
	70, 193, 6, 71, 72, 196, 5, 2, 0, 47,
	200, 39, 0, 50, 52, 42, 43, 44, 45, 46,
	73, 0, 122, 47, 58, 165, 0, 40, 0, 41,
	54, 56, 48, 49, 51, 53, 55, 57, 58, 0,
	0, 40, 0, 41, 54, 56, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 0, 39, 0, 50, 52,
	42, 43, 44, 45, 46, 0, 0, 158, 0, 58,
	0, 0, 40, 0, 41, 54, 56, 48, 49, 51,
	53, 55, 57, 0, 0, 0, 0, 67, 68, 69,
	70, 0, 0, 71, 72, 0, 0, 0, 0, 47,
	0, 39, 0, 50, 52, 42, 43, 44, 45, 46,
	73, 0, 0, 0, 58, 140, 0, 40, 0, 41,
	54, 56, 48, 49, 51, 53, 55, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 0, 39, 0, 50, 52,
	42, 43, 44, 45, 46, 0, 0, 0, 0, 58,
	0, 0, 40, 136, 41, 54, 56, 48, 49, 51,
	53, 55, 57, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 47,
	0, 39, 135, 50, 52, 42, 43, 44, 45, 46,
	0, 0, 0, 0, 58, 0, 0, 40, 0, 41,
	54, 56, 48, 49, 51, 53, 55, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 0, 39, 0, 50, 52,
	42, 43, 44, 45, 46, 0, 0, 0, 0, 58,
	129, 0, 40, 0, 41, 54, 56, 48, 49, 51,
	53, 55, 57, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 47,
	0, 39, 0, 50, 52, 42, 43, 44, 45, 46,
	0, 0, 0, 0, 58, 0, 121, 40, 0, 41,
	54, 56, 48, 49, 51, 53, 55, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 0, 39, 0, 50, 52,
	42, 43, 44, 45, 46, 48, 49, 51, 53, 58,
	57, 0, 40, 0, 41, 54, 56, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 47, 0, 0,
	0, 50, 52, 42, 43, 44, 45, 46, 0, 0,
	0, 0, 58, 0, 0, 40, 0, 41, 54, 56,
	18, 17, 20, 0, 0, 25, 10, 13, 11, 14,
	36, 15, 0, 0, 0, 0, 0, 0, 0, 28,
	21, 22, 23, 12, 16, 0, 0, 0, 0, 0,
	0, 3, 4, 0, 0, 0, 0, 0, 18, 17,
	20, 0, 19, 25, 10, 13, 11, 14, 26, 15,
	27, 0, 0, 24, 0, 0, 0, 28, 21, 22,
	23, 12, 16, 0, 0, 0, 0, 0, 0, 3,
	4, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	19, 0, 0, 0, 0, 31, 26, 0, 27, 0,
	0, 24, 18, 17, 20, 0, 0, 25, 10, 13,
	11, 14, 0, 15, 0, 0, 0, 0, 0, 0,
	0, 28, 21, 22, 23, 12, 16, 0, 0, 0,
	0, 0, 0, 3, 4, 48, 49, 51, 53, 0,
	0, 0, 0, 0, 19, 0, 0, 0, 0, 0,
	26, 0, 27, 0, 0, 24, 0, 47, 0, 0,
	0, 50, 52, 42, 43, 44, 45, 46, 18, 17,
	20, 0, 58, 25, 0, 40, 0, 41, 54, 56,
	77, 17, 20, 0, 0, 25, 0, 28, 21, 22,
	23, 0, 0, 0, 0, 0, 0, 0, 0, 28,
	21, 22, 23, 0, 0, 0, 0, 0, 0, 0,
	19, 0, 0, 0, 0, 0, 26, 0, 27, 0,
	0, 24, 19, 0, 0, 0, 0, 0, 26, 0,
	27, 0, 0, 24,
}
var yyPact = []int{

	658, -1000, 604, 658, 658, 658, 566, 658, 658, 475,
	724, 724, 180, 172, 25, 170, 81, -1000, 246, 724,
	-1000, -1000, -1000, -1000, 736, 60, 140, 724, 167, 22,
	-1000, 658, -1000, -1000, -1000, -1000, 104, -1000, -1000, 724,
	724, 165, 724, 724, 724, 724, 724, 724, 724, 724,
	724, 724, 724, 724, 724, 724, 724, 724, 736, 475,
	475, 79, -18, -1000, 724, 126, 658, 724, 724, 724,
	724, -1000, -1000, 736, 174, -48, 430, 156, 164, 20,
	-22, -1000, 92, 385, 19, 158, 736, -1000, 658, 14,
	340, 295, -1000, -24, -24, 174, 174, 174, 475, 106,
	106, 38, 38, 38, 38, 475, 508, 475, 678, 10,
	658, 736, 250, 724, 55, 475, 475, 475, 475, -12,
	-1000, 736, 736, -28, 136, 153, 140, -1000, 724, -1000,
	736, -1000, -1000, 54, 724, 724, -1000, -1000, 53, -1000,
	75, 205, 112, -1000, -1000, -1000, 70, -15, -37, 135,
	-1000, 475, -16, -1000, 160, 475, -1000, 658, 658, 9,
	658, 69, 68, -35, -1000, 67, 52, 50, 143, 658,
	49, 658, 658, 66, 658, -1000, -1000, -41, 48, -1000,
	47, 43, 658, 41, 64, 105, -1000, -1000, 40, -1000,
	658, 58, -1000, 37, 658, 82, 29, 57, -1000, 658,
	27, -1000,
}
var yyPgo = []int{

	0, 0, 197, 196, 192, 186, 185, 1, 17, 11,
	164, 2,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 3, 4, 4, 4, 5, 6,
	6, 6, 6, 9, 10, 10, 10, 11, 11, 11,
	8, 8, 8, 8, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 5, 4, 7, 5, 9, 7, 15,
	12, 11, 8, 3, 0, 1, 3, 0, 1, 3,
	0, 1, 3, 3, 1, 1, 2, 1, 1, 1,
	1, 5, 4, 3, 7, 8, 8, 9, 3, 3,
	3, 5, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 3, 3, 3, 3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, 35, 36, -3, -4, -5, -6, -7,
	10, 12, 27, 11, 13, 15, 28, 5, 4, 46,
	6, 24, 25, 26, 57, 9, 52, 54, 23, -11,
	-1, 51, -1, -1, -1, -1, 14, -1, -1, 41,
	57, 59, 45, 46, 47, 48, 49, 39, 17, 18,
	43, 19, 44, 20, 60, 21, 61, 22, 54, -7,
	-7, 4, -11, 4, 54, 4, 52, 31, 32, 33,
	34, 37, 38, 54, -7, -8, -7, 4, 54, 4,
	-10, -9, 6, -7, 4, 56, 40, -1, 52, 13,
	-7, -7, 4, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -8,
	52, 40, -7, 16, -1, -7, -7, -7, -7, -8,
	58, 56, 56, -11, 4, 54, 56, 53, 42, 55,
	54, 4, -8, -1, 54, 42, 58, 55, -1, -8,
	55, -7, 53, 55, -8, -8, 55, 8, -11, 4,
	-9, -7, -8, 53, -7, -7, 53, 52, 52, 29,
	52, 55, 55, 8, 55, 55, -1, -1, 54, 52,
	-1, 52, 52, 55, 52, 53, 53, 4, -1, 53,
	-1, -1, 52, -1, 55, 53, 53, 53, -1, 53,
	52, 30, 53, -1, 52, 53, -1, 30, 53, 52,
	-1, 53,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, 10,
	27, 27, 0, 27, 0, 0, 0, 34, -2, 27,
	37, 38, 39, 40, -2, 0, 24, 27, 0, 0,
	2, -2, 4, 5, 6, 7, 0, 8, 9, 27,
	27, 0, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, -2, 11,
	12, 0, 0, 28, 27, 0, -2, 27, 27, 27,
	27, 69, 70, -2, 36, 0, 31, -2, 27, 0,
	0, 25, 0, 0, 0, 0, -2, 3, -2, 0,
	0, 0, 49, 52, 53, 54, 55, 56, 57, -2,
	-2, 60, 61, 62, 63, 71, 72, 73, 74, 0,
	-2, -2, 0, 27, 0, 65, 66, 67, 68, 0,
	43, -2, -2, 0, 28, 27, 0, 48, 27, 50,
	-2, 29, 64, 0, 27, 27, 42, 76, 0, 14,
	0, 0, 0, 75, 32, 33, 0, 0, 0, 28,
	26, 23, 0, 16, 0, 41, 13, -2, -2, 0,
	-2, 0, 0, 0, 51, 0, 0, 0, 0, -2,
	0, -2, -2, 0, -2, 15, 18, 0, 0, 44,
	0, 0, -2, 0, 0, 22, 45, 46, 0, 17,
	-2, 0, 47, 0, -2, 21, 0, 0, 20, -2,
	0, 19,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 49, 61, 3,
	54, 55, 47, 45, 56, 46, 59, 48, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 42, 51,
	44, 40, 43, 41, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 57, 3, 58, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 52, 60, 53,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 50,
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
		//line parser.go.y:55
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:62
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:69
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:76
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.BreakStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:83
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.ContinueStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:90
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:97
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:104
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 9:
		//line parser.go.y:111
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 10:
		//line parser.go.y:119
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:126
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 12:
		//line parser.go.y:133
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 13:
		//line parser.go.y:140
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:145
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 15:
		//line parser.go.y:150
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 16:
		//line parser.go.y:154
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 17:
		//line parser.go.y:162
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 18:
		//line parser.go.y:167
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:172
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:176
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:180
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:184
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:189
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 24:
		//line parser.go.y:194
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 25:
		//line parser.go.y:198
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 26:
		//line parser.go.y:202
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 27:
		//line parser.go.y:207
		{
			yyVAL.idents = []string{}
		}
	case 28:
		//line parser.go.y:211
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 29:
		//line parser.go.y:215
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 30:
		//line parser.go.y:220
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 31:
		//line parser.go.y:224
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 32:
		//line parser.go.y:228
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 33:
		//line parser.go.y:232
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 34:
		//line parser.go.y:237
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 35:
		//line parser.go.y:241
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 36:
		//line parser.go.y:245
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 37:
		//line parser.go.y:249
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 38:
		//line parser.go.y:253
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 39:
		//line parser.go.y:257
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 40:
		//line parser.go.y:261
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
		}
	case 41:
		//line parser.go.y:265
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	case 42:
		//line parser.go.y:269
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 43:
		//line parser.go.y:273
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 44:
		//line parser.go.y:277
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 45:
		//line parser.go.y:281
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 46:
		//line parser.go.y:285
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 47:
		//line parser.go.y:289
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 48:
		//line parser.go.y:293
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 49:
		//line parser.go.y:301
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 50:
		//line parser.go.y:305
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 51:
		//line parser.go.y:309
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 52:
		//line parser.go.y:313
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 53:
		//line parser.go.y:317
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 54:
		//line parser.go.y:321
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:325
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:329
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:333
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:337
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:341
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:345
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:353
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 63:
		//line parser.go.y:357
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 64:
		//line parser.go.y:361
		{
			yyVAL.expr = &ast.LetExpr{Names: yyS[yypt-2].idents, Operator: "=", Exprs: yyS[yypt-0].exprs}
		}
	case 65:
		//line parser.go.y:365
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "+", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 66:
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "-", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 67:
		//line parser.go.y:373
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "*", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 68:
		//line parser.go.y:377
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "/", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 69:
		//line parser.go.y:381
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
		}
	case 70:
		//line parser.go.y:385
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
		}
	case 71:
		//line parser.go.y:389
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 72:
		//line parser.go.y:393
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 73:
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 74:
		//line parser.go.y:401
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 75:
		//line parser.go.y:405
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 76:
		//line parser.go.y:409
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	}
	goto yystack /* stack new state and value */
}
