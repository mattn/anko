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
const ANDEQ = 57377
const OREQ = 57378
const BREAK = 57379
const CONTINUE = 57380
const PLUSPLUS = 57381
const MINUSMINUS = 57382
const POW = 57383
const SHIFTLEFT = 57384
const SHIFTRIGHT = 57385
const UNARY = 57386

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
	"ANDEQ",
	"OREQ",
	"BREAK",
	"CONTINUE",
	"PLUSPLUS",
	"MINUSMINUS",
	"POW",
	"SHIFTLEFT",
	"SHIFTRIGHT",
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

//line parser.go.y:439

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 29,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	44, 29,
	60, 29,
	-2, 1,
	-1, 3,
	44, 29,
	60, 29,
	-2, 1,
	-1, 4,
	44, 29,
	60, 29,
	-2, 1,
	-1, 5,
	44, 29,
	60, 29,
	-2, 1,
	-1, 6,
	44, 29,
	60, 29,
	-2, 1,
	-1, 7,
	44, 29,
	60, 29,
	-2, 1,
	-1, 8,
	44, 29,
	60, 29,
	-2, 1,
	-1, 10,
	44, 29,
	60, 29,
	-2, 32,
	-1, 18,
	44, 30,
	60, 30,
	-2, 37,
	-1, 26,
	64, 32,
	-2, 29,
	-1, 33,
	44, 29,
	60, 29,
	-2, 1,
	-1, 60,
	59, 32,
	-2, 29,
	-1, 63,
	44, 30,
	-2, 37,
	-1, 70,
	57, 1,
	-2, 29,
	-1, 72,
	57, 1,
	-2, 29,
	-1, 81,
	59, 32,
	-2, 29,
	-1, 94,
	44, 29,
	60, 29,
	-2, 32,
	-1, 97,
	57, 1,
	-2, 29,
	-1, 107,
	17, 0,
	18, 0,
	-2, 62,
	-1, 108,
	17, 0,
	18, 0,
	-2, 63,
	-1, 118,
	44, 29,
	60, 29,
	-2, 32,
	-1, 119,
	44, 29,
	60, 29,
	-2, 32,
	-1, 120,
	57, 1,
	-2, 29,
	-1, 121,
	44, 29,
	60, 29,
	-2, 32,
	-1, 142,
	59, 32,
	-2, 29,
	-1, 171,
	57, 1,
	-2, 29,
	-1, 172,
	57, 1,
	-2, 29,
	-1, 173,
	57, 1,
	-2, 29,
	-1, 175,
	57, 1,
	-2, 29,
	-1, 185,
	57, 1,
	-2, 29,
	-1, 187,
	57, 1,
	-2, 29,
	-1, 188,
	57, 1,
	-2, 29,
	-1, 190,
	57, 1,
	-2, 29,
	-1, 199,
	57, 1,
	-2, 29,
	-1, 207,
	57, 1,
	-2, 29,
	-1, 211,
	57, 1,
	-2, 29,
	-1, 216,
	57, 1,
	-2, 29,
}

const yyNprod = 83
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 969

