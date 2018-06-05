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

//line parser.go.y:823

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 139,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	58, 58,
	-2, 1,
	-1, 10,
	58, 59,
	-2, 28,
	-1, 46,
	58, 58,
	-2, 140,
	-1, 90,
	68, 3,
	-2, 139,
	-1, 93,
	58, 59,
	-2, 55,
	-1, 94,
	16, 50,
	58, 50,
	-2, 68,
	-1, 96,
	68, 3,
	-2, 139,
	-1, 105,
	1, 73,
	8, 73,
	46, 73,
	47, 73,
	55, 73,
	57, 73,
	58, 73,
	67, 73,
	68, 73,
	69, 73,
	71, 73,
	73, 73,
	79, 73,
	-2, 68,
	-1, 107,
	1, 75,
	8, 75,
	46, 75,
	47, 75,
	55, 75,
	57, 75,
	58, 75,
	67, 75,
	68, 75,
	69, 75,
	71, 75,
	73, 75,
	79, 75,
	-2, 68,
	-1, 138,
	17, 0,
	18, 0,
	-2, 101,
	-1, 139,
	17, 0,
	18, 0,
	-2, 102,
	-1, 158,
	58, 59,
	-2, 55,
	-1, 160,
	68, 3,
	-2, 139,
	-1, 162,
	68, 3,
	-2, 139,
	-1, 164,
	68, 1,
	-2, 46,
	-1, 167,
	68, 3,
	-2, 139,
	-1, 194,
	68, 3,
	-2, 139,
	-1, 243,
	58, 60,
	-2, 56,
	-1, 244,
	1, 57,
	46, 57,
	47, 57,
	55, 57,
	57, 57,
	58, 61,
	68, 57,
	69, 57,
	79, 57,
	-2, 68,
	-1, 253,
	1, 61,
	8, 61,
	46, 61,
	47, 61,
	58, 61,
	68, 61,
	69, 61,
	71, 61,
	73, 61,
	79, 61,
	-2, 68,
	-1, 255,
	68, 3,
	-2, 139,
	-1, 257,
	68, 3,
	-2, 139,
	-1, 271,
	1, 25,
	46, 25,
	47, 25,
	68, 25,
	69, 25,
	79, 25,
	-2, 120,
	-1, 273,
	1, 27,
	46, 27,
	47, 27,
	68, 27,
	69, 27,
	79, 27,
	-2, 122,
	-1, 278,
	68, 3,
	-2, 139,
	-1, 301,
	68, 3,
	-2, 139,
	-1, 305,
	1, 24,
	46, 24,
	47, 24,
	68, 24,
	69, 24,
	79, 24,
	-2, 119,
	-1, 306,
	1, 26,
	46, 26,
	47, 26,
	68, 26,
	69, 26,
	79, 26,
	-2, 121,
	-1, 309,
	68, 3,
	-2, 139,
	-1, 310,
	68, 3,
	-2, 139,
	-1, 321,
	68, 3,
	-2, 139,
	-1, 322,
	68, 3,
	-2, 139,
	-1, 326,
	46, 3,
	47, 3,
	68, 3,
	-2, 139,
	-1, 330,
	68, 3,
	-2, 139,
	-1, 338,
	46, 3,
	47, 3,
	68, 3,
	-2, 139,
	-1, 339,
	46, 3,
	47, 3,
	68, 3,
	-2, 139,
	-1, 353,
	68, 3,
	-2, 139,
	-1, 354,
	68, 3,
	-2, 139,
}

const yyPrivate = 57344

const yyLast = 3172

