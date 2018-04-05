//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:29
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
	expr_slice      ast.Expr
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
	"'('",
	"')'",
	"'['",
	"']'",
	"'.'",
	"'!'",
	"'^'",
	"'&'",
	"'|'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:817

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 138,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	57, 58,
	-2, 1,
	-1, 10,
	57, 59,
	-2, 28,
	-1, 46,
	57, 58,
	-2, 139,
	-1, 89,
	67, 3,
	-2, 138,
	-1, 92,
	57, 59,
	-2, 55,
	-1, 93,
	16, 50,
	57, 50,
	-2, 68,
	-1, 95,
	67, 3,
	-2, 138,
	-1, 104,
	1, 73,
	8, 73,
	45, 73,
	46, 73,
	54, 73,
	56, 73,
	57, 73,
	66, 73,
	67, 73,
	68, 73,
	70, 73,
	72, 73,
	78, 73,
	-2, 68,
	-1, 106,
	1, 75,
	8, 75,
	45, 75,
	46, 75,
	54, 75,
	56, 75,
	57, 75,
	66, 75,
	67, 75,
	68, 75,
	70, 75,
	72, 75,
	78, 75,
	-2, 68,
	-1, 136,
	17, 0,
	18, 0,
	-2, 100,
	-1, 137,
	17, 0,
	18, 0,
	-2, 101,
	-1, 156,
	57, 59,
	-2, 55,
	-1, 158,
	67, 3,
	-2, 138,
	-1, 160,
	67, 3,
	-2, 138,
	-1, 162,
	67, 1,
	-2, 46,
	-1, 165,
	67, 3,
	-2, 138,
	-1, 192,
	67, 3,
	-2, 138,
	-1, 241,
	57, 60,
	-2, 56,
	-1, 242,
	1, 57,
	45, 57,
	46, 57,
	54, 57,
	56, 57,
	57, 61,
	67, 57,
	68, 57,
	78, 57,
	-2, 68,
	-1, 251,
	1, 61,
	8, 61,
	45, 61,
	46, 61,
	57, 61,
	67, 61,
	68, 61,
	70, 61,
	72, 61,
	78, 61,
	-2, 68,
	-1, 253,
	67, 3,
	-2, 138,
	-1, 255,
	67, 3,
	-2, 138,
	-1, 269,
	1, 25,
	45, 25,
	46, 25,
	67, 25,
	68, 25,
	78, 25,
	-2, 119,
	-1, 271,
	1, 27,
	45, 27,
	46, 27,
	67, 27,
	68, 27,
	78, 27,
	-2, 121,
	-1, 276,
	67, 3,
	-2, 138,
	-1, 299,
	67, 3,
	-2, 138,
	-1, 303,
	1, 24,
	45, 24,
	46, 24,
	67, 24,
	68, 24,
	78, 24,
	-2, 118,
	-1, 304,
	1, 26,
	45, 26,
	46, 26,
	67, 26,
	68, 26,
	78, 26,
	-2, 120,
	-1, 307,
	67, 3,
	-2, 138,
	-1, 308,
	67, 3,
	-2, 138,
	-1, 319,
	67, 3,
	-2, 138,
	-1, 320,
	67, 3,
	-2, 138,
	-1, 324,
	45, 3,
	46, 3,
	67, 3,
	-2, 138,
	-1, 328,
	67, 3,
	-2, 138,
	-1, 336,
	45, 3,
	46, 3,
	67, 3,
	-2, 138,
	-1, 337,
	45, 3,
	46, 3,
	67, 3,
	-2, 138,
	-1, 351,
	67, 3,
	-2, 138,
	-1, 352,
	67, 3,
	-2, 138,
}

const yyPrivate = 57344

const yyLast = 3009

