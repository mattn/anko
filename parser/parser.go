//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
	"unsafe"
)

//line parser.go.y:23
type yySymType struct {
	yys                    int
	stmt_var               ast.Stmt
	stmt_if                ast.Stmt
	stmt_for               ast.Stmt
	stmt_try_catch_finally ast.Stmt
	stmts                  []ast.Stmt
	stmt                   ast.Stmt
	teim                   ast.Expr
	expr                   ast.Expr
	tok                    Token
	idents                 []string
	exprs                  []ast.Expr
	pair                   ast.Expr
	pairs                  []ast.Expr
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

//line parser.go.y:447

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 29,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	44, 29,
	60, 29,
	-2, 1,
	-1, 3,
	44, 29,
	60, 29,
	-2, 1,
	-1, 4,
	44, 29,
	60, 29,
	-2, 1,
	-1, 5,
	44, 29,
	60, 29,
	-2, 1,
	-1, 6,
	44, 29,
	60, 29,
	-2, 1,
	-1, 7,
	44, 29,
	60, 29,
	-2, 1,
	-1, 8,
	44, 29,
	60, 29,
	-2, 1,
	-1, 10,
	44, 29,
	60, 29,
	-2, 32,
	-1, 18,
	44, 30,
	60, 30,
	-2, 37,
	-1, 26,
	64, 32,
	-2, 29,
	-1, 33,
	44, 29,
	60, 29,
	-2, 1,
	-1, 62,
	59, 32,
	-2, 29,
	-1, 65,
	44, 30,
	-2, 37,
	-1, 72,
	57, 1,
	-2, 29,
	-1, 74,
	57, 1,
	-2, 29,
	-1, 83,
	59, 32,
	-2, 29,
	-1, 96,
	44, 29,
	60, 29,
	-2, 32,
	-1, 99,
	57, 1,
	-2, 29,
	-1, 111,
	17, 0,
	18, 0,
	-2, 64,
	-1, 112,
	17, 0,
	18, 0,
	-2, 65,
	-1, 122,
	44, 29,
	60, 29,
	-2, 32,
	-1, 123,
	44, 29,
	60, 29,
	-2, 32,
	-1, 124,
	57, 1,
	-2, 29,
	-1, 125,
	44, 29,
	60, 29,
	-2, 32,
	-1, 146,
	59, 32,
	-2, 29,
	-1, 175,
	57, 1,
	-2, 29,
	-1, 176,
	57, 1,
	-2, 29,
	-1, 177,
	57, 1,
	-2, 29,
	-1, 179,
	57, 1,
	-2, 29,
	-1, 189,
	57, 1,
	-2, 29,
	-1, 191,
	57, 1,
	-2, 29,
	-1, 192,
	57, 1,
	-2, 29,
	-1, 194,
	57, 1,
	-2, 29,
	-1, 203,
	57, 1,
	-2, 29,
	-1, 211,
	57, 1,
	-2, 29,
	-1, 215,
	57, 1,
	-2, 29,
	-1, 220,
	57, 1,
	-2, 29,
}

const yyNprod = 85
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 987

