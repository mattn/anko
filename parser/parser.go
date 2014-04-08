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

//line parser.go.y:586

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
	-1, 30,
	66, 48,
	-2, 45,
	-1, 38,
	47, 45,
	-2, 1,
	-1, 68,
	52, 48,
	-2, 45,
	-1, 71,
	47, 42,
	-2, 53,
	-1, 78,
	63, 1,
	-2, 45,
	-1, 80,
	63, 1,
	-2, 45,
	-1, 90,
	52, 48,
	-2, 45,
	-1, 94,
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
	53, 53,
	54, 53,
	55, 53,
	56, 53,
	57, 53,
	58, 53,
	59, 53,
	64, 53,
	65, 53,
	69, 53,
	70, 53,
	-2, 57,
	-1, 96,
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
	53, 53,
	54, 53,
	55, 53,
	56, 53,
	57, 53,
	58, 53,
	59, 53,
	64, 53,
	65, 53,
	69, 53,
	70, 53,
	-2, 59,
	-1, 106,
	47, 45,
	-2, 48,
	-1, 110,
	63, 1,
	-2, 45,
	-1, 111,
	47, 43,
	50, 43,
	-2, 67,
	-1, 122,
	17, 0,
	18, 0,
	51, 0,
	-2, 84,
	-1, 123,
	17, 0,
	18, 0,
	51, 0,
	-2, 85,
	-1, 133,
	47, 45,
	-2, 48,
	-1, 134,
	47, 45,
	-2, 48,
	-1, 135,
	63, 1,
	-2, 45,
	-1, 136,
	47, 45,
	-2, 48,
	-1, 161,
	52, 48,
	-2, 45,
	-1, 167,
	47, 44,
	50, 44,
	-2, 66,
	-1, 187,
	47, 43,
	50, 43,
	-2, 58,
	-1, 188,
	47, 43,
	50, 43,
	-2, 60,
	-1, 200,
	63, 1,
	-2, 45,
	-1, 201,
	63, 1,
	-2, 45,
	-1, 202,
	63, 1,
	-2, 45,
	-1, 209,
	47, 45,
	-2, 1,
	-1, 210,
	63, 1,
	-2, 45,
	-1, 221,
	63, 1,
	-2, 45,
	-1, 222,
	47, 45,
	-2, 1,
	-1, 225,
	63, 1,
	-2, 45,
	-1, 226,
	63, 1,
	-2, 45,
	-1, 228,
	63, 1,
	-2, 45,
	-1, 239,
	63, 1,
	-2, 45,
	-1, 248,
	63, 1,
	-2, 45,
	-1, 249,
	63, 1,
	-2, 45,
	-1, 254,
	63, 1,
	-2, 45,
	-1, 260,
	63, 1,
	-2, 45,
}

const yyNprod = 105
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1465

