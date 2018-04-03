//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:28
type yySymType struct {
	yys             int
	compstmt        []ast.Stmt
	stmt_if         ast.Stmt
	stmt_default    ast.Stmt
	stmt_case       ast.Stmt
	stmt_cases      []ast.Stmt
	stmts           []ast.Stmt
	stmt            ast.Stmt
	expr            ast.Expr
	exprs           []ast.Expr
	expr_many       []ast.Expr
	expr_lets       ast.Expr
	expr_pair       ast.Expr
	expr_pairs      []ast.Expr
	expr_idents     []string
	expr_type       string
	tok             ast.Token
	term            ast.Token
	terms           ast.Token
	opt_terms       ast.Token
	array_count     ast.ArrayCount
	stmt_multi_case []ast.Stmt
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
const TYPE = 57393
const LEN = 57394
const DELETE = 57395
const UNARY = 57396

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
	"TYPE",
	"LEN",
	"DELETE",
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
	"'['",
	"']'",
	"'.'",
	"'!'",
	"'^'",
	"'&'",
	"'('",
	"')'",
	"'|'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:803

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 135,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	57, 54,
	-2, 1,
	-1, 10,
	57, 55,
	-2, 24,
	-1, 45,
	57, 54,
	-2, 136,
	-1, 87,
	67, 3,
	-2, 135,
	-1, 90,
	57, 55,
	-2, 51,
	-1, 91,
	16, 46,
	57, 46,
	-2, 58,
	-1, 93,
	67, 3,
	-2, 135,
	-1, 100,
	1, 63,
	8, 63,
	45, 63,
	46, 63,
	54, 63,
	56, 63,
	57, 63,
	66, 63,
	67, 63,
	68, 63,
	70, 63,
	76, 63,
	78, 63,
	-2, 58,
	-1, 102,
	1, 65,
	8, 65,
	45, 65,
	46, 65,
	54, 65,
	56, 65,
	57, 65,
	66, 65,
	67, 65,
	68, 65,
	70, 65,
	76, 65,
	78, 65,
	-2, 58,
	-1, 132,
	17, 0,
	18, 0,
	-2, 90,
	-1, 133,
	17, 0,
	18, 0,
	-2, 91,
	-1, 153,
	57, 55,
	-2, 51,
	-1, 155,
	67, 3,
	-2, 135,
	-1, 157,
	67, 3,
	-2, 135,
	-1, 159,
	67, 1,
	-2, 42,
	-1, 162,
	67, 3,
	-2, 135,
	-1, 189,
	67, 3,
	-2, 135,
	-1, 237,
	57, 56,
	-2, 52,
	-1, 238,
	1, 53,
	45, 53,
	46, 53,
	54, 53,
	56, 53,
	57, 57,
	67, 53,
	68, 53,
	78, 53,
	-2, 58,
	-1, 247,
	1, 57,
	8, 57,
	45, 57,
	46, 57,
	57, 57,
	67, 57,
	68, 57,
	70, 57,
	76, 57,
	78, 57,
	-2, 58,
	-1, 249,
	67, 3,
	-2, 135,
	-1, 251,
	67, 3,
	-2, 135,
	-1, 268,
	67, 3,
	-2, 135,
	-1, 278,
	1, 111,
	8, 111,
	45, 111,
	46, 111,
	54, 111,
	56, 111,
	57, 111,
	66, 111,
	67, 111,
	68, 111,
	70, 111,
	76, 111,
	78, 111,
	-2, 109,
	-1, 280,
	1, 115,
	8, 115,
	45, 115,
	46, 115,
	54, 115,
	56, 115,
	57, 115,
	66, 115,
	67, 115,
	68, 115,
	70, 115,
	76, 115,
	78, 115,
	-2, 113,
	-1, 295,
	67, 3,
	-2, 135,
	-1, 301,
	67, 3,
	-2, 135,
	-1, 302,
	67, 3,
	-2, 135,
	-1, 307,
	1, 110,
	8, 110,
	45, 110,
	46, 110,
	54, 110,
	56, 110,
	57, 110,
	66, 110,
	67, 110,
	68, 110,
	70, 110,
	76, 110,
	78, 110,
	-2, 108,
	-1, 308,
	1, 114,
	8, 114,
	45, 114,
	46, 114,
	54, 114,
	56, 114,
	57, 114,
	66, 114,
	67, 114,
	68, 114,
	70, 114,
	76, 114,
	78, 114,
	-2, 112,
	-1, 315,
	67, 3,
	-2, 135,
	-1, 316,
	67, 3,
	-2, 135,
	-1, 320,
	45, 3,
	46, 3,
	67, 3,
	-2, 135,
	-1, 324,
	67, 3,
	-2, 135,
	-1, 332,
	45, 3,
	46, 3,
	67, 3,
	-2, 135,
	-1, 333,
	45, 3,
	46, 3,
	67, 3,
	-2, 135,
	-1, 347,
	67, 3,
	-2, 135,
	-1, 348,
	67, 3,
	-2, 135,
}

