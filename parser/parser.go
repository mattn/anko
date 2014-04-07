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
	" (",
	" )",
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

//line parser.go.y:522

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
	49, 32,
	-2, 42,
	-1, 26,
	44, 35,
	-2, 38,
	-1, 34,
	44, 35,
	-2, 1,
	-1, 63,
	44, 35,
	-2, 38,
	-1, 72,
	60, 1,
	-2, 35,
	-1, 74,
	60, 1,
	-2, 35,
	-1, 83,
	44, 35,
	-2, 38,
	-1, 95,
	44, 35,
	-2, 38,
	-1, 99,
	60, 1,
	-2, 35,
	-1, 100,
	44, 33,
	49, 33,
	-2, 52,
	-1, 111,
	17, 0,
	18, 0,
	47, 0,
	-2, 69,
	-1, 112,
	17, 0,
	18, 0,
	47, 0,
	-2, 70,
	-1, 123,
	60, 1,
	-2, 35,
	-1, 124,
	44, 35,
	-2, 38,
	-1, 146,
	44, 35,
	-2, 38,
	-1, 152,
	44, 34,
	49, 34,
	-2, 51,
	-1, 155,
	17, 0,
	18, 0,
	47, 0,
	-2, 40,
	-1, 176,
	60, 1,
	-2, 35,
	-1, 177,
	60, 1,
	-2, 35,
	-1, 178,
	60, 1,
	-2, 35,
	-1, 180,
	60, 1,
	-2, 35,
	-1, 190,
	60, 1,
	-2, 35,
	-1, 192,
	60, 1,
	-2, 35,
	-1, 193,
	60, 1,
	-2, 35,
	-1, 195,
	60, 1,
	-2, 35,
	-1, 204,
	60, 1,
	-2, 35,
	-1, 212,
	60, 1,
	-2, 35,
	-1, 216,
	60, 1,
	-2, 35,
	-1, 221,
	60, 1,
	-2, 35,
}

const yyNprod = 90
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 864

