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

//line parser.go.y:423

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 29,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	40, 29,
	56, 29,
	-2, 1,
	-1, 3,
	40, 29,
	56, 29,
	-2, 1,
	-1, 4,
	40, 29,
	56, 29,
	-2, 1,
	-1, 5,
	40, 29,
	56, 29,
	-2, 1,
	-1, 6,
	40, 29,
	56, 29,
	-2, 1,
	-1, 7,
	40, 29,
	56, 29,
	-2, 1,
	-1, 8,
	40, 29,
	56, 29,
	-2, 1,
	-1, 10,
	40, 29,
	56, 29,
	-2, 32,
	-1, 18,
	40, 30,
	56, 30,
	-2, 37,
	-1, 24,
	58, 32,
	-2, 29,
	-1, 31,
	40, 29,
	56, 29,
	-2, 1,
	-1, 58,
	55, 32,
	-2, 29,
	-1, 61,
	40, 30,
	-2, 37,
	-1, 68,
	53, 1,
	-2, 29,
	-1, 70,
	53, 1,
	-2, 29,
	-1, 77,
	55, 32,
	-2, 29,
	-1, 88,
	40, 29,
	56, 29,
	-2, 32,
	-1, 91,
	53, 1,
	-2, 29,
	-1, 101,
	17, 0,
	18, 0,
	-2, 60,
	-1, 102,
	17, 0,
	18, 0,
	-2, 61,
	-1, 112,
	40, 29,
	56, 29,
	-2, 32,
	-1, 113,
	40, 29,
	56, 29,
	-2, 32,
	-1, 114,
	53, 1,
	-2, 29,
	-1, 115,
	40, 29,
	56, 29,
	-2, 32,
	-1, 134,
	55, 32,
	-2, 29,
	-1, 163,
	53, 1,
	-2, 29,
	-1, 164,
	53, 1,
	-2, 29,
	-1, 165,
	53, 1,
	-2, 29,
	-1, 167,
	53, 1,
	-2, 29,
	-1, 177,
	53, 1,
	-2, 29,
	-1, 179,
	53, 1,
	-2, 29,
	-1, 180,
	53, 1,
	-2, 29,
	-1, 182,
	53, 1,
	-2, 29,
	-1, 191,
	53, 1,
	-2, 29,
	-1, 199,
	53, 1,
	-2, 29,
	-1, 203,
	53, 1,
	-2, 29,
	-1, 208,
	53, 1,
	-2, 29,
}

const yyNprod = 79
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 846

