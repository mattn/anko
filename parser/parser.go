//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
	"unsafe"
)

//line parser.go.y:29
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

//line parser.go.y:592

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
	66, 48,
	-2, 45,
	-1, 38,
	47, 45,
	-2, 1,
	-1, 68,
	63, 48,
	-2, 45,
	-1, 71,
	47, 42,
	-2, 53,
	-1, 78,
	61, 1,
	-2, 45,
	-1, 80,
	61, 1,
	-2, 45,
	-1, 90,
	63, 48,
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
	64, 53,
	69, 53,
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
	64, 53,
	69, 53,
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
	-1, 163,
	63, 48,
	-2, 45,
	-1, 169,
	47, 44,
	50, 44,
	-2, 104,
	-1, 190,
	47, 43,
	50, 43,
	-2, 58,
	-1, 191,
	47, 43,
	50, 43,
	-2, 60,
	-1, 203,
	61, 1,
	-2, 45,
	-1, 204,
	61, 1,
	-2, 45,
	-1, 205,
	61, 1,
	-2, 45,
	-1, 212,
	47, 45,
	-2, 1,
	-1, 213,
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
	-1, 231,
	61, 1,
	-2, 45,
	-1, 242,
	61, 1,
	-2, 45,
	-1, 251,
	61, 1,
	-2, 45,
	-1, 252,
	61, 1,
	-2, 45,
	-1, 257,
	61, 1,
	-2, 45,
	-1, 263,
	61, 1,
	-2, 45,
}

const yyNprod = 106
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1476