var yyAct = []int{

	1, 91, 138, 32, 34, 35, 36, 37, 39, 40,
	181, 95, 31, 49, 164, 95, 143, 125, 71, 142,
	205, 44, 45, 46, 47, 48, 68, 49, 193, 96,
	62, 220, 149, 95, 97, 42, 9, 43, 58, 60,
	183, 180, 63, 163, 62, 95, 153, 64, 66, 42,
	89, 43, 58, 60, 146, 141, 84, 85, 86, 189,
	70, 188, 222, 64, 219, 216, 93, 213, 210, 87,
	72, 215, 73, 128, 208, 130, 207, 206, 100, 101,
	200, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 64,
	150, 139, 197, 211, 88, 121, 196, 126, 195, 174,
	129, 172, 131, 132, 133, 134, 135, 136, 162, 160,
	64, 203, 194, 192, 49, 156, 137, 98, 191, 179,
	177, 175, 124, 64, 46, 47, 48, 74, 144, 148,
	218, 62, 212, 127, 168, 178, 42, 182, 43, 58,
	60, 165, 92, 198, 166, 167, 147, 140, 102, 64,
	64, 94, 64, 69, 159, 154, 155, 67, 157, 90,
	99, 8, 7, 6, 5, 2, 185, 186, 187, 0,
	190, 169, 0, 64, 0, 0, 171, 0, 173, 170,
	199, 0, 201, 202, 0, 204, 0, 0, 52, 53,
	55, 57, 59, 61, 209, 0, 0, 0, 0, 0,
	0, 0, 214, 0, 0, 0, 217, 0, 0, 0,
	0, 221, 49, 50, 51, 0, 41, 0, 54, 56,
	44, 45, 46, 47, 48, 0, 0, 0, 0, 62,
	184, 0, 0, 0, 42, 0, 43, 58, 60, 52,
	53, 55, 57, 59, 61, 0, 0, 0, 0, 0,
	75, 76, 77, 78, 79, 80, 0, 0, 81, 82,
	0, 0, 0, 49, 50, 51, 0, 41, 0, 54,
	56, 44, 45, 46, 47, 48, 0, 83, 176, 123,
	62, 0, 0, 0, 0, 42, 0, 43, 58, 60,
	52, 53, 55, 57, 59, 61, 0, 0, 0, 0,
	75, 76, 77, 78, 79, 80, 0, 0, 81, 82,
	0, 0, 0, 0, 49, 50, 51, 0, 41, 0,
	54, 56, 44, 45, 46, 47, 48, 83, 0, 0,
	0, 62, 161, 0, 0, 0, 42, 0, 43, 58,
	60, 52, 53, 55, 57, 59, 61, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 49, 50, 51, 0, 41,
	0, 54, 56, 44, 45, 46, 47, 48, 0, 0,
	0, 0, 62, 158, 0, 0, 0, 42, 0, 43,
	58, 60, 52, 53, 55, 57, 59, 61, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 49, 50, 51, 0,
	41, 0, 54, 56, 44, 45, 46, 47, 48, 0,
	0, 0, 0, 62, 0, 0, 0, 0, 42, 152,
	43, 58, 60, 52, 53, 55, 57, 59, 61, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 49, 50, 51,
	0, 41, 151, 54, 56, 44, 45, 46, 47, 48,
	0, 0, 0, 0, 62, 0, 0, 0, 0, 42,
	0, 43, 58, 60, 52, 53, 55, 57, 59, 61,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 49, 50,
	51, 0, 41, 0, 54, 56, 44, 45, 46, 47,
	48, 0, 0, 0, 0, 62, 145, 0, 0, 0,
	42, 0, 43, 58, 60, 52, 53, 55, 57, 59,
	61, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 49,
	50, 51, 0, 41, 0, 54, 56, 44, 45, 46,
	47, 48, 0, 0, 0, 0, 62, 0, 122, 0,
	0, 42, 0, 43, 58, 60, 52, 53, 55, 57,
	59, 61, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	49, 50, 51, 0, 41, 0, 54, 56, 44, 45,
	46, 47, 48, 52, 53, 55, 57, 62, 61, 0,
	0, 0, 42, 0, 43, 58, 60, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 49, 50, 51,
	0, 0, 0, 54, 56, 44, 45, 46, 47, 48,
	0, 0, 0, 0, 62, 0, 0, 0, 0, 42,
	0, 43, 58, 60, 18, 17, 22, 0, 0, 27,
	10, 13, 11, 14, 38, 15, 0, 0, 0, 0,
	0, 0, 0, 30, 23, 24, 25, 12, 16, 0,
	0, 0, 0, 0, 0, 0, 0, 3, 4, 0,
	0, 0, 0, 0, 0, 0, 18, 17, 22, 0,
	19, 27, 10, 13, 11, 14, 28, 15, 29, 0,
	0, 20, 21, 26, 0, 30, 23, 24, 25, 12,
	16, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	4, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 19, 0, 0, 0, 0, 33, 28, 0,
	29, 0, 0, 20, 21, 26, 18, 17, 22, 0,
	0, 27, 10, 13, 11, 14, 0, 15, 0, 0,
	0, 0, 0, 0, 0, 30, 23, 24, 25, 12,
	16, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	4, 0, 52, 53, 55, 57, 0, 0, 0, 0,
	0, 0, 19, 0, 0, 0, 0, 0, 28, 0,
	29, 0, 0, 20, 21, 26, 49, 50, 51, 0,
	0, 0, 54, 56, 44, 45, 46, 47, 48, 55,
	57, 0, 0, 62, 0, 0, 0, 0, 42, 0,
	43, 58, 60, 0, 0, 18, 17, 22, 0, 0,
	27, 49, 50, 51, 0, 0, 0, 54, 56, 44,
	45, 46, 47, 48, 30, 23, 24, 25, 62, 0,
	0, 0, 0, 42, 0, 43, 58, 60, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 17, 22,
	0, 19, 27, 0, 0, 0, 0, 28, 0, 29,
	0, 0, 20, 21, 26, 0, 30, 23, 24, 25,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 19, 0, 0, 0, 0, 0, 28,
	0, 29, 0, 0, 20, 21, 26,
}
var yyPact = []int{

	792, -1000, 732, 792, 792, 792, 690, 792, 792, 589,
	923, 881, 163, 159, 2, 14, 81, -1000, 279, 881,
	881, 881, -1000, -1000, -1000, -1000, 923, 46, 146, 881,
	157, -15, -1000, 792, -1000, -1000, -1000, -1000, 114, -1000,
	-1000, 881, 881, 154, 881, 881, 881, 881, 881, 881,
	881, 881, 881, 881, 881, 881, 881, 881, 881, 881,
	881, 881, 923, -1000, 538, 229, 589, 76, -27, -1000,
	881, 127, 792, 881, 792, 881, 881, 881, 881, 881,
	881, -1000, -1000, 923, -14, -14, -14, -62, 153, -3,
	-41, -1000, 92, 487, -4, 152, 923, -1000, -26, 792,
	436, 385, -1000, 83, 83, -14, -14, -14, 589, -28,
	-28, 850, 850, -28, -28, -28, -28, 589, 626, 589,
	815, -13, 923, 923, 792, 923, 334, 881, 62, 283,
	61, 589, 589, 589, 589, 589, 589, -16, -1000, -45,
	143, 151, 146, -1000, 881, -1000, 923, -1000, -1000, 881,
	54, 881, -1000, -1000, -1000, -1000, 52, -1000, 75, 232,
	-1000, 74, 116, -1000, 73, -18, -49, 139, -1000, 589,
	-19, 181, -1000, 589, -1000, 792, 792, 792, 3, 792,
	72, 67, -31, -1000, 66, 51, 49, 45, 149, 792,
	23, 792, 792, 65, 792, -1000, -1000, -1000, -39, 20,
	-1000, 19, 17, 792, 11, 47, 112, -1000, -1000, 10,
	-1000, 792, 15, -1000, 8, 792, 110, 7, -25, -1000,
	792, 5, -1000,
}
var yyPgo = []int{

	0, 0, 175, 174, 173, 172, 171, 36, 42, 1,
	169, 12,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 3, 4, 4, 4, 5, 5,
	5, 6, 6, 6, 6, 9, 10, 10, 10, 11,
	11, 11, 8, 8, 8, 8, 7, 7, 7, 7,
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
	1, 3, 0, 1, 3, 3, 1, 1, 2, 2,
	2, 1, 1, 1, 1, 5, 4, 3, 7, 8,
	8, 9, 3, 3, 3, 5, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 3,
	3, 3, 3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, 37, 38, -3, -4, -5, -6, -7,
	10, 12, 27, 11, 13, 15, 28, 5, 4, 50,
	61, 62, 6, 24, 25, 26, 63, 9, 56, 58,
	23, -11, -1, 55, -1, -1, -1, -1, 14, -1,
	-1, 45, 63, 65, 49, 50, 51, 52, 53, 41,
	42, 43, 17, 18, 47, 19, 48, 20, 66, 21,
	67, 22, 58, -8, -7, 4, -7, 4, -11, 4,
	58, 4, 56, 58, 56, 31, 32, 33, 34, 35,
	36, 39, 40, 58, -7, -7, -7, -8, 58, 4,
	-10, -9, 6, -7, 4, 60, 44, -1, 13, 56,
	-7, -7, 4, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -8, 60, 60, 56, 44, -7, 16, -1, -7,
	-1, -7, -7, -7, -7, -7, -7, -8, 64, -11,
	4, 58, 60, 57, 46, 59, 58, 4, -8, 58,
	-1, 46, 64, 59, -8, -8, -1, -8, 59, -7,
	57, 59, 57, 59, 59, 8, -11, 4, -9, -7,
	-8, -7, 57, -7, 57, 56, 56, 56, 29, 56,
	59, 59, 8, 59, 59, -1, -1, -1, 58, 56,
	-1, 56, 56, 59, 56, 57, 57, 57, 4, -1,
	57, -1, -1, 56, -1, 59, 57, 57, 57, -1,
	57, 56, 30, 57, -1, 56, 57, -1, 30, 57,
	56, -1, 57,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, 10,
	-2, 29, 0, 29, 0, 0, 0, 36, -2, 29,
	29, 29, 41, 42, 43, 44, -2, 0, 26, 29,
	0, 0, 2, -2, 4, 5, 6, 7, 0, 8,
	9, 29, 29, 0, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, -2, 11, 33, -2, 12, 0, 0, 30,
	29, 0, -2, 29, -2, 29, 29, 29, 29, 29,
	29, 77, 78, -2, 38, 39, 40, 0, 29, 0,
	0, 27, 0, 0, 0, 0, -2, 3, 0, -2,
	0, 0, 53, 56, 57, 58, 59, 60, 61, 62,
	63, -2, -2, 66, 67, 68, 69, 79, 80, 81,
	82, 0, -2, -2, -2, -2, 0, 29, 0, 0,
	0, 71, 72, 73, 74, 75, 76, 0, 47, 0,
	30, 29, 0, 52, 29, 54, -2, 31, 70, 29,
	0, 29, 46, 84, 34, 35, 0, 14, 0, 0,
	19, 0, 0, 83, 0, 0, 0, 30, 28, 25,
	0, 0, 16, 45, 13, -2, -2, -2, 0, -2,
	0, 0, 0, 55, 0, 0, 0, 0, 0, -2,
	0, -2, -2, 0, -2, 17, 18, 20, 0, 0,
	48, 0, 0, -2, 0, 0, 24, 49, 50, 0,
	15, -2, 0, 51, 0, -2, 23, 0, 0, 22,
	-2, 0, 21,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 61, 3, 3, 3, 53, 67, 3,
	58, 59, 51, 49, 60, 50, 65, 52, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 46, 55,
	48, 44, 47, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 63, 3, 64, 62, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 56, 66, 57,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 54,
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
		//line parser.go.y:55
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:62
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:69
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:76
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.BreakStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:83
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.ContinueStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:90
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:97
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:104
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 9:
		//line parser.go.y:111
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 10:
		//line parser.go.y:119
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:126
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				if len(yyS[yypt-0].exprs) > 0 {
					yyS[yypt-0].exprs[0].SetPos(l.pos)
				}
			}
		}
	case 12:
		//line parser.go.y:135
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 13:
		//line parser.go.y:142
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:147
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 15:
		//line parser.go.y:152
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 16:
		//line parser.go.y:156
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 17:
		//line parser.go.y:164
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:169
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:173
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:177
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:182
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:186
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:190
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:194
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:199
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:204
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 27:
		//line parser.go.y:208
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 28:
		//line parser.go.y:212
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 29:
		//line parser.go.y:217
		{
			yyVAL.idents = []string{}
		}
	case 30:
		//line parser.go.y:221
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 31:
		//line parser.go.y:225
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 32:
		//line parser.go.y:230
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 33:
		//line parser.go.y:234
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 34:
		//line parser.go.y:238
		{
			yyVAL.exprs = append([]ast.Expr{yyS[yypt-2].expr}, yyS[yypt-0].exprs...)
		}
	case 35:
		//line parser.go.y:242
		{
			yyVAL.exprs = append([]ast.Expr{&ast.IdentExpr{Lit: yyS[yypt-2].tok.lit}}, yyS[yypt-0].exprs...)
		}
	case 36:
		//line parser.go.y:247
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 37:
		//line parser.go.y:251
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 38:
		//line parser.go.y:255
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
		}
	case 39:
		//line parser.go.y:259
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
		}
	case 40:
		//line parser.go.y:263
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
		}
	case 41:
		//line parser.go.y:267
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 42:
		//line parser.go.y:271
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 43:
		//line parser.go.y:275
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 44:
		//line parser.go.y:279
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
		}
	case 45:
		//line parser.go.y:283
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	case 46:
		//line parser.go.y:287
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 47:
		//line parser.go.y:291
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 48:
		//line parser.go.y:295
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 49:
		//line parser.go.y:299
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 50:
		//line parser.go.y:303
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 51:
		//line parser.go.y:307
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 52:
		//line parser.go.y:311
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 53:
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 54:
		//line parser.go.y:323
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 55:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 56:
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:335
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:351
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:355
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
		}
	case 63:
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
		}
	case 64:
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 65:
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 66:
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 67:
		//line parser.go.y:375
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 68:
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 69:
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 70:
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.LetExpr{Names: yyS[yypt-2].idents, Operator: "=", Exprs: yyS[yypt-0].exprs}
		}
	case 71:
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "+", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 72:
		//line parser.go.y:395
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "-", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 73:
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "*", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 74:
		//line parser.go.y:403
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "/", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 75:
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "&", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 76:
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.LetExpr{Names: []string{yyS[yypt-2].tok.lit}, Operator: "|", Exprs: []ast.Expr{yyS[yypt-0].expr}}
		}
	case 77:
		//line parser.go.y:415
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
		}
	case 78:
		//line parser.go.y:419
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
		}
	case 79:
		//line parser.go.y:423
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 80:
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 81:
		//line parser.go.y:431
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 82:
		//line parser.go.y:435
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 83:
		//line parser.go.y:439
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 84:
		//line parser.go.y:443
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	}
	goto yystack /* stack new state and value */
}
