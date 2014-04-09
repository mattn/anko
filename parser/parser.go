//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:28
type yySymType struct {
	yys                    int
	stmt_var               ast.Stmt
	stmt_if                ast.Stmt
	stmt_for               ast.Stmt
	stmt_try_catch_finally ast.Stmt
	stmt_switch            ast.Stmt
	stmt_default           ast.Stmt
	stmt_case              ast.Stmt
	stmt_cases             []ast.Stmt
	stmts                  []ast.Stmt
	stmt                   ast.Stmt
	expr                   ast.Expr
	exprs                  []ast.Expr
	expr_lhs               ast.Expr
	expr_lhss              []ast.Expr
	expr_pair              ast.Expr
	expr_pairs             []ast.Expr
	expr_idents            []string
	tok                    Token
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

//line parser.go.y:591

//line yacctab:1
var yyExca = []int{
	-1, 0,
	47, 45,
	-2, 1,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	47, 45,
	-2, 1,
	-1, 3,
	47, 45,
	-2, 1,
	-1, 4,
	47, 45,
	-2, 1,
	-1, 5,
	47, 45,
	-2, 1,
	-1, 6,
	47, 45,
	-2, 1,
	-1, 7,
	47, 45,
	-2, 1,
	-1, 8,
	47, 45,
	-2, 1,
	-1, 9,
	47, 45,
	-2, 1,
	-1, 11,
	47, 45,
	-2, 48,
	-1, 20,
	47, 42,
	50, 42,
	-2, 53,
	-1, 31,
	64, 48,
	-2, 45,
	-1, 38,
	47, 45,
	-2, 1,
	-1, 68,
	69, 48,
	-2, 45,
	-1, 71,
	47, 42,
	-2, 53,
	-1, 77,
	47, 42,
	50, 42,
	-2, 53,
	-1, 78,
	50, 36,
	61, 1,
	-2, 45,
	-1, 80,
	61, 1,
	-2, 45,
	-1, 90,
	69, 48,
	-2, 45,
	-1, 95,
	17, 53,
	18, 53,
	19, 53,
	20, 53,
	21, 53,
	22, 53,
	41, 53,
	42, 53,
	43, 53,
	47, 42,
	48, 53,
	50, 42,
	51, 53,
	52, 53,
	53, 53,
	54, 53,
	55, 53,
	56, 53,
	57, 53,
	62, 53,
	67, 53,
	70, 53,
	-2, 57,
	-1, 97,
	17, 53,
	18, 53,
	19, 53,
	20, 53,
	21, 53,
	22, 53,
	41, 53,
	42, 53,
	43, 53,
	47, 42,
	48, 53,
	50, 42,
	51, 53,
	52, 53,
	53, 53,
	54, 53,
	55, 53,
	56, 53,
	57, 53,
	62, 53,
	67, 53,
	70, 53,
	-2, 59,
	-1, 107,
	47, 45,
	-2, 48,
	-1, 111,
	61, 1,
	-2, 45,
	-1, 112,
	47, 43,
	50, 43,
	-2, 66,
	-1, 123,
	17, 0,
	18, 0,
	-2, 83,
	-1, 124,
	17, 0,
	18, 0,
	-2, 84,
	-1, 134,
	47, 45,
	-2, 48,
	-1, 135,
	47, 45,
	-2, 48,
	-1, 136,
	61, 1,
	-2, 45,
	-1, 137,
	47, 45,
	-2, 48,
	-1, 139,
	61, 1,
	-2, 45,
	-1, 143,
	61, 1,
	-2, 45,
	-1, 165,
	69, 48,
	-2, 45,
	-1, 171,
	47, 44,
	50, 44,
	-2, 104,
	-1, 192,
	47, 43,
	50, 43,
	-2, 58,
	-1, 193,
	47, 43,
	50, 43,
	-2, 60,
	-1, 201,
	61, 1,
	-2, 45,
	-1, 206,
	61, 1,
	-2, 45,
	-1, 214,
	47, 45,
	-2, 1,
	-1, 215,
	61, 1,
	-2, 45,
	-1, 224,
	61, 1,
	-2, 45,
	-1, 225,
	47, 45,
	-2, 1,
	-1, 228,
	61, 1,
	-2, 45,
	-1, 229,
	61, 1,
	-2, 45,
	-1, 233,
	61, 1,
	-2, 45,
	-1, 234,
	61, 1,
	-2, 45,
	-1, 240,
	61, 1,
	-2, 45,
	-1, 252,
	61, 1,
	-2, 45,
	-1, 253,
	61, 1,
	-2, 45,
}

const yyNprod = 106
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1493

var yyAct = []int{

	1, 35, 230, 37, 39, 40, 41, 42, 44, 45,
	46, 186, 187, 219, 103, 138, 138, 216, 190, 74,
	173, 20, 19, 26, 165, 159, 30, 11, 14, 12,
	15, 43, 16, 100, 217, 194, 55, 69, 160, 109,
	34, 27, 28, 29, 13, 17, 257, 256, 50, 51,
	52, 53, 54, 161, 3, 4, 253, 47, 48, 188,
	189, 18, 66, 68, 162, 64, 250, 248, 247, 101,
	245, 21, 25, 55, 244, 212, 243, 32, 223, 141,
	31, 145, 22, 23, 24, 33, 237, 52, 53, 54,
	108, 232, 231, 252, 47, 48, 207, 99, 205, 66,
	68, 10, 64, 204, 202, 184, 133, 181, 240, 234,
	167, 229, 170, 70, 72, 228, 215, 76, 79, 157,
	81, 136, 80, 92, 93, 94, 96, 98, 153, 137,
	110, 107, 138, 70, 224, 105, 214, 176, 163, 251,
	179, 188, 189, 249, 182, 166, 218, 209, 195, 104,
	113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
	123, 124, 125, 126, 127, 128, 129, 130, 131, 132,
	70, 197, 174, 175, 193, 177, 198, 111, 192, 196,
	178, 158, 112, 106, 147, 148, 149, 150, 151, 152,
	75, 73, 70, 154, 102, 36, 185, 210, 211, 20,
	19, 26, 220, 200, 30, 9, 8, 221, 7, 70,
	168, 6, 169, 5, 2, 226, 227, 0, 34, 27,
	28, 29, 0, 0, 0, 235, 236, 0, 0, 238,
	239, 0, 0, 0, 241, 242, 70, 70, 0, 70,
	0, 246, 180, 0, 0, 0, 183, 0, 0, 21,
	25, 0, 0, 254, 255, 32, 0, 0, 31, 0,
	22, 23, 24, 33, 0, 199, 0, 70, 0, 61,
	63, 0, 0, 0, 203, 20, 19, 26, 0, 0,
	30, 11, 14, 12, 15, 0, 16, 0, 0, 0,
	213, 55, 56, 57, 34, 27, 28, 29, 13, 17,
	0, 60, 62, 50, 51, 52, 53, 54, 3, 4,
	222, 0, 47, 48, 0, 18, 0, 66, 68, 0,
	64, 0, 0, 0, 0, 21, 25, 0, 0, 0,
	38, 32, 0, 0, 31, 0, 22, 23, 24, 33,
	20, 19, 26, 0, 0, 30, 11, 14, 12, 15,
	55, 16, 0, 0, 0, 0, 0, 0, 0, 34,
	27, 28, 29, 13, 17, 0, 0, 0, 0, 0,
	0, 47, 48, 3, 4, 0, 66, 68, 0, 64,
	18, 58, 59, 61, 63, 65, 67, 0, 0, 0,
	21, 25, 0, 0, 0, 0, 32, 0, 0, 31,
	0, 22, 23, 24, 33, 55, 56, 57, 0, 0,
	0, 0, 49, 0, 0, 60, 62, 50, 51, 52,
	53, 54, 0, 144, 143, 0, 47, 48, 0, 0,
	0, 66, 68, 0, 64, 20, 19, 142, 0, 0,
	30, 11, 14, 12, 15, 0, 16, 0, 0, 0,
	0, 0, 0, 0, 34, 27, 28, 29, 13, 17,
	0, 0, 0, 0, 0, 0, 0, 0, 3, 4,
	0, 0, 0, 0, 0, 18, 58, 59, 61, 63,
	65, 67, 0, 0, 0, 21, 25, 0, 0, 0,
	0, 32, 0, 0, 31, 0, 22, 23, 24, 33,
	55, 56, 57, 0, 0, 0, 0, 49, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 233,
	0, 47, 48, 0, 0, 140, 66, 68, 0, 64,
	58, 59, 61, 63, 65, 67, 0, 0, 0, 0,
	82, 83, 84, 85, 86, 87, 0, 0, 88, 89,
	0, 0, 0, 0, 55, 56, 57, 0, 0, 0,
	0, 49, 225, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 91, 0, 0, 47, 48, 90, 0, 0,
	66, 68, 0, 64, 58, 59, 61, 63, 65, 67,
	0, 0, 0, 0, 0, 82, 83, 84, 85, 86,
	87, 0, 0, 88, 89, 0, 0, 0, 55, 56,
	57, 0, 0, 0, 135, 49, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 208, 91, 0, 47,
	48, 0, 90, 0, 66, 68, 0, 64, 58, 59,
	61, 63, 65, 67, 0, 0, 0, 0, 82, 83,
	84, 85, 86, 87, 0, 0, 88, 89, 0, 0,
	0, 0, 55, 56, 57, 0, 0, 0, 0, 49,
	0, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	91, 206, 0, 47, 48, 90, 0, 0, 66, 68,
	0, 64, 58, 59, 61, 63, 65, 67, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 55, 56, 57, 0,
	0, 0, 0, 49, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 201, 0, 47, 48, 0,
	0, 0, 66, 68, 0, 64, 58, 59, 61, 63,
	65, 67, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	55, 56, 57, 0, 0, 0, 0, 49, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 47, 48, 191, 0, 0, 66, 68, 0, 64,
	58, 59, 61, 63, 65, 67, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 56, 57, 0, 0, 0,
	0, 49, 172, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 47, 48, 0, 0, 0,
	66, 68, 0, 64, 58, 59, 61, 63, 65, 67,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 55, 56,
	57, 0, 0, 0, 0, 49, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 47,
	48, 171, 0, 0, 66, 68, 0, 64, 58, 59,
	61, 63, 65, 67, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 55, 56, 57, 0, 0, 0, 0, 49,
	0, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 47, 48, 0, 0, 0, 66, 68,
	164, 64, 58, 59, 61, 63, 65, 67, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 55, 56, 57, 0,
	0, 0, 0, 49, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 146, 0, 47, 48, 0,
	0, 0, 66, 68, 0, 64, 58, 59, 61, 63,
	65, 67, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	55, 56, 57, 0, 0, 0, 0, 49, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 139,
	0, 47, 48, 0, 0, 0, 66, 68, 0, 64,
	58, 59, 61, 63, 65, 67, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 56, 57, 0, 0, 0,
	0, 49, 0, 134, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 47, 48, 0, 0, 0,
	66, 68, 0, 64, 58, 59, 61, 63, 65, 67,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 55, 56,
	57, 0, 0, 0, 0, 49, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 47,
	48, 0, 0, 0, 66, 68, 0, 64, 58, 59,
	61, 63, 65, 67, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 55, 56, 57, 0, 0, 0, 0, 49,
	0, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 156, 48, 0, 0, 0, 66, 68,
	0, 64, 58, 59, 61, 63, 65, 67, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 55, 56, 57, 0,
	0, 0, 0, 49, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 0, 155, 48, 0,
	0, 0, 66, 68, 0, 64, 58, 59, 61, 63,
	0, 67, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	55, 56, 57, 0, 0, 0, 58, 59, 61, 63,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 47, 48, 0, 0, 0, 66, 68, 0, 64,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 71, 19,
	26, 47, 48, 30, 0, 0, 66, 68, 0, 64,
	0, 0, 0, 0, 0, 0, 0, 34, 27, 28,
	29, 97, 19, 26, 0, 0, 30, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	34, 27, 28, 29, 0, 95, 19, 26, 21, 25,
	30, 0, 0, 0, 32, 0, 0, 31, 0, 22,
	23, 24, 33, 0, 34, 27, 28, 29, 77, 19,
	26, 21, 25, 30, 0, 0, 0, 32, 0, 0,
	31, 0, 22, 23, 24, 33, 0, 34, 27, 28,
	29, 0, 0, 0, 0, 21, 25, 0, 0, 0,
	0, 32, 0, 0, 31, 0, 22, 23, 24, 33,
	0, 0, 0, 0, 0, 0, 0, 0, 21, 25,
	0, 0, 0, 0, 78, 0, 0, 31, 0, 22,
	23, 24, 33,
}
var yyPact = []int{

	336, -1000, 271, 336, 336, 336, 17, 336, 336, 336,
	1107, 1354, 195, 187, 186, 195, 1424, 62, 195, -1000,
	617, 195, 195, 195, 1401, 1377, -1000, -1000, -1000, -1000,
	29, 1354, 143, 195, 179, 84, 40, -1000, 336, -1000,
	-1000, -1000, -1000, 117, -1000, -1000, -1000, 178, 195, 195,
	195, 195, 195, 195, 195, 195, 195, 195, 195, 195,
	195, 195, 195, 195, 195, 195, 195, 195, 1354, -1000,
	1053, 564, 1107, 61, 82, -1000, 999, 509, 431, 364,
	336, 945, 195, 195, 195, 195, 195, 195, -1000, -1000,
	1354, 195, 309, 309, 309, 617, 1215, 617, 1161, 177,
	-43, -26, 3, -1000, 89, 891, -44, 1354, 195, -1000,
	195, 336, -1000, 837, 783, 32, 32, 309, 309, 309,
	1107, -5, -5, 250, 250, -5, -5, -5, -5, 1107,
	1269, 1107, 1299, -49, 1354, 1354, 336, 1354, 176, 336,
	195, 46, 89, 336, 195, 44, 96, 1107, 1107, 1107,
	1107, 1107, 1107, -51, 729, 174, 170, -34, 140, 167,
	-1000, 143, -1000, 195, -1000, 1354, -1000, -1000, 1107, 675,
	43, -1000, 195, -1000, -1000, -1000, 42, -1000, -1000, 37,
	621, -1000, 35, 567, 118, 14, -1000, -1000, 195, 87,
	-1000, -1000, -1000, -1000, 56, -52, -35, 138, -1000, 1107,
	-56, 336, -1000, 1107, -1000, -1000, 336, -1000, 195, 74,
	-1000, -1000, -1000, 513, 336, 336, 55, 51, -67, -1000,
	31, 30, 459, 49, 336, 336, -1000, 25, 336, 336,
	48, -1000, -1000, 336, 336, 15, -1000, -1000, 13, 9,
	336, 7, 6, 113, -1000, -1000, 5, -1000, 109, 33,
	-1000, -4, 336, 336, -14, -15, -1000, -1000,
}
var yyPgo = []int{

	0, 0, 214, 213, 211, 208, 206, 205, 12, 11,
	196, 101, 37, 195, 1, 14, 194, 19,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 2, 2, 3, 4, 4, 4, 5,
	5, 5, 5, 6, 6, 6, 6, 10, 10, 10,
	10, 10, 9, 8, 7, 15, 16, 16, 16, 17,
	17, 17, 13, 13, 13, 14, 14, 14, 12, 12,
	12, 12, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	2, 1, 2, 2, 5, 4, 7, 5, 5, 7,
	4, 5, 9, 13, 12, 9, 8, 0, 1, 2,
	1, 2, 4, 3, 5, 3, 0, 1, 3, 0,
	1, 3, 1, 3, 4, 0, 1, 3, 0, 1,
	3, 3, 1, 1, 2, 2, 2, 2, 4, 2,
	4, 1, 1, 1, 1, 5, 3, 7, 8, 8,
	9, 3, 3, 3, 5, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 3, 3,
	3, 3, 4, 4, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, 37, 38, -3, -4, -5, -6, -7,
	-11, 10, 12, 27, 11, 13, 15, 28, 44, 5,
	4, 54, 65, 66, 67, 55, 6, 24, 25, 26,
	9, 63, 60, 68, 23, -14, -13, -1, 59, -1,
	-1, -1, -1, 14, -1, -1, -1, 62, 63, 48,
	53, 54, 55, 56, 57, 41, 42, 43, 17, 18,
	51, 19, 52, 20, 70, 21, 67, 22, 68, -12,
	-11, 4, -11, 4, -17, 4, -11, 4, 60, -11,
	60, -11, 31, 32, 33, 34, 35, 36, 39, 40,
	68, 63, -11, -11, -11, 4, -11, 4, -11, 68,
	4, -12, -16, -15, 6, -11, 4, 47, 50, -1,
	13, 60, 4, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -12, 50, 50, 60, 47, 50, 60,
	16, -1, 6, 60, 59, -1, 60, -11, -11, -11,
	-11, -11, -11, -12, -11, 62, 62, -17, 4, 68,
	64, 50, 61, 49, 69, 68, -12, -14, -11, -11,
	-1, 64, 49, 69, -12, -12, -1, -12, 4, -1,
	-11, 61, -1, -11, 61, -10, -9, -8, 45, 46,
	69, 64, 4, 4, 69, 8, -17, 4, -15, -11,
	-12, 60, 61, -11, 61, 61, 60, 61, 59, 29,
	-9, -8, 61, -11, 49, 60, 69, 69, 8, 69,
	-1, -1, -11, 4, 60, 49, -1, -1, 60, 60,
	69, 61, 61, 60, 60, -1, -1, 61, -1, -1,
	60, -1, -1, 61, 61, 61, -1, 61, 61, 30,
	61, 30, 60, 60, -1, -1, 61, 61,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	11, -2, 45, 0, 39, 45, 45, 0, 45, 52,
	-2, 45, 45, 45, 45, 45, 61, 62, 63, 64,
	0, -2, 36, 45, 0, 0, 46, 2, -2, 4,
	5, 6, 7, 0, 8, 9, 10, 0, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, -2, 12,
	49, -2, 13, 0, 0, 40, 0, -2, -2, 0,
	-2, 0, 45, 45, 45, 45, 45, 45, 96, 97,
	-2, 45, 54, 55, 56, -2, 0, -2, 0, 39,
	0, 0, 0, 37, 0, 0, 0, -2, 45, 3,
	45, -2, -2, 0, 0, 75, 76, 77, 78, 79,
	80, 81, 82, -2, -2, 85, 86, 87, 88, 98,
	99, 100, 101, 0, -2, -2, -2, -2, 0, -2,
	45, 0, 61, -2, 45, 0, 27, 90, 91, 92,
	93, 94, 95, 0, 0, 0, 0, 0, 40, 39,
	71, 0, 72, 45, 73, -2, 89, 47, 0, 0,
	0, -2, 45, 105, 50, 51, 0, 15, 41, 0,
	0, 20, 0, 0, 0, 0, 28, 30, 45, 0,
	102, 103, -2, -2, 0, 0, 0, 40, 38, 35,
	0, -2, 17, 65, 14, 18, -2, 21, 45, 0,
	29, 31, 34, 0, -2, -2, 0, 0, 0, 74,
	0, 0, 0, 0, -2, -2, 33, 0, -2, -2,
	0, 16, 19, -2, -2, 0, 32, 67, 0, 0,
	-2, 0, 0, 26, 68, 69, 0, 22, 25, 0,
	70, 0, -2, -2, 0, 0, 24, 23,
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
		//line parser.go.y:67
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:74
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-1].stmt.SetPos(l.pos)
			}
		}
	case 3:
		//line parser.go.y:82
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-2].stmt.SetPos(l.pos)
			}
		}
	case 4:
		//line parser.go.y:90
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.BreakStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:95
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.ContinueStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:100
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:105
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:110
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 9:
		//line parser.go.y:115
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 10:
		//line parser.go.y:120
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_switch}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 11:
		//line parser.go.y:126
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 12:
		//line parser.go.y:130
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 13:
		//line parser.go.y:134
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
		}
	case 14:
		//line parser.go.y:138
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 15:
		//line parser.go.y:143
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].expr_idents, Exprs: yyS[yypt-0].exprs}
		}
	case 16:
		//line parser.go.y:148
		{
			yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-6].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts})
		}
	case 17:
		//line parser.go.y:152
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyS[yypt-1].stmts...)
			}
		}
	case 18:
		//line parser.go.y:160
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-3].expr, Then: yyS[yypt-1].stmts, Else: nil}
		}
	case 19:
		//line parser.go.y:165
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:169
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:173
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:177
		{
			yyVAL.stmt_for = &ast.CForStmt{Expr1: yyS[yypt-7].expr, Expr2: yyS[yypt-5].expr, Expr3: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:182
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-10].stmts, Var: yyS[yypt-7].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:186
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:190
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-6].stmts, Var: yyS[yypt-3].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 26:
		//line parser.go.y:194
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 27:
		//line parser.go.y:199
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 28:
		//line parser.go.y:203
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_case}
		}
	case 29:
		//line parser.go.y:207
		{
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_case)
		}
	case 30:
		//line parser.go.y:211
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_default}
		}
	case 31:
		//line parser.go.y:215
		{
			for _, stmt := range yyS[yypt-1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_default)
		}
	case 32:
		//line parser.go.y:225
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyS[yypt-2].expr, Stmts: yyS[yypt-0].stmts}
		}
	case 33:
		//line parser.go.y:230
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyS[yypt-0].stmts}
		}
	case 34:
		//line parser.go.y:235
		{
			yyVAL.stmt_switch = &ast.SwitchStmt{Expr: yyS[yypt-3].expr, Cases: yyS[yypt-1].stmt_cases}
		}
	case 35:
		//line parser.go.y:240
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 36:
		//line parser.go.y:245
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 37:
		//line parser.go.y:249
		{
			yyVAL.expr_pairs = []ast.Expr{yyS[yypt-0].expr_pair}
		}
	case 38:
		//line parser.go.y:253
		{
			yyVAL.expr_pairs = append(yyS[yypt-2].expr_pairs, yyS[yypt-0].expr_pair)
		}
	case 39:
		//line parser.go.y:258
		{
			yyVAL.expr_idents = []string{}
		}
	case 40:
		//line parser.go.y:262
		{
			yyVAL.expr_idents = []string{yyS[yypt-0].tok.lit}
		}
	case 41:
		//line parser.go.y:266
		{
			yyVAL.expr_idents = append(yyS[yypt-2].expr_idents, yyS[yypt-0].tok.lit)
		}
	case 42:
		//line parser.go.y:271
		{
			yyVAL.expr_lhs = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPos(l.pos)
			}
		}
	case 43:
		//line parser.go.y:276
		{
			yyVAL.expr_lhs = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPos(l.pos)
			}
		}
	case 44:
		//line parser.go.y:281
		{
			yyVAL.expr_lhs = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPos(l.pos)
			}
		}
	case 45:
		//line parser.go.y:287
		{
			yyVAL.expr_lhss = []ast.Expr{}
		}
	case 46:
		//line parser.go.y:291
		{
			yyVAL.expr_lhss = []ast.Expr{yyS[yypt-0].expr_lhs}
		}
	case 47:
		//line parser.go.y:295
		{
			yyVAL.expr_lhss = append([]ast.Expr{yyS[yypt-2].expr_lhs}, yyS[yypt-0].expr_lhss...)
		}
	case 48:
		//line parser.go.y:300
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 49:
		//line parser.go.y:304
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 50:
		//line parser.go.y:308
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 51:
		//line parser.go.y:312
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 52:
		//line parser.go.y:317
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:322
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:332
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:337
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:342
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 59:
		//line parser.go.y:352
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 60:
		//line parser.go.y:357
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 61:
		//line parser.go.y:362
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 62:
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:372
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 64:
		//line parser.go.y:377
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 65:
		//line parser.go.y:382
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 66:
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 67:
		//line parser.go.y:392
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 68:
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 69:
		//line parser.go.y:402
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 70:
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 71:
		//line parser.go.y:412
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 72:
		//line parser.go.y:417
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
	case 73:
		//line parser.go.y:426
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 74:
		//line parser.go.y:431
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 75:
		//line parser.go.y:436
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 76:
		//line parser.go.y:441
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 77:
		//line parser.go.y:446
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 78:
		//line parser.go.y:451
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 79:
		//line parser.go.y:456
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 80:
		//line parser.go.y:461
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 81:
		//line parser.go.y:466
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 82:
		//line parser.go.y:471
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 83:
		//line parser.go.y:476
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 84:
		//line parser.go.y:481
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 85:
		//line parser.go.y:486
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 86:
		//line parser.go.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 87:
		//line parser.go.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 88:
		//line parser.go.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 89:
		//line parser.go.y:506
		{
			yyVAL.expr = &ast.LetsExpr{Lhss: yyS[yypt-2].expr_lhss, Operator: "=", Rhss: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 90:
		//line parser.go.y:511
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "+=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 91:
		//line parser.go.y:516
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "-=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 92:
		//line parser.go.y:521
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "*=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 93:
		//line parser.go.y:526
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "/=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 94:
		//line parser.go.y:531
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "&=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 95:
		//line parser.go.y:536
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "|=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 96:
		//line parser.go.y:541
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 97:
		//line parser.go.y:546
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 98:
		//line parser.go.y:551
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 99:
		//line parser.go.y:556
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 100:
		//line parser.go.y:561
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 101:
		//line parser.go.y:566
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 102:
		//line parser.go.y:571
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 103:
		//line parser.go.y:576
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyS[yypt-3].tok.lit}, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 104:
		//line parser.go.y:581
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 105:
		//line parser.go.y:586
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
