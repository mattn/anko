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
const DEFER = 57390
const CHAN = 57391
const MAKE = 57392
const OPCHAN = 57393
const TYPE = 57394
const LEN = 57395
const DELETE = 57396
const CLOSE = 57397
const MAP = 57398
const UNARY = 57399

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
	"DEFER",
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

//line parser.go.y:1055

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	57, 59,
	75, 59,
	76, 5,
	-2, 1,
	-1, 24,
	75, 60,
	-2, 30,
	-1, 28,
	16, 96,
	-2, 59,
	-1, 66,
	57, 59,
	75, 59,
	-2, 5,
	-1, 120,
	16, 97,
	75, 97,
	-2, 110,
	-1, 124,
	4, 105,
	49, 105,
	56, 105,
	-2, 71,
	-1, 262,
	72, 177,
	78, 177,
	-2, 169,
	-1, 280,
	72, 177,
	-2, 169,
	-1, 284,
	1, 62,
	8, 62,
	45, 62,
	46, 62,
	57, 62,
	58, 62,
	72, 62,
	74, 62,
	75, 62,
	76, 62,
	78, 62,
	81, 62,
	-2, 108,
	-1, 288,
	1, 17,
	45, 17,
	46, 17,
	72, 17,
	76, 17,
	81, 17,
	-2, 76,
	-1, 290,
	1, 19,
	45, 19,
	46, 19,
	72, 19,
	76, 19,
	81, 19,
	-2, 78,
	-1, 292,
	1, 21,
	45, 21,
	46, 21,
	72, 21,
	76, 21,
	81, 21,
	-2, 76,
	-1, 294,
	1, 23,
	45, 23,
	46, 23,
	72, 23,
	76, 23,
	81, 23,
	-2, 78,
	-1, 324,
	72, 175,
	78, 175,
	-2, 170,
	-1, 341,
	1, 16,
	45, 16,
	46, 16,
	72, 16,
	76, 16,
	81, 16,
	-2, 75,
	-1, 342,
	1, 18,
	45, 18,
	46, 18,
	72, 18,
	76, 18,
	81, 18,
	-2, 77,
	-1, 343,
	1, 20,
	45, 20,
	46, 20,
	72, 20,
	76, 20,
	81, 20,
	-2, 75,
	-1, 344,
	1, 22,
	45, 22,
	46, 22,
	72, 22,
	76, 22,
	81, 22,
	-2, 77,
}

const yyPrivate = 57344

const yyLast = 3828

