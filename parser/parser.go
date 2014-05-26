//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:25
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

//line parser.go.y:593

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 15,
	47, 39,
	50, 39,
	-2, 49,
	-1, 67,
	47, 39,
	50, 39,
	-2, 49,
	-1, 68,
	61, 1,
	-2, 33,
	-1, 87,
	17, 49,
	18, 49,
	19, 49,
	20, 49,
	21, 49,
	22, 49,
	41, 49,
	42, 49,
	43, 49,
	48, 49,
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
	-2, 54,
	-1, 89,
	17, 49,
	18, 49,
	19, 49,
	20, 49,
	21, 49,
	22, 49,
	41, 49,
	42, 49,
	43, 49,
	48, 49,
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
	-2, 56,
	-1, 103,
	47, 40,
	50, 40,
	-2, 63,
	-1, 114,
	17, 0,
	18, 0,
	-2, 80,
	-1, 115,
	17, 0,
	18, 0,
	-2, 81,
	-1, 165,
	47, 41,
	50, 41,
	-2, 101,
}

const yyNprod = 103
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1646

var yyAct = []int{

	1, 183, 184, 34, 64, 32, 229, 217, 95, 131,
	214, 44, 131, 73, 74, 75, 76, 77, 78, 187,
	167, 80, 81, 39, 40, 41, 42, 43, 215, 79,
	155, 191, 126, 127, 92, 3, 102, 55, 57, 160,
	53, 154, 59, 61, 256, 83, 255, 69, 156, 72,
	82, 249, 247, 84, 85, 86, 88, 90, 246, 157,
	244, 243, 242, 59, 236, 97, 58, 99, 231, 135,
	230, 218, 139, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 59, 93, 222, 152, 205, 91, 203,
	185, 186, 201, 181, 178, 132, 44, 163, 252, 141,
	142, 143, 144, 145, 146, 147, 208, 251, 59, 149,
	41, 42, 43, 239, 124, 101, 233, 126, 127, 228,
	172, 227, 55, 57, 176, 53, 59, 164, 179, 73,
	74, 75, 76, 77, 78, 213, 129, 80, 81, 148,
	71, 223, 133, 138, 130, 79, 100, 131, 128, 193,
	212, 59, 198, 170, 59, 195, 59, 162, 175, 158,
	177, 83, 185, 186, 180, 250, 82, 207, 73, 74,
	75, 76, 77, 78, 209, 210, 80, 81, 248, 66,
	44, 216, 168, 192, 196, 171, 59, 173, 96, 194,
	190, 189, 199, 219, 174, 220, 14, 169, 153, 103,
	83, 126, 127, 225, 226, 82, 55, 57, 70, 53,
	98, 211, 65, 63, 234, 235, 94, 197, 237, 238,
	33, 182, 10, 240, 241, 2, 0, 0, 0, 0,
	245, 0, 221, 50, 52, 0, 0, 0, 0, 15,
	16, 22, 253, 254, 26, 6, 9, 7, 31, 0,
	11, 0, 0, 0, 0, 44, 45, 46, 30, 23,
	24, 25, 8, 12, 0, 49, 51, 39, 40, 41,
	42, 43, 4, 5, 0, 0, 126, 127, 0, 13,
	0, 55, 57, 0, 53, 0, 0, 0, 0, 17,
	21, 0, 0, 0, 35, 28, 0, 0, 27, 0,
	18, 19, 20, 29, 15, 16, 22, 0, 0, 26,
	6, 9, 7, 31, 0, 11, 0, 0, 0, 0,
	0, 0, 0, 30, 23, 24, 25, 8, 12, 0,
	0, 0, 0, 0, 0, 0, 0, 4, 5, 0,
	0, 0, 0, 0, 13, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 17, 21, 0, 0, 0, 0,
	28, 0, 0, 27, 0, 18, 19, 20, 29, 15,
	16, 136, 0, 0, 26, 6, 9, 7, 31, 0,
	11, 0, 0, 0, 0, 0, 0, 0, 30, 23,
	24, 25, 8, 12, 0, 0, 0, 0, 0, 0,
	0, 0, 4, 5, 0, 0, 0, 0, 0, 13,
	47, 48, 50, 52, 54, 56, 0, 0, 0, 17,
	21, 0, 0, 0, 0, 28, 0, 0, 27, 0,
	18, 19, 20, 29, 44, 45, 46, 0, 0, 0,
	0, 38, 0, 0, 49, 51, 39, 40, 41, 42,
	43, 0, 0, 232, 0, 126, 127, 0, 0, 134,
	55, 57, 0, 53, 47, 48, 50, 52, 54, 56,
	0, 0, 0, 0, 73, 74, 75, 76, 77, 78,
	0, 0, 80, 81, 0, 0, 0, 0, 44, 45,
	46, 0, 0, 0, 0, 38, 224, 0, 49, 51,
	39, 40, 41, 42, 43, 0, 83, 0, 0, 126,
	127, 82, 0, 0, 55, 57, 0, 53, 47, 48,
	50, 52, 54, 56, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 44, 45, 46, 0, 0, 0, 0, 38,
	0, 0, 49, 51, 39, 40, 41, 42, 43, 0,
	206, 0, 0, 126, 127, 0, 0, 0, 55, 57,
	0, 53, 47, 48, 50, 52, 54, 56, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 45, 46, 0,
	0, 0, 0, 38, 0, 0, 49, 51, 39, 40,
	41, 42, 43, 0, 0, 204, 0, 126, 127, 0,
	0, 0, 55, 57, 0, 53, 47, 48, 50, 52,
	54, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 45, 46, 0, 0, 0, 0, 38, 0, 0,
	49, 51, 39, 40, 41, 42, 43, 0, 0, 202,
	0, 126, 127, 0, 0, 0, 55, 57, 0, 53,
	47, 48, 50, 52, 54, 56, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 44, 45, 46, 0, 0, 0,
	0, 38, 0, 0, 49, 51, 39, 40, 41, 42,
	43, 0, 0, 0, 0, 126, 127, 200, 0, 0,
	55, 57, 0, 53, 47, 48, 50, 52, 54, 56,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 44, 45,
	46, 0, 0, 0, 0, 38, 0, 0, 49, 51,
	39, 40, 41, 42, 43, 0, 0, 0, 0, 126,
	127, 188, 0, 0, 55, 57, 0, 53, 47, 48,
	50, 52, 54, 56, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 44, 45, 46, 0, 0, 0, 0, 38,
	166, 0, 49, 51, 39, 40, 41, 42, 43, 0,
	0, 0, 0, 126, 127, 0, 0, 0, 55, 57,
	0, 53, 47, 48, 50, 52, 54, 56, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 45, 46, 0,
	0, 0, 0, 38, 0, 0, 49, 51, 39, 40,
	41, 42, 43, 0, 0, 0, 0, 126, 127, 165,
	0, 0, 55, 57, 0, 53, 47, 48, 50, 52,
	54, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 45, 46, 0, 0, 0, 0, 38, 0, 0,
	49, 51, 39, 40, 41, 42, 43, 0, 0, 161,
	0, 126, 127, 0, 0, 0, 55, 57, 0, 53,
	47, 48, 50, 52, 54, 56, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 44, 45, 46, 0, 0, 0,
	0, 38, 0, 0, 49, 51, 39, 40, 41, 42,
	43, 0, 0, 0, 0, 126, 127, 0, 0, 0,
	55, 57, 159, 53, 47, 48, 50, 52, 54, 56,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 44, 45,
	46, 0, 0, 0, 0, 38, 0, 0, 49, 51,
	39, 40, 41, 42, 43, 0, 0, 140, 0, 126,
	127, 0, 0, 0, 55, 57, 0, 53, 47, 48,
	50, 52, 54, 56, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 44, 45, 46, 0, 0, 0, 0, 38,
	0, 0, 49, 51, 39, 40, 41, 42, 43, 0,
	0, 137, 0, 36, 37, 0, 0, 0, 55, 57,
	0, 53, 47, 48, 50, 52, 54, 56, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 45, 46, 0,
	0, 0, 0, 38, 0, 125, 49, 51, 39, 40,
	41, 42, 43, 0, 0, 0, 0, 126, 127, 0,
	0, 0, 55, 57, 0, 53, 47, 48, 50, 52,
	54, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	44, 45, 46, 0, 0, 0, 0, 38, 0, 0,
	49, 51, 39, 40, 41, 42, 43, 0, 0, 0,
	0, 126, 127, 0, 0, 0, 55, 57, 0, 53,
	47, 48, 50, 52, 54, 56, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 44, 45, 46, 0, 0, 0,
	0, 38, 0, 0, 49, 51, 39, 40, 41, 42,
	43, 0, 0, 0, 0, 36, 37, 0, 0, 0,
	55, 57, 0, 53, 47, 48, 50, 52, 54, 56,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 44, 45,
	46, 0, 0, 0, 0, 38, 0, 0, 49, 51,
	39, 40, 41, 42, 43, 0, 0, 0, 0, 151,
	127, 0, 0, 0, 55, 57, 0, 53, 47, 48,
	50, 52, 54, 56, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 44, 45, 46, 0, 0, 0, 0, 38,
	0, 0, 49, 51, 39, 40, 41, 42, 43, 0,
	0, 0, 0, 150, 127, 0, 0, 0, 55, 57,
	0, 53, 47, 48, 50, 52, 0, 56, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 44, 45, 46, 0,
	0, 0, 47, 48, 50, 52, 49, 51, 39, 40,
	41, 42, 43, 0, 0, 0, 0, 126, 127, 0,
	0, 0, 55, 57, 0, 53, 44, 45, 46, 0,
	0, 0, 0, 0, 0, 0, 49, 51, 39, 40,
	41, 42, 43, 0, 62, 16, 22, 126, 127, 26,
	0, 0, 55, 57, 0, 53, 0, 0, 0, 0,
	0, 0, 0, 30, 23, 24, 25, 60, 16, 22,
	0, 0, 26, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 30, 23, 24, 25,
	0, 15, 16, 22, 17, 21, 26, 0, 0, 0,
	28, 0, 0, 27, 0, 18, 19, 20, 29, 0,
	30, 23, 24, 25, 89, 16, 22, 17, 21, 26,
	0, 0, 0, 28, 0, 0, 27, 0, 18, 19,
	20, 29, 0, 30, 23, 24, 25, 0, 87, 16,
	22, 17, 21, 26, 0, 0, 0, 28, 0, 0,
	27, 0, 18, 19, 20, 29, 0, 30, 23, 24,
	25, 67, 16, 22, 17, 21, 26, 0, 0, 0,
	28, 0, 0, 27, 0, 18, 19, 20, 29, 0,
	30, 23, 24, 25, 0, 0, 0, 0, 17, 21,
	0, 0, 0, 0, 28, 0, 0, 27, 0, 18,
	19, 20, 29, 0, 0, 0, 0, 0, 0, 0,
	0, 17, 21, 0, 0, 0, 0, 68, 0, 0,
	27, 0, 18, 19, 20, 29,
}
var yyPact = []int{

	310, -1000, 245, 1213, -1000, -1000, 1483, 1460, 219, 218,
	175, 1577, 90, 1460, -1000, 147, -1000, 1460, 1460, 1460,
	1554, 1530, -1000, -1000, -1000, -1000, 30, 1483, 192, 1460,
	216, 1460, 109, 75, -1000, 310, 205, 1460, 1460, 1460,
	1460, 1460, 1460, 1460, 1460, 1460, 1460, 1460, 1460, 1460,
	1460, 1460, 1460, 1460, 1460, 1460, 1460, 1483, -1000, 1105,
	108, 1159, -18, 86, 107, -1000, 92, 453, 375, 1051,
	94, 310, 997, 1460, 1460, 1460, 1460, 1460, 1460, 1460,
	-1000, -1000, 1483, 1460, 149, 149, 149, -18, 1321, -18,
	1267, 204, -27, -34, -2, -1000, 120, 943, -29, 889,
	1483, 1507, -1000, -1000, 835, 781, 65, 65, 149, 149,
	149, 1159, -30, -30, 224, 224, -30, -30, -30, -30,
	1159, 1375, 1159, 1405, -49, 1483, 203, 1460, 1483, 310,
	1483, 200, 1460, 310, 1460, 43, 120, 310, 1460, 42,
	127, 1159, 1159, 1159, 1159, 1159, 1159, 1159, -50, 727,
	197, 196, -38, 185, 195, -1000, 192, -1000, 1460, -1000,
	1483, 310, -1000, -1000, 1213, -1000, 1460, -1000, -1000, -1000,
	673, -1000, 41, -1000, -1000, 619, 38, 565, -1000, 36,
	511, 148, 55, -1000, -1000, 1460, 111, -1000, -1000, -1000,
	-1000, 85, -59, -41, 183, -1000, 1159, -62, 10, 1159,
	-1000, -1000, 310, -1000, 310, -1000, 1460, 91, -1000, -1000,
	-1000, 457, 310, 310, 71, 69, -63, -1000, -1000, 9,
	7, 403, 66, 310, 310, -1000, 3, 310, 310, 63,
	-1000, -1000, 310, 310, 1, -1000, -1000, 0, -1, 310,
	-3, -9, 158, -1000, -1000, -10, -1000, 145, 57, -1000,
	48, 310, 310, -15, -17, -1000, -1000,
}
var yyPgo = []int{

	0, 0, 235, 232, 2, 1, 231, 35, 66, 230,
	5, 206, 8, 226, 4,
}
var yyR1 = []int{

	0, 1, 1, 1, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 3, 3, 3, 6, 6, 6, 6, 6,
	5, 4, 12, 13, 13, 13, 14, 14, 14, 9,
	9, 9, 11, 10, 10, 8, 8, 8, 8, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 1, 1, 1, 2, 2, 5,
	4, 1, 7, 4, 5, 9, 13, 12, 9, 8,
	5, 1, 7, 5, 5, 0, 1, 2, 1, 2,
	4, 3, 3, 0, 1, 3, 0, 1, 3, 1,
	3, 4, 3, 1, 3, 0, 1, 3, 3, 1,
	1, 2, 2, 2, 2, 4, 2, 4, 1, 1,
	1, 1, 5, 3, 7, 8, 8, 9, 3, 3,
	3, 5, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 3, 3, 3, 3, 4,
	4, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, -7, 37, 38, 10, 12, 27, 11,
	-3, 15, 28, 44, -11, 4, 5, 54, 65, 66,
	67, 55, 6, 24, 25, 26, 9, 63, 60, 68,
	23, 13, -10, -9, -1, 59, 62, 63, 48, 53,
	54, 55, 56, 57, 41, 42, 43, 17, 18, 51,
	19, 52, 20, 70, 21, 67, 22, 68, -8, -7,
	4, -7, 4, 4, -14, 4, 14, 4, 60, -7,
	-11, 60, -7, 31, 32, 33, 34, 35, 36, 47,
	39, 40, 68, 63, -7, -7, -7, 4, -7, 4,
	-7, 68, 4, -8, -13, -12, 6, -7, 4, -7,
	47, 50, -1, 4, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -8, 50, 62, 63, 50, 60,
	47, 50, 13, 60, 16, -1, 6, 60, 59, -1,
	60, -7, -7, -7, -7, -7, -7, -7, -8, -7,
	62, 62, -14, 4, 68, 64, 50, 61, 49, 69,
	68, 60, -8, -10, -7, 64, 49, 69, -8, 4,
	-7, -8, -1, -8, 4, -7, -1, -7, 61, -1,
	-7, 61, -6, -5, -4, 45, 46, 69, 64, 4,
	4, 69, 8, -14, 4, -12, -7, -8, -1, -7,
	64, 61, 60, 61, 60, 61, 59, 29, 61, -5,
	-4, -7, 49, 60, 69, 69, 8, 69, 61, -1,
	-1, -7, 4, 60, 49, -1, -1, 60, 60, 69,
	61, 61, 60, 60, -1, -1, 61, -1, -1, 60,
	-1, -1, 61, 61, 61, -1, 61, 61, 30, 61,
	30, 60, 60, -1, -1, 61, 61,
}
var yyDef = []int{

	1, -2, 1, 4, 5, 6, 45, 0, 0, 36,
	11, 0, 0, 0, 21, -2, 50, 0, 0, 0,
	0, 0, 58, 59, 60, 61, 0, 45, 33, 0,
	0, 0, 0, 43, 2, 1, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 45, 7, 46,
	49, 8, 49, 0, 0, 37, 0, -2, -2, 0,
	0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
	93, 94, 45, 0, 51, 52, 53, -2, 0, -2,
	0, 36, 0, 0, 0, 34, 0, 0, 0, 0,
	45, 0, 3, -2, 0, 0, 72, 73, 74, 75,
	76, 77, 78, 79, -2, -2, 82, 83, 84, 85,
	95, 96, 97, 98, 0, 45, 0, 0, 45, 1,
	45, 0, 0, 1, 0, 0, 58, 1, 0, 0,
	25, 86, 87, 88, 89, 90, 91, 92, 0, 0,
	0, 0, 0, 37, 36, 68, 0, 69, 0, 70,
	45, 1, 42, 44, 0, -2, 0, 102, 47, 63,
	0, 48, 0, 10, 38, 0, 0, 0, 13, 0,
	0, 0, 0, 26, 28, 0, 0, 99, 100, 55,
	57, 0, 0, 0, 37, 35, 32, 0, 0, 62,
	101, 9, 1, 23, 1, 14, 0, 0, 20, 27,
	29, 0, 1, 1, 0, 0, 0, 71, 24, 0,
	0, 0, 0, 1, 1, 31, 0, 1, 1, 0,
	22, 12, 1, 1, 0, 30, 64, 0, 0, 1,
	0, 0, 19, 65, 66, 0, 15, 18, 0, 67,
	0, 1, 1, 0, 0, 17, 16,
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
				for _, s := range yyS[yypt-0].stmts {
					if yyS[yypt-1].stmt.Position().Line == s.Position().Line {
						l.pos = yyS[yypt-1].stmt.Position()
						yylex.Error("syntax error")
					}
				}
			}
		}
	case 3:
		//line parser.go.y:80
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-2].stmt.SetPosition(l.pos)
			}
		}
	case 4:
		//line parser.go.y:89
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			yyVAL.stmt.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 5:
		//line parser.go.y:94
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 6:
		//line parser.go.y:99
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 7:
		//line parser.go.y:104
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
			yyVAL.stmt.SetPosition(yyS[yypt-1].tok.Position())
		}
	case 8:
		//line parser.go.y:109
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			yyVAL.stmt.SetPosition(yyS[yypt-1].tok.Position())
		}
	case 9:
		//line parser.go.y:114
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.Lit, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-4].tok.Position())
		}
	case 10:
		//line parser.go.y:119
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyS[yypt-2].expr_idents, Exprs: yyS[yypt-0].exprs}
			yyVAL.stmt.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 11:
		//line parser.go.y:124
		{
			yyVAL.stmt = yyS[yypt-0].stmt_if
			yyVAL.stmt.SetPosition(yyS[yypt-0].stmt_if.Position())
		}
	case 12:
		//line parser.go.y:129
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyS[yypt-5].tok.Lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-6].tok.Position())
		}
	case 13:
		//line parser.go.y:134
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 14:
		//line parser.go.y:139
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-4].tok.Position())
		}
	case 15:
		//line parser.go.y:144
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyS[yypt-7].expr_lets, Expr2: yyS[yypt-5].expr, Expr3: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-8].tok.Position())
		}
	case 16:
		//line parser.go.y:149
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-10].stmts, Var: yyS[yypt-7].tok.Lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-12].tok.Position())
		}
	case 17:
		//line parser.go.y:154
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-11].tok.Position())
		}
	case 18:
		//line parser.go.y:159
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-6].stmts, Var: yyS[yypt-3].tok.Lit, Catch: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-8].tok.Position())
		}
	case 19:
		//line parser.go.y:164
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
			yyVAL.stmt.SetPosition(yyS[yypt-7].tok.Position())
		}
	case 20:
		//line parser.go.y:169
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyS[yypt-3].expr, Cases: yyS[yypt-1].stmt_cases}
			yyVAL.stmt.SetPosition(yyS[yypt-4].tok.Position())
		}
	case 21:
		//line parser.go.y:174
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr_lets}
			yyVAL.stmt.SetPosition(yyS[yypt-0].expr_lets.Position())
		}
	case 22:
		//line parser.go.y:180
		{
			yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts})
		}
	case 23:
		//line parser.go.y:184
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyS[yypt-1].stmts...)
			}
		}
	case 24:
		//line parser.go.y:192
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts, Else: nil}
		}
	case 25:
		//line parser.go.y:197
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 26:
		//line parser.go.y:201
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_case}
		}
	case 27:
		//line parser.go.y:205
		{
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_case)
		}
	case 28:
		//line parser.go.y:209
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_default}
		}
	case 29:
		//line parser.go.y:213
		{
			for _, stmt := range yyS[yypt-1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_default)
		}
	case 30:
		//line parser.go.y:223
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyS[yypt-2].expr, Stmts: yyS[yypt-0].stmts}
		}
	case 31:
		//line parser.go.y:228
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyS[yypt-0].stmts}
		}
	case 32:
		//line parser.go.y:233
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyS[yypt-2].tok.Lit, Value: yyS[yypt-0].expr}
		}
	case 33:
		//line parser.go.y:238
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 34:
		//line parser.go.y:242
		{
			yyVAL.expr_pairs = []ast.Expr{yyS[yypt-0].expr_pair}
		}
	case 35:
		//line parser.go.y:246
		{
			yyVAL.expr_pairs = append(yyS[yypt-2].expr_pairs, yyS[yypt-0].expr_pair)
		}
	case 36:
		//line parser.go.y:251
		{
			yyVAL.expr_idents = []string{}
		}
	case 37:
		//line parser.go.y:255
		{
			yyVAL.expr_idents = []string{yyS[yypt-0].tok.Lit}
		}
	case 38:
		//line parser.go.y:259
		{
			yyVAL.expr_idents = append(yyS[yypt-2].expr_idents, yyS[yypt-0].tok.Lit)
		}
	case 39:
		//line parser.go.y:264
		{
			yyVAL.expr_lhs = &ast.IdentExpr{Lit: yyS[yypt-0].tok.Lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPosition(l.pos)
			}
		}
	case 40:
		//line parser.go.y:269
		{
			yyVAL.expr_lhs = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.Lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPosition(l.pos)
			}
		}
	case 41:
		//line parser.go.y:274
		{
			yyVAL.expr_lhs = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPosition(l.pos)
			}
		}
	case 42:
		//line parser.go.y:280
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyS[yypt-2].expr_lhss, Operator: "=", Rhss: yyS[yypt-0].exprs}
			yyVAL.expr_lets.SetPosition(yyS[yypt-2].expr_lhss[0].Position())
		}
	case 43:
		//line parser.go.y:292
		{
			yyVAL.expr_lhss = []ast.Expr{yyS[yypt-0].expr_lhs}
		}
	case 44:
		//line parser.go.y:296
		{
			yyVAL.expr_lhss = append([]ast.Expr{yyS[yypt-2].expr_lhs}, yyS[yypt-0].expr_lhss...)
		}
	case 45:
		//line parser.go.y:301
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 46:
		//line parser.go.y:305
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 47:
		//line parser.go.y:309
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 48:
		//line parser.go.y:313
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.Lit}}, yyS[yypt-0].exprs...)
		}
	case 49:
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 50:
		//line parser.go.y:324
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 51:
		//line parser.go.y:329
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 52:
		//line parser.go.y:334
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 53:
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-0].expr.Position())
		}
	case 54:
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.Lit}}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 55:
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.Lit}}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 56:
		//line parser.go.y:354
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.Lit}}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 57:
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.Lit}}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 58:
		//line parser.go.y:364
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 59:
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 60:
		//line parser.go.y:374
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 61:
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-0].tok.Position())
		}
	case 62:
		//line parser.go.y:384
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-4].expr.Position())
		}
	case 63:
		//line parser.go.y:389
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.Lit}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 64:
		//line parser.go.y:394
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			yyVAL.expr.SetPosition(yyS[yypt-6].tok.Position())
		}
	case 65:
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.Lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			yyVAL.expr.SetPosition(yyS[yypt-7].tok.Position())
		}
	case 66:
		//line parser.go.y:404
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.Lit, Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			yyVAL.expr.SetPosition(yyS[yypt-7].tok.Position())
		}
	case 67:
		//line parser.go.y:409
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.Lit, Args: []string{yyS[yypt-5].tok.Lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			yyVAL.expr.SetPosition(yyS[yypt-8].tok.Position())
		}
	case 68:
		//line parser.go.y:414
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 69:
		//line parser.go.y:419
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
	case 70:
		//line parser.go.y:428
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 71:
		//line parser.go.y:433
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.Lit, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPosition(yyS[yypt-4].tok.Position())
		}
	case 72:
		//line parser.go.y:438
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 73:
		//line parser.go.y:443
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 74:
		//line parser.go.y:448
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 75:
		//line parser.go.y:453
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 76:
		//line parser.go.y:458
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 77:
		//line parser.go.y:463
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 78:
		//line parser.go.y:468
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 79:
		//line parser.go.y:473
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 80:
		//line parser.go.y:478
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 81:
		//line parser.go.y:483
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 82:
		//line parser.go.y:488
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 83:
		//line parser.go.y:493
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 84:
		//line parser.go.y:498
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 85:
		//line parser.go.y:503
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 86:
		//line parser.go.y:508
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.Lit, Operator: "+=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].tok.Position())
		}
	case 87:
		//line parser.go.y:513
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.Lit, Operator: "-=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].tok.Position())
		}
	case 88:
		//line parser.go.y:518
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.Lit, Operator: "*=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].tok.Position())
		}
	case 89:
		//line parser.go.y:523
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.Lit, Operator: "/=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].tok.Position())
		}
	case 90:
		//line parser.go.y:528
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.Lit, Operator: "&=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].tok.Position())
		}
	case 91:
		//line parser.go.y:533
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.Lit, Operator: "|=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].tok.Position())
		}
	case 92:
		//line parser.go.y:538
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.Lit, Operator: "=", Expr: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].tok.Position())
		}
	case 93:
		//line parser.go.y:543
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.Lit, Operator: "++"}
			yyVAL.expr.SetPosition(yyS[yypt-1].tok.Position())
		}
	case 94:
		//line parser.go.y:548
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.Lit, Operator: "--"}
			yyVAL.expr.SetPosition(yyS[yypt-1].tok.Position())
		}
	case 95:
		//line parser.go.y:553
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 96:
		//line parser.go.y:558
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 97:
		//line parser.go.y:563
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 98:
		//line parser.go.y:568
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			yyVAL.expr.SetPosition(yyS[yypt-2].expr.Position())
		}
	case 99:
		//line parser.go.y:573
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.Lit, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 100:
		//line parser.go.y:578
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyS[yypt-3].tok.Lit}, Index: yyS[yypt-1].expr}
			yyVAL.expr.SetPosition(yyS[yypt-3].tok.Position())
		}
	case 101:
		//line parser.go.y:583
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			yyVAL.expr.SetPosition(yyS[yypt-3].expr.Position())
		}
	case 102:
		//line parser.go.y:588
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			yyVAL.expr.SetPosition(yyS[yypt-3].expr.Position())
		}
	}
	goto yystack /* stack new state and value */
}
