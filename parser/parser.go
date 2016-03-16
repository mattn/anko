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
const GO = 57389
const CHAN = 57390
const MAKE = 57391
const OPCHAN = 57392
const ARRAYLIT = 57393
const UNARY = 57394

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
	"GO",
	"CHAN",
	"MAKE",
	"OPCHAN",
	"ARRAYLIT",
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
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:690

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 2,
	-2, 118,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 9,
	55, 45,
	-2, 119,
	-1, 12,
	55, 46,
	-2, 23,
	-1, 86,
	65, 2,
	-2, 118,
	-1, 89,
	55, 46,
	-2, 42,
	-1, 91,
	65, 2,
	-2, 118,
	-1, 98,
	1, 54,
	8, 54,
	45, 54,
	46, 54,
	52, 54,
	54, 54,
	55, 54,
	64, 54,
	65, 54,
	66, 54,
	72, 54,
	74, 54,
	76, 54,
	-2, 49,
	-1, 100,
	1, 56,
	8, 56,
	45, 56,
	46, 56,
	52, 56,
	54, 56,
	55, 56,
	64, 56,
	65, 56,
	66, 56,
	72, 56,
	74, 56,
	76, 56,
	-2, 49,
	-1, 127,
	17, 0,
	18, 0,
	-2, 82,
	-1, 128,
	17, 0,
	18, 0,
	-2, 83,
	-1, 147,
	55, 46,
	-2, 42,
	-1, 149,
	65, 2,
	-2, 118,
	-1, 151,
	65, 2,
	-2, 118,
	-1, 156,
	65, 2,
	-2, 118,
	-1, 178,
	65, 2,
	-2, 118,
	-1, 219,
	55, 47,
	-2, 43,
	-1, 220,
	1, 44,
	45, 44,
	46, 44,
	52, 44,
	55, 48,
	65, 44,
	66, 44,
	76, 44,
	-2, 49,
	-1, 227,
	1, 48,
	8, 48,
	45, 48,
	46, 48,
	55, 48,
	65, 48,
	66, 48,
	72, 48,
	74, 48,
	76, 48,
	-2, 49,
	-1, 229,
	65, 2,
	-2, 118,
	-1, 231,
	65, 2,
	-2, 118,
	-1, 244,
	65, 2,
	-2, 118,
	-1, 255,
	1, 103,
	8, 103,
	45, 103,
	46, 103,
	52, 103,
	54, 103,
	55, 103,
	64, 103,
	65, 103,
	66, 103,
	72, 103,
	74, 103,
	76, 103,
	-2, 101,
	-1, 257,
	1, 111,
	8, 111,
	45, 111,
	46, 111,
	52, 111,
	54, 111,
	55, 111,
	64, 111,
	65, 111,
	66, 111,
	72, 111,
	74, 111,
	76, 111,
	-2, 109,
	-1, 267,
	65, 2,
	-2, 118,
	-1, 272,
	65, 2,
	-2, 118,
	-1, 273,
	65, 2,
	-2, 118,
	-1, 278,
	1, 102,
	8, 102,
	45, 102,
	46, 102,
	52, 102,
	54, 102,
	55, 102,
	64, 102,
	65, 102,
	66, 102,
	72, 102,
	74, 102,
	76, 102,
	-2, 100,
	-1, 279,
	1, 110,
	8, 110,
	45, 110,
	46, 110,
	52, 110,
	54, 110,
	55, 110,
	64, 110,
	65, 110,
	66, 110,
	72, 110,
	74, 110,
	76, 110,
	-2, 108,
	-1, 284,
	65, 2,
	-2, 118,
	-1, 285,
	65, 2,
	-2, 118,
	-1, 288,
	45, 2,
	46, 2,
	65, 2,
	-2, 118,
	-1, 292,
	65, 2,
	-2, 118,
	-1, 299,
	45, 2,
	46, 2,
	65, 2,
	-2, 118,
	-1, 312,
	65, 2,
	-2, 118,
	-1, 313,
	65, 2,
	-2, 118,
}

const yyNprod = 124
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 2078

