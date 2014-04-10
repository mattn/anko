//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:24
type yySymType struct {
	yys          int
	stmt_if      ast.Stmt
	stmt_default ast.Stmt
	stmt_case    ast.Stmt
	stmt_cases   []ast.Stmt
	stmts        []ast.Stmt
	stmt         ast.Stmt
	expr         ast.Expr
	exprs        []ast.Expr
	expr_lhs     ast.Expr
	expr_lhss    []ast.Expr
	expr_pair    ast.Expr
	expr_pairs   []ast.Expr
	expr_idents  []string
	tok          Token
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
const SWITCH = 57386
const CASE = 57387
const DEFAULT = 57388
const UNARY = 57389

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
	"SWITCH",
	"CASE",
	"DEFAULT",
	" =",
	" ?",
	" :",
	" ,",
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

//line parser.go.y:578

//line yacctab:1
var yyExca = []int{
	-1, 0,
	47, 41,
	-2, 1,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	47, 41,
	-2, 1,
	-1, 6,
	47, 41,
	-2, 44,
	-1, 15,
	47, 38,
	50, 38,
	-2, 49,
	-1, 26,
	64, 44,
	-2, 41,
	-1, 34,
	47, 41,
	-2, 1,
	-1, 56,
	69, 44,
	-2, 41,
	-1, 59,
	47, 38,
	-2, 49,
	-1, 65,
	47, 38,
	50, 38,
	-2, 49,
	-1, 66,
	50, 32,
	61, 1,
	-2, 41,
	-1, 68,
	61, 1,
	-2, 41,
	-1, 78,
	69, 44,
	-2, 41,
	-1, 83,
	17, 49,
	18, 49,
	19, 49,
	20, 49,
	21, 49,
	22, 49,
	41, 49,
	42, 49,
	43, 49,
	47, 38,
	48, 49,
	50, 38,
	51, 49,
	52, 49,
	53, 49,
	54, 49,
	55, 49,
	56, 49,
	57, 49,
	62, 49,
	67, 49,
	70, 49,
	-2, 53,
	-1, 85,
	17, 49,
	18, 49,
	19, 49,
	20, 49,
	21, 49,
	22, 49,
	41, 49,
	42, 49,
	43, 49,
	47, 38,
	48, 49,
	50, 38,
	51, 49,
	52, 49,
	53, 49,
	54, 49,
	55, 49,
	56, 49,
	57, 49,
	62, 49,
	67, 49,
	70, 49,
	-2, 55,
	-1, 95,
	47, 41,
	-2, 44,
	-1, 99,
	47, 39,
	50, 39,
	-2, 62,
	-1, 110,
	17, 0,
	18, 0,
	-2, 79,
	-1, 111,
	17, 0,
	18, 0,
	-2, 80,
	-1, 121,
	47, 41,
	-2, 44,
	-1, 122,
	47, 41,
	-2, 44,
	-1, 123,
	61, 1,
	-2, 41,
	-1, 124,
	47, 41,
	-2, 44,
	-1, 127,
	61, 1,
	-2, 41,
	-1, 131,
	61, 1,
	-2, 41,
	-1, 153,
	69, 44,
	-2, 41,
	-1, 155,
	61, 1,
	-2, 41,
	-1, 158,
	47, 40,
	50, 40,
	-2, 100,
	-1, 180,
	47, 39,
	50, 39,
	-2, 54,
	-1, 181,
	47, 39,
	50, 39,
	-2, 56,
	-1, 192,
	61, 1,
	-2, 41,
	-1, 194,
	61, 1,
	-2, 41,
	-1, 202,
	47, 41,
	-2, 1,
	-1, 203,
	61, 1,
	-2, 41,
	-1, 213,
	61, 1,
	-2, 41,
	-1, 214,
	47, 41,
	-2, 1,
	-1, 217,
	61, 1,
	-2, 41,
	-1, 218,
	61, 1,
	-2, 41,
	-1, 222,
	61, 1,
	-2, 41,
	-1, 223,
	61, 1,
	-2, 41,
	-1, 229,
	61, 1,
	-2, 41,
	-1, 241,
	61, 1,
	-2, 41,
	-1, 242,
	61, 1,
	-2, 41,
}

const yyNprod = 102
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1493

var yyAct = []int{

	1, 174, 175, 33, 57, 91, 62, 30, 219, 125,
	125, 43, 70, 71, 72, 73, 74, 75, 207, 204,
	76, 77, 178, 38, 39, 40, 41, 42, 205, 182,
	160, 89, 35, 36, 88, 98, 153, 54, 56, 147,
	52, 148, 242, 149, 79, 246, 245, 176, 177, 78,
	239, 15, 14, 21, 150, 237, 25, 6, 9, 7,
	31, 120, 11, 198, 236, 234, 43, 129, 233, 133,
	29, 22, 23, 24, 8, 12, 232, 226, 221, 220,
	40, 41, 42, 141, 4, 5, 212, 35, 36, 208,
	195, 13, 54, 56, 145, 52, 43, 193, 87, 191,
	154, 16, 20, 172, 169, 156, 34, 27, 241, 229,
	26, 223, 17, 18, 19, 28, 126, 35, 36, 218,
	217, 128, 54, 56, 163, 52, 161, 162, 167, 164,
	203, 123, 170, 68, 97, 202, 70, 71, 72, 73,
	74, 75, 213, 151, 76, 77, 46, 47, 49, 51,
	53, 55, 176, 177, 184, 186, 189, 124, 188, 95,
	125, 240, 238, 127, 197, 64, 206, 183, 79, 92,
	43, 44, 45, 78, 185, 199, 200, 37, 181, 180,
	48, 50, 38, 39, 40, 41, 42, 165, 132, 131,
	146, 35, 36, 209, 99, 210, 54, 56, 94, 52,
	63, 61, 90, 215, 216, 32, 173, 10, 2, 0,
	0, 0, 0, 0, 224, 225, 0, 0, 227, 228,
	3, 0, 0, 230, 231, 0, 0, 58, 60, 0,
	235, 0, 67, 0, 69, 0, 0, 80, 81, 82,
	84, 86, 243, 244, 0, 0, 0, 58, 0, 93,
	0, 0, 96, 0, 0, 0, 0, 100, 101, 102,
	103, 104, 105, 106, 107, 108, 109, 110, 111, 112,
	113, 114, 115, 116, 117, 118, 119, 58, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 135, 136, 137, 138, 139, 140, 0, 0, 58,
	142, 15, 14, 21, 0, 0, 25, 6, 9, 7,
	31, 0, 11, 0, 0, 0, 58, 0, 157, 0,
	29, 22, 23, 24, 8, 12, 0, 0, 0, 0,
	0, 0, 0, 0, 4, 5, 0, 0, 0, 0,
	0, 13, 58, 58, 0, 58, 0, 166, 0, 168,
	0, 16, 20, 171, 0, 0, 0, 27, 0, 0,
	26, 0, 17, 18, 19, 28, 0, 0, 0, 0,
	0, 0, 187, 0, 58, 0, 0, 0, 0, 0,
	190, 15, 14, 130, 0, 0, 25, 6, 9, 7,
	31, 0, 11, 0, 0, 0, 0, 201, 0, 0,
	29, 22, 23, 24, 8, 12, 0, 0, 0, 0,
	0, 0, 0, 0, 4, 5, 0, 211, 0, 0,
	0, 13, 46, 47, 49, 51, 53, 55, 0, 0,
	0, 16, 20, 0, 0, 0, 0, 27, 0, 0,
	26, 0, 17, 18, 19, 28, 43, 44, 45, 0,
	0, 0, 0, 37, 0, 0, 48, 50, 38, 39,
	40, 41, 42, 0, 0, 222, 0, 35, 36, 0,
	0, 0, 54, 56, 0, 52, 46, 47, 49, 51,
	53, 55, 0, 0, 0, 0, 70, 71, 72, 73,
	74, 75, 0, 0, 76, 77, 0, 0, 0, 0,
	43, 44, 45, 0, 0, 122, 0, 37, 214, 0,
	48, 50, 38, 39, 40, 41, 42, 0, 79, 0,
	0, 35, 36, 78, 0, 0, 54, 56, 0, 52,
	46, 47, 49, 51, 53, 55, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 44, 45, 0, 0, 0,
	0, 37, 0, 0, 48, 50, 38, 39, 40, 41,
	42, 0, 196, 0, 0, 35, 36, 0, 0, 0,
	54, 56, 0, 52, 46, 47, 49, 51, 53, 55,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 43, 44,
	45, 0, 0, 0, 0, 37, 0, 0, 48, 50,
	38, 39, 40, 41, 42, 0, 0, 194, 0, 35,
	36, 0, 0, 0, 54, 56, 0, 52, 46, 47,
	49, 51, 53, 55, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 43, 44, 45, 0, 0, 0, 0, 37,
	0, 0, 48, 50, 38, 39, 40, 41, 42, 0,
	0, 192, 0, 35, 36, 0, 0, 0, 54, 56,
	0, 52, 46, 47, 49, 51, 53, 55, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 43, 44, 45, 0,
	0, 0, 0, 37, 0, 0, 48, 50, 38, 39,
	40, 41, 42, 0, 0, 0, 0, 35, 36, 179,
	0, 0, 54, 56, 0, 52, 46, 47, 49, 51,
	53, 55, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	43, 44, 45, 0, 0, 0, 0, 37, 159, 0,
	48, 50, 38, 39, 40, 41, 42, 0, 0, 0,
	0, 35, 36, 0, 0, 0, 54, 56, 0, 52,
	46, 47, 49, 51, 53, 55, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 44, 45, 0, 0, 0,
	0, 37, 0, 0, 48, 50, 38, 39, 40, 41,
	42, 0, 0, 0, 0, 35, 36, 158, 0, 0,
	54, 56, 0, 52, 46, 47, 49, 51, 53, 55,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 43, 44,
	45, 0, 0, 0, 0, 37, 0, 0, 48, 50,
	38, 39, 40, 41, 42, 0, 0, 155, 0, 35,
	36, 0, 0, 0, 54, 56, 0, 52, 46, 47,
	49, 51, 53, 55, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 43, 44, 45, 0, 0, 0, 0, 37,
	0, 0, 48, 50, 38, 39, 40, 41, 42, 0,
	0, 0, 0, 35, 36, 0, 0, 0, 54, 56,
	152, 52, 46, 47, 49, 51, 53, 55, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 43, 44, 45, 0,
	0, 0, 0, 37, 0, 0, 48, 50, 38, 39,
	40, 41, 42, 0, 0, 134, 0, 35, 36, 0,
	0, 0, 54, 56, 0, 52, 46, 47, 49, 51,
	53, 55, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	43, 44, 45, 0, 0, 0, 0, 37, 0, 121,
	48, 50, 38, 39, 40, 41, 42, 0, 0, 0,
	0, 35, 36, 0, 0, 0, 54, 56, 0, 52,
	46, 47, 49, 51, 53, 55, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 44, 45, 0, 0, 0,
	0, 37, 0, 0, 48, 50, 38, 39, 40, 41,
	42, 0, 0, 0, 0, 35, 36, 0, 0, 0,
	54, 56, 0, 52, 46, 47, 49, 51, 53, 55,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 43, 44,
	45, 0, 0, 0, 0, 37, 0, 0, 48, 50,
	38, 39, 40, 41, 42, 0, 0, 0, 0, 144,
	36, 0, 0, 0, 54, 56, 0, 52, 46, 47,
	49, 51, 53, 55, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 43, 44, 45, 0, 0, 0, 0, 37,
	0, 0, 48, 50, 38, 39, 40, 41, 42, 0,
	0, 0, 0, 143, 36, 0, 0, 0, 54, 56,
	0, 52, 46, 47, 49, 51, 0, 55, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 43, 44, 45, 0,
	0, 0, 46, 47, 49, 51, 48, 50, 38, 39,
	40, 41, 42, 0, 0, 0, 0, 35, 36, 0,
	0, 0, 54, 56, 0, 52, 43, 44, 45, 0,
	0, 0, 0, 0, 49, 51, 48, 50, 38, 39,
	40, 41, 42, 0, 0, 0, 0, 35, 36, 0,
	0, 0, 54, 56, 0, 52, 43, 44, 45, 0,
	0, 0, 0, 0, 0, 0, 48, 50, 38, 39,
	40, 41, 42, 0, 15, 14, 21, 35, 36, 25,
	0, 0, 54, 56, 0, 52, 0, 0, 0, 0,
	0, 0, 0, 29, 22, 23, 24, 59, 14, 21,
	0, 0, 25, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 29, 22, 23, 24,
	0, 85, 14, 21, 16, 20, 25, 0, 0, 0,
	27, 0, 0, 26, 0, 17, 18, 19, 28, 0,
	29, 22, 23, 24, 83, 14, 21, 16, 20, 25,
	0, 0, 0, 27, 0, 0, 26, 0, 17, 18,
	19, 28, 0, 29, 22, 23, 24, 0, 65, 14,
	21, 16, 20, 25, 0, 0, 0, 27, 0, 0,
	26, 0, 17, 18, 19, 28, 0, 29, 22, 23,
	24, 0, 0, 0, 16, 20, 0, 0, 0, 0,
	27, 0, 0, 26, 0, 17, 18, 19, 28, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 16, 20,
	0, 0, 0, 0, 66, 0, 0, 26, 0, 17,
	18, 19, 28,
}
var yyPact = []int{

	297, -1000, 47, 1053, -1000, -1000, 1353, 1330, 197, 196,
	151, 1424, 73, 1330, -1000, -19, 1330, 1330, 1330, 1400,
	1377, -1000, -1000, -1000, -1000, 30, 1353, 163, 1330, 194,
	112, 1330, 84, -1000, 297, 190, 1330, 1330, 1330, 1330,
	1330, 1330, 1330, 1330, 1330, 1330, 1330, 1330, 1330, 1330,
	1330, 1330, 1330, 1330, 1330, 1330, 1353, -1000, 999, 455,
	1053, 71, 110, -1000, 103, 105, 377, 129, 297, 945,
	1330, 1330, 1330, 1330, 1330, 1330, -1000, -1000, 1353, 1330,
	55, 55, 55, -19, 1161, -19, 1107, 186, -29, -23,
	-7, -1000, 94, 891, -32, 1353, 837, 1330, -1000, -1000,
	783, 729, 25, 25, 55, 55, 55, 1053, -30, -30,
	1275, 1275, -30, -30, -30, -30, 1053, 1215, 1053, 1245,
	-39, 1353, 1353, 297, 1353, 183, 1330, 297, 1330, 43,
	94, 297, 1330, 42, 107, 1053, 1053, 1053, 1053, 1053,
	1053, -47, 675, 175, 174, -40, 159, 170, -1000, 163,
	-1000, 1330, -1000, 1353, -1000, 297, -1000, 1053, -1000, 1330,
	-1000, -1000, -1000, 38, -1000, -1000, 621, 36, 567, -1000,
	29, 513, 135, 2, -1000, -1000, 1330, 86, -1000, -1000,
	-1000, -1000, 70, -50, -41, 158, -1000, 1053, -51, 28,
	1053, -1000, 297, -1000, 297, -1000, 1330, 82, -1000, -1000,
	-1000, 459, 297, 297, 60, 59, -61, -1000, -1000, 18,
	17, 405, 51, 297, 297, -1000, 16, 297, 297, 49,
	-1000, -1000, 297, 297, 15, -1000, -1000, 7, 4, 297,
	3, -6, 132, -1000, -1000, -11, -1000, 131, 48, -1000,
	-18, 297, 297, -15, -16, -1000, -1000,
}
var yyPgo = []int{

	0, 0, 208, 207, 2, 1, 206, 220, 4, 205,
	7, 5, 202, 6,
}
var yyR1 = []int{

	0, 1, 1, 1, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 3, 3, 3, 6, 6, 6, 6, 6, 5,
	4, 11, 12, 12, 12, 13, 13, 13, 9, 9,
	9, 10, 10, 10, 8, 8, 8, 8, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 1, 1, 1, 2, 2, 5,
	4, 1, 7, 4, 5, 9, 13, 12, 9, 8,
	5, 7, 5, 5, 0, 1, 2, 1, 2, 4,
	3, 3, 0, 1, 3, 0, 1, 3, 1, 3,
	4, 0, 1, 3, 0, 1, 3, 3, 1, 1,
	2, 2, 2, 2, 4, 2, 4, 1, 1, 1,
	1, 5, 3, 7, 8, 8, 9, 3, 3, 3,
	5, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 3, 3, 3, 3, 4, 4,
	4, 4,
}
var yyChk = []int{

	-1000, -1, -2, -7, 37, 38, 10, 12, 27, 11,
	-3, 15, 28, 44, 5, 4, 54, 65, 66, 67,
	55, 6, 24, 25, 26, 9, 63, 60, 68, 23,
	-10, 13, -9, -1, 59, 62, 63, 48, 53, 54,
	55, 56, 57, 41, 42, 43, 17, 18, 51, 19,
	52, 20, 70, 21, 67, 22, 68, -8, -7, 4,
	-7, 4, -13, 4, 14, 4, 60, -7, 60, -7,
	31, 32, 33, 34, 35, 36, 39, 40, 68, 63,
	-7, -7, -7, 4, -7, 4, -7, 68, 4, -8,
	-12, -11, 6, -7, 4, 47, -7, 50, -1, 4,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-8, 50, 50, 60, 47, 50, 13, 60, 16, -1,
	6, 60, 59, -1, 60, -7, -7, -7, -7, -7,
	-7, -8, -7, 62, 62, -13, 4, 68, 64, 50,
	61, 49, 69, 68, -8, 60, -10, -7, 64, 49,
	69, -8, -8, -1, -8, 4, -7, -1, -7, 61,
	-1, -7, 61, -6, -5, -4, 45, 46, 69, 64,
	4, 4, 69, 8, -13, 4, -11, -7, -8, -1,
	-7, 61, 60, 61, 60, 61, 59, 29, 61, -5,
	-4, -7, 49, 60, 69, 69, 8, 69, 61, -1,
	-1, -7, 4, 60, 49, -1, -1, 60, 60, 69,
	61, 61, 60, 60, -1, -1, 61, -1, -1, 60,
	-1, -1, 61, 61, 61, -1, 61, 61, 30, 61,
	30, 60, 60, -1, -1, 61, 61,
}
var yyDef = []int{

	-2, -2, -2, 4, 5, 6, -2, 41, 0, 35,
	11, 41, 0, 41, 48, -2, 41, 41, 41, 41,
	41, 57, 58, 59, 60, 0, -2, 32, 41, 0,
	0, 41, 42, 2, -2, 0, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, -2, 7, 45, -2,
	8, 0, 0, 36, 0, -2, -2, 0, -2, 0,
	41, 41, 41, 41, 41, 41, 92, 93, -2, 41,
	50, 51, 52, -2, 0, -2, 0, 35, 0, 0,
	0, 33, 0, 0, 0, -2, 0, 41, 3, -2,
	0, 0, 71, 72, 73, 74, 75, 76, 77, 78,
	-2, -2, 81, 82, 83, 84, 94, 95, 96, 97,
	0, -2, -2, -2, -2, 0, 41, -2, 41, 0,
	57, -2, 41, 0, 24, 86, 87, 88, 89, 90,
	91, 0, 0, 0, 0, 0, 36, 35, 67, 0,
	68, 41, 69, -2, 85, -2, 43, 0, -2, 41,
	101, 46, 47, 0, 10, 37, 0, 0, 0, 13,
	0, 0, 0, 0, 25, 27, 41, 0, 98, 99,
	-2, -2, 0, 0, 0, 36, 34, 31, 0, 0,
	61, 9, -2, 22, -2, 14, 41, 0, 20, 26,
	28, 0, -2, -2, 0, 0, 0, 70, 23, 0,
	0, 0, 0, -2, -2, 30, 0, -2, -2, 0,
	21, 12, -2, -2, 0, 29, 63, 0, 0, -2,
	0, 0, 19, 64, 65, 0, 15, 18, 0, 66,
	0, -2, -2, 0, 0, 17, 16,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 65, 3, 3, 3, 57, 67, 3,
	68, 69, 55, 53, 50, 54, 62, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 59,
	52, 47, 51, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 63, 3, 64, 66, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 60, 70, 61,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 58,
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
				for _, s := range yyS[yypt-0].stmts {
					if yyS[yypt-1].stmt.GetPos().Line == s.GetPos().Line {
						l.pos = yyS[yypt-1].stmt.GetPos()
						yylex.Error("syntax error")
					}
				}
			}
		}
	case 3:
		//line parser.go.y:78
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-2].stmt.SetPos(l.pos)
			}
		}
	case 4:
		//line parser.go.y:87
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			yyVAL.stmt.SetPos(yyS[yypt-0].expr.GetPos())
		}
	case 5:
		//line parser.go.y:92
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 6:
		//line parser.go.y:97
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 7:
		//line parser.go.y:102
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
			yyVAL.stmt.SetPos(yyS[yypt-1].tok.GetPos())
		}
	case 8:
		//line parser.go.y:107
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			yyVAL.stmt.SetPos(yyS[yypt-1].tok.GetPos())
		}
	case 9:
		//line parser.go.y:112
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPos(yyS[yypt-4].tok.GetPos())
		}
	case 10:
		//line parser.go.y:117
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyS[yypt-2].expr_idents, Exprs: yyS[yypt-0].exprs}
			yyVAL.stmt.SetPos(yyS[yypt-3].tok.GetPos())
		}
	case 11:
		//line parser.go.y:122
		{
			yyVAL.stmt = yyS[yypt-0].stmt_if
			yyVAL.stmt.SetPos(yyS[yypt-0].stmt_if.GetPos())
		}
	case 12:
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPos(yyS[yypt-6].tok.GetPos())
		}
	case 13:
		//line parser.go.y:132
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPos(yyS[yypt-3].tok.GetPos())
		}
	case 14:
		//line parser.go.y:137
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPos(yyS[yypt-4].tok.GetPos())
		}
	case 15:
		//line parser.go.y:142
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyS[yypt-7].expr, Expr2: yyS[yypt-5].expr, Expr3: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPos(yyS[yypt-8].tok.GetPos())
		}
	case 16:
		//line parser.go.y:147
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-10].stmts, Var: yyS[yypt-7].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPos(yyS[yypt-12].tok.GetPos())
		}
	case 17:
		//line parser.go.y:152
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPos(yyS[yypt-11].tok.GetPos())
		}
	case 18:
		//line parser.go.y:157
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-6].stmts, Var: yyS[yypt-3].tok.lit, Catch: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPos(yyS[yypt-8].tok.GetPos())
		}
	case 19:
		//line parser.go.y:162
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPos(yyS[yypt-7].tok.GetPos())
		}
	case 20:
		//line parser.go.y:168
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyS[yypt-3].expr, Cases: yyS[yypt-1].stmt_cases}
			yyVAL.stmt.SetPos(yyS[yypt-4].tok.GetPos())
		}
	case 21:
		//line parser.go.y:174
		{
			yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts})
		}
	case 22:
		//line parser.go.y:178
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyS[yypt-1].stmts...)
			}
		}
	case 23:
		//line parser.go.y:186
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts, Else: nil}
		}
	case 24:
		//line parser.go.y:191
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 25:
		//line parser.go.y:195
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_case}
		}
	case 26:
		//line parser.go.y:199
		{
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_case)
		}
	case 27:
		//line parser.go.y:203
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_default}
		}
	case 28:
		//line parser.go.y:207
		{
			for _, stmt := range yyS[yypt-1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_default)
		}
	case 29:
		//line parser.go.y:217
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyS[yypt-2].expr, Stmts: yyS[yypt-0].stmts}
		}
	case 30:
		//line parser.go.y:222
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyS[yypt-0].stmts}
		}
	case 31:
		//line parser.go.y:227
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 32:
		//line parser.go.y:232
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 33:
		//line parser.go.y:236
		{
			yyVAL.expr_pairs = []ast.Expr{yyS[yypt-0].expr_pair}
		}
	case 34:
		//line parser.go.y:240
		{
			yyVAL.expr_pairs = append(yyS[yypt-2].expr_pairs, yyS[yypt-0].expr_pair)
		}
	case 35:
		//line parser.go.y:245
		{
			yyVAL.expr_idents = []string{}
		}
	case 36:
		//line parser.go.y:249
		{
			yyVAL.expr_idents = []string{yyS[yypt-0].tok.lit}
		}
	case 37:
		//line parser.go.y:253
		{
			yyVAL.expr_idents = append(yyS[yypt-2].expr_idents, yyS[yypt-0].tok.lit)
		}
	case 38:
		//line parser.go.y:258
		{
			yyVAL.expr_lhs = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPos(l.pos)
			}
		}
	case 39:
		//line parser.go.y:263
		{
			yyVAL.expr_lhs = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPos(l.pos)
			}
		}
	case 40:
		//line parser.go.y:268
		{
			yyVAL.expr_lhs = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPos(l.pos)
			}
		}
	case 41:
		//line parser.go.y:274
		{
			yyVAL.expr_lhss = []ast.Expr{}
		}
	case 42:
		//line parser.go.y:278
		{
			yyVAL.expr_lhss = []ast.Expr{yyS[yypt-0].expr_lhs}
		}
	case 43:
		//line parser.go.y:282
		{
			yyVAL.expr_lhss = append([]ast.Expr{yyS[yypt-2].expr_lhs}, yyS[yypt-0].expr_lhss...)
		}
	case 44:
		//line parser.go.y:287
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 45:
		//line parser.go.y:291
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 46:
		//line parser.go.y:295
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 47:
		//line parser.go.y:299
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 48:
		//line parser.go.y:304
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
			yyVAL.expr.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 49:
		//line parser.go.y:309
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			yyVAL.expr.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 50:
		//line parser.go.y:314
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-0].expr.GetPos())
		}
	case 51:
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-0].expr.GetPos())
		}
	case 52:
		//line parser.go.y:324
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-0].expr.GetPos())
		}
	case 53:
		//line parser.go.y:329
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}}
			yyVAL.expr.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 54:
		//line parser.go.y:334
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 55:
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}}
			yyVAL.expr.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 56:
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 57:
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
			yyVAL.expr.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 58:
		//line parser.go.y:354
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.lit}
			yyVAL.expr.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 59:
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.lit}
			yyVAL.expr.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 60:
		//line parser.go.y:364
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.lit}
			yyVAL.expr.SetPos(yyS[yypt-0].tok.GetPos())
		}
	case 61:
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-4].expr.GetPos())
		}
	case 62:
		//line parser.go.y:374
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 63:
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			yyVAL.expr.SetPos(yyS[yypt-6].tok.GetPos())
		}
	case 64:
		//line parser.go.y:384
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			yyVAL.expr.SetPos(yyS[yypt-7].tok.GetPos())
		}
	case 65:
		//line parser.go.y:389
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			yyVAL.expr.SetPos(yyS[yypt-7].tok.GetPos())
		}
	case 66:
		//line parser.go.y:394
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			yyVAL.expr.SetPos(yyS[yypt-8].tok.GetPos())
		}
	case 67:
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 68:
		//line parser.go.y:404
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 69:
		//line parser.go.y:413
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 70:
		//line parser.go.y:418
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPos(yyS[yypt-4].tok.GetPos())
		}
	case 71:
		//line parser.go.y:423
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 72:
		//line parser.go.y:428
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 73:
		//line parser.go.y:433
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 74:
		//line parser.go.y:438
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 75:
		//line parser.go.y:443
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 76:
		//line parser.go.y:448
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 77:
		//line parser.go.y:453
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 78:
		//line parser.go.y:458
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 79:
		//line parser.go.y:463
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 80:
		//line parser.go.y:468
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 81:
		//line parser.go.y:473
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 82:
		//line parser.go.y:478
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 83:
		//line parser.go.y:483
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 84:
		//line parser.go.y:488
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 85:
		//line parser.go.y:493
		{
			yyVAL.expr = &ast.LetsExpr{Lhss: yyS[yypt-2].expr_lhss, Operator: "=", Rhss: yyS[yypt-0].exprs}
			yyVAL.expr.SetPos(yyS[yypt-2].expr_lhss[0].GetPos())
		}
	case 86:
		//line parser.go.y:498
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "+=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].tok.GetPos())
		}
	case 87:
		//line parser.go.y:503
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "-=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].tok.GetPos())
		}
	case 88:
		//line parser.go.y:508
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "*=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].tok.GetPos())
		}
	case 89:
		//line parser.go.y:513
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "/=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].tok.GetPos())
		}
	case 90:
		//line parser.go.y:518
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "&=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].tok.GetPos())
		}
	case 91:
		//line parser.go.y:523
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "|=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].tok.GetPos())
		}
	case 92:
		//line parser.go.y:528
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
			yyVAL.expr.SetPos(yyS[yypt-1].tok.GetPos())
		}
	case 93:
		//line parser.go.y:533
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
			yyVAL.expr.SetPos(yyS[yypt-1].tok.GetPos())
		}
	case 94:
		//line parser.go.y:538
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 95:
		//line parser.go.y:543
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 96:
		//line parser.go.y:548
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 97:
		//line parser.go.y:553
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPos(yyS[yypt-2].expr.GetPos())
		}
	case 98:
		//line parser.go.y:558
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPos(yyS[yypt-3].tok.GetPos())
		}
	case 99:
		//line parser.go.y:563
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyS[yypt-3].tok.lit}, Index: yyS[yypt-1].expr}
			yyVAL.expr.SetPos(yyS[yypt-3].tok.GetPos())
		}
	case 100:
		//line parser.go.y:568
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			yyVAL.expr.SetPos(yyS[yypt-3].expr.GetPos())
		}
	case 101:
		//line parser.go.y:573
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPos(yyS[yypt-3].expr.GetPos())
		}
	}
	goto yystack /* stack new state and value */
}
