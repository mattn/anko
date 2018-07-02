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
	map_expr        map[ast.Expr]ast.Expr
	expr_idents     []string
	expr_type       string
	tok             ast.Token
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
const NILCOALESCE = 57369
const MODULE = 57370
const TRY = 57371
const CATCH = 57372
const FINALLY = 57373
const PLUSEQ = 57374
const MINUSEQ = 57375
const MULEQ = 57376
const DIVEQ = 57377
const ANDEQ = 57378
const OREQ = 57379
const BREAK = 57380
const CONTINUE = 57381
const PLUSPLUS = 57382
const MINUSMINUS = 57383
const POW = 57384
const SHIFTLEFT = 57385
const SHIFTRIGHT = 57386
const SWITCH = 57387
const CASE = 57388
const DEFAULT = 57389
const GO = 57390
const CHAN = 57391
const MAKE = 57392
const OPCHAN = 57393
const TYPE = 57394
const LEN = 57395
const DELETE = 57396
const UNARY = 57397

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
	"NILCOALESCE",
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

//line parser.go.y:803

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	58, 57,
	69, 5,
	-2, 1,
	-1, 11,
	58, 58,
	-2, 28,
	-1, 47,
	58, 57,
	-2, 5,
	-1, 95,
	58, 58,
	-2, 54,
	-1, 96,
	16, 49,
	58, 49,
	-2, 67,
	-1, 107,
	1, 72,
	8, 72,
	46, 72,
	47, 72,
	55, 72,
	57, 72,
	58, 72,
	67, 72,
	68, 72,
	69, 72,
	71, 72,
	73, 72,
	79, 72,
	-2, 67,
	-1, 109,
	1, 74,
	8, 74,
	46, 74,
	47, 74,
	55, 74,
	57, 74,
	58, 74,
	67, 74,
	68, 74,
	69, 74,
	71, 74,
	73, 74,
	79, 74,
	-2, 67,
	-1, 141,
	17, 0,
	18, 0,
	-2, 98,
	-1, 142,
	17, 0,
	18, 0,
	-2, 99,
	-1, 161,
	58, 58,
	-2, 54,
	-1, 228,
	73, 149,
	-2, 141,
	-1, 231,
	68, 149,
	-2, 141,
	-1, 244,
	58, 59,
	-2, 55,
	-1, 245,
	1, 56,
	46, 56,
	47, 56,
	55, 56,
	57, 56,
	58, 60,
	68, 56,
	69, 56,
	79, 56,
	-2, 67,
	-1, 254,
	1, 60,
	8, 60,
	46, 60,
	47, 60,
	58, 60,
	68, 60,
	69, 60,
	71, 60,
	73, 60,
	79, 60,
	-2, 67,
	-1, 269,
	1, 25,
	46, 25,
	47, 25,
	68, 25,
	69, 25,
	79, 25,
	-2, 117,
	-1, 271,
	1, 27,
	46, 27,
	47, 27,
	68, 27,
	69, 27,
	79, 27,
	-2, 119,
	-1, 280,
	68, 147,
	73, 147,
	-2, 142,
	-1, 307,
	1, 24,
	46, 24,
	47, 24,
	68, 24,
	69, 24,
	79, 24,
	-2, 116,
	-1, 308,
	1, 26,
	46, 26,
	47, 26,
	68, 26,
	69, 26,
	79, 26,
	-2, 118,
}

const yyPrivate = 57344

const yyLast = 3015

