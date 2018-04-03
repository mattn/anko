//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:28
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
	expr_type    string
	tok          ast.Token
	term         ast.Token
	terms        ast.Token
	opt_terms    ast.Token
	array_count  ast.ArrayCount
	expr_slice   ast.SliceExpr
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

//line parser.go.y:797

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 135,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	57, 51,
	-2, 1,
	-1, 10,
	57, 52,
	-2, 24,
	-1, 46,
	57, 51,
	-2, 136,
	-1, 89,
	67, 3,
	-2, 135,
	-1, 92,
	57, 52,
	-2, 48,
	-1, 93,
	16, 43,
	57, 43,
	-2, 61,
	-1, 95,
	67, 3,
	-2, 135,
	-1, 102,
	1, 66,
	8, 66,
	45, 66,
	46, 66,
	54, 66,
	56, 66,
	57, 66,
	66, 66,
	67, 66,
	68, 66,
	70, 66,
	76, 66,
	78, 66,
	-2, 61,
	-1, 104,
	1, 68,
	8, 68,
	45, 68,
	46, 68,
	54, 68,
	56, 68,
	57, 68,
	66, 68,
	67, 68,
	68, 68,
	70, 68,
	76, 68,
	78, 68,
	-2, 61,
	-1, 136,
	17, 0,
	18, 0,
	-2, 93,
	-1, 137,
	17, 0,
	18, 0,
	-2, 94,
	-1, 157,
	57, 52,
	-2, 48,
	-1, 159,
	67, 3,
	-2, 135,
	-1, 161,
	67, 3,
	-2, 135,
	-1, 163,
	67, 1,
	-2, 39,
	-1, 166,
	67, 3,
	-2, 135,
	-1, 193,
	67, 3,
	-2, 135,
	-1, 241,
	57, 53,
	-2, 49,
	-1, 242,
	1, 50,
	45, 50,
	46, 50,
	54, 50,
	57, 54,
	67, 50,
	68, 50,
	78, 50,
	-2, 61,
	-1, 251,
	1, 54,
	8, 54,
	45, 54,
	46, 54,
	57, 54,
	67, 54,
	68, 54,
	70, 54,
	76, 54,
	78, 54,
	-2, 61,
	-1, 253,
	67, 3,
	-2, 135,
	-1, 255,
	67, 3,
	-2, 135,
	-1, 270,
	67, 3,
	-2, 135,
	-1, 280,
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
	-1, 282,
	1, 118,
	8, 118,
	45, 118,
	46, 118,
	54, 118,
	56, 118,
	57, 118,
	66, 118,
	67, 118,
	68, 118,
	70, 118,
	76, 118,
	78, 118,
	-2, 116,
	-1, 297,
	67, 3,
	-2, 135,
	-1, 302,
	67, 3,
	-2, 135,
	-1, 303,
	67, 3,
	-2, 135,
	-1, 308,
	1, 113,
	8, 113,
	45, 113,
	46, 113,
	54, 113,
	56, 113,
	57, 113,
	66, 113,
	67, 113,
	68, 113,
	70, 113,
	76, 113,
	78, 113,
	-2, 111,
	-1, 309,
	1, 117,
	8, 117,
	45, 117,
	46, 117,
	54, 117,
	56, 117,
	57, 117,
	66, 117,
	67, 117,
	68, 117,
	70, 117,
	76, 117,
	78, 117,
	-2, 115,
	-1, 316,
	67, 3,
	-2, 135,
	-1, 317,
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
	-1, 345,
	67, 3,
	-2, 135,
	-1, 346,
	67, 3,
	-2, 135,
}

const yyPrivate = 57344

const yyLast = 3078

