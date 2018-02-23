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
const UNARY = 57393

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

//line parser.go.y:737

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 127,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	54, 48,
	-2, 1,
	-1, 10,
	54, 49,
	-2, 24,
	-1, 43,
	54, 48,
	-2, 128,
	-1, 85,
	64, 3,
	-2, 127,
	-1, 88,
	54, 49,
	-2, 45,
	-1, 89,
	16, 42,
	54, 42,
	-2, 52,
	-1, 91,
	64, 3,
	-2, 127,
	-1, 98,
	1, 57,
	8, 57,
	45, 57,
	46, 57,
	51, 57,
	53, 57,
	54, 57,
	63, 57,
	64, 57,
	65, 57,
	67, 57,
	73, 57,
	75, 57,
	-2, 52,
	-1, 100,
	1, 59,
	8, 59,
	45, 59,
	46, 59,
	51, 59,
	53, 59,
	54, 59,
	63, 59,
	64, 59,
	65, 59,
	67, 59,
	73, 59,
	75, 59,
	-2, 52,
	-1, 128,
	17, 0,
	18, 0,
	-2, 84,
	-1, 129,
	17, 0,
	18, 0,
	-2, 85,
	-1, 149,
	54, 49,
	-2, 45,
	-1, 151,
	64, 3,
	-2, 127,
	-1, 153,
	64, 3,
	-2, 127,
	-1, 155,
	64, 1,
	-2, 38,
	-1, 158,
	64, 3,
	-2, 127,
	-1, 183,
	64, 3,
	-2, 127,
	-1, 229,
	54, 50,
	-2, 46,
	-1, 230,
	1, 47,
	45, 47,
	46, 47,
	51, 47,
	54, 51,
	64, 47,
	65, 47,
	75, 47,
	-2, 52,
	-1, 239,
	1, 51,
	8, 51,
	45, 51,
	46, 51,
	54, 51,
	64, 51,
	65, 51,
	67, 51,
	73, 51,
	75, 51,
	-2, 52,
	-1, 241,
	64, 3,
	-2, 127,
	-1, 243,
	64, 3,
	-2, 127,
	-1, 258,
	64, 3,
	-2, 127,
	-1, 268,
	1, 105,
	8, 105,
	45, 105,
	46, 105,
	51, 105,
	53, 105,
	54, 105,
	63, 105,
	64, 105,
	65, 105,
	67, 105,
	73, 105,
	75, 105,
	-2, 103,
	-1, 270,
	1, 109,
	8, 109,
	45, 109,
	46, 109,
	51, 109,
	53, 109,
	54, 109,
	63, 109,
	64, 109,
	65, 109,
	67, 109,
	73, 109,
	75, 109,
	-2, 107,
	-1, 282,
	64, 3,
	-2, 127,
	-1, 287,
	64, 3,
	-2, 127,
	-1, 288,
	64, 3,
	-2, 127,
	-1, 293,
	1, 104,
	8, 104,
	45, 104,
	46, 104,
	51, 104,
	53, 104,
	54, 104,
	63, 104,
	64, 104,
	65, 104,
	67, 104,
	73, 104,
	75, 104,
	-2, 102,
	-1, 294,
	1, 108,
	8, 108,
	45, 108,
	46, 108,
	51, 108,
	53, 108,
	54, 108,
	63, 108,
	64, 108,
	65, 108,
	67, 108,
	73, 108,
	75, 108,
	-2, 106,
	-1, 299,
	64, 3,
	-2, 127,
	-1, 300,
	64, 3,
	-2, 127,
	-1, 303,
	45, 3,
	46, 3,
	64, 3,
	-2, 127,
	-1, 307,
	64, 3,
	-2, 127,
	-1, 314,
	45, 3,
	46, 3,
	64, 3,
	-2, 127,
	-1, 327,
	64, 3,
	-2, 127,
	-1, 328,
	64, 3,
	-2, 127,
}

const yyPrivate = 57344

const yyLast = 2850