var yyAct = []int{

	1, 89, 134, 32, 34, 35, 36, 37, 39, 40,
	177, 93, 31, 49, 160, 93, 139, 121, 145, 138,
	201, 44, 45, 46, 47, 48, 66, 49, 189, 94,
	60, 216, 142, 93, 95, 42, 9, 43, 56, 58,
	179, 176, 61, 159, 60, 93, 149, 62, 64, 42,
	49, 43, 56, 58, 137, 69, 82, 83, 84, 68,
	46, 47, 48, 62, 218, 87, 91, 60, 185, 85,
	184, 124, 42, 126, 43, 56, 58, 215, 98, 99,
	212, 101, 102, 103, 104, 105, 106, 107, 108, 109,
	110, 111, 112, 113, 114, 115, 116, 62, 146, 135,
	209, 206, 204, 117, 203, 122, 202, 70, 125, 71,
	127, 128, 129, 130, 131, 132, 196, 193, 62, 86,
	192, 152, 191, 170, 133, 73, 74, 75, 76, 77,
	78, 62, 168, 79, 80, 158, 156, 144, 96, 211,
	164, 207, 199, 190, 188, 187, 175, 173, 171, 120,
	162, 72, 81, 140, 119, 62, 62, 214, 62, 208,
	155, 150, 151, 174, 153, 123, 178, 161, 90, 194,
	163, 143, 181, 182, 183, 88, 186, 165, 136, 62,
	100, 97, 167, 92, 169, 166, 195, 67, 197, 198,
	65, 200, 50, 51, 53, 55, 57, 59, 8, 7,
	205, 6, 73, 74, 75, 76, 77, 78, 210, 5,
	79, 80, 213, 2, 0, 0, 49, 217, 0, 0,
	41, 0, 52, 54, 44, 45, 46, 47, 48, 81,
	0, 0, 0, 60, 180, 0, 0, 0, 42, 0,
	43, 56, 58, 50, 51, 53, 55, 57, 59, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 49, 0, 0,
	0, 41, 0, 52, 54, 44, 45, 46, 47, 48,
	0, 0, 172, 0, 60, 0, 0, 0, 0, 42,
	0, 43, 56, 58, 50, 51, 53, 55, 57, 59,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 49, 0,
	0, 0, 41, 0, 52, 54, 44, 45, 46, 47,
	48, 0, 0, 0, 0, 60, 157, 0, 0, 0,
	42, 0, 43, 56, 58, 50, 51, 53, 55, 57,
	59, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 49,
	0, 0, 0, 41, 0, 52, 54, 44, 45, 46,
	47, 48, 0, 0, 0, 0, 60, 154, 0, 0,
	0, 42, 0, 43, 56, 58, 50, 51, 53, 55,
	57, 59, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	49, 0, 0, 0, 41, 0, 52, 54, 44, 45,
	46, 47, 48, 0, 0, 0, 0, 60, 0, 0,
	0, 0, 42, 148, 43, 56, 58, 50, 51, 53,
	55, 57, 59, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 49, 0, 0, 0, 41, 147, 52, 54, 44,
	45, 46, 47, 48, 0, 0, 0, 0, 60, 0,
	0, 0, 0, 42, 0, 43, 56, 58, 50, 51,
	53, 55, 57, 59, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 49, 0, 0, 0, 41, 0, 52, 54,
	44, 45, 46, 47, 48, 0, 0, 0, 0, 60,
	141, 0, 0, 0, 42, 0, 43, 56, 58, 50,
	51, 53, 55, 57, 59, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 49, 0, 0, 0, 41, 0, 52,
	54, 44, 45, 46, 47, 48, 0, 0, 0, 0,
	60, 0, 118, 0, 0, 42, 0, 43, 56, 58,
	50, 51, 53, 55, 57, 59, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 49, 0, 0, 0, 41, 0,
	52, 54, 44, 45, 46, 47, 48, 0, 0, 0,
	0, 60, 0, 0, 0, 0, 42, 0, 43, 56,
	58, 18, 17, 22, 0, 0, 27, 10, 13, 11,
	14, 38, 15, 0, 0, 0, 0, 0, 0, 0,
	30, 23, 24, 25, 12, 16, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 4, 0, 0, 0, 0,
	0, 0, 0, 18, 17, 22, 0, 19, 27, 10,
	13, 11, 14, 28, 15, 29, 0, 0, 20, 21,
	26, 0, 30, 23, 24, 25, 12, 16, 0, 0,
	0, 0, 0, 0, 0, 0, 3, 4, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 19,
	0, 0, 0, 0, 33, 28, 0, 29, 0, 0,
	20, 21, 26, 18, 17, 22, 0, 0, 27, 10,
	13, 11, 14, 0, 15, 0, 0, 0, 0, 0,
	0, 0, 30, 23, 24, 25, 12, 16, 0, 0,
	0, 0, 0, 0, 0, 0, 3, 4, 0, 50,
	51, 53, 55, 0, 59, 0, 0, 0, 0, 19,
	0, 0, 0, 0, 0, 28, 0, 29, 0, 0,
	20, 21, 26, 49, 0, 0, 0, 0, 0, 52,
	54, 44, 45, 46, 47, 48, 50, 51, 53, 55,
	60, 0, 0, 0, 0, 42, 0, 43, 56, 58,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	49, 53, 55, 0, 0, 0, 52, 54, 44, 45,
	46, 47, 48, 0, 0, 0, 0, 60, 0, 0,
	0, 0, 42, 49, 43, 56, 58, 0, 0, 52,
	54, 44, 45, 46, 47, 48, 18, 17, 22, 0,
	60, 27, 0, 0, 0, 42, 0, 43, 56, 58,
	0, 0, 0, 0, 0, 30, 23, 24, 25, 63,
	17, 22, 0, 0, 27, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 30, 23,
	24, 25, 19, 0, 0, 0, 0, 0, 28, 0,
	29, 0, 0, 20, 21, 26, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 19, 0, 0, 0, 0,
	0, 28, 0, 29, 0, 0, 20, 21, 26,
}
var yyPact = []int{

	749, -1000, 689, 749, 749, 749, 647, 749, 749, 583,
	905, 882, 186, 183, 1, 51, 95, -1000, 171, 882,
	882, 882, -1000, -1000, -1000, -1000, 905, 61, 162, 882,
	179, -15, -1000, 749, -1000, -1000, -1000, -1000, 125, -1000,
	-1000, 882, 882, 176, 882, 882, 882, 882, 882, 882,
	882, 882, 882, 882, 882, 882, 882, 882, 882, 882,
	905, -1000, 532, 94, 583, 93, -27, -1000, 882, 149,
	749, 882, 749, 882, 882, 882, 882, 882, 882, -1000,
	-1000, 905, -14, -14, -14, -62, 174, -4, -41, -1000,
	107, 481, -26, 167, 905, -1000, -40, 749, 430, 379,
	-1000, 9, 9, -14, -14, -14, 583, 832, 832, -28,
	-28, -28, -28, 583, 772, 583, 809, -13, 905, 905,
	749, 905, 328, 882, 79, 277, 78, 583, 583, 583,
	583, 583, 583, -16, -1000, -45, 159, 166, 162, -1000,
	882, -1000, 905, -1000, -1000, 882, 75, 882, -1000, -1000,
	-1000, -1000, 66, -1000, 92, 226, -1000, 91, 134, -1000,
	90, -18, -49, 158, -1000, 583, -19, 175, -1000, 583,
	-1000, 749, 749, 749, 12, 749, 89, 88, -31, -1000,
	87, 65, 63, 60, 165, 749, 59, 749, 749, 86,
	749, -1000, -1000, -1000, -39, 49, -1000, 47, 45, 749,
	44, 85, 129, -1000, -1000, 43, -1000, 749, 83, -1000,
	23, 749, 127, 20, -25, -1000, 749, 7, -1000,
}
var yyPgo = []int{

	0, 0, 213, 209, 201, 199, 198, 36, 42, 1,
	175, 12,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 3, 4, 4, 4, 5, 5,
	5, 6, 6, 6, 6, 9, 10, 10, 10, 11,
	11, 11, 8, 8, 8, 8, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 5, 4, 9, 5, 7, 7, 4,
	7, 15, 12, 11, 8, 3, 0, 1, 3, 0,
	1, 3, 0, 1, 3, 3, 1, 1, 2, 2,
	2, 1, 1, 1, 1, 5, 4, 3, 7, 8,
	8, 9, 3, 3, 3, 5, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 3, 3, 3,
	3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, 37, 38, -3, -4, -5, -6, -7,
	10, 12, 27, 11, 13, 15, 28, 5, 4, 50,
	61, 62, 6, 24, 25, 26, 63, 9, 56, 58,
	23, -11, -1, 55, -1, -1, -1, -1, 14, -1,
	-1, 45, 63, 65, 49, 50, 51, 52, 53, 41,
	17, 18, 47, 19, 48, 20, 66, 21, 67, 22,
	58, -8, -7, 4, -7, 4, -11, 4, 58, 4,
	56, 58, 56, 31, 32, 33, 34, 35, 36, 39,
	40, 58, -7, -7, -7, -8, 58, 4, -10, -9,
	6, -7, 4, 60, 44, -1, 13, 56, -7, -7,
	4, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -8, 60, 60,
	56, 44, -7, 16, -1, -7, -1, -7, -7, -7,
	-7, -7, -7, -8, 64, -11, 4, 58, 60, 57,
	46, 59, 58, 4, -8, 58, -1, 46, 64, 59,
	-8, -8, -1, -8, 59, -7, 57, 59, 57, 59,
	59, 8, -11, 4, -9, -7, -8, -7, 57, -7,
	57, 56, 56, 56, 29, 56, 59, 59, 8, 59,
	59, -1, -1, -1, 58, 56, -1, 56, 56, 59,
	56, 57, 57, 57, 4, -1, 57, -1, -1, 56,
	-1, 59, 57, 57, 57, -1, 57, 56, 30, 57,
	-1, 56, 57, -1, 30, 57, 56, -1, 57,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, 10,
	-2, 29, 0, 29, 0, 0, 0, 36, -2, 29,
	29, 29, 41, 42, 43, 44, -2, 0, 26, 29,
	0, 0, 2, -2, 4, 5, 6, 7, 0, 8,
	9, 29, 29, 0, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	-2, 11, 33, -2, 12, 0, 0, 30, 29, 0,
	-2, 29, -2, 29, 29, 29, 29, 29, 29, 75,
	76, -2, 38, 39, 40, 0, 29, 0, 0, 27,
	0, 0, 0, 0, -2, 3, 0, -2, 0, 0,
	53, 56, 57, 58, 59, 60, 61, -2, -2, 64,
	65, 66, 67, 77, 78, 79, 80, 0, -2, -2,
	-2, -2, 0, 29, 0, 0, 0, 69, 70, 71,
	72, 73, 74, 0, 47, 0, 30, 29, 0, 52,
	29, 54, -2, 31, 68, 29, 0, 29, 46, 82,
	34, 35, 0, 14, 0, 0, 19, 0, 0, 81,
	0, 0, 0, 30, 28, 25, 0, 0, 16, 45,
	13, -2, -2, -2, 0, -2, 0, 0, 0, 55,
	0, 0, 0, 0, 0, -2, 0, -2, -2, 0,
	-2, 17, 18, 20, 0, 0, 48, 0, 0, -2,
	0, 0, 24, 49, 50, 0, 15, -2, 0, 51,
	0, -2, 23, 0, 0, 22, -2, 0, 21,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 61, 3, 3, 3, 53, 67, 3,
	58, 59, 51, 49, 60, 50, 65, 52, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 46, 55,
	48, 44, 47, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 63, 3, 64, 62, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 56, 66, 57,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 54,
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
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
		}
	case 39:
		//line parser.go.y:259
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
		}
	case 40:
		//line parser.go.y:263
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
		}
	case 41:
		//line parser.go.y:267
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 42:
		//line parser.go.y:271
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 43:
		//line parser.go.y:275
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 44:
		//line parser.go.y:279
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
		}
	case 45:
		//line parser.go.y:283
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	case 46:
		//line parser.go.y:287
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 47:
		//line parser.go.y:291
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 48:
		//line parser.go.y:295
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 49:
		//line parser.go.y:299
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 50:
		//line parser.go.y:303
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 51:
		//line parser.go.y:307
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 52:
		//line parser.go.y:311
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 53:
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 54:
		//line parser.go.y:323
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 55:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 56:
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:335
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:351
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:355
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 63:
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 64:
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 65:
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 66:
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 67:
		//line parser.go.y:375
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 68:
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.LetExpr{Names: yyS[yypt-2].idents, Operator: "=", Exprs: yyS[yypt-0].exprs}
		}
	case 69:
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "+", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 70:
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "-", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 71:
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "*", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 72:
		//line parser.go.y:395
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "/", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 73:
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "&", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 74:
		//line parser.go.y:403
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "|", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 75:
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
		}
	case 76:
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
		}
	case 77:
		//line parser.go.y:415
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 78:
		//line parser.go.y:419
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 79:
		//line parser.go.y:423
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 80:
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 81:
		//line parser.go.y:431
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 82:
		//line parser.go.y:435
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	}
	goto yystack /* stack new state and value */
}