var yyAct = [...]int{

	82, 169, 236, 237, 12, 46, 93, 6, 94, 279,
	12, 3, 1, 211, 8, 42, 13, 7, 83, 249,
	209, 89, 278, 92, 6, 274, 95, 96, 97, 99,
	101, 6, 81, 154, 7, 245, 256, 90, 106, 254,
	109, 7, 111, 173, 113, 94, 199, 104, 105, 116,
	117, 242, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 136,
	137, 138, 184, 225, 139, 140, 141, 142, 143, 259,
	145, 147, 103, 148, 148, 115, 148, 172, 93, 115,
	94, 166, 110, 148, 144, 161, 258, 146, 153, 152,
	257, 253, 246, 255, 158, 155, 205, 170, 164, 160,
	200, 317, 238, 239, 316, 147, 309, 306, 305, 148,
	167, 302, 301, 298, 177, 289, 283, 180, 282, 157,
	261, 179, 235, 266, 251, 233, 185, 230, 228, 196,
	190, 150, 260, 313, 312, 292, 285, 273, 272, 102,
	244, 188, 149, 91, 12, 192, 193, 148, 147, 112,
	186, 269, 187, 213, 189, 114, 277, 175, 115, 194,
	176, 198, 207, 80, 195, 238, 239, 311, 219, 210,
	212, 307, 223, 234, 224, 85, 10, 226, 214, 215,
	216, 221, 151, 267, 5, 247, 45, 206, 170, 44,
	222, 240, 241, 243, 44, 218, 217, 208, 204, 203,
	165, 118, 107, 84, 252, 47, 4, 168, 88, 9,
	197, 248, 19, 250, 2, 0, 0, 0, 0, 0,
	0, 0, 0, 265, 0, 0, 0, 0, 0, 268,
	0, 0, 263, 0, 264, 0, 0, 0, 0, 226,
	0, 0, 276, 0, 0, 0, 0, 271, 0, 0,
	280, 281, 0, 0, 65, 66, 67, 68, 69, 70,
	0, 0, 71, 72, 56, 0, 0, 0, 0, 0,
	286, 288, 0, 79, 0, 290, 291, 0, 0, 0,
	0, 51, 52, 53, 54, 55, 304, 296, 297, 299,
	0, 300, 75, 50, 78, 303, 77, 0, 73, 0,
	0, 0, 308, 59, 60, 62, 64, 74, 76, 0,
	0, 0, 0, 0, 0, 314, 315, 65, 66, 67,
	68, 69, 70, 0, 0, 71, 72, 56, 57, 58,
	0, 0, 0, 0, 0, 0, 79, 0, 0, 49,
	0, 295, 61, 63, 51, 52, 53, 54, 55, 59,
	60, 62, 64, 74, 76, 75, 50, 78, 294, 77,
	0, 73, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 49, 202, 0, 61, 63,
	51, 52, 53, 54, 55, 59, 60, 62, 64, 74,
	76, 75, 50, 78, 0, 77, 201, 73, 0, 65,
	66, 67, 68, 69, 70, 0, 0, 71, 72, 56,
	57, 58, 0, 0, 0, 0, 0, 0, 79, 0,
	0, 49, 183, 0, 61, 63, 51, 52, 53, 54,
	55, 59, 60, 62, 64, 74, 76, 75, 50, 78,
	0, 77, 182, 73, 0, 65, 66, 67, 68, 69,
	70, 0, 0, 71, 72, 56, 57, 58, 0, 0,
	0, 0, 0, 0, 79, 0, 0, 49, 0, 0,
	61, 63, 51, 52, 53, 54, 55, 59, 60, 62,
	64, 74, 76, 75, 50, 78, 310, 77, 0, 73,
	0, 65, 66, 67, 68, 69, 70, 0, 0, 71,
	72, 56, 57, 58, 0, 0, 0, 0, 0, 0,
	79, 0, 0, 49, 0, 0, 61, 63, 51, 52,
	53, 54, 55, 59, 60, 62, 64, 74, 76, 75,
	50, 78, 293, 77, 0, 73, 0, 65, 66, 67,
	68, 69, 70, 0, 0, 71, 72, 56, 57, 58,
	0, 0, 0, 0, 0, 0, 79, 0, 0, 49,
	287, 0, 61, 63, 51, 52, 53, 54, 55, 59,
	60, 62, 64, 74, 76, 75, 50, 78, 0, 77,
	0, 73, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 65, 66, 67, 68,
	69, 70, 79, 0, 0, 49, 56, 0, 61, 63,
	51, 52, 53, 54, 55, 79, 284, 0, 0, 0,
	0, 75, 50, 78, 0, 77, 0, 73, 59, 60,
	62, 64, 74, 76, 75, 50, 78, 0, 77, 0,
	73, 0, 65, 66, 67, 68, 69, 70, 0, 0,
	71, 72, 56, 57, 58, 0, 0, 0, 0, 0,
	0, 79, 0, 0, 49, 0, 0, 61, 63, 51,
	52, 53, 54, 55, 59, 60, 62, 64, 74, 76,
	75, 50, 78, 0, 77, 270, 73, 0, 65, 66,
	67, 68, 69, 70, 0, 0, 71, 72, 56, 57,
	58, 0, 0, 0, 0, 0, 0, 79, 0, 0,
	49, 0, 0, 61, 63, 51, 52, 53, 54, 55,
	59, 60, 62, 64, 74, 76, 75, 50, 78, 0,
	77, 262, 73, 0, 65, 66, 67, 68, 69, 70,
	0, 0, 71, 72, 56, 57, 58, 0, 0, 0,
	0, 0, 0, 79, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 232,
	0, 0, 75, 50, 78, 0, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 49, 0, 0, 61, 63,
	51, 52, 53, 54, 55, 0, 231, 0, 0, 0,
	0, 75, 50, 78, 0, 77, 0, 73, 59, 60,
	62, 64, 74, 76, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 66, 67, 68, 69, 70, 0, 0,
	71, 72, 56, 57, 58, 0, 0, 0, 0, 0,
	0, 79, 0, 0, 49, 0, 0, 61, 63, 51,
	52, 53, 54, 55, 0, 229, 0, 0, 0, 0,
	75, 50, 78, 0, 77, 0, 73, 59, 60, 62,
	64, 74, 76, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 66, 67, 68, 69, 70, 0, 0, 71,
	72, 56, 57, 58, 0, 0, 0, 0, 0, 0,
	79, 0, 0, 49, 181, 0, 61, 63, 51, 52,
	53, 54, 55, 59, 60, 62, 64, 74, 76, 75,
	50, 78, 0, 77, 0, 73, 0, 65, 66, 67,
	68, 69, 70, 0, 0, 71, 72, 56, 57, 58,
	0, 0, 0, 0, 0, 0, 79, 0, 0, 49,
	0, 0, 61, 63, 51, 52, 53, 54, 55, 0,
	178, 0, 0, 0, 0, 75, 50, 78, 0, 77,
	0, 73, 59, 60, 62, 64, 74, 76, 0, 0,
	0, 0, 0, 0, 0, 0, 65, 66, 67, 68,
	69, 70, 0, 0, 71, 72, 56, 57, 58, 0,
	0, 0, 0, 0, 0, 79, 0, 0, 49, 0,
	0, 61, 63, 51, 52, 53, 54, 55, 59, 60,
	62, 64, 74, 76, 75, 50, 78, 171, 77, 0,
	73, 0, 65, 66, 67, 68, 69, 70, 0, 0,
	71, 72, 56, 57, 58, 0, 0, 0, 0, 0,
	0, 79, 0, 0, 49, 0, 0, 61, 63, 51,
	52, 53, 54, 55, 0, 159, 0, 0, 0, 0,
	75, 50, 78, 0, 77, 0, 73, 59, 60, 62,
	64, 74, 76, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 66, 67, 68, 69, 70, 0, 0, 71,
	72, 56, 57, 58, 0, 0, 0, 0, 0, 0,
	79, 0, 0, 49, 0, 0, 61, 63, 51, 52,
	53, 54, 55, 0, 156, 0, 0, 0, 0, 75,
	50, 78, 0, 77, 0, 73, 59, 60, 62, 64,
	74, 76, 0, 0, 0, 0, 0, 0, 0, 0,
	65, 66, 67, 68, 69, 70, 0, 0, 71, 72,
	56, 57, 58, 0, 0, 0, 0, 0, 0, 79,
	0, 48, 49, 0, 0, 61, 63, 51, 52, 53,
	54, 55, 0, 0, 0, 0, 0, 0, 75, 50,
	78, 0, 77, 0, 73, 23, 24, 30, 0, 0,
	34, 16, 11, 17, 43, 0, 20, 0, 0, 0,
	0, 0, 0, 0, 38, 31, 32, 33, 18, 21,
	0, 0, 0, 0, 0, 0, 0, 0, 14, 15,
	0, 0, 0, 0, 0, 22, 0, 0, 39, 0,
	40, 41, 0, 0, 0, 0, 0, 0, 0, 0,
	25, 29, 0, 0, 0, 36, 0, 6, 26, 27,
	28, 0, 37, 0, 35, 0, 0, 7, 59, 60,
	62, 64, 74, 76, 0, 0, 0, 0, 0, 0,
	0, 0, 65, 66, 67, 68, 69, 70, 0, 0,
	71, 72, 56, 57, 58, 0, 0, 0, 0, 0,
	0, 79, 0, 0, 49, 0, 0, 61, 63, 51,
	52, 53, 54, 55, 59, 60, 62, 64, 74, 76,
	75, 50, 78, 0, 77, 0, 73, 0, 65, 66,
	67, 68, 69, 70, 0, 0, 71, 72, 56, 57,
	58, 0, 0, 0, 0, 0, 0, 79, 0, 0,
	49, 0, 0, 61, 63, 51, 52, 53, 54, 55,
	59, 60, 62, 64, 74, 76, 75, 50, 174, 0,
	77, 0, 73, 0, 65, 66, 67, 68, 69, 70,
	0, 0, 71, 72, 56, 57, 58, 0, 0, 0,
	0, 0, 0, 79, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 59, 60, 62, 64,
	74, 76, 75, 163, 78, 0, 77, 0, 73, 0,
	65, 66, 67, 68, 69, 70, 0, 0, 71, 72,
	56, 57, 58, 0, 0, 0, 0, 0, 0, 79,
	0, 0, 49, 0, 0, 61, 63, 51, 52, 53,
	54, 55, 59, 60, 62, 64, 0, 76, 75, 162,
	78, 0, 77, 0, 73, 0, 65, 66, 67, 68,
	69, 70, 0, 0, 71, 72, 56, 57, 58, 0,
	0, 0, 0, 0, 0, 79, 0, 0, 0, 0,
	0, 61, 63, 51, 52, 53, 54, 55, 59, 60,
	62, 64, 0, 0, 75, 50, 78, 0, 77, 0,
	73, 0, 65, 66, 67, 68, 69, 70, 0, 0,
	71, 72, 56, 57, 58, 0, 0, 0, 0, 0,
	0, 79, 0, 0, 0, 0, 0, 61, 63, 51,
	52, 53, 54, 55, 0, 0, 0, 0, 0, 0,
	75, 50, 78, 0, 77, 0, 73, 23, 24, 191,
	0, 0, 34, 16, 11, 17, 43, 0, 20, 0,
	0, 0, 0, 0, 0, 0, 38, 31, 32, 33,
	18, 21, 0, 0, 0, 0, 0, 0, 0, 0,
	14, 15, 0, 0, 0, 0, 0, 22, 0, 0,
	39, 0, 40, 41, 0, 0, 0, 0, 0, 0,
	0, 0, 25, 29, 0, 0, 0, 36, 0, 0,
	26, 27, 28, 0, 37, 0, 35, 23, 24, 30,
	0, 0, 34, 16, 11, 17, 43, 0, 20, 0,
	0, 0, 0, 0, 0, 0, 38, 31, 32, 33,
	18, 21, 0, 0, 0, 0, 0, 0, 0, 0,
	14, 15, 0, 0, 0, 0, 0, 22, 0, 0,
	39, 0, 40, 41, 0, 0, 0, 0, 0, 0,
	0, 0, 25, 29, 0, 62, 64, 36, 0, 0,
	26, 27, 28, 0, 37, 0, 35, 65, 66, 67,
	68, 69, 70, 0, 0, 71, 72, 56, 57, 58,
	227, 24, 30, 0, 0, 34, 79, 0, 0, 0,
	0, 0, 61, 63, 51, 52, 53, 54, 55, 38,
	31, 32, 33, 0, 0, 75, 50, 78, 0, 77,
	0, 73, 23, 24, 30, 0, 0, 34, 0, 0,
	0, 0, 0, 39, 0, 40, 41, 0, 0, 0,
	0, 38, 31, 32, 33, 25, 29, 0, 0, 0,
	36, 0, 0, 26, 27, 28, 0, 37, 0, 35,
	275, 0, 0, 0, 0, 39, 0, 40, 41, 0,
	0, 0, 0, 0, 227, 24, 30, 25, 29, 34,
	0, 0, 36, 0, 0, 26, 27, 28, 0, 37,
	0, 35, 0, 38, 31, 32, 33, 0, 0, 0,
	0, 0, 220, 24, 30, 0, 0, 34, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 39, 0, 40,
	41, 38, 31, 32, 33, 0, 0, 0, 0, 25,
	29, 108, 24, 30, 36, 0, 34, 26, 27, 28,
	0, 37, 0, 35, 0, 39, 0, 40, 41, 0,
	38, 31, 32, 33, 0, 0, 0, 25, 29, 100,
	24, 30, 36, 0, 34, 26, 27, 28, 0, 37,
	0, 35, 0, 0, 39, 0, 40, 41, 38, 31,
	32, 33, 0, 0, 0, 0, 25, 29, 98, 24,
	30, 36, 0, 34, 26, 27, 28, 0, 37, 0,
	35, 0, 39, 0, 40, 41, 0, 38, 31, 32,
	33, 0, 0, 0, 25, 29, 87, 24, 30, 36,
	0, 34, 26, 27, 28, 0, 37, 0, 35, 0,
	0, 39, 0, 40, 41, 38, 31, 32, 33, 0,
	0, 0, 0, 25, 29, 0, 0, 0, 36, 0,
	0, 26, 27, 28, 0, 37, 0, 35, 0, 39,
	0, 40, 41, 65, 66, 67, 68, 69, 70, 0,
	0, 25, 29, 56, 0, 0, 86, 0, 0, 26,
	27, 28, 79, 37, 0, 35, 0, 0, 0, 0,
	0, 0, 53, 54, 55, 0, 0, 0, 0, 0,
	0, 75, 50, 78, 0, 77, 0, 73,
}
var yyPact = [...]int{

	-59, -1000, -59, 1673, -59, -1000, -1000, -1000, -1000, 1241,
	-1000, 211, 1169, 121, -1000, -1000, 1788, 1788, 209, 171,
	1982, 89, 1788, -65, -1000, 1788, 1788, 1788, 1954, 1925,
	-1000, -1000, -1000, -1000, 78, -59, -59, 1788, 208, 1897,
	21, 1788, 104, 1788, -1000, -1000, 113, -1000, 1788, 1788,
	207, 1788, 1788, 1788, 1788, 1788, 1788, 1788, 1788, 1788,
	1788, 1788, 1788, 1788, 1788, 1788, 1788, 1788, 1788, 1788,
	1788, -1000, -1000, 1788, 1788, 1788, 1788, 1788, 1788, 1788,
	1788, 102, 1301, 1301, 88, 128, -59, 17, 39, 1110,
	77, -59, 1051, 1788, 1788, 585, 585, 585, -65, 1439,
	-65, 1393, 206, 20, 1788, 192, 1005, 16, -28, 1347,
	119, 1301, -59, 946, 1788, -59, 1301, 900, -1000, 2002,
	2002, 585, 585, 585, 1301, 233, 233, 1716, 1716, 233,
	233, 233, 233, 1301, 1301, 1301, 1301, 1301, 1301, 1301,
	1485, 1301, 1531, 388, 64, 1301, -1000, 1301, -59, -59,
	1788, -59, 75, 1603, 1788, 1788, -59, 1788, 74, -59,
	38, 342, 205, 204, 34, 189, 203, -35, -42, -1000,
	109, -1000, 1788, 1788, 1788, 202, 201, 1868, -59, -1000,
	196, 1788, -1000, 1788, 1, -1000, 1840, 73, 841, 72,
	-1000, 109, 782, 723, 70, -1000, 154, 67, 130, -21,
	-1000, -1000, 1788, -1000, -1000, 86, -37, 30, 187, -59,
	-55, -59, 69, 1788, 29, 31, 28, 24, 87, 1301,
	-65, 65, -1000, 1301, 677, -1000, 1301, -65, -1000, -59,
	-1000, -59, 1788, -1000, 129, -1000, -1000, -1000, 1788, 107,
	-1000, -1000, -1000, 631, -59, 84, 83, -47, 1756, -1000,
	101, -1000, 1301, -1000, -50, -1000, -63, -1000, -1000, 1788,
	1788, -1000, -1000, 63, 61, 572, 82, -59, 526, -59,
	-1000, 60, -59, -59, 81, -1000, -1000, -1000, -1000, -1000,
	480, 296, -1000, -1000, -59, -59, 58, -59, -59, -1000,
	57, 56, -59, -1000, -1000, 1788, 53, 52, 151, -59,
	-1000, -1000, -1000, 51, 434, -1000, 147, 80, -1000, -1000,
	-1000, 79, -59, -59, 49, 46, -1000, -1000,
}
var yyPgo = [...]int{

	0, 12, 224, 186, 222, 3, 2, 220, 0, 15,
	16, 218, 1, 217, 5, 11, 216, 194,
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
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 15, 15,
	16, 16, 17, 17,
}
var yyR2 = [...]int{

	0, 2, 0, 2, 3, 4, 3, 3, 1, 1,
	2, 2, 5, 1, 4, 7, 9, 5, 13, 12,
	9, 8, 5, 1, 7, 5, 5, 0, 2, 2,
	2, 2, 5, 4, 3, 0, 1, 4, 0, 1,
	4, 3, 1, 4, 4, 0, 1, 4, 4, 1,
	1, 2, 2, 2, 2, 4, 2, 4, 1, 1,
	1, 1, 5, 3, 7, 8, 8, 9, 5, 6,
	5, 6, 3, 5, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 3, 3, 3, 3,
	5, 4, 6, 5, 4, 4, 6, 6, 5, 4,
	6, 5, 5, 7, 7, 9, 3, 2, 0, 1,
	1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -15, -16, -17, 66, 76, -15, -16,
	-3, 11, -8, -10, 37, 38, 10, 12, 27, -4,
	15, 28, 44, 4, 5, 59, 67, 68, 69, 60,
	6, 24, 25, 26, 9, 73, 64, 71, 23, 47,
	49, 50, -9, 13, -17, -3, -14, 4, 52, 53,
	70, 58, 59, 60, 61, 62, 41, 42, 43, 17,
	18, 56, 19, 57, 20, 31, 32, 33, 34, 35,
	36, 39, 40, 75, 21, 69, 22, 73, 71, 50,
	52, -9, -8, -8, 4, 14, 64, 4, -11, -8,
	-10, 64, -8, 71, 73, -8, -8, -8, 4, -8,
	4, -8, 71, 4, -15, -15, -8, 4, 4, -8,
	71, -8, 55, -8, 52, 55, -8, -8, 4, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -9, -8, -10, -8, 55, 64,
	13, 64, -1, -15, 16, 66, 64, 52, -1, 64,
	-9, -8, 70, 70, -14, 4, 71, -9, -13, -12,
	6, 72, 71, 71, 71, 48, 51, -15, 64, -10,
	-15, 54, 74, 54, 8, 72, -15, -1, -8, -1,
	65, 6, -8, -8, -1, -10, 65, -7, -15, 8,
	72, 74, 54, 4, 4, 72, 8, -14, 4, 55,
	-15, 55, -15, 54, -9, -9, -9, 4, 4, -8,
	4, -1, 4, -8, -8, 72, -8, 4, 65, 64,
	65, 64, 66, 65, 29, 65, -6, -5, 45, 46,
	-6, -5, 72, -8, 64, 72, 72, 8, -15, 74,
	-15, 65, -8, 72, 8, 72, 8, 72, 72, 55,
	55, 65, 74, -1, -1, -8, 4, 64, -8, 54,
	74, -1, 64, 64, 72, 74, -12, 65, 72, 72,
	-8, -8, 65, 65, 64, 64, -1, 54, -15, 65,
	-1, -1, 64, 72, 72, 55, -1, -1, 65, -15,
	-1, 65, 65, -1, -8, 65, 65, 30, -1, 65,
	72, 30, 64, 64, -1, -1, 65, 65,
}
var yyDef = [...]int{

	-2, -2, 118, 45, 119, 120, 122, 123, 1, -2,
	3, 38, -2, 0, 8, 9, 45, 0, 0, 13,
	45, 0, 0, 49, 50, 0, 0, 0, 0, 0,
	58, 59, 60, 61, 0, 118, 118, 0, 0, 0,
	0, 0, 0, 0, 121, 4, 0, 39, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 94, 95, 0, 0, 0, 0, 0, 45, 0,
	45, 10, 46, 11, 0, 0, -2, 49, 0, -2,
	0, -2, 0, 45, 0, 51, 52, 53, -2, 0,
	-2, 0, 38, 0, 45, 35, 0, 0, 49, 0,
	0, 117, 118, 0, 45, 118, 6, 0, 63, 74,
	75, 76, 77, 78, 79, 80, 81, -2, -2, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 96,
	97, 98, 99, 0, 0, 116, 7, -2, 118, -2,
	0, -2, 0, 35, 0, 0, -2, 45, 0, 27,
	0, 0, 0, 0, 0, 39, 38, 118, 118, 36,
	0, 72, 45, 45, 45, 0, 0, 0, -2, 5,
	0, 0, 105, 0, 0, 109, 0, 0, 0, 0,
	14, 58, 0, 0, 0, 41, 0, 0, 0, 0,
	101, 104, 0, 55, 57, 0, 0, 0, 39, 118,
	0, 118, 0, 0, 0, 0, 0, 0, 0, -2,
	-2, 0, 40, 62, 0, 108, 47, -2, 12, -2,
	25, -2, 0, 17, 0, 22, 30, 31, 0, 0,
	28, 29, 100, 0, -2, 0, 0, 0, 0, 68,
	0, 70, 34, 73, 0, -2, 0, -2, 112, 0,
	0, 26, 107, 0, 0, 0, 0, -2, 0, 118,
	106, 0, -2, -2, 0, 69, 37, 71, -2, -2,
	0, 0, 24, 15, -2, -2, 0, 118, -2, 64,
	0, 0, -2, 113, 114, 0, 0, 0, 21, -2,
	33, 65, 66, 0, 0, 16, 20, 0, 32, 67,
	115, 0, -2, -2, 0, 0, 19, 18,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	76, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 67, 3, 3, 3, 62, 69, 3,
	71, 72, 60, 58, 55, 59, 70, 61, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 54, 66,
	57, 52, 56, 53, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 73, 3, 74, 68, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 64, 75, 65,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	63,
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
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
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
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			yyrcvr.char = -1
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:75
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
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
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:141
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyDollar[2].tok.Lit, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:146
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:151
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
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
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:579
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 101:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:584
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 102:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:589
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:594
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:599
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:604
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:609
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 107:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:614
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:619
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:624
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:629
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:634
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:639
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].tok.Lit, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:644
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].tok.Lit, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:649
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].tok.Lit, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:654
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].tok.Lit, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:659
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:664
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:675
		{
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:678
		{
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:683
		{
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:686
		{
		}
	}
	goto yystack /* stack new state and value */
}
