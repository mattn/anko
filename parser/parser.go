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
	expr_many    []ast.Expr
	expr_lets    ast.Expr
	expr_pair    ast.Expr
	expr_pairs   []ast.Expr
	expr_idents  []string
	tok          ast.Token
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
	"'='",
	"'?'",
	"':'",
	"','",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:572

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 3,
	47, 40,
	-2, 4,
	-1, 76,
	61, 1,
	-2, 33,
	-1, 87,
	17, 47,
	18, 47,
	19, 47,
	20, 47,
	21, 47,
	22, 47,
	31, 47,
	32, 47,
	33, 47,
	34, 47,
	35, 47,
	36, 47,
	39, 47,
	40, 47,
	41, 47,
	42, 47,
	43, 47,
	48, 47,
	51, 47,
	52, 47,
	53, 47,
	54, 47,
	55, 47,
	56, 47,
	57, 47,
	64, 47,
	65, 47,
	70, 47,
	-2, 52,
	-1, 89,
	17, 47,
	18, 47,
	19, 47,
	20, 47,
	21, 47,
	22, 47,
	31, 47,
	32, 47,
	33, 47,
	34, 47,
	35, 47,
	36, 47,
	39, 47,
	40, 47,
	41, 47,
	42, 47,
	43, 47,
	48, 47,
	51, 47,
	52, 47,
	53, 47,
	54, 47,
	55, 47,
	56, 47,
	57, 47,
	64, 47,
	65, 47,
	70, 47,
	-2, 54,
	-1, 113,
	17, 0,
	18, 0,
	-2, 78,
	-1, 114,
	17, 0,
	18, 0,
	-2, 79,
}

const yyNprod = 100
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1509