var yyAct = [...]int{

	85, 180, 261, 10, 260, 262, 184, 6, 288, 281,
	238, 11, 48, 1, 309, 304, 86, 7, 43, 92,
	303, 96, 98, 228, 277, 101, 102, 103, 105, 107,
	94, 90, 233, 84, 6, 232, 100, 112, 99, 169,
	177, 99, 116, 275, 7, 119, 249, 10, 185, 117,
	115, 123, 124, 126, 114, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 138, 139, 140, 141, 142,
	143, 144, 145, 146, 147, 109, 290, 148, 149, 150,
	151, 2, 153, 154, 156, 45, 251, 23, 29, 289,
	226, 33, 232, 279, 113, 155, 164, 287, 356, 152,
	171, 6, 355, 161, 348, 39, 30, 31, 32, 167,
	286, 7, 344, 232, 183, 235, 110, 111, 190, 173,
	270, 176, 156, 181, 186, 343, 197, 188, 340, 178,
	339, 40, 41, 193, 38, 42, 268, 263, 264, 335,
	108, 325, 122, 24, 28, 318, 224, 189, 35, 298,
	317, 36, 219, 34, 310, 278, 25, 26, 27, 259,
	203, 199, 293, 10, 207, 208, 283, 156, 257, 157,
	254, 162, 202, 218, 204, 252, 211, 205, 210, 209,
	352, 351, 271, 163, 312, 157, 328, 159, 214, 215,
	225, 320, 241, 234, 236, 122, 245, 308, 269, 248,
	191, 157, 250, 307, 194, 276, 243, 158, 223, 95,
	157, 299, 121, 291, 220, 122, 267, 272, 265, 266,
	157, 118, 323, 200, 122, 302, 230, 166, 350, 83,
	345, 284, 68, 69, 70, 71, 72, 73, 5, 201,
	160, 292, 59, 47, 263, 264, 8, 258, 88, 181,
	213, 81, 285, 244, 237, 185, 49, 297, 222, 221,
	227, 229, 127, 87, 300, 37, 187, 295, 179, 296,
	80, 91, 51, 212, 53, 301, 4, 78, 76, 17,
	46, 250, 3, 0, 311, 47, 0, 0, 313, 0,
	306, 314, 315, 120, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 280, 0,
	282, 0, 0, 321, 0, 0, 0, 0, 0, 0,
	0, 326, 327, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 342, 333, 334, 0, 0, 0, 338, 0,
	0, 0, 341, 0, 0, 0, 0, 0, 0, 0,
	346, 347, 0, 82, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 353, 354, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 324, 0, 0, 81, 0, 0,
	0, 0, 52, 0, 331, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 336, 337, 80, 330, 51, 0,
	53, 0, 0, 78, 76, 82, 62, 63, 65, 67,
	77, 79, 0, 0, 0, 0, 0, 0, 0, 0,
	68, 69, 70, 71, 72, 73, 0, 0, 74, 75,
	59, 60, 61, 0, 0, 0, 0, 0, 0, 81,
	0, 0, 0, 0, 52, 0, 240, 64, 66, 54,
	55, 56, 57, 58, 0, 0, 0, 0, 80, 239,
	51, 0, 53, 0, 0, 78, 76, 82, 62, 63,
	65, 67, 77, 79, 0, 0, 0, 0, 0, 0,
	0, 0, 68, 69, 70, 71, 72, 73, 0, 0,
	74, 75, 59, 60, 61, 0, 0, 0, 0, 0,
	0, 81, 0, 0, 0, 0, 52, 216, 0, 64,
	66, 54, 55, 56, 57, 58, 0, 0, 0, 0,
	80, 0, 51, 217, 53, 0, 0, 78, 76, 82,
	62, 63, 65, 67, 77, 79, 0, 0, 0, 0,
	0, 0, 0, 0, 68, 69, 70, 71, 72, 73,
	0, 0, 74, 75, 59, 60, 61, 0, 0, 0,
	0, 0, 0, 81, 0, 0, 0, 0, 52, 195,
	0, 64, 66, 54, 55, 56, 57, 58, 0, 0,
	0, 0, 80, 0, 51, 196, 53, 0, 0, 78,
	76, 22, 23, 29, 0, 0, 33, 14, 9, 15,
	44, 0, 18, 0, 0, 0, 0, 0, 0, 0,
	39, 30, 31, 32, 16, 19, 0, 0, 0, 0,
	0, 0, 0, 0, 12, 13, 0, 0, 0, 0,
	0, 20, 0, 0, 21, 0, 40, 41, 0, 38,
	42, 0, 0, 0, 0, 0, 0, 0, 24, 28,
	0, 0, 0, 35, 0, 6, 36, 0, 34, 0,
	0, 25, 26, 27, 0, 7, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 80,
	349, 51, 0, 53, 0, 0, 78, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 80, 332, 51, 0, 53, 0, 0, 78, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 80, 329, 51, 0, 53, 0, 0,
	78, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 322, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 80, 0, 51, 0, 53,
	0, 0, 78, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 319, 0, 0, 80, 0, 51,
	0, 53, 0, 0, 78, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 80,
	316, 51, 0, 53, 0, 0, 78, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 80, 0, 51, 305, 53, 0, 0, 78, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 80, 0, 51, 294, 53, 0, 0,
	78, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 80, 0, 51, 274, 53,
	0, 0, 78, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 256, 80, 0, 51,
	0, 53, 0, 0, 78, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 255, 0, 0, 80,
	0, 51, 0, 53, 0, 0, 78, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 253, 0,
	0, 80, 0, 51, 0, 53, 0, 0, 78, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 80, 0, 51, 247, 53, 0, 0,
	78, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 80, 231, 51, 0, 53,
	0, 0, 78, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 198, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 80, 0, 51,
	0, 53, 0, 0, 78, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 192, 0, 0, 80,
	0, 51, 0, 53, 0, 0, 78, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 80, 182, 51, 0, 53, 0, 0, 78, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	168, 0, 0, 80, 0, 51, 0, 53, 0, 0,
	78, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 165, 0, 0, 80, 0, 51, 0, 53,
	0, 0, 78, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 50, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 80, 0, 51,
	0, 53, 0, 0, 78, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 80,
	0, 51, 0, 53, 0, 0, 78, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 80, 0, 51, 0, 175, 0, 0, 78, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 80, 0, 51, 0, 174, 0, 0,
	78, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 170, 0, 51, 0, 53,
	0, 0, 78, 76, 22, 23, 206, 0, 0, 33,
	14, 9, 15, 44, 0, 18, 0, 0, 0, 0,
	0, 0, 0, 39, 30, 31, 32, 16, 19, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 13, 0,
	0, 0, 0, 0, 20, 0, 0, 21, 0, 40,
	41, 0, 38, 42, 0, 0, 0, 0, 0, 0,
	0, 24, 28, 0, 0, 0, 35, 0, 0, 36,
	0, 34, 0, 0, 25, 26, 27, 82, 62, 63,
	65, 67, 0, 79, 0, 0, 0, 0, 0, 0,
	0, 0, 68, 69, 70, 71, 72, 73, 0, 0,
	74, 75, 59, 60, 61, 0, 0, 0, 0, 0,
	0, 81, 0, 0, 0, 0, 0, 0, 0, 64,
	66, 54, 55, 56, 57, 58, 0, 0, 0, 0,
	80, 0, 51, 0, 53, 0, 0, 78, 76, 22,
	23, 29, 0, 0, 33, 14, 9, 15, 44, 0,
	18, 0, 0, 0, 0, 0, 0, 0, 39, 30,
	31, 32, 16, 19, 0, 0, 0, 0, 0, 0,
	0, 0, 12, 13, 0, 0, 0, 0, 0, 20,
	0, 0, 21, 0, 40, 41, 0, 38, 42, 0,
	0, 0, 0, 0, 0, 0, 24, 28, 0, 0,
	0, 35, 0, 0, 36, 0, 34, 0, 0, 25,
	26, 27, 82, 62, 63, 65, 67, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 65, 67, 0, 80, 0, 51, 0, 53,
	0, 0, 78, 76, 68, 69, 70, 71, 72, 73,
	0, 0, 74, 75, 59, 60, 61, 0, 0, 22,
	23, 29, 0, 81, 33, 0, 0, 0, 0, 0,
	0, 64, 66, 54, 55, 56, 57, 58, 39, 30,
	31, 32, 80, 0, 51, 0, 53, 0, 0, 78,
	76, 22, 23, 29, 0, 0, 33, 0, 0, 0,
	0, 0, 0, 0, 40, 41, 0, 38, 42, 0,
	39, 30, 31, 32, 0, 0, 24, 28, 0, 0,
	0, 35, 0, 0, 36, 0, 34, 273, 0, 25,
	26, 27, 0, 0, 0, 0, 40, 41, 0, 38,
	42, 0, 0, 0, 0, 0, 0, 0, 24, 28,
	0, 0, 0, 35, 0, 0, 36, 0, 34, 246,
	0, 25, 26, 27, 68, 69, 70, 71, 72, 73,
	0, 0, 74, 75, 59, 0, 0, 0, 0, 22,
	23, 29, 0, 81, 33, 0, 0, 0, 0, 0,
	0, 0, 0, 54, 55, 56, 57, 58, 39, 30,
	31, 32, 80, 0, 51, 0, 53, 0, 0, 78,
	76, 0, 0, 22, 23, 29, 0, 0, 33, 0,
	0, 0, 0, 0, 40, 41, 0, 38, 42, 0,
	0, 172, 39, 30, 31, 32, 24, 28, 0, 0,
	0, 35, 0, 0, 36, 0, 34, 0, 0, 25,
	26, 27, 0, 0, 0, 0, 0, 0, 40, 41,
	0, 38, 42, 0, 0, 125, 0, 22, 23, 29,
	24, 28, 33, 0, 0, 35, 0, 0, 36, 0,
	34, 0, 0, 25, 26, 27, 39, 30, 31, 32,
	0, 0, 0, 0, 0, 0, 0, 0, 251, 23,
	29, 0, 0, 33, 0, 0, 0, 0, 0, 0,
	0, 0, 40, 41, 0, 38, 42, 39, 30, 31,
	32, 0, 0, 0, 24, 28, 0, 0, 0, 35,
	0, 0, 36, 0, 34, 0, 0, 25, 26, 27,
	0, 0, 0, 40, 41, 0, 38, 42, 0, 0,
	0, 0, 242, 23, 29, 24, 28, 33, 0, 0,
	35, 0, 0, 36, 0, 34, 0, 0, 25, 26,
	27, 39, 30, 31, 32, 0, 0, 0, 0, 0,
	0, 0, 0, 106, 23, 29, 0, 0, 33, 0,
	0, 0, 0, 0, 0, 0, 0, 40, 41, 0,
	38, 42, 39, 30, 31, 32, 0, 0, 0, 24,
	28, 0, 0, 0, 35, 0, 0, 36, 0, 34,
	0, 0, 25, 26, 27, 0, 0, 0, 40, 41,
	0, 38, 42, 0, 0, 0, 0, 104, 23, 29,
	24, 28, 33, 0, 0, 35, 0, 0, 36, 0,
	34, 0, 0, 25, 26, 27, 39, 30, 31, 32,
	0, 0, 0, 0, 0, 0, 0, 0, 97, 23,
	29, 0, 0, 33, 0, 0, 0, 0, 0, 0,
	0, 0, 40, 41, 0, 38, 42, 39, 30, 31,
	32, 0, 0, 0, 24, 28, 0, 0, 0, 35,
	0, 0, 36, 0, 34, 0, 0, 25, 26, 27,
	0, 0, 0, 40, 41, 0, 38, 42, 0, 0,
	0, 0, 93, 23, 29, 24, 28, 33, 0, 0,
	35, 0, 0, 36, 0, 34, 0, 0, 25, 26,
	27, 39, 30, 31, 32, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 40, 41, 0,
	38, 42, 68, 69, 70, 71, 72, 73, 0, 24,
	28, 0, 59, 0, 89, 0, 0, 36, 0, 34,
	0, 81, 25, 26, 27, 0, 0, 0, 0, 0,
	0, 0, 0, 56, 57, 58, 0, 0, 0, 0,
	80, 0, 51, 0, 53, 0, 0, 78, 76,
}
var yyPact = [...]int{

	-61, -1000, 2295, -61, -61, -1000, -1000, -1000, -1000, 252,
	1838, 175, -1000, -1000, 2653, 2653, 259, 234, 2908, 143,
	2653, 2854, -33, -1000, 2653, 2653, 2653, 2823, 2769, -1000,
	-1000, -1000, -1000, 71, -61, -61, 2653, -1000, 25, -15,
	-19, 2653, -20, 164, 2653, -1000, 597, -1000, 158, -1000,
	2653, 2599, 2653, 258, 2653, 2653, 2653, 2653, 2653, 2653,
	2653, 2653, 2653, 2653, 2653, 2653, 2653, 2653, 2653, 2653,
	2653, 2653, 2653, 2653, -1000, -1000, 2653, 2653, 2653, 2653,
	2653, 2653, 2653, 2653, 163, 1900, 1900, 141, 174, -61,
	167, 28, 1776, -33, 173, -61, 1714, -30, 2086, 2565,
	2653, 201, 201, 201, -33, 2024, -33, 1962, 252, -29,
	2653, 243, 1652, 2653, 251, 76, 1900, 2653, -61, 1590,
	-1000, 2653, -61, 1900, 523, 2653, 1528, -1000, 2931, 2931,
	201, 201, 201, 1900, 2523, 2523, 2403, 2403, 2523, 2523,
	2523, 2523, 1900, 1900, 1900, 1900, 1900, 1900, 1900, 2221,
	1900, 2356, 153, 1900, 2356, -1000, 1900, -61, -61, 2653,
	-61, 110, 2160, 2653, 2653, -61, 2653, 109, -61, 2653,
	2653, 461, 2653, 144, 255, 254, 138, 252, 33, -34,
	-1000, 170, -1000, 1466, -38, -1000, 251, 44, 250, -62,
	399, 2738, -61, -1000, 249, 2477, -1000, 1404, 2653, -24,
	-1000, 2684, 108, 1342, 103, -1000, 170, 1280, 1218, 101,
	-1000, 218, 92, 199, 128, 112, 2445, -1000, 1156, -27,
	-1000, -1000, -1000, 139, -46, 85, -61, -63, -61, 99,
	2653, -1000, 248, -1000, 40, -64, 19, 156, -1000, -1000,
	2653, 1900, -33, 95, -1000, 1094, -1000, -1000, 1900, -1000,
	1900, -33, -1000, -61, -1000, -61, 2653, -1000, 145, -1000,
	-1000, -1000, -1000, 2653, 169, -1000, -1000, -1000, -50, -1000,
	-55, -1000, 1032, -1000, -1000, -1000, -61, 137, 131, -56,
	82, -1000, 117, -1000, 1900, -1000, -1000, 2653, -1000, -1000,
	2653, 2653, 970, -1000, -1000, 83, 78, 908, 125, -61,
	846, 166, -61, -1000, -1000, -1000, 74, -61, -61, 120,
	-1000, -1000, -1000, 784, 337, 722, -1000, -1000, -1000, -61,
	-61, 72, -61, -61, -61, -1000, 63, 61, -61, -1000,
	-1000, 2653, -1000, 58, 45, 200, -61, -61, -1000, -1000,
	-1000, 37, 660, -1000, 198, 115, -1000, -1000, -1000, -1000,
	114, -61, -61, 35, 31, -1000, -1000,
}
var yyPgo = [...]int{

	0, 13, 282, 246, 279, 5, 4, 273, 0, 18,
	11, 271, 1, 268, 12, 6, 266, 265, 2, 81,
	276, 238,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	4, 4, 7, 7, 7, 7, 7, 7, 7, 6,
	18, 5, 16, 16, 16, 12, 13, 13, 13, 14,
	14, 14, 15, 15, 11, 10, 10, 10, 9, 9,
	9, 9, 17, 17, 17, 17, 17, 17, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 19, 19,
	20, 20, 21, 21,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 3, 3, 1,
	1, 2, 2, 5, 1, 4, 7, 9, 5, 13,
	12, 9, 8, 5, 6, 5, 6, 5, 1, 7,
	5, 5, 0, 2, 2, 2, 2, 2, 2, 5,
	5, 4, 0, 2, 3, 3, 0, 1, 4, 0,
	1, 4, 1, 3, 3, 1, 4, 4, 0, 1,
	4, 4, 6, 5, 5, 6, 5, 5, 1, 1,
	2, 2, 2, 2, 4, 2, 4, 1, 1, 1,
	1, 5, 3, 7, 8, 8, 9, 5, 6, 5,
	6, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 3, 3, 3, 3, 5, 4,
	5, 4, 4, 4, 1, 4, 4, 5, 7, 5,
	7, 9, 7, 3, 2, 4, 6, 3, 0, 1,
	1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -19, -2, -20, -21, 68, 78, -3, 11,
	-8, -10, 37, 38, 10, 12, 27, -4, 15, 28,
	44, 47, 4, 5, 61, 74, 75, 76, 62, 6,
	24, 25, 26, 9, 71, 66, 69, -17, 52, 23,
	49, 50, 53, -9, 13, -19, -20, -21, -14, 4,
	54, 71, 55, 73, 60, 61, 62, 63, 64, 41,
	42, 43, 17, 18, 58, 19, 59, 20, 31, 32,
	33, 34, 35, 36, 39, 40, 77, 21, 76, 22,
	69, 50, 16, 54, -9, -8, -8, 4, 14, 66,
	-14, -11, -8, 4, -10, 66, -8, 4, -8, 71,
	69, -8, -8, -8, 4, -8, 4, -8, 69, 4,
	-19, -19, -8, 69, 69, 69, -8, 69, 57, -8,
	-3, 54, 57, -8, -8, 56, -8, 4, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -9, -8, -8, -10, -8, 57, 66, 13,
	66, -1, -19, 16, 68, 66, 54, -1, 66, 69,
	69, -8, 56, -9, 73, 73, -14, 69, -9, -13,
	-12, 6, 70, -8, -15, 4, 48, -16, 51, 71,
	-8, -19, 66, -10, -19, 56, 72, -8, 56, 8,
	70, -19, -1, -8, -1, 67, 6, -8, -8, -1,
	-10, 67, -7, -19, -9, -9, 56, 72, -8, 8,
	70, 4, 4, 70, 8, -14, 57, -19, 57, -19,
	56, 70, 73, 70, -15, 71, -15, 4, 72, 70,
	57, -8, 4, -1, 4, -8, 72, 72, -8, 70,
	-8, 4, 67, 66, 67, 66, 68, 67, 29, 67,
	-6, -18, -5, 45, 46, -6, -5, -18, 8, 70,
	8, 70, -8, 72, 72, 70, 66, 70, 70, 8,
	-19, 72, -19, 67, -8, 4, 70, 57, 72, 70,
	57, 57, -8, 67, 72, -1, -1, -8, 4, 66,
	-8, -10, 56, 70, 70, 72, -1, 66, 66, 70,
	72, -12, 67, -8, -8, -8, 70, 67, 67, 66,
	66, -1, 56, 56, -19, 67, -1, -1, 66, 70,
	70, 57, 70, -1, -1, 67, -19, -19, -1, 67,
	67, -1, -8, 67, 67, 30, -1, -1, 67, 70,
	30, 66, 66, -1, -1, 67, 67,
}
var yyDef = [...]int{

	-2, -2, -2, 138, 139, 140, 142, 143, 4, 49,
	-2, 0, 9, 10, 58, 0, 0, 14, 49, 0,
	0, 0, 68, 69, 0, 0, 0, 0, 0, 77,
	78, 79, 80, 0, 138, 138, 0, 124, 0, 0,
	0, 0, 0, 0, 0, 2, -2, 141, 0, 50,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 112, 113, 0, 0, 0, 0,
	58, 0, 0, 58, 11, 59, 12, 0, 0, -2,
	0, 0, -2, -2, 0, -2, 0, 68, 0, 0,
	58, 70, 71, 72, -2, 0, -2, 0, 49, 0,
	58, 46, 0, 0, 0, 42, 134, 0, 138, 0,
	5, 58, 138, 7, 0, 0, 0, 82, 92, 93,
	94, 95, 96, 97, 98, 99, -2, -2, 102, 103,
	104, 105, 106, 107, 108, 109, 110, 111, 114, 115,
	116, 117, 0, 133, 137, 8, -2, 138, -2, 0,
	-2, 0, -2, 0, 0, -2, 58, 0, 32, 58,
	58, 0, 0, 0, 0, 0, 0, 49, 138, 138,
	47, 0, 91, 0, 0, 52, 0, 0, 0, 0,
	0, 0, -2, 6, 0, 0, 123, 0, 0, 0,
	121, 0, 0, 0, 0, 15, 77, 0, 0, 0,
	54, 0, 0, 0, 0, 0, 0, 122, 0, 0,
	119, 74, 76, 0, 0, 0, 138, 0, 138, 0,
	0, 125, 0, 126, 0, 0, 0, 0, 43, 135,
	0, -2, -2, 0, 51, 0, 66, 67, 81, 120,
	60, -2, 13, -2, 30, -2, 0, 18, 0, 23,
	35, 37, 38, 58, 0, 33, 34, 36, 0, -2,
	0, -2, 0, 63, 64, 118, -2, 0, 0, 0,
	0, 87, 0, 89, 45, 53, 127, 0, 44, 129,
	0, 0, 0, 31, 65, 0, 0, 0, 0, -2,
	59, 0, 138, -2, -2, 62, 0, -2, -2, 0,
	88, 48, 90, 0, 0, 0, 136, 29, 16, -2,
	-2, 0, 138, 138, -2, 83, 0, 0, -2, 128,
	130, 0, 132, 0, 0, 22, -2, -2, 41, 84,
	85, 0, 0, 17, 21, 0, 39, 40, 86, 131,
	0, -2, -2, 0, 0, 20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	78, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 74, 3, 3, 3, 64, 76, 3,
	69, 70, 62, 60, 57, 61, 73, 63, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 56, 68,
	59, 54, 58, 55, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 71, 3, 72, 75, 3, 3, 3, 3, 3,
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
		//line parser.go.y:73
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:82
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:89
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:96
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
		//line parser.go.y:107
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:116
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
		//line parser.go.y:128
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:133
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:138
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:148
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:163
		{
			yyVAL.stmt = &ast.ForStmt{Vars: yyDollar[2].expr_idents, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:188
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:198
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:203
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:208
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:213
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].expr.Position())
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:218
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:223
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:231
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:236
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyDollar[4].compstmt...)
			}
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:245
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:251
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:255
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:259
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:263
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:267
		{
			yyVAL.stmt_cases = yyDollar[2].stmt_multi_case
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_multi_case...)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:275
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:292
		{
			var cases = []ast.Stmt{}
			for _, stmt := range yyDollar[2].expr_many {
				cases = append(cases, &ast.CaseStmt{Expr: stmt, Stmts: yyDollar[5].compstmt})
			}
			yyVAL.stmt_multi_case = cases
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:307
		{
			yyVAL.array_count = ast.ArrayCount{Count: 0}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:311
		{
			yyVAL.array_count = ast.ArrayCount{Count: 1}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:315
		{
			yyVAL.array_count.Count = yyVAL.array_count.Count + 1
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:321
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:326
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:330
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:334
		{
			if len(yyDollar[1].expr_pairs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:342
		{
			yyVAL.expr_idents = []string{}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:346
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:350
		{
			if len(yyDollar[1].expr_idents) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:359
		{
			yyVAL.expr_type = yyDollar[1].tok.Lit
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:363
		{
			yyVAL.expr_type = yyVAL.expr_type + "." + yyDollar[3].tok.Lit
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:368
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:374
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:378
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:382
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:387
		{
			yyVAL.exprs = nil
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:391
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:395
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:402
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 62:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:411
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:415
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: nil}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:419
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: nil, End: yyDollar[4].expr}
		}
	case 65:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:423
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:427
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
		}
	case 67:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:431
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:437
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:442
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:447
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:452
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:487
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:492
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:497
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:502
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:507
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:512
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 84:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:517
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 85:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:522
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 86:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:527
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:532
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 88:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:537
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 89:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:542
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
	case 90:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:551
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
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:560
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:565
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:570
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:575
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:580
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:585
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:590
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:595
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:600
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:605
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:610
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:615
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:620
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:625
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:630
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:635
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:640
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:645
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:650
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:655
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:660
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:665
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:670
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:675
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:680
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:685
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:690
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:695
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:700
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:705
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:710
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:715
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:720
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:725
		{
			yyVAL.expr = yyDollar[1].expr_slice
			yyVAL.expr.SetPosition(yyDollar[1].expr_slice.Position())
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:730
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:735
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:740
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 128:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:745
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 129:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:750
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 130:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:755
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 131:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:760
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 132:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:765
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].tok.Lit, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:770
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:775
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:780
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 136:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:785
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr, KeyExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:790
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: &ast.SliceExpr{Value: yyDollar[3].expr, Begin: nil, End: nil}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:803
		{
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:806
		{
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:811
		{
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:814
		{
		}
	}
	goto yystack /* stack new state and value */
}
