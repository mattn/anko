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

//line parser.go.y:541

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
	-1, 28,
	63, 38,
	-2, 35,
	-1, 36,
	44, 35,
	-2, 1,
	-1, 65,
	49, 38,
	-2, 35,
	-1, 68,
	44, 32,
	-2, 43,
	-1, 75,
	60, 1,
	-2, 35,
	-1, 77,
	60, 1,
	-2, 35,
	-1, 86,
	49, 38,
	-2, 35,
	-1, 90,
	17, 43,
	18, 43,
	19, 43,
	20, 43,
	21, 43,
	22, 43,
	41, 43,
	42, 43,
	43, 43,
	44, 32,
	45, 43,
	47, 32,
	50, 43,
	51, 43,
	52, 43,
	53, 43,
	54, 43,
	55, 43,
	56, 43,
	61, 43,
	62, 43,
	66, 43,
	67, 43,
	-2, 47,
	-1, 92,
	17, 43,
	18, 43,
	19, 43,
	20, 43,
	21, 43,
	22, 43,
	41, 43,
	42, 43,
	43, 43,
	44, 32,
	45, 43,
	47, 32,
	50, 43,
	51, 43,
	52, 43,
	53, 43,
	54, 43,
	55, 43,
	56, 43,
	61, 43,
	62, 43,
	66, 43,
	67, 43,
	-2, 49,
	-1, 102,
	44, 35,
	-2, 38,
	-1, 106,
	60, 1,
	-2, 35,
	-1, 107,
	44, 33,
	47, 33,
	-2, 57,
	-1, 118,
	17, 0,
	18, 0,
	48, 0,
	-2, 74,
	-1, 119,
	17, 0,
	18, 0,
	48, 0,
	-2, 75,
	-1, 129,
	44, 35,
	-2, 38,
	-1, 130,
	44, 35,
	-2, 38,
	-1, 131,
	60, 1,
	-2, 35,
	-1, 132,
	44, 35,
	-2, 38,
	-1, 156,
	49, 38,
	-2, 35,
	-1, 162,
	44, 34,
	47, 34,
	-2, 56,
	-1, 176,
	44, 33,
	47, 33,
	-2, 48,
	-1, 177,
	44, 33,
	47, 33,
	-2, 50,
	-1, 189,
	60, 1,
	-2, 35,
	-1, 190,
	60, 1,
	-2, 35,
	-1, 191,
	60, 1,
	-2, 35,
	-1, 193,
	60, 1,
	-2, 35,
	-1, 203,
	60, 1,
	-2, 35,
	-1, 205,
	60, 1,
	-2, 35,
	-1, 206,
	60, 1,
	-2, 35,
	-1, 208,
	60, 1,
	-2, 35,
	-1, 217,
	60, 1,
	-2, 35,
	-1, 225,
	60, 1,
	-2, 35,
	-1, 229,
	60, 1,
	-2, 35,
	-1, 234,
	60, 1,
	-2, 35,
}

const yyNprod = 95
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1217