const yyPrivate = 57344

const yyLast = 3018

var yyAct = [...]int{

	83, 175, 181, 10, 229, 257, 308, 256, 258, 230,
	47, 223, 11, 6, 1, 307, 84, 303, 269, 90,
	286, 94, 6, 7, 97, 98, 99, 101, 103, 88,
	42, 92, 7, 283, 229, 279, 108, 110, 277, 285,
	264, 114, 271, 219, 117, 82, 10, 229, 242, 211,
	121, 122, 282, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 193, 221, 144, 145, 146, 147, 105,
	149, 151, 153, 96, 154, 6, 172, 154, 115, 95,
	113, 120, 120, 96, 152, 7, 2, 167, 154, 178,
	44, 112, 158, 280, 111, 176, 278, 284, 164, 148,
	270, 218, 180, 273, 235, 171, 187, 212, 183, 161,
	153, 185, 154, 182, 259, 260, 166, 352, 351, 344,
	106, 107, 190, 340, 339, 336, 335, 173, 331, 186,
	321, 194, 156, 314, 313, 294, 255, 289, 275, 253,
	104, 197, 250, 248, 208, 202, 348, 200, 347, 324,
	10, 204, 205, 316, 153, 160, 306, 302, 301, 215,
	199, 268, 201, 155, 93, 119, 207, 206, 120, 287,
	154, 116, 319, 220, 159, 163, 231, 233, 232, 237,
	298, 225, 81, 241, 8, 157, 346, 243, 5, 246,
	259, 260, 341, 46, 239, 254, 120, 295, 86, 226,
	227, 176, 281, 188, 240, 265, 263, 191, 261, 262,
	234, 182, 48, 217, 216, 123, 276, 85, 4, 184,
	174, 89, 45, 209, 17, 3, 0, 288, 0, 0,
	118, 0, 0, 0, 46, 0, 0, 0, 0, 0,
	0, 198, 0, 293, 0, 0, 0, 0, 0, 0,
	296, 0, 210, 0, 291, 0, 292, 0, 0, 0,
	222, 224, 297, 246, 0, 0, 305, 0, 0, 0,
	0, 0, 0, 300, 309, 0, 0, 310, 311, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 0, 0, 0, 0, 0,
	317, 0, 0, 80, 0, 0, 322, 323, 272, 0,
	274, 0, 0, 52, 53, 54, 55, 56, 338, 0,
	329, 330, 79, 0, 51, 334, 0, 76, 78, 337,
	74, 0, 0, 0, 0, 0, 0, 342, 343, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 349, 350, 21, 22, 28, 0, 0, 32,
	14, 9, 15, 43, 0, 18, 0, 0, 0, 0,
	0, 0, 0, 38, 29, 30, 31, 16, 19, 0,
	0, 0, 0, 0, 0, 320, 0, 12, 13, 0,
	0, 0, 0, 0, 20, 0, 0, 36, 0, 39,
	40, 0, 37, 41, 0, 332, 333, 0, 0, 0,
	0, 23, 27, 0, 0, 0, 34, 0, 6, 33,
	0, 0, 24, 25, 26, 35, 0, 0, 7, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 50, 0, 327,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	0, 79, 0, 51, 0, 0, 76, 78, 326, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 50, 214,
	0, 62, 64, 52, 53, 54, 55, 56, 0, 0,
	0, 0, 79, 213, 51, 0, 0, 76, 78, 0,
	74, 60, 61, 63, 65, 75, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 50,
	196, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	0, 0, 0, 79, 195, 51, 0, 0, 76, 78,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	50, 0, 0, 62, 64, 52, 53, 54, 55, 56,
	0, 0, 0, 0, 79, 0, 51, 0, 0, 76,
	78, 345, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 0, 79, 0, 51, 0, 0,
	76, 78, 328, 74, 60, 61, 63, 65, 75, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 72, 73, 57, 58,
	59, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 50, 0, 0, 62, 64, 52, 53, 54,
	55, 56, 0, 0, 0, 0, 79, 0, 51, 0,
	0, 76, 78, 325, 74, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 50, 318, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 79, 0, 51,
	0, 0, 76, 78, 0, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 50, 0, 0, 62, 64, 52,
	53, 54, 55, 56, 0, 315, 0, 0, 79, 0,
	51, 0, 0, 76, 78, 0, 74, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 50, 0, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 0, 79,
	0, 51, 0, 0, 76, 78, 312, 74, 60, 61,
	63, 65, 75, 77, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 67, 68, 69, 70, 71, 0, 0,
	72, 73, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 50, 0, 0, 62,
	64, 52, 53, 54, 55, 56, 0, 0, 0, 0,
	79, 299, 51, 0, 0, 76, 78, 0, 74, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 50, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	0, 79, 290, 51, 0, 0, 76, 78, 0, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 50, 0,
	0, 62, 64, 52, 53, 54, 55, 56, 0, 0,
	0, 0, 79, 267, 51, 0, 0, 76, 78, 0,
	74, 60, 61, 63, 65, 75, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 50,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	0, 0, 252, 79, 0, 51, 0, 0, 76, 78,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	50, 0, 0, 62, 64, 52, 53, 54, 55, 56,
	0, 251, 0, 0, 79, 0, 51, 0, 0, 76,
	78, 0, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 249, 0, 0, 79, 0, 51, 0, 0,
	76, 78, 0, 74, 60, 61, 63, 65, 75, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 72, 73, 57, 58,
	59, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 50, 0, 0, 62, 64, 52, 53, 54,
	55, 56, 0, 0, 0, 0, 79, 245, 51, 0,
	0, 76, 78, 0, 74, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 50, 0, 236, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 79, 0, 51,
	0, 0, 76, 78, 0, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 50, 0, 0, 62, 64, 52,
	53, 54, 55, 56, 0, 0, 0, 0, 79, 0,
	51, 0, 0, 76, 78, 228, 74, 60, 61, 63,
	65, 75, 77, 0, 0, 0, 0, 0, 0, 0,
	0, 66, 67, 68, 69, 70, 71, 0, 0, 72,
	73, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 50, 192, 0, 62, 64,
	52, 53, 54, 55, 56, 0, 0, 0, 0, 79,
	0, 51, 0, 0, 76, 78, 0, 74, 60, 61,
	63, 65, 75, 77, 0, 0, 0, 0, 0, 0,
	0, 0, 66, 67, 68, 69, 70, 71, 0, 0,
	72, 73, 57, 58, 59, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 50, 0, 0, 62,
	64, 52, 53, 54, 55, 56, 0, 189, 0, 0,
	79, 0, 51, 0, 0, 76, 78, 0, 74, 60,
	61, 63, 65, 75, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 50, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	0, 79, 0, 51, 0, 0, 76, 78, 177, 74,
	60, 61, 63, 65, 75, 77, 0, 0, 0, 0,
	0, 0, 0, 0, 66, 67, 68, 69, 70, 71,
	0, 0, 72, 73, 57, 58, 59, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 50, 0,
	0, 62, 64, 52, 53, 54, 55, 56, 0, 165,
	0, 0, 79, 0, 51, 0, 0, 76, 78, 0,
	74, 60, 61, 63, 65, 75, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 66, 67, 68, 69, 70,
	71, 0, 0, 72, 73, 57, 58, 59, 0, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 50,
	0, 0, 62, 64, 52, 53, 54, 55, 56, 0,
	162, 0, 0, 79, 0, 51, 0, 0, 76, 78,
	0, 74, 60, 61, 63, 65, 75, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 49,
	50, 0, 0, 62, 64, 52, 53, 54, 55, 56,
	0, 0, 0, 0, 79, 0, 51, 0, 0, 76,
	78, 0, 74, 60, 61, 63, 65, 75, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 66, 67, 68,
	69, 70, 71, 0, 0, 72, 73, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 50, 0, 0, 62, 64, 52, 53, 54, 55,
	56, 0, 0, 0, 0, 79, 0, 51, 0, 0,
	76, 78, 0, 74, 60, 61, 63, 65, 75, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 72, 73, 57, 58,
	59, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 50, 0, 0, 62, 64, 52, 53, 54,
	55, 56, 0, 0, 0, 0, 79, 0, 51, 0,
	0, 76, 179, 0, 74, 60, 61, 63, 65, 75,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, 0, 0, 72, 73, 57,
	58, 59, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 50, 0, 0, 62, 64, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 79, 0, 170,
	0, 0, 76, 78, 0, 74, 60, 61, 63, 65,
	75, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 67, 68, 69, 70, 71, 0, 0, 72, 73,
	57, 58, 59, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 50, 0, 0, 62, 64, 52,
	53, 54, 55, 56, 0, 0, 0, 0, 79, 0,
	169, 0, 0, 76, 78, 0, 74, 21, 22, 203,
	0, 0, 32, 14, 9, 15, 43, 0, 18, 0,
	0, 0, 0, 0, 0, 0, 38, 29, 30, 31,
	16, 19, 0, 0, 0, 0, 0, 0, 0, 0,
	12, 13, 0, 0, 0, 0, 0, 20, 0, 0,
	36, 0, 39, 40, 0, 37, 41, 0, 0, 0,
	0, 0, 0, 0, 23, 27, 0, 0, 0, 34,
	0, 0, 33, 0, 0, 24, 25, 26, 35, 60,
	61, 63, 65, 0, 77, 0, 0, 0, 0, 0,
	0, 0, 0, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 0, 0, 0,
	0, 79, 0, 51, 0, 0, 76, 78, 0, 74,
	21, 22, 28, 0, 0, 32, 14, 9, 15, 43,
	0, 18, 0, 0, 0, 0, 0, 0, 0, 38,
	29, 30, 31, 16, 19, 0, 0, 0, 0, 0,
	0, 0, 0, 12, 13, 0, 0, 0, 0, 0,
	20, 0, 0, 36, 0, 39, 40, 0, 37, 41,
	0, 0, 0, 0, 0, 0, 0, 23, 27, 0,
	0, 0, 34, 0, 0, 33, 0, 0, 24, 25,
	26, 35, 60, 61, 63, 65, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 66, 67, 68, 69,
	70, 71, 0, 0, 72, 73, 57, 58, 59, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 62, 64, 52, 53, 54, 55, 56,
	0, 63, 65, 0, 79, 0, 51, 0, 0, 76,
	78, 0, 74, 66, 67, 68, 69, 70, 71, 0,
	0, 72, 73, 57, 58, 59, 0, 0, 247, 22,
	28, 0, 80, 32, 0, 0, 0, 0, 0, 0,
	62, 64, 52, 53, 54, 55, 56, 38, 29, 30,
	31, 79, 0, 51, 0, 0, 76, 78, 0, 74,
	21, 22, 28, 0, 0, 32, 0, 0, 0, 0,
	0, 36, 0, 39, 40, 0, 37, 41, 0, 38,
	29, 30, 31, 0, 0, 23, 27, 0, 0, 0,
	34, 0, 0, 33, 304, 0, 24, 25, 26, 35,
	0, 0, 0, 36, 0, 39, 40, 0, 37, 41,
	0, 0, 0, 0, 21, 22, 28, 23, 27, 32,
	0, 0, 34, 0, 0, 33, 266, 0, 24, 25,
	26, 35, 0, 38, 29, 30, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 21, 22, 28, 0, 0,
	32, 0, 0, 0, 0, 0, 0, 36, 0, 39,
	40, 0, 37, 41, 38, 29, 30, 31, 0, 0,
	0, 23, 27, 0, 0, 0, 34, 0, 0, 33,
	244, 0, 24, 25, 26, 35, 0, 0, 36, 0,
	39, 40, 0, 37, 41, 0, 0, 168, 0, 21,
	22, 28, 23, 27, 32, 0, 0, 34, 0, 0,
	33, 0, 0, 24, 25, 26, 35, 0, 38, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 21, 22, 28, 0, 0, 32, 0,
	0, 0, 36, 0, 39, 40, 0, 37, 41, 0,
	0, 150, 38, 29, 30, 31, 23, 27, 0, 0,
	0, 34, 0, 0, 33, 0, 0, 24, 25, 26,
	35, 0, 0, 0, 0, 0, 36, 0, 39, 40,
	0, 37, 41, 0, 0, 0, 0, 247, 22, 28,
	23, 27, 32, 0, 0, 34, 0, 0, 33, 0,
	0, 24, 25, 26, 35, 0, 38, 29, 30, 31,
	0, 0, 0, 0, 0, 0, 0, 0, 238, 22,
	28, 0, 0, 32, 0, 0, 0, 0, 0, 0,
	36, 0, 39, 40, 0, 37, 41, 38, 29, 30,
	31, 0, 0, 0, 23, 27, 0, 0, 0, 34,
	0, 0, 33, 0, 0, 24, 25, 26, 35, 0,
	0, 36, 0, 39, 40, 0, 37, 41, 0, 0,
	0, 0, 109, 22, 28, 23, 27, 32, 0, 0,
	34, 0, 0, 33, 0, 0, 24, 25, 26, 35,
	0, 38, 29, 30, 31, 0, 0, 0, 0, 0,
	0, 0, 0, 102, 22, 28, 0, 0, 32, 0,
	0, 0, 0, 0, 0, 36, 0, 39, 40, 0,
	37, 41, 38, 29, 30, 31, 0, 0, 0, 23,
	27, 0, 0, 0, 34, 0, 0, 33, 0, 0,
	24, 25, 26, 35, 0, 0, 36, 0, 39, 40,
	0, 37, 41, 0, 0, 0, 0, 100, 22, 28,
	23, 27, 32, 0, 0, 34, 0, 0, 33, 0,
	0, 24, 25, 26, 35, 0, 38, 29, 30, 31,
	0, 0, 0, 0, 0, 0, 0, 0, 91, 22,
	28, 0, 0, 32, 0, 0, 0, 0, 0, 0,
	36, 0, 39, 40, 0, 37, 41, 38, 29, 30,
	31, 0, 0, 0, 23, 27, 0, 0, 0, 34,
	0, 0, 33, 0, 0, 24, 25, 26, 35, 0,
	0, 36, 0, 39, 40, 0, 37, 41, 0, 0,
	0, 0, 0, 0, 0, 23, 27, 0, 0, 0,
	87, 0, 0, 33, 0, 0, 24, 25, 26, 35,
	66, 67, 68, 69, 70, 71, 0, 0, 0, 0,
	57, 66, 67, 68, 69, 70, 71, 0, 0, 80,
	0, 57, 0, 0, 0, 0, 0, 0, 0, 0,
	80, 54, 55, 56, 0, 0, 0, 0, 79, 0,
	51, 0, 0, 76, 78, 0, 74, 0, 0, 79,
	0, 51, 0, 0, 76, 78, 0, 74,
}
var yyPact = [...]int{

	-55, -1000, 2276, -55, -55, -1000, -1000, -1000, -1000, 218,
	1825, 138, -1000, -1000, 2629, 2629, 223, 194, 2884, 108,
	2629, 14, -1000, 2629, 2629, 2629, 2853, 2799, -1000, -1000,
	-1000, -1000, 75, -55, -55, 2629, 2768, 29, 26, 15,
	2629, 13, 124, 2629, -1000, 360, -1000, 121, -1000, 2629,
	2629, 221, 2629, 2629, 2629, 2629, 2629, 2629, 2629, 2629,
	2629, 2629, 2629, 2629, 2629, 2629, 2629, 2629, 2629, 2629,
	2629, 2629, -1000, -1000, 2629, 2629, 2629, 2629, 2629, 2595,
	2629, 2629, 123, 1886, 1886, 107, 129, -55, 149, 51,
	1764, 14, 131, -55, 1703, 2629, 2541, 2940, 2940, 2940,
	14, 2069, 14, 2008, 218, 11, 2629, 205, 1642, 24,
	1947, 2629, 217, 70, 1886, 2629, -55, 1581, -1000, 2629,
	-55, 1886, 1520, -1000, 2929, 2929, 2940, 2940, 2940, 1886,
	263, 263, 2382, 2382, 263, 263, 263, 263, 1886, 1886,
	1886, 1886, 1886, 1886, 1886, 2202, 1886, 2335, 65, 544,
	2629, 1886, -1000, 1886, -55, -55, 2629, -55, 88, 2143,
	2629, 2629, -55, 2629, 87, -55, 41, 483, 2629, 220,
	219, 35, 218, 17, -46, -1000, 135, -1000, 2629, 2629,
	1459, -67, -1000, 217, 119, 216, 44, 1398, 2714, -55,
	-1000, 210, 2629, -28, -1000, -1000, 2510, 1337, 2683, 86,
	1276, 85, -1000, 135, 1215, 1154, 82, -1000, 176, 79,
	155, -36, -1000, -1000, 2456, 1093, -1000, -1000, 105, -58,
	34, -55, 43, -55, 81, 2629, 30, 27, -1000, 208,
	-1000, -24, 37, -37, 122, -1000, 2629, 1886, 14, 80,
	-1000, 1886, -1000, 1032, -1000, -1000, 1886, 14, -1000, -55,
	-1000, -55, 2629, -1000, 141, -1000, -1000, -1000, -1000, 2629,
	134, -1000, -1000, -1000, -1000, 971, -1000, -1000, -55, 102,
	101, -59, 2424, -1000, 99, -1000, 1886, -61, -1000, -70,
	-1000, -1000, -1000, 2629, -1000, -1000, 2629, 2629, 910, -1000,
	-1000, 77, 76, 849, 97, -55, 788, 126, -55, -1000,
	73, -55, -55, 93, -1000, -1000, -1000, -1000, -1000, 727,
	422, 666, -1000, -1000, -1000, -55, -55, 71, -55, -55,
	-55, -1000, 69, 68, -55, -1000, -1000, 2629, -1000, 67,
	66, 172, -55, -55, -1000, -1000, -1000, 62, 605, -1000,
	166, 92, -1000, -1000, -1000, -1000, 90, -55, -55, 61,
	60, -1000, -1000,
}
var yyPgo = [...]int{

	0, 14, 235, 194, 234, 8, 7, 233, 0, 30,
	12, 231, 1, 230, 10, 2, 229, 5, 96, 228,
	198,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 7, 7,
	7, 7, 7, 7, 7, 6, 17, 5, 16, 16,
	16, 12, 13, 13, 13, 14, 14, 14, 15, 15,
	11, 10, 10, 10, 9, 9, 9, 9, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 18, 18, 19, 19, 20,
	20,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 3, 3, 1,
	1, 2, 2, 5, 1, 4, 7, 9, 5, 13,
	12, 9, 8, 5, 1, 7, 5, 5, 0, 2,
	2, 2, 2, 2, 2, 5, 5, 4, 0, 2,
	3, 3, 0, 1, 4, 0, 1, 4, 1, 3,
	3, 1, 4, 4, 0, 1, 4, 4, 1, 1,
	2, 2, 2, 2, 4, 2, 4, 1, 1, 1,
	1, 5, 3, 7, 8, 8, 9, 5, 6, 5,
	6, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 3, 3, 3, 3, 5, 4,
	6, 5, 5, 4, 6, 5, 4, 4, 6, 5,
	5, 6, 5, 5, 4, 4, 5, 7, 5, 7,
	9, 7, 3, 2, 6, 0, 1, 1, 2, 1,
	1,
}
var yyChk = [...]int{

	-1000, -1, -18, -2, -19, -20, 68, 78, -3, 11,
	-8, -10, 37, 38, 10, 12, 27, -4, 15, 28,
	44, 4, 5, 61, 72, 73, 74, 62, 6, 24,
	25, 26, 9, 69, 66, 75, 47, 52, 23, 49,
	50, 53, -9, 13, -18, -19, -20, -14, 4, 54,
	55, 71, 60, 61, 62, 63, 64, 41, 42, 43,
	17, 18, 58, 19, 59, 20, 31, 32, 33, 34,
	35, 36, 39, 40, 77, 21, 74, 22, 75, 69,
	50, 54, -9, -8, -8, 4, 14, 66, -14, -11,
	-8, 4, -10, 66, -8, 75, 69, -8, -8, -8,
	4, -8, 4, -8, 75, 4, -18, -18, -8, 4,
	-8, 75, 75, 75, -8, 75, 57, -8, -3, 54,
	57, -8, -8, 4, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -9, -8,
	56, -8, -10, -8, 57, 66, 13, 66, -1, -18,
	16, 68, 66, 54, -1, 66, -9, -8, 56, 71,
	71, -14, 75, -9, -13, -12, 6, 76, 75, 75,
	-8, -15, 4, 48, -16, 51, 69, -8, -18, 66,
	-10, -18, 56, 8, 76, 70, 56, -8, -18, -1,
	-8, -1, 67, 6, -8, -8, -1, -10, 67, -7,
	-18, 8, 76, 70, 56, -8, 4, 4, 76, 8,
	-14, 57, -18, 57, -18, 56, -9, -9, 76, 71,
	76, -15, 69, -15, 4, 70, 57, -8, 4, -1,
	4, -8, 76, -8, 70, 70, -8, 4, 67, 66,
	67, 66, 68, 67, 29, 67, -6, -17, -5, 45,
	46, -6, -5, -17, 76, -8, 70, 70, 66, 76,
	76, 8, -18, 70, -18, 67, -8, 8, 76, 8,
	76, 4, 76, 57, 70, 76, 57, 57, -8, 67,
	70, -1, -1, -8, 4, 66, -8, -10, 56, 70,
	-1, 66, 66, 76, 70, -12, 67, 76, 76, -8,
	-8, -8, 76, 67, 67, 66, 66, -1, 56, 56,
	-18, 67, -1, -1, 66, 76, 76, 57, 76, -1,
	-1, 67, -18, -18, -1, 67, 67, -1, -8, 67,
	67, 30, -1, -1, 67, 76, 30, 66, 66, -1,
	-1, 67, 67,
}
var yyDef = [...]int{

	-2, -2, -2, 135, 136, 137, 139, 140, 4, 45,
	-2, 0, 9, 10, 54, 0, 0, 14, 45, 0,
	0, 58, 59, 0, 0, 0, 0, 0, 67, 68,
	69, 70, 0, 135, 135, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 2, -2, 138, 0, 46, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 102, 103, 0, 0, 0, 0, 54, 0,
	0, 54, 11, 55, 12, 0, 0, -2, 0, 0,
	-2, -2, 0, -2, 0, 54, 0, 60, 61, 62,
	-2, 0, -2, 0, 45, 0, 54, 42, 0, 58,
	0, 0, 0, 38, 133, 0, 135, 0, 5, 54,
	135, 7, 0, 72, 82, 83, 84, 85, 86, 87,
	88, 89, -2, -2, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 104, 105, 106, 107, 0, 0,
	0, 132, 8, -2, 135, -2, 0, -2, 0, -2,
	0, 0, -2, 54, 0, 28, 0, 0, 0, 0,
	0, 0, 45, 135, 135, 43, 0, 81, 54, 54,
	0, 0, 48, 0, 0, 0, 0, 0, 0, -2,
	6, 0, 0, 0, 113, 117, 0, 0, 0, 0,
	0, 0, 15, 67, 0, 0, 0, 50, 0, 0,
	0, 0, 109, 116, 0, 0, 64, 66, 0, 0,
	0, 135, 0, 135, 0, 0, 0, 0, 124, 0,
	125, 0, 0, 0, 0, 39, 0, -2, -2, 0,
	47, 71, 112, 0, 122, 123, 56, -2, 13, -2,
	26, -2, 0, 18, 0, 23, 31, 33, 34, 54,
	0, 29, 30, 32, 108, 0, 119, 120, -2, 0,
	0, 0, 0, 77, 0, 79, 41, 0, -2, 0,
	-2, 49, 126, 0, 40, 128, 0, 0, 0, 27,
	121, 0, 0, 0, 0, -2, 55, 0, 135, 118,
	0, -2, -2, 0, 78, 44, 80, -2, -2, 0,
	0, 0, 134, 25, 16, -2, -2, 0, 135, 135,
	-2, 73, 0, 0, -2, 127, 129, 0, 131, 0,
	0, 22, -2, -2, 37, 74, 75, 0, 0, 17,
	21, 0, 35, 36, 76, 130, 0, -2, -2, 0,
	0, 20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	78, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 72, 3, 3, 3, 64, 74, 3,
	75, 76, 62, 60, 57, 61, 71, 63, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 56, 68,
	59, 54, 58, 55, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 70, 73, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 66, 77, 67,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 65,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:70
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:74
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:79
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:86
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:93
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
				if l, ok := yylex.(*Lexer); ok {
					l.stmts = yyVAL.stmts
				}
			}
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:104
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:109
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:113
		{
			if len(yyDollar[1].expr_many) == 2 && len(yyDollar[3].expr_many) == 1 {
				if _, ok := yyDollar[3].expr_many[0].(*ast.ItemExpr); ok {
					yyVAL.stmt = &ast.LetMapItemStmt{Lhss: yyDollar[1].expr_many, Rhs: yyDollar[3].expr_many[0]}
				} else {
					yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
				}
			} else {
				yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
			}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:125
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:130
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:135
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:140
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:145
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:150
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:155
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:160
		{
			yyVAL.stmt = &ast.ForStmt{Vars: yyDollar[2].expr_idents, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:165
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:170
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:175
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:180
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:185
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:190
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:195
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:200
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:208
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:213
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyDollar[4].compstmt...)
			}
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:222
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:228
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:232
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:236
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:240
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:244
		{
			yyVAL.stmt_cases = yyDollar[2].stmt_multi_case
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:248
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_multi_case...)
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:252
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:263
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:269
		{
			var cases = []ast.Stmt{}
			for _, stmt := range yyDollar[2].expr_many {
				cases = append(cases, &ast.CaseStmt{Expr: stmt, Stmts: yyDollar[5].compstmt})
			}
			yyVAL.stmt_multi_case = cases
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:284
		{
			yyVAL.array_count = ast.ArrayCount{Count: 0}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:288
		{
			yyVAL.array_count = ast.ArrayCount{Count: 1}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:292
		{
			yyVAL.array_count.Count = yyVAL.array_count.Count + 1
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:298
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:303
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:307
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:311
		{
			if len(yyDollar[1].expr_pairs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:319
		{
			yyVAL.expr_idents = []string{}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:323
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:327
		{
			if len(yyDollar[1].expr_idents) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:336
		{
			yyVAL.expr_type = yyDollar[1].tok.Lit
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:340
		{
			yyVAL.expr_type = yyVAL.expr_type + "." + yyDollar[3].tok.Lit
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:345
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:351
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:355
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:359
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:364
		{
			yyVAL.exprs = nil
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:368
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:372
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:379
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:388
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:393
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:398
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:403
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:408
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:413
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:418
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:423
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:428
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:433
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:438
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:443
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:448
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:453
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:458
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 73:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:463
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:468
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 75:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:473
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 76:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:478
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:483
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 78:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:488
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:493
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
	case 80:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:502
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
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:511
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:516
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:521
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:526
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:531
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:536
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:541
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:546
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:551
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:556
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:561
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:566
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:571
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:576
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:581
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:586
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:591
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:596
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:601
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:606
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:611
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:616
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:621
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:626
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:631
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:636
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:641
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:646
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:651
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:656
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:661
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:666
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:671
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:676
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:681
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:686
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:691
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:696
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:701
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:706
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:711
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 122:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:716
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:721
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:726
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:731
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:736
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:741
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 128:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:746
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 129:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:751
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 130:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:756
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 131:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:761
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].tok.Lit, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:766
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:771
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 134:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:776
		{
			yyVAL.expr = &ast.DeleteExpr{MapExpr: yyDollar[3].expr, KeyExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:789
		{
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:792
		{
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:797
		{
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:800
		{
		}
	}
	goto yystack /* stack new state and value */
}