var yyAct = []int{

	1, 126, 83, 30, 32, 33, 34, 35, 37, 38,
	169, 87, 152, 87, 29, 131, 193, 137, 130, 181,
	177, 59, 176, 134, 71, 72, 73, 74, 64, 115,
	75, 76, 89, 47, 171, 48, 49, 51, 53, 55,
	57, 44, 45, 46, 88, 87, 79, 77, 58, 113,
	168, 40, 67, 41, 54, 56, 151, 47, 141, 39,
	87, 50, 52, 42, 43, 44, 45, 46, 47, 118,
	164, 120, 58, 9, 129, 40, 66, 41, 54, 56,
	111, 210, 208, 58, 60, 62, 40, 207, 41, 54,
	56, 81, 138, 78, 204, 127, 201, 198, 60, 125,
	68, 85, 69, 196, 195, 194, 188, 185, 184, 183,
	136, 162, 160, 92, 93, 144, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 104, 105, 106, 107, 108,
	109, 110, 60, 156, 142, 143, 150, 145, 148, 132,
	116, 80, 203, 119, 154, 121, 122, 123, 124, 90,
	199, 60, 71, 72, 73, 74, 158, 191, 75, 76,
	182, 180, 60, 179, 173, 174, 175, 167, 178, 165,
	163, 114, 70, 206, 200, 77, 166, 117, 187, 170,
	189, 190, 153, 192, 84, 186, 60, 60, 91, 60,
	155, 147, 197, 48, 49, 51, 53, 55, 57, 135,
	202, 128, 94, 86, 205, 65, 157, 63, 60, 209,
	82, 159, 8, 161, 7, 47, 6, 39, 5, 50,
	52, 42, 43, 44, 45, 46, 2, 0, 0, 0,
	58, 172, 0, 40, 0, 41, 54, 56, 48, 49,
	51, 53, 55, 57, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	47, 0, 39, 0, 50, 52, 42, 43, 44, 45,
	46, 0, 0, 0, 0, 58, 149, 0, 40, 0,
	41, 54, 56, 48, 49, 51, 53, 55, 57, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 47, 0, 39, 0, 50,
	52, 42, 43, 44, 45, 46, 0, 0, 0, 0,
	58, 146, 0, 40, 0, 41, 54, 56, 48, 49,
	51, 53, 55, 57, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	47, 0, 39, 0, 50, 52, 42, 43, 44, 45,
	46, 0, 0, 0, 0, 58, 0, 0, 40, 140,
	41, 54, 56, 48, 49, 51, 53, 55, 57, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 47, 0, 39, 139, 50,
	52, 42, 43, 44, 45, 46, 0, 0, 0, 0,
	58, 0, 0, 40, 0, 41, 54, 56, 48, 49,
	51, 53, 55, 57, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	47, 0, 39, 0, 50, 52, 42, 43, 44, 45,
	46, 0, 0, 0, 0, 58, 133, 0, 40, 0,
	41, 54, 56, 48, 49, 51, 53, 55, 57, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 47, 0, 39, 0, 50,
	52, 42, 43, 44, 45, 46, 0, 0, 0, 0,
	58, 0, 112, 40, 0, 41, 54, 56, 48, 49,
	51, 53, 55, 57, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	47, 0, 39, 0, 50, 52, 42, 43, 44, 45,
	46, 48, 49, 51, 53, 58, 57, 0, 40, 0,
	41, 54, 56, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 47, 0, 0, 0, 50, 52, 42,
	43, 44, 45, 46, 0, 0, 0, 0, 58, 0,
	0, 40, 0, 41, 54, 56, 18, 17, 20, 0,
	0, 25, 10, 13, 11, 14, 36, 15, 0, 0,
	0, 0, 0, 0, 0, 28, 21, 22, 23, 12,
	16, 0, 0, 0, 0, 0, 0, 3, 4, 0,
	0, 0, 0, 0, 18, 17, 20, 0, 19, 25,
	10, 13, 11, 14, 26, 15, 27, 0, 0, 24,
	0, 0, 0, 28, 21, 22, 23, 12, 16, 0,
	0, 0, 0, 0, 0, 3, 4, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 19, 0, 0, 0,
	0, 31, 26, 0, 27, 0, 0, 24, 18, 17,
	20, 0, 0, 25, 10, 13, 11, 14, 0, 15,
	0, 0, 0, 0, 0, 0, 0, 28, 21, 22,
	23, 12, 16, 0, 0, 0, 0, 0, 0, 3,
	4, 48, 49, 51, 53, 0, 0, 0, 0, 0,
	19, 0, 0, 0, 0, 0, 26, 0, 27, 0,
	0, 24, 0, 47, 0, 0, 0, 50, 52, 42,
	43, 44, 45, 46, 51, 53, 0, 0, 58, 0,
	0, 40, 0, 41, 54, 56, 18, 17, 20, 0,
	0, 25, 0, 0, 47, 0, 0, 0, 50, 52,
	42, 43, 44, 45, 46, 28, 21, 22, 23, 58,
	0, 47, 40, 0, 41, 54, 56, 42, 43, 44,
	45, 46, 61, 17, 20, 0, 58, 25, 19, 40,
	0, 41, 54, 56, 26, 0, 27, 0, 0, 24,
	0, 28, 21, 22, 23, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 19, 0, 0, 0, 0, 0,
	26, 0, 27, 0, 0, 24,
}
var yyPact = []int{

	674, -1000, 620, 674, 674, 674, 582, 674, 674, 491,
	788, 752, 203, 201, 22, 48, 120, -1000, 121, 752,
	-1000, -1000, -1000, -1000, 788, 87, 178, 752, 199, 4,
	-1000, 674, -1000, -1000, -1000, -1000, 136, -1000, -1000, 752,
	752, 198, 752, 752, 752, 752, 752, 752, 752, 752,
	752, 752, 752, 752, 752, 752, 752, 752, 788, -1000,
	446, -7, 491, 119, -11, -1000, 752, 161, 674, 752,
	674, 752, 752, 752, 752, -1000, -1000, 788, 29, -57,
	197, 20, -38, -1000, 97, 401, -31, 195, 788, -1000,
	-37, 674, 356, 311, -1000, -6, -6, 29, 29, 29,
	491, 725, 725, 742, 742, 742, 742, 491, 524, 491,
	694, 3, 788, 788, 674, 788, 266, 752, 85, 221,
	83, 491, 491, 491, 491, 1, -1000, -43, 174, 186,
	178, -1000, 752, -1000, 788, -1000, -1000, 752, 59, 752,
	-1000, -1000, -1000, -1000, 58, -1000, 118, 18, -1000, 117,
	147, -1000, 115, -5, -45, 171, -1000, 491, -21, 176,
	-1000, 491, -1000, 674, 674, 674, -32, 674, 111, 109,
	-36, -1000, 108, 56, 55, 54, 181, 674, 53, 674,
	674, 105, 674, -1000, -1000, -1000, -39, 52, -1000, 51,
	50, 674, 44, 98, 144, -1000, -1000, 43, -1000, 674,
	90, -1000, 41, 674, 143, 34, 30, -1000, 674, 28,
	-1000,
}
var yyPgo = []int{

	0, 0, 226, 218, 216, 214, 212, 73, 21, 2,
	210, 14,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 3, 4, 4, 4, 5, 5,
	5, 6, 6, 6, 6, 9, 10, 10, 10, 11,
	11, 11, 8, 8, 8, 8, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 5, 4, 9, 5, 7, 7, 4,
	7, 15, 12, 11, 8, 3, 0, 1, 3, 0,
	1, 3, 0, 1, 3, 3, 1, 1, 2, 1,
	1, 1, 1, 5, 4, 3, 7, 8, 8, 9,
	3, 3, 3, 5, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 3, 3, 3, 3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, 35, 36, -3, -4, -5, -6, -7,
	10, 12, 27, 11, 13, 15, 28, 5, 4, 46,
	6, 24, 25, 26, 57, 9, 52, 54, 23, -11,
	-1, 51, -1, -1, -1, -1, 14, -1, -1, 41,
	57, 59, 45, 46, 47, 48, 49, 39, 17, 18,
	43, 19, 44, 20, 60, 21, 61, 22, 54, -8,
	-7, 4, -7, 4, -11, 4, 54, 4, 52, 54,
	52, 31, 32, 33, 34, 37, 38, 54, -7, -8,
	54, 4, -10, -9, 6, -7, 4, 56, 40, -1,
	13, 52, -7, -7, 4, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -8, 56, 56, 52, 40, -7, 16, -1, -7,
	-1, -7, -7, -7, -7, -8, 58, -11, 4, 54,
	56, 53, 42, 55, 54, 4, -8, 54, -1, 42,
	58, 55, -8, -8, -1, -8, 55, -7, 53, 55,
	53, 55, 55, 8, -11, 4, -9, -7, -8, -7,
	53, -7, 53, 52, 52, 52, 29, 52, 55, 55,
	8, 55, 55, -1, -1, -1, 54, 52, -1, 52,
	52, 55, 52, 53, 53, 53, 4, -1, 53, -1,
	-1, 52, -1, 55, 53, 53, 53, -1, 53, 52,
	30, 53, -1, 52, 53, -1, 30, 53, 52, -1,
	53,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, 10,
	-2, 29, 0, 29, 0, 0, 0, 36, -2, 29,
	39, 40, 41, 42, -2, 0, 26, 29, 0, 0,
	2, -2, 4, 5, 6, 7, 0, 8, 9, 29,
	29, 0, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, -2, 11,
	33, -2, 12, 0, 0, 30, 29, 0, -2, 29,
	-2, 29, 29, 29, 29, 71, 72, -2, 38, 0,
	29, 0, 0, 27, 0, 0, 0, 0, -2, 3,
	0, -2, 0, 0, 51, 54, 55, 56, 57, 58,
	59, -2, -2, 62, 63, 64, 65, 73, 74, 75,
	76, 0, -2, -2, -2, -2, 0, 29, 0, 0,
	0, 67, 68, 69, 70, 0, 45, 0, 30, 29,
	0, 50, 29, 52, -2, 31, 66, 29, 0, 29,
	44, 78, 34, 35, 0, 14, 0, 0, 19, 0,
	0, 77, 0, 0, 0, 30, 28, 25, 0, 0,
	16, 43, 13, -2, -2, -2, 0, -2, 0, 0,
	0, 53, 0, 0, 0, 0, 0, -2, 0, -2,
	-2, 0, -2, 17, 18, 20, 0, 0, 46, 0,
	0, -2, 0, 0, 24, 47, 48, 0, 15, -2,
	0, 49, 0, -2, 23, 0, 0, 22, -2, 0,
	21,
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
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				if len(yyS[yypt-0].exprs) > 0 {
					yyS[yypt-0].exprs[0].SetPos(l.pos)
				}
			}
		}
	case 12:
		//line parser.go.y:135
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 13:
		//line parser.go.y:142
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:147
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 15:
		//line parser.go.y:152
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 16:
		//line parser.go.y:156
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 17:
		//line parser.go.y:164
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:169
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:173
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:177
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:182
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:186
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:190
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:194
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:199
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:204
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 27:
		//line parser.go.y:208
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 28:
		//line parser.go.y:212
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 29:
		//line parser.go.y:217
		{
			yyVAL.idents = []string{}
		}
	case 30:
		//line parser.go.y:221
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 31:
		//line parser.go.y:225
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 32:
		//line parser.go.y:230
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 33:
		//line parser.go.y:234
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 34:
		//line parser.go.y:238
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 35:
		//line parser.go.y:242
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 36:
		//line parser.go.y:247
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 37:
		//line parser.go.y:251
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 38:
		//line parser.go.y:255
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 39:
		//line parser.go.y:259
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 40:
		//line parser.go.y:263
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 41:
		//line parser.go.y:267
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 42:
		//line parser.go.y:271
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
		}
	case 43:
		//line parser.go.y:275
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	case 44:
		//line parser.go.y:279
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 45:
		//line parser.go.y:283
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 46:
		//line parser.go.y:287
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 47:
		//line parser.go.y:291
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 48:
		//line parser.go.y:295
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 49:
		//line parser.go.y:299
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 50:
		//line parser.go.y:303
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 51:
		//line parser.go.y:311
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 52:
		//line parser.go.y:315
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 53:
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 54:
		//line parser.go.y:323
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:335
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:351
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:355
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 63:
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 64:
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 65:
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 66:
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.LetExpr{Names: yyS[yypt-2].idents, Operator: "=", Exprs: yyS[yypt-0].exprs}
		}
	case 67:
		//line parser.go.y:375
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "+", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 68:
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "-", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 69:
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "*", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 70:
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "/", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 71:
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
		}
	case 72:
		//line parser.go.y:395
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
		}
	case 73:
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 74:
		//line parser.go.y:403
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 75:
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 76:
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 77:
		//line parser.go.y:415
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 78:
		//line parser.go.y:419
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	}
	goto yystack /* stack new state and value */
}
