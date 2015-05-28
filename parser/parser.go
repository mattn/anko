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
	compstmt     []ast.Stmt
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
	term         ast.Token
	terms        ast.Token
	opt_terms    ast.Token
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
	"'{'",
	"'}'",
	"';'",
	"'!'",
	"'^'",
	"'&'",
	"'.'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'|'",
	"'\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:620

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	50, 45,
	-2, 2,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 5,
	50, 46,
	-2, 23,
	-1, 35,
	50, 45,
	-2, 105,
	-1, 79,
	60, 2,
	-2, 45,
	-1, 80,
	50, 46,
	-2, 42,
	-1, 83,
	50, 45,
	-2, 2,
	-1, 90,
	1, 54,
	45, 54,
	46, 54,
	47, 54,
	49, 54,
	50, 54,
	59, 54,
	60, 54,
	61, 54,
	67, 54,
	69, 54,
	71, 54,
	-2, 49,
	-1, 92,
	1, 56,
	45, 56,
	46, 56,
	47, 56,
	49, 56,
	50, 56,
	59, 56,
	60, 56,
	61, 56,
	67, 56,
	69, 56,
	71, 56,
	-2, 49,
	-1, 118,
	17, 0,
	18, 0,
	-2, 82,
	-1, 119,
	17, 0,
	18, 0,
	-2, 83,
	-1, 137,
	50, 46,
	-2, 42,
	-1, 139,
	50, 45,
	-2, 2,
	-1, 141,
	50, 45,
	-2, 2,
	-1, 144,
	50, 45,
	-2, 2,
	-1, 163,
	50, 45,
	-2, 2,
	-1, 195,
	50, 47,
	-2, 43,
	-1, 196,
	1, 44,
	45, 44,
	46, 44,
	47, 44,
	50, 48,
	60, 44,
	61, 44,
	71, 44,
	-2, 49,
	-1, 201,
	1, 48,
	45, 48,
	46, 48,
	50, 48,
	60, 48,
	61, 48,
	67, 48,
	69, 48,
	71, 48,
	-2, 49,
	-1, 203,
	50, 45,
	-2, 2,
	-1, 205,
	50, 45,
	-2, 2,
	-1, 216,
	50, 45,
	-2, 2,
	-1, 231,
	50, 45,
	-2, 2,
	-1, 235,
	50, 45,
	-2, 2,
	-1, 236,
	50, 45,
	-2, 2,
	-1, 243,
	50, 45,
	-2, 2,
	-1, 244,
	50, 45,
	-2, 2,
	-1, 247,
	50, 45,
	-2, 2,
	-1, 251,
	50, 45,
	-2, 2,
	-1, 255,
	50, 45,
	-2, 2,
	-1, 266,
	50, 45,
	-2, 2,
	-1, 267,
	50, 45,
	-2, 2,
}

const yyNprod = 110
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1497