var yyAct = []int{

	1, 98, 52, 35, 37, 38, 39, 40, 42, 43,
	71, 148, 33, 47, 48, 49, 50, 51, 152, 52,
	74, 52, 44, 45, 236, 233, 230, 63, 61, 227,
	224, 153, 49, 50, 51, 222, 221, 104, 9, 44,
	45, 44, 45, 220, 63, 61, 63, 61, 66, 67,
	69, 214, 211, 210, 209, 202, 188, 186, 87, 88,
	89, 91, 93, 174, 76, 172, 203, 67, 234, 229,
	100, 225, 217, 208, 206, 75, 136, 94, 138, 205,
	105, 193, 191, 189, 108, 109, 110, 111, 112, 113,
	114, 115, 116, 117, 118, 119, 120, 121, 122, 123,
	124, 125, 126, 127, 67, 131, 149, 161, 77, 133,
	219, 195, 134, 207, 128, 137, 158, 139, 140, 141,
	142, 143, 144, 197, 194, 67, 106, 133, 175, 178,
	164, 96, 167, 160, 156, 145, 151, 132, 73, 103,
	133, 67, 159, 154, 78, 79, 80, 81, 82, 83,
	102, 157, 84, 85, 182, 232, 226, 192, 135, 196,
	130, 86, 180, 179, 99, 212, 181, 177, 67, 67,
	176, 67, 169, 150, 171, 95, 107, 101, 165, 166,
	72, 168, 70, 97, 34, 8, 7, 6, 5, 2,
	199, 200, 201, 183, 204, 67, 0, 0, 0, 185,
	0, 0, 187, 0, 213, 184, 215, 216, 0, 218,
	0, 78, 79, 80, 81, 82, 83, 0, 223, 84,
	85, 0, 0, 0, 0, 0, 228, 0, 86, 0,
	231, 18, 17, 24, 0, 235, 29, 10, 13, 11,
	14, 41, 15, 0, 0, 0, 0, 0, 0, 0,
	32, 25, 26, 27, 12, 16, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 4, 0, 0, 78, 79,
	80, 81, 82, 83, 0, 31, 84, 85, 0, 0,
	19, 23, 0, 0, 0, 0, 30, 0, 0, 28,
	0, 20, 21, 22, 18, 17, 24, 0, 0, 29,
	10, 13, 11, 14, 0, 15, 0, 0, 0, 0,
	0, 0, 0, 32, 25, 26, 27, 12, 16, 0,
	0, 0, 0, 0, 0, 0, 0, 3, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 31, 0,
	0, 0, 0, 19, 23, 0, 0, 0, 36, 30,
	0, 0, 28, 0, 20, 21, 22, 18, 17, 24,
	0, 0, 29, 10, 13, 11, 14, 0, 15, 0,
	0, 0, 0, 0, 0, 0, 32, 25, 26, 27,
	12, 16, 0, 0, 0, 0, 0, 0, 0, 0,
	3, 4, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 31, 0, 0, 0, 0, 19, 23, 0, 0,
	0, 0, 30, 0, 0, 28, 0, 20, 21, 22,
	55, 56, 58, 60, 62, 64, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 52, 53, 54, 0, 46, 0,
	0, 65, 198, 57, 59, 47, 48, 49, 50, 51,
	0, 0, 0, 0, 44, 45, 0, 0, 0, 63,
	61, 55, 56, 58, 60, 62, 64, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 52, 53, 54, 0, 46,
	0, 0, 65, 0, 57, 59, 47, 48, 49, 50,
	51, 0, 0, 190, 0, 44, 45, 0, 0, 0,
	63, 61, 55, 56, 58, 60, 62, 64, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 52, 53, 54, 0,
	46, 0, 0, 65, 173, 57, 59, 47, 48, 49,
	50, 51, 0, 0, 0, 0, 44, 45, 0, 0,
	0, 63, 61, 55, 56, 58, 60, 62, 64, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 52, 53, 54,
	0, 46, 0, 0, 65, 170, 57, 59, 47, 48,
	49, 50, 51, 0, 0, 0, 0, 44, 45, 0,
	0, 0, 63, 61, 55, 56, 58, 60, 62, 64,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 52, 53,
	54, 0, 46, 163, 0, 65, 0, 57, 59, 47,
	48, 49, 50, 51, 0, 0, 0, 0, 44, 45,
	0, 0, 0, 63, 61, 55, 56, 58, 60, 62,
	64, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 52,
	53, 54, 0, 46, 0, 0, 65, 0, 57, 59,
	47, 48, 49, 50, 51, 0, 0, 0, 0, 44,
	45, 162, 0, 0, 63, 61, 55, 56, 58, 60,
	62, 64, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	52, 53, 54, 0, 46, 0, 0, 65, 155, 57,
	59, 47, 48, 49, 50, 51, 0, 0, 0, 0,
	44, 45, 0, 0, 0, 63, 61, 55, 56, 58,
	60, 62, 64, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 52, 53, 54, 0, 46, 0, 129, 65, 0,
	57, 59, 47, 48, 49, 50, 51, 0, 0, 0,
	0, 44, 45, 0, 0, 0, 63, 61, 55, 56,
	58, 60, 62, 64, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 52, 53, 54, 0, 46, 0, 0, 65,
	0, 57, 59, 47, 48, 49, 50, 51, 0, 0,
	0, 0, 44, 45, 0, 0, 0, 63, 61, 55,
	56, 58, 60, 62, 64, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 52, 53, 54, 0, 46, 0, 0,
	65, 0, 57, 59, 47, 48, 49, 50, 51, 0,
	0, 0, 0, 147, 45, 0, 0, 0, 63, 61,
	55, 56, 58, 60, 62, 64, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 52, 53, 54, 0, 46, 0,
	0, 65, 0, 57, 59, 47, 48, 49, 50, 51,
	0, 0, 0, 0, 146, 45, 0, 0, 0, 63,
	61, 55, 56, 58, 60, 0, 64, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 52, 53, 54, 55, 56,
	58, 60, 65, 0, 57, 59, 47, 48, 49, 50,
	51, 0, 0, 0, 0, 44, 45, 0, 0, 0,
	63, 61, 52, 53, 54, 0, 0, 58, 60, 65,
	0, 57, 59, 47, 48, 49, 50, 51, 0, 0,
	0, 0, 44, 45, 0, 0, 0, 63, 61, 52,
	53, 54, 0, 0, 0, 0, 0, 0, 57, 59,
	47, 48, 49, 50, 51, 0, 18, 17, 24, 44,
	45, 29, 0, 0, 63, 61, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 32, 25, 26, 27, 0,
	0, 0, 68, 17, 24, 0, 0, 29, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	31, 32, 25, 26, 27, 19, 23, 0, 92, 17,
	24, 30, 0, 29, 28, 0, 20, 21, 22, 0,
	0, 0, 0, 0, 0, 0, 31, 32, 25, 26,
	27, 19, 23, 0, 90, 17, 24, 30, 0, 29,
	28, 0, 20, 21, 22, 0, 0, 0, 0, 0,
	0, 0, 31, 32, 25, 26, 27, 19, 23, 0,
	0, 0, 0, 30, 0, 0, 28, 0, 20, 21,
	22, 0, 0, 0, 0, 0, 0, 0, 31, 0,
	0, 0, 0, 19, 23, 0, 0, 0, 0, 30,
	0, 0, 28, 0, 20, 21, 22,
}
var yyPact = []int{

	353, -1000, 290, 353, 353, 353, 227, 353, 353, 811,
	1098, 1072, 178, 176, 90, 16, 49, -1000, 180, 1072,
	1072, 1072, 1150, 1124, -1000, -1000, -1000, -1000, 1098, 127,
	158, 1072, 173, 106, 92, -1000, 353, -1000, -1000, -1000,
	-1000, 67, -1000, -1000, 172, 1072, 1072, 1072, 1072, 1072,
	1072, 1072, 1072, 1072, 1072, 1072, 1072, 1072, 1072, 1072,
	1072, 1072, 1072, 1072, 1072, 1098, -1000, 760, 113, 811,
	46, 93, -1000, 1072, 142, 353, 1072, 353, 1072, 1072,
	1072, 1072, 1072, 1072, -1000, -1000, 1098, -20, -20, -20,
	237, 913, 237, 862, -52, 169, 88, -29, -1000, 97,
	709, 86, 1098, 1072, -1000, 85, 353, -1000, 658, 607,
	-22, -22, -20, -20, -20, 811, -39, -39, 1018, 1018,
	-39, -39, -39, -39, 811, 964, 811, 991, 81, 1098,
	1098, 353, 1098, 168, 556, 1072, 5, 505, 3, 811,
	811, 811, 811, 811, 811, 79, 166, 163, -1000, 80,
	155, 162, 158, -1000, 1072, -1000, 1098, -1000, -1000, 811,
	1072, -3, -1000, 1072, -1000, -1000, -1000, -4, -1000, -1000,
	24, 454, -1000, 23, 128, -1000, -1000, -1000, 22, 75,
	62, 151, -1000, 811, 74, 403, -1000, 811, -1000, 353,
	353, 353, 7, 353, 20, 15, 64, -1000, 14, -6,
	-7, -8, 161, 353, -9, 353, 353, 13, 353, -1000,
	-1000, -1000, 61, -17, -1000, -24, -25, 353, -30, 12,
	126, -1000, -1000, -31, -1000, 353, 10, -1000, -34, 353,
	125, -35, 9, -1000, 353, -36, -1000,
}
var yyPgo = []int{

	0, 0, 189, 188, 187, 186, 185, 38, 48, 184,
	12, 1, 183, 10,
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
	7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 5, 4, 9, 5, 7, 7, 4,
	7, 15, 12, 11, 8, 3, 0, 1, 3, 0,
	1, 3, 1, 3, 4, 0, 1, 3, 0, 1,
	3, 3, 1, 1, 2, 2, 2, 2, 4, 2,
	4, 1, 1, 1, 1, 5, 4, 3, 3, 7,
	8, 8, 9, 3, 3, 5, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 3,
	3, 3, 3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, 37, 38, -3, -4, -5, -6, -7,
	10, 12, 27, 11, 13, 15, 28, 5, 4, 53,
	64, 65, 66, 54, 6, 24, 25, 26, 62, 9,
	59, 48, 23, -10, -9, -1, 58, -1, -1, -1,
	-1, 14, -1, -1, 61, 62, 45, 52, 53, 54,
	55, 56, 41, 42, 43, 17, 18, 50, 19, 51,
	20, 67, 21, 66, 22, 48, -8, -7, 4, -7,
	4, -13, 4, 48, 4, 59, 48, 59, 31, 32,
	33, 34, 35, 36, 39, 40, 48, -7, -7, -7,
	4, -7, 4, -7, -8, 48, 4, -12, -11, 6,
	-7, 4, 44, 47, -1, 13, 59, 4, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -8, 47,
	47, 59, 44, 47, -7, 16, -1, -7, -1, -7,
	-7, -7, -7, -7, -7, -8, 61, 61, 63, -13,
	4, 48, 47, 60, 46, 49, 48, -8, -10, -7,
	48, -1, 63, 46, 49, -8, -8, -1, -8, 4,
	49, -7, 60, 49, 60, 49, 4, 4, 49, 8,
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
	35, 35, 35, 35, 51, 52, 53, 54, -2, 0,
	26, 35, 0, 0, 36, 2, -2, 4, 5, 6,
	7, 0, 8, 9, 0, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, -2, 11, 39, -2, 12,
	0, 0, 30, 35, 0, -2, 35, -2, 35, 35,
	35, 35, 35, 35, 87, 88, -2, 44, 45, 46,
	-2, 0, -2, 0, 0, 29, 0, 0, 27, 0,
	0, 0, -2, 35, 3, 0, -2, -2, 0, 0,
	66, 67, 68, 69, 70, 71, 72, 73, -2, -2,
	76, 77, 78, 79, 89, 90, 91, 92, 0, -2,
	-2, -2, -2, 0, 0, 35, 0, 0, 0, 81,
	82, 83, 84, 85, 86, 0, 0, 0, 58, 0,
	30, 29, 0, 63, 35, 64, -2, 80, 37, 0,
	35, 0, -2, 35, 94, 40, 41, 0, 14, 31,
	0, 0, 19, 0, 0, 93, -2, -2, 0, 0,
	0, 30, 28, 25, 0, 0, 16, 55, 13, -2,
	-2, -2, 0, -2, 0, 0, 0, 65, 0, 0,
	0, 0, 0, -2, 0, -2, -2, 0, -2, 17,
	18, 20, 0, 0, 59, 0, 0, -2, 0, 0,
	24, 60, 61, 0, 15, -2, 0, 62, 0, -2,
	23, 0, 0, 22, -2, 0, 21,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 64, 3, 3, 3, 56, 66, 3,
	48, 49, 54, 52, 47, 53, 61, 55, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 46, 58,
	51, 44, 50, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 62, 3, 63, 65, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 59, 67, 60,
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
			yyVAL.lhs = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}
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
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 48:
		//line parser.go.y:302
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 49:
		//line parser.go.y:307
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 50:
		//line parser.go.y:312
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 51:
		//line parser.go.y:317
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 52:
		//line parser.go.y:322
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:332
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:337
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:342
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:352
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 59:
		//line parser.go.y:357
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 60:
		//line parser.go.y:362
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 61:
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 62:
		//line parser.go.y:372
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:377
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
	case 64:
		//line parser.go.y:386
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 65:
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 66:
		//line parser.go.y:396
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 67:
		//line parser.go.y:401
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 68:
		//line parser.go.y:406
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 69:
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 70:
		//line parser.go.y:416
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 71:
		//line parser.go.y:421
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 72:
		//line parser.go.y:426
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 73:
		//line parser.go.y:431
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 74:
		//line parser.go.y:436
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 75:
		//line parser.go.y:441
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 76:
		//line parser.go.y:446
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 77:
		//line parser.go.y:451
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 78:
		//line parser.go.y:456
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 79:
		//line parser.go.y:461
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 80:
		//line parser.go.y:466
		{
			yyVAL.expr = &ast.LetsExpr{Lhss: yyS[yypt-2].lhss, Operator: "=", Rhss: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 81:
		//line parser.go.y:471
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "+=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 82:
		//line parser.go.y:476
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "-=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 83:
		//line parser.go.y:481
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "*=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 84:
		//line parser.go.y:486
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "/=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 85:
		//line parser.go.y:491
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "&=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 86:
		//line parser.go.y:496
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "|=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 87:
		//line parser.go.y:501
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 88:
		//line parser.go.y:506
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 89:
		//line parser.go.y:511
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 90:
		//line parser.go.y:516
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 91:
		//line parser.go.y:521
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 92:
		//line parser.go.y:526
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 93:
		//line parser.go.y:531
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 94:
		//line parser.go.y:536
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