var yyAct = []int{

	1, 153, 55, 37, 39, 40, 41, 42, 44, 45,
	46, 182, 183, 77, 74, 102, 157, 35, 52, 53,
	54, 262, 259, 256, 55, 47, 48, 255, 251, 158,
	66, 64, 247, 69, 245, 20, 19, 26, 244, 108,
	31, 11, 14, 12, 15, 43, 16, 47, 48, 184,
	185, 243, 66, 64, 34, 27, 28, 29, 13, 17,
	79, 236, 231, 230, 98, 229, 220, 207, 3, 4,
	199, 78, 242, 197, 180, 18, 177, 221, 109, 140,
	260, 142, 33, 254, 249, 248, 239, 21, 25, 55,
	228, 226, 225, 32, 210, 202, 30, 10, 22, 23,
	24, 200, 132, 50, 51, 52, 53, 54, 135, 70,
	72, 166, 47, 48, 154, 80, 81, 66, 64, 91,
	92, 93, 95, 97, 150, 163, 227, 110, 70, 214,
	137, 104, 212, 137, 211, 189, 172, 186, 169, 100,
	162, 165, 161, 156, 76, 107, 112, 113, 114, 115,
	116, 117, 118, 119, 120, 121, 122, 123, 124, 125,
	126, 127, 128, 129, 130, 131, 70, 170, 171, 136,
	173, 191, 137, 193, 138, 209, 159, 141, 184, 185,
	144, 145, 146, 147, 148, 149, 99, 106, 70, 258,
	250, 204, 139, 205, 206, 195, 213, 190, 103, 233,
	192, 216, 217, 218, 70, 164, 188, 187, 174, 155,
	223, 224, 111, 105, 75, 73, 82, 83, 84, 85,
	86, 87, 234, 235, 88, 89, 237, 238, 101, 240,
	36, 70, 70, 181, 70, 134, 90, 176, 9, 8,
	246, 7, 6, 82, 83, 84, 85, 86, 87, 252,
	253, 88, 89, 5, 2, 257, 0, 194, 0, 70,
	0, 261, 0, 196, 0, 0, 198, 20, 19, 26,
	0, 0, 31, 11, 14, 12, 15, 203, 16, 0,
	0, 0, 208, 0, 0, 0, 34, 27, 28, 29,
	13, 17, 0, 0, 0, 0, 0, 0, 0, 0,
	3, 4, 0, 0, 0, 0, 0, 18, 0, 0,
	0, 0, 0, 0, 33, 0, 0, 232, 0, 21,
	25, 0, 0, 0, 38, 32, 0, 0, 30, 0,
	22, 23, 24, 20, 19, 26, 0, 0, 31, 11,
	14, 12, 15, 0, 16, 0, 0, 0, 0, 0,
	0, 0, 34, 27, 28, 29, 13, 17, 0, 0,
	0, 0, 0, 0, 0, 0, 3, 4, 0, 0,
	0, 0, 0, 18, 0, 0, 0, 0, 0, 0,
	33, 0, 0, 0, 0, 21, 25, 0, 0, 0,
	0, 32, 0, 0, 30, 0, 22, 23, 24, 58,
	59, 61, 63, 65, 67, 0, 82, 83, 84, 85,
	86, 87, 0, 0, 88, 89, 0, 0, 0, 0,
	0, 0, 0, 55, 56, 57, 90, 0, 0, 0,
	49, 0, 0, 68, 178, 60, 62, 50, 51, 52,
	53, 54, 0, 179, 0, 0, 47, 48, 0, 0,
	0, 66, 64, 58, 59, 61, 63, 65, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 55, 56, 57,
	0, 0, 0, 0, 49, 0, 0, 68, 241, 60,
	62, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	47, 48, 0, 0, 0, 66, 64, 58, 59, 61,
	63, 65, 67, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 55, 56, 57, 0, 0, 0, 0, 49, 222,
	0, 68, 0, 60, 62, 50, 51, 52, 53, 54,
	0, 0, 0, 0, 47, 48, 0, 0, 0, 66,
	64, 58, 59, 61, 63, 65, 67, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 55, 56, 57, 0, 0,
	0, 0, 49, 0, 0, 68, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 219, 0, 0, 47, 48,
	0, 0, 0, 66, 64, 58, 59, 61, 63, 65,
	67, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 55,
	56, 57, 0, 0, 0, 0, 49, 0, 0, 68,
	215, 60, 62, 50, 51, 52, 53, 54, 0, 0,
	0, 0, 47, 48, 0, 0, 0, 66, 64, 58,
	59, 61, 63, 65, 67, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 55, 56, 57, 0, 0, 0, 0,
	49, 0, 0, 68, 0, 60, 62, 50, 51, 52,
	53, 54, 0, 0, 201, 0, 47, 48, 0, 0,
	0, 66, 64, 58, 59, 61, 63, 65, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 55, 56, 57,
	0, 0, 0, 0, 49, 0, 0, 68, 175, 60,
	62, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	47, 48, 0, 0, 0, 66, 64, 58, 59, 61,
	63, 65, 67, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 55, 56, 57, 0, 0, 0, 0, 49, 168,
	0, 68, 0, 60, 62, 50, 51, 52, 53, 54,
	0, 0, 0, 0, 47, 48, 0, 0, 0, 66,
	64, 58, 59, 61, 63, 65, 67, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 55, 56, 57, 0, 0,
	0, 0, 49, 0, 0, 68, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 0, 0, 0, 47, 48,
	167, 0, 0, 66, 64, 58, 59, 61, 63, 65,
	67, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 55,
	56, 57, 0, 0, 0, 0, 49, 0, 0, 68,
	160, 60, 62, 50, 51, 52, 53, 54, 0, 0,
	0, 0, 47, 48, 0, 0, 0, 66, 64, 58,
	59, 61, 63, 65, 67, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 55, 56, 57, 0, 0, 0, 0,
	49, 0, 0, 68, 0, 60, 62, 50, 51, 52,
	53, 54, 0, 0, 143, 0, 47, 48, 0, 0,
	0, 66, 64, 58, 59, 61, 63, 65, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 55, 56, 57,
	0, 0, 0, 0, 49, 0, 133, 68, 0, 60,
	62, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	47, 48, 0, 0, 0, 66, 64, 58, 59, 61,
	63, 65, 67, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 55, 56, 57, 0, 0, 0, 0, 49, 0,
	0, 68, 0, 60, 62, 50, 51, 52, 53, 54,
	0, 0, 0, 0, 47, 48, 0, 0, 0, 66,
	64, 58, 59, 61, 63, 65, 67, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 55, 56, 57, 0, 0,
	0, 0, 49, 0, 0, 68, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 0, 0, 0, 152, 48,
	0, 0, 0, 66, 64, 58, 59, 61, 63, 65,
	67, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 55,
	56, 57, 0, 0, 0, 0, 49, 0, 0, 68,
	0, 60, 62, 50, 51, 52, 53, 54, 0, 0,
	0, 0, 151, 48, 0, 0, 0, 66, 64, 58,
	59, 61, 63, 0, 67, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 58,
	59, 61, 63, 55, 56, 57, 0, 0, 0, 0,
	0, 0, 0, 68, 0, 60, 62, 50, 51, 52,
	53, 54, 0, 55, 56, 57, 47, 48, 0, 0,
	0, 66, 64, 68, 0, 60, 62, 50, 51, 52,
	53, 54, 61, 63, 0, 0, 47, 48, 20, 19,
	26, 66, 64, 31, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 56, 57, 34, 27, 28,
	29, 0, 0, 0, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 0, 47, 48, 0,
	0, 0, 66, 64, 0, 33, 0, 71, 19, 26,
	21, 25, 31, 0, 0, 0, 32, 0, 0, 30,
	0, 22, 23, 24, 0, 0, 34, 27, 28, 29,
	96, 19, 26, 0, 0, 31, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 34,
	27, 28, 29, 0, 33, 0, 0, 0, 0, 21,
	25, 0, 0, 0, 0, 32, 0, 0, 30, 0,
	22, 23, 24, 0, 0, 0, 0, 33, 0, 94,
	19, 26, 21, 25, 31, 0, 0, 0, 32, 0,
	0, 30, 0, 22, 23, 24, 0, 0, 34, 27,
	28, 29, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 33, 0, 0, 0,
	0, 21, 25, 0, 0, 0, 0, 32, 0, 0,
	30, 0, 22, 23, 24,
}
var yyPact = []int{

	329, -1000, 263, 329, 329, 329, 31, 329, 329, 329,
	1030, 1323, 1274, 211, 210, 93, 9, 53, 1274, -1000,
	375, 1274, 1274, 1274, 1395, 1346, -1000, -1000, -1000, -1000,
	1323, 135, 192, 1274, 209, 140, 95, -1000, 329, -1000,
	-1000, -1000, -1000, 65, -1000, -1000, -1000, 208, 1274, 1274,
	1274, 1274, 1274, 1274, 1274, 1274, 1274, 1274, 1274, 1274,
	1274, 1274, 1274, 1274, 1274, 1274, 1274, 1274, 1323, -1000,
	976, 185, 1030, 46, 122, -1000, 1274, 176, 329, 1274,
	329, 922, 1274, 1274, 1274, 1274, 1274, 1274, -1000, -1000,
	1323, -17, -17, -17, 212, 1138, 212, 1084, -65, 205,
	92, -34, -1000, 127, 868, 91, 1323, 1274, -1000, 90,
	329, -1000, 814, 760, -39, -39, -17, -17, -17, 1030,
	48, 48, 1253, 1253, 48, 48, 48, 48, 1030, 1192,
	1030, 1212, 86, 1323, 1323, 329, 1323, 204, 706, 1274,
	13, 382, 11, 133, 1030, 1030, 1030, 1030, 1030, 1030,
	85, 203, 202, -1000, 83, 189, 196, 192, -1000, 1274,
	-1000, 1323, -1000, -1000, 1030, 1274, 10, -1000, 1274, -1000,
	-1000, -1000, 7, -1000, -1000, 39, 652, -1000, 33, 1274,
	162, 4, -1000, -1000, 1274, 126, -1000, -1000, -1000, 32,
	82, 80, 188, -1000, 1030, 77, 598, -1000, 1030, -1000,
	329, 329, 329, 544, 15, -1000, -1000, -1000, 490, 329,
	329, 30, 29, 74, -1000, 28, 2, 0, -1, 1274,
	195, 329, 329, -1000, -2, 329, 329, 24, 329, -1000,
	-1000, -1000, 436, 20, -12, -1000, -1000, -25, -29, 329,
	-31, 23, 22, 160, -1000, -1000, -35, -1000, 329, 329,
	21, -1000, -36, -40, 329, -1000, 159, -41, 18, -1000,
	329, -42, -1000,
}
var yyPgo = []int{

	0, 0, 254, 253, 242, 241, 239, 238, 12, 11,
	233, 97, 33, 230, 17, 15, 228, 14,
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
	11, 11, 11, 11, 11,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	2, 1, 2, 2, 5, 4, 9, 5, 7, 7,
	4, 7, 11, 15, 12, 11, 8, 0, 1, 2,
	1, 2, 4, 3, 5, 3, 0, 1, 3, 0,
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
	-11, 10, 12, 27, 11, 13, 15, 28, 44, 5,
	4, 56, 67, 68, 69, 57, 6, 24, 25, 26,
	65, 9, 62, 51, 23, -14, -13, -1, 61, -1,
	-1, -1, -1, 14, -1, -1, -1, 64, 65, 48,
	55, 56, 57, 58, 59, 41, 42, 43, 17, 18,
	53, 19, 54, 20, 70, 21, 69, 22, 51, -12,
	-11, 4, -11, 4, -17, 4, 51, 4, 62, 51,
	62, -11, 31, 32, 33, 34, 35, 36, 39, 40,
	51, -11, -11, -11, 4, -11, 4, -11, -12, 51,
	4, -16, -15, 6, -11, 4, 47, 50, -1, 13,
	62, 4, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -12, 50, 50, 62, 47, 50, -11, 16,
	-1, -11, -1, 62, -11, -11, -11, -11, -11, -11,
	-12, 64, 64, 66, -17, 4, 51, 50, 63, 49,
	52, 51, -12, -14, -11, 51, -1, 66, 49, 52,
	-12, -12, -1, -12, 4, 52, -11, 63, 52, 61,
	63, -10, -9, -8, 45, 46, 52, 4, 4, 52,
	8, -17, 4, -15, -11, -12, -11, 63, -11, 63,
	62, 62, 62, -11, 29, -9, -8, 63, -11, 49,
	62, 52, 52, 8, 52, 52, -1, -1, -1, 61,
	51, 62, 49, -1, -1, 62, 62, 52, 62, 63,
	63, 63, -11, 4, -1, -1, 63, -1, -1, 62,
	-1, 52, 52, 63, 63, 63, -1, 63, 62, 62,
	30, 63, -1, -1, 62, 63, 63, -1, 30, 63,
	62, -1, 63,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	11, -2, 45, 0, 39, 0, 0, 0, 45, 52,
	-2, 45, 45, 45, 45, 45, 61, 62, 63, 64,
	-2, 0, 36, 45, 0, 0, 46, 2, -2, 4,
	5, 6, 7, 0, 8, 9, 10, 0, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, -2, 12,
	49, -2, 13, 0, 0, 40, 45, 0, -2, 45,
	-2, 0, 45, 45, 45, 45, 45, 45, 97, 98,
	-2, 54, 55, 56, -2, 0, -2, 0, 0, 39,
	0, 0, 37, 0, 0, 0, -2, 45, 3, 0,
	-2, -2, 0, 0, 76, 77, 78, 79, 80, 81,
	82, 83, -2, -2, 86, 87, 88, 89, 99, 100,
	101, 102, 0, -2, -2, -2, -2, 0, 0, 45,
	0, 0, 0, 27, 91, 92, 93, 94, 95, 96,
	0, 0, 0, 68, 0, 40, 39, 0, 73, 45,
	74, -2, 90, 47, 0, 45, 0, -2, 45, 104,
	50, 51, 0, 15, 41, 0, 0, 20, 0, 45,
	0, 0, 28, 30, 45, 0, 103, -2, -2, 0,
	0, 0, 40, 38, 35, 0, 0, 17, 65, 14,
	-2, -2, -2, 0, 0, 29, 31, 34, 0, -2,
	-2, 0, 0, 0, 75, 0, 0, 0, 0, 45,
	0, -2, -2, 33, 0, -2, -2, 0, -2, 18,
	19, 21, 0, 0, 0, 32, 69, 0, 0, -2,
	0, 0, 0, 26, 70, 71, 0, 16, -2, -2,
	0, 72, 0, 0, -2, 22, 25, 0, 0, 24,
	-2, 0, 23,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 67, 3, 3, 3, 59, 69, 3,
	51, 52, 57, 55, 50, 56, 64, 58, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 61,
	54, 47, 53, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 65, 3, 66, 68, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 62, 70, 63,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 60,
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
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 17:
		//line parser.go.y:152
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 18:
		//line parser.go.y:160
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
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
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:177
		{
			yyVAL.stmt_for = &ast.CForStmt{Expr1: yyS[yypt-8].expr, Expr2: yyS[yypt-6].expr, Expr3: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:182
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:186
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:190
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
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
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:372
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 64:
		//line parser.go.y:377
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
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
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 67:
		//line parser.go.y:392
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 68:
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 69:
		//line parser.go.y:402
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 70:
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 71:
		//line parser.go.y:412
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].expr_idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 72:
		//line parser.go.y:417
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 73:
		//line parser.go.y:422
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
	case 74:
		//line parser.go.y:431
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 75:
		//line parser.go.y:436
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 76:
		//line parser.go.y:441
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 77:
		//line parser.go.y:446
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 78:
		//line parser.go.y:451
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 79:
		//line parser.go.y:456
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 80:
		//line parser.go.y:461
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 81:
		//line parser.go.y:466
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 82:
		//line parser.go.y:471
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 83:
		//line parser.go.y:476
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 84:
		//line parser.go.y:481
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 85:
		//line parser.go.y:486
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 86:
		//line parser.go.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 87:
		//line parser.go.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 88:
		//line parser.go.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 89:
		//line parser.go.y:506
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 90:
		//line parser.go.y:511
		{
			yyVAL.expr = &ast.LetsExpr{Lhss: yyS[yypt-2].expr_lhss, Operator: "=", Rhss: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 91:
		//line parser.go.y:516
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "+=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 92:
		//line parser.go.y:521
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "-=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 93:
		//line parser.go.y:526
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "*=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 94:
		//line parser.go.y:531
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "/=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 95:
		//line parser.go.y:536
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "&=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 96:
		//line parser.go.y:541
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "|=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 97:
		//line parser.go.y:546
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 98:
		//line parser.go.y:551
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 99:
		//line parser.go.y:556
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 100:
		//line parser.go.y:561
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 101:
		//line parser.go.y:566
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 102:
		//line parser.go.y:571
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 103:
		//line parser.go.y:576
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 104:
		//line parser.go.y:581
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