var yyAct = [...]int{

	1, 210, 158, 5, 98, 211, 32, 34, 37, 191,
	189, 221, 138, 74, 75, 39, 73, 80, 38, 84,
	37, 37, 87, 88, 89, 91, 93, 142, 6, 225,
	38, 38, 36, 96, 99, 237, 85, 102, 86, 217,
	161, 155, 82, 145, 271, 107, 108, 270, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 126, 127, 128, 129, 104, 106,
	130, 131, 132, 133, 134, 74, 137, 85, 135, 86,
	143, 159, 264, 261, 147, 106, 218, 260, 138, 74,
	150, 138, 149, 95, 58, 59, 60, 61, 62, 63,
	74, 136, 185, 156, 49, 181, 162, 258, 168, 137,
	153, 165, 257, 254, 248, 242, 241, 230, 46, 47,
	48, 212, 213, 140, 226, 223, 206, 68, 43, 71,
	104, 70, 204, 66, 164, 240, 209, 202, 178, 174,
	170, 267, 172, 169, 171, 175, 173, 266, 146, 176,
	137, 251, 244, 180, 236, 94, 235, 216, 139, 83,
	233, 190, 192, 193, 197, 74, 195, 138, 194, 141,
	199, 187, 231, 200, 105, 177, 101, 106, 212, 213,
	265, 72, 214, 3, 262, 208, 215, 77, 219, 186,
	159, 198, 188, 184, 220, 183, 222, 224, 154, 109,
	100, 76, 40, 157, 227, 97, 228, 81, 35, 179,
	12, 229, 2, 0, 0, 0, 232, 234, 0, 103,
	0, 0, 0, 0, 200, 239, 0, 0, 0, 0,
	0, 0, 245, 0, 0, 0, 249, 250, 247, 0,
	0, 0, 0, 0, 252, 253, 0, 0, 256, 0,
	0, 255, 259, 0, 0, 0, 263, 52, 53, 55,
	57, 67, 69, 0, 0, 0, 0, 268, 269, 0,
	0, 58, 59, 60, 61, 62, 63, 0, 0, 64,
	65, 49, 50, 51, 0, 0, 0, 0, 42, 246,
	0, 54, 56, 44, 45, 46, 47, 48, 52, 53,
	55, 57, 67, 69, 68, 43, 71, 0, 70, 0,
	66, 0, 58, 59, 60, 61, 62, 63, 0, 0,
	64, 65, 49, 50, 51, 0, 0, 0, 0, 42,
	0, 0, 54, 56, 44, 45, 46, 47, 48, 0,
	243, 0, 0, 0, 0, 68, 43, 71, 0, 70,
	0, 66, 52, 53, 55, 57, 67, 69, 0, 0,
	0, 0, 0, 0, 0, 0, 58, 59, 60, 61,
	62, 63, 0, 0, 64, 65, 49, 50, 51, 0,
	0, 0, 0, 42, 0, 0, 54, 56, 44, 45,
	46, 47, 48, 0, 0, 0, 207, 0, 0, 68,
	43, 71, 0, 70, 0, 66, 52, 53, 55, 57,
	67, 69, 0, 0, 0, 0, 0, 0, 0, 0,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 42, 0, 0,
	54, 56, 44, 45, 46, 47, 48, 0, 205, 0,
	0, 0, 0, 68, 43, 71, 0, 70, 0, 66,
	52, 53, 55, 57, 67, 69, 0, 0, 0, 0,
	0, 0, 0, 0, 58, 59, 60, 61, 62, 63,
	0, 0, 64, 65, 49, 50, 51, 0, 0, 0,
	0, 42, 0, 0, 54, 56, 44, 45, 46, 47,
	48, 0, 203, 0, 0, 0, 0, 68, 43, 71,
	0, 70, 0, 66, 52, 53, 55, 57, 67, 69,
	0, 0, 0, 0, 0, 0, 0, 0, 58, 59,
	60, 61, 62, 63, 0, 0, 64, 65, 49, 50,
	51, 0, 0, 0, 0, 42, 0, 0, 54, 56,
	44, 45, 46, 47, 48, 52, 53, 55, 57, 67,
	69, 68, 43, 71, 0, 70, 182, 66, 0, 58,
	59, 60, 61, 62, 63, 0, 0, 64, 65, 49,
	50, 51, 0, 0, 0, 0, 42, 0, 0, 54,
	56, 44, 45, 46, 47, 48, 52, 53, 55, 57,
	67, 69, 68, 43, 71, 0, 70, 167, 66, 0,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 42, 166, 0,
	54, 56, 44, 45, 46, 47, 48, 52, 53, 55,
	57, 67, 69, 68, 43, 71, 0, 70, 0, 66,
	0, 58, 59, 60, 61, 62, 63, 0, 0, 64,
	65, 49, 50, 51, 0, 0, 0, 0, 42, 0,
	0, 54, 56, 44, 45, 46, 47, 48, 0, 163,
	0, 0, 0, 0, 68, 43, 71, 0, 70, 0,
	66, 52, 53, 55, 57, 67, 69, 0, 0, 0,
	0, 0, 0, 0, 0, 58, 59, 60, 61, 62,
	63, 0, 0, 64, 65, 49, 50, 51, 0, 0,
	0, 0, 42, 0, 0, 54, 56, 44, 45, 46,
	47, 48, 52, 53, 55, 57, 67, 69, 68, 43,
	71, 160, 70, 0, 66, 0, 58, 59, 60, 61,
	62, 63, 0, 0, 64, 65, 49, 50, 51, 0,
	0, 0, 0, 42, 0, 0, 54, 56, 44, 45,
	46, 47, 48, 0, 148, 0, 0, 0, 0, 68,
	43, 71, 0, 70, 0, 66, 52, 53, 55, 57,
	67, 69, 0, 0, 0, 0, 0, 0, 0, 0,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 42, 0, 0,
	54, 56, 44, 45, 46, 47, 48, 0, 144, 0,
	0, 0, 0, 68, 43, 71, 0, 70, 0, 66,
	52, 53, 55, 57, 67, 69, 0, 0, 0, 0,
	0, 0, 0, 0, 58, 59, 60, 61, 62, 63,
	0, 0, 64, 65, 49, 50, 51, 0, 0, 0,
	41, 42, 0, 0, 54, 56, 44, 45, 46, 47,
	48, 52, 53, 55, 57, 67, 69, 68, 43, 71,
	0, 70, 0, 66, 0, 58, 59, 60, 61, 62,
	63, 0, 0, 64, 65, 49, 50, 51, 0, 0,
	0, 0, 42, 0, 0, 54, 56, 44, 45, 46,
	47, 48, 52, 53, 55, 57, 67, 69, 68, 43,
	71, 0, 70, 0, 66, 0, 58, 59, 60, 61,
	62, 63, 0, 0, 64, 65, 49, 50, 51, 0,
	0, 0, 0, 42, 0, 0, 54, 56, 44, 45,
	46, 47, 48, 52, 53, 55, 57, 67, 69, 68,
	152, 71, 0, 70, 0, 66, 0, 58, 59, 60,
	61, 62, 63, 0, 0, 64, 65, 49, 50, 51,
	0, 0, 0, 0, 42, 0, 0, 54, 56, 44,
	45, 46, 47, 48, 0, 0, 0, 0, 0, 0,
	68, 151, 71, 0, 70, 0, 66, 16, 17, 23,
	0, 0, 27, 9, 4, 10, 33, 0, 13, 0,
	0, 0, 0, 0, 0, 0, 31, 24, 25, 26,
	11, 14, 0, 0, 0, 0, 0, 0, 0, 0,
	7, 8, 0, 0, 0, 0, 0, 15, 0, 58,
	59, 60, 61, 62, 63, 0, 0, 18, 22, 49,
	0, 0, 29, 0, 37, 19, 20, 21, 0, 30,
	0, 28, 0, 0, 38, 52, 53, 55, 57, 0,
	69, 0, 68, 43, 71, 0, 70, 0, 66, 58,
	59, 60, 61, 62, 63, 0, 0, 64, 65, 49,
	50, 51, 0, 0, 0, 0, 0, 0, 0, 54,
	56, 44, 45, 46, 47, 48, 52, 53, 55, 57,
	0, 0, 68, 43, 71, 0, 70, 0, 66, 0,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 0, 0, 0,
	54, 56, 44, 45, 46, 47, 48, 0, 55, 57,
	0, 0, 0, 68, 43, 71, 0, 70, 0, 66,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 0, 0, 0,
	54, 56, 44, 45, 46, 47, 48, 0, 0, 0,
	0, 0, 0, 68, 43, 71, 0, 70, 0, 66,
	16, 17, 23, 0, 0, 27, 9, 4, 10, 33,
	0, 13, 0, 0, 0, 0, 0, 0, 0, 31,
	24, 25, 26, 11, 14, 0, 0, 0, 0, 0,
	0, 0, 0, 7, 8, 0, 0, 0, 0, 0,
	15, 0, 0, 0, 58, 59, 60, 61, 62, 63,
	18, 22, 64, 65, 49, 29, 0, 0, 19, 20,
	21, 0, 30, 0, 28, 0, 44, 45, 46, 47,
	48, 201, 17, 23, 0, 0, 27, 68, 43, 71,
	0, 70, 0, 66, 0, 0, 0, 0, 0, 0,
	31, 24, 25, 26, 16, 17, 23, 0, 0, 27,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 31, 24, 25, 26, 0, 201, 17,
	23, 18, 22, 27, 0, 0, 29, 0, 0, 19,
	20, 21, 0, 30, 0, 28, 238, 31, 24, 25,
	26, 196, 17, 23, 18, 22, 27, 0, 0, 29,
	0, 0, 19, 20, 21, 0, 30, 0, 28, 0,
	31, 24, 25, 26, 0, 92, 17, 23, 18, 22,
	27, 0, 0, 29, 0, 0, 19, 20, 21, 0,
	30, 0, 28, 0, 31, 24, 25, 26, 90, 17,
	23, 18, 22, 27, 0, 0, 29, 0, 0, 19,
	20, 21, 0, 30, 0, 28, 0, 31, 24, 25,
	26, 0, 78, 17, 23, 18, 22, 27, 0, 0,
	29, 0, 0, 19, 20, 21, 0, 30, 0, 28,
	0, 31, 24, 25, 26, 0, 0, 0, 18, 22,
	0, 0, 0, 29, 0, 0, 19, 20, 21, 0,
	30, 0, 28, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 18, 22, 0, 0, 0, 79, 0, 0,
	19, 20, 21, 0, 30, 0, 28,
}
var yyPact = [...]int{

	1216, -1000, -53, -1000, 198, 823, 134, -1000, -1000, 1310,
	1310, 197, 173, 1428, 100, 1310, -30, -1000, 1310, 1310,
	1310, 1404, 1381, -1000, -1000, -1000, -1000, 89, -53, -53,
	1310, 196, 126, 1310, -1000, 1013, -1000, -1000, -1000, 127,
	-1000, 1310, 1310, 195, 1310, 1310, 1310, 1310, 1310, 1310,
	1310, 1310, 1310, 1310, 1310, 1310, 1310, 1310, 1310, 1310,
	1310, 1310, 1310, 1310, -1000, -1000, 1310, 1310, 1310, 1310,
	1310, 1310, 1310, 117, 864, 864, 99, 110, 11, 1013,
	769, -18, 101, 1216, 715, 1310, 1310, 1028, 1028, 1028,
	-30, 946, -30, 905, 194, -25, 1310, -53, 184, 674,
	-26, -53, 620, -1000, -1000, 1310, -53, 864, 579, -1000,
	63, 63, 1028, 1028, 1028, 864, 1233, 1233, 1149, 1149,
	1233, 1233, 1233, 1233, 864, 864, 864, 864, 864, 864,
	864, 1068, 864, 1109, 538, 41, -1000, 864, -53, 1216,
	1310, 1216, 1310, 79, 1216, 1310, 1310, 78, -53, 38,
	497, 191, 189, 35, 181, 188, -40, -41, -1000, 114,
	-1000, 1310, 1357, 1216, -1000, 187, 1310, -1000, -1000, 1334,
	77, 443, 72, 389, -1000, 66, 335, -1000, 156, 76,
	133, -1000, -1000, -1000, -1000, 98, -28, 19, 180, -53,
	-58, -53, 65, 1310, -38, 864, -30, 64, -1000, 864,
	864, -30, -1000, 1216, -1000, 1216, -1000, 1310, 113, -1000,
	-1000, -1000, 1310, 111, -1000, -1000, 1216, 97, 95, -32,
	1287, -1000, 75, -1000, 864, -1000, -1000, 56, 55, 281,
	93, 1216, 240, -53, 54, 1216, 1216, 92, -1000, -1000,
	-1000, -1000, -1000, 1216, 1216, 53, -53, 1216, -1000, 52,
	47, 1216, 27, 23, 154, 1216, -1000, -1000, -1000, 22,
	-1000, 150, 88, -1000, -1000, 82, 1216, 1216, -13, -16,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 0, 212, 183, 210, 5, 1, 209, 3, 6,
	28, 207, 2, 203, 15, 4, 205, 32,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 4, 4, 7, 7, 7,
	7, 7, 6, 5, 12, 13, 13, 13, 14, 14,
	14, 11, 10, 10, 10, 9, 9, 9, 9, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 15, 15, 16, 16, 17, 17,
}
var yyR2 = [...]int{

	0, 2, 0, 1, 3, 4, 3, 3, 1, 1,
	2, 2, 5, 1, 7, 4, 5, 9, 13, 12,
	9, 8, 5, 1, 7, 5, 5, 0, 2, 2,
	2, 2, 5, 4, 3, 0, 1, 4, 0, 1,
	4, 3, 1, 4, 4, 0, 1, 4, 4, 1,
	1, 2, 2, 2, 2, 4, 2, 4, 1, 1,
	1, 1, 5, 3, 7, 8, 8, 9, 5, 6,
	5, 6, 3, 5, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 3, 3, 3, 3,
	4, 4, 4, 4, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 11, -8, -10, 37, 38, 10,
	12, 27, -4, 15, 28, 44, 4, 5, 54, 62,
	63, 64, 55, 6, 24, 25, 26, 9, 68, 59,
	66, 23, -9, 13, -15, -16, -17, 61, 71, -14,
	4, 47, 48, 65, 53, 54, 55, 56, 57, 41,
	42, 43, 17, 18, 51, 19, 52, 20, 31, 32,
	33, 34, 35, 36, 39, 40, 70, 21, 64, 22,
	68, 66, 47, -9, -8, -8, 4, 14, 4, 59,
	-8, -11, -10, 59, -8, 66, 68, -8, -8, -8,
	4, -8, 4, -8, 66, 4, -15, -16, -15, -8,
	4, 50, -8, -3, -17, 47, 50, -8, -8, 4,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -9, -10, -8, 50, 59,
	13, 59, 16, -1, 59, 61, 47, -1, 59, -9,
	-8, 65, 65, -14, 4, 66, -9, -13, -12, 6,
	67, 66, -15, 59, -10, -15, 49, 69, 67, -15,
	-1, -8, -1, -8, 60, -1, -8, -10, 60, -7,
	-15, 67, 69, 4, 4, 67, 8, -14, 4, 50,
	-15, 50, -15, 49, -9, -8, 4, -1, 4, -8,
	-8, 4, 60, 59, 60, 59, 60, 61, 29, 60,
	-6, -5, 45, 46, -6, -5, 59, 67, 67, 8,
	-15, 69, -15, 60, -8, 67, 60, -1, -1, -8,
	4, 59, -8, 49, -1, 59, 59, 67, 69, -12,
	60, 60, 60, 59, 59, -1, 49, -15, 60, -1,
	-1, 59, -1, -1, 60, -15, -1, 60, 60, -1,
	60, 60, 30, -1, 60, 30, 59, 59, -1, -1,
	60, 60,
}
var yyDef = [...]int{

	-2, -2, 104, 3, 38, -2, 0, 8, 9, 45,
	0, 0, 13, 45, 0, 0, 49, 50, 0, 0,
	0, 0, 0, 58, 59, 60, 61, 0, 104, 104,
	0, 0, 0, 0, 1, -2, 106, 108, 109, 0,
	39, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 94, 95, 0, 0, 0, 0,
	0, 45, 45, 10, 46, 11, 0, 0, 49, -2,
	-2, 0, 0, -2, 0, 45, 0, 51, 52, 53,
	-2, 0, -2, 0, 38, 0, 45, 105, 35, 0,
	0, 104, 0, 4, 107, 45, 104, 6, 0, 63,
	74, 75, 76, 77, 78, 79, 80, 81, -2, -2,
	84, 85, 86, 87, 88, 89, 90, 91, 92, 93,
	96, 97, 98, 99, 0, 0, 7, -2, 104, -2,
	0, -2, 0, 0, -2, 0, 45, 0, 27, 0,
	0, 0, 0, 0, 39, 38, 104, 104, 36, 0,
	72, 45, 0, -2, 5, 0, 0, 102, 103, 0,
	0, 0, 0, 0, 15, 0, 0, 41, 0, 0,
	0, 100, 101, 55, 57, 0, 0, 0, 39, 104,
	0, 104, 0, 0, 0, -2, -2, 0, 40, 62,
	47, -2, 12, -2, 25, -2, 16, 0, 0, 22,
	30, 31, 0, 0, 28, 29, -2, 0, 0, 0,
	0, 68, 0, 70, 34, 73, 26, 0, 0, 0,
	0, -2, 0, 104, 0, -2, -2, 0, 69, 37,
	71, 24, 14, -2, -2, 0, 104, -2, 64, 0,
	0, -2, 0, 0, 21, -2, 33, 65, 66, 0,
	17, 20, 0, 32, 67, 0, -2, -2, 0, 0,
	19, 18,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	71, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 62, 3, 3, 3, 57, 64, 3,
	66, 67, 55, 53, 50, 54, 65, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 61,
	52, 47, 51, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 68, 3, 69, 63, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 59, 70, 60,
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

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

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
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
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

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:63
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:68
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:75
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[1].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:82
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
				if l, ok := yylex.(*Lexer); ok {
					l.stmts = yyVAL.stmts
				}
			}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:93
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:98
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:106
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:111
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:116
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:121
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:126
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 14:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyDollar[2].tok.Lit, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:141
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:146
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:151
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:156
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:166
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:171
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:176
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 24:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:189
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:194
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyDollar[4].compstmt...)
			}
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:203
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:209
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:213
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:217
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:221
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:225
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:236
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:242
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:248
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:253
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:257
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:261
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:266
		{
			yyVAL.expr_idents = []string{}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:270
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:285
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:289
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:293
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:298
		{
			yyVAL.exprs = nil
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:310
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:316
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:321
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:326
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:336
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:341
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:346
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:351
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:356
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:361
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:366
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:376
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:381
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:386
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 64:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:396
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:401
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:406
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 69:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:416
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:421
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[3].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:430
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[3].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:439
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:444
		{
			yyVAL.expr = &ast.NewExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:449
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:454
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:459
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:464
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:469
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:474
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:479
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:484
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:489
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:494
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:499
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:504
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:509
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:514
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:519
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:524
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:529
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:534
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:539
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:544
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:549
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:554
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:559
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:564
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:569
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:574
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:579
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 101:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:584
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:589
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:594
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:605
		{
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:608
		{
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:613
		{
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:616
		{
		}
	}
	goto yystack /* stack new state and value */
}
