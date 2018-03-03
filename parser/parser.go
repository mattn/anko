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
	"TYPE",
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

//line parser.go.y:740

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 128,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	55, 49,
	-2, 1,
	-1, 10,
	55, 50,
	-2, 24,
	-1, 43,
	55, 49,
	-2, 129,
	-1, 85,
	65, 3,
	-2, 128,
	-1, 88,
	55, 50,
	-2, 46,
	-1, 89,
	16, 43,
	55, 43,
	-2, 53,
	-1, 91,
	65, 3,
	-2, 128,
	-1, 98,
	1, 58,
	8, 58,
	45, 58,
	46, 58,
	52, 58,
	54, 58,
	55, 58,
	64, 58,
	65, 58,
	66, 58,
	68, 58,
	74, 58,
	76, 58,
	-2, 53,
	-1, 100,
	1, 60,
	8, 60,
	45, 60,
	46, 60,
	52, 60,
	54, 60,
	55, 60,
	64, 60,
	65, 60,
	66, 60,
	68, 60,
	74, 60,
	76, 60,
	-2, 53,
	-1, 128,
	17, 0,
	18, 0,
	-2, 85,
	-1, 129,
	17, 0,
	18, 0,
	-2, 86,
	-1, 149,
	55, 50,
	-2, 46,
	-1, 151,
	65, 3,
	-2, 128,
	-1, 153,
	65, 3,
	-2, 128,
	-1, 155,
	65, 1,
	-2, 39,
	-1, 158,
	65, 3,
	-2, 128,
	-1, 182,
	65, 3,
	-2, 128,
	-1, 227,
	55, 51,
	-2, 47,
	-1, 228,
	1, 48,
	45, 48,
	46, 48,
	52, 48,
	55, 52,
	65, 48,
	66, 48,
	76, 48,
	-2, 53,
	-1, 237,
	1, 52,
	8, 52,
	45, 52,
	46, 52,
	55, 52,
	65, 52,
	66, 52,
	68, 52,
	74, 52,
	76, 52,
	-2, 53,
	-1, 239,
	65, 3,
	-2, 128,
	-1, 241,
	65, 3,
	-2, 128,
	-1, 256,
	65, 3,
	-2, 128,
	-1, 266,
	1, 106,
	8, 106,
	45, 106,
	46, 106,
	52, 106,
	54, 106,
	55, 106,
	64, 106,
	65, 106,
	66, 106,
	68, 106,
	74, 106,
	76, 106,
	-2, 104,
	-1, 268,
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
	68, 110,
	74, 110,
	76, 110,
	-2, 108,
	-1, 281,
	65, 3,
	-2, 128,
	-1, 286,
	65, 3,
	-2, 128,
	-1, 287,
	65, 3,
	-2, 128,
	-1, 292,
	1, 105,
	8, 105,
	45, 105,
	46, 105,
	52, 105,
	54, 105,
	55, 105,
	64, 105,
	65, 105,
	66, 105,
	68, 105,
	74, 105,
	76, 105,
	-2, 103,
	-1, 293,
	1, 109,
	8, 109,
	45, 109,
	46, 109,
	52, 109,
	54, 109,
	55, 109,
	64, 109,
	65, 109,
	66, 109,
	68, 109,
	74, 109,
	76, 109,
	-2, 107,
	-1, 299,
	65, 3,
	-2, 128,
	-1, 300,
	65, 3,
	-2, 128,
	-1, 303,
	45, 3,
	46, 3,
	65, 3,
	-2, 128,
	-1, 307,
	65, 3,
	-2, 128,
	-1, 315,
	45, 3,
	46, 3,
	65, 3,
	-2, 128,
	-1, 328,
	65, 3,
	-2, 128,
	-1, 329,
	65, 3,
	-2, 128,
}

const yyPrivate = 57344

const yyLast = 2908