var yyAct = [...]int{

	88, 115, 1, 11, 229, 12, 265, 263, 7, 5,
	264, 266, 267, 231, 8, 49, 188, 89, 228, 8,
	95, 50, 99, 101, 288, 97, 104, 105, 106, 108,
	110, 103, 281, 102, 8, 44, 290, 113, 116, 8,
	241, 93, 236, 120, 8, 235, 123, 313, 11, 289,
	279, 87, 235, 127, 128, 130, 131, 49, 133, 134,
	135, 136, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 146, 147, 148, 149, 150, 151, 152, 6, 270,
	153, 154, 155, 156, 48, 158, 159, 161, 23, 24,
	30, 308, 160, 34, 189, 166, 174, 112, 102, 307,
	126, 172, 190, 176, 287, 192, 277, 40, 31, 32,
	33, 275, 252, 278, 114, 114, 185, 286, 187, 157,
	235, 182, 194, 49, 195, 193, 161, 121, 198, 162,
	201, 197, 268, 181, 41, 42, 119, 39, 43, 178,
	118, 169, 271, 117, 358, 357, 25, 29, 350, 183,
	348, 36, 291, 347, 37, 343, 35, 273, 354, 26,
	27, 28, 238, 111, 205, 207, 206, 342, 208, 210,
	211, 167, 161, 212, 226, 215, 49, 213, 220, 339,
	328, 324, 162, 71, 72, 73, 74, 75, 76, 232,
	221, 77, 78, 62, 320, 269, 244, 319, 293, 246,
	248, 114, 84, 251, 227, 114, 253, 237, 239, 283,
	216, 217, 57, 58, 59, 60, 61, 298, 203, 272,
	260, 83, 257, 53, 126, 56, 255, 214, 81, 79,
	205, 209, 164, 282, 284, 353, 331, 225, 322, 49,
	162, 114, 312, 311, 292, 276, 163, 98, 168, 162,
	122, 125, 114, 222, 126, 326, 306, 171, 86, 295,
	297, 296, 230, 230, 300, 9, 352, 304, 162, 302,
	301, 349, 305, 303, 266, 267, 261, 91, 285, 310,
	299, 204, 247, 314, 240, 189, 165, 51, 315, 49,
	126, 316, 317, 224, 223, 132, 90, 4, 38, 2,
	191, 47, 323, 46, 184, 94, 262, 280, 18, 327,
	280, 3, 0, 124, 329, 330, 0, 0, 0, 0,
	0, 0, 0, 0, 337, 338, 0, 0, 340, 341,
	0, 0, 0, 345, 344, 0, 346, 0, 0, 0,
	0, 114, 85, 65, 66, 68, 70, 80, 82, 0,
	0, 0, 0, 55, 0, 0, 355, 356, 71, 72,
	73, 74, 75, 76, 0, 0, 77, 78, 62, 63,
	64, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 54, 0, 335, 67, 69, 57, 58, 59,
	60, 61, 0, 0, 0, 0, 83, 334, 53, 0,
	56, 0, 0, 81, 79, 85, 65, 66, 68, 70,
	80, 82, 0, 0, 0, 0, 55, 0, 0, 0,
	0, 71, 72, 73, 74, 75, 76, 0, 0, 77,
	78, 62, 63, 64, 0, 0, 0, 0, 0, 0,
	84, 0, 0, 0, 0, 54, 0, 243, 67, 69,
	57, 58, 59, 60, 61, 0, 0, 0, 0, 83,
	242, 53, 0, 56, 0, 0, 81, 79, 85, 65,
	66, 68, 70, 80, 82, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 71, 72, 73, 74, 75, 76,
	0, 0, 77, 78, 62, 63, 64, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 0, 0, 54, 218,
	0, 67, 69, 57, 58, 59, 60, 61, 0, 0,
	0, 0, 83, 0, 53, 219, 56, 0, 0, 81,
	79, 85, 65, 66, 68, 70, 80, 82, 0, 0,
	0, 0, 55, 0, 0, 0, 0, 71, 72, 73,
	74, 75, 76, 0, 0, 77, 78, 62, 63, 64,
	0, 0, 0, 0, 0, 0, 84, 0, 0, 0,
	0, 54, 199, 0, 67, 69, 57, 58, 59, 60,
	61, 0, 0, 0, 0, 83, 0, 53, 200, 56,
	0, 0, 81, 79, 85, 65, 66, 68, 70, 80,
	82, 0, 0, 0, 0, 55, 0, 0, 0, 0,
	71, 72, 73, 74, 75, 76, 0, 0, 77, 78,
	62, 63, 64, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 0, 0, 54, 0, 0, 67, 69, 57,
	58, 59, 60, 61, 0, 0, 0, 0, 83, 351,
	53, 0, 56, 0, 0, 81, 79, 85, 65, 66,
	68, 70, 80, 82, 0, 0, 0, 0, 55, 0,
	0, 0, 0, 71, 72, 73, 74, 75, 76, 0,
	0, 77, 78, 62, 63, 64, 0, 0, 0, 0,
	0, 0, 84, 0, 0, 0, 0, 54, 0, 0,
	67, 69, 57, 58, 59, 60, 61, 0, 0, 0,
	0, 83, 336, 53, 0, 56, 0, 0, 81, 79,
	85, 65, 66, 68, 70, 80, 82, 0, 0, 0,
	0, 55, 0, 0, 0, 0, 71, 72, 73, 74,
	75, 76, 0, 0, 77, 78, 62, 63, 64, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	54, 0, 0, 67, 69, 57, 58, 59, 60, 61,
	0, 0, 0, 0, 83, 333, 53, 0, 56, 0,
	0, 81, 79, 85, 65, 66, 68, 70, 80, 82,
	0, 0, 0, 0, 55, 0, 0, 0, 0, 71,
	72, 73, 74, 75, 76, 0, 0, 77, 78, 62,
	63, 64, 0, 0, 0, 0, 0, 0, 84, 0,
	0, 0, 0, 54, 332, 0, 67, 69, 57, 58,
	59, 60, 61, 0, 0, 0, 0, 83, 0, 53,
	0, 56, 0, 0, 81, 79, 85, 65, 66, 68,
	70, 80, 82, 0, 0, 0, 0, 55, 0, 0,
	0, 0, 71, 72, 73, 74, 75, 76, 0, 0,
	77, 78, 62, 63, 64, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 54, 325, 0, 67,
	69, 57, 58, 59, 60, 61, 0, 0, 0, 0,
	83, 0, 53, 0, 56, 0, 0, 81, 79, 85,
	65, 66, 68, 70, 80, 82, 0, 0, 0, 0,
	55, 0, 0, 0, 0, 71, 72, 73, 74, 75,
	76, 0, 0, 77, 78, 62, 63, 64, 0, 0,
	0, 0, 0, 0, 84, 0, 0, 0, 0, 54,
	0, 0, 67, 69, 57, 58, 59, 60, 61, 0,
	321, 0, 0, 83, 0, 53, 0, 56, 0, 0,
	81, 79, 85, 65, 66, 68, 70, 80, 82, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 71, 72,
	73, 74, 75, 76, 0, 0, 77, 78, 62, 63,
	64, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 54, 0, 0, 67, 69, 57, 58, 59,
	60, 61, 0, 0, 0, 0, 83, 318, 53, 0,
	56, 0, 0, 81, 79, 85, 65, 66, 68, 70,
	80, 82, 0, 0, 0, 0, 55, 0, 0, 0,
	0, 71, 72, 73, 74, 75, 76, 0, 0, 77,
	78, 62, 63, 64, 0, 0, 0, 0, 0, 0,
	84, 0, 0, 0, 0, 54, 0, 0, 67, 69,
	57, 58, 59, 60, 61, 0, 0, 0, 0, 83,
	0, 53, 309, 56, 0, 0, 81, 79, 85, 65,
	66, 68, 70, 80, 82, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 71, 72, 73, 74, 75, 76,
	0, 0, 77, 78, 62, 63, 64, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 0, 0, 54, 0,
	0, 67, 69, 57, 58, 59, 60, 61, 0, 0,
	0, 0, 83, 0, 53, 294, 56, 0, 0, 81,
	79, 85, 65, 66, 68, 70, 80, 82, 0, 0,
	0, 0, 55, 0, 0, 0, 0, 71, 72, 73,
	74, 75, 76, 0, 0, 77, 78, 62, 63, 64,
	0, 0, 0, 0, 0, 0, 84, 0, 0, 0,
	0, 54, 0, 0, 67, 69, 57, 58, 59, 60,
	61, 0, 0, 0, 0, 83, 0, 53, 274, 56,
	0, 0, 81, 79, 85, 65, 66, 68, 70, 80,
	82, 0, 0, 0, 0, 55, 0, 0, 0, 0,
	71, 72, 73, 74, 75, 76, 0, 0, 77, 78,
	62, 63, 64, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 0, 0, 54, 0, 0, 67, 69, 57,
	58, 59, 60, 61, 0, 0, 0, 259, 83, 0,
	53, 0, 56, 0, 0, 81, 79, 85, 65, 66,
	68, 70, 80, 82, 0, 0, 0, 0, 55, 0,
	0, 0, 0, 71, 72, 73, 74, 75, 76, 0,
	0, 77, 78, 62, 63, 64, 0, 0, 0, 0,
	0, 0, 84, 0, 0, 0, 0, 54, 0, 0,
	67, 69, 57, 58, 59, 60, 61, 0, 258, 0,
	0, 83, 0, 53, 0, 56, 0, 0, 81, 79,
	85, 65, 66, 68, 70, 80, 82, 0, 0, 0,
	0, 55, 0, 0, 0, 0, 71, 72, 73, 74,
	75, 76, 0, 0, 77, 78, 62, 63, 64, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	54, 0, 0, 67, 69, 57, 58, 59, 60, 61,
	0, 256, 0, 0, 83, 0, 53, 0, 56, 0,
	0, 81, 79, 85, 65, 66, 68, 70, 80, 82,
	0, 0, 0, 0, 55, 0, 0, 0, 0, 71,
	72, 73, 74, 75, 76, 0, 0, 77, 78, 62,
	63, 64, 0, 0, 0, 0, 0, 0, 84, 0,
	0, 0, 0, 54, 0, 0, 67, 69, 57, 58,
	59, 60, 61, 0, 0, 0, 0, 83, 0, 53,
	250, 56, 0, 0, 81, 79, 85, 65, 66, 68,
	70, 80, 82, 0, 0, 0, 0, 55, 0, 0,
	0, 0, 71, 72, 73, 74, 75, 76, 0, 0,
	77, 78, 62, 63, 64, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 54, 0, 0, 67,
	69, 57, 58, 59, 60, 61, 0, 0, 0, 0,
	83, 234, 53, 0, 56, 0, 0, 81, 79, 85,
	65, 66, 68, 70, 80, 82, 0, 0, 0, 0,
	55, 0, 0, 0, 0, 71, 72, 73, 74, 75,
	76, 0, 0, 77, 78, 62, 63, 64, 0, 0,
	0, 0, 0, 0, 84, 0, 0, 0, 0, 54,
	233, 0, 67, 69, 57, 58, 59, 60, 61, 0,
	0, 0, 0, 83, 0, 53, 0, 56, 0, 0,
	81, 79, 85, 65, 66, 68, 70, 80, 82, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 71, 72,
	73, 74, 75, 76, 0, 0, 77, 78, 62, 63,
	64, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 54, 202, 0, 67, 69, 57, 58, 59,
	60, 61, 0, 0, 0, 0, 83, 0, 53, 0,
	56, 0, 0, 81, 79, 85, 65, 66, 68, 70,
	80, 82, 0, 0, 0, 0, 55, 0, 0, 0,
	0, 71, 72, 73, 74, 75, 76, 0, 0, 77,
	78, 62, 63, 64, 0, 0, 0, 0, 0, 0,
	84, 0, 0, 0, 0, 54, 0, 0, 67, 69,
	57, 58, 59, 60, 61, 0, 196, 0, 0, 83,
	0, 53, 0, 56, 0, 0, 81, 79, 85, 65,
	66, 68, 70, 80, 82, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 71, 72, 73, 74, 75, 76,
	0, 0, 77, 78, 62, 63, 64, 0, 0, 0,
	0, 0, 0, 84, 0, 0, 0, 0, 54, 0,
	0, 67, 69, 57, 58, 59, 60, 61, 0, 0,
	0, 0, 83, 186, 53, 0, 56, 0, 0, 81,
	79, 85, 65, 66, 68, 70, 80, 82, 0, 0,
	0, 0, 55, 0, 0, 0, 0, 71, 72, 73,
	74, 75, 76, 0, 0, 77, 78, 62, 63, 64,
	0, 0, 0, 0, 0, 0, 84, 0, 0, 0,
	0, 54, 0, 0, 67, 69, 57, 58, 59, 60,
	61, 0, 173, 0, 0, 83, 0, 53, 0, 56,
	0, 0, 81, 79, 85, 65, 66, 68, 70, 80,
	82, 0, 0, 0, 0, 55, 0, 0, 0, 0,
	71, 72, 73, 74, 75, 76, 0, 0, 77, 78,
	62, 63, 64, 0, 0, 0, 0, 0, 0, 84,
	0, 0, 0, 0, 54, 0, 0, 67, 69, 57,
	58, 59, 60, 61, 0, 170, 0, 0, 83, 0,
	53, 0, 56, 0, 0, 81, 79, 85, 65, 66,
	68, 70, 80, 82, 0, 0, 0, 0, 55, 0,
	0, 0, 0, 71, 72, 73, 74, 75, 76, 0,
	0, 77, 78, 62, 63, 64, 0, 0, 0, 0,
	0, 0, 84, 0, 0, 0, 52, 54, 0, 0,
	67, 69, 57, 58, 59, 60, 61, 0, 0, 0,
	0, 83, 0, 53, 0, 56, 0, 0, 81, 79,
	85, 65, 66, 68, 70, 80, 82, 0, 0, 0,
	0, 55, 0, 0, 0, 0, 71, 72, 73, 74,
	75, 76, 0, 0, 77, 78, 62, 63, 64, 0,
	0, 0, 0, 0, 0, 84, 0, 0, 0, 0,
	54, 0, 0, 67, 69, 57, 58, 59, 60, 61,
	0, 0, 0, 0, 83, 0, 53, 0, 56, 0,
	0, 81, 79, 85, 65, 66, 68, 70, 80, 82,
	0, 0, 0, 0, 55, 0, 0, 0, 0, 71,
	72, 73, 74, 75, 76, 0, 0, 77, 78, 62,
	63, 64, 0, 0, 0, 0, 0, 0, 84, 0,
	0, 0, 0, 54, 0, 0, 67, 69, 57, 58,
	59, 60, 61, 0, 0, 0, 0, 83, 0, 53,
	0, 180, 0, 0, 81, 79, 85, 65, 66, 68,
	70, 80, 82, 0, 0, 0, 0, 55, 0, 0,
	0, 0, 71, 72, 73, 74, 75, 76, 0, 0,
	77, 78, 62, 63, 64, 0, 0, 0, 0, 0,
	0, 84, 0, 0, 0, 0, 54, 0, 0, 67,
	69, 57, 58, 59, 60, 61, 0, 0, 0, 0,
	83, 0, 53, 0, 179, 0, 0, 81, 79, 85,
	65, 66, 68, 70, 80, 82, 0, 0, 0, 0,
	55, 0, 0, 0, 0, 71, 72, 73, 74, 75,
	76, 0, 0, 77, 78, 62, 63, 64, 0, 0,
	0, 0, 0, 0, 84, 0, 0, 0, 0, 54,
	0, 0, 67, 69, 57, 58, 59, 60, 61, 0,
	0, 0, 0, 175, 0, 53, 0, 56, 0, 0,
	81, 79, 85, 65, 66, 68, 70, 80, 82, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 71, 72,
	73, 74, 75, 76, 0, 0, 77, 78, 62, 63,
	64, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 0, 0, 0, 67, 69, 57, 58, 59,
	60, 61, 0, 0, 0, 0, 83, 0, 53, 0,
	56, 0, 0, 81, 79, 85, 65, 66, 68, 70,
	0, 82, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 71, 72, 73, 74, 75, 76, 0, 0, 77,
	78, 62, 63, 64, 0, 0, 0, 0, 0, 0,
	84, 0, 0, 0, 0, 0, 0, 0, 67, 69,
	57, 58, 59, 60, 61, 0, 0, 0, 0, 83,
	0, 53, 0, 56, 0, 0, 81, 79, 23, 24,
	30, 0, 0, 34, 15, 10, 16, 45, 0, 19,
	0, 0, 0, 0, 0, 0, 0, 40, 31, 32,
	33, 0, 17, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 13, 14, 0, 0, 0, 0, 0, 21,
	0, 0, 22, 0, 41, 42, 0, 39, 43, 0,
	0, 0, 0, 0, 0, 0, 25, 29, 0, 0,
	0, 36, 0, 0, 37, 0, 35, 0, 0, 26,
	27, 28, 85, 65, 66, 68, 70, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 71, 72,
	73, 74, 75, 76, 0, 0, 77, 78, 62, 63,
	64, 0, 0, 0, 0, 0, 0, 84, 0, 0,
	0, 0, 0, 0, 0, 67, 69, 57, 58, 59,
	60, 61, 68, 70, 0, 0, 83, 0, 53, 0,
	56, 0, 0, 81, 79, 71, 72, 73, 74, 75,
	76, 0, 0, 77, 78, 62, 63, 64, 0, 0,
	23, 24, 30, 0, 84, 34, 0, 0, 0, 0,
	0, 0, 67, 69, 57, 58, 59, 60, 61, 40,
	31, 32, 33, 83, 0, 53, 0, 56, 0, 0,
	81, 79, 23, 24, 30, 0, 0, 34, 0, 0,
	0, 0, 0, 0, 0, 0, 41, 42, 0, 39,
	43, 40, 31, 32, 33, 0, 0, 0, 25, 29,
	0, 0, 0, 36, 0, 0, 37, 0, 35, 249,
	0, 26, 27, 28, 0, 0, 0, 0, 41, 42,
	0, 39, 43, 0, 0, 177, 0, 23, 24, 30,
	25, 29, 34, 0, 0, 36, 0, 0, 37, 0,
	35, 0, 0, 26, 27, 28, 40, 31, 32, 33,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 23, 24, 30, 0, 0, 34, 0, 0,
	0, 0, 0, 41, 42, 0, 39, 43, 0, 0,
	129, 40, 31, 32, 33, 25, 29, 0, 0, 0,
	36, 0, 0, 37, 0, 35, 0, 0, 26, 27,
	28, 0, 0, 0, 0, 0, 0, 0, 41, 42,
	0, 39, 43, 0, 0, 0, 0, 254, 24, 30,
	25, 29, 34, 0, 0, 36, 0, 0, 37, 0,
	35, 0, 0, 26, 27, 28, 40, 31, 32, 33,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 245,
	24, 30, 0, 0, 34, 0, 0, 0, 0, 0,
	0, 0, 0, 41, 42, 0, 39, 43, 40, 31,
	32, 33, 0, 0, 0, 25, 29, 0, 0, 0,
	36, 0, 0, 37, 0, 35, 0, 0, 26, 27,
	28, 0, 0, 0, 0, 41, 42, 0, 39, 43,
	0, 0, 0, 0, 109, 24, 30, 25, 29, 34,
	0, 0, 36, 0, 0, 37, 0, 35, 0, 0,
	26, 27, 28, 40, 31, 32, 33, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 107, 24, 30, 0,
	0, 34, 0, 0, 0, 0, 0, 0, 0, 0,
	41, 42, 0, 39, 43, 40, 31, 32, 33, 0,
	0, 0, 25, 29, 0, 0, 0, 36, 0, 0,
	37, 0, 35, 0, 0, 26, 27, 28, 0, 0,
	0, 0, 41, 42, 0, 39, 43, 0, 0, 0,
	0, 100, 24, 30, 25, 29, 34, 0, 0, 36,
	0, 0, 37, 0, 35, 0, 0, 26, 27, 28,
	40, 31, 32, 33, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 96, 24, 30, 0, 0, 34, 0,
	0, 0, 0, 0, 0, 0, 0, 41, 42, 0,
	39, 43, 40, 31, 32, 33, 0, 0, 0, 25,
	29, 0, 0, 0, 36, 0, 0, 37, 0, 35,
	0, 0, 26, 27, 28, 0, 0, 0, 0, 41,
	42, 0, 39, 43, 71, 72, 73, 74, 75, 76,
	0, 25, 29, 0, 62, 0, 92, 0, 0, 37,
	0, 35, 0, 84, 26, 27, 28, 0, 71, 72,
	73, 74, 75, 76, 0, 59, 60, 61, 62, 0,
	0, 0, 83, 0, 53, 0, 56, 84, 0, 81,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 83, 0, 53, 0,
	56, 0, 0, 81, 79,
}
var yyPact = [...]int{

	-60, -1000, 2354, -60, -1000, -65, -65, -1000, -1000, -1000,
	283, 1901, 203, -1000, -1000, 2628, 2628, 292, 263, 2889,
	180, 2628, 2857, -39, -1000, 2628, 2628, 2628, 2802, 2770,
	-1000, -1000, -1000, -1000, 93, -65, -65, 2628, -1000, 73,
	70, 66, 2628, 57, 192, 2628, -1000, 2354, -65, -1000,
	196, -1000, 2628, 2593, 2628, 2628, 291, 2628, 2628, 2628,
	2628, 2628, 2628, 2628, 2628, 2628, 2628, 2628, 2628, 2628,
	2628, 2628, 2628, 2628, 2628, 2628, 2628, -1000, -1000, 2628,
	2628, 2628, 2628, 2628, 2628, 2628, 2628, 191, 1964, 1964,
	179, 219, -60, 232, 72, 1838, -39, 202, -60, 1775,
	26, 2153, 2538, 2628, 2936, 2936, 2936, -39, 2090, -39,
	2027, 283, 51, 2628, -65, 2628, 1712, 2628, 281, 53,
	1964, 2628, -65, 1649, -1000, 2628, -65, 1964, 515, 2628,
	1586, 2216, -1000, 2912, 2912, 2936, 2936, 2936, 1964, 151,
	151, 2463, 2463, 151, 151, 151, 151, 1964, 1964, 1964,
	1964, 1964, 1964, 1964, 2279, 1964, 2416, 210, 1964, 2416,
	-1000, 1964, -65, -60, 2628, -60, 163, -65, 2628, 2628,
	-60, 2628, 159, -65, 2628, 2628, 452, 2628, 182, 290,
	289, 166, 283, -40, -45, 1523, -1000, 1460, -29, -1000,
	281, 90, 280, -33, 389, 2715, -60, -1000, 278, 2506,
	-1000, 1397, 2628, 41, -1000, 2683, 158, 1334, 154, -1000,
	1271, 1208, 152, -1000, 246, 228, 124, 71, 84, -1000,
	1145, 40, -1000, -1000, -1000, 178, 35, 42, -65, -41,
	-65, -65, 141, 2628, -1000, 274, -1000, 46, -49, -22,
	94, -1000, -1000, 2628, 1964, -39, 130, -1000, 1082, -1000,
	-1000, 1964, -1000, 1964, -39, -1000, -60, -1000, -60, 2628,
	-1000, 213, -35, -1000, -1000, -1000, 2628, 199, 28, -1000,
	20, -1000, 1019, -1000, -1000, -1000, -60, 176, 175, -24,
	-65, -1000, 2628, -1000, 1964, -1000, -1000, 2628, -1000, -1000,
	2628, 2628, 956, -1000, -1000, 129, 126, 893, 171, -60,
	113, -1000, -1000, -1000, 830, 198, -60, -1000, -1000, -1000,
	112, -60, -60, 169, 767, 704, 326, 641, -1000, -1000,
	-1000, -60, -60, 111, -1000, -60, -60, -1000, -1000, 99,
	87, -60, 2628, -1000, -1000, 2628, -1000, 85, 82, 240,
	-1000, -1000, -1000, -1000, 80, 1964, 578, -1000, 235, 168,
	-1000, -1000, 91, -60, -60, 77, 76, -1000, -1000,
}
var yyPgo = [...]int{

	0, 2, 311, 265, 308, 10, 7, 306, 0, 35,
	5, 305, 304, 21, 16, 300, 298, 6, 299, 297,
	1, 4, 78, 8,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	4, 4, 7, 7, 7, 7, 7, 7, 7, 6,
	17, 5, 15, 15, 15, 12, 12, 12, 13, 13,
	13, 14, 14, 11, 10, 10, 10, 9, 9, 9,
	9, 16, 16, 16, 16, 16, 16, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 18, 18, 19, 19,
	19, 20, 20, 22, 22, 23, 21, 21, 21, 21,
}
var yyR2 = [...]int{

	0, 1, 2, 2, 3, 0, 4, 3, 3, 1,
	1, 2, 2, 5, 1, 4, 7, 9, 5, 13,
	12, 9, 8, 7, 6, 5, 6, 5, 1, 5,
	7, 5, 0, 1, 1, 2, 1, 2, 2, 4,
	4, 3, 0, 2, 3, 0, 3, 6, 0, 1,
	4, 1, 3, 3, 1, 4, 4, 0, 1, 4,
	4, 6, 5, 5, 6, 5, 5, 1, 1, 2,
	2, 2, 2, 4, 2, 4, 1, 1, 1, 1,
	5, 3, 3, 7, 8, 8, 9, 5, 5, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 3, 3, 3, 3, 5, 4, 5, 4,
	4, 4, 1, 4, 4, 5, 7, 5, 7, 9,
	7, 3, 2, 4, 6, 3, 0, 1, 2, 1,
	1, 0, 1, 1, 2, 1, 0, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -18, -2, -19, 69, -22, -23, 79, -3,
	11, -8, -10, 38, 39, 10, 12, 28, -4, 15,
	29, 45, 48, 4, 5, 62, 75, 76, 77, 63,
	6, 24, 25, 26, 9, 72, 67, 70, -16, 53,
	23, 50, 51, 54, -9, 13, -18, -19, -22, -23,
	-13, 4, 55, 72, 56, 27, 74, 61, 62, 63,
	64, 65, 42, 43, 44, 17, 18, 59, 19, 60,
	20, 32, 33, 34, 35, 36, 37, 40, 41, 78,
	21, 77, 22, 70, 51, 16, 55, -9, -8, -8,
	4, 14, 67, -13, -11, -8, 4, -10, 67, -8,
	4, -8, 72, 70, -8, -8, -8, 4, -8, 4,
	-8, 70, 4, -20, -22, -20, -8, 70, 70, 70,
	-8, 70, 58, -8, -3, 55, 58, -8, -8, 57,
	-8, -8, 4, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -9, -8, -8,
	-10, -8, 58, 67, 13, 67, -1, -22, 16, 69,
	67, 55, -1, 67, 70, 70, -8, 57, -9, 74,
	74, -13, 70, -9, -12, -8, 71, -8, -14, 4,
	49, -15, 52, 72, -8, -20, 67, -10, -20, 57,
	73, -8, 57, 8, 71, -20, -1, -8, -1, 68,
	-8, -8, -1, -10, 68, -20, -9, -9, 57, 73,
	-8, 8, 71, 4, 4, 71, 8, -13, 58, -21,
	-22, 58, -21, 57, 71, 74, 71, -14, 72, -14,
	4, 73, 71, 58, -8, 4, -1, 4, -8, 73,
	73, -8, 71, -8, 4, 68, 67, 68, 67, 69,
	68, 30, -7, -6, -5, -17, 46, 47, 8, 71,
	8, 71, -8, 73, 73, 71, 67, 71, 71, 8,
	-22, 73, -20, 68, -8, 4, 71, 58, 73, 71,
	58, 58, -8, 68, 73, -1, -1, -8, 4, 67,
	-20, -6, -17, -5, -8, -10, 57, 71, 71, 73,
	-1, 67, 67, 71, -8, -8, -8, -8, 71, 68,
	68, 67, 67, -1, 68, 57, 57, -1, 68, -1,
	-1, 67, 57, 71, 71, 58, 71, -1, -1, 68,
	-1, -1, 68, 68, -1, -8, -8, 68, 68, 31,
	68, 71, 31, 67, 67, -1, -1, 68, 68,
}
var yyDef = [...]int{

	136, -2, -2, 136, 137, 140, 139, 143, 145, 3,
	48, -2, 0, 9, 10, 57, 0, 0, 14, 48,
	0, 0, 0, 67, 68, 0, 0, 0, 0, 0,
	76, 77, 78, 79, 0, 141, 141, 0, 122, 0,
	0, 0, 0, 0, 0, 0, 2, -2, 138, 144,
	0, 49, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 110, 111, 0,
	0, 0, 0, 57, 0, 0, 57, 11, 58, 12,
	0, 0, 136, 0, 0, -2, -2, 0, 136, 0,
	67, 0, 0, 57, 69, 70, 71, -2, 0, -2,
	0, 48, 0, 57, 142, 45, 0, 0, 0, 42,
	132, 0, 141, 0, 4, 57, 141, 7, 0, 0,
	0, 81, 82, 90, 91, 92, 93, 94, 95, 96,
	97, -2, -2, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 112, 113, 114, 115, 0, 131, 135,
	8, -2, 141, 136, 0, 136, 0, 139, 0, 0,
	136, 57, 0, 141, 57, 57, 0, 0, 0, 0,
	0, 0, 48, 146, 146, 0, 89, 0, 0, 51,
	0, 0, 0, 0, 0, 0, 136, 6, 0, 0,
	121, 0, 0, 0, 119, 0, 0, 0, 0, 15,
	0, 0, 0, 53, 0, 32, 0, 0, 0, 120,
	0, 0, 117, 73, 75, 0, 0, 0, -2, 0,
	148, -2, 0, 0, 123, 0, 124, 0, 0, 0,
	0, 43, 133, 0, -2, -2, 0, 50, 0, 65,
	66, 80, 118, 59, -2, 13, 136, 31, 136, 0,
	18, 0, 141, 33, 34, 36, 57, 0, 0, -2,
	0, -2, 0, 62, 63, 116, 136, 0, 0, 0,
	-2, 87, 0, 88, 46, 52, 125, 0, 44, 127,
	0, 0, 0, 29, 64, 0, 0, 0, 0, 136,
	0, 35, 37, 38, 58, 0, 136, -2, -2, 61,
	0, 136, 136, 0, 0, 0, 0, 0, 134, 30,
	16, 136, 136, 0, 23, 136, 136, 41, 83, 0,
	0, 136, 0, 126, 128, 0, 130, 0, 0, 22,
	39, 40, 84, 85, 0, 47, 0, 17, 21, 0,
	86, 129, 0, 136, 136, 0, 0, 20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	79, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 75, 3, 3, 3, 65, 77, 3,
	70, 71, 63, 61, 58, 62, 74, 64, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 57, 69,
	60, 55, 59, 56, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 72, 3, 73, 76, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 67, 78, 68,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 66,
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
		//line parser.go.y:69
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:73
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:79
		{
			if yyDollar[2].stmt != nil {
				yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			} else {
				yyVAL.stmts = []ast.Stmt{}
			}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:90
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
				if l, ok := yylex.(*Lexer); ok {
					l.stmts = yyVAL.stmts
				}
			}
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:101
		{
			yyVAL.stmt = nil
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:106
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:111
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:115
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
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:132
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:137
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:142
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:157
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.stmt = &ast.ForStmt{Vars: yyDollar[2].expr_idents, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:172
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:177
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:182
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:192
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:197
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[5].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:202
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:207
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:212
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].expr.Position())
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:217
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:222
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:230
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 30:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:235
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:240
		{
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = yyDollar[4].compstmt
			}
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:250
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:254
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[1].stmt_case}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:258
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[1].stmt_default}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:262
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:266
		{
			yyVAL.stmt_cases = yyDollar[1].stmt_multi_case
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:270
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_multi_case...)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:274
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:285
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:291
		{
			var cases = []ast.Stmt{}
			for _, stmt := range yyDollar[2].expr_many {
				cases = append(cases, &ast.CaseStmt{Expr: stmt, Stmts: yyDollar[4].compstmt})
			}
			yyVAL.stmt_multi_case = cases
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:301
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[3].compstmt}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.array_count = ast.ArrayCount{Count: 0}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:310
		{
			yyVAL.array_count = ast.ArrayCount{Count: 1}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:314
		{
			yyVAL.array_count.Count = yyVAL.array_count.Count + 1
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:319
		{
			yyVAL.map_expr = make(map[ast.Expr]ast.Expr)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:323
		{
			mapExpr := make(map[ast.Expr]ast.Expr)
			mapExpr[yyDollar[1].expr] = yyDollar[3].expr
			yyVAL.map_expr = mapExpr
		}
	case 47:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:329
		{
			if len(yyDollar[1].map_expr) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyDollar[1].map_expr[yyDollar[4].expr] = yyDollar[6].expr
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:337
		{
			yyVAL.expr_idents = []string{}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:341
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:345
		{
			if len(yyDollar[1].expr_idents) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:354
		{
			yyVAL.expr_type = yyDollar[1].tok.Lit
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:358
		{
			yyVAL.expr_type = yyVAL.expr_type + "." + yyDollar[3].tok.Lit
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:363
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:369
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:373
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:377
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:382
		{
			yyVAL.exprs = nil
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:386
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:390
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:397
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 61:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:406
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:410
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: nil}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:414
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: nil, End: yyDollar[4].expr}
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:418
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:422
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:426
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:432
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:437
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:442
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:447
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:452
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
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
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:497
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:502
		{
			yyVAL.expr = &ast.NilCoalescingOpExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
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
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:537
		{
			yyVAL.expr = &ast.MapExpr{MapExpr: yyDollar[3].map_expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:542
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:547
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:552
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:557
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:562
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:567
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:572
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:577
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:582
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:587
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:592
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:597
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:602
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:607
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:612
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:617
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:622
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:627
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:632
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:637
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:642
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:647
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:652
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:657
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:662
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:667
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:672
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:677
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:682
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:687
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:692
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:697
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:702
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:707
		{
			yyVAL.expr = yyDollar[1].expr_slice
			yyVAL.expr.SetPosition(yyDollar[1].expr_slice.Position())
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:712
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:717
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 125:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:722
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:727
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:732
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 128:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:737
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 129:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:742
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 130:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:747
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].tok.Lit, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:752
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:757
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 133:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:762
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 134:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:767
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr, KeyExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:772
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: &ast.SliceExpr{Value: yyDollar[3].expr, Begin: nil, End: nil}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	}
	goto yystack /* stack new state and value */
}