var yyAct = []int{

	1, 178, 179, 33, 3, 72, 32, 95, 153, 223,
	135, 67, 69, 135, 211, 132, 77, 82, 80, 83,
	208, 81, 84, 85, 86, 88, 90, 209, 182, 138,
	186, 82, 67, 83, 97, 101, 99, 82, 158, 83,
	67, 103, 164, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 81, 92, 125, 126, 127, 128, 129,
	67, 152, 154, 180, 181, 250, 249, 139, 243, 82,
	143, 83, 66, 155, 241, 240, 67, 67, 147, 202,
	52, 53, 54, 55, 56, 57, 238, 150, 58, 59,
	43, 237, 236, 230, 225, 161, 224, 160, 212, 199,
	93, 197, 38, 39, 40, 41, 42, 195, 102, 176,
	216, 136, 173, 62, 37, 65, 91, 64, 246, 60,
	245, 233, 227, 222, 167, 221, 67, 67, 171, 67,
	207, 170, 174, 172, 133, 79, 142, 175, 130, 52,
	53, 54, 55, 56, 57, 206, 156, 100, 188, 43,
	193, 191, 190, 67, 145, 146, 134, 194, 137, 135,
	180, 181, 244, 40, 41, 42, 217, 242, 201, 203,
	204, 74, 62, 37, 65, 205, 64, 210, 60, 14,
	187, 96, 189, 185, 184, 169, 151, 213, 104, 214,
	98, 78, 73, 71, 94, 215, 177, 219, 220, 10,
	2, 0, 0, 0, 165, 166, 0, 168, 228, 229,
	0, 0, 231, 232, 0, 0, 0, 234, 235, 0,
	0, 0, 0, 0, 239, 46, 47, 49, 51, 61,
	63, 192, 0, 0, 0, 0, 247, 248, 0, 52,
	53, 54, 55, 56, 57, 0, 0, 58, 59, 43,
	44, 45, 0, 0, 0, 0, 36, 0, 35, 48,
	50, 38, 39, 40, 41, 42, 0, 0, 141, 0,
	0, 0, 62, 37, 65, 0, 64, 0, 60, 46,
	47, 49, 51, 61, 63, 0, 0, 0, 0, 0,
	0, 0, 0, 52, 53, 54, 55, 56, 57, 0,
	0, 58, 59, 43, 44, 45, 0, 0, 0, 0,
	36, 0, 0, 48, 50, 38, 39, 40, 41, 42,
	0, 0, 226, 0, 0, 0, 62, 37, 65, 0,
	64, 0, 60, 46, 47, 49, 51, 61, 63, 0,
	0, 0, 0, 0, 0, 0, 0, 52, 53, 54,
	55, 56, 57, 0, 0, 58, 59, 43, 44, 45,
	0, 0, 0, 0, 36, 218, 0, 48, 50, 38,
	39, 40, 41, 42, 46, 47, 49, 51, 61, 63,
	62, 37, 65, 0, 64, 0, 60, 0, 52, 53,
	54, 55, 56, 57, 0, 0, 58, 59, 43, 44,
	45, 0, 0, 0, 0, 36, 0, 0, 48, 50,
	38, 39, 40, 41, 42, 0, 200, 0, 0, 0,
	0, 62, 37, 65, 0, 64, 0, 60, 46, 47,
	49, 51, 61, 63, 0, 0, 0, 0, 0, 0,
	0, 0, 52, 53, 54, 55, 56, 57, 0, 0,
	58, 59, 43, 44, 45, 0, 0, 0, 0, 36,
	0, 0, 48, 50, 38, 39, 40, 41, 42, 0,
	0, 198, 0, 0, 0, 62, 37, 65, 0, 64,
	0, 60, 46, 47, 49, 51, 61, 63, 0, 0,
	0, 0, 0, 0, 0, 0, 52, 53, 54, 55,
	56, 57, 0, 0, 58, 59, 43, 44, 45, 0,
	0, 0, 0, 36, 0, 0, 48, 50, 38, 39,
	40, 41, 42, 0, 0, 196, 0, 0, 0, 62,
	37, 65, 0, 64, 0, 60, 46, 47, 49, 51,
	61, 63, 0, 0, 0, 0, 0, 0, 0, 0,
	52, 53, 54, 55, 56, 57, 0, 0, 58, 59,
	43, 44, 45, 0, 0, 0, 0, 36, 0, 35,
	48, 50, 38, 39, 40, 41, 42, 46, 47, 49,
	51, 61, 63, 62, 37, 65, 0, 64, 0, 60,
	0, 52, 53, 54, 55, 56, 57, 0, 0, 58,
	59, 43, 44, 45, 0, 0, 0, 0, 36, 0,
	0, 48, 50, 38, 39, 40, 41, 42, 46, 47,
	49, 51, 61, 63, 62, 37, 65, 0, 64, 183,
	60, 0, 52, 53, 54, 55, 56, 57, 0, 0,
	58, 59, 43, 44, 45, 0, 0, 0, 0, 36,
	0, 0, 48, 50, 38, 39, 40, 41, 42, 46,
	47, 49, 51, 61, 63, 62, 37, 65, 0, 64,
	163, 60, 0, 52, 53, 54, 55, 56, 57, 0,
	0, 58, 59, 43, 44, 45, 0, 0, 0, 0,
	36, 162, 0, 48, 50, 38, 39, 40, 41, 42,
	46, 47, 49, 51, 61, 63, 62, 37, 65, 0,
	64, 0, 60, 0, 52, 53, 54, 55, 56, 57,
	0, 0, 58, 59, 43, 44, 45, 0, 0, 0,
	0, 36, 0, 0, 48, 50, 38, 39, 40, 41,
	42, 0, 0, 159, 0, 0, 0, 62, 37, 65,
	0, 64, 0, 60, 46, 47, 49, 51, 61, 63,
	0, 0, 0, 0, 0, 0, 0, 0, 52, 53,
	54, 55, 56, 57, 0, 0, 58, 59, 43, 44,
	45, 0, 0, 0, 0, 36, 0, 0, 48, 50,
	38, 39, 40, 41, 42, 46, 47, 49, 51, 61,
	63, 62, 37, 65, 157, 64, 0, 60, 0, 52,
	53, 54, 55, 56, 57, 0, 0, 58, 59, 43,
	44, 45, 0, 0, 0, 0, 36, 0, 0, 48,
	50, 38, 39, 40, 41, 42, 0, 0, 144, 0,
	0, 0, 62, 37, 65, 0, 64, 0, 60, 46,
	47, 49, 51, 61, 63, 0, 0, 0, 0, 0,
	0, 0, 0, 52, 53, 54, 55, 56, 57, 0,
	0, 58, 59, 43, 44, 45, 0, 0, 0, 0,
	36, 0, 131, 48, 50, 38, 39, 40, 41, 42,
	46, 47, 49, 51, 61, 63, 62, 37, 65, 0,
	64, 0, 60, 0, 52, 53, 54, 55, 56, 57,
	0, 0, 58, 59, 43, 44, 45, 0, 0, 0,
	0, 36, 0, 0, 48, 50, 38, 39, 40, 41,
	42, 46, 47, 49, 51, 61, 63, 62, 37, 65,
	0, 64, 0, 60, 0, 52, 53, 54, 55, 56,
	57, 0, 0, 58, 59, 43, 44, 45, 0, 0,
	0, 0, 36, 0, 0, 48, 50, 38, 39, 40,
	41, 42, 46, 47, 49, 51, 61, 63, 62, 149,
	65, 0, 64, 0, 60, 0, 52, 53, 54, 55,
	56, 57, 0, 0, 58, 59, 43, 44, 45, 0,
	0, 0, 0, 36, 0, 0, 48, 50, 38, 39,
	40, 41, 42, 46, 47, 49, 51, 0, 63, 62,
	148, 65, 0, 64, 0, 60, 0, 52, 53, 54,
	55, 56, 57, 0, 0, 58, 59, 43, 44, 45,
	0, 0, 0, 0, 0, 0, 0, 48, 50, 38,
	39, 40, 41, 42, 46, 47, 49, 51, 0, 0,
	62, 37, 65, 0, 64, 0, 60, 0, 52, 53,
	54, 55, 56, 57, 0, 0, 58, 59, 43, 44,
	45, 0, 0, 0, 0, 0, 0, 0, 48, 50,
	38, 39, 40, 41, 42, 0, 49, 51, 0, 0,
	0, 62, 37, 65, 0, 64, 0, 60, 52, 53,
	54, 55, 56, 57, 0, 0, 58, 59, 43, 44,
	45, 0, 0, 0, 0, 0, 0, 0, 48, 50,
	38, 39, 40, 41, 42, 0, 0, 0, 0, 0,
	0, 62, 37, 65, 0, 64, 0, 60, 15, 16,
	22, 0, 0, 26, 6, 9, 7, 31, 0, 11,
	0, 0, 0, 0, 0, 0, 0, 30, 23, 24,
	25, 8, 12, 0, 0, 0, 0, 0, 0, 0,
	0, 4, 5, 0, 0, 0, 0, 0, 13, 0,
	0, 0, 52, 53, 54, 55, 56, 57, 17, 21,
	0, 0, 43, 34, 28, 0, 18, 19, 20, 0,
	29, 0, 27, 15, 16, 22, 0, 0, 26, 6,
	9, 7, 31, 0, 11, 62, 37, 65, 0, 64,
	0, 60, 30, 23, 24, 25, 8, 12, 0, 0,
	0, 0, 0, 0, 0, 0, 4, 5, 0, 0,
	0, 0, 0, 13, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 17, 21, 0, 0, 0, 0, 28,
	0, 18, 19, 20, 0, 29, 0, 27, 15, 16,
	140, 0, 0, 26, 6, 9, 7, 31, 0, 11,
	0, 0, 0, 0, 0, 0, 0, 30, 23, 24,
	25, 8, 12, 70, 16, 22, 0, 0, 26, 0,
	0, 4, 5, 0, 0, 0, 0, 0, 13, 0,
	0, 0, 30, 23, 24, 25, 0, 0, 17, 21,
	68, 16, 22, 0, 28, 26, 18, 19, 20, 0,
	29, 0, 27, 0, 0, 0, 0, 0, 0, 30,
	23, 24, 25, 17, 21, 15, 16, 22, 0, 28,
	26, 18, 19, 20, 0, 29, 0, 27, 0, 0,
	0, 0, 0, 0, 30, 23, 24, 25, 0, 0,
	17, 21, 89, 16, 22, 0, 28, 26, 18, 19,
	20, 0, 29, 0, 27, 0, 0, 0, 0, 0,
	0, 30, 23, 24, 25, 17, 21, 87, 16, 22,
	0, 28, 26, 18, 19, 20, 0, 29, 0, 27,
	0, 0, 0, 0, 0, 0, 30, 23, 24, 25,
	0, 0, 17, 21, 75, 16, 22, 0, 28, 26,
	18, 19, 20, 0, 29, 0, 27, 0, 0, 0,
	0, 0, 0, 30, 23, 24, 25, 17, 21, 0,
	0, 0, 0, 28, 0, 18, 19, 20, 0, 29,
	0, 27, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 17, 21, 0, 0, 0, 0,
	76, 0, 18, 19, 20, 0, 29, 0, 27,
}
var yyPact = []int{

	1219, -1000, 1154, 529, -1000, -1000, 1336, 1309, 199, 198,
	167, 1440, 85, 1309, -1000, -29, -1000, 1309, 1309, 1309,
	1413, 1388, -1000, -1000, -1000, -1000, 60, 1336, 185, 1309,
	196, 1309, 110, -1000, 1219, 1336, 1309, 194, 1309, 1309,
	1309, 1309, 1309, 1309, 1309, 1309, 1309, 1309, 1309, 1309,
	1309, 1309, 1309, 1309, 1309, 1309, 1309, 1309, -1000, -1000,
	1309, 1309, 1309, 1309, 1309, 1336, -1000, 842, -35, 883,
	-49, 84, 119, -1000, 108, 13, 1284, 218, 87, 1219,
	788, 1336, 1336, 1309, 1171, 1171, 1171, -49, 965, -49,
	924, 192, 5, -61, 22, -1000, 107, 747, -28, 693,
	1361, -1000, -1000, 652, -1000, 118, 118, 1171, 1171, 1171,
	883, 59, 59, 1087, 1087, 59, 59, 59, 59, 883,
	883, 883, 883, 883, 883, 883, 1006, 883, 1047, 611,
	-25, 1336, 1336, 1219, 1336, 191, 1309, 1219, 1309, 61,
	107, 1219, 1309, 58, 125, -1000, -39, 570, 190, 189,
	-37, 182, 188, -1000, 185, -1000, 1309, -1000, 1336, 1219,
	-1000, 529, 1309, -1000, -1000, -1000, -1000, 56, -1000, -1000,
	475, 50, 421, -1000, 48, 367, 149, 28, -1000, -1000,
	1309, 106, -1000, -1000, -1000, -1000, 80, -47, -40, 179,
	-1000, 883, -53, 47, 883, -1000, 1219, -1000, 1219, -1000,
	1309, 116, -1000, -1000, -1000, 326, 1219, 1219, 75, 73,
	-58, -1000, -1000, 45, 43, 272, 72, 1219, 1219, -1000,
	42, 1219, 1219, 71, -1000, -1000, 1219, 1219, 41, -1000,
	-1000, 40, 35, 1219, 24, 23, 147, -1000, -1000, 17,
	-1000, 142, 70, -1000, 68, 1219, 1219, 15, 14, -1000,
	-1000,
}
var yyPgo = []int{

	0, 0, 210, 209, 2, 1, 206, 4, 82, 6,
	189, 7, 204, 5,
}
var yyR1 = []int{

	0, 1, 1, 1, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 3, 3, 3, 6, 6, 6, 6, 6,
	5, 4, 11, 12, 12, 12, 13, 13, 13, 10,
	9, 9, 9, 8, 8, 8, 8, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 1, 1, 1, 2, 2, 5,
	4, 1, 7, 4, 5, 9, 13, 12, 9, 8,
	5, 1, 7, 5, 5, 0, 1, 2, 1, 2,
	4, 3, 3, 0, 1, 3, 0, 1, 3, 3,
	1, 3, 3, 0, 1, 3, 3, 1, 1, 2,
	2, 2, 2, 4, 2, 4, 1, 1, 1, 1,
	5, 3, 7, 8, 8, 9, 3, 3, 3, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 3, 3, 3, 3, 4, 4, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, -7, 37, 38, 10, 12, 27, 11,
	-3, 15, 28, 44, -10, 4, 5, 54, 62, 63,
	64, 55, 6, 24, 25, 26, 9, 68, 60, 66,
	23, 13, -9, -1, 59, 50, 48, 65, 53, 54,
	55, 56, 57, 41, 42, 43, 17, 18, 51, 19,
	52, 20, 31, 32, 33, 34, 35, 36, 39, 40,
	70, 21, 64, 22, 68, 66, -8, -7, 4, -7,
	4, 4, -13, 4, 14, 4, 60, -7, -10, 60,
	-7, 50, 66, 68, -7, -7, -7, 4, -7, 4,
	-7, 66, 4, -8, -12, -11, 6, -7, 4, -7,
	47, -1, -8, -7, 4, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-8, 50, 50, 60, 47, 50, 13, 60, 16, -1,
	6, 60, 59, -1, 60, -8, -8, -7, 65, 65,
	-13, 4, 66, 69, 50, 61, 49, 67, 66, 60,
	-9, -7, 49, 69, 67, -8, -8, -1, -8, 4,
	-7, -1, -7, 61, -1, -7, 61, -6, -5, -4,
	45, 46, 67, 69, 4, 4, 67, 8, -13, 4,
	-11, -7, -8, -1, -7, 61, 60, 61, 60, 61,
	59, 29, 61, -5, -4, -7, 49, 60, 67, 67,
	8, 67, 61, -1, -1, -7, 4, 60, 49, -1,
	-1, 60, 60, 67, 61, 61, 60, 60, -1, -1,
	61, -1, -1, 60, -1, -1, 61, 61, 61, -1,
	61, 61, 30, 61, 30, 60, 60, -1, -1, 61,
	61,
}
var yyDef = []int{

	1, -2, 1, -2, 5, 6, 43, 0, 0, 36,
	11, 0, 0, 0, 21, 47, 48, 0, 0, 0,
	0, 0, 56, 57, 58, 59, 0, 43, 33, 0,
	0, 0, 0, 2, 1, 43, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 90, 91,
	0, 0, 0, 0, 0, 43, 7, 44, 47, 8,
	47, 0, 0, 37, 0, 47, -2, 40, 0, 1,
	0, 43, 43, 0, 49, 50, 51, -2, 0, -2,
	0, 36, 0, 0, 0, 34, 0, 0, 0, 0,
	0, 3, 41, 0, 61, 70, 71, 72, 73, 74,
	75, 76, 77, -2, -2, 80, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 92, 93, 94, 95, 0,
	0, 43, 43, 1, 43, 0, 0, 1, 0, 0,
	56, 1, 0, 0, 25, 42, 0, 0, 0, 0,
	0, 37, 36, 66, 0, 67, 0, 68, 43, 1,
	39, 40, 0, 98, 99, 45, 46, 0, 10, 38,
	0, 0, 0, 13, 0, 0, 0, 0, 26, 28,
	0, 0, 96, 97, 53, 55, 0, 0, 0, 37,
	35, 32, 0, 0, 60, 9, 1, 23, 1, 14,
	0, 0, 20, 27, 29, 0, 1, 1, 0, 0,
	0, 69, 24, 0, 0, 0, 0, 1, 1, 31,
	0, 1, 1, 0, 22, 12, 1, 1, 0, 30,
	62, 0, 0, 1, 0, 0, 19, 63, 64, 0,
	15, 18, 0, 65, 0, 1, 1, 0, 0, 17,
	16,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 62, 3, 3, 3, 57, 64, 3,
	66, 67, 55, 53, 50, 54, 65, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 59,
	52, 47, 51, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 68, 3, 69, 63, 3, 3, 3, 3, 3,
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
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
					if yyS[yypt-1].stmt.Position().Line == s.Position().Line {
						l.pos = yyS[yypt-1].stmt.Position()
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
				yyS[yypt-2].stmt.SetPosition(l.pos)
			}
		}
	case 4:
		//line parser.go.y:87
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			yyVAL.stmt.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 5:
		//line parser.go.y:92
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 6:
		//line parser.go.y:97
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 7:
		//line parser.go.y:102
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
			yyVAL.stmt.SetPosition(yyS[yypt-1].tok.Position())
		}
	case 8:
		//line parser.go.y:107
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			yyVAL.stmt.SetPosition(yyS[yypt-1].tok.Position())
		}
	case 9:
		//line parser.go.y:112
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.Lit, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-4].tok.Position())
		}
	case 10:
		//line parser.go.y:117
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyS[yypt-2].expr_idents, Exprs: yyS[yypt-0].exprs}
			yyVAL.stmt.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 11:
		//line parser.go.y:122
		{
			yyVAL.stmt = yyS[yypt-0].stmt_if
			yyVAL.stmt.SetPosition(yyS[yypt-0].stmt_if.Position())
		}
	case 12:
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyS[yypt-5].tok.Lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-6].tok.Position())
		}
	case 13:
		//line parser.go.y:132
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 14:
		//line parser.go.y:137
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-4].tok.Position())
		}
	case 15:
		//line parser.go.y:142
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyS[yypt-7].expr_lets, Expr2: yyS[yypt-5].expr, Expr3: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-8].tok.Position())
		}
	case 16:
		//line parser.go.y:147
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-10].stmts, Var: yyS[yypt-7].tok.Lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-12].tok.Position())
		}
	case 17:
		//line parser.go.y:152
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-11].tok.Position())
		}
	case 18:
		//line parser.go.y:157
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-6].stmts, Var: yyS[yypt-3].tok.Lit, Catch: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-8].tok.Position())
		}
	case 19:
		//line parser.go.y:162
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-7].tok.Position())
		}
	case 20:
		//line parser.go.y:167
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyS[yypt-3].expr, Cases: yyS[yypt-1].stmt_cases}
			yyVAL.stmt.SetPosition(yyS[yypt-4].tok.Position())
		}
	case 21:
		//line parser.go.y:172
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr_lets}
			yyVAL.stmt.SetPosition(yyS[yypt-0].expr_lets.Position())
		}
	case 22:
		//line parser.go.y:178
		{
			yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts})
			yyVAL.stmt_if.SetPosition(yyS[yypt-6].stmt_if.Position())
		}
	case 23:
		//line parser.go.y:183
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyS[yypt-1].stmts...)
			}
			yyVAL.stmt_if.SetPosition(yyS[yypt-4].stmt_if.Position())
		}
	case 24:
		//line parser.go.y:192
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts, Else: nil}
			yyVAL.stmt_if.SetPosition(yyS[yypt-4].tok.Position())
		}
	case 25:
		//line parser.go.y:198
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 26:
		//line parser.go.y:202
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_case}
		}
	case 27:
		//line parser.go.y:206
		{
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_case)
		}
	case 28:
		//line parser.go.y:210
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_default}
		}
	case 29:
		//line parser.go.y:214
		{
			for _, stmt := range yyS[yypt-1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_default)
		}
	case 30:
		//line parser.go.y:224
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyS[yypt-2].expr, Stmts: yyS[yypt-0].stmts}
		}
	case 31:
		//line parser.go.y:229
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyS[yypt-0].stmts}
		}
	case 32:
		//line parser.go.y:234
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyS[yypt-2].tok.Lit, Value: yyS[yypt-0].expr}
		}
	case 33:
		//line parser.go.y:239
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 34:
		//line parser.go.y:243
		{
			yyVAL.expr_pairs = []ast.Expr{yyS[yypt-0].expr_pair}
		}
	case 35:
		//line parser.go.y:247
		{
			yyVAL.expr_pairs = append(yyS[yypt-2].expr_pairs, yyS[yypt-0].expr_pair)
		}
	case 36:
		//line parser.go.y:252
		{
			yyVAL.expr_idents = []string{}
		}
	case 37:
		//line parser.go.y:256
		{
			yyVAL.expr_idents = []string{yyS[yypt-0].tok.Lit}
		}
	case 38:
		//line parser.go.y:260
		{
			yyVAL.expr_idents = append(yyS[yypt-2].expr_idents, yyS[yypt-0].tok.Lit)
		}
	case 39:
		//line parser.go.y:265
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyS[yypt-2].expr_many, Operator: "=", Rhss: yyS[yypt-0].expr_many}
			yyVAL.expr_lets.SetPosition(yyS[yypt-2].expr_many[0].Position())
		}
	case 40:
		//line parser.go.y:272
		{
			yyVAL.expr_many = []ast.Expr{yyS[yypt-0].expr}
		}
	case 41:
		//line parser.go.y:276
		{
			yyVAL.expr_many = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 42:
		//line parser.go.y:280
		{
			yyVAL.expr_many = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.Lit}}, yyS[yypt-0].exprs...)
		}
	case 43:
		//line parser.go.y:285
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 44:
		//line parser.go.y:289
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 45:
		//line parser.go.y:293
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 46:
		//line parser.go.y:297
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.Lit}}, yyS[yypt-0].exprs...)
		}
	case 47:
		//line parser.go.y:303
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 48:
		//line parser.go.y:308
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 49:
		//line parser.go.y:313
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 50:
		//line parser.go.y:318
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 51:
		//line parser.go.y:323
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 52:
		//line parser.go.y:328
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.Lit}}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 53:
		//line parser.go.y:333
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.Lit}}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 54:
		//line parser.go.y:338
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.Lit}}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 55:
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.Lit}}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 56:
		//line parser.go.y:348
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 57:
		//line parser.go.y:353
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 58:
		//line parser.go.y:358
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 59:
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 60:
		//line parser.go.y:368
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-4].expr.Position())
		}
	case 61:
		//line parser.go.y:373
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 62:
		//line parser.go.y:378
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			yyVAL.expr.SetPosition(yyS[yypt-6].tok.Position())
		}
	case 63:
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.Lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			yyVAL.expr.SetPosition(yyS[yypt-7].tok.Position())
		}
	case 64:
		//line parser.go.y:388
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.Lit, Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			yyVAL.expr.SetPosition(yyS[yypt-7].tok.Position())
		}
	case 65:
		//line parser.go.y:393
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.Lit, Args: []string{yyS[yypt-5].tok.Lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			yyVAL.expr.SetPosition(yyS[yypt-8].tok.Position())
		}
	case 66:
		//line parser.go.y:398
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 67:
		//line parser.go.y:403
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 68:
		//line parser.go.y:412
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 69:
		//line parser.go.y:417
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.Lit, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPosition(yyS[yypt-4].tok.Position())
		}
	case 70:
		//line parser.go.y:422
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 71:
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 72:
		//line parser.go.y:432
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 73:
		//line parser.go.y:437
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 74:
		//line parser.go.y:442
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 75:
		//line parser.go.y:447
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 76:
		//line parser.go.y:452
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 77:
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 78:
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 79:
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 80:
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 81:
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 82:
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 83:
		//line parser.go.y:487
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 84:
		//line parser.go.y:492
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "+=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 85:
		//line parser.go.y:497
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "-=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 86:
		//line parser.go.y:502
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "*=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 87:
		//line parser.go.y:507
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "/=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 88:
		//line parser.go.y:512
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "&=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 89:
		//line parser.go.y:517
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-2].expr, Operator: "|=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 90:
		//line parser.go.y:522
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyS[yypt-1].expr.Position())
		}
	case 91:
		//line parser.go.y:527
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyS[yypt-1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyS[yypt-1].expr.Position())
		}
	case 92:
		//line parser.go.y:532
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 93:
		//line parser.go.y:537
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 94:
		//line parser.go.y:542
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 95:
		//line parser.go.y:547
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 96:
		//line parser.go.y:552
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.Lit, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 97:
		//line parser.go.y:557
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyS[yypt-3].tok.Lit}, Index: yyS[yypt-1].expr}
			yyVAL.expr.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 98:
		//line parser.go.y:562
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			yyVAL.expr.SetPosition(yyS[yypt-3].expr.Position())
		}
	case 99:
		//line parser.go.y:567
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPosition(yyS[yypt-3].expr.Position())
		}
	}
	goto yystack /* stack new state and value */
}