var yyAct = []int{

	1, 64, 68, 33, 35, 36, 37, 38, 40, 41,
	122, 50, 91, 223, 53, 54, 56, 58, 60, 62,
	220, 221, 142, 217, 138, 31, 214, 71, 87, 211,
	216, 42, 43, 143, 209, 97, 59, 61, 50, 51,
	52, 50, 44, 208, 63, 185, 207, 55, 57, 45,
	46, 47, 48, 49, 47, 48, 49, 9, 42, 43,
	212, 42, 43, 59, 61, 121, 59, 61, 65, 66,
	73, 201, 198, 128, 197, 130, 196, 84, 85, 86,
	189, 175, 72, 98, 65, 137, 173, 93, 163, 161,
	204, 139, 190, 195, 193, 192, 180, 147, 178, 176,
	151, 101, 102, 103, 104, 105, 106, 107, 108, 109,
	110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
	120, 65, 148, 123, 156, 74, 157, 124, 126, 99,
	122, 129, 125, 131, 132, 133, 134, 135, 136, 184,
	122, 65, 182, 125, 167, 50, 165, 125, 171, 164,
	122, 154, 122, 65, 149, 169, 45, 46, 47, 48,
	49, 96, 206, 194, 181, 42, 43, 89, 150, 146,
	59, 61, 141, 70, 144, 95, 219, 186, 187, 188,
	155, 191, 65, 213, 179, 160, 127, 183, 166, 92,
	199, 200, 168, 202, 203, 158, 205, 140, 100, 94,
	69, 67, 170, 90, 65, 210, 32, 8, 172, 7,
	88, 174, 6, 215, 5, 2, 0, 218, 0, 0,
	0, 0, 222, 53, 54, 56, 58, 60, 62, 75,
	76, 77, 78, 79, 80, 0, 0, 81, 82, 0,
	0, 0, 0, 0, 0, 83, 0, 50, 51, 52,
	0, 44, 0, 63, 0, 0, 55, 57, 45, 46,
	47, 48, 49, 0, 0, 177, 0, 42, 43, 0,
	0, 0, 59, 61, 53, 54, 56, 58, 60, 62,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 50, 51,
	52, 0, 44, 0, 63, 162, 0, 55, 57, 45,
	46, 47, 48, 49, 0, 0, 0, 0, 42, 43,
	0, 0, 0, 59, 61, 53, 54, 56, 58, 60,
	62, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 50,
	51, 52, 0, 44, 0, 63, 159, 0, 55, 57,
	45, 46, 47, 48, 49, 0, 0, 0, 0, 42,
	43, 0, 0, 0, 59, 61, 53, 54, 56, 58,
	60, 62, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	50, 51, 52, 0, 44, 153, 63, 0, 0, 55,
	57, 45, 46, 47, 48, 49, 0, 0, 0, 0,
	42, 43, 0, 0, 0, 59, 61, 53, 54, 56,
	58, 60, 62, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 50, 51, 52, 0, 44, 0, 63, 0, 0,
	55, 57, 45, 46, 47, 48, 49, 0, 0, 0,
	0, 42, 43, 152, 0, 0, 59, 61, 53, 54,
	56, 58, 60, 62, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 50, 51, 52, 0, 44, 0, 63, 145,
	0, 55, 57, 45, 46, 47, 48, 49, 0, 0,
	0, 0, 42, 43, 0, 0, 0, 59, 61, 53,
	54, 56, 58, 60, 62, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 50, 51, 52, 0, 44, 0, 63,
	0, 0, 55, 57, 45, 46, 47, 48, 49, 56,
	58, 0, 0, 42, 43, 18, 17, 22, 59, 61,
	27, 10, 13, 11, 14, 39, 15, 0, 0, 0,
	0, 50, 51, 52, 30, 23, 24, 25, 12, 16,
	55, 57, 45, 46, 47, 48, 49, 0, 3, 4,
	0, 42, 43, 0, 0, 0, 59, 61, 29, 0,
	18, 17, 22, 0, 19, 27, 10, 13, 11, 14,
	28, 15, 0, 26, 0, 20, 21, 0, 0, 30,
	23, 24, 25, 12, 16, 0, 0, 0, 0, 0,
	0, 0, 0, 3, 4, 0, 0, 0, 0, 0,
	0, 0, 0, 29, 0, 0, 0, 0, 0, 19,
	0, 0, 0, 0, 34, 28, 0, 0, 26, 0,
	20, 21, 18, 17, 22, 0, 0, 27, 10, 13,
	11, 14, 0, 15, 0, 0, 0, 0, 0, 0,
	0, 30, 23, 24, 25, 12, 16, 0, 0, 0,
	0, 0, 0, 0, 0, 3, 4, 0, 0, 0,
	0, 53, 54, 56, 58, 29, 62, 0, 0, 0,
	0, 19, 0, 0, 0, 0, 0, 28, 0, 0,
	26, 0, 20, 21, 0, 50, 51, 52, 0, 0,
	0, 63, 0, 0, 55, 57, 45, 46, 47, 48,
	49, 53, 54, 56, 58, 42, 43, 0, 0, 0,
	59, 61, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 50, 51, 52, 0, 0,
	0, 63, 0, 0, 55, 57, 45, 46, 47, 48,
	49, 0, 18, 17, 22, 42, 43, 27, 0, 0,
	59, 61, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 30, 23, 24, 25, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 29, 0, 0, 0, 0,
	0, 19, 0, 0, 0, 0, 0, 28, 0, 0,
	26, 0, 20, 21,
}
var yyPact = []int{

	678, -1000, 616, 678, 678, 678, 571, 678, 678, 512,
	798, 798, 197, 196, 126, 23, 66, -1000, 198, 798,
	798, 798, -1000, -1000, -1000, -1000, 798, 163, 183, 798,
	195, 131, 112, -1000, 678, -1000, -1000, -1000, -1000, 70,
	-1000, -1000, 194, 798, 798, 798, 798, 798, 798, 798,
	798, 798, 798, 798, 798, 798, 798, 798, 798, 798,
	798, 798, 798, 798, 81, 512, 512, 64, 83, -1000,
	798, 170, 678, 798, 678, 798, 798, 798, 798, 798,
	798, -1000, -1000, 798, -30, -30, -30, -39, 193, 125,
	-27, -1000, 128, 461, 122, 798, 798, -1000, 121, 678,
	-1000, 410, 359, 0, 0, -30, -30, -30, 512, 104,
	104, 550, 550, 104, 104, 104, 104, 512, 704, 512,
	744, 103, 798, 678, 798, 191, 308, 798, 29, 257,
	28, 512, 512, 512, 512, 512, 512, 101, -1000, 98,
	180, 188, 183, -1000, 798, -1000, 798, 81, -1000, 512,
	798, 26, -1000, 798, -1000, 550, 21, 81, -1000, 40,
	206, -1000, 39, 155, -1000, 37, 116, 94, 179, -1000,
	512, 91, -3, -1000, 512, -1000, 678, 678, 678, 33,
	678, 36, 35, 115, -1000, 34, 16, 14, 12, 186,
	678, 11, 678, 678, 31, 678, -1000, -1000, -1000, 114,
	-14, -1000, -17, -26, 678, -31, 1, 153, -1000, -1000,
	-34, -1000, 678, -29, -1000, -37, 678, 146, -40, -38,
	-1000, 678, -47, -1000,
}
var yyPgo = []int{

	0, 0, 215, 214, 212, 209, 207, 57, 1, 206,
	25, 12, 203, 2,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 3, 4, 4, 4, 5, 5,
	5, 6, 6, 6, 6, 11, 12, 12, 12, 13,
	13, 13, 9, 9, 9, 10, 10, 10, 8, 8,
	8, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 5, 4, 9, 5, 7, 7, 4,
	7, 15, 12, 11, 8, 3, 0, 1, 3, 0,
	1, 3, 1, 3, 4, 0, 1, 3, 0, 1,
	3, 1, 1, 2, 2, 2, 1, 1, 1, 1,
	5, 4, 3, 3, 7, 8, 8, 9, 3, 3,
	5, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 3, 3, 3, 3, 4, 4,
}
var yyChk = []int{

	-1000, -1, -2, 37, 38, -3, -4, -5, -6, -7,
	10, 12, 27, 11, 13, 15, 28, 5, 4, 53,
	64, 65, 6, 24, 25, 26, 62, 9, 59, 47,
	23, -10, -9, -1, 58, -1, -1, -1, -1, 14,
	-1, -1, 61, 62, 45, 52, 53, 54, 55, 56,
	41, 42, 43, 17, 18, 50, 19, 51, 20, 66,
	21, 67, 22, 47, -8, -7, -7, 4, -13, 4,
	47, 4, 59, 47, 59, 31, 32, 33, 34, 35,
	36, 39, 40, 47, -7, -7, -7, -8, 47, 4,
	-12, -11, 6, -7, 4, 44, 49, -1, 13, 59,
	4, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -8, 49, 59, 44, 49, -7, 16, -1, -7,
	-1, -7, -7, -7, -7, -7, -7, -8, 63, -13,
	4, 47, 49, 60, 46, 48, 47, -8, -10, -7,
	47, -1, 63, 46, 48, -7, -1, -8, 4, 48,
	-7, 60, 48, 60, 48, 48, 8, -13, 4, -11,
	-7, -8, -7, 60, -7, 60, 59, 59, 59, 29,
	59, 48, 48, 8, 48, 48, -1, -1, -1, 47,
	59, -1, 59, 59, 48, 59, 60, 60, 60, 4,
	-1, 60, -1, -1, 59, -1, 48, 60, 60, 60,
	-1, 60, 59, 30, 60, -1, 59, 60, -1, 30,
	60, 59, -1, 60,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, -2, 10,
	-2, 35, 0, 29, 0, 0, 0, 41, -2, 35,
	35, 35, 46, 47, 48, 49, -2, 0, 26, 35,
	0, 0, 36, 2, -2, 4, 5, 6, 7, 0,
	8, 9, 0, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, -2, 11, 39, 12, 0, 0, 30,
	35, 0, -2, 35, -2, 35, 35, 35, 35, 35,
	35, 82, 83, -2, 43, 44, 45, 0, 29, 0,
	0, 27, 0, 0, 0, -2, 35, 3, 0, -2,
	-2, 0, 0, 61, 62, 63, 64, 65, 66, 67,
	68, -2, -2, 71, 72, 73, 74, 84, 85, 86,
	87, 0, 35, -2, -2, 0, 0, 35, 0, 0,
	0, 76, 77, 78, 79, 80, 81, 0, 53, 0,
	30, 29, 0, 58, 35, 59, -2, 75, 37, 0,
	35, 0, -2, 35, 89, -2, 0, 14, 31, 0,
	0, 19, 0, 0, 88, 0, 0, 0, 30, 28,
	25, 0, 0, 16, 50, 13, -2, -2, -2, 0,
	-2, 0, 0, 0, 60, 0, 0, 0, 0, 0,
	-2, 0, -2, -2, 0, -2, 17, 18, 20, 0,
	0, 54, 0, 0, -2, 0, 0, 24, 55, 56,
	0, 15, -2, 0, 57, 0, -2, 23, 0, 0,
	22, -2, 0, 21,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 64, 3, 3, 3, 56, 67, 3,
	47, 48, 54, 52, 49, 53, 61, 55, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 46, 58,
	51, 44, 50, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 62, 3, 63, 65, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 59, 66, 60,
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
		//line parser.go.y:59
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:66
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-1].stmt.SetPos(l.pos)
			}
		}
	case 3:
		//line parser.go.y:74
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
				yyS[yypt-2].stmt.SetPos(l.pos)
			}
		}
	case 4:
		//line parser.go.y:82
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.BreakStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:89
		{
			yyVAL.stmts = append([]ast.Stmt{&ast.ContinueStmt{}}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:96
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:103
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
		//line parser.go.y:117
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 10:
		//line parser.go.y:125
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 11:
		//line parser.go.y:129
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 12:
		//line parser.go.y:133
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
		}
	case 13:
		//line parser.go.y:137
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:142
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 15:
		//line parser.go.y:147
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 16:
		//line parser.go.y:151
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 17:
		//line parser.go.y:159
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:164
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:168
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmts: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:172
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyS[yypt-4].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:177
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:181
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:185
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:189
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:194
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:199
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 27:
		//line parser.go.y:203
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 28:
		//line parser.go.y:207
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 29:
		//line parser.go.y:212
		{
			yyVAL.idents = []string{}
		}
	case 30:
		//line parser.go.y:216
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 31:
		//line parser.go.y:220
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 32:
		//line parser.go.y:225
		{
			yyVAL.lhs = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.lhs.SetPos(l.pos)
			}
		}
	case 33:
		//line parser.go.y:230
		{
			yyVAL.lhs = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.lhs.SetPos(l.pos)
			}
		}
	case 34:
		//line parser.go.y:235
		{
			yyVAL.lhs = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.lhs.SetPos(l.pos)
			}
		}
	case 35:
		//line parser.go.y:241
		{
			yyVAL.lhss = []ast.Expr{}
		}
	case 36:
		//line parser.go.y:245
		{
			yyVAL.lhss = []ast.Expr{yyS[yypt-0].lhs}
		}
	case 37:
		//line parser.go.y:249
		{
			yyVAL.lhss = append([]ast.Expr{yyS[yypt-2].lhs}, yyS[yypt-0].lhss...)
		}
	case 38:
		//line parser.go.y:254
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 39:
		//line parser.go.y:258
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 40:
		//line parser.go.y:262
		{
			//$$ = append([]ast.Expr{$1}, $3...)
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 41:
		//line parser.go.y:268
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 42:
		//line parser.go.y:273
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 43:
		//line parser.go.y:278
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 44:
		//line parser.go.y:283
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 45:
		//line parser.go.y:288
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 46:
		//line parser.go.y:293
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 47:
		//line parser.go.y:298
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 48:
		//line parser.go.y:303
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 49:
		//line parser.go.y:308
		{
			yyVAL.expr = &ast.ConstExpr{Value: unsafe.Pointer(nil)}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 50:
		//line parser.go.y:313
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 51:
		//line parser.go.y:318
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 52:
		//line parser.go.y:323
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:328
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:333
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:338
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:348
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:353
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
	case 59:
		//line parser.go.y:362
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 60:
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 61:
		//line parser.go.y:372
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 62:
		//line parser.go.y:377
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:382
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 64:
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 65:
		//line parser.go.y:392
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 66:
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "**", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 67:
		//line parser.go.y:402
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 68:
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">>", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 69:
		//line parser.go.y:412
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 70:
		//line parser.go.y:417
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 71:
		//line parser.go.y:422
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 72:
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 73:
		//line parser.go.y:432
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 74:
		//line parser.go.y:437
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 75:
		//line parser.go.y:442
		{
			yyVAL.expr = &ast.LetsExpr{Lhss: yyS[yypt-2].lhss, Operator: "=", Rhss: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 76:
		//line parser.go.y:452
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "+=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 77:
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "-=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 78:
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "*=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 79:
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "/=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 80:
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "&=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 81:
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-2].tok.lit, Operator: "|=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 82:
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "++"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 83:
		//line parser.go.y:487
		{
			yyVAL.expr = &ast.AssocExpr{Name: yyS[yypt-1].tok.lit, Operator: "--"}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 84:
		//line parser.go.y:492
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 85:
		//line parser.go.y:497
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 86:
		//line parser.go.y:502
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 87:
		//line parser.go.y:507
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 88:
		//line parser.go.y:512
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 89:
		//line parser.go.y:517
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