var yyAct = []int{

	1, 35, 158, 37, 39, 40, 41, 42, 44, 45,
	46, 184, 185, 138, 103, 245, 230, 224, 74, 223,
	167, 138, 20, 19, 26, 217, 215, 30, 11, 14,
	12, 15, 43, 16, 192, 77, 55, 69, 214, 109,
	188, 34, 27, 28, 29, 13, 17, 55, 50, 51,
	52, 53, 54, 171, 163, 3, 4, 68, 100, 47,
	48, 157, 18, 76, 66, 64, 186, 187, 68, 101,
	47, 48, 21, 25, 55, 66, 64, 265, 32, 141,
	33, 143, 210, 31, 262, 22, 23, 24, 52, 53,
	54, 78, 110, 79, 159, 68, 259, 47, 48, 258,
	254, 10, 66, 64, 250, 160, 133, 248, 247, 246,
	165, 239, 168, 70, 72, 263, 99, 234, 155, 233,
	81, 232, 202, 92, 93, 94, 96, 98, 151, 200,
	182, 179, 257, 70, 252, 105, 251, 174, 242, 111,
	231, 229, 228, 213, 205, 164, 203, 136, 80, 108,
	113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
	123, 124, 125, 126, 127, 128, 129, 130, 131, 132,
	70, 212, 172, 173, 196, 175, 194, 137, 139, 107,
	138, 142, 161, 261, 145, 146, 147, 148, 149, 150,
	186, 187, 70, 152, 253, 208, 209, 207, 140, 216,
	193, 198, 104, 236, 219, 220, 221, 195, 191, 70,
	166, 190, 176, 226, 227, 82, 83, 84, 85, 86,
	87, 156, 112, 88, 89, 237, 238, 106, 75, 240,
	241, 73, 243, 102, 135, 36, 70, 70, 183, 70,
	9, 8, 178, 249, 7, 6, 90, 5, 2, 91,
	0, 0, 255, 256, 0, 0, 0, 0, 260, 0,
	0, 0, 0, 197, 264, 70, 0, 0, 0, 199,
	0, 0, 201, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 206, 0, 20, 19, 26, 211, 0,
	30, 11, 14, 12, 15, 0, 16, 0, 0, 0,
	0, 0, 0, 0, 34, 27, 28, 29, 13, 17,
	0, 0, 0, 0, 0, 0, 0, 0, 3, 4,
	0, 0, 0, 0, 235, 18, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 21, 25, 0, 0, 0,
	38, 32, 0, 33, 0, 0, 31, 0, 22, 23,
	24, 20, 19, 26, 0, 0, 30, 11, 14, 12,
	15, 0, 16, 0, 0, 0, 0, 0, 0, 0,
	34, 27, 28, 29, 13, 17, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 4, 0, 0, 0, 0,
	0, 18, 0, 58, 59, 61, 63, 65, 67, 0,
	0, 21, 25, 0, 0, 0, 0, 32, 0, 33,
	0, 0, 31, 0, 22, 23, 24, 55, 56, 57,
	0, 0, 0, 0, 49, 0, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 181, 0, 0, 68, 180,
	47, 48, 0, 0, 0, 66, 64, 58, 59, 61,
	63, 65, 67, 0, 0, 0, 0, 82, 83, 84,
	85, 86, 87, 0, 0, 88, 89, 0, 0, 0,
	0, 55, 56, 57, 0, 0, 0, 0, 49, 0,
	0, 60, 62, 50, 51, 52, 53, 54, 90, 0,
	0, 91, 68, 244, 47, 48, 0, 0, 0, 66,
	64, 58, 59, 61, 63, 65, 67, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 55, 56, 57, 0, 0,
	0, 0, 49, 225, 0, 60, 62, 50, 51, 52,
	53, 54, 0, 0, 0, 0, 68, 0, 47, 48,
	0, 0, 0, 66, 64, 58, 59, 61, 63, 65,
	67, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 55,
	56, 57, 0, 0, 0, 0, 49, 0, 0, 60,
	62, 50, 51, 52, 53, 54, 0, 222, 0, 0,
	68, 0, 47, 48, 0, 0, 0, 66, 64, 58,
	59, 61, 63, 65, 67, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 55, 56, 57, 0, 0, 0, 0,
	49, 0, 0, 60, 62, 50, 51, 52, 53, 54,
	0, 0, 0, 0, 68, 218, 47, 48, 0, 0,
	0, 66, 64, 58, 59, 61, 63, 65, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 55, 56, 57,
	0, 0, 0, 0, 49, 0, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 0, 204, 0, 68, 0,
	47, 48, 0, 0, 0, 66, 64, 58, 59, 61,
	63, 65, 67, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 55, 56, 57, 0, 0, 0, 0, 49, 0,
	0, 60, 62, 50, 51, 52, 53, 54, 0, 0,
	0, 0, 68, 0, 47, 48, 189, 0, 0, 66,
	64, 58, 59, 61, 63, 65, 67, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 55, 56, 57, 0, 0,
	0, 0, 49, 0, 0, 60, 62, 50, 51, 52,
	53, 54, 0, 0, 0, 0, 68, 177, 47, 48,
	0, 0, 0, 66, 64, 58, 59, 61, 63, 65,
	67, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 55,
	56, 57, 0, 0, 0, 0, 49, 170, 0, 60,
	62, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	68, 0, 47, 48, 0, 0, 0, 66, 64, 58,
	59, 61, 63, 65, 67, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 55, 56, 57, 0, 0, 0, 0,
	49, 0, 0, 60, 62, 50, 51, 52, 53, 54,
	0, 0, 0, 0, 68, 0, 47, 48, 169, 0,
	0, 66, 64, 58, 59, 61, 63, 65, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 55, 56, 57,
	0, 0, 0, 0, 49, 0, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 0, 0, 0, 68, 162,
	47, 48, 0, 0, 0, 66, 64, 58, 59, 61,
	63, 65, 67, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 55, 56, 57, 0, 0, 0, 0, 49, 0,
	0, 60, 62, 50, 51, 52, 53, 54, 0, 0,
	144, 0, 68, 0, 47, 48, 0, 0, 0, 66,
	64, 58, 59, 61, 63, 65, 67, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 55, 56, 57, 0, 0,
	0, 0, 49, 0, 134, 60, 62, 50, 51, 52,
	53, 54, 0, 0, 0, 0, 68, 0, 47, 48,
	0, 0, 0, 66, 64, 58, 59, 61, 63, 65,
	67, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 55,
	56, 57, 0, 0, 0, 0, 49, 0, 0, 60,
	62, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	68, 0, 47, 48, 0, 0, 0, 66, 64, 58,
	59, 61, 63, 65, 67, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 55, 56, 57, 0, 0, 0, 0,
	49, 0, 0, 60, 62, 50, 51, 52, 53, 54,
	0, 0, 0, 0, 68, 0, 154, 48, 0, 0,
	0, 66, 64, 58, 59, 61, 63, 65, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 55, 56, 57,
	0, 0, 0, 0, 49, 0, 0, 60, 62, 50,
	51, 52, 53, 54, 58, 59, 61, 63, 68, 67,
	153, 48, 0, 0, 0, 66, 64, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 55, 56,
	57, 0, 0, 0, 58, 59, 61, 63, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 68,
	0, 47, 48, 0, 0, 0, 66, 64, 55, 56,
	57, 61, 63, 0, 0, 0, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 68,
	0, 47, 48, 55, 56, 57, 66, 64, 0, 0,
	0, 0, 0, 60, 62, 50, 51, 52, 53, 54,
	20, 19, 26, 0, 68, 30, 47, 48, 0, 0,
	0, 66, 64, 0, 0, 0, 0, 0, 0, 34,
	27, 28, 29, 71, 19, 26, 0, 0, 30, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 34, 27, 28, 29, 0, 97, 19, 26,
	21, 25, 30, 0, 0, 0, 32, 0, 33, 0,
	0, 31, 0, 22, 23, 24, 34, 27, 28, 29,
	95, 19, 26, 21, 25, 30, 0, 0, 0, 32,
	0, 33, 0, 0, 31, 0, 22, 23, 24, 34,
	27, 28, 29, 0, 0, 0, 0, 21, 25, 0,
	0, 0, 0, 32, 0, 33, 0, 0, 31, 0,
	22, 23, 24, 0, 0, 0, 0, 0, 0, 0,
	21, 25, 0, 0, 0, 0, 32, 0, 33, 0,
	0, 31, 0, 22, 23, 24,
}
var yyPact = []int{

	347, -1000, 281, 347, 347, 347, 18, 347, 347, 347,
	1078, 1359, 1336, 227, 224, 1, 31, 88, 1336, -1000,
	426, 1336, 1336, 1336, 1406, 1383, -1000, -1000, -1000, -1000,
	54, 1359, 196, 1336, 223, 132, 99, -1000, 347, -1000,
	-1000, -1000, -1000, 79, -1000, -1000, -1000, 218, 1336, 1336,
	1336, 1336, 1336, 1336, 1336, 1336, 1336, 1336, 1336, 1336,
	1336, 1336, 1336, 1336, 1336, 1336, 1336, 1336, 1359, -1000,
	1024, 184, 1078, 87, 130, -1000, 1336, 182, 347, 1336,
	347, 970, 1336, 1336, 1336, 1336, 1336, 1336, -1000, -1000,
	1359, 1336, 6, 6, 6, 426, 1186, 426, 1132, 217,
	-1, -64, 44, -1000, 133, 916, -8, 1359, 1336, -1000,
	-42, 347, -1000, 862, 808, 33, 33, 6, 6, 6,
	1078, -5, -5, 1282, 1282, -5, -5, -5, -5, 1078,
	1227, 1078, 1257, -10, 1359, 1359, 347, 1359, 208, 754,
	1336, 70, 376, 69, 145, 1078, 1078, 1078, 1078, 1078,
	1078, -23, 700, 207, 204, -29, 192, 203, -1000, 196,
	-1000, 1336, -1000, 1359, -1000, -1000, 1078, 1336, 68, -1000,
	1336, -1000, -1000, -1000, 61, -1000, -1000, 86, 646, -1000,
	84, 1336, 168, 21, -1000, -1000, 1336, 122, -1000, -1000,
	-1000, -1000, 83, -25, -37, 191, -1000, 1078, -38, 592,
	-1000, 1078, -1000, 347, 347, 347, 538, -43, -1000, -1000,
	-1000, 484, 347, 347, 82, 81, -47, -1000, 80, 60,
	58, 56, 1336, 199, 347, 347, -1000, 50, 347, 347,
	78, 347, -1000, -1000, -1000, 430, -48, 48, -1000, -1000,
	47, 46, 347, 43, 76, 74, 164, -1000, -1000, 39,
	-1000, 347, 347, 72, -1000, 38, 35, 347, -1000, 153,
	23, 55, -1000, 347, 16, -1000,
}
var yyPgo = []int{

	0, 0, 248, 247, 245, 244, 241, 240, 12, 11,
	238, 101, 37, 235, 1, 14, 233, 18,
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
	2, 1, 2, 2, 5, 4, 9, 5, 7, 7,
	4, 7, 11, 15, 12, 11, 8, 0, 1, 2,
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
	4, 54, 67, 68, 69, 55, 6, 24, 25, 26,
	9, 65, 60, 62, 23, -14, -13, -1, 59, -1,
	-1, -1, -1, 14, -1, -1, -1, 64, 65, 48,
	53, 54, 55, 56, 57, 41, 42, 43, 17, 18,
	51, 19, 52, 20, 70, 21, 69, 22, 62, -12,
	-11, 4, -11, 4, -17, 4, 62, 4, 60, 62,
	60, -11, 31, 32, 33, 34, 35, 36, 39, 40,
	62, 65, -11, -11, -11, 4, -11, 4, -11, 62,
	4, -12, -16, -15, 6, -11, 4, 47, 50, -1,
	13, 60, 4, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -12, 50, 50, 60, 47, 50, -11,
	16, -1, -11, -1, 60, -11, -11, -11, -11, -11,
	-11, -12, -11, 64, 64, -17, 4, 62, 66, 50,
	61, 49, 63, 62, -12, -14, -11, 62, -1, 66,
	49, 63, -12, -12, -1, -12, 4, 63, -11, 61,
	63, 59, 61, -10, -9, -8, 45, 46, 63, 66,
	4, 4, 63, 8, -17, 4, -15, -11, -12, -11,
	61, -11, 61, 60, 60, 60, -11, 29, -9, -8,
	61, -11, 49, 60, 63, 63, 8, 63, 63, -1,
	-1, -1, 59, 62, 60, 49, -1, -1, 60, 60,
	63, 60, 61, 61, 61, -11, 4, -1, -1, 61,
	-1, -1, 60, -1, 63, 63, 61, 61, 61, -1,
	61, 60, 60, 30, 61, -1, -1, 60, 61, 61,
	-1, 30, 61, 60, -1, 61,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	11, -2, 45, 0, 39, 0, 0, 0, 45, 52,
	-2, 45, 45, 45, 45, 45, 61, 62, 63, 64,
	0, -2, 36, 45, 0, 0, 46, 2, -2, 4,
	5, 6, 7, 0, 8, 9, 10, 0, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, -2, 12,
	49, -2, 13, 0, 0, 40, 45, 0, -2, 45,
	-2, 0, 45, 45, 45, 45, 45, 45, 96, 97,
	-2, 45, 54, 55, 56, -2, 0, -2, 0, 39,
	0, 0, 0, 37, 0, 0, 0, -2, 45, 3,
	0, -2, -2, 0, 0, 75, 76, 77, 78, 79,
	80, 81, 82, -2, -2, 85, 86, 87, 88, 98,
	99, 100, 101, 0, -2, -2, -2, -2, 0, 0,
	45, 0, 0, 0, 27, 90, 91, 92, 93, 94,
	95, 0, 0, 0, 0, 0, 40, 39, 71, 0,
	72, 45, 73, -2, 89, 47, 0, 45, 0, -2,
	45, 105, 50, 51, 0, 15, 41, 0, 0, 20,
	0, 45, 0, 0, 28, 30, 45, 0, 102, 103,
	-2, -2, 0, 0, 0, 40, 38, 35, 0, 0,
	17, 65, 14, -2, -2, -2, 0, 0, 29, 31,
	34, 0, -2, -2, 0, 0, 0, 74, 0, 0,
	0, 0, 45, 0, -2, -2, 33, 0, -2, -2,
	0, -2, 18, 19, 21, 0, 0, 0, 32, 67,
	0, 0, -2, 0, 0, 0, 26, 68, 69, 0,
	16, -2, -2, 0, 70, 0, 0, -2, 22, 25,
	0, 0, 24, -2, 0, 23,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 67, 3, 3, 3, 57, 69, 3,
	62, 63, 55, 53, 50, 54, 64, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 59,
	52, 47, 51, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 65, 3, 66, 68, 3, 3, 3, 3, 3,
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
		//line parser.go.y:68
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:75
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-1].stmt.SetPos(l.pos)
			}
		}
	case 3:
		//line parser.go.y:83
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-2].stmt.SetPos(l.pos)
			}
		}
	case 4:
		//line parser.go.y:91
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.BreakStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:96
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.ContinueStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:101
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:106
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
		//line parser.go.y:116
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 10:
		//line parser.go.y:121
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_switch}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 11:
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 12:
		//line parser.go.y:131
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 13:
		//line parser.go.y:135
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
		}
	case 14:
		//line parser.go.y:139
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 15:
		//line parser.go.y:144
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].expr_idents, Exprs: yyS[yypt-0].exprs}
		}
	case 16:
		//line parser.go.y:149
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 17:
		//line parser.go.y:153
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 18:
		//line parser.go.y:161
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:166
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:170
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:174
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:178
		{
			yyVAL.stmt_for = &ast.CForStmt{Expr1: yyS[yypt-8].expr, Expr2: yyS[yypt-6].expr, Expr3: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:183
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:187
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:191
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 26:
		//line parser.go.y:195
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 27:
		//line parser.go.y:200
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 28:
		//line parser.go.y:204
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_case}
		}
	case 29:
		//line parser.go.y:208
		{
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_case)
		}
	case 30:
		//line parser.go.y:212
		{
			yyVAL.stmt_cases = []ast.Stmt{yyS[yypt-0].stmt_default}
		}
	case 31:
		//line parser.go.y:216
		{
			for _, stmt := range yyS[yypt-1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyS[yypt-1].stmt_cases, yyS[yypt-0].stmt_default)
		}
	case 32:
		//line parser.go.y:226
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyS[yypt-2].expr, Stmts: yyS[yypt-0].stmts}
		}
	case 33:
		//line parser.go.y:231
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyS[yypt-0].stmts}
		}
	case 34:
		//line parser.go.y:236
		{
			yyVAL.stmt_switch = &ast.SwitchStmt{Expr: yyS[yypt-3].expr, Cases: yyS[yypt-1].stmt_cases}
		}
	case 35:
		//line parser.go.y:241
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 36:
		//line parser.go.y:246
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 37:
		//line parser.go.y:250
		{
			yyVAL.expr_pairs = []ast.Expr{yyS[yypt-0].expr_pair}
		}
	case 38:
		//line parser.go.y:254
		{
			yyVAL.expr_pairs = append(yyS[yypt-2].expr_pairs, yyS[yypt-0].expr_pair)
		}
	case 39:
		//line parser.go.y:259
		{
			yyVAL.expr_idents = []string{}
		}
	case 40:
		//line parser.go.y:263
		{
			yyVAL.expr_idents = []string{yyS[yypt-0].tok.lit}
		}
	case 41:
		//line parser.go.y:267
		{
			yyVAL.expr_idents = append(yyS[yypt-2].expr_idents, yyS[yypt-0].tok.lit)
		}
	case 42:
		//line parser.go.y:272
		{
			yyVAL.expr_lhs = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPos(l.pos)
			}
		}
	case 43:
		//line parser.go.y:277
		{
			yyVAL.expr_lhs = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPos(l.pos)
			}
		}
	case 44:
		//line parser.go.y:282
		{
			yyVAL.expr_lhs = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr_lhs.SetPos(l.pos)
			}
		}
	case 45:
		//line parser.go.y:288
		{
			yyVAL.expr_lhss = []ast.Expr{}
		}
	case 46:
		//line parser.go.y:292
		{
			yyVAL.expr_lhss = []ast.Expr{yyS[yypt-0].expr_lhs}
		}
	case 47:
		//line parser.go.y:296
		{
			yyVAL.expr_lhss = append([]ast.Expr{yyS[yypt-2].expr_lhs}, yyS[yypt-0].expr_lhss...)
		}
	case 48:
		//line parser.go.y:301
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 49:
		//line parser.go.y:305
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 50:
		//line parser.go.y:309
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 51:
		//line parser.go.y:313
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 52:
		//line parser.go.y:318
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:323
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:328
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:333
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:338
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:348
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 59:
		//line parser.go.y:353
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 60:
		//line parser.go.y:358
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 61:
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 62:
		//line parser.go.y:368
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:373
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 64:
		//line parser.go.y:378
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 65:
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 66:
		//line parser.go.y:388
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 67:
		//line parser.go.y:393
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 68:
		//line parser.go.y:398
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 69:
		//line parser.go.y:403
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 70:
		//line parser.go.y:408
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 71:
		//line parser.go.y:413
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 72:
		//line parser.go.y:418
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
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 74:
		//line parser.go.y:432
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 75:
		//line parser.go.y:437
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 76:
		//line parser.go.y:442
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 77:
		//line parser.go.y:447
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 78:
		//line parser.go.y:452
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 79:
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 80:
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 81:
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 82:
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 83:
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 84:
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 85:
		//line parser.go.y:487
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 86:
		//line parser.go.y:492
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 87:
		//line parser.go.y:497
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 88:
		//line parser.go.y:502
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 89:
		//line parser.go.y:507
		{
			yyVAL.expr = &ast.LetsExpr{Lhss: yyS[yypt-2].expr_lhss, Operator: "=", Rhss: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 90:
		//line parser.go.y:512
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "+=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 91:
		//line parser.go.y:517
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "-=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 92:
		//line parser.go.y:522
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "*=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 93:
		//line parser.go.y:527
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "/=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 94:
		//line parser.go.y:532
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "&=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 95:
		//line parser.go.y:537
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "|=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 96:
		//line parser.go.y:542
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 97:
		//line parser.go.y:547
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 98:
		//line parser.go.y:552
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 99:
		//line parser.go.y:557
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 100:
		//line parser.go.y:562
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 101:
		//line parser.go.y:567
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 102:
		//line parser.go.y:572
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 103:
		//line parser.go.y:577
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyS[yypt-3].tok.lit}, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 104:
		//line parser.go.y:582
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 105:
		//line parser.go.y:587
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