var yyAct = [...]int{

	81, 172, 248, 10, 6, 249, 273, 6, 40, 228,
	11, 218, 1, 294, 7, 116, 82, 7, 293, 88,
	45, 92, 6, 80, 95, 96, 97, 99, 101, 90,
	169, 6, 7, 289, 260, 269, 106, 108, 267, 86,
	111, 7, 113, 205, 10, 259, 254, 187, 117, 118,
	234, 120, 121, 122, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 135, 136, 137, 138,
	139, 216, 103, 140, 141, 142, 143, 116, 145, 147,
	149, 150, 6, 110, 150, 144, 109, 2, 94, 150,
	148, 42, 7, 150, 93, 163, 212, 173, 154, 94,
	270, 263, 162, 268, 160, 175, 250, 251, 206, 157,
	177, 179, 188, 170, 332, 331, 149, 324, 321, 320,
	317, 104, 105, 167, 316, 247, 184, 313, 304, 298,
	297, 276, 152, 281, 265, 245, 242, 240, 202, 196,
	102, 328, 327, 307, 300, 288, 287, 191, 258, 151,
	91, 156, 150, 194, 112, 292, 10, 198, 199, 115,
	149, 8, 116, 284, 193, 209, 195, 220, 159, 79,
	201, 200, 5, 155, 326, 250, 251, 44, 322, 224,
	246, 227, 153, 229, 221, 222, 84, 233, 261, 116,
	214, 235, 282, 238, 213, 173, 231, 232, 215, 211,
	182, 210, 168, 119, 185, 114, 83, 252, 46, 255,
	253, 4, 180, 171, 87, 43, 44, 203, 17, 3,
	0, 266, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 192, 0,
	0, 0, 0, 0, 0, 280, 0, 0, 0, 204,
	0, 283, 0, 0, 278, 0, 279, 0, 217, 219,
	0, 0, 0, 238, 0, 0, 291, 0, 0, 104,
	0, 286, 0, 295, 0, 0, 296, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	65, 66, 67, 68, 69, 301, 0, 0, 0, 55,
	305, 306, 0, 0, 262, 0, 264, 0, 78, 0,
	0, 319, 311, 312, 104, 0, 315, 52, 53, 54,
	318, 0, 0, 0, 77, 0, 0, 323, 74, 49,
	76, 0, 72, 0, 0, 0, 0, 0, 0, 0,
	329, 330, 58, 59, 61, 63, 73, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 303, 0, 0, 78, 0, 48, 0, 310,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	314, 77, 0, 0, 0, 74, 49, 76, 309, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 48, 0, 275, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 274, 72, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 48, 0, 272, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 0, 77, 0, 0,
	0, 74, 49, 76, 271, 72, 58, 59, 61, 63,
	73, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 78,
	0, 48, 208, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 77, 207, 0, 0, 74,
	49, 76, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 65,
	66, 67, 68, 69, 0, 0, 70, 71, 55, 56,
	57, 0, 0, 0, 0, 0, 0, 78, 0, 48,
	190, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 77, 189, 0, 0, 74, 49, 76,
	0, 72, 58, 59, 61, 63, 73, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 77, 0, 0, 0, 74, 49, 76, 325, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 48, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 308, 72, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 48, 302, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 0, 77, 0, 0,
	0, 74, 49, 76, 0, 72, 58, 59, 61, 63,
	73, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 78,
	0, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 299, 0, 0, 77, 0, 0, 0, 74,
	49, 76, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 65,
	66, 67, 68, 69, 0, 0, 70, 71, 55, 56,
	57, 0, 0, 0, 0, 0, 0, 78, 0, 48,
	0, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 77, 285, 0, 0, 74, 49, 76,
	0, 72, 58, 59, 61, 63, 73, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 77, 277, 0, 0, 74, 49, 76, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 48, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 77,
	257, 0, 0, 74, 49, 76, 0, 72, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 48, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 244, 77, 0, 0,
	0, 74, 49, 76, 0, 72, 58, 59, 61, 63,
	73, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 78,
	0, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 243, 0, 0, 77, 0, 0, 0, 74,
	49, 76, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 65,
	66, 67, 68, 69, 0, 0, 70, 71, 55, 56,
	57, 0, 0, 0, 0, 0, 0, 78, 0, 48,
	0, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	241, 0, 0, 77, 0, 0, 0, 74, 49, 76,
	0, 72, 58, 59, 61, 63, 73, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 77, 237, 0, 0, 74, 49, 76, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 48, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 225, 72, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 48, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 0, 77, 0, 0,
	0, 74, 49, 76, 223, 72, 58, 59, 61, 63,
	73, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 78,
	0, 48, 186, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 77, 0, 0, 0, 74,
	49, 76, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 65,
	66, 67, 68, 69, 0, 0, 70, 71, 55, 56,
	57, 0, 0, 0, 0, 0, 0, 78, 0, 48,
	0, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	183, 0, 0, 77, 0, 0, 0, 74, 49, 76,
	0, 72, 58, 59, 61, 63, 73, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 77, 0, 0, 0, 74, 49, 76, 174, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 48, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 161, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 0, 72, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 48, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 158, 0, 0, 77, 0, 0,
	0, 74, 49, 76, 0, 72, 21, 22, 28, 0,
	0, 32, 14, 9, 15, 41, 0, 18, 0, 0,
	0, 0, 0, 0, 0, 37, 29, 30, 31, 16,
	19, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	13, 0, 0, 0, 0, 0, 20, 0, 0, 36,
	0, 38, 39, 0, 0, 0, 0, 0, 0, 0,
	23, 27, 0, 0, 0, 34, 0, 6, 33, 0,
	24, 25, 26, 0, 35, 0, 0, 7, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 47, 48, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 0, 77, 0, 0,
	0, 74, 49, 76, 0, 72, 58, 59, 61, 63,
	73, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 78,
	0, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 77, 0, 0, 0, 74,
	49, 76, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 65,
	66, 67, 68, 69, 0, 0, 70, 71, 55, 56,
	57, 0, 0, 0, 0, 0, 0, 78, 0, 48,
	0, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 77, 0, 0, 0, 74, 49, 176,
	0, 72, 58, 59, 61, 63, 73, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 77, 0, 0, 0, 74, 166, 76, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 48, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 58, 59, 61, 63, 77,
	75, 0, 0, 74, 165, 76, 0, 72, 0, 64,
	65, 66, 67, 68, 69, 0, 0, 70, 71, 55,
	56, 57, 0, 0, 0, 0, 0, 0, 78, 0,
	0, 0, 0, 60, 62, 50, 51, 52, 53, 54,
	58, 59, 61, 63, 77, 0, 0, 0, 74, 49,
	76, 0, 72, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 0, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 0, 72, 21, 22,
	197, 0, 0, 32, 14, 9, 15, 41, 0, 18,
	0, 0, 0, 0, 0, 0, 0, 37, 29, 30,
	31, 16, 19, 0, 0, 0, 0, 0, 0, 0,
	0, 12, 13, 0, 0, 0, 0, 0, 20, 0,
	0, 36, 0, 38, 39, 0, 0, 0, 0, 0,
	0, 0, 23, 27, 0, 0, 0, 34, 0, 0,
	33, 0, 24, 25, 26, 0, 35, 21, 22, 28,
	0, 0, 32, 14, 9, 15, 41, 0, 18, 0,
	0, 0, 0, 0, 0, 0, 37, 29, 30, 31,
	16, 19, 0, 0, 0, 0, 0, 0, 0, 0,
	12, 13, 0, 0, 0, 0, 0, 20, 0, 0,
	36, 0, 38, 39, 0, 0, 0, 0, 0, 0,
	0, 23, 27, 0, 61, 63, 34, 0, 0, 33,
	0, 24, 25, 26, 0, 35, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 0, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 77, 0, 0, 0, 74, 49, 76, 0, 72,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 0, 239, 22, 28, 0, 0, 32, 0, 78,
	0, 0, 0, 0, 0, 0, 50, 51, 52, 53,
	54, 37, 29, 30, 31, 77, 0, 0, 0, 74,
	49, 76, 0, 72, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 36, 0, 38, 39, 0,
	0, 0, 0, 21, 22, 28, 23, 27, 32, 0,
	0, 34, 0, 0, 33, 290, 24, 25, 26, 0,
	35, 0, 37, 29, 30, 31, 0, 0, 0, 0,
	0, 21, 22, 28, 0, 0, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 36, 0, 38, 39,
	37, 29, 30, 31, 0, 0, 0, 23, 27, 21,
	22, 28, 34, 0, 32, 33, 256, 24, 25, 26,
	0, 35, 0, 0, 36, 0, 38, 39, 37, 29,
	30, 31, 0, 0, 0, 23, 27, 21, 22, 28,
	34, 0, 32, 33, 236, 24, 25, 26, 0, 35,
	0, 0, 36, 178, 38, 39, 37, 29, 30, 31,
	0, 0, 0, 23, 27, 0, 0, 0, 34, 0,
	0, 181, 0, 24, 25, 26, 0, 35, 0, 0,
	36, 0, 38, 39, 0, 0, 164, 0, 21, 22,
	28, 23, 27, 32, 0, 0, 34, 0, 0, 33,
	0, 24, 25, 26, 0, 35, 0, 37, 29, 30,
	31, 0, 0, 0, 0, 0, 0, 0, 0, 21,
	22, 28, 0, 0, 32, 0, 0, 0, 0, 0,
	0, 36, 0, 38, 39, 0, 0, 146, 37, 29,
	30, 31, 23, 27, 0, 0, 0, 34, 0, 0,
	33, 0, 24, 25, 26, 0, 35, 0, 0, 0,
	0, 0, 36, 0, 38, 39, 0, 0, 0, 0,
	239, 22, 28, 23, 27, 32, 0, 0, 34, 0,
	0, 33, 0, 24, 25, 26, 0, 35, 0, 37,
	29, 30, 31, 0, 0, 0, 0, 0, 230, 22,
	28, 0, 0, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 36, 0, 38, 39, 37, 29, 30,
	31, 0, 0, 0, 23, 27, 21, 22, 28, 34,
	0, 32, 33, 0, 24, 25, 26, 0, 35, 0,
	0, 36, 0, 38, 39, 37, 29, 30, 31, 0,
	0, 0, 23, 27, 107, 22, 28, 34, 0, 32,
	33, 0, 24, 25, 26, 0, 35, 0, 0, 36,
	0, 38, 39, 37, 29, 30, 31, 0, 0, 0,
	23, 27, 100, 22, 28, 34, 0, 32, 226, 0,
	24, 25, 26, 0, 35, 0, 0, 36, 0, 38,
	39, 37, 29, 30, 31, 0, 0, 0, 23, 27,
	98, 22, 28, 34, 0, 32, 33, 0, 24, 25,
	26, 0, 35, 0, 0, 36, 0, 38, 39, 37,
	29, 30, 31, 0, 0, 0, 23, 27, 89, 22,
	28, 34, 0, 32, 33, 0, 24, 25, 26, 0,
	35, 0, 0, 36, 0, 38, 39, 37, 29, 30,
	31, 0, 0, 0, 23, 27, 0, 0, 0, 34,
	0, 0, 33, 0, 24, 25, 26, 0, 35, 0,
	0, 36, 0, 38, 39, 0, 64, 65, 66, 67,
	68, 69, 23, 27, 0, 0, 55, 85, 0, 0,
	33, 0, 24, 25, 26, 78, 35, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 77, 0, 0, 0, 74, 49, 76, 0, 72,
}
var yyPact = [...]int{

	-34, -1000, 2193, -34, -34, -1000, -1000, -1000, -1000, 204,
	1731, 118, -1000, -1000, 2535, 2535, 202, 172, 2754, 87,
	2535, 22, -1000, 2535, 2535, 2535, 2726, 2698, -1000, -1000,
	-1000, -1000, 68, -34, -34, 2535, 2670, 14, 11, 2535,
	100, 2535, -1000, 1672, -1000, 108, -1000, 2535, 2535, 199,
	2535, 2535, 2535, 2535, 2535, 2535, 2535, 2535, 2535, 2535,
	2535, 2535, 2535, 2535, 2535, 2535, 2535, 2535, 2535, 2535,
	-1000, -1000, 2535, 2535, 2535, 2535, 2535, 2504, 2535, 2535,
	98, 1789, 1789, 86, 119, -34, 135, 44, 1601, 22,
	117, -34, 1543, 2535, 2453, 2775, 2775, 2775, 22, 1963,
	22, 1905, 198, -42, 2535, 189, 1485, 33, 1847, 2535,
	2425, 1789, -34, 1427, -1000, 2535, -34, 1789, 1369, -1000,
	258, 258, 2775, 2775, 2775, 1789, 2279, 2279, 2235, 2235,
	2279, 2279, 2279, 2279, 1789, 1789, 1789, 1789, 1789, 1789,
	1789, 2008, 1789, 2053, 39, 557, 2535, 1789, -1000, 1789,
	-34, -34, 2535, -34, 75, 2124, 2535, 2535, -34, 2535,
	74, -34, 35, 499, 2535, 197, 195, 23, 186, 194,
	17, -43, -1000, 114, -1000, 2535, 2535, 1311, 2535, 1253,
	2642, -58, 2614, -34, -1000, 193, 2535, -23, -1000, -1000,
	2397, 1195, 2586, 73, 1137, 72, -1000, 114, 1079, 1021,
	71, -1000, 151, 61, 130, -27, -1000, -1000, 2369, 963,
	-1000, -1000, 85, -28, -39, 180, -34, 34, -34, 70,
	2535, 30, 27, -1000, 441, -1000, -61, 383, -1000, 1789,
	22, 67, -1000, 1789, -1000, 905, -1000, -1000, 1789, 22,
	-1000, -34, -1000, -34, 2535, -1000, 129, -1000, -1000, -1000,
	2535, 110, -1000, -1000, -1000, 847, -1000, -1000, -34, 83,
	82, -40, 2318, -1000, 91, -1000, 1789, -55, -1000, -60,
	-1000, -1000, 2535, -1000, -1000, 2535, -1000, -1000, 66, 65,
	789, 81, -34, 731, -34, -1000, 64, -34, -34, 80,
	-1000, -1000, -1000, -1000, -1000, 673, 325, -1000, -1000, -34,
	-34, 63, -34, -34, -1000, 60, 56, -34, -1000, -1000,
	2535, 55, 54, 148, -34, -1000, -1000, -1000, 53, 615,
	-1000, 144, 79, -1000, -1000, -1000, 78, -34, -34, 51,
	50, -1000, -1000,
}
var yyPgo = [...]int{

	0, 12, 219, 161, 218, 5, 2, 217, 0, 8,
	10, 214, 1, 213, 20, 212, 87, 211, 172,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 7, 7,
	7, 7, 7, 6, 5, 15, 15, 12, 13, 13,
	13, 14, 14, 14, 11, 10, 10, 10, 9, 9,
	9, 9, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 16, 16, 17,
	17, 18, 18,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 3, 3, 1,
	1, 2, 2, 5, 1, 4, 7, 9, 5, 13,
	12, 9, 8, 5, 1, 7, 5, 5, 0, 2,
	2, 2, 2, 5, 4, 3, 2, 3, 0, 1,
	4, 0, 1, 4, 3, 1, 4, 4, 0, 1,
	4, 4, 1, 1, 2, 2, 2, 2, 4, 2,
	4, 1, 1, 1, 1, 5, 3, 7, 8, 8,
	9, 5, 6, 5, 6, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 3, 3,
	3, 3, 5, 4, 6, 5, 5, 4, 6, 5,
	4, 4, 6, 5, 5, 6, 5, 5, 4, 5,
	7, 4, 5, 7, 9, 3, 2, 0, 1, 1,
	2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -16, -2, -17, -18, 65, 75, -3, 11,
	-8, -10, 37, 38, 10, 12, 27, -4, 15, 28,
	44, 4, 5, 58, 68, 69, 70, 59, 6, 24,
	25, 26, 9, 66, 63, 72, 47, 23, 49, 50,
	-9, 13, -16, -17, -18, -14, 4, 51, 52, 71,
	57, 58, 59, 60, 61, 41, 42, 43, 17, 18,
	55, 19, 56, 20, 31, 32, 33, 34, 35, 36,
	39, 40, 74, 21, 70, 22, 72, 66, 50, 51,
	-9, -8, -8, 4, 14, 63, -14, -11, -8, 4,
	-10, 63, -8, 72, 66, -8, -8, -8, 4, -8,
	4, -8, 72, 4, -16, -16, -8, 4, -8, 72,
	72, -8, 54, -8, -3, 51, 54, -8, -8, 4,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -9, -8, 53, -8, -10, -8,
	54, 63, 13, 63, -1, -16, 16, 65, 63, 51,
	-1, 63, -9, -8, 53, 71, 71, -14, 4, 72,
	-9, -13, -12, 6, 73, 72, 72, -8, 48, -8,
	-15, 66, -16, 63, -10, -16, 53, 8, 73, 67,
	53, -8, -16, -1, -8, -1, 64, 6, -8, -8,
	-1, -10, 64, -7, -16, 8, 73, 67, 53, -8,
	4, 4, 73, 8, -14, 4, 54, -16, 54, -16,
	53, -9, -9, 73, -8, 73, 66, -8, 67, -8,
	4, -1, 4, -8, 73, -8, 67, 67, -8, 4,
	64, 63, 64, 63, 65, 64, 29, 64, -6, -5,
	45, 46, -6, -5, 73, -8, 67, 67, 63, 73,
	73, 8, -16, 67, -16, 64, -8, 8, 73, 8,
	73, 73, 54, 67, 73, 54, 64, 67, -1, -1,
	-8, 4, 63, -8, 53, 67, -1, 63, 63, 73,
	67, -12, 64, 73, 73, -8, -8, 64, 64, 63,
	63, -1, 53, -16, 64, -1, -1, 63, 73, 73,
	54, -1, -1, 64, -16, -1, 64, 64, -1, -8,
	64, 64, 30, -1, 64, 73, 30, 63, 63, -1,
	-1, 64, 64,
}
var yyDef = [...]int{

	-2, -2, -2, 127, 128, 129, 131, 132, 4, 41,
	-2, 0, 9, 10, 48, 0, 0, 14, 41, 0,
	0, 52, 53, 0, 0, 0, 0, 0, 61, 62,
	63, 64, 0, 127, 127, 0, 0, 0, 0, 0,
	0, 0, 2, -2, 130, 0, 42, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	96, 97, 0, 0, 0, 0, 48, 0, 0, 48,
	11, 49, 12, 0, 0, -2, 0, 0, -2, -2,
	0, -2, 0, 48, 0, 54, 55, 56, -2, 0,
	-2, 0, 41, 0, 48, 38, 0, 52, 0, 0,
	0, 126, 127, 0, 5, 48, 127, 7, 0, 66,
	76, 77, 78, 79, 80, 81, 82, 83, -2, -2,
	86, 87, 88, 89, 90, 91, 92, 93, 94, 95,
	98, 99, 100, 101, 0, 0, 0, 125, 8, -2,
	127, -2, 0, -2, 0, -2, 0, 0, -2, 48,
	0, 28, 0, 0, 0, 0, 0, 0, 42, 41,
	127, 127, 39, 0, 75, 48, 48, 0, 0, 0,
	0, 127, 0, -2, 6, 0, 0, 0, 107, 111,
	0, 0, 0, 0, 0, 0, 15, 61, 0, 0,
	0, 44, 0, 0, 0, 0, 103, 110, 0, 0,
	58, 60, 0, 0, 0, 42, 127, 0, 127, 0,
	0, 0, 0, 118, 0, 121, 127, 0, 36, -2,
	-2, 0, 43, 65, 106, 0, 116, 117, 50, -2,
	13, -2, 26, -2, 0, 18, 0, 23, 31, 32,
	0, 0, 29, 30, 102, 0, 113, 114, -2, 0,
	0, 0, 0, 71, 0, 73, 37, 0, -2, 0,
	-2, 119, 0, 35, 122, 0, 27, 115, 0, 0,
	0, 0, -2, 0, 127, 112, 0, -2, -2, 0,
	72, 40, 74, -2, -2, 0, 0, 25, 16, -2,
	-2, 0, 127, -2, 67, 0, 0, -2, 120, 123,
	0, 0, 0, 22, -2, 34, 68, 69, 0, 0,
	17, 21, 0, 33, 70, 124, 0, -2, -2, 0,
	0, 20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	75, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 68, 3, 3, 3, 61, 70, 3,
	72, 73, 59, 57, 54, 58, 71, 60, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 53, 65,
	56, 51, 55, 52, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 66, 3, 67, 69, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 63, 74, 64,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 62,
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:255
		{
			yyVAL.array_count = ast.ArrayCount{Count: yyDollar[1].array_count.Count + 1}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:259
		{
			yyVAL.array_count = ast.ArrayCount{Count: 1}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:265
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:270
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:278
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 41:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:283
		{
			yyVAL.expr_idents = []string{}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:291
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:296
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:310
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:315
		{
			yyVAL.exprs = nil
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:319
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:323
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:333
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:338
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:348
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:353
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:358
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:368
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:373
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:378
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:388
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:393
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:398
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:403
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:408
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:413
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:418
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:423
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:428
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:433
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:438
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
	case 74:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:447
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:456
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:461
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:466
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:471
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:476
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:481
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:486
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:491
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:496
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:501
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:506
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:511
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:516
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:521
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:526
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:531
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:536
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:541
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:546
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:551
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:556
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:561
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:566
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:571
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:576
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:581
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:586
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:591
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:596
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 104:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:601
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:606
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:611
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:616
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:621
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:626
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:631
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:636
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:641
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:646
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:651
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:656
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:661
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:666
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:671
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:676
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:681
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:686
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:691
		{
			yyVAL.expr = &ast.MakeArrayExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:696
		{
			yyVAL.expr = &ast.MakeArrayExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:701
		{
			yyVAL.expr = &ast.MakeArrayExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:706
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:711
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:722
		{
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:725
		{
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:730
		{
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:733
		{
		}
	}
	goto yystack /* stack new state and value */
}
