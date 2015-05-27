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

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
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
	"';'",
	"'{'",
	"'}'",
	"'!'",
	"'^'",
	"'&'",
	"'.'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'|'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:571

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 3,
	47, 40,
	-2, 4,
	-1, 75,
	61, 1,
	-2, 33,
	-1, 88,
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
	-1, 90,
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

const yyLast = 1583

var yyAct = [...]int{

	1, 180, 181, 32, 96, 71, 14, 132, 3, 156,
	83, 161, 84, 225, 135, 66, 68, 82, 78, 213,
	76, 210, 80, 83, 135, 84, 85, 86, 87, 89,
	91, 211, 138, 83, 101, 84, 66, 184, 98, 165,
	100, 188, 93, 66, 103, 155, 105, 106, 107, 108,
	109, 110, 111, 112, 113, 114, 115, 116, 117, 118,
	119, 120, 121, 122, 123, 124, 82, 252, 125, 126,
	127, 128, 129, 66, 157, 251, 139, 218, 182, 183,
	144, 65, 83, 245, 84, 158, 243, 242, 146, 240,
	147, 66, 66, 150, 204, 239, 238, 232, 153, 246,
	227, 226, 214, 201, 92, 199, 197, 178, 174, 94,
	248, 136, 247, 235, 229, 224, 102, 223, 209, 133,
	51, 52, 53, 54, 55, 56, 79, 142, 134, 208,
	42, 135, 159, 219, 168, 182, 183, 244, 172, 143,
	66, 66, 175, 66, 81, 171, 130, 173, 203, 73,
	177, 176, 147, 61, 36, 64, 212, 63, 137, 59,
	189, 190, 192, 195, 148, 149, 97, 191, 193, 187,
	66, 186, 196, 170, 154, 104, 99, 72, 70, 95,
	77, 205, 206, 179, 10, 2, 0, 0, 0, 0,
	0, 207, 0, 0, 0, 0, 0, 0, 0, 215,
	0, 216, 0, 0, 0, 0, 0, 0, 0, 221,
	222, 217, 0, 166, 167, 0, 169, 0, 0, 0,
	230, 231, 0, 0, 233, 234, 0, 0, 0, 236,
	237, 0, 0, 0, 0, 0, 241, 45, 46, 48,
	50, 60, 62, 194, 0, 0, 0, 0, 249, 250,
	0, 51, 52, 53, 54, 55, 56, 0, 0, 57,
	58, 42, 43, 44, 0, 0, 0, 0, 35, 0,
	34, 47, 49, 37, 38, 39, 40, 41, 0, 0,
	141, 0, 0, 0, 61, 36, 64, 0, 63, 0,
	59, 45, 46, 48, 50, 60, 62, 0, 0, 0,
	0, 0, 0, 0, 0, 51, 52, 53, 54, 55,
	56, 0, 0, 57, 58, 42, 43, 44, 0, 0,
	0, 0, 35, 0, 0, 47, 49, 37, 38, 39,
	40, 41, 0, 0, 228, 0, 0, 0, 61, 36,
	64, 0, 63, 0, 59, 45, 46, 48, 50, 60,
	62, 0, 0, 0, 0, 0, 0, 0, 0, 51,
	52, 53, 54, 55, 56, 0, 0, 57, 58, 42,
	43, 44, 0, 0, 0, 0, 35, 220, 0, 47,
	49, 37, 38, 39, 40, 41, 45, 46, 48, 50,
	60, 62, 61, 36, 64, 0, 63, 0, 59, 0,
	51, 52, 53, 54, 55, 56, 0, 0, 57, 58,
	42, 43, 44, 0, 0, 0, 0, 35, 0, 0,
	47, 49, 37, 38, 39, 40, 41, 0, 202, 0,
	0, 0, 0, 61, 36, 64, 0, 63, 0, 59,
	45, 46, 48, 50, 60, 62, 0, 0, 0, 0,
	0, 0, 0, 0, 51, 52, 53, 54, 55, 56,
	0, 0, 57, 58, 42, 43, 44, 0, 0, 0,
	0, 35, 0, 0, 47, 49, 37, 38, 39, 40,
	41, 0, 0, 200, 0, 0, 0, 61, 36, 64,
	0, 63, 0, 59, 45, 46, 48, 50, 60, 62,
	0, 0, 0, 0, 0, 0, 0, 0, 51, 52,
	53, 54, 55, 56, 0, 0, 57, 58, 42, 43,
	44, 0, 0, 0, 0, 35, 0, 0, 47, 49,
	37, 38, 39, 40, 41, 0, 0, 198, 0, 0,
	0, 61, 36, 64, 0, 63, 0, 59, 45, 46,
	48, 50, 60, 62, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 52, 53, 54, 55, 56, 0, 0,
	57, 58, 42, 43, 44, 0, 0, 0, 0, 35,
	0, 0, 47, 49, 37, 38, 39, 40, 41, 45,
	46, 48, 50, 60, 62, 61, 36, 64, 0, 63,
	185, 59, 0, 51, 52, 53, 54, 55, 56, 0,
	0, 57, 58, 42, 43, 44, 0, 0, 0, 0,
	35, 0, 34, 47, 49, 37, 38, 39, 40, 41,
	45, 46, 48, 50, 60, 62, 61, 36, 64, 0,
	63, 0, 59, 0, 51, 52, 53, 54, 55, 56,
	0, 0, 57, 58, 42, 43, 44, 0, 0, 0,
	0, 35, 0, 0, 47, 49, 37, 38, 39, 40,
	41, 45, 46, 48, 50, 60, 62, 61, 36, 64,
	0, 63, 164, 59, 0, 51, 52, 53, 54, 55,
	56, 0, 0, 57, 58, 42, 43, 44, 0, 0,
	0, 0, 35, 163, 0, 47, 49, 37, 38, 39,
	40, 41, 45, 46, 48, 50, 60, 62, 61, 36,
	64, 0, 63, 0, 59, 0, 51, 52, 53, 54,
	55, 56, 0, 0, 57, 58, 42, 43, 44, 0,
	0, 0, 0, 35, 0, 0, 47, 49, 37, 38,
	39, 40, 41, 0, 0, 162, 0, 0, 0, 61,
	36, 64, 0, 63, 0, 59, 45, 46, 48, 50,
	60, 62, 0, 0, 0, 0, 0, 0, 0, 0,
	51, 52, 53, 54, 55, 56, 0, 0, 57, 58,
	42, 43, 44, 0, 0, 0, 0, 35, 0, 0,
	47, 49, 37, 38, 39, 40, 41, 45, 46, 48,
	50, 60, 62, 61, 36, 64, 160, 63, 0, 59,
	0, 51, 52, 53, 54, 55, 56, 0, 0, 57,
	58, 42, 43, 44, 0, 0, 0, 0, 35, 0,
	0, 47, 49, 37, 38, 39, 40, 41, 0, 0,
	145, 0, 0, 0, 61, 36, 64, 0, 63, 0,
	59, 45, 46, 48, 50, 60, 62, 0, 0, 0,
	0, 0, 0, 0, 0, 51, 52, 53, 54, 55,
	56, 0, 0, 57, 58, 42, 43, 44, 0, 0,
	0, 0, 35, 0, 131, 47, 49, 37, 38, 39,
	40, 41, 45, 46, 48, 50, 60, 62, 61, 36,
	64, 0, 63, 0, 59, 0, 51, 52, 53, 54,
	55, 56, 0, 0, 57, 58, 42, 43, 44, 0,
	0, 0, 0, 35, 0, 0, 47, 49, 37, 38,
	39, 40, 41, 45, 46, 48, 50, 60, 62, 61,
	36, 64, 0, 63, 0, 59, 0, 51, 52, 53,
	54, 55, 56, 0, 0, 57, 58, 42, 43, 44,
	0, 0, 0, 0, 35, 0, 0, 47, 49, 37,
	38, 39, 40, 41, 45, 46, 48, 50, 60, 62,
	61, 152, 64, 0, 63, 0, 59, 0, 51, 52,
	53, 54, 55, 56, 0, 0, 57, 58, 42, 43,
	44, 0, 0, 0, 0, 35, 0, 0, 47, 49,
	37, 38, 39, 40, 41, 45, 46, 48, 50, 0,
	62, 61, 151, 64, 0, 63, 0, 59, 0, 51,
	52, 53, 54, 55, 56, 0, 0, 57, 58, 42,
	43, 44, 0, 0, 0, 0, 0, 0, 0, 47,
	49, 37, 38, 39, 40, 41, 45, 46, 48, 50,
	0, 0, 61, 36, 64, 0, 63, 0, 59, 0,
	51, 52, 53, 54, 55, 56, 0, 0, 57, 58,
	42, 43, 44, 0, 0, 0, 0, 0, 0, 0,
	47, 49, 37, 38, 39, 40, 41, 0, 48, 50,
	0, 0, 0, 61, 36, 64, 0, 63, 0, 59,
	51, 52, 53, 54, 55, 56, 0, 0, 57, 58,
	42, 43, 44, 0, 0, 0, 0, 0, 0, 0,
	47, 49, 37, 38, 39, 40, 41, 0, 0, 0,
	0, 0, 0, 61, 36, 64, 0, 63, 0, 59,
	15, 16, 22, 0, 0, 26, 6, 9, 7, 31,
	0, 11, 0, 0, 0, 0, 0, 0, 0, 30,
	23, 24, 25, 8, 12, 0, 0, 0, 0, 0,
	0, 0, 0, 4, 5, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	17, 21, 0, 0, 0, 33, 28, 0, 18, 19,
	20, 0, 29, 0, 27, 15, 16, 22, 0, 0,
	26, 6, 9, 7, 31, 0, 11, 0, 0, 0,
	0, 0, 0, 0, 30, 23, 24, 25, 8, 12,
	0, 0, 0, 0, 0, 0, 0, 0, 4, 5,
	0, 0, 0, 0, 0, 13, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 17, 21, 0, 0, 0,
	0, 28, 0, 18, 19, 20, 0, 29, 0, 27,
	15, 16, 140, 0, 0, 26, 6, 9, 7, 31,
	0, 11, 0, 0, 0, 0, 0, 0, 0, 30,
	23, 24, 25, 8, 12, 0, 0, 0, 0, 0,
	0, 0, 0, 4, 5, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 51, 52, 53, 54, 55, 56,
	17, 21, 57, 58, 42, 0, 28, 0, 18, 19,
	20, 0, 29, 0, 27, 0, 37, 38, 39, 40,
	41, 0, 0, 0, 0, 0, 0, 61, 36, 64,
	0, 63, 0, 59, 51, 52, 53, 54, 55, 56,
	0, 0, 0, 0, 42, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 39, 40,
	41, 69, 16, 22, 0, 0, 26, 61, 36, 64,
	0, 63, 0, 59, 0, 0, 0, 0, 0, 0,
	30, 23, 24, 25, 67, 16, 22, 0, 0, 26,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 30, 23, 24, 25, 0, 15, 16,
	22, 17, 21, 26, 0, 0, 0, 28, 0, 18,
	19, 20, 0, 29, 0, 27, 0, 30, 23, 24,
	25, 90, 16, 22, 17, 21, 26, 0, 0, 0,
	28, 0, 18, 19, 20, 0, 29, 0, 27, 0,
	30, 23, 24, 25, 0, 88, 16, 22, 17, 21,
	26, 0, 0, 0, 28, 0, 18, 19, 20, 0,
	29, 0, 27, 0, 30, 23, 24, 25, 74, 16,
	22, 17, 21, 26, 0, 0, 0, 28, 0, 18,
	19, 20, 0, 29, 0, 27, 0, 30, 23, 24,
	25, 0, 0, 0, 0, 17, 21, 0, 0, 0,
	0, 28, 0, 18, 19, 20, 0, 29, 0, 27,
	0, 0, 0, 0, 0, 0, 0, 0, 17, 21,
	0, 0, 0, 0, 75, 0, 18, 19, 20, 0,
	29, 0, 27,
}
var yyPact = [...]int{

	1221, -1000, 1156, 572, -1000, -1000, 1420, 1397, 174, 173,
	135, 1514, 66, 1397, 97, -33, -1000, 1397, 1397, 1397,
	1491, 1467, -1000, -1000, -1000, -1000, 38, 1420, 160, 1397,
	172, 1397, -1000, 1221, 1420, 1397, 171, 1397, 1397, 1397,
	1397, 1397, 1397, 1397, 1397, 1397, 1397, 1397, 1397, 1397,
	1397, 1397, 1397, 1397, 1397, 1397, 1397, -1000, -1000, 1397,
	1397, 1397, 1397, 1397, 1420, -1000, 844, -43, 885, -56,
	59, 81, -1000, 98, 16, 1286, 220, 68, 92, 1221,
	790, 1444, 1420, 1420, 1397, 89, 89, 89, -56, 967,
	-56, 926, 170, -21, -60, 24, -1000, 83, 749, -55,
	695, -1000, -1000, 654, -1000, 1343, 1343, 89, 89, 89,
	885, 1303, 1303, 1089, 1089, 1303, 1303, 1303, 1303, 885,
	885, 885, 885, 885, 885, 885, 1008, 885, 1049, 613,
	-28, 1420, 1420, 1221, 1420, 169, 1397, 1221, 1397, 47,
	83, 1221, 1397, 1444, 46, 90, -1000, 572, -1000, -30,
	531, 167, 165, -26, 152, 163, -1000, 160, -1000, 1397,
	-1000, 1420, 1221, 1397, -1000, -1000, -1000, -1000, 45, -1000,
	-1000, 477, 44, 423, -1000, 42, 369, -1000, 119, 33,
	-1000, -1000, 1397, 80, -1000, -1000, -1000, -1000, 58, -46,
	-36, 148, -1000, 885, -48, 41, 885, -1000, 1221, -1000,
	1221, -1000, 1397, 73, -1000, -1000, -1000, 328, 1221, 1221,
	57, 55, -54, -1000, -1000, 40, 39, 274, 54, 1221,
	1221, -1000, 36, 1221, 1221, 53, -1000, -1000, 1221, 1221,
	35, -1000, -1000, 34, 28, 1221, 26, 25, 107, -1000,
	-1000, 22, -1000, 69, 52, -1000, 50, 1221, 1221, 14,
	6, -1000, -1000,
}
var yyPgo = [...]int{

	0, 0, 185, 184, 2, 1, 183, 8, 81, 6,
	180, 4, 179, 5,
}
var yyR1 = [...]int{

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
var yyR2 = [...]int{

	0, 0, 2, 3, 1, 1, 1, 2, 2, 5,
	4, 1, 7, 4, 5, 9, 13, 12, 9, 8,
	5, 3, 7, 5, 5, 0, 1, 2, 1, 2,
	4, 3, 3, 0, 1, 3, 0, 1, 3, 3,
	1, 3, 3, 0, 1, 3, 3, 1, 1, 2,
	2, 2, 2, 4, 2, 4, 1, 1, 1, 1,
	5, 3, 7, 8, 8, 9, 3, 3, 3, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 3, 3, 3, 3, 4, 4, 4, 4,
}
var yyChk = [...]int{

	-1000, -1, -2, -7, 37, 38, 10, 12, 27, 11,
	-3, 15, 28, 44, -9, 4, 5, 54, 62, 63,
	64, 55, 6, 24, 25, 26, 9, 68, 60, 66,
	23, 13, -1, 59, 50, 48, 65, 53, 54, 55,
	56, 57, 41, 42, 43, 17, 18, 51, 19, 52,
	20, 31, 32, 33, 34, 35, 36, 39, 40, 70,
	21, 64, 22, 68, 66, -8, -7, 4, -7, 4,
	4, -13, 4, 14, 4, 60, -7, -10, -9, 60,
	-7, 47, 50, 66, 68, -7, -7, -7, 4, -7,
	4, -7, 66, 4, -8, -12, -11, 6, -7, 4,
	-7, -1, -8, -7, 4, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-8, 50, 50, 60, 47, 50, 13, 60, 16, -1,
	6, 60, 59, 47, -1, 60, -9, -7, -8, -8,
	-7, 65, 65, -13, 4, 66, 69, 50, 61, 49,
	67, 66, 60, 49, 69, 67, -8, -8, -1, -8,
	4, -7, -1, -7, 61, -1, -7, -9, 61, -6,
	-5, -4, 45, 46, 67, 69, 4, 4, 67, 8,
	-13, 4, -11, -7, -8, -1, -7, 61, 60, 61,
	60, 61, 59, 29, 61, -5, -4, -7, 49, 60,
	67, 67, 8, 67, 61, -1, -1, -7, 4, 60,
	49, -1, -1, 60, 60, 67, 61, 61, 60, 60,
	-1, -1, 61, -1, -1, 60, -1, -1, 61, 61,
	61, -1, 61, 61, 30, 61, 30, 60, 60, -1,
	-1, 61, 61,
}
var yyDef = [...]int{

	1, -2, 1, -2, 5, 6, 43, 0, 0, 36,
	11, 0, 0, 0, 0, 47, 48, 0, 0, 0,
	0, 0, 56, 57, 58, 59, 0, 43, 33, 0,
	0, 0, 2, 1, 43, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 90, 91, 0,
	0, 0, 0, 0, 43, 7, 44, 47, 8, 47,
	0, 0, 37, 0, 47, -2, 40, 0, 0, 1,
	0, 0, 43, 43, 0, 49, 50, 51, -2, 0,
	-2, 0, 36, 0, 0, 0, 34, 0, 0, 0,
	0, 3, 41, 0, 61, 70, 71, 72, 73, 74,
	75, 76, 77, -2, -2, 80, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 92, 93, 94, 95, 0,
	0, 43, 43, 1, 43, 0, 0, 1, 0, 0,
	56, 1, 0, 0, 0, 25, 21, 40, 42, 0,
	0, 0, 0, 0, 37, 36, 66, 0, 67, 0,
	68, 43, 1, 0, 98, 99, 45, 46, 0, 10,
	38, 0, 0, 0, 13, 0, 0, 39, 0, 0,
	26, 28, 0, 0, 96, 97, 53, 55, 0, 0,
	0, 37, 35, 32, 0, 0, 60, 9, 1, 23,
	1, 14, 0, 0, 20, 27, 29, 0, 1, 1,
	0, 0, 0, 69, 24, 0, 0, 0, 0, 1,
	1, 31, 0, 1, 1, 0, 22, 12, 1, 1,
	0, 30, 62, 0, 0, 1, 0, 0, 19, 63,
	64, 0, 15, 18, 0, 65, 0, 1, 1, 0,
	0, 17, 16,
}
var yyTok1 = [...]int{

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
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 58,
}
var yyTok3 = [...]int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
	state     func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
		state:     func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}
	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.state = func() int { return yystate }
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
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
			yychar, yytoken = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yytoken {
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
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
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
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:58
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:65
		{
			yyVAL.stmts = append([]ast.Stmt{yyDollar[1].stmt}, yyDollar[2].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				for _, s := range yyDollar[2].stmts {
					if yyDollar[1].stmt.Position().Line == s.Position().Line {
						l.pos = yyDollar[1].stmt.Position()
						yylex.Error("syntax error")
					}
				}
			}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:78
		{
			yyVAL.stmts = append([]ast.Stmt{yyDollar[1].stmt}, yyDollar[3].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyDollar[1].stmt.SetPosition(l.pos)
			}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:87
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:92
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:97
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].stmts}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:117
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:122
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyDollar[2].tok.Lit, Value: yyDollar[4].expr, Stmts: yyDollar[6].stmts}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:132
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:137
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].stmts}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:142
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].stmts, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].stmts, Finally: yyDollar[12].stmts}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].stmts, Catch: yyDollar[7].stmts, Finally: yyDollar[11].stmts}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:157
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].stmts, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].stmts}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].stmts, Catch: yyDollar[7].stmts}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:172
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 22:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:178
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].stmts})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:183
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyDollar[4].stmts...)
			}
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:192
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].stmts, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:198
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:202
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[1].stmt_case}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:206
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:210
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[1].stmt_default}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:214
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].stmts}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:229
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[3].stmts}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:247
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[3].expr_pair)
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:252
		{
			yyVAL.expr_idents = []string{}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:260
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[3].tok.Lit)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:265
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:275
		{
			yyVAL.expr_many = append([]ast.Expr{yyDollar[1].expr}, yyDollar[3].exprs...)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.expr_many = append([]ast.Expr{&ast.IdentExpr{Lit: yyDollar[1].tok.Lit}}, yyDollar[3].exprs...)
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:284
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:288
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:292
		{
			yyVAL.exprs = append([]ast.Expr{yyDollar[1].expr}, yyDollar[3].exprs...)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:296
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyDollar[1].tok.Lit}}, yyDollar[3].exprs...)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:307
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:312
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:317
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:322
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:332
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:337
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:342
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:352
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:357
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:362
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:372
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 62:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:377
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].stmts}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:382
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].stmts, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].stmts}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:392
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].stmts, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[2].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:402
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[2].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:416
		{
			yyVAL.expr = &ast.NewExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:421
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:426
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:431
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:436
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:441
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:446
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:451
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:456
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:461
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:466
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:471
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:476
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:481
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:486
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:491
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:496
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:501
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:506
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:511
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:516
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:521
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:526
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:531
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:536
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:541
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:546
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:551
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:556
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:561
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:566
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	}
	goto yystack /* stack new state and value */
}