var yyAct = [...]int{

	85, 179, 260, 10, 261, 6, 185, 43, 37, 233,
	48, 11, 227, 1, 234, 7, 86, 309, 308, 92,
	225, 96, 84, 6, 99, 100, 101, 103, 105, 90,
	94, 6, 281, 7, 304, 288, 110, 112, 271, 176,
	97, 7, 116, 279, 273, 119, 98, 10, 269, 233,
	249, 123, 124, 126, 287, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 138, 139, 140, 141, 142,
	143, 144, 145, 146, 147, 285, 223, 148, 149, 150,
	151, 158, 153, 155, 157, 218, 2, 200, 152, 233,
	45, 154, 158, 122, 284, 156, 97, 117, 170, 107,
	282, 115, 182, 162, 114, 113, 172, 187, 286, 168,
	189, 280, 272, 275, 184, 239, 177, 175, 191, 186,
	108, 109, 157, 262, 263, 122, 198, 180, 190, 165,
	350, 349, 342, 194, 158, 346, 158, 339, 338, 335,
	334, 331, 321, 315, 222, 259, 296, 314, 291, 277,
	257, 254, 252, 219, 212, 201, 206, 160, 345, 324,
	317, 204, 303, 302, 10, 208, 209, 270, 157, 159,
	106, 95, 217, 203, 164, 205, 163, 289, 121, 211,
	210, 122, 158, 118, 236, 299, 229, 224, 307, 167,
	230, 231, 83, 241, 235, 237, 344, 245, 340, 5,
	248, 262, 263, 250, 47, 192, 8, 243, 297, 195,
	161, 258, 88, 180, 283, 122, 266, 264, 244, 265,
	238, 186, 49, 221, 220, 127, 87, 4, 188, 178,
	278, 46, 91, 68, 69, 70, 71, 72, 73, 213,
	17, 290, 3, 59, 0, 202, 47, 0, 0, 0,
	0, 0, 81, 120, 0, 0, 214, 295, 0, 0,
	0, 0, 0, 298, 226, 228, 0, 293, 0, 294,
	0, 51, 0, 53, 0, 250, 78, 80, 306, 76,
	0, 0, 0, 0, 301, 0, 310, 0, 0, 311,
	312, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 318, 274, 0, 276, 0, 322, 323, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 337, 0,
	329, 330, 0, 0, 333, 0, 0, 0, 336, 0,
	0, 0, 0, 0, 0, 0, 341, 0, 0, 82,
	62, 63, 65, 67, 77, 79, 0, 0, 0, 347,
	348, 0, 0, 0, 68, 69, 70, 71, 72, 73,
	0, 0, 74, 75, 59, 60, 61, 0, 0, 0,
	0, 0, 0, 81, 0, 0, 320, 0, 52, 0,
	327, 64, 66, 54, 55, 56, 57, 58, 0, 0,
	0, 0, 51, 0, 53, 0, 332, 78, 80, 326,
	76, 82, 62, 63, 65, 67, 77, 79, 0, 0,
	0, 0, 0, 0, 0, 0, 68, 69, 70, 71,
	72, 73, 0, 0, 74, 75, 59, 60, 61, 0,
	0, 0, 0, 0, 0, 81, 0, 0, 0, 0,
	52, 215, 0, 64, 66, 54, 55, 56, 57, 58,
	0, 0, 0, 0, 51, 216, 53, 0, 0, 78,
	80, 0, 76, 82, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 81, 0, 0,
	0, 0, 52, 196, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 0, 51, 197, 53, 0,
	0, 78, 80, 0, 76, 21, 22, 28, 0, 0,
	32, 14, 9, 15, 44, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 39, 29, 30, 31, 16, 19,
	0, 0, 0, 0, 0, 0, 0, 0, 12, 13,
	0, 0, 0, 0, 0, 20, 0, 0, 36, 0,
	40, 41, 0, 38, 42, 0, 0, 0, 0, 0,
	0, 0, 23, 27, 0, 0, 0, 34, 0, 6,
	33, 0, 0, 24, 25, 26, 35, 0, 0, 7,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	343, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 51, 0, 53, 0, 0,
	78, 80, 328, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 0, 53,
	0, 0, 78, 80, 325, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 319, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 51,
	0, 53, 0, 0, 78, 80, 0, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 316, 0,
	0, 51, 0, 53, 0, 0, 78, 80, 0, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	313, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 51, 300, 53, 0, 0,
	78, 80, 0, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 292, 53,
	0, 0, 78, 80, 0, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 51,
	268, 53, 0, 0, 78, 80, 0, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	256, 51, 0, 53, 0, 0, 78, 80, 0, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	255, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	0, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 253, 0, 0, 51, 0, 53, 0, 0,
	78, 80, 0, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 247, 53,
	0, 0, 78, 80, 0, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 240, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 51,
	0, 53, 0, 0, 78, 80, 0, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 51, 0, 53, 0, 0, 78, 80, 232, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	199, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	0, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 193, 0, 0, 51, 0, 53, 0, 0,
	78, 80, 0, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 0, 53,
	0, 0, 78, 80, 181, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 169, 0, 0, 51,
	0, 53, 0, 0, 78, 80, 0, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 166, 0,
	0, 51, 0, 53, 0, 0, 78, 80, 0, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 50, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	0, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 51, 0, 53, 0, 0,
	78, 80, 0, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 0, 53,
	0, 0, 78, 183, 0, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 51,
	0, 174, 0, 0, 78, 80, 0, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 51, 0, 173, 0, 0, 78, 80, 0, 76,
	82, 62, 63, 65, 67, 0, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	0, 76, 21, 22, 207, 0, 0, 32, 14, 9,
	15, 44, 0, 18, 0, 0, 0, 0, 0, 0,
	0, 39, 29, 30, 31, 16, 19, 0, 0, 0,
	0, 0, 0, 0, 0, 12, 13, 0, 0, 0,
	0, 0, 20, 0, 0, 36, 0, 40, 41, 0,
	38, 42, 0, 0, 0, 0, 0, 0, 0, 23,
	27, 0, 0, 0, 34, 0, 0, 33, 0, 0,
	24, 25, 26, 35, 82, 62, 63, 65, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 0, 53,
	0, 0, 78, 80, 0, 76, 21, 22, 28, 0,
	0, 32, 14, 9, 15, 44, 0, 18, 0, 0,
	0, 0, 0, 0, 0, 39, 29, 30, 31, 16,
	19, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	13, 0, 0, 0, 0, 0, 20, 0, 0, 36,
	0, 40, 41, 0, 38, 42, 0, 0, 0, 0,
	0, 0, 0, 23, 27, 0, 65, 67, 34, 0,
	0, 33, 0, 0, 24, 25, 26, 35, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 251, 22, 28, 0, 81, 32, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 39, 29, 30, 31, 51, 0, 53, 0,
	0, 78, 80, 0, 76, 21, 22, 28, 0, 0,
	32, 0, 0, 0, 0, 0, 36, 0, 40, 41,
	0, 38, 42, 0, 39, 29, 30, 31, 0, 0,
	23, 27, 0, 0, 0, 34, 0, 0, 33, 305,
	0, 24, 25, 26, 35, 0, 0, 0, 36, 0,
	40, 41, 0, 38, 42, 0, 0, 0, 0, 21,
	22, 28, 23, 27, 32, 0, 0, 34, 0, 0,
	33, 267, 0, 24, 25, 26, 35, 0, 39, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 36, 0, 40, 41, 0, 38, 42, 0,
	0, 0, 0, 0, 0, 0, 23, 27, 0, 0,
	0, 34, 0, 0, 33, 246, 0, 24, 25, 26,
	35, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 0, 0, 0, 0, 21, 22, 28, 0,
	81, 32, 0, 0, 0, 0, 0, 0, 0, 0,
	54, 55, 56, 57, 58, 39, 29, 30, 31, 51,
	0, 53, 0, 0, 78, 80, 0, 76, 0, 0,
	21, 22, 28, 0, 0, 32, 0, 0, 0, 36,
	0, 40, 41, 0, 38, 42, 0, 0, 171, 39,
	29, 30, 31, 23, 27, 0, 0, 0, 34, 0,
	0, 33, 0, 0, 24, 25, 26, 35, 0, 0,
	0, 0, 0, 36, 0, 40, 41, 0, 38, 42,
	0, 0, 125, 0, 21, 22, 28, 23, 27, 32,
	0, 0, 34, 0, 0, 33, 0, 0, 24, 25,
	26, 35, 0, 39, 29, 30, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 251, 22, 28, 0, 0,
	32, 0, 0, 0, 0, 0, 0, 36, 0, 40,
	41, 0, 38, 42, 39, 29, 30, 31, 0, 0,
	0, 23, 27, 0, 0, 0, 34, 0, 0, 33,
	0, 0, 24, 25, 26, 35, 0, 0, 36, 0,
	40, 41, 0, 38, 42, 0, 0, 0, 0, 242,
	22, 28, 23, 27, 32, 0, 0, 34, 0, 0,
	33, 0, 0, 24, 25, 26, 35, 0, 39, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	111, 22, 28, 0, 0, 32, 0, 0, 0, 0,
	0, 0, 36, 0, 40, 41, 0, 38, 42, 39,
	29, 30, 31, 0, 0, 0, 23, 27, 0, 0,
	0, 34, 0, 0, 33, 0, 0, 24, 25, 26,
	35, 0, 0, 36, 0, 40, 41, 0, 38, 42,
	0, 0, 0, 0, 104, 22, 28, 23, 27, 32,
	0, 0, 34, 0, 0, 33, 0, 0, 24, 25,
	26, 35, 0, 39, 29, 30, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 102, 22, 28, 0, 0,
	32, 0, 0, 0, 0, 0, 0, 36, 0, 40,
	41, 0, 38, 42, 39, 29, 30, 31, 0, 0,
	0, 23, 27, 0, 0, 0, 34, 0, 0, 33,
	0, 0, 24, 25, 26, 35, 0, 0, 36, 0,
	40, 41, 0, 38, 42, 0, 0, 0, 0, 93,
	22, 28, 23, 27, 32, 0, 0, 34, 0, 0,
	33, 0, 0, 24, 25, 26, 35, 0, 39, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 36, 0, 40, 41, 0, 38, 42, 0,
	0, 0, 0, 0, 0, 0, 23, 27, 0, 0,
	0, 89, 0, 0, 33, 0, 0, 24, 25, 26,
	35, 68, 69, 70, 71, 72, 73, 0, 0, 0,
	0, 59, 0, 0, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 56, 57, 58, 0, 0, 0, 0, 51,
	0, 53, 0, 0, 78, 80, 0, 76,
}
var yyPact = [...]int{

	-63, -1000, 2352, -63, -63, -1000, -1000, -1000, -1000, 218,
	1834, 138, -1000, -1000, 2700, 2700, 222, 198, 2955, 105,
	2700, -29, -1000, 2700, 2700, 2700, 2901, 2870, -1000, -1000,
	-1000, -1000, 95, -63, -63, 2700, 2816, -1000, 30, 29,
	26, 2700, 22, 126, 2700, -1000, 531, -1000, 124, -1000,
	2700, 2646, 2700, 221, 2700, 2700, 2700, 2700, 2700, 2700,
	2700, 2700, 2700, 2700, 2700, 2700, 2700, 2700, 2700, 2700,
	2700, 2700, 2700, 2700, -1000, -1000, 2700, 2700, 2700, 2700,
	2700, 2700, 2700, 2700, 125, 1896, 1896, 103, 144, -63,
	158, 61, 1772, -29, 135, -63, 1710, 2612, 2700, 202,
	202, 202, -29, 2082, -29, 2020, 218, -36, 2700, 207,
	1648, 27, 1958, 2700, 217, 59, 1896, 2700, -63, 1586,
	-1000, 2700, -63, 1896, 457, 2700, 1524, -1000, 3000, 3000,
	202, 202, 202, 1896, 2570, 2570, 2397, 2397, 2570, 2570,
	2570, 2570, 1896, 1896, 1896, 1896, 1896, 1896, 1896, 2144,
	1896, 2278, 79, 1896, -1000, 2278, -1000, 1896, -63, -63,
	2700, -63, 89, 2218, 2700, 2700, -63, 2700, 87, -63,
	395, 2700, 77, 220, 219, 68, 218, -37, -45, -1000,
	130, -1000, 2700, 2700, 1462, -62, -1000, 217, 115, 216,
	45, 1400, 2785, -63, -1000, 214, 2525, -1000, 1338, 2700,
	-26, -1000, 2731, 85, 1276, 84, -1000, 130, 1214, 1152,
	83, -1000, 182, 78, 156, 2471, -1000, 1090, -28, -1000,
	-1000, -1000, 101, -38, 36, -63, 43, -63, 82, 2700,
	35, 24, -1000, 210, -1000, 18, 38, -22, 120, -1000,
	2700, 1896, -29, 81, -1000, 1028, -1000, -1000, 1896, -1000,
	1896, -29, -1000, -63, -1000, -63, 2700, -1000, 142, -1000,
	-1000, -1000, 2700, 129, -1000, -1000, 966, -1000, -1000, -1000,
	-63, 97, 96, -42, 2439, -1000, 121, -1000, 1896, -58,
	-1000, -59, -1000, -1000, -1000, 2700, -1000, -1000, 2700, 2700,
	904, -1000, -1000, 80, 76, 842, 94, -63, 780, -63,
	-1000, 75, -63, -63, 93, -1000, -1000, -1000, -1000, -1000,
	718, 333, 656, -1000, -1000, -1000, -63, -63, 74, -63,
	-63, -1000, 73, 72, -63, -1000, -1000, 2700, -1000, 71,
	70, 168, -63, -1000, -1000, -1000, 65, 594, -1000, 166,
	92, -1000, -1000, -1000, 69, -63, -63, 64, 63, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 13, 242, 206, 240, 4, 2, 239, 0, 7,
	11, 232, 1, 229, 10, 6, 228, 8, 86, 227,
	199,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 7, 7,
	7, 7, 7, 6, 5, 16, 16, 16, 12, 13,
	13, 13, 14, 14, 14, 15, 15, 11, 10, 10,
	10, 9, 9, 9, 9, 17, 17, 17, 17, 17,
	17, 8, 8, 8, 8, 8, 8, 8, 8, 8,
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
	2, 2, 2, 5, 4, 0, 2, 3, 3, 0,
	1, 4, 0, 1, 4, 1, 3, 3, 1, 4,
	4, 0, 1, 4, 4, 6, 5, 5, 6, 5,
	5, 1, 1, 2, 2, 2, 2, 4, 2, 4,
	1, 1, 1, 1, 5, 3, 7, 8, 8, 9,
	5, 6, 5, 6, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 3, 3, 3,
	3, 5, 4, 6, 5, 5, 4, 6, 5, 4,
	4, 1, 4, 4, 5, 7, 5, 7, 9, 7,
	3, 2, 6, 3, 3, 0, 1, 1, 2, 1,
	1,
}
var yyChk = [...]int{

	-1000, -1, -18, -2, -19, -20, 68, 78, -3, 11,
	-8, -10, 37, 38, 10, 12, 27, -4, 15, 28,
	44, 4, 5, 61, 72, 73, 74, 62, 6, 24,
	25, 26, 9, 69, 66, 75, 47, -17, 52, 23,
	49, 50, 53, -9, 13, -18, -19, -20, -14, 4,
	54, 69, 55, 71, 60, 61, 62, 63, 64, 41,
	42, 43, 17, 18, 58, 19, 59, 20, 31, 32,
	33, 34, 35, 36, 39, 40, 77, 21, 74, 22,
	75, 50, 16, 54, -9, -8, -8, 4, 14, 66,
	-14, -11, -8, 4, -10, 66, -8, 69, 75, -8,
	-8, -8, 4, -8, 4, -8, 75, 4, -18, -18,
	-8, 4, -8, 75, 75, 75, -8, 75, 57, -8,
	-3, 54, 57, -8, -8, 56, -8, 4, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -9, -8, -17, -8, -10, -8, 57, 66,
	13, 66, -1, -18, 16, 68, 66, 54, -1, 66,
	-8, 56, -9, 71, 71, -14, 75, -9, -13, -12,
	6, 76, 75, 75, -8, -15, 4, 48, -16, 51,
	69, -8, -18, 66, -10, -18, 56, 70, -8, 56,
	8, 76, -18, -1, -8, -1, 67, 6, -8, -8,
	-1, -10, 67, -7, -18, 56, 70, -8, 8, 76,
	4, 4, 76, 8, -14, 57, -18, 57, -18, 56,
	-9, -9, 76, 71, 76, -15, 69, -15, 4, 70,
	57, -8, 4, -1, 4, -8, 70, 70, -8, 76,
	-8, 4, 67, 66, 67, 66, 68, 67, 29, 67,
	-6, -5, 45, 46, -6, -5, -8, 70, 70, 76,
	66, 76, 76, 8, -18, 70, -18, 67, -8, 8,
	76, 8, 76, 4, 76, 57, 70, 76, 57, 57,
	-8, 67, 70, -1, -1, -8, 4, 66, -8, 56,
	70, -1, 66, 66, 76, 70, -12, 67, 76, 76,
	-8, -8, -8, 76, 67, 67, 66, 66, -1, 56,
	-18, 67, -1, -1, 66, 76, 76, 57, 76, -1,
	-1, 67, -18, -1, 67, 67, -1, -8, 67, 67,
	30, -1, 67, 76, 30, 66, 66, -1, -1, 67,
	67,
}
var yyDef = [...]int{

	-2, -2, -2, 135, 136, 137, 139, 140, 4, 42,
	-2, 0, 9, 10, 51, 0, 0, 14, 42, 0,
	0, 61, 62, 0, 0, 0, 0, 0, 70, 71,
	72, 73, 0, 135, 135, 0, 0, 121, 0, 0,
	0, 0, 0, 0, 0, 2, -2, 138, 0, 43,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 105, 106, 0, 0, 0, 0,
	51, 0, 0, 51, 11, 52, 12, 0, 0, -2,
	0, 0, -2, -2, 0, -2, 0, 0, 51, 63,
	64, 65, -2, 0, -2, 0, 42, 0, 51, 39,
	0, 61, 0, 0, 0, 35, 131, 0, 135, 0,
	5, 51, 135, 7, 0, 0, 0, 75, 85, 86,
	87, 88, 89, 90, 91, 92, -2, -2, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 107, 108,
	109, 110, 0, 130, 121, 134, 8, -2, 135, -2,
	0, -2, 0, -2, 0, 0, -2, 51, 0, 28,
	0, 0, 0, 0, 0, 0, 42, 135, 135, 40,
	0, 84, 51, 51, 0, 0, 45, 0, 0, 0,
	0, 0, 0, -2, 6, 0, 0, 120, 0, 0,
	0, 116, 0, 0, 0, 0, 15, 70, 0, 0,
	0, 47, 0, 0, 0, 0, 119, 0, 0, 112,
	67, 69, 0, 0, 0, 135, 0, 135, 0, 0,
	0, 0, 122, 0, 123, 0, 0, 0, 0, 36,
	0, -2, -2, 0, 44, 0, 59, 60, 74, 115,
	53, -2, 13, -2, 26, -2, 0, 18, 0, 23,
	31, 32, 0, 0, 29, 30, 0, 56, 57, 111,
	-2, 0, 0, 0, 0, 80, 0, 82, 38, 0,
	-2, 0, -2, 46, 124, 0, 37, 126, 0, 0,
	0, 27, 58, 0, 0, 0, 0, -2, 0, 135,
	55, 0, -2, -2, 0, 81, 41, 83, -2, -2,
	0, 0, 0, 132, 25, 16, -2, -2, 0, 135,
	-2, 76, 0, 0, -2, 125, 127, 0, 129, 0,
	0, 22, -2, 34, 77, 78, 0, 0, 17, 21,
	0, 33, 79, 128, 0, -2, -2, 0, 0, 20,
	19,
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
		//line parser.go.y:71
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:75
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:80
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:87
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:94
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
		//line parser.go.y:105
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:110
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:114
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
		//line parser.go.y:126
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:141
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:146
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:151
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:156
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.stmt = &ast.ForStmt{Vars: yyDollar[2].expr_idents, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:166
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:171
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:176
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:186
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:191
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:196
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:201
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:209
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:214
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
		//line parser.go.y:223
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:229
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:233
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:237
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:241
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:245
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:262
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:267
		{
			yyVAL.array_count = ast.ArrayCount{Count: 0}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.array_count = ast.ArrayCount{Count: 1}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:275
		{
			yyVAL.array_count.Count = yyVAL.array_count.Count + 1
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:281
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:290
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:294
		{
			if len(yyDollar[1].expr_pairs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.expr_idents = []string{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:310
		{
			if len(yyDollar[1].expr_idents) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:319
		{
			yyVAL.expr_type = yyDollar[1].tok.Lit
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:323
		{
			yyVAL.expr_type = yyVAL.expr_type + "." + yyDollar[3].tok.Lit
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:328
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:334
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:338
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:342
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:347
		{
			yyVAL.exprs = nil
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:351
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:355
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:362
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:371
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 56:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:375
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: nil}
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:379
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: nil, End: yyDollar[4].expr}
		}
	case 58:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:383
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:387
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
		}
	case 60:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:391
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:402
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:407
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:412
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:417
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:422
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:432
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:437
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:442
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:447
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:452
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 76:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 77:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 78:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 79:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:487
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:492
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 81:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:497
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
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
	case 83:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:511
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
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:520
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:525
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:530
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:535
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:540
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:545
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:550
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:555
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:560
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:565
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:570
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:575
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:580
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:585
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:590
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:595
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:600
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:605
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:610
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:615
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:620
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:625
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:630
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:635
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:640
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:645
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:650
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:655
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:660
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:665
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:670
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:675
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:680
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:685
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:690
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:695
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:700
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:705
		{
			yyVAL.expr = &yyDollar[1].expr_slice
			yyVAL.expr.SetPosition(yyDollar[1].expr_slice.Position())
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:710
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:715
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:720
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:725
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:730
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:735
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 128:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:740
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 129:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:745
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].tok.Lit, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:750
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:755
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 132:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:760
		{
			yyVAL.expr = &ast.DeleteExpr{MapExpr: yyDollar[3].expr, KeyExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:765
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: yyDollar[3].expr_slice}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:770
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: ast.SliceExpr{Value: yyDollar[3].expr, Begin: nil, End: nil}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:783
		{
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:786
		{
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:791
		{
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:794
		{
		}
	}
	goto yystack /* stack new state and value */
}
