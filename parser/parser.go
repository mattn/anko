//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:28
type yySymType struct {
	yys           int
	compstmt      []ast.Stmt
	stmt_if       ast.Stmt
	stmt_default  ast.Stmt
	stmt_case     ast.Stmt
	stmt_cases    []ast.Stmt
	stmts         []ast.Stmt
	stmt          ast.Stmt
	typ           ast.Type
	expr          ast.Expr
	exprs         []ast.Expr
	expr_many     []ast.Expr
	expr_lets     ast.Expr
	expr_pair     ast.Expr
	expr_pairs    []ast.Expr
	expr_idents   []string
	tok           ast.Token
	term          ast.Token
	terms         ast.Token
	opt_terms     ast.Token
	struct_field  ast.StructField
	struct_fields []ast.StructField
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
const STRUCT = 57370
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
const ARRAYLIT = 57394
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
	"STRUCT",
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
	"ARRAYLIT",
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
	"'.'",
	"'!'",
	"'^'",
	"'&'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'|'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:739

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 127,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	56, 49,
	-2, 1,
	-1, 10,
	56, 50,
	-2, 24,
	-1, 43,
	56, 49,
	-2, 128,
	-1, 85,
	66, 3,
	-2, 127,
	-1, 88,
	56, 50,
	-2, 43,
	-1, 90,
	66, 3,
	-2, 127,
	-1, 97,
	1, 58,
	8, 58,
	46, 58,
	47, 58,
	53, 58,
	55, 58,
	56, 58,
	65, 58,
	66, 58,
	67, 58,
	73, 58,
	75, 58,
	77, 58,
	-2, 53,
	-1, 99,
	1, 60,
	8, 60,
	46, 60,
	47, 60,
	53, 60,
	55, 60,
	56, 60,
	65, 60,
	66, 60,
	67, 60,
	73, 60,
	75, 60,
	77, 60,
	-2, 53,
	-1, 127,
	17, 0,
	18, 0,
	-2, 86,
	-1, 128,
	17, 0,
	18, 0,
	-2, 87,
	-1, 147,
	56, 50,
	-2, 43,
	-1, 149,
	66, 3,
	-2, 127,
	-1, 151,
	66, 3,
	-2, 127,
	-1, 153,
	66, 1,
	-2, 36,
	-1, 156,
	66, 3,
	-2, 127,
	-1, 181,
	66, 3,
	-2, 127,
	-1, 225,
	56, 51,
	-2, 44,
	-1, 226,
	1, 45,
	46, 45,
	47, 45,
	53, 45,
	56, 52,
	66, 45,
	67, 45,
	77, 45,
	-2, 53,
	-1, 233,
	1, 52,
	8, 52,
	46, 52,
	47, 52,
	56, 52,
	66, 52,
	67, 52,
	73, 52,
	75, 52,
	77, 52,
	-2, 53,
	-1, 235,
	66, 3,
	-2, 127,
	-1, 237,
	66, 3,
	-2, 127,
	-1, 250,
	66, 3,
	-2, 127,
	-1, 262,
	1, 107,
	8, 107,
	46, 107,
	47, 107,
	53, 107,
	55, 107,
	56, 107,
	65, 107,
	66, 107,
	67, 107,
	73, 107,
	75, 107,
	77, 107,
	-2, 105,
	-1, 264,
	1, 111,
	8, 111,
	46, 111,
	47, 111,
	53, 111,
	55, 111,
	56, 111,
	65, 111,
	66, 111,
	67, 111,
	73, 111,
	75, 111,
	77, 111,
	-2, 109,
	-1, 274,
	66, 3,
	-2, 127,
	-1, 279,
	66, 3,
	-2, 127,
	-1, 280,
	66, 3,
	-2, 127,
	-1, 288,
	1, 106,
	8, 106,
	46, 106,
	47, 106,
	53, 106,
	55, 106,
	56, 106,
	65, 106,
	66, 106,
	67, 106,
	73, 106,
	75, 106,
	77, 106,
	-2, 104,
	-1, 289,
	1, 110,
	8, 110,
	46, 110,
	47, 110,
	53, 110,
	55, 110,
	56, 110,
	65, 110,
	66, 110,
	67, 110,
	73, 110,
	75, 110,
	77, 110,
	-2, 108,
	-1, 294,
	66, 3,
	-2, 127,
	-1, 295,
	66, 3,
	-2, 127,
	-1, 298,
	46, 3,
	47, 3,
	66, 3,
	-2, 127,
	-1, 302,
	66, 3,
	-2, 127,
	-1, 311,
	46, 3,
	47, 3,
	66, 3,
	-2, 127,
	-1, 326,
	66, 3,
	-2, 127,
	-1, 327,
	66, 3,
	-2, 127,
}