var yyAct = [...]int{

	81, 171, 246, 10, 6, 247, 271, 6, 40, 293,
	292, 288, 1, 11, 7, 168, 82, 7, 257, 88,
	45, 92, 216, 80, 95, 96, 97, 99, 101, 94,
	267, 252, 90, 6, 265, 93, 106, 108, 232, 86,
	111, 259, 113, 7, 10, 103, 110, 109, 117, 118,
	261, 120, 121, 122, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 135, 136, 137, 138,
	139, 214, 212, 140, 141, 142, 143, 150, 145, 147,
	149, 150, 6, 204, 186, 144, 2, 226, 116, 172,
	42, 94, 7, 148, 157, 163, 268, 174, 154, 333,
	266, 332, 162, 177, 160, 325, 179, 258, 248, 249,
	176, 322, 321, 169, 102, 318, 149, 317, 314, 116,
	104, 105, 180, 167, 304, 298, 297, 280, 245, 183,
	150, 150, 275, 263, 243, 240, 238, 201, 211, 195,
	152, 329, 328, 307, 300, 287, 286, 190, 291, 205,
	187, 256, 151, 193, 91, 156, 10, 197, 198, 115,
	149, 150, 116, 112, 192, 208, 194, 283, 218, 159,
	79, 199, 155, 200, 248, 249, 5, 8, 222, 224,
	225, 44, 227, 219, 220, 327, 231, 281, 323, 213,
	233, 153, 236, 244, 116, 229, 84, 172, 230, 181,
	46, 210, 209, 184, 119, 83, 250, 4, 253, 251,
	178, 43, 170, 87, 202, 17, 3, 0, 0, 264,
	44, 114, 0, 0, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 0, 0, 55, 191, 0, 0,
	0, 0, 0, 279, 0, 78, 0, 0, 203, 282,
	0, 0, 277, 0, 278, 0, 215, 217, 0, 0,
	0, 236, 77, 0, 290, 0, 74, 49, 76, 285,
	72, 294, 0, 0, 295, 296, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 301, 0, 0, 0, 0, 305,
	306, 260, 0, 262, 0, 0, 0, 0, 0, 0,
	104, 320, 312, 313, 0, 0, 316, 0, 0, 0,
	319, 58, 59, 61, 63, 73, 75, 0, 324, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 330, 331, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 310,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	303, 77, 0, 0, 0, 74, 49, 76, 309, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 315,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 0, 48, 0, 273, 60,
	62, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	77, 0, 0, 0, 74, 49, 76, 272, 72, 58,
	59, 61, 63, 73, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 65, 66, 67, 68, 69, 0,
	0, 70, 71, 55, 56, 57, 0, 0, 0, 0,
	0, 0, 78, 0, 0, 48, 0, 270, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 269, 72, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 0, 48, 207, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 0, 0, 0, 77, 206,
	0, 0, 74, 49, 76, 0, 72, 58, 59, 61,
	63, 73, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 65, 66, 67, 68, 69, 0, 0, 70,
	71, 55, 56, 57, 0, 0, 0, 0, 0, 0,
	78, 0, 0, 48, 189, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 0, 77, 188, 0,
	0, 74, 49, 76, 0, 72, 58, 59, 61, 63,
	73, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 78,
	0, 0, 48, 0, 0, 60, 62, 50, 51, 52,
	53, 54, 0, 0, 0, 0, 77, 0, 0, 0,
	74, 49, 76, 326, 72, 58, 59, 61, 63, 73,
	75, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	65, 66, 67, 68, 69, 0, 0, 70, 71, 55,
	56, 57, 0, 0, 0, 0, 0, 0, 78, 0,
	0, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 77, 0, 0, 0, 74,
	49, 76, 311, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 65,
	66, 67, 68, 69, 0, 0, 70, 71, 55, 56,
	57, 0, 0, 0, 0, 0, 0, 78, 0, 0,
	48, 0, 0, 60, 62, 50, 51, 52, 53, 54,
	0, 0, 0, 0, 77, 0, 0, 0, 74, 49,
	76, 308, 72, 58, 59, 61, 63, 73, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 65, 66,
	67, 68, 69, 0, 0, 70, 71, 55, 56, 57,
	0, 0, 0, 0, 0, 0, 78, 0, 0, 48,
	302, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 77, 0, 0, 0, 74, 49, 76,
	0, 72, 58, 59, 61, 63, 73, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 0, 48, 0,
	0, 60, 62, 50, 51, 52, 53, 54, 0, 299,
	0, 0, 77, 0, 0, 0, 74, 49, 76, 0,
	72, 58, 59, 61, 63, 73, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 77, 284, 0, 0, 74, 49, 76, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 0, 48, 0, 0, 60,
	62, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	77, 276, 0, 0, 74, 49, 76, 0, 72, 58,
	59, 61, 63, 73, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 65, 66, 67, 68, 69, 0,
	0, 70, 71, 55, 56, 57, 0, 0, 0, 0,
	0, 0, 78, 0, 0, 48, 0, 274, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 0, 72, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 0, 48, 0, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 0, 0, 0, 77, 255,
	0, 0, 74, 49, 76, 0, 72, 58, 59, 61,
	63, 73, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 65, 66, 67, 68, 69, 0, 0, 70,
	71, 55, 56, 57, 0, 0, 0, 0, 0, 0,
	78, 0, 0, 48, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 242, 77, 0, 0,
	0, 74, 49, 76, 0, 72, 58, 59, 61, 63,
	73, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 78,
	0, 0, 48, 0, 0, 60, 62, 50, 51, 52,
	53, 54, 0, 241, 0, 0, 77, 0, 0, 0,
	74, 49, 76, 0, 72, 58, 59, 61, 63, 73,
	75, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	65, 66, 67, 68, 69, 0, 0, 70, 71, 55,
	56, 57, 0, 0, 0, 0, 0, 0, 78, 0,
	0, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 239, 0, 0, 77, 0, 0, 0, 74,
	49, 76, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 64, 65,
	66, 67, 68, 69, 0, 0, 70, 71, 55, 56,
	57, 0, 0, 0, 0, 0, 0, 78, 0, 0,
	48, 0, 0, 60, 62, 50, 51, 52, 53, 54,
	0, 0, 0, 0, 77, 235, 0, 0, 74, 49,
	76, 0, 72, 58, 59, 61, 63, 73, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 65, 66,
	67, 68, 69, 0, 0, 70, 71, 55, 56, 57,
	0, 0, 0, 0, 0, 0, 78, 0, 0, 48,
	0, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 77, 0, 0, 0, 74, 49, 76,
	221, 72, 58, 59, 61, 63, 73, 75, 0, 0,
	0, 0, 0, 0, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 0, 48, 185,
	0, 60, 62, 50, 51, 52, 53, 54, 0, 0,
	0, 0, 77, 0, 0, 0, 74, 49, 76, 0,
	72, 58, 59, 61, 63, 73, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 182, 0,
	0, 77, 0, 0, 0, 74, 49, 76, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 0, 48, 0, 0, 60,
	62, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	77, 0, 0, 0, 74, 49, 76, 173, 72, 58,
	59, 61, 63, 73, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 65, 66, 67, 68, 69, 0,
	0, 70, 71, 55, 56, 57, 0, 0, 0, 0,
	0, 0, 78, 0, 0, 48, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 161, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 0, 72, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 0, 48, 0, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 158, 0, 0, 77, 0,
	0, 0, 74, 49, 76, 0, 72, 21, 22, 28,
	0, 0, 32, 14, 9, 15, 41, 0, 18, 0,
	0, 0, 0, 0, 0, 0, 37, 29, 30, 31,
	16, 19, 0, 0, 0, 0, 0, 0, 0, 0,
	12, 13, 0, 0, 0, 0, 0, 20, 0, 0,
	36, 0, 38, 39, 0, 0, 0, 0, 0, 0,
	0, 0, 23, 27, 0, 0, 0, 34, 0, 6,
	33, 0, 24, 25, 26, 0, 35, 0, 0, 7,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 70, 71, 55, 56, 57, 0, 0, 0,
	0, 0, 0, 78, 0, 47, 48, 0, 0, 60,
	62, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	77, 0, 0, 0, 74, 49, 76, 0, 72, 58,
	59, 61, 63, 73, 75, 0, 0, 0, 0, 0,
	0, 0, 0, 64, 65, 66, 67, 68, 69, 0,
	0, 70, 71, 55, 56, 57, 0, 0, 0, 0,
	0, 0, 78, 0, 0, 48, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 0, 72, 58, 59,
	61, 63, 73, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 0, 48, 0, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 0, 0, 0, 77, 0,
	0, 0, 74, 49, 175, 0, 72, 58, 59, 61,
	63, 73, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 64, 65, 66, 67, 68, 69, 0, 0, 70,
	71, 55, 56, 57, 0, 0, 0, 0, 0, 0,
	78, 0, 0, 48, 0, 0, 60, 62, 50, 51,
	52, 53, 54, 0, 0, 0, 0, 77, 0, 0,
	0, 74, 166, 76, 0, 72, 58, 59, 61, 63,
	73, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 56, 57, 0, 0, 0, 0, 0, 0, 78,
	0, 0, 48, 0, 0, 60, 62, 50, 51, 52,
	53, 54, 58, 59, 61, 63, 77, 75, 0, 0,
	74, 165, 76, 0, 72, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 0, 0, 0,
	0, 60, 62, 50, 51, 52, 53, 54, 58, 59,
	61, 63, 77, 0, 0, 0, 74, 49, 76, 0,
	72, 0, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 56, 57, 0, 0, 0, 0, 0,
	0, 78, 0, 0, 0, 0, 0, 60, 62, 50,
	51, 52, 53, 54, 0, 0, 0, 0, 77, 0,
	0, 0, 74, 49, 76, 0, 72, 21, 22, 196,
	0, 0, 32, 14, 9, 15, 41, 0, 18, 0,
	0, 0, 0, 0, 0, 0, 37, 29, 30, 31,
	16, 19, 0, 0, 0, 0, 0, 0, 0, 0,
	12, 13, 0, 0, 0, 0, 0, 20, 0, 0,
	36, 0, 38, 39, 0, 0, 0, 0, 0, 0,
	0, 0, 23, 27, 0, 0, 0, 34, 0, 0,
	33, 0, 24, 25, 26, 0, 35, 21, 22, 28,
	0, 0, 32, 14, 9, 15, 41, 0, 18, 0,
	0, 0, 0, 0, 0, 0, 37, 29, 30, 31,
	16, 19, 0, 0, 0, 0, 0, 0, 0, 0,
	12, 13, 0, 0, 0, 0, 0, 20, 0, 0,
	36, 0, 38, 39, 0, 0, 0, 0, 0, 0,
	0, 0, 23, 27, 0, 61, 63, 34, 0, 0,
	33, 0, 24, 25, 26, 0, 35, 64, 65, 66,
	67, 68, 69, 0, 0, 70, 71, 55, 56, 57,
	0, 0, 0, 0, 0, 0, 78, 0, 0, 0,
	0, 0, 60, 62, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 77, 0, 0, 0, 74, 49, 76,
	0, 72, 64, 65, 66, 67, 68, 69, 0, 0,
	70, 71, 55, 0, 0, 237, 22, 28, 0, 0,
	32, 78, 0, 0, 0, 0, 0, 0, 0, 50,
	51, 52, 53, 54, 37, 29, 30, 31, 77, 0,
	0, 0, 74, 49, 76, 0, 72, 21, 22, 28,
	0, 0, 32, 0, 0, 0, 0, 0, 36, 0,
	38, 39, 0, 0, 0, 0, 37, 29, 30, 31,
	23, 27, 0, 0, 0, 34, 0, 0, 33, 289,
	24, 25, 26, 0, 35, 0, 0, 0, 0, 0,
	36, 0, 38, 39, 0, 0, 0, 0, 0, 21,
	22, 28, 23, 27, 32, 0, 0, 34, 0, 0,
	33, 254, 24, 25, 26, 0, 35, 0, 37, 29,
	30, 31, 0, 0, 0, 0, 0, 21, 22, 28,
	0, 0, 32, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 36, 0, 38, 39, 37, 29, 30, 31,
	0, 0, 0, 0, 23, 27, 0, 0, 0, 34,
	0, 0, 33, 234, 24, 25, 26, 0, 35, 0,
	36, 0, 38, 39, 0, 0, 0, 164, 0, 21,
	22, 28, 23, 27, 32, 0, 0, 34, 0, 0,
	33, 0, 24, 25, 26, 0, 35, 0, 37, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 21, 22, 28, 0, 0, 32, 0, 0, 0,
	0, 0, 36, 0, 38, 39, 0, 0, 0, 146,
	37, 29, 30, 31, 23, 27, 0, 0, 0, 34,
	0, 0, 33, 0, 24, 25, 26, 0, 35, 0,
	0, 0, 0, 0, 36, 0, 38, 39, 0, 0,
	0, 0, 0, 237, 22, 28, 23, 27, 32, 0,
	0, 34, 0, 0, 33, 0, 24, 25, 26, 0,
	35, 0, 37, 29, 30, 31, 0, 0, 0, 0,
	0, 228, 22, 28, 0, 0, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 36, 0, 38, 39,
	37, 29, 30, 31, 0, 0, 0, 0, 23, 27,
	21, 22, 28, 34, 0, 32, 33, 0, 24, 25,
	26, 0, 35, 0, 36, 0, 38, 39, 0, 37,
	29, 30, 31, 0, 0, 0, 23, 27, 107, 22,
	28, 34, 0, 32, 33, 0, 24, 25, 26, 0,
	35, 0, 0, 36, 0, 38, 39, 37, 29, 30,
	31, 0, 0, 0, 0, 23, 27, 100, 22, 28,
	34, 0, 32, 223, 0, 24, 25, 26, 0, 35,
	0, 36, 0, 38, 39, 0, 37, 29, 30, 31,
	0, 0, 0, 23, 27, 98, 22, 28, 34, 0,
	32, 33, 0, 24, 25, 26, 0, 35, 0, 0,
	36, 0, 38, 39, 37, 29, 30, 31, 0, 0,
	0, 0, 23, 27, 89, 22, 28, 34, 0, 32,
	33, 0, 24, 25, 26, 0, 35, 0, 36, 0,
	38, 39, 0, 37, 29, 30, 31, 0, 0, 0,
	23, 27, 0, 0, 0, 34, 0, 0, 33, 0,
	24, 25, 26, 0, 35, 0, 0, 36, 0, 38,
	39, 0, 0, 64, 65, 66, 67, 68, 69, 23,
	27, 0, 0, 55, 85, 0, 0, 33, 0, 24,
	25, 26, 78, 35, 0, 0, 0, 0, 0, 0,
	0, 0, 52, 53, 54, 0, 0, 0, 0, 77,
	0, 0, 0, 74, 49, 76, 0, 72,
}
var yyPact = [...]int{

	-59, -1000, 2263, -59, -59, -1000, -1000, -1000, -1000, 196,
	1793, 118, -1000, -1000, 2587, 2587, 201, 182, 2810, 90,
	2587, -38, -1000, 2587, 2587, 2587, 2781, 2753, -1000, -1000,
	-1000, -1000, 41, -59, -59, 2587, 2724, -26, -27, 2587,
	108, 2587, -1000, 1733, -1000, 107, -1000, 2587, 2587, 200,
	2587, 2587, 2587, 2587, 2587, 2587, 2587, 2587, 2587, 2587,
	2587, 2587, 2587, 2587, 2587, 2587, 2587, 2587, 2587, 2587,
	-1000, -1000, 2587, 2587, 2587, 2587, 2587, 2555, 2587, 2587,
	106, 1852, 1852, 88, 127, -59, 139, 28, 1661, -38,
	117, -59, 1602, 2587, 2503, 195, 195, 195, -38, 2029,
	-38, 1970, 196, -58, 2587, 191, 1543, 24, 1911, 2587,
	55, 1852, -59, 1484, -1000, 2587, -59, 1852, 1425, -1000,
	2832, 2832, 195, 195, 195, 1852, 2351, 2351, 2306, 2306,
	2351, 2351, 2351, 2351, 1852, 1852, 1852, 1852, 1852, 1852,
	1852, 2075, 1852, 2121, 76, 540, 2587, 1852, -1000, 1852,
	-59, -59, 2587, -59, 74, 2193, 2587, 2587, -59, 2587,
	72, -59, 75, 481, 2587, 198, 197, 64, 196, 16,
	-33, -1000, 114, -1000, 2587, 2587, 1366, 2587, 2696, 2587,
	19, 2667, -59, -1000, 194, 2587, -36, -1000, -1000, 2475,
	1307, 2639, 71, 1248, 70, -1000, 114, 1189, 1130, 69,
	-1000, 164, 63, 129, -43, -1000, -1000, 2423, 1071, -1000,
	-1000, 87, -56, 33, -59, -18, -59, 68, 2587, 26,
	22, -1000, 422, -62, 363, 1012, -1000, 1852, -38, 67,
	-1000, 1852, -1000, 953, -1000, -1000, 1852, -38, -1000, -59,
	-1000, -59, 2587, -1000, 123, -1000, -1000, -1000, 2587, 113,
	-1000, -1000, -1000, 894, -1000, -1000, -59, 82, 81, -63,
	2391, -1000, 83, -1000, 1852, -64, -1000, -65, -1000, -1000,
	2587, -1000, -1000, 2587, 2587, -1000, -1000, 61, 60, 835,
	80, -59, 776, -59, -1000, 59, -59, -59, 79, -1000,
	-1000, -1000, -1000, -1000, 717, 304, 658, -1000, -1000, -59,
	-59, 53, -59, -59, -1000, 52, 50, -59, -1000, -1000,
	2587, -1000, 47, 46, 158, -59, -1000, -1000, -1000, 40,
	599, -1000, 155, 78, -1000, -1000, -1000, 77, -59, -59,
	36, 34, -1000, -1000,
}
var yyPgo = [...]int{

	0, 12, 216, 177, 215, 5, 2, 214, 0, 8,
	13, 213, 1, 212, 20, 210, 86, 207, 176,
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
	8, 8, 8, 8, 8, 8, 8, 8, 16, 16,
	17, 17, 18, 18,
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
	5, 7, 5, 7, 9, 7, 3, 2, 0, 1,
	1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -16, -2, -17, -18, 66, 76, -3, 11,
	-8, -10, 37, 38, 10, 12, 27, -4, 15, 28,
	44, 4, 5, 59, 69, 70, 71, 60, 6, 24,
	25, 26, 9, 67, 64, 73, 47, 23, 49, 50,
	-9, 13, -16, -17, -18, -14, 4, 52, 53, 72,
	58, 59, 60, 61, 62, 41, 42, 43, 17, 18,
	56, 19, 57, 20, 31, 32, 33, 34, 35, 36,
	39, 40, 75, 21, 71, 22, 73, 67, 50, 52,
	-9, -8, -8, 4, 14, 64, -14, -11, -8, 4,
	-10, 64, -8, 73, 67, -8, -8, -8, 4, -8,
	4, -8, 73, 4, -16, -16, -8, 4, -8, 73,
	73, -8, 55, -8, -3, 52, 55, -8, -8, 4,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -9, -8, 54, -8, -10, -8,
	55, 64, 13, 64, -1, -16, 16, 66, 64, 52,
	-1, 64, -9, -8, 54, 72, 72, -14, 73, -9,
	-13, -12, 6, 74, 73, 73, -8, 48, -15, 51,
	67, -16, 64, -10, -16, 54, 8, 74, 68, 54,
	-8, -16, -1, -8, -1, 65, 6, -8, -8, -1,
	-10, 65, -7, -16, 8, 74, 68, 54, -8, 4,
	4, 74, 8, -14, 55, -16, 55, -16, 54, -9,
	-9, 74, -8, 67, -8, -8, 68, -8, 4, -1,
	4, -8, 74, -8, 68, 68, -8, 4, 65, 64,
	65, 64, 66, 65, 29, 65, -6, -5, 45, 46,
	-6, -5, 74, -8, 68, 68, 64, 74, 74, 8,
	-16, 68, -16, 65, -8, 8, 74, 8, 74, 74,
	55, 68, 74, 55, 55, 65, 68, -1, -1, -8,
	4, 64, -8, 54, 68, -1, 64, 64, 74, 68,
	-12, 65, 74, 74, -8, -8, -8, 65, 65, 64,
	64, -1, 54, -16, 65, -1, -1, 64, 74, 74,
	55, 74, -1, -1, 65, -16, -1, 65, 65, -1,
	-8, 65, 65, 30, -1, 65, 74, 30, 64, 64,
	-1, -1, 65, 65,
}
var yyDef = [...]int{

	-2, -2, -2, 128, 129, 130, 132, 133, 4, 42,
	-2, 0, 9, 10, 49, 0, 0, 14, 42, 0,
	0, 53, 54, 0, 0, 0, 0, 0, 62, 63,
	64, 65, 0, 128, 128, 0, 0, 0, 0, 0,
	0, 0, 2, -2, 131, 0, 43, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	97, 98, 0, 0, 0, 0, 49, 0, 0, 49,
	11, 50, 12, 0, 0, -2, 0, 0, -2, -2,
	0, -2, 0, 49, 0, 55, 56, 57, -2, 0,
	-2, 0, 42, 0, 49, 39, 0, 53, 0, 0,
	35, 127, 128, 0, 5, 49, 128, 7, 0, 67,
	77, 78, 79, 80, 81, 82, 83, 84, -2, -2,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	99, 100, 101, 102, 0, 0, 0, 126, 8, -2,
	128, -2, 0, -2, 0, -2, 0, 0, -2, 49,
	0, 28, 0, 0, 0, 0, 0, 0, 42, 128,
	128, 40, 0, 76, 49, 49, 0, 0, 0, 0,
	0, 0, -2, 6, 0, 0, 0, 108, 112, 0,
	0, 0, 0, 0, 0, 15, 62, 0, 0, 0,
	45, 0, 0, 0, 0, 104, 111, 0, 0, 59,
	61, 0, 0, 0, 128, 0, 128, 0, 0, 0,
	0, 119, 0, 128, 0, 0, 36, -2, -2, 0,
	44, 66, 107, 0, 117, 118, 51, -2, 13, -2,
	26, -2, 0, 18, 0, 23, 31, 32, 0, 0,
	29, 30, 103, 0, 114, 115, -2, 0, 0, 0,
	0, 72, 0, 74, 38, 0, -2, 0, -2, 120,
	0, 37, 122, 0, 0, 27, 116, 0, 0, 0,
	0, -2, 0, 128, 113, 0, -2, -2, 0, 73,
	41, 75, -2, -2, 0, 0, 0, 25, 16, -2,
	-2, 0, 128, -2, 68, 0, 0, -2, 121, 123,
	0, 125, 0, 0, 22, -2, 34, 69, 70, 0,
	0, 17, 21, 0, 33, 71, 124, 0, -2, -2,
	0, 0, 20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	76, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 69, 3, 3, 3, 62, 71, 3,
	73, 74, 60, 58, 55, 59, 72, 61, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 54, 66,
	57, 52, 56, 53, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 67, 3, 68, 70, 3, 3, 3, 3, 3,
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
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:679
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:684
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:689
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:694
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:699
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:704
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].expr, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:709
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:714
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:725
		{
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:728
		{
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:733
		{
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:736
		{
		}
	}
	goto yystack /* stack new state and value */
}