var yyAct = [...]int{

	86, 182, 263, 10, 262, 264, 186, 6, 290, 283,
	240, 11, 48, 1, 311, 306, 87, 7, 43, 93,
	305, 97, 99, 230, 279, 102, 103, 104, 106, 108,
	95, 91, 235, 85, 6, 234, 101, 113, 100, 187,
	277, 171, 117, 100, 7, 120, 251, 10, 110, 281,
	179, 124, 125, 127, 128, 118, 130, 131, 132, 133,
	134, 135, 136, 137, 138, 139, 140, 141, 142, 143,
	144, 145, 146, 147, 148, 149, 228, 292, 150, 151,
	152, 153, 2, 155, 156, 158, 45, 6, 188, 289,
	291, 190, 272, 234, 270, 226, 157, 7, 221, 123,
	154, 173, 288, 116, 163, 234, 115, 237, 114, 201,
	169, 191, 280, 358, 109, 185, 183, 111, 112, 192,
	175, 166, 178, 158, 265, 266, 357, 199, 350, 346,
	180, 345, 342, 341, 195, 337, 327, 320, 319, 295,
	300, 285, 159, 259, 159, 123, 261, 256, 159, 254,
	213, 207, 354, 161, 168, 273, 353, 271, 225, 159,
	330, 222, 205, 322, 310, 10, 209, 210, 165, 158,
	309, 278, 202, 164, 204, 220, 206, 293, 314, 160,
	212, 211, 96, 122, 159, 119, 123, 325, 304, 232,
	216, 217, 227, 84, 243, 236, 238, 8, 247, 265,
	266, 250, 193, 301, 252, 352, 196, 162, 245, 347,
	123, 260, 69, 70, 71, 72, 73, 74, 269, 274,
	267, 268, 60, 5, 89, 183, 287, 246, 47, 239,
	187, 82, 49, 286, 224, 223, 129, 88, 37, 189,
	181, 4, 203, 294, 121, 46, 92, 214, 17, 3,
	81, 0, 51, 215, 54, 0, 0, 79, 77, 299,
	0, 0, 0, 229, 231, 0, 302, 0, 0, 297,
	47, 298, 0, 0, 0, 0, 0, 303, 0, 0,
	0, 0, 0, 252, 0, 0, 313, 0, 0, 0,
	315, 0, 308, 316, 317, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 282, 0, 284, 0, 323, 0, 0, 0, 0,
	0, 0, 0, 328, 329, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 344, 335, 336, 0, 0, 0,
	340, 0, 0, 0, 343, 0, 0, 0, 0, 0,
	0, 0, 348, 349, 0, 83, 63, 64, 66, 68,
	78, 80, 0, 0, 0, 0, 53, 355, 356, 0,
	0, 69, 70, 71, 72, 73, 74, 0, 0, 75,
	76, 60, 61, 62, 0, 0, 0, 326, 0, 0,
	82, 0, 0, 0, 0, 52, 0, 333, 65, 67,
	55, 56, 57, 58, 59, 0, 0, 338, 339, 81,
	332, 51, 0, 54, 0, 0, 79, 77, 83, 63,
	64, 66, 68, 78, 80, 0, 0, 0, 0, 53,
	0, 0, 0, 0, 69, 70, 71, 72, 73, 74,
	0, 0, 75, 76, 60, 61, 62, 0, 0, 0,
	0, 0, 0, 82, 0, 0, 0, 0, 52, 0,
	242, 65, 67, 55, 56, 57, 58, 59, 0, 0,
	0, 0, 81, 241, 51, 0, 54, 0, 0, 79,
	77, 83, 63, 64, 66, 68, 78, 80, 0, 0,
	0, 0, 53, 0, 0, 0, 0, 69, 70, 71,
	72, 73, 74, 0, 0, 75, 76, 60, 61, 62,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 52, 218, 0, 65, 67, 55, 56, 57, 58,
	59, 0, 0, 0, 0, 81, 0, 51, 219, 54,
	0, 0, 79, 77, 83, 63, 64, 66, 68, 78,
	80, 0, 0, 0, 0, 53, 0, 0, 0, 0,
	69, 70, 71, 72, 73, 74, 0, 0, 75, 76,
	60, 61, 62, 0, 0, 0, 0, 0, 0, 82,
	0, 0, 0, 0, 52, 197, 0, 65, 67, 55,
	56, 57, 58, 59, 0, 0, 0, 0, 81, 0,
	51, 198, 54, 0, 0, 79, 77, 83, 63, 64,
	66, 68, 78, 80, 0, 0, 0, 0, 53, 0,
	0, 0, 0, 69, 70, 71, 72, 73, 74, 0,
	0, 75, 76, 60, 61, 62, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 52, 0, 0,
	65, 67, 55, 56, 57, 58, 59, 0, 0, 0,
	0, 81, 351, 51, 0, 54, 0, 0, 79, 77,
	83, 63, 64, 66, 68, 78, 80, 0, 0, 0,
	0, 53, 0, 0, 0, 0, 69, 70, 71, 72,
	73, 74, 0, 0, 75, 76, 60, 61, 62, 0,
	0, 0, 0, 0, 0, 82, 0, 0, 0, 0,
	52, 0, 0, 65, 67, 55, 56, 57, 58, 59,
	0, 0, 0, 0, 81, 334, 51, 0, 54, 0,
	0, 79, 77, 83, 63, 64, 66, 68, 78, 80,
	0, 0, 0, 0, 53, 0, 0, 0, 0, 69,
	70, 71, 72, 73, 74, 0, 0, 75, 76, 60,
	61, 62, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 52, 0, 0, 65, 67, 55, 56,
	57, 58, 59, 0, 0, 0, 0, 81, 331, 51,
	0, 54, 0, 0, 79, 77, 83, 63, 64, 66,
	68, 78, 80, 0, 0, 0, 0, 53, 0, 0,
	0, 0, 69, 70, 71, 72, 73, 74, 0, 0,
	75, 76, 60, 61, 62, 0, 0, 0, 0, 0,
	0, 82, 0, 0, 0, 0, 52, 324, 0, 65,
	67, 55, 56, 57, 58, 59, 0, 0, 0, 0,
	81, 0, 51, 0, 54, 0, 0, 79, 77, 83,
	63, 64, 66, 68, 78, 80, 0, 0, 0, 0,
	53, 0, 0, 0, 0, 69, 70, 71, 72, 73,
	74, 0, 0, 75, 76, 60, 61, 62, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 52,
	0, 0, 65, 67, 55, 56, 57, 58, 59, 0,
	321, 0, 0, 81, 0, 51, 0, 54, 0, 0,
	79, 77, 83, 63, 64, 66, 68, 78, 80, 0,
	0, 0, 0, 53, 0, 0, 0, 0, 69, 70,
	71, 72, 73, 74, 0, 0, 75, 76, 60, 61,
	62, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 52, 0, 0, 65, 67, 55, 56, 57,
	58, 59, 0, 0, 0, 0, 81, 318, 51, 0,
	54, 0, 0, 79, 77, 83, 63, 64, 66, 68,
	78, 80, 0, 0, 0, 0, 53, 0, 0, 0,
	0, 69, 70, 71, 72, 73, 74, 0, 0, 75,
	76, 60, 61, 62, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 52, 0, 0, 65, 67,
	55, 56, 57, 58, 59, 0, 0, 0, 0, 81,
	0, 51, 307, 54, 0, 0, 79, 77, 83, 63,
	64, 66, 68, 78, 80, 0, 0, 0, 0, 53,
	0, 0, 0, 0, 69, 70, 71, 72, 73, 74,
	0, 0, 75, 76, 60, 61, 62, 0, 0, 0,
	0, 0, 0, 82, 0, 0, 0, 0, 52, 0,
	0, 65, 67, 55, 56, 57, 58, 59, 0, 0,
	0, 0, 81, 0, 51, 296, 54, 0, 0, 79,
	77, 83, 63, 64, 66, 68, 78, 80, 0, 0,
	0, 0, 53, 0, 0, 0, 0, 69, 70, 71,
	72, 73, 74, 0, 0, 75, 76, 60, 61, 62,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 52, 0, 0, 65, 67, 55, 56, 57, 58,
	59, 0, 0, 0, 0, 81, 0, 51, 276, 54,
	0, 0, 79, 77, 83, 63, 64, 66, 68, 78,
	80, 0, 0, 0, 0, 53, 0, 0, 0, 0,
	69, 70, 71, 72, 73, 74, 0, 0, 75, 76,
	60, 61, 62, 0, 0, 0, 0, 0, 0, 82,
	0, 0, 0, 0, 52, 0, 0, 65, 67, 55,
	56, 57, 58, 59, 0, 0, 0, 258, 81, 0,
	51, 0, 54, 0, 0, 79, 77, 83, 63, 64,
	66, 68, 78, 80, 0, 0, 0, 0, 53, 0,
	0, 0, 0, 69, 70, 71, 72, 73, 74, 0,
	0, 75, 76, 60, 61, 62, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 52, 0, 0,
	65, 67, 55, 56, 57, 58, 59, 0, 257, 0,
	0, 81, 0, 51, 0, 54, 0, 0, 79, 77,
	83, 63, 64, 66, 68, 78, 80, 0, 0, 0,
	0, 53, 0, 0, 0, 0, 69, 70, 71, 72,
	73, 74, 0, 0, 75, 76, 60, 61, 62, 0,
	0, 0, 0, 0, 0, 82, 0, 0, 0, 0,
	52, 0, 0, 65, 67, 55, 56, 57, 58, 59,
	0, 255, 0, 0, 81, 0, 51, 0, 54, 0,
	0, 79, 77, 83, 63, 64, 66, 68, 78, 80,
	0, 0, 0, 0, 53, 0, 0, 0, 0, 69,
	70, 71, 72, 73, 74, 0, 0, 75, 76, 60,
	61, 62, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 52, 0, 0, 65, 67, 55, 56,
	57, 58, 59, 0, 0, 0, 0, 81, 0, 51,
	249, 54, 0, 0, 79, 77, 83, 63, 64, 66,
	68, 78, 80, 0, 0, 0, 0, 53, 0, 0,
	0, 0, 69, 70, 71, 72, 73, 74, 0, 0,
	75, 76, 60, 61, 62, 0, 0, 0, 0, 0,
	0, 82, 0, 0, 0, 0, 52, 0, 0, 65,
	67, 55, 56, 57, 58, 59, 0, 0, 0, 0,
	81, 233, 51, 0, 54, 0, 0, 79, 77, 83,
	63, 64, 66, 68, 78, 80, 0, 0, 0, 0,
	53, 0, 0, 0, 0, 69, 70, 71, 72, 73,
	74, 0, 0, 75, 76, 60, 61, 62, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 52,
	200, 0, 65, 67, 55, 56, 57, 58, 59, 0,
	0, 0, 0, 81, 0, 51, 0, 54, 0, 0,
	79, 77, 83, 63, 64, 66, 68, 78, 80, 0,
	0, 0, 0, 53, 0, 0, 0, 0, 69, 70,
	71, 72, 73, 74, 0, 0, 75, 76, 60, 61,
	62, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 52, 0, 0, 65, 67, 55, 56, 57,
	58, 59, 0, 194, 0, 0, 81, 0, 51, 0,
	54, 0, 0, 79, 77, 83, 63, 64, 66, 68,
	78, 80, 0, 0, 0, 0, 53, 0, 0, 0,
	0, 69, 70, 71, 72, 73, 74, 0, 0, 75,
	76, 60, 61, 62, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 52, 0, 0, 65, 67,
	55, 56, 57, 58, 59, 0, 0, 0, 0, 81,
	184, 51, 0, 54, 0, 0, 79, 77, 83, 63,
	64, 66, 68, 78, 80, 0, 0, 0, 0, 53,
	0, 0, 0, 0, 69, 70, 71, 72, 73, 74,
	0, 0, 75, 76, 60, 61, 62, 0, 0, 0,
	0, 0, 0, 82, 0, 0, 0, 0, 52, 0,
	0, 65, 67, 55, 56, 57, 58, 59, 0, 170,
	0, 0, 81, 0, 51, 0, 54, 0, 0, 79,
	77, 83, 63, 64, 66, 68, 78, 80, 0, 0,
	0, 0, 53, 0, 0, 0, 0, 69, 70, 71,
	72, 73, 74, 0, 0, 75, 76, 60, 61, 62,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 52, 0, 0, 65, 67, 55, 56, 57, 58,
	59, 0, 167, 0, 0, 81, 0, 51, 0, 54,
	0, 0, 79, 77, 22, 23, 29, 0, 0, 33,
	14, 9, 15, 44, 0, 18, 0, 0, 0, 0,
	0, 0, 0, 39, 30, 31, 32, 0, 16, 19,
	0, 0, 0, 0, 0, 0, 0, 0, 12, 13,
	0, 0, 0, 0, 0, 20, 0, 0, 21, 0,
	40, 41, 0, 38, 42, 0, 0, 0, 0, 0,
	0, 0, 24, 28, 0, 0, 0, 35, 0, 6,
	36, 0, 34, 0, 0, 25, 26, 27, 0, 7,
	83, 63, 64, 66, 68, 78, 80, 0, 0, 0,
	0, 53, 0, 0, 0, 0, 69, 70, 71, 72,
	73, 74, 0, 0, 75, 76, 60, 61, 62, 0,
	0, 0, 0, 0, 0, 82, 0, 0, 0, 50,
	52, 0, 0, 65, 67, 55, 56, 57, 58, 59,
	0, 0, 0, 0, 81, 0, 51, 0, 54, 0,
	0, 79, 77, 83, 63, 64, 66, 68, 78, 80,
	0, 0, 0, 0, 53, 0, 0, 0, 0, 69,
	70, 71, 72, 73, 74, 0, 0, 75, 76, 60,
	61, 62, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 52, 0, 0, 65, 67, 55, 56,
	57, 58, 59, 0, 0, 0, 0, 81, 0, 51,
	0, 54, 0, 0, 79, 77, 83, 63, 64, 66,
	68, 78, 80, 0, 0, 0, 0, 53, 0, 0,
	0, 0, 69, 70, 71, 72, 73, 74, 0, 0,
	75, 76, 60, 61, 62, 0, 0, 0, 0, 0,
	0, 82, 0, 0, 0, 0, 52, 0, 0, 65,
	67, 55, 56, 57, 58, 59, 0, 0, 0, 0,
	81, 0, 51, 0, 177, 0, 0, 79, 77, 83,
	63, 64, 66, 68, 78, 80, 0, 0, 0, 0,
	53, 0, 0, 0, 0, 69, 70, 71, 72, 73,
	74, 0, 0, 75, 76, 60, 61, 62, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 52,
	0, 0, 65, 67, 55, 56, 57, 58, 59, 0,
	0, 0, 0, 81, 0, 51, 0, 176, 0, 0,
	79, 77, 83, 63, 64, 66, 68, 78, 80, 0,
	0, 0, 0, 53, 0, 0, 0, 0, 69, 70,
	71, 72, 73, 74, 0, 0, 75, 76, 60, 61,
	62, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 52, 0, 0, 65, 67, 55, 56, 57,
	58, 59, 0, 0, 0, 0, 172, 0, 51, 0,
	54, 0, 0, 79, 77, 83, 63, 64, 66, 68,
	78, 80, 0, 0, 0, 0, 53, 0, 0, 0,
	0, 69, 70, 71, 72, 73, 74, 0, 0, 75,
	76, 60, 61, 62, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 0, 0, 0, 65, 67,
	55, 56, 57, 58, 59, 0, 0, 0, 0, 81,
	0, 51, 0, 54, 0, 0, 79, 77, 22, 23,
	208, 0, 0, 33, 14, 9, 15, 44, 0, 18,
	0, 0, 0, 0, 0, 0, 0, 39, 30, 31,
	32, 0, 16, 19, 0, 0, 0, 0, 0, 0,
	0, 0, 12, 13, 0, 0, 0, 0, 0, 20,
	0, 0, 21, 0, 40, 41, 0, 38, 42, 0,
	0, 0, 0, 0, 0, 0, 24, 28, 0, 0,
	0, 35, 0, 0, 36, 0, 34, 0, 0, 25,
	26, 27, 83, 63, 64, 66, 68, 0, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 70,
	71, 72, 73, 74, 0, 0, 75, 76, 60, 61,
	62, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 65, 67, 55, 56, 57,
	58, 59, 0, 0, 0, 0, 81, 0, 51, 0,
	54, 0, 0, 79, 77, 22, 23, 29, 0, 0,
	33, 14, 9, 15, 44, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 39, 30, 31, 32, 0, 16,
	19, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	13, 0, 0, 0, 0, 0, 20, 0, 0, 21,
	0, 40, 41, 0, 38, 42, 0, 0, 0, 0,
	0, 0, 0, 24, 28, 0, 0, 0, 35, 0,
	0, 36, 0, 34, 0, 0, 25, 26, 27, 83,
	63, 64, 66, 68, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 70, 71, 72, 73,
	74, 0, 0, 75, 76, 60, 61, 62, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 0,
	0, 0, 65, 67, 55, 56, 57, 58, 59, 66,
	68, 0, 0, 81, 0, 51, 0, 54, 0, 0,
	79, 77, 69, 70, 71, 72, 73, 74, 0, 0,
	75, 76, 60, 61, 62, 0, 0, 253, 23, 29,
	0, 82, 33, 0, 0, 0, 0, 0, 0, 65,
	67, 55, 56, 57, 58, 59, 39, 30, 31, 32,
	81, 0, 51, 0, 54, 0, 0, 79, 77, 22,
	23, 29, 0, 0, 33, 0, 0, 0, 0, 0,
	0, 0, 0, 40, 41, 0, 38, 42, 39, 30,
	31, 32, 0, 0, 0, 24, 28, 0, 0, 0,
	35, 0, 0, 36, 0, 34, 312, 0, 25, 26,
	27, 0, 0, 0, 0, 40, 41, 0, 38, 42,
	0, 0, 0, 0, 22, 23, 29, 24, 28, 33,
	0, 0, 35, 0, 0, 36, 0, 34, 275, 0,
	25, 26, 27, 39, 30, 31, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	40, 41, 0, 38, 42, 0, 0, 0, 0, 0,
	0, 0, 24, 28, 0, 0, 0, 35, 0, 0,
	36, 0, 34, 248, 0, 25, 26, 27, 69, 70,
	71, 72, 73, 74, 0, 0, 75, 76, 60, 0,
	0, 0, 0, 22, 23, 29, 0, 82, 33, 0,
	0, 0, 0, 0, 0, 0, 0, 55, 56, 57,
	58, 59, 39, 30, 31, 32, 81, 0, 51, 0,
	54, 0, 0, 79, 77, 0, 0, 0, 22, 23,
	29, 0, 0, 33, 0, 0, 0, 0, 0, 40,
	41, 0, 38, 42, 0, 0, 174, 39, 30, 31,
	32, 24, 28, 0, 0, 0, 35, 0, 0, 36,
	0, 34, 0, 0, 25, 26, 27, 0, 0, 0,
	0, 0, 0, 0, 40, 41, 0, 38, 42, 0,
	0, 126, 0, 22, 23, 29, 24, 28, 33, 0,
	0, 35, 0, 0, 36, 0, 34, 0, 0, 25,
	26, 27, 39, 30, 31, 32, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 253, 23, 29, 0, 0,
	33, 0, 0, 0, 0, 0, 0, 0, 0, 40,
	41, 0, 38, 42, 39, 30, 31, 32, 0, 0,
	0, 24, 28, 0, 0, 0, 35, 0, 0, 36,
	0, 34, 0, 0, 25, 26, 27, 0, 0, 0,
	0, 40, 41, 0, 38, 42, 0, 0, 0, 0,
	244, 23, 29, 24, 28, 33, 0, 0, 35, 0,
	0, 36, 0, 34, 0, 0, 25, 26, 27, 39,
	30, 31, 32, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 107, 23, 29, 0, 0, 33, 0, 0,
	0, 0, 0, 0, 0, 0, 40, 41, 0, 38,
	42, 39, 30, 31, 32, 0, 0, 0, 24, 28,
	0, 0, 0, 35, 0, 0, 36, 0, 34, 0,
	0, 25, 26, 27, 0, 0, 0, 0, 40, 41,
	0, 38, 42, 0, 0, 0, 0, 105, 23, 29,
	24, 28, 33, 0, 0, 35, 0, 0, 36, 0,
	34, 0, 0, 25, 26, 27, 39, 30, 31, 32,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 98,
	23, 29, 0, 0, 33, 0, 0, 0, 0, 0,
	0, 0, 0, 40, 41, 0, 38, 42, 39, 30,
	31, 32, 0, 0, 0, 24, 28, 0, 0, 0,
	35, 0, 0, 36, 0, 34, 0, 0, 25, 26,
	27, 0, 0, 0, 0, 40, 41, 0, 38, 42,
	0, 0, 0, 0, 94, 23, 29, 24, 28, 33,
	0, 0, 35, 0, 0, 36, 0, 34, 0, 0,
	25, 26, 27, 39, 30, 31, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	40, 41, 0, 38, 42, 69, 70, 71, 72, 73,
	74, 0, 24, 28, 0, 60, 0, 90, 0, 0,
	36, 0, 34, 0, 82, 25, 26, 27, 0, 0,
	0, 0, 0, 0, 0, 0, 57, 58, 59, 0,
	0, 0, 0, 81, 0, 51, 0, 54, 0, 0,
	79, 77,
}
var yyPact = [...]int{

	-62, -1000, 2391, -62, -62, -1000, -1000, -1000, -1000, 228,
	1864, 138, -1000, -1000, 2809, 2809, 233, 210, 3070, 115,
	2809, 3015, -34, -1000, 2809, 2809, 2809, 2983, 2928, -1000,
	-1000, -1000, -1000, 44, -62, -62, 2809, -1000, 38, 36,
	33, 2809, -15, 127, 2809, -1000, 1800, -1000, 128, -1000,
	2809, 2754, 2809, 2809, 232, 2809, 2809, 2809, 2809, 2809,
	2809, 2809, 2809, 2809, 2809, 2809, 2809, 2809, 2809, 2809,
	2809, 2809, 2809, 2809, 2809, -1000, -1000, 2809, 2809, 2809,
	2809, 2809, 2809, 2809, 2809, 126, 1927, 1927, 112, 140,
	-62, 152, 52, 1725, -34, 99, -62, 1662, -29, 2116,
	2719, 2809, 180, 180, 180, -34, 2053, -34, 1990, 228,
	-20, 2809, 219, 1599, 2809, 226, 39, 1927, 2809, -62,
	1536, -1000, 2809, -62, 1927, 528, 2809, 1473, 2179, -1000,
	3093, 3093, 180, 180, 180, 1927, 2676, 2676, 2500, 2500,
	2676, 2676, 2676, 2676, 1927, 1927, 1927, 1927, 1927, 1927,
	1927, 2316, 1927, 2453, 101, 1927, 2453, -1000, 1927, -62,
	-62, 2809, -62, 83, 2254, 2809, 2809, -62, 2809, 82,
	-62, 2809, 2809, 465, 2809, 90, 231, 230, 87, 228,
	18, -35, -1000, 132, -1000, 1410, -39, -1000, 226, 35,
	225, -63, 402, 2896, -62, -1000, 223, 2630, -1000, 1347,
	2809, -25, -1000, 2841, 81, 1284, 79, -1000, 132, 1221,
	1158, 75, -1000, 181, 78, 153, 86, 84, 2575, -1000,
	1095, -31, -1000, -1000, -1000, 104, -47, 41, -62, -64,
	-62, 73, 2809, -1000, 222, -1000, 31, -65, 19, 119,
	-1000, -1000, 2809, 1927, -34, 71, -1000, 1032, -1000, -1000,
	1927, -1000, 1927, -34, -1000, -62, -1000, -62, 2809, -1000,
	136, -1000, -1000, -1000, -1000, 2809, 131, -1000, -1000, -1000,
	-51, -1000, -56, -1000, 969, -1000, -1000, -1000, -62, 103,
	97, -57, 2543, -1000, 110, -1000, 1927, -1000, -1000, 2809,
	-1000, -1000, 2809, 2809, 906, -1000, -1000, 70, 69, 843,
	96, -62, 780, 130, -62, -1000, -1000, -1000, 68, -62,
	-62, 93, -1000, -1000, -1000, 717, 339, 654, -1000, -1000,
	-1000, -62, -62, 67, -62, -62, -62, -1000, 65, 64,
	-62, -1000, -1000, 2809, -1000, 63, 61, 178, -62, -62,
	-1000, -1000, -1000, 60, 591, -1000, 174, 89, -1000, -1000,
	-1000, -1000, 85, -62, -62, 58, 45, -1000, -1000,
}
var yyPgo = [...]int{

	0, 13, 249, 197, 248, 5, 4, 247, 0, 18,
	11, 246, 1, 240, 12, 6, 239, 238, 2, 82,
	241, 223,
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
	8, 8, 8, 8, 8, 8, 8, 8, 8, 19,
	19, 20, 20, 21, 21,
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
	1, 5, 3, 3, 7, 8, 8, 9, 5, 6,
	5, 6, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 3, 3, 3, 3, 5,
	4, 5, 4, 4, 4, 1, 4, 4, 5, 7,
	5, 7, 9, 7, 3, 2, 4, 6, 3, 0,
	1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -19, -2, -20, -21, 69, 79, -3, 11,
	-8, -10, 38, 39, 10, 12, 28, -4, 15, 29,
	45, 48, 4, 5, 62, 75, 76, 77, 63, 6,
	24, 25, 26, 9, 72, 67, 70, -17, 53, 23,
	50, 51, 54, -9, 13, -19, -20, -21, -14, 4,
	55, 72, 56, 27, 74, 61, 62, 63, 64, 65,
	42, 43, 44, 17, 18, 59, 19, 60, 20, 32,
	33, 34, 35, 36, 37, 40, 41, 78, 21, 77,
	22, 70, 51, 16, 55, -9, -8, -8, 4, 14,
	67, -14, -11, -8, 4, -10, 67, -8, 4, -8,
	72, 70, -8, -8, -8, 4, -8, 4, -8, 70,
	4, -19, -19, -8, 70, 70, 70, -8, 70, 58,
	-8, -3, 55, 58, -8, -8, 57, -8, -8, 4,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -9, -8, -8, -10, -8, 58,
	67, 13, 67, -1, -19, 16, 69, 67, 55, -1,
	67, 70, 70, -8, 57, -9, 74, 74, -14, 70,
	-9, -13, -12, 6, 71, -8, -15, 4, 49, -16,
	52, 72, -8, -19, 67, -10, -19, 57, 73, -8,
	57, 8, 71, -19, -1, -8, -1, 68, 6, -8,
	-8, -1, -10, 68, -7, -19, -9, -9, 57, 73,
	-8, 8, 71, 4, 4, 71, 8, -14, 58, -19,
	58, -19, 57, 71, 74, 71, -15, 72, -15, 4,
	73, 71, 58, -8, 4, -1, 4, -8, 73, 73,
	-8, 71, -8, 4, 68, 67, 68, 67, 69, 68,
	30, 68, -6, -18, -5, 46, 47, -6, -5, -18,
	8, 71, 8, 71, -8, 73, 73, 71, 67, 71,
	71, 8, -19, 73, -19, 68, -8, 4, 71, 58,
	73, 71, 58, 58, -8, 68, 73, -1, -1, -8,
	4, 67, -8, -10, 57, 71, 71, 73, -1, 67,
	67, 71, 73, -12, 68, -8, -8, -8, 71, 68,
	68, 67, 67, -1, 57, 57, -19, 68, -1, -1,
	67, 71, 71, 58, 71, -1, -1, 68, -19, -19,
	-1, 68, 68, -1, -8, 68, 68, 31, -1, -1,
	68, 71, 31, 67, 67, -1, -1, 68, 68,
}
var yyDef = [...]int{

	-2, -2, -2, 139, 140, 141, 143, 144, 4, 49,
	-2, 0, 9, 10, 58, 0, 0, 14, 49, 0,
	0, 0, 68, 69, 0, 0, 0, 0, 0, 77,
	78, 79, 80, 0, 139, 139, 0, 125, 0, 0,
	0, 0, 0, 0, 0, 2, -2, 142, 0, 50,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 113, 114, 0, 0, 0,
	0, 58, 0, 0, 58, 11, 59, 12, 0, 0,
	-2, 0, 0, -2, -2, 0, -2, 0, 68, 0,
	0, 58, 70, 71, 72, -2, 0, -2, 0, 49,
	0, 58, 46, 0, 0, 0, 42, 135, 0, 139,
	0, 5, 58, 139, 7, 0, 0, 0, 82, 83,
	93, 94, 95, 96, 97, 98, 99, 100, -2, -2,
	103, 104, 105, 106, 107, 108, 109, 110, 111, 112,
	115, 116, 117, 118, 0, 134, 138, 8, -2, 139,
	-2, 0, -2, 0, -2, 0, 0, -2, 58, 0,
	32, 58, 58, 0, 0, 0, 0, 0, 0, 49,
	139, 139, 47, 0, 92, 0, 0, 52, 0, 0,
	0, 0, 0, 0, -2, 6, 0, 0, 124, 0,
	0, 0, 122, 0, 0, 0, 0, 15, 77, 0,
	0, 0, 54, 0, 0, 0, 0, 0, 0, 123,
	0, 0, 120, 74, 76, 0, 0, 0, 139, 0,
	139, 0, 0, 126, 0, 127, 0, 0, 0, 0,
	43, 136, 0, -2, -2, 0, 51, 0, 66, 67,
	81, 121, 60, -2, 13, -2, 30, -2, 0, 18,
	0, 23, 35, 37, 38, 58, 0, 33, 34, 36,
	0, -2, 0, -2, 0, 63, 64, 119, -2, 0,
	0, 0, 0, 88, 0, 90, 45, 53, 128, 0,
	44, 130, 0, 0, 0, 31, 65, 0, 0, 0,
	0, -2, 59, 0, 139, -2, -2, 62, 0, -2,
	-2, 0, 89, 48, 91, 0, 0, 0, 137, 29,
	16, -2, -2, 0, 139, 139, -2, 84, 0, 0,
	-2, 129, 131, 0, 133, 0, 0, 22, -2, -2,
	41, 85, 86, 0, 0, 17, 21, 0, 39, 40,
	87, 132, 0, -2, -2, 0, 0, 20, 19,
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
		//line parser.go.y:74
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:78
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:83
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:90
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:97
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
		//line parser.go.y:108
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:113
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:117
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
		//line parser.go.y:129
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:134
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:149
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:154
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:159
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:164
		{
			yyVAL.stmt = &ast.ForStmt{Vars: yyDollar[2].expr_idents, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:169
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:174
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:179
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:184
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:189
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:194
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:199
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:204
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:209
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:214
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].expr.Position())
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:219
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:232
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:237
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
		//line parser.go.y:246
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:252
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:260
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:264
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:268
		{
			yyVAL.stmt_cases = yyDollar[2].stmt_multi_case
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:272
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_multi_case...)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:276
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
		//line parser.go.y:287
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:293
		{
			var cases = []ast.Stmt{}
			for _, stmt := range yyDollar[2].expr_many {
				cases = append(cases, &ast.CaseStmt{Expr: stmt, Stmts: yyDollar[5].compstmt})
			}
			yyVAL.stmt_multi_case = cases
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:303
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:308
		{
			yyVAL.array_count = ast.ArrayCount{Count: 0}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:312
		{
			yyVAL.array_count = ast.ArrayCount{Count: 1}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:316
		{
			yyVAL.array_count.Count = yyVAL.array_count.Count + 1
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:322
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:331
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:335
		{
			if len(yyDollar[1].expr_pairs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:343
		{
			yyVAL.expr_idents = []string{}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:347
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:351
		{
			if len(yyDollar[1].expr_idents) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:360
		{
			yyVAL.expr_type = yyDollar[1].tok.Lit
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:364
		{
			yyVAL.expr_type = yyVAL.expr_type + "." + yyDollar[3].tok.Lit
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:369
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:375
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:379
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:383
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:388
		{
			yyVAL.exprs = nil
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:392
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:396
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:403
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 62:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:412
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:416
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: nil}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:420
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: nil, End: yyDollar[4].expr}
		}
	case 65:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:424
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:428
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
		}
	case 67:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:432
		{
			yyVAL.expr_slice = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:438
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:443
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:448
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:453
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:458
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:463
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:468
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:473
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:478
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:483
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:488
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:493
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:498
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:503
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:508
		{
			yyVAL.expr = &ast.NilCoalescingOpExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:513
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:518
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 85:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:523
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 86:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:528
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 87:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:533
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 88:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:538
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 89:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:543
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:548
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
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:557
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
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:566
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:571
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:576
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:581
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:586
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:591
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:596
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:601
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:606
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:611
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:616
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:621
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:626
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:631
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:636
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:641
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:646
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:651
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:656
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:661
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:666
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:671
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:676
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:681
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:686
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:691
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:696
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:701
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:706
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:711
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:716
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:721
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:726
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:731
		{
			yyVAL.expr = yyDollar[1].expr_slice
			yyVAL.expr.SetPosition(yyDollar[1].expr_slice.Position())
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:736
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:741
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 128:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:746
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 129:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:751
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 130:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:756
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 131:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:761
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 132:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:766
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 133:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:771
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].tok.Lit, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:776
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:781
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 136:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:786
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 137:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:791
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr, KeyExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:796
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: &ast.SliceExpr{Value: yyDollar[3].expr, Begin: nil, End: nil}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:809
		{
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:812
		{
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:817
		{
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:820
		{
		}
	}
	goto yystack /* stack new state and value */
}