const yyNprod = 133
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 2291

var yyAct = [...]int{

	81, 169, 242, 10, 172, 243, 286, 45, 255, 289,
	1, 11, 214, 92, 166, 93, 82, 115, 212, 88,
	40, 91, 6, 6, 94, 95, 96, 98, 100, 6,
	89, 154, 7, 7, 252, 80, 105, 288, 108, 7,
	110, 175, 112, 93, 10, 217, 281, 251, 116, 117,
	222, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 136, 137,
	138, 263, 261, 139, 140, 141, 142, 248, 144, 145,
	147, 266, 2, 115, 230, 217, 42, 92, 109, 93,
	218, 146, 202, 217, 161, 106, 152, 143, 265, 185,
	208, 158, 64, 65, 66, 67, 68, 69, 102, 164,
	70, 71, 55, 160, 177, 147, 103, 104, 217, 148,
	148, 78, 267, 155, 167, 287, 182, 331, 170, 50,
	51, 52, 53, 54, 217, 330, 264, 262, 49, 323,
	148, 74, 76, 320, 77, 319, 72, 148, 314, 313,
	273, 191, 244, 245, 10, 195, 196, 203, 147, 310,
	190, 299, 192, 293, 186, 292, 268, 197, 153, 198,
	150, 257, 241, 239, 210, 236, 101, 234, 199, 193,
	327, 225, 326, 223, 224, 229, 276, 316, 284, 231,
	232, 302, 227, 295, 180, 280, 220, 221, 183, 279,
	250, 219, 149, 90, 246, 114, 249, 247, 115, 148,
	111, 274, 157, 233, 22, 28, 216, 258, 32, 79,
	244, 245, 151, 325, 321, 240, 84, 8, 173, 173,
	253, 189, 36, 29, 30, 31, 209, 170, 5, 272,
	287, 259, 201, 44, 228, 275, 270, 211, 271, 207,
	213, 215, 174, 174, 206, 232, 165, 37, 283, 38,
	39, 278, 118, 83, 46, 285, 168, 290, 291, 23,
	27, 113, 87, 178, 34, 200, 179, 17, 24, 25,
	26, 35, 44, 33, 282, 296, 3, 0, 0, 4,
	300, 301, 304, 43, 0, 254, 0, 256, 0, 0,
	0, 0, 260, 0, 0, 308, 309, 0, 318, 312,
	317, 0, 0, 315, 0, 0, 0, 0, 0, 0,
	0, 0, 322, 0, 0, 0, 0, 0, 0, 0,
	58, 59, 61, 63, 73, 75, 0, 328, 329, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 298,
	0, 0, 0, 0, 78, 0, 0, 48, 303, 307,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	311, 49, 0, 0, 74, 76, 306, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 205, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 204, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 188, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 187, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 49, 0, 0, 74, 76, 324, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 49, 0, 0, 74, 76, 305, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 297, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 294, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 277, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 269, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	238, 49, 0, 0, 74, 76, 0, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 237, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 235, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 184, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 181, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 49, 0, 0, 74, 76, 171, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 159, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 0, 72,
	58, 59, 61, 63, 73, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 64, 65, 66, 67, 68,
	69, 0, 0, 70, 71, 55, 56, 57, 0, 0,
	0, 0, 0, 0, 78, 0, 0, 48, 0, 0,
	60, 62, 50, 51, 52, 53, 54, 0, 156, 0,
	0, 49, 0, 0, 74, 76, 0, 77, 0, 72,
	21, 22, 28, 0, 0, 32, 14, 9, 15, 41,
	0, 18, 0, 0, 0, 0, 0, 0, 0, 36,
	29, 30, 31, 16, 0, 19, 0, 0, 0, 0,
	0, 0, 0, 0, 12, 13, 0, 0, 0, 0,
	0, 20, 0, 0, 37, 0, 38, 39, 0, 0,
	0, 0, 0, 0, 0, 0, 23, 27, 0, 0,
	0, 34, 0, 6, 0, 24, 25, 26, 35, 0,
	33, 0, 0, 7, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	65, 66, 67, 68, 69, 0, 0, 70, 71, 55,
	56, 57, 0, 0, 0, 0, 0, 0, 78, 0,
	47, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 49, 0, 0, 74, 76,
	0, 77, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	65, 66, 67, 68, 69, 0, 0, 70, 71, 55,
	56, 57, 0, 0, 0, 0, 0, 0, 78, 0,
	0, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 49, 0, 0, 74, 76,
	0, 77, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	65, 66, 67, 68, 69, 0, 0, 70, 71, 55,
	56, 57, 0, 0, 0, 0, 0, 0, 78, 0,
	0, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 49, 0, 0, 74, 176,
	0, 77, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	65, 66, 67, 68, 69, 0, 0, 70, 71, 55,
	56, 57, 0, 0, 0, 0, 0, 0, 78, 0,
	0, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 0, 0, 0, 0, 163, 0, 0, 74, 76,
	0, 77, 0, 72, 58, 59, 61, 63, 73, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 64,
	65, 66, 67, 68, 69, 0, 0, 70, 71, 55,
	56, 57, 0, 0, 0, 0, 0, 0, 78, 0,
	0, 48, 0, 0, 60, 62, 50, 51, 52, 53,
	54, 58, 59, 61, 63, 162, 75, 0, 74, 76,
	0, 77, 0, 72, 0, 0, 64, 65, 66, 67,
	68, 69, 0, 0, 70, 71, 55, 56, 57, 0,
	0, 0, 0, 0, 0, 78, 0, 0, 0, 0,
	0, 60, 62, 50, 51, 52, 53, 54, 58, 59,
	61, 63, 49, 0, 0, 74, 76, 0, 77, 0,
	72, 0, 0, 64, 65, 66, 67, 68, 69, 0,
	0, 70, 71, 55, 56, 57, 0, 0, 0, 0,
	0, 0, 78, 0, 0, 0, 0, 0, 60, 62,
	50, 51, 52, 53, 54, 0, 0, 0, 0, 49,
	0, 0, 74, 76, 0, 77, 0, 72, 21, 22,
	194, 0, 0, 32, 14, 9, 15, 41, 0, 18,
	0, 0, 0, 0, 0, 0, 0, 36, 29, 30,
	31, 16, 0, 19, 0, 0, 0, 0, 0, 0,
	0, 0, 12, 13, 0, 0, 0, 0, 0, 20,
	0, 0, 37, 0, 38, 39, 0, 0, 0, 0,
	0, 0, 0, 0, 23, 27, 0, 0, 0, 34,
	0, 0, 0, 24, 25, 26, 35, 0, 33, 21,
	22, 28, 0, 0, 32, 14, 9, 15, 41, 0,
	18, 0, 0, 0, 0, 0, 0, 0, 36, 29,
	30, 31, 16, 0, 19, 0, 0, 0, 0, 0,
	0, 0, 0, 12, 13, 0, 0, 0, 0, 0,
	20, 0, 0, 37, 0, 38, 39, 0, 0, 0,
	0, 0, 0, 0, 0, 23, 27, 61, 63, 0,
	34, 0, 0, 0, 24, 25, 26, 35, 0, 33,
	64, 65, 66, 67, 68, 69, 0, 0, 70, 71,
	55, 56, 57, 21, 22, 28, 0, 0, 32, 78,
	0, 0, 0, 0, 0, 60, 62, 50, 51, 52,
	53, 54, 36, 29, 30, 31, 49, 0, 0, 74,
	76, 0, 77, 0, 72, 233, 22, 28, 0, 0,
	32, 0, 0, 0, 0, 0, 0, 37, 0, 38,
	39, 0, 0, 0, 36, 29, 30, 31, 0, 23,
	27, 0, 0, 0, 34, 0, 0, 0, 24, 25,
	26, 35, 0, 33, 0, 0, 0, 0, 0, 37,
	0, 38, 39, 0, 0, 0, 0, 0, 226, 22,
	28, 23, 27, 32, 0, 0, 34, 0, 0, 0,
	24, 25, 26, 35, 0, 33, 0, 36, 29, 30,
	31, 0, 0, 0, 0, 0, 0, 107, 22, 28,
	0, 0, 32, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 37, 0, 38, 39, 36, 29, 30, 31,
	0, 0, 0, 0, 23, 27, 99, 22, 28, 34,
	0, 32, 0, 24, 25, 26, 35, 0, 33, 0,
	0, 37, 0, 38, 39, 36, 29, 30, 31, 0,
	0, 0, 0, 23, 27, 97, 22, 28, 34, 0,
	32, 0, 24, 25, 26, 35, 0, 33, 0, 0,
	37, 0, 38, 39, 36, 29, 30, 31, 0, 0,
	0, 0, 23, 27, 86, 22, 28, 34, 0, 32,
	0, 24, 25, 26, 35, 0, 33, 0, 0, 37,
	0, 38, 39, 36, 29, 30, 31, 0, 0, 0,
	0, 23, 27, 0, 0, 0, 34, 0, 0, 0,
	24, 25, 26, 35, 0, 33, 0, 0, 37, 0,
	38, 39, 0, 0, 64, 65, 66, 67, 68, 69,
	23, 27, 0, 0, 55, 85, 0, 0, 0, 24,
	25, 26, 35, 78, 33, 0, 64, 65, 66, 67,
	68, 69, 0, 52, 53, 54, 55, 0, 0, 0,
	49, 0, 0, 74, 76, 78, 77, 0, 72, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 49, 0, 0, 74, 76, 0, 77, 0,
	72,
}
var yyPact = [...]int{

	-45, -1000, 1885, -45, -45, -1000, -1000, -1000, -1000, 260,
	1407, 166, -1000, -1000, 1969, 1969, 259, 212, 2170, 138,
	1969, -59, -1000, 1969, 1969, 1969, 2141, 2112, -1000, -1000,
	-1000, -1000, 104, -45, -45, 1969, 23, 2083, 16, 1969,
	154, 1969, -1000, 1346, -1000, 152, -1000, 1969, 1969, 258,
	1969, 1969, 1969, 1969, 1969, 1969, 1969, 1969, 1969, 1969,
	1969, 1969, 1969, 1969, 1969, 1969, 1969, 1969, 1969, 1969,
	-1000, -1000, 1969, 1969, 1969, 1969, 1969, 1969, 1969, 1969,
	153, 1467, 1467, 137, 157, -45, 15, 56, 1273, 159,
	-45, 1213, 1969, 1969, 2214, 2214, 2214, -59, 1647, -59,
	1587, 252, -58, 1969, 231, 1153, 225, -31, 1527, 224,
	1467, -45, 1093, -1000, 1969, -45, 1467, 1033, -1000, 2192,
	2192, 2214, 2214, 2214, 1467, 70, 70, 1928, 1928, 70,
	70, 70, 70, 1467, 1467, 1467, 1467, 1467, 1467, 1467,
	1694, 1467, 1741, 91, 433, 1467, -1000, 1467, -45, -45,
	1969, -45, 113, 1814, 1969, 1969, -45, 1969, 112, -45,
	84, 373, 250, 245, 27, 228, 243, -38, -44, -1000,
	161, -1000, 17, -1000, 136, 1969, 1969, -23, 225, 225,
	2054, -45, -1000, 240, 1969, 11, -1000, -1000, 1969, 2001,
	111, 973, 109, -1000, 161, 913, 853, 107, -1000, 195,
	106, 174, 4, -1000, -1000, 1969, -1000, -1000, 135, -26,
	-39, 222, -45, -67, -45, 105, 1969, 237, -1000, -45,
	64, 63, -1000, 25, 66, 1467, -59, 100, -1000, 1467,
	-1000, 793, 1467, -59, -1000, -45, -1000, -45, 1969, -1000,
	146, -1000, -1000, -1000, 1969, 131, -1000, -1000, -1000, 733,
	-45, 134, 130, -27, 209, -1000, 122, -1000, 1467, -1000,
	236, -36, -1000, -64, -1000, -1000, 1969, 1969, -1000, -1000,
	99, 97, 673, 128, -45, 613, -45, -1000, 95, -45,
	-45, 126, -1000, -1000, -1000, -45, -1000, 225, -1000, -1000,
	553, 313, -1000, -1000, -45, -45, 93, -45, -45, -1000,
	83, 82, -45, 121, 50, -1000, -1000, 1969, 79, 77,
	193, -45, -1000, -1000, -1000, 73, -1000, -1000, 493, -1000,
	192, 117, -1000, -1000, -1000, 115, -45, -45, 69, 61,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 10, 286, 227, 277, 5, 2, 275, 4, 0,
	20, 11, 272, 1, 266, 7, 6, 265, 82, 289,
	238,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 7, 7,
	7, 7, 7, 6, 5, 13, 14, 14, 14, 15,
	15, 15, 12, 11, 11, 11, 8, 8, 8, 10,
	10, 10, 10, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 17, 17, 17, 16, 18, 18, 19,
	19, 20, 20,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 3, 3, 1,
	1, 2, 2, 5, 1, 4, 7, 9, 5, 13,
	12, 9, 8, 5, 1, 7, 5, 5, 0, 2,
	2, 2, 2, 5, 4, 3, 0, 1, 4, 0,
	1, 4, 3, 1, 4, 4, 1, 3, 6, 0,
	1, 4, 4, 1, 1, 2, 2, 2, 2, 4,
	2, 4, 1, 1, 1, 1, 5, 3, 7, 8,
	8, 9, 5, 6, 5, 6, 3, 4, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	3, 3, 3, 3, 5, 4, 6, 5, 5, 4,
	6, 5, 4, 4, 6, 6, 4, 5, 7, 7,
	9, 3, 2, 0, 1, 3, 2, 0, 1, 1,
	2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -18, -2, -19, -20, 67, 77, -3, 11,
	-9, -11, 38, 39, 10, 12, 27, -4, 15, 29,
	45, 4, 5, 60, 69, 70, 71, 61, 6, 24,
	25, 26, 9, 74, 65, 72, 23, 48, 50, 51,
	-10, 13, -18, -19, -20, -15, 4, 53, 54, 68,
	59, 60, 61, 62, 63, 42, 43, 44, 17, 18,
	57, 19, 58, 20, 32, 33, 34, 35, 36, 37,
	40, 41, 76, 21, 71, 22, 72, 74, 51, 53,
	-10, -9, -9, 4, 14, 65, 4, -12, -9, -11,
	65, -9, 72, 74, -9, -9, -9, 4, -9, 4,
	-9, 72, 4, -18, -18, -9, 72, 4, -9, 72,
	-9, 56, -9, -3, 53, 56, -9, -9, 4, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -10, -9, -9, -11, -9, 56, 65,
	13, 65, -1, -18, 16, 67, 65, 53, -1, 65,
	-10, -9, 68, 68, -15, 4, 72, -10, -14, -13,
	6, 73, -8, 4, 28, 72, 72, -8, 49, 52,
	-18, 65, -11, -18, 55, 8, 73, 75, 55, -18,
	-1, -9, -1, 66, 6, -9, -9, -1, -11, 66,
	-7, -18, 8, 73, 75, 55, 4, 4, 73, 8,
	-15, 4, 56, -18, 56, -18, 55, 68, 73, 65,
	-10, -10, 73, -8, -8, -9, 4, -1, 4, -9,
	73, -9, -9, 4, 66, 65, 66, 65, 67, 66,
	30, 66, -6, -5, 46, 47, -6, -5, 73, -9,
	65, 73, 73, 8, -18, 75, -18, 66, -9, 4,
	-18, 8, 73, 8, 73, 73, 56, 56, 66, 75,
	-1, -1, -9, 4, 65, -9, 55, 75, -1, 65,
	65, 73, 75, -13, 66, -17, -16, 4, 73, 73,
	-9, -9, 66, 66, 65, 65, -1, 55, -18, 66,
	-1, -1, 65, -18, -8, 73, 73, 56, -1, -1,
	66, -18, -1, 66, 66, -1, 66, -16, -9, 66,
	66, 31, -1, 66, 73, 31, 65, 65, -1, -1,
	66, 66,
}
var yyDef = [...]int{

	-2, -2, -2, 127, 128, 129, 131, 132, 4, 39,
	-2, 0, 9, 10, 49, 0, 0, 14, 49, 0,
	0, 53, 54, 0, 0, 0, 0, 0, 62, 63,
	64, 65, 0, 127, 127, 0, 0, 0, 0, 0,
	0, 0, 2, -2, 130, 0, 40, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	98, 99, 0, 0, 0, 0, 49, 0, 0, 49,
	11, 50, 12, 0, 0, -2, 53, 0, -2, 0,
	-2, 0, 49, 0, 55, 56, 57, -2, 0, -2,
	0, 39, 0, 49, 36, 0, 0, 53, 0, 0,
	122, 127, 0, 5, 49, 127, 7, 0, 67, 78,
	79, 80, 81, 82, 83, 84, 85, -2, -2, 88,
	89, 90, 91, 92, 93, 94, 95, 96, 97, 100,
	101, 102, 103, 0, 0, 121, 8, -2, 127, -2,
	0, -2, 0, -2, 0, 0, -2, 49, 0, 28,
	0, 0, 0, 0, 0, 40, 39, 127, 127, 37,
	0, 76, 0, 46, 0, 49, 49, 0, 0, 0,
	0, -2, 6, 0, 0, 0, 109, 113, 0, 0,
	0, 0, 0, 15, 62, 0, 0, 0, 42, 0,
	0, 0, 0, 105, 112, 0, 59, 61, 0, 0,
	0, 40, 127, 0, 127, 0, 0, 0, 77, 127,
	0, 0, 116, 0, 0, -2, -2, 0, 41, 66,
	108, 0, 51, -2, 13, -2, 26, -2, 0, 18,
	0, 23, 31, 32, 0, 0, 29, 30, 104, 0,
	-2, 0, 0, 0, 0, 72, 0, 74, 35, 47,
	123, 0, -2, 0, -2, 117, 0, 0, 27, 115,
	0, 0, 0, 0, -2, 0, 127, 114, 0, -2,
	-2, 0, 73, 38, 75, 127, 124, 0, -2, -2,
	0, 0, 25, 16, -2, -2, 0, 127, -2, 68,
	0, 0, -2, 0, 126, 118, 119, 0, 0, 0,
	22, -2, 34, 69, 70, 0, 48, 125, 0, 17,
	21, 0, 33, 71, 120, 0, -2, -2, 0, 0,
	20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	77, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 69, 3, 3, 3, 63, 71, 3,
	72, 73, 61, 59, 56, 60, 68, 62, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 55, 67,
	58, 53, 57, 54, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 74, 3, 75, 70, 3, 3, 3, 3, 3,
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
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:78
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:85
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:92
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
		//line parser.go.y:103
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:108
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:116
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:121
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:126
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:141
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:146
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:151
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyDollar[2].tok.Lit, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:156
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:166
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:171
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:176
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:186
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:191
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:199
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:204
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
		//line parser.go.y:213
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:219
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:223
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:227
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:231
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:235
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
		//line parser.go.y:246
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:252
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:258
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:263
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:267
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:276
		{
			yyVAL.expr_idents = []string{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:280
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:284
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:289
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:295
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:299
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:303
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:308
		{
			yyVAL.typ = ast.Type{Name: yyDollar[1].tok.Lit}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:312
		{
			yyVAL.typ = ast.Type{Name: yyDollar[1].typ.Name + "." + yyDollar[3].tok.Lit}
		}
	case 48:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:316
		{
			yyVAL.typ = ast.Type{Fields: yyDollar[4].struct_fields}
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:321
		{
			yyVAL.exprs = nil
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:325
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:329
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:333
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:354
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:364
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:374
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:384
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:389
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:394
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:404
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:409
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 68:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:414
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:419
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:424
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 71:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:429
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:434
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 73:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:439
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:444
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
		//line parser.go.y:453
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
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:487
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:492
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:497
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:502
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:507
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:512
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:517
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:522
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:527
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:532
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:537
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:542
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:547
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:552
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:557
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:562
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:567
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:572
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:577
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:582
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:587
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:592
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:597
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:602
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:607
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 106:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:612
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:617
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:622
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:627
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:632
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:637
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:642
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:647
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:652
		{
			yyVAL.expr = &ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:657
		{
			yyVAL.expr = &ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:662
		{
			yyVAL.expr = &ast.MakeExpr{Type: yyDollar[3].typ.Name}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:667
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].typ.Name, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:672
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].typ.Name, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 119:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:677
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 120:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:682
		{
			yyVAL.expr = &ast.MakeArrayExpr{Type: yyDollar[4].typ.Name, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:687
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:692
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 123:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:698
		{
			yyVAL.struct_fields = nil
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:702
		{
			yyVAL.struct_fields = []ast.StructField{yyDollar[1].struct_field}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:706
		{
			yyVAL.struct_fields = append(yyDollar[1].struct_fields, yyDollar[3].struct_field)
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:711
		{
			yyVAL.struct_field = ast.StructField{
				Name: yyDollar[1].tok.Lit,
				Type: yyDollar[2].typ.Name,
			}
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:724
		{
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:727
		{
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:732
		{
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:735
		{
		}
	}
	goto yystack /* stack new state and value */
}
