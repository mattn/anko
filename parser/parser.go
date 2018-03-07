//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:26
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
	array_count  ast.ArrayCount
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
const UNARY = 57395

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
	"'!'",
	"'^'",
	"'&'",
	"'.'",
	"'('",
	"')'",
	"'|'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:745

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 129,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	56, 49,
	-2, 1,
	-1, 10,
	56, 50,
	-2, 24,
	-1, 44,
	56, 49,
	-2, 130,
	-1, 86,
	66, 3,
	-2, 129,
	-1, 89,
	56, 50,
	-2, 46,
	-1, 90,
	16, 43,
	56, 43,
	-2, 53,
	-1, 92,
	66, 3,
	-2, 129,
	-1, 99,
	1, 58,
	8, 58,
	45, 58,
	46, 58,
	53, 58,
	55, 58,
	56, 58,
	65, 58,
	66, 58,
	67, 58,
	69, 58,
	75, 58,
	77, 58,
	-2, 53,
	-1, 101,
	1, 60,
	8, 60,
	45, 60,
	46, 60,
	53, 60,
	55, 60,
	56, 60,
	65, 60,
	66, 60,
	67, 60,
	69, 60,
	75, 60,
	77, 60,
	-2, 53,
	-1, 130,
	17, 0,
	18, 0,
	-2, 85,
	-1, 131,
	17, 0,
	18, 0,
	-2, 86,
	-1, 151,
	56, 50,
	-2, 46,
	-1, 153,
	66, 3,
	-2, 129,
	-1, 155,
	66, 3,
	-2, 129,
	-1, 157,
	66, 1,
	-2, 39,
	-1, 160,
	66, 3,
	-2, 129,
	-1, 185,
	66, 3,
	-2, 129,
	-1, 231,
	56, 51,
	-2, 47,
	-1, 232,
	1, 48,
	45, 48,
	46, 48,
	53, 48,
	56, 52,
	66, 48,
	67, 48,
	77, 48,
	-2, 53,
	-1, 241,
	1, 52,
	8, 52,
	45, 52,
	46, 52,
	56, 52,
	66, 52,
	67, 52,
	69, 52,
	75, 52,
	77, 52,
	-2, 53,
	-1, 243,
	66, 3,
	-2, 129,
	-1, 245,
	66, 3,
	-2, 129,
	-1, 260,
	66, 3,
	-2, 129,
	-1, 270,
	1, 106,
	8, 106,
	45, 106,
	46, 106,
	53, 106,
	55, 106,
	56, 106,
	65, 106,
	66, 106,
	67, 106,
	69, 106,
	75, 106,
	77, 106,
	-2, 104,
	-1, 272,
	1, 110,
	8, 110,
	45, 110,
	46, 110,
	53, 110,
	55, 110,
	56, 110,
	65, 110,
	66, 110,
	67, 110,
	69, 110,
	75, 110,
	77, 110,
	-2, 108,
	-1, 285,
	66, 3,
	-2, 129,
	-1, 290,
	66, 3,
	-2, 129,
	-1, 291,
	66, 3,
	-2, 129,
	-1, 296,
	1, 105,
	8, 105,
	45, 105,
	46, 105,
	53, 105,
	55, 105,
	56, 105,
	65, 105,
	66, 105,
	67, 105,
	69, 105,
	75, 105,
	77, 105,
	-2, 103,
	-1, 297,
	1, 109,
	8, 109,
	45, 109,
	46, 109,
	53, 109,
	55, 109,
	56, 109,
	65, 109,
	66, 109,
	67, 109,
	69, 109,
	75, 109,
	77, 109,
	-2, 107,
	-1, 303,
	66, 3,
	-2, 129,
	-1, 304,
	66, 3,
	-2, 129,
	-1, 307,
	45, 3,
	46, 3,
	66, 3,
	-2, 129,
	-1, 311,
	66, 3,
	-2, 129,
	-1, 319,
	45, 3,
	46, 3,
	66, 3,
	-2, 129,
	-1, 332,
	66, 3,
	-2, 129,
	-1, 333,
	66, 3,
	-2, 129,
}

