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

//line parser.go.y:421

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
	-1, 66,
	53, 1,
	-2, 29,
	-1, 68,
	53, 1,
	-2, 29,
	-1, 75,
	55, 32,
	-2, 29,
	-1, 79,
	40, 30,
	-2, 37,
	-1, 88,
	40, 29,
	56, 29,
	-2, 32,
	-1, 90,
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
	53, 1,
	-2, 29,
	-1, 113,
	40, 29,
	56, 29,
	-2, 32,
	-1, 125,
	40, 29,
	56, 29,
	-2, 32,
	-1, 126,
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

const yyLast = 856

var yyAct = []int{

	1, 124, 193, 30, 32, 33, 34, 35, 37, 38,
	83, 169, 87, 113, 29, 48, 49, 51, 53, 55,
	57, 152, 87, 131, 88, 138, 130, 181, 62, 87,
	171, 177, 89, 176, 208, 168, 149, 47, 141, 39,
	87, 50, 52, 42, 43, 44, 45, 46, 51, 53,
	134, 129, 58, 172, 64, 40, 210, 41, 54, 56,
	65, 81, 207, 204, 201, 198, 196, 116, 47, 118,
	195, 194, 50, 52, 42, 43, 44, 45, 46, 188,
	185, 9, 47, 58, 77, 184, 40, 183, 41, 54,
	56, 137, 59, 60, 162, 127, 203, 58, 159, 148,
	40, 76, 41, 54, 56, 146, 78, 199, 66, 85,
	67, 80, 191, 142, 182, 180, 91, 179, 167, 165,
	163, 92, 93, 112, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	78, 156, 68, 111, 154, 206, 114, 132, 200, 117,
	166, 119, 120, 121, 122, 90, 115, 78, 170, 153,
	123, 84, 186, 155, 173, 174, 175, 47, 178, 135,
	78, 128, 94, 136, 86, 44, 45, 46, 187, 63,
	189, 190, 58, 192, 61, 40, 82, 41, 54, 56,
	8, 7, 197, 6, 5, 78, 2, 145, 143, 0,
	202, 0, 0, 0, 205, 0, 0, 78, 78, 209,
	150, 151, 0, 0, 157, 0, 78, 0, 0, 158,
	160, 161, 48, 49, 51, 53, 55, 57, 0, 0,
	0, 0, 0, 69, 70, 71, 72, 0, 0, 73,
	74, 0, 0, 0, 47, 0, 39, 0, 50, 52,
	42, 43, 44, 45, 46, 0, 75, 164, 126, 58,
	0, 0, 40, 0, 41, 54, 56, 48, 49, 51,
	53, 55, 57, 0, 0, 0, 0, 69, 70, 71,
	72, 0, 0, 73, 74, 0, 0, 0, 0, 47,
	0, 39, 0, 50, 52, 42, 43, 44, 45, 46,
	75, 0, 0, 0, 58, 147, 0, 40, 0, 41,
	54, 56, 48, 49, 51, 53, 55, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 0, 39, 0, 50, 52,
	42, 43, 44, 45, 46, 0, 0, 0, 0, 58,
	144, 0, 40, 0, 41, 54, 56, 48, 49, 51,
	53, 55, 57, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 47,
	0, 39, 0, 50, 52, 42, 43, 44, 45, 46,
	0, 0, 0, 0, 58, 0, 0, 40, 140, 41,
	54, 56, 48, 49, 51, 53, 55, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 0, 39, 139, 50, 52,
	42, 43, 44, 45, 46, 0, 0, 0, 0, 58,
	0, 0, 40, 0, 41, 54, 56, 48, 49, 51,
	53, 55, 57, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 47,
	0, 39, 0, 50, 52, 42, 43, 44, 45, 46,
	0, 0, 0, 0, 58, 133, 0, 40, 0, 41,
	54, 56, 48, 49, 51, 53, 55, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 0, 39, 0, 50, 52,
	42, 43, 44, 45, 46, 0, 0, 0, 0, 58,
	0, 125, 40, 0, 41, 54, 56, 48, 49, 51,
	53, 55, 57, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 47,
	0, 39, 0, 50, 52, 42, 43, 44, 45, 46,
	48, 49, 51, 53, 58, 57, 0, 40, 0, 41,
	54, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 47, 0, 0, 0, 50, 52, 42, 43,
	44, 45, 46, 0, 0, 0, 0, 58, 0, 0,
	40, 0, 41, 54, 56, 18, 17, 20, 0, 0,
	25, 10, 13, 11, 14, 36, 15, 0, 0, 0,
	0, 0, 0, 0, 28, 21, 22, 23, 12, 16,
	0, 0, 0, 0, 0, 0, 3, 4, 0, 0,
	0, 0, 0, 18, 17, 20, 0, 19, 25, 10,
	13, 11, 14, 26, 15, 27, 0, 0, 24, 0,
	0, 0, 28, 21, 22, 23, 12, 16, 0, 0,
	0, 0, 0, 0, 3, 4, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 19, 0, 0, 0, 0,
	31, 26, 0, 27, 0, 0, 24, 18, 17, 20,
	0, 0, 25, 10, 13, 11, 14, 0, 15, 0,
	0, 0, 0, 0, 0, 0, 28, 21, 22, 23,
	12, 16, 0, 0, 0, 0, 0, 0, 3, 4,
	48, 49, 51, 53, 0, 0, 0, 0, 0, 19,
	0, 0, 0, 0, 0, 26, 0, 27, 0, 0,
	24, 0, 47, 0, 0, 0, 50, 52, 42, 43,
	44, 45, 46, 0, 0, 0, 0, 58, 0, 47,
	40, 0, 41, 54, 56, 42, 43, 44, 45, 46,
	18, 17, 20, 0, 58, 25, 0, 40, 0, 41,
	54, 56, 79, 17, 20, 0, 0, 25, 0, 28,
	21, 22, 23, 0, 0, 0, 0, 0, 0, 0,
	0, 28, 21, 22, 23, 0, 0, 0, 0, 0,
	0, 0, 19, 0, 0, 0, 0, 0, 26, 0,
	27, 0, 0, 24, 19, 0, 0, 0, 0, 0,
	26, 0, 27, 0, 0, 24,
}
var yyPact = []int{

	703, -1000, 649, 703, 703, 703, 611, 703, 703, 520,
	786, 786, 180, 175, 0, 56, 90, -1000, 246, 786,
	-1000, -1000, -1000, -1000, 798, 57, 155, 786, 170, -16,
	-1000, 703, -1000, -1000, -1000, -1000, 103, -1000, -1000, 786,
	786, 168, 786, 786, 786, 786, 786, 786, 786, 786,
	786, 786, 786, 786, 786, 786, 786, 786, 798, 520,
	520, 71, -27, -1000, 786, 140, 703, 786, 703, 786,
	786, 786, 786, -1000, -1000, 798, 43, -57, 475, 202,
	167, -3, -30, -1000, 105, 430, -4, 165, 798, -1000,
	703, -29, 385, 340, -1000, 128, 128, 43, 43, 43,
	520, 29, 29, 740, 740, 740, 740, 520, 553, 520,
	723, -17, 703, 798, 295, 786, 52, 250, 46, 520,
	520, 520, 520, -19, -1000, 798, 798, -34, 151, 159,
	155, -1000, 786, -1000, 798, -1000, -1000, 45, 786, 786,
	-1000, -1000, 41, -1000, 68, 205, -1000, 67, 121, -1000,
	-1000, -1000, 66, -20, -44, 150, -1000, 520, -25, -1000,
	-2, 520, -1000, 703, 703, 703, -21, 703, 65, 63,
	-28, -1000, 62, 34, 32, 27, 158, 703, 26, 703,
	703, 60, 703, -1000, -1000, -1000, -53, 18, -1000, 17,
	13, 703, 12, 55, 118, -1000, -1000, 11, -1000, 703,
	44, -1000, 10, 703, 115, 9, -18, -1000, 703, 3,
	-1000,
}
var yyPgo = []int{

	0, 0, 196, 194, 193, 191, 190, 81, 84, 10,
	186, 14,
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
	1, 2, 2, 5, 4, 7, 5, 9, 7, 4,
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
	43, 19, 44, 20, 60, 21, 61, 22, 54, -7,
	-7, 4, -11, 4, 54, 4, 52, 54, 52, 31,
	32, 33, 34, 37, 38, 54, -7, -8, -7, 4,
	54, 4, -10, -9, 6, -7, 4, 56, 40, -1,
	52, 13, -7, -7, 4, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -8, 52, 40, -7, 16, -1, -7, -1, -7,
	-7, -7, -7, -8, 58, 56, 56, -11, 4, 54,
	56, 53, 42, 55, 54, 4, -8, -1, 54, 42,
	58, 55, -1, -8, 55, -7, 53, 55, 53, 55,
	-8, -8, 55, 8, -11, 4, -9, -7, -8, 53,
	-7, -7, 53, 52, 52, 52, 29, 52, 55, 55,
	8, 55, 55, -1, -1, -1, 54, 52, -1, 52,
	52, 55, 52, 53, 53, 53, 4, -1, 53, -1,
	-1, 52, -1, 55, 53, 53, 53, -1, 53, 52,
	30, 53, -1, 52, 53, -1, 30, 53, 52, -1,
	53,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, 10,
	29, 29, 0, 29, 0, 0, 0, 36, -2, 29,
	39, 40, 41, 42, -2, 0, 26, 29, 0, 0,
	2, -2, 4, 5, 6, 7, 0, 8, 9, 29,
	29, 0, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, -2, 11,
	12, 0, 0, 30, 29, 0, -2, 29, -2, 29,
	29, 29, 29, 71, 72, -2, 38, 0, 33, -2,
	29, 0, 0, 27, 0, 0, 0, 0, -2, 3,
	-2, 0, 0, 0, 51, 54, 55, 56, 57, 58,
	59, -2, -2, 62, 63, 64, 65, 73, 74, 75,
	76, 0, -2, -2, 0, 29, 0, 0, 0, 67,
	68, 69, 70, 0, 45, -2, -2, 0, 30, 29,
	0, 50, 29, 52, -2, 31, 66, 0, 29, 29,
	44, 78, 0, 14, 0, 0, 19, 0, 0, 77,
	34, 35, 0, 0, 0, 30, 28, 25, 0, 16,
	0, 43, 13, -2, -2, -2, 0, -2, 0, 0,
	0, 53, 0, 0, 0, 0, 0, -2, 0, -2,
	-2, 0, -2, 15, 18, 20, 0, 0, 46, 0,
	0, -2, 0, 0, 24, 47, 48, 0, 17, -2,
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
		//line parser.go.y:171
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:175
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:180
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:184
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:188
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:192
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:197
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:202
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 27:
		//line parser.go.y:206
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 28:
		//line parser.go.y:210
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 29:
		//line parser.go.y:215
		{
			yyVAL.idents = []string{}
		}
	case 30:
		//line parser.go.y:219
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 31:
		//line parser.go.y:223
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 32:
		//line parser.go.y:228
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 33:
		//line parser.go.y:232
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 34:
		//line parser.go.y:236
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 35:
		//line parser.go.y:240
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 36:
		//line parser.go.y:245
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 37:
		//line parser.go.y:249
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 38:
		//line parser.go.y:253
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 39:
		//line parser.go.y:257
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 40:
		//line parser.go.y:261
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 41:
		//line parser.go.y:265
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 42:
		//line parser.go.y:269
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
		}
	case 43:
		//line parser.go.y:273
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	case 44:
		//line parser.go.y:277
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 45:
		//line parser.go.y:281
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 46:
		//line parser.go.y:285
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 47:
		//line parser.go.y:289
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 48:
		//line parser.go.y:293
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 49:
		//line parser.go.y:297
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 50:
		//line parser.go.y:301
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 51:
		//line parser.go.y:309
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 52:
		//line parser.go.y:313
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 53:
		//line parser.go.y:317
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 54:
		//line parser.go.y:321
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:325
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:329
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:333
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:337
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:341
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:345
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:353
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 63:
		//line parser.go.y:357
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 64:
		//line parser.go.y:361
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 65:
		//line parser.go.y:365
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 66:
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.LetExpr{Names: yyS[yypt-2].idents, Operator: "=", Exprs: yyS[yypt-0].exprs}
		}
	case 67:
		//line parser.go.y:373
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "+", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 68:
		//line parser.go.y:377
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "-", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 69:
		//line parser.go.y:381
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "*", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 70:
		//line parser.go.y:385
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "/", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 71:
		//line parser.go.y:389
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
		}
	case 72:
		//line parser.go.y:393
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
		}
	case 73:
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 74:
		//line parser.go.y:401
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 75:
		//line parser.go.y:405
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 76:
		//line parser.go.y:409
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 77:
		//line parser.go.y:413
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 78:
		//line parser.go.y:417
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	}
	goto yystack /* stack new state and value */
}
