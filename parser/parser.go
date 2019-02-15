//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:43
type yySymType struct {
	yys int
	tok ast.Token

	compstmt            ast.Stmt
	stmts               ast.Stmt
	stmt                ast.Stmt
	stmt_var_or_lets    ast.Stmt
	stmt_var            ast.Stmt
	stmt_lets           ast.Stmt
	stmt_if             ast.Stmt
	stmt_for            ast.Stmt
	stmt_switch         ast.Stmt
	stmt_switch_cases   ast.Stmt
	stmt_switch_case    ast.Stmt
	stmt_switch_default ast.Stmt

	exprs                []ast.Expr
	expr                 ast.Expr
	expr_idents          []string
	type_data            *ast.TypeStruct
	slice_count          int
	expr_member_or_ident ast.Expr
	expr_member          *ast.MemberExpr
	expr_ident           *ast.IdentExpr
	expr_literals        ast.Expr
	expr_map             *ast.MapExpr
	expr_slice           ast.Expr
	expr_unary           ast.Expr
	expr_binary          ast.Expr
	expr_lets            ast.Expr

	op_binary     ast.Operator
	op_comparison ast.Operator
	op_add        ast.Operator
	op_multiply   ast.Operator
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
const CLOSE = 57396
const MAP = 57397
const UNARY = 57398

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
	"CLOSE",
	"MAP",
	"'='",
	"':'",
	"'?'",
	"'<'",
	"'>'",
	"'+'",
	"'-'",
	"'|'",
	"'^'",
	"'*'",
	"'/'",
	"'%'",
	"'&'",
	"UNARY",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"','",
	"';'",
	"'['",
	"']'",
	"'.'",
	"'!'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:1027

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	56, 55,
	74, 55,
	75, 5,
	-2, 1,
	-1, 23,
	74, 56,
	-2, 26,
	-1, 27,
	16, 92,
	-2, 55,
	-1, 65,
	56, 55,
	74, 55,
	-2, 5,
	-1, 117,
	16, 93,
	74, 93,
	-2, 106,
	-1, 121,
	4, 101,
	48, 101,
	55, 101,
	-2, 67,
	-1, 255,
	71, 171,
	77, 171,
	-2, 163,
	-1, 273,
	71, 171,
	-2, 163,
	-1, 277,
	1, 58,
	8, 58,
	45, 58,
	46, 58,
	56, 58,
	57, 58,
	71, 58,
	73, 58,
	74, 58,
	75, 58,
	77, 58,
	80, 58,
	-2, 104,
	-1, 281,
	1, 17,
	45, 17,
	46, 17,
	71, 17,
	75, 17,
	80, 17,
	-2, 72,
	-1, 283,
	1, 19,
	45, 19,
	46, 19,
	71, 19,
	75, 19,
	80, 19,
	-2, 74,
	-1, 312,
	71, 169,
	77, 169,
	-2, 164,
	-1, 328,
	1, 16,
	45, 16,
	46, 16,
	71, 16,
	75, 16,
	80, 16,
	-2, 71,
	-1, 329,
	1, 18,
	45, 18,
	46, 18,
	71, 18,
	75, 18,
	80, 18,
	-2, 73,
}

const yyPrivate = 57344

const yyLast = 3559