const yyPrivate = 57344

const yyLast = 3109

var yyAct = [...]int{

	82, 173, 250, 10, 46, 251, 297, 6, 6, 275,
	296, 292, 1, 261, 41, 265, 83, 7, 7, 89,
	256, 93, 236, 87, 96, 97, 98, 100, 102, 81,
	271, 104, 11, 170, 112, 111, 107, 109, 95, 269,
	263, 113, 95, 115, 94, 10, 110, 215, 176, 119,
	120, 91, 122, 123, 124, 125, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	140, 141, 219, 217, 142, 143, 144, 145, 152, 147,
	149, 151, 2, 6, 6, 207, 43, 152, 118, 230,
	159, 174, 146, 7, 7, 118, 165, 272, 189, 156,
	337, 103, 252, 253, 336, 162, 270, 262, 169, 164,
	329, 178, 179, 150, 214, 326, 105, 106, 151, 325,
	171, 180, 322, 249, 182, 321, 318, 308, 302, 301,
	279, 267, 247, 152, 244, 65, 66, 67, 68, 69,
	70, 183, 284, 152, 242, 56, 152, 204, 198, 193,
	186, 295, 208, 154, 79, 196, 333, 332, 10, 200,
	201, 311, 151, 304, 291, 190, 195, 211, 197, 157,
	290, 260, 78, 202, 153, 216, 75, 50, 77, 92,
	73, 226, 228, 229, 158, 231, 114, 117, 287, 235,
	118, 222, 223, 237, 203, 240, 221, 184, 233, 161,
	80, 187, 331, 285, 8, 155, 252, 253, 327, 254,
	248, 257, 255, 85, 174, 234, 47, 213, 5, 212,
	121, 84, 268, 45, 118, 181, 172, 88, 205, 65,
	66, 67, 68, 69, 70, 194, 4, 17, 3, 56,
	44, 0, 0, 0, 0, 0, 206, 283, 79, 116,
	0, 0, 0, 286, 218, 220, 281, 0, 282, 53,
	54, 55, 0, 45, 0, 240, 78, 0, 294, 0,
	75, 50, 77, 289, 73, 298, 0, 0, 299, 300,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 305, 0,
	264, 0, 266, 309, 310, 0, 0, 0, 0, 0,
	105, 0, 0, 0, 0, 324, 316, 317, 0, 0,
	320, 0, 0, 0, 323, 0, 0, 0, 0, 0,
	0, 0, 328, 0, 0, 59, 60, 62, 64, 74,
	76, 0, 0, 0, 0, 334, 335, 0, 0, 65,
	66, 67, 68, 69, 70, 0, 0, 71, 72, 56,
	57, 58, 0, 0, 0, 0, 0, 0, 79, 0,
	307, 0, 49, 0, 314, 61, 63, 51, 52, 53,
	54, 55, 0, 0, 0, 0, 78, 0, 0, 319,
	75, 50, 77, 313, 73, 59, 60, 62, 64, 74,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	66, 67, 68, 69, 70, 0, 0, 71, 72, 56,
	57, 58, 0, 0, 0, 0, 0, 0, 79, 0,
	0, 0, 49, 0, 277, 61, 63, 51, 52, 53,
	54, 55, 0, 0, 0, 0, 78, 0, 0, 0,
	75, 50, 77, 276, 73, 59, 60, 62, 64, 74,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	66, 67, 68, 69, 70, 0, 0, 71, 72, 56,
	57, 58, 0, 0, 0, 0, 0, 0, 79, 0,
	0, 0, 49, 0, 274, 61, 63, 51, 52, 53,
	54, 55, 0, 0, 0, 0, 78, 0, 0, 0,
	75, 50, 77, 273, 73, 59, 60, 62, 64, 74,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	66, 67, 68, 69, 70, 0, 0, 71, 72, 56,
	57, 58, 0, 0, 0, 0, 0, 0, 79, 0,
	0, 0, 49, 210, 0, 61, 63, 51, 52, 53,
	54, 55, 0, 0, 0, 0, 78, 209, 0, 0,
	75, 50, 77, 0, 73, 59, 60, 62, 64, 74,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 65,
	66, 67, 68, 69, 70, 0, 0, 71, 72, 56,
	57, 58, 0, 0, 0, 0, 0, 0, 79, 0,
	0, 0, 49, 192, 0, 61, 63, 51, 52, 53,
	54, 55, 0, 0, 0, 0, 78, 191, 0, 0,
	75, 50, 77, 0, 73, 21, 22, 28, 0, 0,
	32, 14, 9, 15, 42, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 38, 29, 30, 31, 16, 19,
	0, 0, 0, 0, 0, 0, 0, 0, 12, 13,
	0, 0, 0, 0, 0, 20, 0, 0, 36, 0,
	39, 40, 0, 37, 0, 0, 0, 0, 0, 0,
	0, 23, 27, 0, 0, 0, 34, 0, 6, 33,
	0, 24, 25, 26, 0, 35, 0, 0, 7, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 330, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 315, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 312, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 306, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 303, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 288, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 280, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 278, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 259, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 246,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 245, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 243, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 239, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 225, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 224, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 188, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 185, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 175, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 163, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 160, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 48, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 50, 177, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 0, 0, 0, 0,
	78, 0, 0, 0, 75, 168, 77, 0, 73, 59,
	60, 62, 64, 74, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 67, 68, 69, 70, 0,
	0, 71, 72, 56, 57, 58, 0, 0, 0, 0,
	0, 0, 79, 0, 0, 0, 49, 0, 0, 61,
	63, 51, 52, 53, 54, 55, 59, 60, 62, 64,
	78, 76, 0, 0, 75, 167, 77, 0, 73, 0,
	65, 66, 67, 68, 69, 70, 0, 0, 71, 72,
	56, 57, 58, 0, 0, 0, 0, 0, 0, 79,
	0, 0, 0, 0, 0, 0, 61, 63, 51, 52,
	53, 54, 55, 0, 0, 0, 0, 78, 0, 0,
	0, 75, 50, 77, 0, 73, 21, 22, 199, 0,
	0, 32, 14, 9, 15, 42, 0, 18, 0, 0,
	0, 0, 0, 0, 0, 38, 29, 30, 31, 16,
	19, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	13, 0, 0, 0, 0, 0, 20, 0, 0, 36,
	0, 39, 40, 0, 37, 0, 0, 0, 0, 0,
	0, 0, 23, 27, 0, 0, 0, 34, 0, 0,
	33, 0, 24, 25, 26, 0, 35, 59, 60, 62,
	64, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 65, 66, 67, 68, 69, 70, 0, 0, 71,
	72, 56, 57, 58, 0, 0, 0, 0, 0, 0,
	79, 0, 0, 0, 0, 0, 0, 61, 63, 51,
	52, 53, 54, 55, 0, 0, 0, 0, 78, 0,
	0, 0, 75, 50, 77, 0, 73, 21, 22, 28,
	0, 0, 32, 14, 9, 15, 42, 0, 18, 0,
	0, 0, 0, 0, 0, 0, 38, 29, 30, 31,
	16, 19, 0, 0, 0, 0, 0, 0, 0, 0,
	12, 13, 0, 0, 0, 0, 0, 20, 0, 0,
	36, 0, 39, 40, 0, 37, 0, 0, 0, 0,
	0, 0, 0, 23, 27, 0, 62, 64, 34, 0,
	0, 33, 0, 24, 25, 26, 0, 35, 65, 66,
	67, 68, 69, 70, 0, 0, 71, 72, 56, 57,
	58, 0, 0, 0, 0, 0, 0, 79, 0, 0,
	0, 0, 0, 0, 61, 63, 51, 52, 53, 54,
	55, 0, 0, 0, 0, 78, 0, 0, 0, 75,
	50, 77, 0, 73, 65, 66, 67, 68, 69, 70,
	0, 0, 71, 72, 56, 0, 0, 0, 0, 0,
	0, 0, 0, 79, 0, 0, 0, 0, 0, 0,
	0, 0, 51, 52, 53, 54, 55, 241, 22, 28,
	0, 78, 32, 0, 0, 75, 50, 77, 0, 73,
	0, 0, 0, 0, 0, 0, 38, 29, 30, 31,
	0, 0, 0, 0, 0, 0, 0, 21, 22, 28,
	0, 0, 32, 0, 0, 0, 0, 0, 0, 0,
	36, 0, 39, 40, 0, 37, 38, 29, 30, 31,
	0, 0, 0, 23, 27, 0, 0, 0, 34, 0,
	0, 33, 293, 24, 25, 26, 0, 35, 0, 0,
	36, 0, 39, 40, 0, 37, 0, 0, 0, 0,
	21, 22, 28, 23, 27, 32, 0, 0, 34, 0,
	0, 33, 258, 24, 25, 26, 0, 35, 0, 38,
	29, 30, 31, 0, 0, 0, 0, 0, 0, 0,
	21, 22, 28, 0, 0, 32, 0, 0, 0, 0,
	0, 0, 0, 36, 0, 39, 40, 0, 37, 38,
	29, 30, 31, 0, 0, 0, 23, 27, 0, 0,
	0, 34, 0, 0, 33, 238, 24, 25, 26, 0,
	35, 0, 0, 36, 0, 39, 40, 0, 37, 0,
	0, 166, 0, 21, 22, 28, 23, 27, 32, 0,
	0, 34, 0, 0, 33, 0, 24, 25, 26, 0,
	35, 0, 38, 29, 30, 31, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 21, 22, 28, 0,
	0, 32, 0, 0, 0, 0, 36, 0, 39, 40,
	0, 37, 0, 0, 148, 38, 29, 30, 31, 23,
	27, 0, 0, 0, 34, 0, 0, 33, 0, 24,
	25, 26, 0, 35, 0, 0, 0, 0, 0, 36,
	0, 39, 40, 0, 37, 0, 0, 0, 0, 241,
	22, 28, 23, 27, 32, 0, 0, 34, 0, 0,
	33, 0, 24, 25, 26, 0, 35, 0, 38, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 232,
	22, 28, 0, 0, 32, 0, 0, 0, 0, 0,
	0, 0, 36, 0, 39, 40, 0, 37, 38, 29,
	30, 31, 0, 0, 0, 23, 27, 0, 0, 0,
	34, 0, 0, 33, 0, 24, 25, 26, 0, 35,
	0, 0, 36, 0, 39, 40, 0, 37, 0, 0,
	0, 0, 21, 22, 28, 23, 27, 32, 0, 0,
	34, 0, 0, 33, 0, 24, 25, 26, 0, 35,
	0, 38, 29, 30, 31, 0, 0, 0, 0, 0,
	0, 0, 108, 22, 28, 0, 0, 32, 0, 0,
	0, 0, 0, 0, 0, 36, 0, 39, 40, 0,
	37, 38, 29, 30, 31, 0, 0, 0, 23, 27,
	0, 0, 0, 34, 0, 0, 227, 0, 24, 25,
	26, 0, 35, 0, 0, 36, 0, 39, 40, 0,
	37, 0, 0, 0, 0, 101, 22, 28, 23, 27,
	32, 0, 0, 34, 0, 0, 33, 0, 24, 25,
	26, 0, 35, 0, 38, 29, 30, 31, 0, 0,
	0, 0, 0, 0, 0, 99, 22, 28, 0, 0,
	32, 0, 0, 0, 0, 0, 0, 0, 36, 0,
	39, 40, 0, 37, 38, 29, 30, 31, 0, 0,
	0, 23, 27, 0, 0, 0, 34, 0, 0, 33,
	0, 24, 25, 26, 0, 35, 0, 0, 36, 0,
	39, 40, 0, 37, 0, 0, 0, 0, 90, 22,
	28, 23, 27, 32, 0, 0, 34, 0, 0, 33,
	0, 24, 25, 26, 0, 35, 0, 38, 29, 30,
	31, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 36, 0, 39, 40, 0, 37, 0, 0, 0,
	0, 0, 0, 0, 23, 27, 0, 0, 0, 86,
	0, 0, 33, 0, 24, 25, 26, 0, 35,
}
var yyPact = [...]int{

	-59, -1000, 2383, -59, -59, -1000, -1000, -1000, -1000, 212,
	1892, 147, -1000, -1000, 2732, 2732, 217, 199, 3034, 114,
	2732, -30, -1000, 2732, 2732, 2732, 2981, 2951, -1000, -1000,
	-1000, -1000, 27, -59, -59, 2732, 2898, -28, -39, -40,
	2732, 130, 2732, -1000, 631, -1000, 134, -1000, 2732, 2732,
	216, 2732, 2732, 2732, 2732, 2732, 2732, 2732, 2732, 2732,
	2732, 2732, 2732, 2732, 2732, 2732, 2732, 2732, 2732, 2732,
	2732, -1000, -1000, 2732, 2732, 2732, 2732, 2732, 2699, 2732,
	2732, 87, 1952, 1952, 109, 140, -59, 168, 23, 1832,
	-30, 146, -59, 1772, 2732, 2646, 104, 104, 104, -30,
	2132, -30, 2072, 212, -41, 2732, 208, 1712, -26, 2012,
	2732, 2732, 73, 1952, -59, 1652, -1000, 2732, -59, 1952,
	1592, -1000, 198, 198, 104, 104, 104, 1952, 2473, 2473,
	2427, 2427, 2473, 2473, 2473, 2473, 1952, 1952, 1952, 1952,
	1952, 1952, 1952, 2179, 1952, 2310, 90, 558, 2732, 1952,
	-1000, 1952, -59, -59, 2732, -59, 82, 2252, 2732, 2732,
	-59, 2732, 81, -59, 77, 498, 2732, 215, 213, 39,
	212, 17, 16, -1000, 141, -1000, 2732, 2732, 1532, 1472,
	2732, 2868, 2732, 20, 2815, -59, -1000, 211, 2732, -53,
	-1000, -1000, 2616, 1412, 2785, 78, 1352, 68, -1000, 141,
	1292, 1232, 66, -1000, 181, 57, 161, -55, -1000, -1000,
	2563, 1172, -1000, -1000, 106, -62, 32, -59, -54, -59,
	65, 2732, 31, 22, -1000, -1000, 438, -60, 378, 1112,
	-1000, 1952, -30, 64, -1000, 1952, -1000, 1052, -1000, -1000,
	1952, -30, -1000, -59, -1000, -59, 2732, -1000, 138, -1000,
	-1000, -1000, 2732, 133, -1000, -1000, -1000, 992, -1000, -1000,
	-59, 105, 99, -64, 2533, -1000, 85, -1000, 1952, -65,
	-1000, -69, -1000, -1000, 2732, -1000, -1000, 2732, 2732, -1000,
	-1000, 63, 62, 932, 98, -59, 872, -59, -1000, 61,
	-59, -59, 96, -1000, -1000, -1000, -1000, -1000, 812, 318,
	752, -1000, -1000, -59, -59, 60, -59, -59, -1000, 59,
	56, -59, -1000, -1000, 2732, -1000, 53, 49, 178, -59,
	-1000, -1000, -1000, 44, 692, -1000, 172, 92, -1000, -1000,
	-1000, 91, -59, -59, 38, 34, -1000, -1000,
}
var yyPgo = [...]int{

	0, 12, 238, 204, 237, 5, 2, 228, 0, 14,
	32, 227, 1, 226, 4, 225, 82, 236, 218,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 7, 7,
	7, 7, 7, 6, 5, 15, 15, 15, 12, 13,
	13, 13, 14, 14, 14, 11, 10, 10, 10, 9,
	9, 9, 9, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 16,
	16, 17, 17, 18, 18,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 3, 3, 1,
	1, 2, 2, 5, 1, 4, 7, 9, 5, 13,
	12, 9, 8, 5, 1, 7, 5, 5, 0, 2,
	2, 2, 2, 5, 4, 0, 2, 3, 3, 0,
	1, 4, 0, 1, 4, 3, 1, 4, 4, 0,
	1, 4, 4, 1, 1, 2, 2, 2, 2, 4,
	2, 4, 1, 1, 1, 1, 5, 3, 7, 8,
	8, 9, 5, 6, 5, 6, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 3,
	3, 3, 3, 5, 4, 6, 5, 5, 4, 6,
	5, 4, 4, 6, 5, 5, 6, 5, 5, 4,
	4, 5, 7, 5, 7, 9, 7, 3, 2, 0,
	1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -16, -2, -17, -18, 67, 77, -3, 11,
	-8, -10, 37, 38, 10, 12, 27, -4, 15, 28,
	44, 4, 5, 60, 70, 71, 72, 61, 6, 24,
	25, 26, 9, 68, 65, 74, 47, 52, 23, 49,
	50, -9, 13, -16, -17, -18, -14, 4, 53, 54,
	73, 59, 60, 61, 62, 63, 41, 42, 43, 17,
	18, 57, 19, 58, 20, 31, 32, 33, 34, 35,
	36, 39, 40, 76, 21, 72, 22, 74, 68, 50,
	53, -9, -8, -8, 4, 14, 65, -14, -11, -8,
	4, -10, 65, -8, 74, 68, -8, -8, -8, 4,
	-8, 4, -8, 74, 4, -16, -16, -8, 4, -8,
	74, 74, 74, -8, 56, -8, -3, 53, 56, -8,
	-8, 4, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -9, -8, 55, -8,
	-10, -8, 56, 65, 13, 65, -1, -16, 16, 67,
	65, 53, -1, 65, -9, -8, 55, 73, 73, -14,
	74, -9, -13, -12, 6, 75, 74, 74, -8, -8,
	48, -15, 51, 68, -16, 65, -10, -16, 55, 8,
	75, 69, 55, -8, -16, -1, -8, -1, 66, 6,
	-8, -8, -1, -10, 66, -7, -16, 8, 75, 69,
	55, -8, 4, 4, 75, 8, -14, 56, -16, 56,
	-16, 55, -9, -9, 75, 75, -8, 68, -8, -8,
	69, -8, 4, -1, 4, -8, 75, -8, 69, 69,
	-8, 4, 66, 65, 66, 65, 67, 66, 29, 66,
	-6, -5, 45, 46, -6, -5, 75, -8, 69, 69,
	65, 75, 75, 8, -16, 69, -16, 66, -8, 8,
	75, 8, 75, 75, 56, 69, 75, 56, 56, 66,
	69, -1, -1, -8, 4, 65, -8, 55, 69, -1,
	65, 65, 75, 69, -12, 66, 75, 75, -8, -8,
	-8, 66, 66, 65, 65, -1, 55, -16, 66, -1,
	-1, 65, 75, 75, 56, 75, -1, -1, 66, -16,
	-1, 66, 66, -1, -8, 66, 66, 30, -1, 66,
	75, 30, 65, 65, -1, -1, 66, 66,
}
var yyDef = [...]int{

	-2, -2, -2, 129, 130, 131, 133, 134, 4, 42,
	-2, 0, 9, 10, 49, 0, 0, 14, 42, 0,
	0, 53, 54, 0, 0, 0, 0, 0, 62, 63,
	64, 65, 0, 129, 129, 0, 0, 0, 0, 0,
	0, 0, 0, 2, -2, 132, 0, 43, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 97, 98, 0, 0, 0, 0, 49, 0, 0,
	49, 11, 50, 12, 0, 0, -2, 0, 0, -2,
	-2, 0, -2, 0, 49, 0, 55, 56, 57, -2,
	0, -2, 0, 42, 0, 49, 39, 0, 53, 0,
	0, 0, 35, 128, 129, 0, 5, 49, 129, 7,
	0, 67, 77, 78, 79, 80, 81, 82, 83, 84,
	-2, -2, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 99, 100, 101, 102, 0, 0, 0, 127,
	8, -2, 129, -2, 0, -2, 0, -2, 0, 0,
	-2, 49, 0, 28, 0, 0, 0, 0, 0, 0,
	42, 129, 129, 40, 0, 76, 49, 49, 0, 0,
	0, 0, 0, 0, 0, -2, 6, 0, 0, 0,
	108, 112, 0, 0, 0, 0, 0, 0, 15, 62,
	0, 0, 0, 45, 0, 0, 0, 0, 104, 111,
	0, 0, 59, 61, 0, 0, 0, 129, 0, 129,
	0, 0, 0, 0, 119, 120, 0, 129, 0, 0,
	36, -2, -2, 0, 44, 66, 107, 0, 117, 118,
	51, -2, 13, -2, 26, -2, 0, 18, 0, 23,
	31, 32, 0, 0, 29, 30, 103, 0, 114, 115,
	-2, 0, 0, 0, 0, 72, 0, 74, 38, 0,
	-2, 0, -2, 121, 0, 37, 123, 0, 0, 27,
	116, 0, 0, 0, 0, -2, 0, 129, 113, 0,
	-2, -2, 0, 73, 41, 75, -2, -2, 0, 0,
	0, 25, 16, -2, -2, 0, 129, -2, 68, 0,
	0, -2, 122, 124, 0, 126, 0, 0, 22, -2,
	34, 69, 70, 0, 0, 17, 21, 0, 33, 71,
	125, 0, -2, -2, 0, 0, 20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	77, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 70, 3, 3, 3, 63, 72, 3,
	74, 75, 61, 59, 56, 60, 73, 62, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 55, 67,
	58, 53, 57, 54, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 68, 3, 69, 71, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 65, 76, 66,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 64,
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
		//line parser.go.y:66
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:70
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:75
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:82
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:89
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
		//line parser.go.y:100
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:105
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:109
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:113
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:118
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:123
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:128
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:133
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:138
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:148
		{
			yyVAL.stmt = &ast.ForStmt{Vars: yyDollar[2].expr_idents, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:163
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:188
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:196
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:201
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
		//line parser.go.y:210
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:216
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:220
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:228
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:232
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
		//line parser.go.y:243
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:249
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:254
		{
			yyVAL.array_count = ast.ArrayCount{Count: 0}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:258
		{
			yyVAL.array_count = ast.ArrayCount{Count: 1}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:262
		{
			yyVAL.array_count.Count = yyVAL.array_count.Count + 1
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:268
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:273
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:277
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:281
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.expr_idents = []string{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:290
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:294
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:299
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:305
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:313
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:318
		{
			yyVAL.exprs = nil
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:322
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:326
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:330
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:336
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:341
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:346
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:351
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:356
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:361
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:366
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:376
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:381
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:386
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:396
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:401
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:406
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 68:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:416
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:421
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:426
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:431
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 73:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:436
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:441
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
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:450
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
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:459
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:464
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:469
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:474
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:479
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:484
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:489
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:494
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:499
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:504
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:509
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:514
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:519
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:524
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:529
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:534
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:539
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:544
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:549
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:554
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:559
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:564
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:569
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:574
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:579
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:584
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:589
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:594
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:599
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:604
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:609
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:614
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:619
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:624
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:629
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:634
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:639
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:644
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:649
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:654
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:659
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:664
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:669
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:674
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:679
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:684
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:689
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:694
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:699
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:704
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:709
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].expr, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:714
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:719
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:730
		{
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:733
		{
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:738
		{
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:741
		{
		}
	}
	goto yystack /* stack new state and value */
}
