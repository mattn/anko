//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
	"unsafe"
)

//line parser.go.y:25
type yySymType struct {
	yys                    int
	stmt_var               ast.Stmt
	stmt_if                ast.Stmt
	stmt_for               ast.Stmt
	stmt_try_catch_finally ast.Stmt
	stmts                  []ast.Stmt
	stmt                   ast.Stmt
	teim                   ast.Expr
	tok                    Token
	expr                   ast.Expr
	exprs                  []ast.Expr
	lhs                    ast.Expr
	lhss                   []ast.Expr
	pair                   ast.Expr
	pairs                  []ast.Expr
	idents                 []string
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
	" ,",
	" (",
	" )",
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

//line parser.go.y:521

//line yacctab:1
var yyExca = []int{
	-1, 0,
	44, 35,
	-2, 1,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	44, 35,
	-2, 1,
	-1, 3,
	44, 35,
	-2, 1,
	-1, 4,
	44, 35,
	-2, 1,
	-1, 5,
	44, 35,
	-2, 1,
	-1, 6,
	44, 35,
	-2, 1,
	-1, 7,
	44, 35,
	-2, 1,
	-1, 8,
	44, 35,
	-2, 1,
	-1, 10,
	44, 35,
	-2, 38,
	-1, 18,
	44, 32,
	47, 32,
	-2, 43,
	-1, 26,
	63, 38,
	-2, 35,
	-1, 34,
	44, 35,
	-2, 1,
	-1, 63,
	49, 38,
	-2, 35,
	-1, 66,
	44, 32,
	-2, 43,
	-1, 73,
	60, 1,
	-2, 35,
	-1, 75,
	60, 1,
	-2, 35,
	-1, 84,
	49, 38,
	-2, 35,
	-1, 96,
	44, 35,
	-2, 38,
	-1, 100,
	60, 1,
	-2, 35,
	-1, 101,
	44, 33,
	47, 33,
	-2, 53,
	-1, 112,
	17, 0,
	18, 0,
	48, 0,
	-2, 70,
	-1, 113,
	17, 0,
	18, 0,
	48, 0,
	-2, 71,
	-1, 123,
	44, 35,
	-2, 38,
	-1, 124,
	44, 35,
	-2, 38,
	-1, 125,
	60, 1,
	-2, 35,
	-1, 126,
	44, 35,
	-2, 38,
	-1, 148,
	49, 38,
	-2, 35,
	-1, 154,
	44, 34,
	47, 34,
	-2, 52,
	-1, 179,
	60, 1,
	-2, 35,
	-1, 180,
	60, 1,
	-2, 35,
	-1, 181,
	60, 1,
	-2, 35,
	-1, 183,
	60, 1,
	-2, 35,
	-1, 193,
	60, 1,
	-2, 35,
	-1, 195,
	60, 1,
	-2, 35,
	-1, 196,
	60, 1,
	-2, 35,
	-1, 198,
	60, 1,
	-2, 35,
	-1, 207,
	60, 1,
	-2, 35,
	-1, 215,
	60, 1,
	-2, 35,
	-1, 219,
	60, 1,
	-2, 35,
	-1, 224,
	60, 1,
	-2, 35,
}

const yyNprod = 91
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 941

var yyAct = []int{

	1, 92, 50, 33, 35, 36, 37, 38, 40, 41,
	140, 31, 69, 45, 46, 47, 48, 49, 50, 226,
	72, 144, 42, 43, 223, 220, 217, 59, 61, 50,
	214, 47, 48, 49, 145, 98, 9, 224, 42, 43,
	212, 211, 210, 59, 61, 204, 64, 65, 67, 42,
	43, 201, 200, 199, 59, 61, 85, 86, 87, 192,
	178, 176, 166, 65, 74, 164, 94, 219, 215, 207,
	193, 198, 196, 88, 130, 73, 132, 99, 195, 183,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	65, 153, 141, 181, 179, 56, 58, 125, 128, 150,
	122, 131, 75, 133, 134, 135, 136, 137, 138, 209,
	127, 65, 185, 100, 197, 187, 159, 50, 51, 52,
	127, 139, 168, 65, 151, 90, 55, 57, 45, 46,
	47, 48, 49, 149, 184, 167, 172, 42, 43, 156,
	152, 148, 59, 61, 143, 126, 170, 71, 127, 97,
	65, 65, 146, 65, 96, 222, 163, 216, 182, 129,
	157, 158, 186, 160, 169, 93, 202, 171, 161, 89,
	189, 190, 191, 173, 194, 65, 142, 101, 95, 175,
	70, 68, 177, 91, 203, 174, 205, 206, 32, 208,
	8, 7, 53, 54, 56, 58, 60, 62, 213, 6,
	5, 2, 0, 0, 0, 0, 218, 0, 0, 0,
	221, 0, 0, 0, 0, 225, 50, 51, 52, 0,
	44, 0, 0, 63, 188, 55, 57, 45, 46, 47,
	48, 49, 0, 0, 0, 0, 42, 43, 0, 0,
	0, 59, 61, 53, 54, 56, 58, 60, 62, 76,
	77, 78, 79, 80, 81, 0, 0, 82, 83, 0,
	0, 0, 0, 0, 0, 124, 84, 50, 51, 52,
	0, 44, 0, 0, 63, 0, 55, 57, 45, 46,
	47, 48, 49, 0, 0, 180, 0, 42, 43, 0,
	0, 0, 59, 61, 53, 54, 56, 58, 60, 62,
	76, 77, 78, 79, 80, 81, 0, 0, 82, 83,
	0, 0, 0, 0, 0, 0, 0, 84, 50, 51,
	52, 0, 44, 0, 0, 63, 165, 55, 57, 45,
	46, 47, 48, 49, 0, 0, 0, 0, 42, 43,
	0, 0, 0, 59, 61, 53, 54, 56, 58, 60,
	62, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 50,
	51, 52, 0, 44, 0, 0, 63, 162, 55, 57,
	45, 46, 47, 48, 49, 0, 0, 0, 0, 42,
	43, 0, 0, 0, 59, 61, 53, 54, 56, 58,
	60, 62, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	50, 51, 52, 0, 44, 155, 0, 63, 0, 55,
	57, 45, 46, 47, 48, 49, 0, 0, 0, 0,
	42, 43, 0, 0, 0, 59, 61, 53, 54, 56,
	58, 60, 62, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 50, 51, 52, 0, 44, 0, 0, 63, 0,
	55, 57, 45, 46, 47, 48, 49, 0, 0, 0,
	0, 42, 43, 154, 0, 0, 59, 61, 53, 54,
	56, 58, 60, 62, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 50, 51, 52, 0, 44, 0, 0, 63,
	147, 55, 57, 45, 46, 47, 48, 49, 0, 0,
	0, 0, 42, 43, 0, 0, 0, 59, 61, 53,
	54, 56, 58, 60, 62, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 50, 51, 52, 0, 44, 0, 123,
	63, 0, 55, 57, 45, 46, 47, 48, 49, 0,
	0, 0, 0, 42, 43, 0, 0, 0, 59, 61,
	53, 54, 56, 58, 60, 62, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 50, 51, 52, 0, 44, 0,
	0, 63, 0, 55, 57, 45, 46, 47, 48, 49,
	0, 0, 0, 0, 42, 43, 18, 17, 22, 59,
	61, 27, 10, 13, 11, 14, 39, 15, 0, 0,
	0, 0, 0, 0, 0, 30, 23, 24, 25, 12,
	16, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	4, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	29, 18, 17, 22, 0, 19, 27, 10, 13, 11,
	14, 28, 15, 0, 26, 0, 20, 21, 0, 0,
	30, 23, 24, 25, 12, 16, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 4, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 29, 0, 0, 0, 0,
	19, 0, 0, 0, 0, 34, 28, 0, 0, 26,
	0, 20, 21, 18, 17, 22, 0, 0, 27, 10,
	13, 11, 14, 0, 15, 0, 0, 0, 0, 0,
	0, 0, 30, 23, 24, 25, 12, 16, 0, 0,
	0, 0, 0, 0, 0, 0, 3, 4, 0, 0,
	0, 53, 54, 56, 58, 0, 62, 29, 0, 0,
	0, 0, 19, 0, 0, 0, 0, 0, 28, 0,
	0, 26, 0, 20, 21, 50, 51, 52, 53, 54,
	56, 58, 63, 0, 55, 57, 45, 46, 47, 48,
	49, 0, 0, 0, 0, 42, 43, 0, 0, 0,
	59, 61, 50, 51, 52, 0, 0, 0, 0, 63,
	0, 55, 57, 45, 46, 47, 48, 49, 0, 18,
	17, 22, 42, 43, 27, 0, 0, 59, 61, 66,
	17, 22, 0, 0, 27, 0, 0, 0, 30, 23,
	24, 25, 0, 0, 0, 0, 0, 0, 30, 23,
	24, 25, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 29, 0, 0, 0, 0, 19, 0,
	0, 0, 0, 29, 28, 0, 0, 26, 19, 20,
	21, 0, 0, 0, 28, 0, 0, 26, 0, 20,
	21,
}
var yyPact = []int{

	759, -1000, 697, 759, 759, 759, 652, 759, 759, 593,
	875, 865, 187, 186, 109, 16, 53, -1000, 279, 865,
	865, 865, -1000, -1000, -1000, -1000, 875, 131, 169, 865,
	184, 120, 112, -1000, 759, -1000, -1000, -1000, -1000, 64,
	-1000, -1000, 183, 865, 865, 865, 865, 865, 865, 865,
	865, 865, 865, 865, 865, 865, 865, 865, 865, 865,
	865, 865, 865, 875, -1000, 542, 228, 593, 48, 111,
	-1000, 865, 153, 759, 865, 759, 865, 865, 865, 865,
	865, 865, -1000, -1000, 875, -12, -12, -12, -53, 182,
	106, -26, -1000, 116, 491, 103, 875, 865, -1000, 102,
	759, -1000, 440, 389, -23, -23, -12, -12, -12, 593,
	-39, -39, 86, 86, -39, -39, -39, -39, 593, 784,
	593, 811, 100, 875, 875, 759, 875, 174, 338, 865,
	5, 287, 2, 593, 593, 593, 593, 593, 593, 96,
	-1000, 83, 166, 173, 169, -1000, 865, -1000, 875, -1000,
	-1000, 593, 865, 1, -1000, 865, -1000, -1000, -1000, 0,
	-1000, -1000, 45, 236, -1000, 44, 139, -1000, 20, 95,
	73, 164, -1000, 593, 76, 185, -1000, 593, -1000, 759,
	759, 759, 11, 759, 19, 13, 75, -1000, 12, -7,
	-8, -9, 172, 759, -15, 759, 759, 10, 759, -1000,
	-1000, -1000, 70, -18, -1000, -19, -20, 759, -30, 9,
	137, -1000, -1000, -34, -1000, 759, 8, -1000, -35, 759,
	135, -36, -22, -1000, 759, -41, -1000,
}
var yyPgo = []int{

	0, 0, 211, 210, 209, 201, 200, 36, 46, 198,
	11, 1, 193, 12,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 3, 4, 4, 4, 5, 5,
	5, 6, 6, 6, 6, 11, 12, 12, 12, 13,
	13, 13, 9, 9, 9, 10, 10, 10, 8, 8,
	8, 8, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 5, 4, 9, 5, 7, 7, 4,
	7, 15, 12, 11, 8, 3, 0, 1, 3, 0,
	1, 3, 1, 3, 4, 0, 1, 3, 0, 1,
	3, 3, 1, 1, 2, 2, 2, 1, 1, 1,
	1, 5, 4, 3, 3, 7, 8, 8, 9, 3,
	3, 5, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 3, 3, 3, 3, 4,
	4,
}
var yyChk = []int{

	-1000, -1, -2, 37, 38, -3, -4, -5, -6, -7,
	10, 12, 27, 11, 13, 15, 28, 5, 4, 53,
	64, 65, 6, 24, 25, 26, 62, 9, 59, 48,
	23, -10, -9, -1, 58, -1, -1, -1, -1, 14,
	-1, -1, 61, 62, 45, 52, 53, 54, 55, 56,
	41, 42, 43, 17, 18, 50, 19, 51, 20, 66,
	21, 67, 22, 48, -8, -7, 4, -7, 4, -13,
	4, 48, 4, 59, 48, 59, 31, 32, 33, 34,
	35, 36, 39, 40, 48, -7, -7, -7, -8, 48,
	4, -12, -11, 6, -7, 4, 44, 47, -1, 13,
	59, 4, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -8, 47, 47, 59, 44, 47, -7, 16,
	-1, -7, -1, -7, -7, -7, -7, -7, -7, -8,
	63, -13, 4, 48, 47, 60, 46, 49, 48, -8,
	-10, -7, 48, -1, 63, 46, 49, -8, -8, -1,
	-8, 4, 49, -7, 60, 49, 60, 49, 49, 8,
	-13, 4, -11, -7, -8, -7, 60, -7, 60, 59,
	59, 59, 29, 59, 49, 49, 8, 49, 49, -1,
	-1, -1, 48, 59, -1, 59, 59, 49, 59, 60,
	60, 60, 4, -1, 60, -1, -1, 59, -1, 49,
	60, 60, 60, -1, 60, 59, 30, 60, -1, 59,
	60, -1, 30, 60, 59, -1, 60,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, 10,
	-2, 35, 0, 29, 0, 0, 0, 42, -2, 35,
	35, 35, 47, 48, 49, 50, -2, 0, 26, 35,
	0, 0, 36, 2, -2, 4, 5, 6, 7, 0,
	8, 9, 0, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, -2, 11, 39, -2, 12, 0, 0,
	30, 35, 0, -2, 35, -2, 35, 35, 35, 35,
	35, 35, 83, 84, -2, 44, 45, 46, 0, 29,
	0, 0, 27, 0, 0, 0, -2, 35, 3, 0,
	-2, -2, 0, 0, 62, 63, 64, 65, 66, 67,
	68, 69, -2, -2, 72, 73, 74, 75, 85, 86,
	87, 88, 0, -2, -2, -2, -2, 0, 0, 35,
	0, 0, 0, 77, 78, 79, 80, 81, 82, 0,
	54, 0, 30, 29, 0, 59, 35, 60, -2, 76,
	37, 0, 35, 0, -2, 35, 90, 40, 41, 0,
	14, 31, 0, 0, 19, 0, 0, 89, 0, 0,
	0, 30, 28, 25, 0, 0, 16, 51, 13, -2,
	-2, -2, 0, -2, 0, 0, 0, 61, 0, 0,
	0, 0, 0, -2, 0, -2, -2, 0, -2, 17,
	18, 20, 0, 0, 55, 0, 0, -2, 0, 0,
	24, 56, 57, 0, 15, -2, 0, 58, 0, -2,
	23, 0, 0, 22, -2, 0, 21,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 64, 3, 3, 3, 56, 67, 3,
	48, 49, 54, 52, 47, 53, 61, 55, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 46, 58,
	51, 44, 50, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 62, 3, 63, 65, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 59, 66, 60,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 57,
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
		//line parser.go.y:60
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:67
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-1].stmt.SetPos(l.pos)
			}
		}
	case 3:
		//line parser.go.y:75
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-2].stmt.SetPos(l.pos)
			}
		}
	case 4:
		//line parser.go.y:83
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.BreakStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:90
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.ContinueStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:97
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:104
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:111
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 9:
		//line parser.go.y:118
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 10:
		//line parser.go.y:126
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 11:
		//line parser.go.y:130
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 12:
		//line parser.go.y:134
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
		}
	case 13:
		//line parser.go.y:138
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:143
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 15:
		//line parser.go.y:148
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 16:
		//line parser.go.y:152
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 17:
		//line parser.go.y:160
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:165
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:169
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:173
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:178
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:182
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:186
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:190
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:195
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:200
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 27:
		//line parser.go.y:204
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 28:
		//line parser.go.y:208
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 29:
		//line parser.go.y:213
		{
			yyVAL.idents = []string{}
		}
	case 30:
		//line parser.go.y:217
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 31:
		//line parser.go.y:221
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 32:
		//line parser.go.y:226
		{
			yyVAL.lhs = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.lhs.SetPos(l.pos)
			}
		}
	case 33:
		//line parser.go.y:231
		{
			yyVAL.lhs = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.lhs.SetPos(l.pos)
			}
		}
	case 34:
		//line parser.go.y:236
		{
			yyVAL.lhs = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.lhs.SetPos(l.pos)
			}
		}
	case 35:
		//line parser.go.y:242
		{
			yyVAL.lhss = []ast.Expr{}
		}
	case 36:
		//line parser.go.y:246
		{
			yyVAL.lhss = []ast.Expr{yyS[yypt-0].lhs}
		}
	case 37:
		//line parser.go.y:250
		{
			yyVAL.lhss = append([]ast.Expr{yyS[yypt-2].lhs}, yyS[yypt-0].lhss...)
		}
	case 38:
		//line parser.go.y:255
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 39:
		//line parser.go.y:259
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 40:
		//line parser.go.y:263
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 41:
		//line parser.go.y:267
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 42:
		//line parser.go.y:272
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 43:
		//line parser.go.y:277
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 44:
		//line parser.go.y:282
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 45:
		//line parser.go.y:287
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 46:
		//line parser.go.y:292
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 47:
		//line parser.go.y:297
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 48:
		//line parser.go.y:302
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 49:
		//line parser.go.y:307
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 50:
		//line parser.go.y:312
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 51:
		//line parser.go.y:317
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 52:
		//line parser.go.y:322
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:332
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:337
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:342
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:352
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 59:
		//line parser.go.y:357
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
	case 60:
		//line parser.go.y:366
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 61:
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 62:
		//line parser.go.y:376
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:381
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 64:
		//line parser.go.y:386
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 65:
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 66:
		//line parser.go.y:396
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 67:
		//line parser.go.y:401
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 68:
		//line parser.go.y:406
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 69:
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 70:
		//line parser.go.y:416
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 71:
		//line parser.go.y:421
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 72:
		//line parser.go.y:426
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 73:
		//line parser.go.y:431
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 74:
		//line parser.go.y:436
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 75:
		//line parser.go.y:441
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 76:
		//line parser.go.y:446
		{
			yyVAL.expr = &ast.LetsExpr{Lhss: yyS[yypt-2].lhss, Operator: "=", Rhss: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 77:
		//line parser.go.y:451
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "+=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 78:
		//line parser.go.y:456
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "-=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 79:
		//line parser.go.y:461
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "*=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 80:
		//line parser.go.y:466
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "/=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 81:
		//line parser.go.y:471
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "&=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 82:
		//line parser.go.y:476
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "|=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 83:
		//line parser.go.y:481
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 84:
		//line parser.go.y:486
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 85:
		//line parser.go.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 86:
		//line parser.go.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 87:
		//line parser.go.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 88:
		//line parser.go.y:506
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 89:
		//line parser.go.y:511
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 90:
		//line parser.go.y:516
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