var yyAct = [...]int{

	69, 256, 7, 23, 221, 48, 6, 139, 304, 67,
	305, 36, 66, 307, 306, 70, 113, 5, 74, 68,
	273, 8, 8, 360, 255, 206, 8, 111, 114, 118,
	8, 206, 313, 121, 33, 131, 8, 315, 206, 123,
	122, 137, 1, 271, 206, 268, 269, 124, 8, 123,
	206, 145, 211, 133, 205, 130, 209, 146, 147, 148,
	149, 81, 206, 138, 140, 82, 23, 85, 127, 67,
	125, 377, 199, 267, 194, 143, 155, 156, 206, 159,
	160, 161, 321, 163, 165, 166, 352, 162, 143, 168,
	169, 170, 171, 172, 173, 174, 175, 176, 177, 178,
	179, 180, 181, 182, 183, 184, 185, 186, 187, 188,
	189, 190, 223, 311, 129, 152, 198, 219, 329, 193,
	84, 128, 282, 142, 280, 328, 67, 261, 204, 202,
	195, 126, 195, 316, 214, 216, 201, 309, 213, 416,
	222, 143, 130, 288, 86, 87, 132, 120, 224, 153,
	123, 225, 253, 125, 136, 192, 203, 135, 233, 228,
	229, 127, 127, 134, 127, 240, 76, 75, 415, 411,
	127, 127, 407, 127, 207, 208, 81, 210, 310, 195,
	82, 406, 85, 217, 218, 326, 220, 283, 143, 281,
	143, 404, 262, 143, 226, 243, 67, 129, 247, 398,
	250, 234, 123, 244, 128, 397, 236, 123, 251, 393,
	392, 257, 123, 258, 126, 119, 265, 252, 195, 254,
	157, 391, 389, 272, 382, 130, 276, 378, 257, 275,
	374, 370, 284, 368, 241, 367, 287, 277, 366, 245,
	289, 364, 84, 335, 127, 323, 203, 295, 292, 299,
	301, 327, 412, 193, 286, 278, 242, 260, 227, 410,
	67, 381, 312, 225, 314, 317, 86, 87, 97, 98,
	320, 237, 143, 362, 351, 325, 350, 158, 308, 151,
	312, 324, 72, 9, 346, 307, 306, 409, 405, 296,
	279, 94, 95, 96, 99, 10, 77, 294, 81, 270,
	259, 141, 82, 342, 85, 193, 127, 193, 347, 167,
	123, 343, 345, 348, 344, 67, 353, 71, 60, 322,
	4, 257, 359, 116, 65, 361, 2, 61, 331, 62,
	64, 63, 46, 45, 365, 44, 334, 43, 30, 49,
	336, 337, 29, 339, 303, 22, 21, 20, 193, 150,
	127, 349, 25, 24, 3, 0, 0, 383, 0, 384,
	0, 0, 0, 354, 0, 0, 0, 123, 386, 0,
	363, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 369, 0, 371, 372, 0, 0, 222, 403, 375,
	0, 402, 0, 379, 380, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 408, 388, 0, 0, 0, 257,
	0, 0, 0, 0, 0, 0, 394, 0, 0, 395,
	396, 0, 0, 0, 399, 84, 103, 104, 108, 106,
	110, 109, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 88, 89, 91, 92, 93, 90, 0, 0, 86,
	87, 97, 98, 413, 0, 414, 0, 0, 0, 83,
	0, 0, 0, 0, 0, 0, 0, 79, 105, 107,
	100, 101, 102, 0, 94, 95, 96, 99, 0, 0,
	0, 81, 357, 358, 0, 82, 0, 85, 84, 103,
	104, 108, 106, 110, 109, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 88, 89, 91, 92, 93, 90,
	0, 0, 86, 87, 97, 98, 0, 0, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 0, 0, 356,
	79, 105, 107, 100, 101, 102, 0, 94, 95, 96,
	99, 0, 0, 0, 81, 0, 0, 0, 82, 355,
	85, 84, 103, 104, 108, 106, 110, 109, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 88, 89, 91,
	92, 93, 90, 0, 0, 86, 87, 97, 98, 0,
	0, 0, 0, 0, 0, 83, 0, 0, 0, 0,
	0, 0, 333, 79, 105, 107, 100, 101, 102, 0,
	94, 95, 96, 99, 0, 0, 0, 81, 0, 0,
	0, 82, 332, 85, 84, 103, 104, 108, 106, 110,
	109, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	88, 89, 91, 92, 93, 90, 0, 0, 86, 87,
	97, 98, 0, 0, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 0, 0, 264, 79, 105, 107, 100,
	101, 102, 0, 94, 95, 96, 99, 0, 0, 0,
	81, 0, 0, 0, 82, 263, 85, 84, 103, 104,
	108, 106, 110, 109, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 88, 89, 91, 92, 93, 90, 0,
	0, 86, 87, 97, 98, 0, 0, 0, 0, 0,
	0, 83, 0, 0, 0, 0, 0, 0, 239, 79,
	105, 107, 100, 101, 102, 0, 94, 95, 96, 99,
	0, 0, 0, 81, 0, 0, 0, 82, 238, 85,
	84, 103, 104, 108, 106, 110, 109, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 88, 89, 91, 92,
	93, 90, 0, 0, 86, 87, 97, 98, 0, 0,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 0,
	0, 0, 79, 105, 107, 100, 101, 102, 0, 94,
	95, 96, 99, 0, 0, 0, 81, 230, 231, 0,
	82, 0, 85, 84, 103, 104, 108, 106, 110, 109,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 88,
	89, 91, 92, 93, 90, 0, 0, 86, 87, 97,
	98, 0, 0, 0, 0, 0, 0, 83, 0, 0,
	0, 0, 0, 78, 0, 79, 105, 107, 100, 101,
	102, 0, 94, 95, 96, 99, 0, 196, 0, 81,
	0, 0, 0, 82, 0, 85, 35, 51, 52, 0,
	0, 31, 13, 47, 14, 26, 0, 27, 0, 0,
	0, 0, 0, 0, 0, 38, 53, 54, 55, 0,
	15, 16, 0, 0, 0, 0, 0, 0, 0, 0,
	11, 12, 0, 0, 0, 0, 28, 0, 0, 17,
	0, 39, 40, 0, 37, 18, 19, 41, 0, 0,
	0, 0, 0, 0, 50, 0, 57, 59, 0, 0,
	58, 0, 42, 0, 34, 0, 0, 0, 32, 0,
	0, 56, 84, 103, 104, 108, 106, 110, 109, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 88, 89,
	91, 92, 93, 90, 0, 0, 86, 87, 97, 98,
	0, 0, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 0, 0, 0, 79, 105, 107, 100, 101, 102,
	0, 94, 95, 96, 99, 0, 0, 0, 81, 401,
	0, 0, 82, 0, 85, 84, 103, 104, 108, 106,
	110, 109, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 88, 89, 91, 92, 93, 90, 0, 0, 86,
	87, 97, 98, 0, 0, 0, 0, 0, 0, 83,
	0, 0, 0, 0, 0, 0, 0, 79, 105, 107,
	100, 101, 102, 0, 94, 95, 96, 99, 0, 0,
	0, 81, 0, 0, 0, 82, 400, 85, 84, 103,
	104, 108, 106, 110, 109, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 88, 89, 91, 92, 93, 90,
	0, 0, 86, 87, 97, 98, 0, 0, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 0, 0, 0,
	79, 105, 107, 100, 101, 102, 0, 94, 95, 96,
	99, 0, 0, 0, 81, 0, 0, 0, 82, 390,
	85, 84, 103, 104, 108, 106, 110, 109, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 88, 89, 91,
	92, 93, 90, 0, 0, 86, 87, 97, 98, 0,
	0, 0, 0, 0, 0, 83, 0, 0, 0, 0,
	0, 0, 387, 79, 105, 107, 100, 101, 102, 0,
	94, 95, 96, 99, 0, 0, 0, 81, 0, 0,
	0, 82, 0, 85, 84, 103, 104, 108, 106, 110,
	109, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	88, 89, 91, 92, 93, 90, 0, 0, 86, 87,
	97, 98, 0, 0, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 0, 0, 0, 79, 105, 107, 100,
	101, 102, 0, 94, 95, 96, 99, 0, 0, 0,
	81, 385, 0, 0, 82, 0, 85, 84, 103, 104,
	108, 106, 110, 109, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 88, 89, 91, 92, 93, 90, 0,
	0, 86, 87, 97, 98, 0, 0, 0, 0, 0,
	0, 83, 0, 0, 0, 0, 0, 0, 376, 79,
	105, 107, 100, 101, 102, 0, 94, 95, 96, 99,
	0, 0, 0, 81, 0, 0, 0, 82, 0, 85,
	84, 103, 104, 108, 106, 110, 109, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 88, 89, 91, 92,
	93, 90, 0, 0, 86, 87, 97, 98, 0, 0,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 0,
	0, 0, 79, 105, 107, 100, 101, 102, 0, 94,
	95, 96, 99, 0, 373, 0, 81, 0, 0, 0,
	82, 0, 85, 84, 103, 104, 108, 106, 110, 109,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 88,
	89, 91, 92, 93, 90, 0, 0, 86, 87, 97,
	98, 0, 0, 0, 0, 0, 0, 83, 0, 0,
	0, 0, 0, 0, 0, 79, 105, 107, 100, 101,
	102, 0, 94, 95, 96, 99, 0, 340, 0, 81,
	0, 0, 0, 82, 0, 85, 84, 103, 104, 108,
	106, 110, 109, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 88, 89, 91, 92, 93, 90, 0, 0,
	86, 87, 97, 98, 0, 0, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 0, 0, 0, 79, 105,
	107, 100, 101, 102, 0, 94, 95, 96, 99, 0,
	338, 0, 81, 0, 0, 0, 82, 0, 85, 84,
	103, 104, 108, 106, 110, 109, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 88, 89, 91, 92, 93,
	90, 0, 0, 86, 87, 97, 98, 0, 0, 0,
	0, 0, 0, 83, 0, 0, 0, 0, 0, 0,
	0, 79, 105, 107, 100, 101, 102, 0, 94, 95,
	96, 99, 0, 0, 0, 81, 330, 0, 0, 82,
	0, 85, 84, 103, 104, 108, 106, 110, 109, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 88, 89,
	91, 92, 93, 90, 0, 0, 86, 87, 97, 98,
	0, 0, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 0, 0, 0, 79, 105, 107, 100, 101, 102,
	0, 94, 95, 96, 99, 0, 0, 0, 81, 0,
	0, 0, 82, 319, 85, 84, 103, 104, 108, 106,
	110, 109, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 88, 89, 91, 92, 93, 90, 0, 0, 86,
	87, 97, 98, 0, 0, 0, 0, 0, 0, 83,
	0, 0, 0, 0, 0, 0, 0, 79, 105, 107,
	100, 101, 102, 0, 94, 95, 96, 99, 0, 0,
	0, 81, 0, 0, 302, 82, 0, 85, 84, 103,
	104, 108, 106, 110, 109, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 88, 89, 91, 92, 93, 90,
	0, 0, 86, 87, 97, 98, 0, 0, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 0, 0, 0,
	79, 105, 107, 100, 101, 102, 0, 94, 95, 96,
	99, 0, 297, 0, 81, 0, 0, 0, 82, 0,
	85, 84, 103, 104, 108, 106, 110, 109, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 88, 89, 91,
	92, 93, 90, 0, 0, 86, 87, 97, 98, 0,
	0, 0, 0, 0, 0, 83, 0, 0, 0, 0,
	0, 0, 0, 79, 105, 107, 100, 101, 102, 0,
	94, 95, 96, 99, 0, 293, 0, 81, 0, 0,
	0, 82, 0, 85, 84, 103, 104, 108, 106, 110,
	109, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	88, 89, 91, 92, 93, 90, 0, 0, 86, 87,
	97, 98, 0, 0, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 0, 0, 0, 79, 105, 107, 100,
	101, 102, 0, 94, 95, 96, 99, 0, 0, 0,
	81, 0, 0, 0, 82, 291, 85, 84, 103, 104,
	108, 106, 110, 109, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 88, 89, 91, 92, 93, 90, 0,
	0, 86, 87, 97, 98, 0, 0, 0, 0, 0,
	0, 83, 0, 0, 0, 0, 0, 0, 0, 79,
	105, 107, 100, 101, 102, 0, 94, 95, 96, 99,
	0, 285, 0, 81, 0, 0, 0, 82, 0, 85,
	84, 103, 104, 108, 106, 110, 109, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 88, 89, 91, 92,
	93, 90, 0, 0, 86, 87, 97, 98, 0, 0,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 0,
	0, 274, 79, 105, 107, 100, 101, 102, 0, 94,
	95, 96, 99, 0, 0, 0, 81, 0, 0, 0,
	82, 0, 85, 84, 103, 104, 108, 106, 110, 109,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 88,
	89, 91, 92, 93, 90, 0, 0, 86, 87, 97,
	98, 0, 0, 0, 0, 0, 0, 83, 0, 0,
	0, 0, 0, 0, 0, 79, 105, 107, 100, 101,
	102, 0, 94, 95, 96, 99, 0, 0, 0, 81,
	266, 0, 0, 82, 0, 85, 84, 103, 104, 108,
	106, 110, 109, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 88, 89, 91, 92, 93, 90, 0, 0,
	86, 87, 97, 98, 0, 0, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 0, 0, 0, 79, 105,
	107, 100, 101, 102, 0, 94, 95, 96, 99, 0,
	0, 0, 81, 0, 0, 248, 82, 0, 85, 84,
	103, 104, 108, 106, 110, 109, 0, 0, 0, 0,
	80, 0, 0, 0, 0, 88, 89, 91, 92, 93,
	90, 0, 0, 86, 87, 97, 98, 0, 0, 0,
	0, 0, 0, 83, 0, 0, 0, 0, 0, 0,
	235, 79, 105, 107, 100, 101, 102, 0, 94, 95,
	96, 99, 0, 0, 0, 81, 0, 0, 0, 82,
	0, 85, 84, 103, 104, 108, 106, 110, 109, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 88, 89,
	91, 92, 93, 90, 0, 0, 86, 87, 97, 98,
	0, 0, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 0, 0, 0, 79, 105, 107, 100, 101, 102,
	0, 94, 95, 96, 99, 0, 0, 0, 81, 232,
	0, 0, 82, 0, 85, 84, 103, 104, 108, 106,
	110, 109, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 88, 89, 91, 92, 93, 90, 0, 0, 86,
	87, 97, 98, 0, 0, 0, 0, 0, 0, 83,
	0, 0, 0, 0, 0, 0, 0, 79, 105, 107,
	100, 101, 102, 0, 94, 95, 96, 99, 0, 0,
	0, 81, 212, 0, 0, 82, 0, 85, 84, 103,
	104, 108, 106, 110, 109, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 88, 89, 91, 92, 93, 90,
	0, 0, 86, 87, 97, 98, 0, 0, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 0, 0, 0,
	79, 105, 107, 100, 101, 102, 0, 94, 95, 96,
	99, 0, 200, 0, 81, 0, 0, 0, 82, 0,
	85, 84, 103, 104, 108, 106, 110, 109, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 88, 89, 91,
	92, 93, 90, 0, 0, 86, 87, 97, 98, 0,
	0, 0, 0, 0, 0, 83, 0, 0, 0, 0,
	0, 0, 0, 79, 105, 107, 100, 101, 102, 0,
	94, 95, 96, 99, 0, 191, 0, 81, 0, 0,
	0, 82, 0, 85, 84, 103, 104, 108, 106, 110,
	109, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	88, 89, 91, 92, 93, 90, 0, 0, 86, 87,
	97, 98, 0, 0, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 0, 78, 0, 79, 105, 107, 100,
	101, 102, 0, 94, 95, 96, 99, 0, 0, 0,
	81, 0, 0, 0, 82, 0, 85, 84, 103, 104,
	108, 106, 110, 109, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 88, 89, 91, 92, 93, 90, 0,
	0, 86, 87, 97, 98, 0, 0, 0, 0, 0,
	0, 83, 0, 0, 0, 0, 0, 0, 0, 79,
	105, 107, 100, 101, 102, 0, 94, 95, 96, 99,
	0, 0, 0, 81, 0, 0, 0, 82, 0, 85,
	84, 103, 104, 108, 106, 110, 109, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 88, 89, 91, 92,
	93, 90, 0, 0, 86, 87, 97, 98, 0, 0,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 0,
	0, 0, 79, 105, 107, 100, 101, 102, 0, 94,
	95, 96, 99, 0, 0, 0, 154, 0, 0, 0,
	82, 0, 85, 84, 103, 104, 108, 106, 110, 109,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 86, 87, 97,
	98, 0, 0, 0, 0, 0, 0, 83, 0, 0,
	0, 0, 0, 0, 0, 79, 105, 107, 100, 101,
	102, 0, 94, 95, 96, 99, 0, 0, 0, 81,
	0, 0, 0, 82, 0, 85, 84, 103, 104, 108,
	106, 110, 109, 0, 0, 0, 0, 80, 0, 117,
	51, 52, 0, 0, 31, 0, 47, 0, 0, 0,
	86, 87, 97, 98, 0, 0, 0, 0, 38, 53,
	54, 55, 0, 0, 0, 0, 0, 0, 79, 105,
	107, 100, 101, 102, 0, 94, 95, 96, 99, 0,
	0, 0, 81, 0, 39, 40, 82, 37, 85, 0,
	41, 0, 0, 0, 0, 0, 0, 50, 0, 57,
	59, 0, 0, 58, 0, 112, 0, 34, 0, 0,
	115, 32, 0, 0, 56, 35, 51, 52, 0, 0,
	31, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 38, 53, 54, 55, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 35, 51,
	52, 0, 0, 31, 0, 0, 0, 0, 0, 0,
	39, 40, 0, 37, 0, 0, 41, 38, 53, 54,
	55, 0, 0, 50, 0, 57, 59, 0, 0, 58,
	0, 42, 0, 34, 0, 0, 0, 32, 318, 0,
	56, 0, 0, 39, 40, 0, 37, 0, 0, 41,
	0, 0, 0, 0, 84, 0, 50, 0, 57, 59,
	0, 0, 58, 0, 42, 0, 34, 35, 51, 52,
	32, 290, 31, 56, 0, 0, 0, 0, 86, 87,
	97, 98, 0, 0, 0, 0, 38, 53, 54, 55,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 100,
	101, 102, 0, 94, 95, 96, 99, 0, 0, 0,
	81, 0, 39, 40, 82, 37, 85, 0, 41, 0,
	84, 103, 104, 108, 106, 50, 109, 57, 59, 0,
	0, 58, 0, 42, 0, 34, 0, 0, 249, 32,
	0, 0, 56, 0, 86, 87, 97, 98, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 105, 107, 100, 101, 102, 0, 94,
	95, 96, 99, 35, 51, 52, 81, 0, 31, 0,
	82, 0, 85, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 38, 53, 54, 55, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 35, 51,
	52, 0, 0, 31, 0, 0, 0, 0, 39, 40,
	0, 37, 0, 0, 41, 0, 215, 38, 53, 54,
	55, 50, 0, 57, 59, 0, 0, 58, 0, 42,
	0, 34, 35, 51, 52, 32, 0, 31, 56, 0,
	0, 0, 0, 39, 40, 0, 37, 0, 0, 41,
	0, 38, 53, 54, 55, 0, 50, 0, 57, 59,
	0, 0, 58, 0, 42, 0, 34, 0, 0, 197,
	32, 0, 0, 56, 0, 0, 0, 39, 40, 0,
	37, 0, 0, 41, 0, 164, 0, 0, 0, 0,
	50, 0, 57, 59, 0, 0, 58, 0, 42, 0,
	34, 35, 51, 52, 32, 0, 31, 56, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	38, 53, 54, 55, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 35, 51, 52, 0, 0, 31,
	0, 0, 0, 0, 0, 0, 39, 40, 0, 37,
	0, 0, 41, 38, 53, 54, 55, 0, 0, 50,
	0, 57, 59, 0, 0, 58, 0, 42, 0, 34,
	0, 0, 0, 32, 0, 0, 56, 0, 0, 39,
	40, 0, 37, 0, 0, 41, 0, 0, 0, 0,
	0, 0, 50, 0, 57, 59, 0, 0, 58, 0,
	341, 0, 34, 35, 51, 52, 32, 0, 31, 56,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 38, 53, 54, 55, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 35, 51, 52, 0,
	0, 31, 0, 0, 0, 0, 0, 0, 39, 40,
	0, 37, 0, 0, 41, 38, 53, 54, 55, 0,
	0, 50, 0, 57, 59, 0, 0, 58, 0, 300,
	0, 34, 0, 0, 0, 32, 0, 0, 56, 0,
	0, 39, 40, 0, 37, 0, 0, 41, 0, 0,
	0, 0, 0, 0, 50, 0, 57, 59, 0, 0,
	58, 0, 298, 0, 34, 35, 51, 52, 32, 0,
	31, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 38, 53, 54, 55, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	39, 40, 0, 37, 0, 0, 41, 84, 103, 104,
	108, 106, 0, 50, 0, 57, 59, 0, 0, 58,
	0, 246, 0, 34, 0, 0, 0, 32, 0, 0,
	56, 86, 87, 97, 98, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	105, 107, 100, 101, 102, 0, 94, 95, 96, 99,
	35, 144, 52, 81, 0, 31, 0, 82, 0, 85,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 38,
	53, 54, 55, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 73, 51, 52, 0, 0, 31, 0,
	0, 0, 0, 0, 0, 39, 40, 0, 37, 0,
	0, 41, 38, 53, 54, 55, 0, 0, 50, 0,
	57, 59, 0, 0, 58, 0, 42, 0, 34, 0,
	0, 0, 32, 0, 0, 56, 0, 0, 39, 40,
	0, 37, 0, 0, 41, 0, 0, 0, 0, 0,
	0, 50, 0, 57, 59, 0, 0, 58, 0, 42,
	0, 34, 0, 0, 0, 32, 0, 0, 56,
}
var yyPact = [...]int{

	-58, -1000, 862, -58, -1000, -59, -59, -1000, -1000, -1000,
	-1000, -1000, -1000, 3137, 3137, 313, 212, 3479, 95, 94,
	282, -1000, -1000, 2438, -1000, -1000, 3137, 2715, 3137, -1000,
	-1000, 143, -44, 149, 3137, 74, -23, 91, 85, 82,
	3137, -13, -59, -1000, -1000, -1000, -1000, 297, 67, -1000,
	3446, -1000, -1000, -1000, -1000, -1000, 3137, 3137, 3137, 3137,
	-1000, -1000, -1000, -1000, -1000, 862, -59, -1000, 1, 2501,
	2501, 209, -58, 77, 2564, 3137, 3137, 207, 3137, 3137,
	3137, 3137, 3068, 3137, 3137, 305, -1000, -1000, 3137, 3137,
	3137, 3137, 3137, 3137, 3137, 3137, 3137, 3137, 3137, 3137,
	3137, 3137, 3137, 3137, 3137, 3137, 3137, 3137, 3137, 3137,
	3137, 2375, -58, 58, 787, 3034, -3, 74, 2312, 297,
	57, -21, 3137, -59, -16, -1000, 149, 149, -20, 149,
	-25, 2249, 3137, 2999, 3137, 149, 66, 2627, 149, 3137,
	56, -1000, 3137, -59, -1000, -11, -11, -11, -11, -11,
	-1000, -58, 187, 3137, 3137, 724, 2186, 3137, -58, 2501,
	2123, 2690, 198, 661, 3137, 2627, 104, -1000, 2501, 2501,
	2501, 2501, 2501, 2501, 104, 104, 104, 104, 104, 104,
	226, 226, 226, 2868, 2868, 2868, 2868, 2868, 2868, 3381,
	2934, -58, 185, -59, 3137, -59, -58, 3341, 2060, 2893,
	-59, 144, 297, -1000, -50, -59, 296, -53, -53, 149,
	-53, -21, -1000, 119, 598, 3137, 1997, 0, -28, 295,
	-34, -54, 1934, 3137, 1, 3137, 184, 260, 116, 114,
	-1000, 3137, -1000, 1871, 183, 3137, 70, -1000, -1000, 2824,
	1808, 177, -1000, 1745, 293, 176, -58, 1682, 3272, 3239,
	1619, 240, 208, 64, 105, -59, -45, -59, 3137, -1000,
	-40, 60, -1000, -1000, 2791, 1556, -1000, -1000, -1000, 3137,
	8, 149, 174, -59, 3137, 1, 2501, -23, -1000, 181,
	52, -1000, 45, -1000, 1493, -58, -1000, 2627, -1000, 535,
	-1000, -1000, -1000, -58, -1000, -1000, 172, -58, -58, 1430,
	-58, 1367, 3170, -32, -1000, -1000, 227, 3137, -58, 206,
	204, 13, -59, -1000, -50, 149, -1000, 472, -1000, -1000,
	409, 3137, -47, -1000, 3137, 2501, 203, -58, -1000, -1000,
	-1000, 170, -1000, 3137, 167, -1000, 164, 162, -58, 160,
	-58, -58, 1304, 159, -1000, -1000, -58, 1241, 14, 156,
	-58, -58, 191, 153, -53, -1000, 3137, -1000, 3137, 1178,
	-59, 1115, -58, 151, -1000, 1052, -1000, -1000, -1000, 150,
	-1000, 139, 138, -58, -1000, -1000, -58, -58, -1000, 134,
	128, -58, -1000, 989, 926, -1000, 3137, 3137, 120, 257,
	-1000, -1000, -1000, -1000, 110, -1000, -1000, -1000, -1000, 101,
	-1000, -1000, -54, 2501, 256, 189, -1000, -1000, 98, 182,
	-58, -1000, -58, 97, 68, -1000, -1000,
}
var yyPgo = [...]int{

	0, 42, 354, 283, 295, 353, 352, 347, 346, 345,
	344, 10, 8, 5, 0, 16, 47, 34, 342, 339,
	11, 338, 4, 337, 335, 333, 332, 331, 329, 327,
	318, 326, 320, 7, 1, 6, 2,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 4, 5,
	6, 6, 7, 7, 7, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 9, 10, 10, 10,
	10, 10, 11, 11, 12, 13, 13, 13, 13, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 15, 15, 15, 16, 16, 16, 16, 16,
	16, 17, 17, 18, 18, 19, 20, 21, 21, 21,
	21, 21, 21, 22, 22, 22, 23, 23, 23, 23,
	23, 23, 23, 23, 24, 24, 24, 24, 24, 25,
	25, 25, 25, 26, 26, 26, 26, 26, 26, 26,
	26, 30, 30, 30, 30, 30, 30, 29, 29, 29,
	28, 28, 28, 28, 28, 28, 27, 27, 31, 31,
	32, 32, 32, 33, 33, 35, 35, 36, 34, 34,
	34, 34,
}
var yyR2 = [...]int{

	0, 1, 2, 2, 3, 0, 1, 1, 1, 2,
	2, 5, 13, 12, 9, 8, 6, 5, 6, 5,
	4, 6, 4, 1, 1, 1, 1, 1, 1, 4,
	3, 3, 5, 7, 5, 4, 7, 5, 6, 7,
	7, 8, 7, 8, 8, 9, 7, 0, 1, 1,
	2, 2, 4, 4, 3, 0, 1, 4, 4, 1,
	1, 5, 3, 7, 8, 8, 9, 2, 5, 7,
	3, 5, 4, 5, 4, 4, 4, 4, 4, 4,
	6, 8, 7, 3, 2, 3, 10, 5, 1, 1,
	1, 1, 0, 1, 4, 1, 3, 2, 2, 5,
	2, 2, 3, 1, 1, 3, 1, 2, 1, 1,
	1, 1, 1, 0, 3, 6, 6, 5, 5, 6,
	5, 5, 8, 8, 2, 2, 2, 2, 2, 1,
	1, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 0, 1,
	2, 1, 1, 0, 1, 1, 2, 1, 0, 2,
	1, 1,
}
var yyChk = [...]int{

	-1000, -1, -31, -2, -32, 75, -35, -36, 80, -3,
	-4, 38, 39, 10, 12, 28, 29, 47, 53, 54,
	-7, -8, -9, -14, -5, -6, 13, 15, 44, -18,
	-21, 9, 76, -17, 72, 4, -20, 52, 23, 49,
	50, 55, 70, -23, -24, -25, -26, 11, -13, -19,
	62, 5, 6, 24, 25, 26, 79, 64, 68, 65,
	-30, -29, -28, -27, -31, -32, -35, -36, -13, -14,
	-14, 4, 70, 4, -14, 72, 72, 14, 56, 58,
	27, 72, 76, 50, 16, 78, 40, 41, 32, 33,
	37, 34, 35, 36, 65, 66, 67, 42, 43, 68,
	61, 62, 63, 17, 18, 59, 20, 60, 19, 22,
	21, -14, 70, -15, -14, 75, -4, 4, -14, 72,
	4, 77, -33, -35, -16, 4, 65, -17, 55, 48,
	76, -14, 72, 76, 72, 72, 72, -14, 76, -33,
	-15, 4, 56, 74, 5, -14, -14, -14, -14, -14,
	-3, 70, -1, 72, 72, -14, -14, 13, 70, -14,
	-14, -14, -13, -14, 57, -14, -14, 4, -14, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
	-14, 70, -1, -35, 16, 74, 70, 75, -14, 75,
	70, -15, 72, -17, -13, 70, 78, -16, -16, 76,
	-16, 77, 73, -13, -14, 57, -14, -16, -16, 51,
	-16, -22, -14, 56, -13, -33, -1, 71, -13, -13,
	73, 74, 73, -14, -1, 57, 8, 73, 77, 57,
	-14, -1, 71, -14, -33, -1, 70, -14, 75, 75,
	-14, -33, 73, 8, -15, 74, -34, -35, -33, 4,
	-16, 8, 73, 77, 57, -14, 73, 73, 73, 74,
	4, 77, -34, 74, 57, -13, -14, -20, 71, 30,
	8, 73, 8, 73, -14, 70, 71, -14, 73, -14,
	77, 77, 71, 70, 4, 71, -1, 70, 70, -14,
	70, -14, 75, -10, -12, -11, 46, 45, 70, 73,
	73, 8, -35, 77, -13, 77, 73, -14, 77, 77,
	-14, 74, -16, 71, -33, -14, 4, 70, 73, 73,
	73, -1, 77, 57, -1, 71, -1, -1, 70, -1,
	70, 70, -14, -33, -11, -12, 57, -14, -13, -1,
	70, 70, 73, -34, -16, 77, 57, 73, 74, -14,
	70, -14, 70, -1, 71, -14, 71, 71, 71, -1,
	71, -1, -1, 70, 71, -1, 57, 57, 71, -1,
	-1, 70, 71, -14, -14, 73, -33, 57, -1, 71,
	77, 71, 71, 71, -1, -1, -1, 71, 71, -1,
	77, 73, -22, -14, 71, 31, 71, 71, -34, 31,
	70, 71, 70, -1, -1, 71, 71,
}
var yyDef = [...]int{

	158, -2, -2, 158, 159, 162, 161, 165, 167, 3,
	6, 7, 8, 55, 0, 0, 0, 0, 0, 0,
	23, 24, 25, -2, 27, 28, 0, -2, 0, 59,
	60, 0, 163, 0, 0, 106, 104, 0, 0, 0,
	0, 0, 163, 88, 89, 90, 91, 92, 0, 103,
	0, 108, 109, 110, 111, 112, 0, 0, 0, 0,
	129, 130, 131, 132, 2, -2, 160, 166, 9, 56,
	10, 0, 158, 106, 0, 0, 0, 0, 0, 0,
	0, 55, 0, 0, 0, 0, 133, 134, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 158, 0, 56, 0, 0, -2, 0, 92,
	0, -2, 55, 164, 0, 95, 0, 0, 0, 0,
	0, 0, 55, 0, 0, 0, 0, 84, 0, 113,
	0, 93, 55, 163, 107, 124, 125, 126, 127, 128,
	4, 158, 0, 55, 55, 0, 0, 0, 158, 30,
	0, 62, 0, 0, 0, 83, 85, 105, 135, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 150, 151, 152, 153, 154, 155, 156,
	157, 158, 0, 161, 0, 163, 158, 0, 0, 0,
	163, 0, 92, 102, 168, 163, 0, 97, 98, 0,
	100, 101, 70, 0, 0, 0, 0, 0, 0, 0,
	0, 168, 0, 55, 31, 0, 0, 0, 0, 0,
	20, 0, 22, 0, 0, 0, 0, 74, 76, 0,
	0, 0, 35, 0, 0, 0, 158, 0, 0, 0,
	0, 47, 0, 0, 0, -2, 0, 170, 55, 96,
	0, 0, 72, 75, 0, 0, 77, 78, 79, 0,
	0, 0, 0, -2, 0, 29, 57, -2, 11, 0,
	0, -2, 0, -2, 0, 158, 34, 61, 73, 0,
	120, 121, 32, 158, 94, 37, 0, 158, 158, 0,
	158, 0, 0, 163, 48, 49, 0, 55, 158, 0,
	0, 0, -2, 68, 168, 0, 71, 0, 117, 118,
	0, 0, 0, 87, 0, 114, 0, 158, -2, -2,
	21, 0, 119, 0, 0, 38, 0, 0, 158, 0,
	158, 158, 0, 0, 50, 51, 158, 56, 0, 0,
	158, 158, 0, 0, 99, 116, 0, 80, 0, 0,
	163, 0, 158, 0, 33, 0, 36, 39, 40, 0,
	42, 0, 0, 158, 46, 54, 158, 158, 63, 0,
	0, 158, 69, 0, 0, 82, 113, 0, 0, 15,
	123, 41, 43, 44, 0, 52, 53, 64, 65, 0,
	122, 81, 168, 115, 14, 0, 45, 66, 0, 0,
	158, 86, 158, 0, 0, 13, 12,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	80, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 79, 3, 3, 3, 67, 68, 3,
	72, 73, 65, 61, 74, 62, 78, 66, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 57, 75,
	59, 56, 60, 58, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 76, 3, 77, 64, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 70, 63, 71,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 69,
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
		//line parser.go.y:104
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:108
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:114
		{
			if yyDollar[2].stmt != nil {
				yyVAL.stmts = &ast.StmtsStmt{Stmts: []ast.Stmt{yyDollar[2].stmt}}
			}
			if l, ok := yylex.(*Lexer); ok {
				l.stmt = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:123
		{
			if yyDollar[3].stmt != nil {
				if yyDollar[1].stmts == nil {
					yyVAL.stmts = &ast.StmtsStmt{Stmts: []ast.Stmt{yyDollar[3].stmt}}
				} else {
					stmts := yyDollar[1].stmts.(*ast.StmtsStmt)
					stmts.Stmts = append(stmts.Stmts, yyDollar[3].stmt)
				}
				if l, ok := yylex.(*Lexer); ok {
					l.stmt = yyVAL.stmts
				}
			}
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.stmt = nil
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.stmt = yyDollar[1].stmt_var_or_lets
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:157
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmt: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:172
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:177
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:182
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 15:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:192
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:197
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:202
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].expr.Position())
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:207
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:212
		{
			yyVAL.stmt = &ast.DeleteStmt{Item: yyDollar[3].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:217
		{
			yyVAL.stmt = &ast.DeleteStmt{Item: yyDollar[3].expr, Key: yyDollar[5].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:222
		{
			yyVAL.stmt = &ast.CloseStmt{Expr: yyDollar[3].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:227
		{
			yyVAL.stmt = yyDollar[1].stmt_if
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:231
		{
			yyVAL.stmt = yyDollar[1].stmt_for
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:235
		{
			yyVAL.stmt = yyDollar[1].stmt_switch
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:246
		{
			yyVAL.stmt_var_or_lets = yyDollar[1].stmt_var
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:250
		{
			yyVAL.stmt_var_or_lets = yyDollar[1].stmt_lets
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].exprs}
			yyVAL.stmt_var.SetPosition(yyDollar[1].tok.Position())
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:263
		{
			yyVAL.stmt_lets = &ast.LetsStmt{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{yyDollar[3].expr}}
			yyVAL.stmt_lets.SetPosition(yyDollar[1].expr.Position())
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:268
		{
			if len(yyDollar[1].exprs) == 2 && len(yyDollar[3].exprs) == 1 {
				if _, ok := yyDollar[3].exprs[0].(*ast.ItemExpr); ok {
					yyVAL.stmt_lets = &ast.LetMapItemStmt{LHSS: yyDollar[1].exprs, RHS: yyDollar[3].exprs[0]}
				} else {
					yyVAL.stmt_lets = &ast.LetsStmt{LHSS: yyDollar[1].exprs, RHSS: yyDollar[3].exprs}
				}
			} else {
				yyVAL.stmt_lets = &ast.LetsStmt{LHSS: yyDollar[1].exprs, RHSS: yyDollar[3].exprs}
			}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:282
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 33:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:287
		{
			ifStmt := yyDollar[1].stmt_if.(*ast.IfStmt)
			ifStmt.ElseIf = append(ifStmt.ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:292
		{
			ifStmt := yyDollar[1].stmt_if.(*ast.IfStmt)
			if ifStmt.Else != nil {
				yylex.Error("multiple else statement")
			}
			ifStmt.Else = yyDollar[4].compstmt
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmt: yyDollar[3].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 36:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:307
		{
			if len(yyDollar[2].expr_idents) < 1 {
				yylex.Error("missing identifier")
			} else if len(yyDollar[2].expr_idents) > 2 {
				yylex.Error("too many identifiers")
			} else {
				yyVAL.stmt_for = &ast.ForStmt{Vars: yyDollar[2].expr_idents, Value: yyDollar[4].expr, Stmt: yyDollar[6].compstmt}
				yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
			}
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:318
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmt: yyDollar[4].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:323
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt: yyDollar[5].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:328
		{
			yyVAL.stmt_for = &ast.CForStmt{Expr3: yyDollar[4].expr, Stmt: yyDollar[6].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:333
		{
			yyVAL.stmt_for = &ast.CForStmt{Expr2: yyDollar[3].expr, Stmt: yyDollar[6].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 41:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:338
		{
			yyVAL.stmt_for = &ast.CForStmt{Expr2: yyDollar[3].expr, Expr3: yyDollar[5].expr, Stmt: yyDollar[7].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 42:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:343
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt1: yyDollar[2].stmt_var_or_lets, Stmt: yyDollar[6].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 43:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:348
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt1: yyDollar[2].stmt_var_or_lets, Expr3: yyDollar[5].expr, Stmt: yyDollar[7].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 44:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:353
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt1: yyDollar[2].stmt_var_or_lets, Expr2: yyDollar[4].expr, Stmt: yyDollar[7].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 45:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:358
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt1: yyDollar[2].stmt_var_or_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmt: yyDollar[8].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 46:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:365
		{
			switchStmt := yyDollar[5].stmt_switch_cases.(*ast.SwitchStmt)
			switchStmt.Expr = yyDollar[2].expr
			yyVAL.stmt_switch = switchStmt
			yyVAL.stmt_switch.SetPosition(yyDollar[1].tok.Position())
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:374
		{
			yyVAL.stmt_switch_cases = &ast.SwitchStmt{}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:378
		{
			yyVAL.stmt_switch_cases = &ast.SwitchStmt{Default: yyDollar[1].stmt_switch_default}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:382
		{
			yyVAL.stmt_switch_cases = &ast.SwitchStmt{Cases: []ast.Stmt{yyDollar[1].stmt_switch_case}}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:386
		{
			switchStmt := yyDollar[1].stmt_switch_cases.(*ast.SwitchStmt)
			switchStmt.Cases = append(switchStmt.Cases, yyDollar[2].stmt_switch_case)
			yyVAL.stmt_switch_cases = switchStmt
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:392
		{
			switchStmt := yyDollar[1].stmt_switch_cases.(*ast.SwitchStmt)
			if switchStmt.Default != nil {
				yylex.Error("multiple default statement")
			}
			switchStmt.Default = yyDollar[2].stmt_switch_default
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:402
		{
			yyVAL.stmt_switch_case = &ast.SwitchCaseStmt{Exprs: []ast.Expr{yyDollar[2].expr}, Stmt: yyDollar[4].compstmt}
			yyVAL.stmt_switch_case.SetPosition(yyDollar[1].tok.Position())
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:407
		{
			yyVAL.stmt_switch_case = &ast.SwitchCaseStmt{Exprs: yyDollar[2].exprs, Stmt: yyDollar[4].compstmt}
			yyVAL.stmt_switch_case.SetPosition(yyDollar[1].tok.Position())
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:414
		{
			yyVAL.stmt_switch_default = yyDollar[3].compstmt
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:421
		{
			yyVAL.exprs = nil
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:425
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:429
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:436
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr_ident)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:445
		{
			yyVAL.expr = yyDollar[1].expr_member_or_ident
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:449
		{
			yyVAL.expr = yyDollar[1].expr_literals
		}
	case 61:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:453
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, LHS: yyDollar[3].expr, RHS: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:458
		{
			yyVAL.expr = &ast.NilCoalescingOpExpr{LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 63:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:463
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmt: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:468
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmt: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:473
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmt: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:478
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmt: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:483
		{
			yyVAL.expr = &ast.ArrayExpr{}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:488
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 69:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:493
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[5].exprs, TypeData: &ast.TypeStruct{Kind: ast.TypeSlice, SubType: yyDollar[2].type_data, Dimensions: yyDollar[1].slice_count}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:498
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:503
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:508
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:513
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:518
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:523
		{
			yyVAL.expr = &ast.ItemExpr{Item: yyDollar[1].expr_ident, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr_ident.Position())
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:528
		{
			yyVAL.expr = &ast.ItemExpr{Item: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:533
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:538
		{
			if yyDollar[3].type_data.Kind == ast.TypeDefault {
				yyDollar[3].type_data.Kind = ast.TypePtr
				yyVAL.expr = &ast.MakeExpr{TypeData: yyDollar[3].type_data}
			} else {
				yyVAL.expr = &ast.MakeExpr{TypeData: &ast.TypeStruct{Kind: ast.TypePtr, SubType: yyDollar[3].type_data}}
			}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:548
		{
			yyVAL.expr = &ast.MakeExpr{TypeData: yyDollar[3].type_data}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 80:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:553
		{
			yyVAL.expr = &ast.MakeExpr{TypeData: yyDollar[3].type_data, LenExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 81:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:558
		{
			yyVAL.expr = &ast.MakeExpr{TypeData: yyDollar[3].type_data, LenExpr: yyDollar[5].expr, CapExpr: yyDollar[7].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 82:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:563
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].tok.Lit, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:568
		{
			yyVAL.expr = &ast.ChanExpr{LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:573
		{
			yyVAL.expr = &ast.ChanExpr{RHS: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:578
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.go.y:583
		{
			yyDollar[8].expr_map.TypeData = &ast.TypeStruct{Kind: ast.TypeMap, Key: yyDollar[3].type_data, SubType: yyDollar[5].type_data}
			yyVAL.expr = yyDollar[8].expr_map
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:589
		{
			yyVAL.expr = yyDollar[3].expr_map
			yyVAL.expr.SetPosition(yyDollar[3].expr_map.Position())
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:594
		{
			yyVAL.expr = yyDollar[1].expr_slice
			yyVAL.expr.SetPosition(yyDollar[1].expr_slice.Position())
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:603
		{
			yyVAL.expr_idents = []string{}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:607
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:611
		{
			if len(yyDollar[1].expr_idents) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:620
		{
			yyVAL.type_data = &ast.TypeStruct{Name: yyDollar[1].tok.Lit}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:624
		{
			if yyDollar[1].type_data.Kind != ast.TypeDefault {
				yylex.Error("blah1")
			} else {
				yyDollar[1].type_data.Env = append(yyDollar[1].type_data.Env, yyDollar[1].type_data.Name)
				yyDollar[1].type_data.Name = yyDollar[3].tok.Lit
			}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:633
		{
			if yyDollar[2].type_data.Kind == ast.TypeDefault {
				yyDollar[2].type_data.Kind = ast.TypePtr
				yyVAL.type_data = yyDollar[2].type_data
			} else {
				yyVAL.type_data = &ast.TypeStruct{Kind: ast.TypePtr, SubType: yyDollar[2].type_data}
			}
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:642
		{
			if yyDollar[2].type_data.Kind == ast.TypeDefault {
				yyDollar[2].type_data.Kind = ast.TypeSlice
				yyDollar[2].type_data.Dimensions = yyDollar[1].slice_count
				yyVAL.type_data = yyDollar[2].type_data
			} else {
				yyVAL.type_data = &ast.TypeStruct{Kind: ast.TypeSlice, SubType: yyDollar[2].type_data, Dimensions: yyDollar[1].slice_count}
			}
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:652
		{
			yyVAL.type_data = &ast.TypeStruct{Kind: ast.TypeMap, Key: yyDollar[3].type_data, SubType: yyDollar[5].type_data}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:656
		{
			if yyDollar[2].type_data.Kind == ast.TypeDefault {
				yyDollar[2].type_data.Kind = ast.TypeChan
				yyVAL.type_data = yyDollar[2].type_data
			} else {
				yyVAL.type_data = &ast.TypeStruct{Kind: ast.TypeChan, SubType: yyDollar[2].type_data}
			}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:668
		{
			yyVAL.slice_count = 1
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:672
		{
			yyVAL.slice_count = yyDollar[3].slice_count + 1
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:678
		{
			yyVAL.expr_member_or_ident = yyDollar[1].expr_member
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:682
		{
			yyVAL.expr_member_or_ident = yyDollar[1].expr_ident
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:688
		{
			yyVAL.expr_member = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr_member.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:695
		{
			yyVAL.expr_ident = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr_ident.SetPosition(yyDollar[1].tok.Position())
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:702
		{
			num, err := toNumber("-" + yyDollar[2].tok.Lit)
			if err != nil {
				yylex.Error("invalid number: -" + yyDollar[2].tok.Lit)
			}
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: num}
			yyVAL.expr_literals.SetPosition(yyDollar[2].tok.Position())
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:711
		{
			num, err := toNumber(yyDollar[1].tok.Lit)
			if err != nil {
				yylex.Error("invalid number: " + yyDollar[1].tok.Lit)
			}
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: num}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:720
		{
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: stringToValue(yyDollar[1].tok.Lit)}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:725
		{
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: trueValue}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:730
		{
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: falseValue}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:735
		{
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: nilValue}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:742
		{
			yyVAL.expr_map = &ast.MapExpr{}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:746
		{
			yyVAL.expr_map = &ast.MapExpr{Keys: []ast.Expr{yyDollar[1].expr}, Values: []ast.Expr{yyDollar[3].expr}}
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:750
		{
			if yyDollar[1].expr_map.Keys == nil {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_map.Keys = append(yyVAL.expr_map.Keys, yyDollar[4].expr)
			yyVAL.expr_map.Values = append(yyVAL.expr_map.Values, yyDollar[6].expr)
		}
	case 116:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:760
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr_ident, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:764
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr_ident, Begin: yyDollar[3].expr, End: nil}
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:768
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr_ident, Begin: nil, End: yyDollar[4].expr}
		}
	case 119:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:772
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:776
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:780
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
		}
	case 122:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:784
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr_ident, Begin: yyDollar[3].expr, End: yyDollar[5].expr, Cap: yyDollar[7].expr}
		}
	case 123:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:788
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr, Cap: yyDollar[7].expr}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:794
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:799
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:804
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:809
		{
			yyVAL.expr = &ast.AddrExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:814
		{
			yyVAL.expr = &ast.DerefExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:821
		{
			yyVAL.expr = &ast.OpExpr{Op: yyDollar[1].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:826
		{
			yyVAL.expr = &ast.OpExpr{Op: yyDollar[1].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:831
		{
			yyVAL.expr = &ast.OpExpr{Op: yyDollar[1].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:836
		{
			yyVAL.expr = &ast.OpExpr{Op: yyDollar[1].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:843
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "+", RHS: oneLiteral}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:851
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "-", RHS: oneLiteral}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:859
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "+", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:867
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "-", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:875
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "|", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:883
		{
			rhs := &ast.OpExpr{Op: &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "*", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:891
		{
			rhs := &ast.OpExpr{Op: &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "/", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:899
		{
			rhs := &ast.OpExpr{Op: &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "&", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:910
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "*", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:915
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "/", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:920
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "%", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:925
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "<<", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:930
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: ">>", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:935
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "&", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:942
		{
			yyVAL.expr = &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "+", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:947
		{
			yyVAL.expr = &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "-", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:952
		{
			yyVAL.expr = &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "|", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:959
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: "==", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:964
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: "!=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:969
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: "<", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:974
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: "<=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:979
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: ">", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:984
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: ">=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:991
		{
			yyVAL.expr = &ast.BinaryOperator{LHS: yyDollar[1].expr, Operator: "&&", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:996
		{
			yyVAL.expr = &ast.BinaryOperator{LHS: yyDollar[1].expr, Operator: "||", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	}
	goto yystack /* stack new state and value */
}