var yyAct = [...]int{

	70, 263, 316, 24, 37, 142, 226, 317, 8, 34,
	211, 5, 325, 87, 127, 71, 8, 280, 75, 77,
	319, 318, 216, 8, 124, 262, 377, 8, 114, 117,
	121, 8, 49, 136, 211, 274, 134, 89, 90, 125,
	211, 116, 140, 133, 130, 214, 69, 7, 327, 211,
	275, 276, 148, 1, 68, 211, 8, 210, 149, 150,
	151, 152, 84, 278, 211, 211, 85, 24, 88, 141,
	84, 204, 395, 128, 85, 146, 88, 199, 334, 160,
	161, 368, 164, 165, 166, 344, 168, 170, 171, 146,
	143, 343, 173, 174, 175, 176, 177, 178, 179, 180,
	181, 182, 183, 184, 185, 186, 187, 188, 189, 190,
	191, 192, 193, 194, 195, 68, 6, 167, 132, 203,
	128, 224, 67, 228, 342, 131, 341, 155, 323, 328,
	293, 321, 291, 299, 208, 129, 200, 219, 221, 130,
	130, 200, 130, 227, 212, 213, 133, 215, 130, 130,
	126, 130, 230, 222, 223, 145, 225, 207, 209, 135,
	126, 123, 158, 240, 206, 132, 156, 289, 218, 197,
	247, 287, 131, 146, 68, 268, 139, 138, 229, 137,
	79, 78, 129, 435, 434, 260, 430, 243, 431, 233,
	234, 235, 236, 133, 322, 200, 294, 146, 292, 146,
	250, 426, 425, 254, 423, 257, 251, 417, 231, 416,
	412, 258, 411, 410, 408, 400, 265, 241, 396, 392,
	388, 272, 386, 385, 130, 384, 208, 381, 279, 267,
	122, 283, 198, 290, 146, 284, 351, 288, 146, 295,
	336, 269, 146, 298, 339, 307, 68, 300, 304, 261,
	248, 259, 200, 244, 146, 252, 311, 313, 297, 285,
	249, 282, 232, 126, 429, 399, 162, 379, 230, 367,
	366, 320, 329, 154, 73, 362, 428, 333, 9, 319,
	318, 424, 338, 80, 286, 10, 337, 306, 130, 277,
	266, 144, 172, 335, 72, 4, 2, 61, 326, 66,
	65, 62, 63, 64, 349, 47, 46, 308, 45, 44,
	31, 340, 68, 50, 119, 358, 30, 126, 361, 315,
	363, 359, 126, 360, 163, 23, 264, 126, 369, 22,
	21, 26, 25, 373, 3, 376, 0, 130, 378, 0,
	0, 0, 370, 264, 0, 153, 0, 0, 0, 382,
	346, 0, 364, 0, 0, 0, 0, 0, 0, 350,
	0, 0, 0, 352, 353, 0, 355, 0, 0, 0,
	198, 0, 68, 401, 365, 0, 403, 0, 0, 324,
	0, 0, 0, 405, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 380, 0, 0, 324, 0, 0,
	0, 0, 0, 0, 0, 0, 227, 422, 387, 0,
	389, 390, 421, 0, 0, 0, 393, 0, 0, 0,
	397, 398, 0, 427, 0, 0, 0, 198, 0, 198,
	0, 0, 126, 407, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 264, 0, 413, 0, 0, 414, 415,
	0, 0, 0, 418, 0, 0, 36, 52, 53, 0,
	0, 32, 13, 48, 14, 27, 0, 28, 0, 0,
	0, 0, 0, 0, 198, 39, 54, 55, 56, 0,
	15, 16, 0, 432, 0, 433, 0, 0, 0, 0,
	11, 12, 0, 0, 126, 0, 29, 0, 0, 17,
	18, 0, 40, 41, 0, 38, 19, 20, 42, 0,
	0, 0, 0, 0, 0, 51, 0, 58, 60, 0,
	0, 59, 0, 43, 0, 35, 0, 0, 0, 33,
	0, 0, 57, 0, 0, 0, 0, 0, 264, 87,
	106, 107, 111, 109, 113, 112, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 91, 92, 94, 95, 96,
	93, 0, 0, 89, 90, 100, 101, 0, 0, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 0, 0,
	0, 0, 82, 108, 110, 103, 104, 105, 0, 97,
	98, 99, 102, 0, 0, 0, 84, 374, 375, 0,
	85, 0, 88, 87, 106, 107, 111, 109, 113, 112,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 91,
	92, 94, 95, 96, 93, 0, 0, 89, 90, 100,
	101, 0, 0, 0, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 0, 372, 82, 108, 110, 103,
	104, 105, 0, 97, 98, 99, 102, 0, 0, 0,
	84, 0, 0, 0, 85, 371, 88, 87, 106, 107,
	111, 109, 113, 112, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 91, 92, 94, 95, 96, 93, 0,
	0, 89, 90, 100, 101, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 0, 0, 0, 0, 0, 348,
	82, 108, 110, 103, 104, 105, 0, 97, 98, 99,
	102, 0, 0, 0, 84, 0, 0, 0, 85, 347,
	88, 87, 106, 107, 111, 109, 113, 112, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 91, 92, 94,
	95, 96, 93, 0, 0, 89, 90, 100, 101, 0,
	0, 0, 0, 0, 0, 0, 86, 0, 0, 0,
	0, 0, 0, 332, 82, 108, 110, 103, 104, 105,
	0, 97, 98, 99, 102, 0, 0, 0, 84, 0,
	0, 0, 85, 331, 88, 87, 106, 107, 111, 109,
	113, 112, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 91, 92, 94, 95, 96, 93, 0, 0, 89,
	90, 100, 101, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 0, 0, 0, 0, 0, 303, 82, 108,
	110, 103, 104, 105, 0, 97, 98, 99, 102, 0,
	0, 0, 84, 0, 0, 0, 85, 302, 88, 87,
	106, 107, 111, 109, 113, 112, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 91, 92, 94, 95, 96,
	93, 0, 0, 89, 90, 100, 101, 0, 0, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 0, 0,
	0, 271, 82, 108, 110, 103, 104, 105, 0, 97,
	98, 99, 102, 0, 0, 0, 84, 0, 0, 0,
	85, 270, 88, 87, 106, 107, 111, 109, 113, 112,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 91,
	92, 94, 95, 96, 93, 0, 0, 89, 90, 100,
	101, 0, 0, 0, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 0, 246, 82, 108, 110, 103,
	104, 105, 0, 97, 98, 99, 102, 0, 0, 0,
	84, 0, 0, 0, 85, 245, 88, 87, 106, 107,
	111, 109, 113, 112, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 91, 92, 94, 95, 96, 93, 0,
	0, 89, 90, 100, 101, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 0, 0, 0, 0, 0, 0,
	82, 108, 110, 103, 104, 105, 0, 97, 98, 99,
	102, 0, 0, 0, 84, 237, 238, 0, 85, 0,
	88, 87, 106, 107, 111, 109, 113, 112, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 91, 92, 94,
	95, 96, 93, 0, 0, 89, 90, 100, 101, 0,
	0, 0, 0, 0, 0, 0, 86, 0, 0, 0,
	0, 0, 81, 0, 82, 108, 110, 103, 104, 105,
	0, 97, 98, 99, 102, 0, 201, 0, 84, 0,
	0, 0, 85, 0, 88, 87, 106, 107, 111, 109,
	113, 112, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 91, 92, 94, 95, 96, 93, 0, 0, 89,
	90, 100, 101, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 0, 0, 0, 0, 0, 0, 82, 108,
	110, 103, 104, 105, 0, 97, 98, 99, 102, 0,
	0, 0, 84, 420, 0, 0, 85, 0, 88, 87,
	106, 107, 111, 109, 113, 112, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 91, 92, 94, 95, 96,
	93, 0, 0, 89, 90, 100, 101, 0, 0, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 0, 0,
	0, 0, 82, 108, 110, 103, 104, 105, 0, 97,
	98, 99, 102, 0, 0, 0, 84, 0, 0, 0,
	85, 419, 88, 87, 106, 107, 111, 109, 113, 112,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 91,
	92, 94, 95, 96, 93, 0, 0, 89, 90, 100,
	101, 0, 0, 0, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 0, 0, 82, 108, 110, 103,
	104, 105, 0, 97, 98, 99, 102, 0, 0, 0,
	84, 0, 0, 0, 85, 409, 88, 87, 106, 107,
	111, 109, 113, 112, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 91, 92, 94, 95, 96, 93, 0,
	0, 89, 90, 100, 101, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 0, 0, 0, 0, 0, 406,
	82, 108, 110, 103, 104, 105, 0, 97, 98, 99,
	102, 0, 0, 0, 84, 0, 0, 0, 85, 0,
	88, 87, 106, 107, 111, 109, 113, 112, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 91, 92, 94,
	95, 96, 93, 0, 0, 89, 90, 100, 101, 0,
	0, 0, 0, 0, 0, 0, 86, 0, 0, 0,
	0, 0, 0, 0, 82, 108, 110, 103, 104, 105,
	0, 97, 98, 99, 102, 0, 0, 0, 84, 404,
	0, 0, 85, 0, 88, 87, 106, 107, 111, 109,
	113, 112, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 91, 92, 94, 95, 96, 93, 0, 0, 89,
	90, 100, 101, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 0, 0, 0, 0, 0, 0, 82, 108,
	110, 103, 104, 105, 0, 97, 98, 99, 102, 0,
	0, 0, 84, 0, 0, 0, 85, 402, 88, 87,
	106, 107, 111, 109, 113, 112, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 91, 92, 94, 95, 96,
	93, 0, 0, 89, 90, 100, 101, 0, 0, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 0, 0,
	0, 394, 82, 108, 110, 103, 104, 105, 0, 97,
	98, 99, 102, 0, 0, 0, 84, 0, 0, 0,
	85, 0, 88, 87, 106, 107, 111, 109, 113, 112,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 91,
	92, 94, 95, 96, 93, 0, 0, 89, 90, 100,
	101, 0, 0, 0, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 0, 0, 82, 108, 110, 103,
	104, 105, 0, 97, 98, 99, 102, 0, 391, 0,
	84, 0, 0, 0, 85, 0, 88, 87, 106, 107,
	111, 109, 113, 112, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 91, 92, 94, 95, 96, 93, 0,
	0, 89, 90, 100, 101, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 0, 0, 0, 0, 0, 0,
	82, 108, 110, 103, 104, 105, 0, 97, 98, 99,
	102, 0, 0, 0, 84, 0, 0, 0, 85, 383,
	88, 87, 106, 107, 111, 109, 113, 112, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 91, 92, 94,
	95, 96, 93, 0, 0, 89, 90, 100, 101, 0,
	0, 0, 0, 0, 0, 0, 86, 0, 0, 0,
	0, 0, 0, 0, 82, 108, 110, 103, 104, 105,
	0, 97, 98, 99, 102, 0, 356, 0, 84, 0,
	0, 0, 85, 0, 88, 87, 106, 107, 111, 109,
	113, 112, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 91, 92, 94, 95, 96, 93, 0, 0, 89,
	90, 100, 101, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 0, 0, 0, 0, 0, 0, 82, 108,
	110, 103, 104, 105, 0, 97, 98, 99, 102, 0,
	354, 0, 84, 0, 0, 0, 85, 0, 88, 87,
	106, 107, 111, 109, 113, 112, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 91, 92, 94, 95, 96,
	93, 0, 0, 89, 90, 100, 101, 0, 0, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 0, 0,
	0, 0, 82, 108, 110, 103, 104, 105, 0, 97,
	98, 99, 102, 0, 0, 0, 84, 345, 0, 0,
	85, 0, 88, 87, 106, 107, 111, 109, 113, 112,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 91,
	92, 94, 95, 96, 93, 0, 0, 89, 90, 100,
	101, 0, 0, 0, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 0, 0, 82, 108, 110, 103,
	104, 105, 0, 97, 98, 99, 102, 0, 0, 0,
	84, 0, 0, 314, 85, 0, 88, 87, 106, 107,
	111, 109, 113, 112, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 91, 92, 94, 95, 96, 93, 0,
	0, 89, 90, 100, 101, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 0, 0, 0, 0, 0, 0,
	82, 108, 110, 103, 104, 105, 0, 97, 98, 99,
	102, 0, 309, 0, 84, 0, 0, 0, 85, 0,
	88, 87, 106, 107, 111, 109, 113, 112, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 91, 92, 94,
	95, 96, 93, 0, 0, 89, 90, 100, 101, 0,
	0, 0, 0, 0, 0, 0, 86, 0, 0, 0,
	0, 0, 0, 0, 82, 108, 110, 103, 104, 105,
	0, 97, 98, 99, 102, 0, 305, 0, 84, 0,
	0, 0, 85, 0, 88, 87, 106, 107, 111, 109,
	113, 112, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 91, 92, 94, 95, 96, 93, 0, 0, 89,
	90, 100, 101, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 0, 0, 0, 0, 0, 0, 82, 108,
	110, 103, 104, 105, 0, 97, 98, 99, 102, 0,
	296, 0, 84, 0, 0, 0, 85, 0, 88, 87,
	106, 107, 111, 109, 113, 112, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 91, 92, 94, 95, 96,
	93, 0, 0, 89, 90, 100, 101, 0, 0, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 0, 0,
	0, 281, 82, 108, 110, 103, 104, 105, 0, 97,
	98, 99, 102, 0, 0, 0, 84, 0, 0, 0,
	85, 0, 88, 87, 106, 107, 111, 109, 113, 112,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 91,
	92, 94, 95, 96, 93, 0, 0, 89, 90, 100,
	101, 0, 0, 0, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 0, 0, 82, 108, 110, 103,
	104, 105, 0, 97, 98, 99, 102, 0, 0, 0,
	84, 273, 0, 0, 85, 0, 88, 87, 106, 107,
	111, 109, 113, 112, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 91, 92, 94, 95, 96, 93, 0,
	0, 89, 90, 100, 101, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 0, 0, 0, 0, 0, 0,
	82, 108, 110, 103, 104, 105, 0, 97, 98, 99,
	102, 0, 0, 0, 84, 0, 0, 255, 85, 0,
	88, 87, 106, 107, 111, 109, 113, 112, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 91, 92, 94,
	95, 96, 93, 0, 0, 89, 90, 100, 101, 0,
	0, 0, 0, 0, 0, 0, 86, 0, 0, 0,
	0, 0, 0, 242, 82, 108, 110, 103, 104, 105,
	0, 97, 98, 99, 102, 0, 0, 0, 84, 0,
	0, 0, 85, 0, 88, 87, 106, 107, 111, 109,
	113, 112, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 91, 92, 94, 95, 96, 93, 0, 0, 89,
	90, 100, 101, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 0, 0, 0, 0, 0, 0, 82, 108,
	110, 103, 104, 105, 0, 97, 98, 99, 102, 0,
	0, 0, 84, 239, 0, 0, 85, 0, 88, 87,
	106, 107, 111, 109, 113, 112, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 91, 92, 94, 95, 96,
	93, 0, 0, 89, 90, 100, 101, 0, 0, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 0, 0,
	0, 0, 82, 108, 110, 103, 104, 105, 0, 97,
	98, 99, 102, 0, 0, 0, 84, 217, 0, 0,
	85, 0, 88, 87, 106, 107, 111, 109, 113, 112,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 91,
	92, 94, 95, 96, 93, 0, 0, 89, 90, 100,
	101, 0, 0, 0, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 0, 0, 82, 108, 110, 103,
	104, 105, 0, 97, 98, 99, 102, 0, 205, 0,
	84, 0, 0, 0, 85, 0, 88, 87, 106, 107,
	111, 109, 113, 112, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 91, 92, 94, 95, 96, 93, 0,
	0, 89, 90, 100, 101, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 0, 0, 0, 0, 0, 0,
	82, 108, 110, 103, 104, 105, 0, 97, 98, 99,
	102, 0, 196, 0, 84, 0, 0, 0, 85, 0,
	88, 87, 106, 107, 111, 109, 113, 112, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 91, 92, 94,
	95, 96, 93, 0, 0, 89, 90, 100, 101, 0,
	0, 0, 0, 0, 0, 0, 86, 0, 0, 0,
	0, 0, 81, 0, 82, 108, 110, 103, 104, 105,
	0, 97, 98, 99, 102, 0, 0, 0, 84, 0,
	0, 0, 85, 0, 88, 87, 106, 107, 111, 109,
	113, 112, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 91, 92, 94, 95, 96, 93, 0, 0, 89,
	90, 100, 101, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 0, 0, 0, 0, 0, 0, 82, 108,
	110, 103, 104, 105, 0, 97, 98, 99, 102, 0,
	0, 0, 84, 0, 0, 0, 85, 0, 88, 87,
	106, 107, 111, 109, 113, 112, 0, 0, 0, 0,
	83, 0, 0, 0, 0, 91, 92, 94, 95, 96,
	93, 0, 0, 89, 90, 100, 101, 0, 0, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 0, 0,
	0, 0, 82, 108, 110, 103, 104, 105, 0, 97,
	98, 99, 102, 0, 0, 0, 159, 0, 0, 0,
	85, 0, 88, 87, 106, 107, 111, 109, 113, 112,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 91,
	92, 94, 95, 96, 93, 0, 0, 89, 90, 100,
	101, 0, 0, 0, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 0, 0, 82, 108, 110, 103,
	104, 105, 0, 97, 98, 99, 102, 0, 0, 0,
	157, 0, 0, 0, 85, 0, 88, 87, 106, 107,
	111, 109, 113, 112, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 89, 90, 100, 101, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 0, 0, 0, 0, 0, 0,
	82, 108, 110, 103, 104, 105, 0, 97, 98, 99,
	102, 0, 0, 0, 84, 0, 0, 0, 85, 0,
	88, 87, 106, 107, 111, 109, 113, 112, 0, 0,
	0, 0, 83, 0, 120, 52, 53, 0, 0, 32,
	0, 48, 0, 0, 0, 89, 90, 100, 101, 0,
	0, 0, 0, 39, 54, 55, 56, 0, 0, 0,
	0, 0, 0, 0, 82, 108, 110, 103, 104, 105,
	0, 97, 98, 99, 102, 0, 0, 0, 84, 0,
	40, 41, 85, 38, 88, 0, 42, 0, 0, 0,
	0, 0, 0, 51, 0, 58, 60, 0, 0, 59,
	0, 115, 0, 35, 0, 0, 118, 33, 0, 0,
	57, 36, 52, 53, 0, 0, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	39, 54, 55, 56, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 36, 52, 53, 0, 0,
	32, 0, 0, 0, 0, 0, 0, 40, 41, 0,
	38, 0, 0, 42, 39, 54, 55, 56, 0, 0,
	51, 0, 58, 60, 0, 0, 59, 0, 43, 0,
	35, 36, 52, 53, 33, 330, 32, 57, 0, 0,
	0, 40, 41, 0, 38, 0, 0, 42, 0, 0,
	39, 54, 55, 56, 51, 0, 58, 60, 0, 0,
	59, 0, 43, 0, 35, 0, 0, 0, 33, 301,
	0, 57, 0, 0, 0, 0, 0, 40, 41, 0,
	38, 0, 0, 42, 0, 87, 106, 107, 111, 109,
	51, 112, 58, 60, 0, 0, 59, 0, 43, 0,
	35, 0, 0, 256, 33, 0, 0, 57, 0, 89,
	90, 100, 101, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
	110, 103, 104, 105, 0, 97, 98, 99, 102, 36,
	52, 53, 84, 0, 32, 0, 85, 0, 88, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 39, 54,
	55, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 36, 52, 53, 0, 0,
	32, 0, 0, 0, 0, 40, 41, 0, 38, 0,
	0, 42, 0, 220, 39, 54, 55, 56, 51, 0,
	58, 60, 0, 0, 59, 0, 43, 0, 35, 36,
	52, 53, 33, 0, 32, 57, 0, 0, 0, 0,
	0, 40, 41, 0, 38, 0, 0, 42, 39, 54,
	55, 56, 0, 0, 51, 0, 58, 60, 0, 0,
	59, 0, 43, 0, 35, 0, 0, 202, 33, 0,
	0, 57, 0, 0, 0, 40, 41, 0, 38, 0,
	0, 42, 0, 169, 0, 0, 0, 0, 51, 0,
	58, 60, 0, 0, 59, 0, 43, 0, 35, 36,
	52, 53, 33, 0, 32, 57, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 39, 54,
	55, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 36, 52, 53, 0, 0, 32, 0,
	0, 0, 0, 0, 0, 40, 41, 0, 38, 0,
	0, 42, 39, 54, 55, 56, 0, 0, 51, 0,
	58, 60, 0, 0, 59, 0, 43, 0, 35, 36,
	52, 53, 33, 0, 32, 57, 0, 0, 0, 40,
	41, 0, 38, 0, 0, 42, 0, 0, 39, 54,
	55, 56, 51, 0, 58, 60, 0, 0, 59, 0,
	357, 0, 35, 36, 52, 53, 33, 0, 32, 57,
	0, 0, 0, 0, 0, 40, 41, 0, 38, 0,
	0, 42, 39, 54, 55, 56, 0, 0, 51, 0,
	58, 60, 0, 0, 59, 0, 312, 0, 35, 36,
	52, 53, 33, 0, 32, 57, 0, 0, 0, 40,
	41, 0, 38, 0, 0, 42, 0, 0, 39, 54,
	55, 56, 51, 0, 58, 60, 0, 0, 59, 0,
	310, 0, 35, 0, 0, 0, 33, 0, 0, 57,
	0, 0, 0, 0, 0, 40, 41, 0, 38, 0,
	0, 42, 87, 106, 107, 111, 109, 0, 51, 0,
	58, 60, 0, 0, 59, 0, 253, 0, 35, 0,
	0, 0, 33, 0, 0, 57, 89, 90, 100, 101,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 108, 110, 103, 104,
	105, 0, 97, 98, 99, 102, 36, 147, 53, 84,
	0, 32, 0, 85, 0, 88, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 39, 54, 55, 56, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	76, 52, 53, 0, 0, 32, 0, 0, 0, 0,
	0, 0, 40, 41, 0, 38, 0, 0, 42, 39,
	54, 55, 56, 0, 0, 51, 0, 58, 60, 0,
	0, 59, 0, 43, 0, 35, 74, 52, 53, 33,
	0, 32, 57, 0, 0, 0, 40, 41, 0, 38,
	0, 0, 42, 0, 0, 39, 54, 55, 56, 51,
	0, 58, 60, 0, 0, 59, 0, 43, 0, 35,
	0, 0, 0, 33, 0, 0, 57, 0, 0, 0,
	0, 0, 40, 41, 0, 38, 0, 0, 42, 87,
	0, 0, 0, 0, 0, 51, 0, 58, 60, 0,
	0, 59, 0, 43, 87, 35, 0, 0, 0, 33,
	0, 0, 57, 89, 90, 100, 101, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 89, 90,
	100, 101, 0, 0, 0, 103, 104, 105, 0, 97,
	98, 99, 102, 0, 0, 0, 84, 0, 0, 0,
	85, 0, 88, 0, 97, 98, 99, 102, 0, 0,
	0, 84, 0, 0, 0, 85, 0, 88,
}
var yyPact = [...]int{

	-65, -1000, 452, -65, -1000, -73, -73, -1000, -1000, -1000,
	-1000, -1000, -1000, 3375, 3375, 290, 203, 3692, 3656, 108,
	107, 269, -1000, -1000, 2635, -1000, -1000, 3375, 2980, 3375,
	-1000, -1000, 157, -54, 116, 3375, 86, -44, 106, 104,
	103, 3375, -8, -73, -1000, -1000, -1000, -1000, 287, 98,
	-1000, 3622, -1000, -1000, -1000, -1000, -1000, 3375, 3375, 3375,
	3375, -1000, -1000, -1000, -1000, -1000, 452, -73, -1000, 0,
	2699, 2699, 202, -65, 93, 2827, 89, 2763, 3375, 3375,
	253, 3375, 3375, 3375, 3375, 3305, 3375, 3375, 288, -1000,
	-1000, 3375, 3375, 3375, 3375, 3375, 3375, 3375, 3375, 3375,
	3375, 3375, 3375, 3375, 3375, 3375, 3375, 3375, 3375, 3375,
	3375, 3375, 3375, 3375, 2571, -65, 61, 1035, 3271, -5,
	86, 2507, 287, 84, -34, 3375, -73, -14, -1000, 116,
	116, -32, 116, -56, 2443, 3375, 3235, 3375, 116, 69,
	2891, 116, 3375, 66, -1000, 3375, -73, -1000, -11, -11,
	-11, -11, -11, -1000, -65, 190, 3375, 3375, 3375, 3375,
	971, 2379, 3375, -65, 2699, 2315, 2955, 179, 907, 3375,
	2891, -3, -1000, 2699, 2699, 2699, 2699, 2699, 2699, -3,
	-3, -3, -3, -3, -3, 3748, 3748, 3748, 3733, 3733,
	3733, 3733, 3733, 3733, 3556, 3169, -65, 188, -73, 3375,
	-73, -65, 3515, 2251, 3127, -73, 177, 287, -1000, -50,
	-73, 286, -69, -69, 116, -69, -34, -1000, 167, 843,
	3375, 2187, -39, -24, 285, -15, -58, 2123, 3375, 0,
	3375, 187, 254, 163, 159, 124, 122, -1000, 3375, -1000,
	2059, 186, 3375, 59, -1000, -1000, 3091, 779, 176, -1000,
	1995, 283, 173, -65, 1931, 3479, 3445, 1867, 234, 200,
	57, 120, -73, -66, -73, 3375, -1000, -30, 55, -1000,
	-1000, 3057, 715, -1000, -1000, -1000, 3375, 3, 116, 168,
	-73, 3375, 0, 2699, -44, -1000, 240, 52, -1000, 50,
	-1000, 17, -1000, 11, -1000, 1803, -65, -1000, 2891, -1000,
	651, -1000, -1000, 3375, -1000, -65, -1000, -1000, 164, -65,
	-65, 1739, -65, 1675, 3409, -25, -1000, -1000, 217, 3375,
	-65, 199, 198, 7, -73, -1000, -50, 116, -1000, 587,
	-1000, -1000, 3375, 523, 3375, -45, -1000, 3375, 2699, 196,
	-65, -1000, -1000, -1000, -1000, -1000, 155, -1000, 3375, 1611,
	153, -1000, 151, 150, -65, 148, -65, -65, 1547, 147,
	-1000, -1000, -65, 1483, 14, 146, -65, -65, 194, 143,
	-69, -1000, 3375, 1419, -1000, 3375, 1355, -73, 1291, -65,
	142, -1000, 1227, -1000, -1000, -1000, -1000, 141, -1000, 140,
	138, -65, -1000, -1000, -65, -65, -1000, 137, 135, -65,
	-1000, 1163, -1000, 1099, -1000, 3375, 3375, 132, 250, -1000,
	-1000, -1000, -1000, 130, -1000, -1000, -1000, -1000, 129, -1000,
	-1000, -58, 2699, 245, 193, -1000, -1000, 114, 117, -65,
	-1000, -65, 112, 111, -1000, -1000,
}
var yyPgo = [...]int{

	0, 53, 334, 278, 285, 332, 331, 330, 329, 325,
	319, 7, 2, 32, 0, 41, 14, 9, 316, 313,
	4, 310, 6, 309, 308, 306, 305, 303, 302, 301,
	297, 296, 295, 5, 1, 116, 47,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 4, 5, 6, 6, 7, 7, 7, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	9, 10, 10, 10, 10, 10, 11, 11, 12, 13,
	13, 13, 13, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 15, 15, 15, 16,
	16, 16, 16, 16, 16, 17, 17, 18, 18, 19,
	20, 21, 21, 21, 21, 21, 21, 22, 22, 22,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	24, 24, 24, 24, 24, 25, 25, 25, 25, 26,
	26, 26, 26, 26, 26, 26, 26, 30, 30, 30,
	30, 30, 30, 29, 29, 29, 28, 28, 28, 28,
	28, 28, 27, 27, 31, 31, 32, 32, 32, 33,
	33, 35, 35, 36, 34, 34, 34, 34,
}
var yyR2 = [...]int{

	0, 1, 2, 2, 3, 0, 1, 1, 1, 2,
	2, 5, 13, 12, 9, 8, 6, 5, 6, 5,
	6, 5, 6, 5, 4, 6, 4, 1, 1, 1,
	1, 1, 1, 4, 3, 3, 5, 7, 5, 4,
	7, 5, 6, 7, 7, 8, 7, 8, 8, 9,
	7, 0, 1, 1, 2, 2, 4, 4, 3, 0,
	1, 4, 4, 1, 1, 5, 3, 7, 8, 8,
	9, 2, 5, 7, 3, 5, 4, 5, 4, 4,
	4, 4, 4, 4, 6, 8, 7, 3, 2, 3,
	10, 5, 1, 1, 1, 1, 0, 1, 4, 1,
	3, 2, 2, 5, 2, 2, 3, 1, 1, 3,
	1, 2, 1, 1, 1, 1, 1, 0, 3, 6,
	6, 5, 5, 7, 8, 6, 5, 5, 7, 8,
	2, 2, 2, 2, 2, 1, 1, 1, 1, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 0, 1, 2, 1, 1, 0,
	1, 1, 2, 1, 0, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -31, -2, -32, 76, -35, -36, 81, -3,
	-4, 38, 39, 10, 12, 28, 29, 47, 48, 54,
	55, -7, -8, -9, -14, -5, -6, 13, 15, 44,
	-18, -21, 9, 77, -17, 73, 4, -20, 53, 23,
	50, 51, 56, 71, -23, -24, -25, -26, 11, -13,
	-19, 63, 5, 6, 24, 25, 26, 80, 65, 69,
	66, -30, -29, -28, -27, -31, -32, -35, -36, -13,
	-14, -14, 4, 71, 4, -14, 4, -14, 73, 73,
	14, 57, 59, 27, 73, 77, 51, 16, 79, 40,
	41, 32, 33, 37, 34, 35, 36, 66, 67, 68,
	42, 43, 69, 62, 63, 64, 17, 18, 60, 20,
	61, 19, 22, 21, -14, 71, -15, -14, 76, -4,
	4, -14, 73, 4, 78, -33, -35, -16, 4, 66,
	-17, 56, 49, 77, -14, 73, 77, 73, 73, 73,
	-14, 77, -33, -15, 4, 57, 75, 5, -14, -14,
	-14, -14, -14, -3, 71, -1, 73, 73, 73, 73,
	-14, -14, 13, 71, -14, -14, -14, -13, -14, 58,
	-14, -14, 4, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -14, 71, -1, -35, 16,
	75, 71, 76, -14, 76, 71, -15, 73, -17, -13,
	71, 79, -16, -16, 77, -16, 78, 74, -13, -14,
	58, -14, -16, -16, 52, -16, -22, -14, 57, -13,
	-33, -1, 72, -13, -13, -13, -13, 74, 75, 74,
	-14, -1, 58, 8, 74, 78, 58, -14, -1, 72,
	-14, -33, -1, 71, -14, 76, 76, -14, -33, 74,
	8, -15, 75, -34, -35, -33, 4, -16, 8, 74,
	78, 58, -14, 74, 74, 74, 75, 4, 78, -34,
	75, 58, -13, -14, -20, 72, 30, 8, 74, 8,
	74, 8, 74, 8, 74, -14, 71, 72, -14, 74,
	-14, 78, 78, 58, 72, 71, 4, 72, -1, 71,
	71, -14, 71, -14, 76, -10, -12, -11, 46, 45,
	71, 74, 74, 8, -35, 78, -13, 78, 74, -14,
	78, 78, 58, -14, 75, -16, 72, -33, -14, 4,
	71, 74, 74, 74, 74, 74, -1, 78, 58, -14,
	-1, 72, -1, -1, 71, -1, 71, 71, -14, -33,
	-11, -12, 58, -14, -13, -1, 71, 71, 74, -34,
	-16, 78, 58, -14, 74, 75, -14, 71, -14, 71,
	-1, 72, -14, 78, 72, 72, 72, -1, 72, -1,
	-1, 71, 72, -1, 58, 58, 72, -1, -1, 71,
	72, -14, 78, -14, 74, -33, 58, -1, 72, 78,
	72, 72, 72, -1, -1, -1, 72, 72, -1, 78,
	74, -22, -14, 72, 31, 72, 72, -34, 31, 71,
	72, 71, -1, -1, 72, 72,
}
var yyDef = [...]int{

	164, -2, -2, 164, 165, 168, 167, 171, 173, 3,
	6, 7, 8, 59, 0, 0, 0, 0, 0, 0,
	0, 27, 28, 29, -2, 31, 32, 0, -2, 0,
	63, 64, 0, 169, 0, 0, 110, 108, 0, 0,
	0, 0, 0, 169, 92, 93, 94, 95, 96, 0,
	107, 0, 112, 113, 114, 115, 116, 0, 0, 0,
	0, 135, 136, 137, 138, 2, -2, 166, 172, 9,
	60, 10, 0, 164, 110, 0, 110, 0, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 139,
	140, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 164, 0, 60, 0, 0,
	-2, 0, 96, 0, -2, 59, 170, 0, 99, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 0, 0,
	88, 0, 117, 0, 97, 59, 169, 111, 130, 131,
	132, 133, 134, 4, 164, 0, 59, 59, 59, 59,
	0, 0, 0, 164, 34, 0, 66, 0, 0, 0,
	87, 89, 109, 141, 142, 143, 144, 145, 146, 147,
	148, 149, 150, 151, 152, 153, 154, 155, 156, 157,
	158, 159, 160, 161, 162, 163, 164, 0, 167, 0,
	169, 164, 0, 0, 0, 169, 0, 96, 106, 174,
	169, 0, 101, 102, 0, 104, 105, 74, 0, 0,
	0, 0, 0, 0, 0, 0, 174, 0, 59, 35,
	0, 0, 0, 0, 0, 0, 0, 24, 0, 26,
	0, 0, 0, 0, 78, 80, 0, 0, 0, 39,
	0, 0, 0, 164, 0, 0, 0, 0, 51, 0,
	0, 0, -2, 0, 176, 59, 100, 0, 0, 76,
	79, 0, 0, 81, 82, 83, 0, 0, 0, 0,
	-2, 0, 33, 61, -2, 11, 0, 0, -2, 0,
	-2, 0, -2, 0, -2, 0, 164, 38, 65, 77,
	0, 126, 127, 0, 36, 164, 98, 41, 0, 164,
	164, 0, 164, 0, 0, 169, 52, 53, 0, 59,
	164, 0, 0, 0, -2, 72, 174, 0, 75, 0,
	121, 122, 0, 0, 0, 0, 91, 0, 118, 0,
	164, -2, -2, -2, -2, 25, 0, 125, 0, 0,
	0, 42, 0, 0, 164, 0, 164, 164, 0, 0,
	54, 55, 164, 60, 0, 0, 164, 164, 0, 0,
	103, 120, 0, 0, 84, 0, 0, 169, 0, 164,
	0, 37, 0, 128, 40, 43, 44, 0, 46, 0,
	0, 164, 50, 58, 164, 164, 67, 0, 0, 164,
	73, 0, 123, 0, 86, 117, 0, 0, 15, 129,
	45, 47, 48, 0, 56, 57, 68, 69, 0, 124,
	85, 174, 119, 14, 0, 49, 70, 0, 0, 164,
	90, 164, 0, 0, 13, 12,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	81, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 80, 3, 3, 3, 68, 69, 3,
	73, 74, 66, 62, 75, 63, 79, 67, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 58, 76,
	60, 57, 61, 59, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 77, 3, 78, 65, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 71, 64, 72,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 70,
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
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:212
		{
			yyVAL.stmt = &ast.DeferStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Defer: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:217
		{
			yyVAL.stmt = &ast.DeferStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Defer: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:222
		{
			yyVAL.stmt = &ast.DeferStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Defer: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].expr.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:227
		{
			yyVAL.stmt = &ast.DeferStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Defer: true}}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:232
		{
			yyVAL.stmt = &ast.DeleteStmt{Item: yyDollar[3].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:237
		{
			yyVAL.stmt = &ast.DeleteStmt{Item: yyDollar[3].expr, Key: yyDollar[5].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:242
		{
			yyVAL.stmt = &ast.CloseStmt{Expr: yyDollar[3].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:247
		{
			yyVAL.stmt = yyDollar[1].stmt_if
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:251
		{
			yyVAL.stmt = yyDollar[1].stmt_for
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:255
		{
			yyVAL.stmt = yyDollar[1].stmt_switch
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:259
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:266
		{
			yyVAL.stmt_var_or_lets = yyDollar[1].stmt_var
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:270
		{
			yyVAL.stmt_var_or_lets = yyDollar[1].stmt_lets
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:276
		{
			yyVAL.stmt_var = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].exprs}
			yyVAL.stmt_var.SetPosition(yyDollar[1].tok.Position())
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:283
		{
			yyVAL.stmt_lets = &ast.LetsStmt{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{yyDollar[3].expr}}
			yyVAL.stmt_lets.SetPosition(yyDollar[1].expr.Position())
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:288
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
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 37:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:307
		{
			ifStmt := yyDollar[1].stmt_if.(*ast.IfStmt)
			ifStmt.ElseIf = append(ifStmt.ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:312
		{
			ifStmt := yyDollar[1].stmt_if.(*ast.IfStmt)
			if ifStmt.Else != nil {
				yylex.Error("multiple else statement")
			}
			ifStmt.Else = yyDollar[4].compstmt
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:322
		{
			yyVAL.stmt_for = &ast.LoopStmt{Stmt: yyDollar[3].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:327
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
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:338
		{
			yyVAL.stmt_for = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmt: yyDollar[4].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:343
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt: yyDollar[5].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:348
		{
			yyVAL.stmt_for = &ast.CForStmt{Expr3: yyDollar[4].expr, Stmt: yyDollar[6].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 44:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:353
		{
			yyVAL.stmt_for = &ast.CForStmt{Expr2: yyDollar[3].expr, Stmt: yyDollar[6].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 45:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:358
		{
			yyVAL.stmt_for = &ast.CForStmt{Expr2: yyDollar[3].expr, Expr3: yyDollar[5].expr, Stmt: yyDollar[7].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 46:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:363
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt1: yyDollar[2].stmt_var_or_lets, Stmt: yyDollar[6].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 47:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:368
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt1: yyDollar[2].stmt_var_or_lets, Expr3: yyDollar[5].expr, Stmt: yyDollar[7].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 48:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:373
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt1: yyDollar[2].stmt_var_or_lets, Expr2: yyDollar[4].expr, Stmt: yyDollar[7].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 49:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:378
		{
			yyVAL.stmt_for = &ast.CForStmt{Stmt1: yyDollar[2].stmt_var_or_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmt: yyDollar[8].compstmt}
			yyVAL.stmt_for.SetPosition(yyDollar[1].tok.Position())
		}
	case 50:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:385
		{
			switchStmt := yyDollar[5].stmt_switch_cases.(*ast.SwitchStmt)
			switchStmt.Expr = yyDollar[2].expr
			yyVAL.stmt_switch = switchStmt
			yyVAL.stmt_switch.SetPosition(yyDollar[1].tok.Position())
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:394
		{
			yyVAL.stmt_switch_cases = &ast.SwitchStmt{}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:398
		{
			yyVAL.stmt_switch_cases = &ast.SwitchStmt{Default: yyDollar[1].stmt_switch_default}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:402
		{
			yyVAL.stmt_switch_cases = &ast.SwitchStmt{Cases: []ast.Stmt{yyDollar[1].stmt_switch_case}}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:406
		{
			switchStmt := yyDollar[1].stmt_switch_cases.(*ast.SwitchStmt)
			switchStmt.Cases = append(switchStmt.Cases, yyDollar[2].stmt_switch_case)
			yyVAL.stmt_switch_cases = switchStmt
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:412
		{
			switchStmt := yyDollar[1].stmt_switch_cases.(*ast.SwitchStmt)
			if switchStmt.Default != nil {
				yylex.Error("multiple default statement")
			}
			switchStmt.Default = yyDollar[2].stmt_switch_default
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:422
		{
			yyVAL.stmt_switch_case = &ast.SwitchCaseStmt{Exprs: []ast.Expr{yyDollar[2].expr}, Stmt: yyDollar[4].compstmt}
			yyVAL.stmt_switch_case.SetPosition(yyDollar[1].tok.Position())
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:427
		{
			yyVAL.stmt_switch_case = &ast.SwitchCaseStmt{Exprs: yyDollar[2].exprs, Stmt: yyDollar[4].compstmt}
			yyVAL.stmt_switch_case.SetPosition(yyDollar[1].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:434
		{
			yyVAL.stmt_switch_default = yyDollar[3].compstmt
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:441
		{
			yyVAL.exprs = nil
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:445
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:449
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:456
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr_ident)
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:465
		{
			yyVAL.expr = yyDollar[1].expr_member_or_ident
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:469
		{
			yyVAL.expr = yyDollar[1].expr_literals
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:473
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, LHS: yyDollar[3].expr, RHS: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:478
		{
			yyVAL.expr = &ast.NilCoalescingOpExpr{LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:483
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmt: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:488
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmt: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:493
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmt: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:498
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmt: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:503
		{
			yyVAL.expr = &ast.ArrayExpr{}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:508
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 73:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:513
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[5].exprs, TypeData: &ast.TypeStruct{Kind: ast.TypeSlice, SubType: yyDollar[2].type_data, Dimensions: yyDollar[1].slice_count}}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:518
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:523
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:528
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:533
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:538
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:543
		{
			yyVAL.expr = &ast.ItemExpr{Item: yyDollar[1].expr_ident, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr_ident.Position())
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:548
		{
			yyVAL.expr = &ast.ItemExpr{Item: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:553
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:558
		{
			if yyDollar[3].type_data.Kind == ast.TypeDefault {
				yyDollar[3].type_data.Kind = ast.TypePtr
				yyVAL.expr = &ast.MakeExpr{TypeData: yyDollar[3].type_data}
			} else {
				yyVAL.expr = &ast.MakeExpr{TypeData: &ast.TypeStruct{Kind: ast.TypePtr, SubType: yyDollar[3].type_data}}
			}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:568
		{
			yyVAL.expr = &ast.MakeExpr{TypeData: yyDollar[3].type_data}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 84:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:573
		{
			yyVAL.expr = &ast.MakeExpr{TypeData: yyDollar[3].type_data, LenExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 85:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:578
		{
			yyVAL.expr = &ast.MakeExpr{TypeData: yyDollar[3].type_data, LenExpr: yyDollar[5].expr, CapExpr: yyDollar[7].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 86:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:583
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].tok.Lit, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:588
		{
			yyVAL.expr = &ast.ChanExpr{LHS: yyDollar[1].expr, RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:593
		{
			yyVAL.expr = &ast.ChanExpr{RHS: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:598
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.go.y:603
		{
			yyDollar[8].expr_map.TypeData = &ast.TypeStruct{Kind: ast.TypeMap, Key: yyDollar[3].type_data, SubType: yyDollar[5].type_data}
			yyVAL.expr = yyDollar[8].expr_map
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:609
		{
			yyVAL.expr = yyDollar[3].expr_map
			yyVAL.expr.SetPosition(yyDollar[3].expr_map.Position())
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:614
		{
			yyVAL.expr = yyDollar[1].expr_slice
			yyVAL.expr.SetPosition(yyDollar[1].expr_slice.Position())
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:623
		{
			yyVAL.expr_idents = []string{}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:627
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:631
		{
			if len(yyDollar[1].expr_idents) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:640
		{
			yyVAL.type_data = &ast.TypeStruct{Name: yyDollar[1].tok.Lit}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:644
		{
			if yyDollar[1].type_data.Kind != ast.TypeDefault {
				yylex.Error("blah1")
			} else {
				yyDollar[1].type_data.Env = append(yyDollar[1].type_data.Env, yyDollar[1].type_data.Name)
				yyDollar[1].type_data.Name = yyDollar[3].tok.Lit
			}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:653
		{
			if yyDollar[2].type_data.Kind == ast.TypeDefault {
				yyDollar[2].type_data.Kind = ast.TypePtr
				yyVAL.type_data = yyDollar[2].type_data
			} else {
				yyVAL.type_data = &ast.TypeStruct{Kind: ast.TypePtr, SubType: yyDollar[2].type_data}
			}
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:662
		{
			if yyDollar[2].type_data.Kind == ast.TypeDefault {
				yyDollar[2].type_data.Kind = ast.TypeSlice
				yyDollar[2].type_data.Dimensions = yyDollar[1].slice_count
				yyVAL.type_data = yyDollar[2].type_data
			} else {
				yyVAL.type_data = &ast.TypeStruct{Kind: ast.TypeSlice, SubType: yyDollar[2].type_data, Dimensions: yyDollar[1].slice_count}
			}
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:672
		{
			yyVAL.type_data = &ast.TypeStruct{Kind: ast.TypeMap, Key: yyDollar[3].type_data, SubType: yyDollar[5].type_data}
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:676
		{
			if yyDollar[2].type_data.Kind == ast.TypeDefault {
				yyDollar[2].type_data.Kind = ast.TypeChan
				yyVAL.type_data = yyDollar[2].type_data
			} else {
				yyVAL.type_data = &ast.TypeStruct{Kind: ast.TypeChan, SubType: yyDollar[2].type_data}
			}
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:688
		{
			yyVAL.slice_count = 1
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:692
		{
			yyVAL.slice_count = yyDollar[3].slice_count + 1
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:698
		{
			yyVAL.expr_member_or_ident = yyDollar[1].expr_member
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:702
		{
			yyVAL.expr_member_or_ident = yyDollar[1].expr_ident
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:708
		{
			yyVAL.expr_member = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr_member.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:715
		{
			yyVAL.expr_ident = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr_ident.SetPosition(yyDollar[1].tok.Position())
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:722
		{
			num, err := toNumber("-" + yyDollar[2].tok.Lit)
			if err != nil {
				yylex.Error("invalid number: -" + yyDollar[2].tok.Lit)
			}
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: num}
			yyVAL.expr_literals.SetPosition(yyDollar[2].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:731
		{
			num, err := toNumber(yyDollar[1].tok.Lit)
			if err != nil {
				yylex.Error("invalid number: " + yyDollar[1].tok.Lit)
			}
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: num}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:740
		{
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: stringToValue(yyDollar[1].tok.Lit)}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:745
		{
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: trueValue}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:750
		{
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: falseValue}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:755
		{
			yyVAL.expr_literals = &ast.LiteralExpr{Literal: nilValue}
			yyVAL.expr_literals.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:762
		{
			yyVAL.expr_map = &ast.MapExpr{}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:766
		{
			yyVAL.expr_map = &ast.MapExpr{Keys: []ast.Expr{yyDollar[1].expr}, Values: []ast.Expr{yyDollar[3].expr}}
		}
	case 119:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:770
		{
			if yyDollar[1].expr_map.Keys == nil {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_map.Keys = append(yyVAL.expr_map.Keys, yyDollar[4].expr)
			yyVAL.expr_map.Values = append(yyVAL.expr_map.Values, yyDollar[6].expr)
		}
	case 120:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:780
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr_ident, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:784
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr_ident, Begin: yyDollar[3].expr, End: nil}
		}
	case 122:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:788
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr_ident, Begin: nil, End: yyDollar[4].expr}
		}
	case 123:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:792
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr_ident, End: yyDollar[4].expr, Cap: yyDollar[6].expr}
		}
	case 124:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:796
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr_ident, Begin: yyDollar[3].expr, End: yyDollar[5].expr, Cap: yyDollar[7].expr}
		}
	case 125:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:800
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 126:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:804
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
		}
	case 127:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:808
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
		}
	case 128:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:812
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr, End: yyDollar[4].expr, Cap: yyDollar[6].expr}
		}
	case 129:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:816
		{
			yyVAL.expr_slice = &ast.SliceExpr{Item: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr, Cap: yyDollar[7].expr}
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:822
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:827
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:832
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:837
		{
			yyVAL.expr = &ast.AddrExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:842
		{
			yyVAL.expr = &ast.DerefExpr{Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:849
		{
			yyVAL.expr = &ast.OpExpr{Op: yyDollar[1].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:854
		{
			yyVAL.expr = &ast.OpExpr{Op: yyDollar[1].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:859
		{
			yyVAL.expr = &ast.OpExpr{Op: yyDollar[1].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:864
		{
			yyVAL.expr = &ast.OpExpr{Op: yyDollar[1].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:871
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "+", RHS: oneLiteral}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:879
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "-", RHS: oneLiteral}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:887
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "+", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:895
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "-", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:903
		{
			rhs := &ast.OpExpr{Op: &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "|", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:911
		{
			rhs := &ast.OpExpr{Op: &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "*", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:919
		{
			rhs := &ast.OpExpr{Op: &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "/", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:927
		{
			rhs := &ast.OpExpr{Op: &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "&", RHS: yyDollar[3].expr}}
			rhs.Op.SetPosition(yyDollar[1].expr.Position())
			rhs.SetPosition(yyDollar[1].expr.Position())
			yyVAL.expr = &ast.LetsExpr{LHSS: []ast.Expr{yyDollar[1].expr}, RHSS: []ast.Expr{rhs}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:938
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "*", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:943
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "/", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:948
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "%", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:953
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "<<", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:958
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: ">>", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:963
		{
			yyVAL.expr = &ast.MultiplyOperator{LHS: yyDollar[1].expr, Operator: "&", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:970
		{
			yyVAL.expr = &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "+", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:975
		{
			yyVAL.expr = &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "-", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:980
		{
			yyVAL.expr = &ast.AddOperator{LHS: yyDollar[1].expr, Operator: "|", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:987
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: "==", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:992
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: "!=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:997
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: "<", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1002
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: "<=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1007
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: ">", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1012
		{
			yyVAL.expr = &ast.ComparisonOperator{LHS: yyDollar[1].expr, Operator: ">=", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1019
		{
			yyVAL.expr = &ast.BinaryOperator{LHS: yyDollar[1].expr, Operator: "&&", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:1024
		{
			yyVAL.expr = &ast.BinaryOperator{LHS: yyDollar[1].expr, Operator: "||", RHS: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	}
	goto yystack /* stack new state and value */
}
