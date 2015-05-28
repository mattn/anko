//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:25
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
const UNARY = 57389

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
	"'!'",
	"'^'",
	"'&'",
	"'.'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'|'",
	"'\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:609

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	50, 46,
	-2, 2,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 5,
	50, 47,
	-2, 24,
	-1, 7,
	1, 8,
	45, 8,
	46, 8,
	60, 8,
	61, 8,
	71, 8,
	-2, 50,
	-1, 35,
	50, 46,
	-2, 104,
	-1, 82,
	50, 36,
	-2, 2,
	-1, 83,
	50, 47,
	-2, 43,
	-1, 86,
	50, 46,
	-2, 2,
	-1, 91,
	1, 55,
	45, 55,
	46, 55,
	47, 55,
	49, 55,
	50, 55,
	59, 55,
	60, 55,
	61, 55,
	67, 55,
	69, 55,
	71, 55,
	-2, 50,
	-1, 93,
	1, 57,
	45, 57,
	46, 57,
	47, 57,
	49, 57,
	50, 57,
	59, 57,
	60, 57,
	61, 57,
	67, 57,
	69, 57,
	71, 57,
	-2, 50,
	-1, 120,
	17, 0,
	18, 0,
	-2, 81,
	-1, 121,
	17, 0,
	18, 0,
	-2, 82,
	-1, 139,
	50, 47,
	-2, 43,
	-1, 143,
	50, 46,
	-2, 2,
	-1, 145,
	50, 46,
	-2, 2,
	-1, 149,
	50, 46,
	-2, 2,
	-1, 165,
	50, 48,
	-2, 44,
	-1, 166,
	1, 45,
	45, 45,
	46, 45,
	47, 45,
	50, 49,
	60, 45,
	61, 45,
	71, 45,
	-2, 50,
	-1, 167,
	50, 46,
	-2, 2,
	-1, 176,
	1, 49,
	45, 49,
	46, 49,
	50, 49,
	60, 49,
	61, 49,
	67, 49,
	69, 49,
	71, 49,
	-2, 50,
	-1, 201,
	50, 46,
	-2, 2,
	-1, 203,
	50, 46,
	-2, 2,
	-1, 214,
	50, 46,
	-2, 2,
	-1, 224,
	50, 46,
	-2, 2,
	-1, 228,
	50, 46,
	-2, 2,
	-1, 229,
	50, 46,
	-2, 2,
	-1, 233,
	50, 46,
	-2, 2,
	-1, 234,
	50, 46,
	-2, 2,
	-1, 237,
	50, 46,
	-2, 2,
	-1, 241,
	50, 46,
	-2, 2,
	-1, 245,
	50, 46,
	-2, 2,
	-1, 256,
	50, 46,
	-2, 2,
	-1, 257,
	50, 46,
	-2, 2,
}

const yyNprod = 109
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1503

var yyAct = [...]int{

	1, 36, 5, 208, 34, 209, 37, 39, 99, 142,
	73, 6, 74, 76, 78, 142, 38, 83, 108, 87,
	230, 88, 89, 90, 92, 94, 85, 215, 159, 146,
	164, 76, 218, 101, 150, 216, 104, 106, 158, 96,
	261, 260, 254, 251, 109, 110, 250, 112, 113, 114,
	115, 116, 117, 118, 119, 120, 121, 122, 123, 124,
	125, 126, 127, 128, 129, 130, 131, 32, 108, 132,
	133, 134, 135, 136, 76, 139, 76, 141, 75, 73,
	248, 74, 247, 147, 138, 191, 244, 152, 142, 58,
	59, 60, 61, 62, 63, 257, 97, 64, 65, 49,
	142, 95, 238, 156, 232, 173, 165, 231, 219, 204,
	76, 44, 45, 46, 47, 48, 223, 172, 210, 211,
	160, 202, 68, 43, 71, 200, 70, 185, 66, 181,
	161, 144, 142, 207, 256, 241, 234, 229, 228, 137,
	214, 140, 143, 86, 177, 175, 179, 178, 107, 180,
	182, 108, 103, 183, 139, 226, 162, 151, 187, 77,
	17, 23, 72, 184, 27, 196, 193, 76, 198, 195,
	3, 224, 255, 199, 252, 168, 206, 145, 31, 24,
	25, 26, 210, 211, 80, 217, 192, 100, 194, 190,
	106, 212, 189, 213, 169, 157, 111, 102, 79, 40,
	188, 98, 220, 35, 221, 84, 105, 186, 222, 18,
	22, 13, 2, 225, 29, 227, 0, 19, 20, 21,
	0, 30, 0, 28, 0, 235, 0, 0, 0, 239,
	240, 237, 197, 0, 242, 243, 0, 0, 246, 0,
	0, 245, 249, 0, 0, 0, 253, 52, 53, 55,
	57, 67, 69, 0, 0, 0, 0, 258, 259, 0,
	0, 58, 59, 60, 61, 62, 63, 0, 0, 64,
	65, 49, 50, 51, 0, 0, 0, 0, 42, 236,
	0, 54, 56, 44, 45, 46, 47, 48, 52, 53,
	55, 57, 67, 69, 68, 43, 71, 0, 70, 0,
	66, 0, 58, 59, 60, 61, 62, 63, 0, 0,
	64, 65, 49, 50, 51, 0, 0, 0, 0, 42,
	0, 0, 54, 56, 44, 45, 46, 47, 48, 0,
	233, 0, 0, 0, 0, 68, 43, 71, 0, 70,
	0, 66, 52, 53, 55, 57, 67, 69, 0, 0,
	0, 0, 0, 0, 0, 0, 58, 59, 60, 61,
	62, 63, 0, 0, 64, 65, 49, 50, 51, 0,
	0, 0, 0, 42, 0, 0, 54, 56, 44, 45,
	46, 47, 48, 0, 0, 0, 205, 0, 0, 68,
	43, 71, 0, 70, 0, 66, 52, 53, 55, 57,
	67, 69, 0, 0, 0, 0, 0, 0, 0, 0,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 42, 0, 0,
	54, 56, 44, 45, 46, 47, 48, 0, 203, 0,
	0, 0, 0, 68, 43, 71, 0, 70, 0, 66,
	52, 53, 55, 57, 67, 69, 0, 0, 0, 0,
	0, 0, 0, 0, 58, 59, 60, 61, 62, 63,
	0, 0, 64, 65, 49, 50, 51, 0, 0, 0,
	0, 42, 0, 0, 54, 56, 44, 45, 46, 47,
	48, 0, 201, 0, 0, 0, 0, 68, 43, 71,
	0, 70, 0, 66, 52, 53, 55, 57, 67, 69,
	0, 0, 0, 0, 0, 0, 0, 0, 58, 59,
	60, 61, 62, 63, 0, 0, 64, 65, 49, 50,
	51, 0, 0, 0, 0, 42, 0, 0, 54, 56,
	44, 45, 46, 47, 48, 52, 53, 55, 57, 67,
	69, 68, 43, 71, 0, 70, 174, 66, 0, 58,
	59, 60, 61, 62, 63, 0, 0, 64, 65, 49,
	50, 51, 0, 0, 0, 0, 42, 0, 0, 54,
	56, 44, 45, 46, 47, 48, 52, 53, 55, 57,
	67, 69, 68, 43, 71, 0, 70, 171, 66, 0,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 42, 170, 0,
	54, 56, 44, 45, 46, 47, 48, 52, 53, 55,
	57, 67, 69, 68, 43, 71, 0, 70, 0, 66,
	0, 58, 59, 60, 61, 62, 63, 0, 0, 64,
	65, 49, 50, 51, 0, 0, 0, 0, 42, 0,
	0, 54, 56, 44, 45, 46, 47, 48, 0, 167,
	0, 0, 0, 0, 68, 43, 71, 0, 70, 0,
	66, 52, 53, 55, 57, 67, 69, 0, 0, 0,
	0, 0, 0, 0, 0, 58, 59, 60, 61, 62,
	63, 0, 0, 64, 65, 49, 50, 51, 0, 0,
	0, 0, 42, 0, 0, 54, 56, 44, 45, 46,
	47, 48, 52, 53, 55, 57, 67, 69, 68, 43,
	71, 163, 70, 0, 66, 0, 58, 59, 60, 61,
	62, 63, 0, 0, 64, 65, 49, 50, 51, 0,
	0, 0, 0, 42, 0, 0, 54, 56, 44, 45,
	46, 47, 48, 0, 153, 0, 0, 0, 0, 68,
	43, 71, 0, 70, 0, 66, 52, 53, 55, 57,
	67, 69, 0, 0, 0, 0, 0, 0, 0, 0,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 42, 0, 0,
	54, 56, 44, 45, 46, 47, 48, 0, 149, 0,
	0, 0, 0, 68, 43, 71, 0, 70, 0, 66,
	52, 53, 55, 57, 67, 69, 0, 0, 0, 0,
	0, 0, 0, 0, 58, 59, 60, 61, 62, 63,
	0, 0, 64, 65, 49, 50, 51, 0, 0, 0,
	41, 42, 0, 0, 54, 56, 44, 45, 46, 47,
	48, 52, 53, 55, 57, 67, 69, 68, 43, 71,
	0, 70, 0, 66, 0, 58, 59, 60, 61, 62,
	63, 0, 0, 64, 65, 49, 50, 51, 0, 0,
	0, 0, 42, 0, 0, 54, 56, 44, 45, 46,
	47, 48, 52, 53, 55, 57, 67, 69, 68, 43,
	71, 0, 70, 0, 66, 0, 58, 59, 60, 61,
	62, 63, 0, 0, 64, 65, 49, 50, 51, 0,
	0, 0, 0, 42, 0, 0, 54, 56, 44, 45,
	46, 47, 48, 52, 53, 55, 57, 67, 69, 68,
	155, 71, 0, 70, 0, 66, 0, 58, 59, 60,
	61, 62, 63, 0, 0, 64, 65, 49, 50, 51,
	0, 0, 0, 0, 42, 0, 0, 54, 56, 44,
	45, 46, 47, 48, 0, 0, 0, 0, 0, 0,
	68, 154, 71, 0, 70, 0, 66, 7, 17, 23,
	0, 0, 27, 10, 4, 11, 33, 0, 14, 0,
	0, 0, 0, 0, 0, 0, 31, 24, 25, 26,
	12, 15, 0, 0, 0, 0, 0, 0, 0, 0,
	8, 9, 0, 0, 0, 0, 0, 16, 0, 58,
	59, 60, 61, 62, 63, 0, 0, 18, 22, 49,
	0, 0, 29, 0, 37, 19, 20, 21, 0, 30,
	0, 28, 0, 0, 38, 52, 53, 55, 57, 0,
	69, 0, 68, 43, 71, 0, 70, 0, 66, 58,
	59, 60, 61, 62, 63, 0, 0, 64, 65, 49,
	50, 51, 0, 0, 0, 0, 0, 0, 0, 54,
	56, 44, 45, 46, 47, 48, 52, 53, 55, 57,
	0, 0, 68, 43, 71, 0, 70, 0, 66, 0,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 0, 0, 0,
	54, 56, 44, 45, 46, 47, 48, 0, 55, 57,
	0, 0, 0, 68, 43, 71, 0, 70, 0, 66,
	58, 59, 60, 61, 62, 63, 0, 0, 64, 65,
	49, 50, 51, 0, 0, 0, 0, 0, 0, 0,
	54, 56, 44, 45, 46, 47, 48, 0, 0, 0,
	0, 0, 0, 68, 43, 71, 0, 70, 0, 66,
	7, 17, 23, 0, 0, 27, 10, 4, 11, 33,
	0, 14, 0, 0, 0, 0, 0, 0, 0, 31,
	24, 25, 26, 12, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 8, 9, 0, 0, 0, 0, 0,
	16, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	18, 22, 0, 0, 0, 29, 0, 0, 19, 20,
	21, 0, 30, 0, 28, 7, 17, 148, 0, 0,
	27, 10, 4, 11, 33, 0, 14, 0, 0, 0,
	0, 0, 0, 0, 31, 24, 25, 26, 12, 15,
	0, 0, 0, 0, 0, 0, 0, 0, 8, 9,
	0, 0, 0, 0, 0, 16, 0, 58, 59, 60,
	61, 62, 63, 0, 0, 18, 22, 49, 0, 0,
	29, 0, 0, 19, 20, 21, 0, 30, 0, 28,
	0, 46, 47, 48, 176, 17, 23, 0, 0, 27,
	68, 43, 71, 0, 70, 0, 66, 0, 0, 0,
	0, 0, 0, 31, 24, 25, 26, 166, 17, 23,
	0, 0, 27, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 31, 24, 25, 26,
	0, 93, 17, 23, 18, 22, 27, 0, 0, 29,
	0, 0, 19, 20, 21, 0, 30, 0, 28, 0,
	31, 24, 25, 26, 91, 17, 23, 18, 22, 27,
	0, 0, 29, 0, 0, 19, 20, 21, 0, 30,
	0, 28, 0, 31, 24, 25, 26, 0, 81, 17,
	23, 18, 22, 27, 0, 0, 29, 0, 0, 19,
	20, 21, 0, 30, 0, 28, 0, 31, 24, 25,
	26, 0, 0, 0, 18, 22, 0, 0, 0, 29,
	0, 0, 19, 20, 21, 0, 30, 0, 28, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 18, 22,
	0, 0, 0, 82, 0, 0, 19, 20, 21, 0,
	30, 0, 28,
}
var yyPact = [...]int{

	1206, -1000, -55, -1000, 195, 813, 115, -56, -1000, -1000,
	155, 155, 194, 170, 1434, 84, 155, -1000, 155, 155,
	155, 1410, 1387, -1000, -1000, -1000, -1000, 35, 155, 181,
	155, 193, 102, 155, -1000, 1003, -1000, -1000, -1000, 101,
	-1000, 155, 155, 192, 155, 155, 155, 155, 155, 155,
	155, 155, 155, 155, 155, 155, 155, 155, 155, 155,
	155, 155, 155, 155, -1000, -1000, 155, 155, 155, 155,
	155, 155, 155, 155, 155, 82, 854, -56, 854, 83,
	118, 13, 1271, 759, -27, 110, 1206, 705, 1018, 1018,
	1018, -56, 936, -56, 895, 191, -28, -41, 70, -1000,
	107, 664, -36, 1363, 610, -1000, -1000, 155, 190, 854,
	569, -1000, 1286, 1286, 1018, 1018, 1018, 854, 58, 58,
	1139, 1139, 58, 58, 58, 58, 854, 854, 854, 854,
	854, 854, 854, 1058, 854, 1099, 528, 50, -1000, 854,
	38, 487, 1340, 1206, 155, 1206, 155, 69, 107, 1206,
	155, 155, 67, -55, 188, 185, 18, 178, 184, -1000,
	181, -1000, 155, -1000, 155, 854, -56, 1206, 82, -1000,
	155, -1000, -1000, -1000, -1000, 854, -56, 65, 433, 61,
	379, -1000, 49, 325, -1000, 147, 73, 137, -55, -1000,
	-1000, 81, -40, -32, 177, -1000, 854, -35, 48, 854,
	-1000, 1206, -1000, 1206, -1000, 155, 112, -1000, -1000, -1000,
	155, 106, -1000, -1000, 1206, 79, 78, -47, -1000, -1000,
	47, 44, 271, 77, 1206, 230, -55, 42, 1206, 1206,
	76, -1000, -1000, 1206, 1206, 26, -55, 1206, -1000, 22,
	20, 1206, -14, -17, 144, 1206, -1000, -1000, -1000, -18,
	-1000, 142, 75, -1000, -1000, 36, 1206, 1206, -19, -20,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 0, 212, 170, 211, 5, 3, 207, 2, 67,
	11, 205, 8, 201, 7, 4, 200, 1,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 7, 7,
	7, 7, 7, 6, 5, 12, 13, 13, 13, 14,
	14, 14, 11, 10, 10, 10, 9, 9, 9, 9,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 15, 15, 16, 16, 17, 17,
}
var yyR2 = [...]int{

	0, 2, 0, 1, 3, 4, 3, 3, 1, 1,
	1, 2, 2, 5, 1, 7, 4, 5, 9, 13,
	12, 9, 8, 5, 1, 7, 5, 5, 0, 2,
	2, 2, 2, 5, 4, 3, 0, 1, 3, 0,
	1, 3, 3, 1, 3, 3, 0, 1, 3, 3,
	1, 1, 2, 2, 2, 2, 4, 2, 4, 1,
	1, 1, 1, 5, 3, 7, 8, 8, 9, 3,
	3, 3, 5, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 3, 3, 3, 3, 4,
	4, 4, 4, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 11, -8, -10, 4, 37, 38,
	10, 12, 27, -4, 15, 28, 44, 5, 54, 62,
	63, 64, 55, 6, 24, 25, 26, 9, 68, 59,
	66, 23, -9, 13, -15, -16, -17, 61, 71, -14,
	4, 47, 48, 65, 53, 54, 55, 56, 57, 41,
	42, 43, 17, 18, 51, 19, 52, 20, 31, 32,
	33, 34, 35, 36, 39, 40, 70, 21, 64, 22,
	68, 66, 47, 66, 68, -9, -8, 4, -8, 4,
	14, 4, 59, -8, -11, -10, 59, -8, -8, -8,
	-8, 4, -8, 4, -8, 66, 4, -9, -13, -12,
	6, -8, 4, 50, -8, -3, -17, 47, 50, -8,
	-8, 4, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -9, -10, -8,
	-9, -8, 50, 59, 13, 59, 16, -1, 6, 59,
	61, 47, -1, 59, 65, 65, -14, 4, 66, 69,
	50, 60, 49, 67, 66, -8, 4, 59, -9, 4,
	49, 69, 67, 67, 69, -8, 4, -1, -8, -1,
	-8, 60, -1, -8, -10, 60, -7, -15, -16, 4,
	4, 67, 8, -14, 4, -12, -8, -9, -1, -8,
	60, 59, 60, 59, 60, 61, 29, 60, -6, -5,
	45, 46, -6, -5, 59, 67, 67, 8, 67, 60,
	-1, -1, -8, 4, 59, -8, 49, -1, 59, 59,
	67, 60, 60, 59, 59, -1, 49, -15, 60, -1,
	-1, 59, -1, -1, 60, -15, -1, 60, 60, -1,
	60, 60, 30, -1, 60, 30, 59, 59, -1, -1,
	60, 60,
}
var yyDef = [...]int{

	-2, -2, 103, 3, 39, -2, 0, -2, 9, 10,
	46, 0, 0, 14, 46, 0, 0, 51, 0, 0,
	0, 0, 0, 59, 60, 61, 62, 0, 46, 36,
	0, 0, 0, 0, 1, -2, 105, 107, 108, 0,
	40, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 93, 94, 0, 0, 0, 0,
	0, 46, 46, 46, 0, 11, 47, 50, 12, 0,
	0, 50, -2, -2, 0, 0, -2, 0, 52, 53,
	54, -2, 0, -2, 0, 39, 0, 0, 0, 37,
	0, 0, 0, 0, 0, 4, 106, 46, 0, 6,
	0, 64, 73, 74, 75, 76, 77, 78, 79, 80,
	-2, -2, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 95, 96, 97, 98, 0, 0, 7, -2,
	0, 0, 0, -2, 0, -2, 0, 0, 59, -2,
	0, 46, 0, 28, 0, 0, 0, 40, 39, 69,
	0, 70, 0, 71, 46, -2, -2, -2, 5, 41,
	0, 101, 102, 99, 100, 48, -2, 0, 0, 0,
	0, 16, 0, 0, 42, 0, 0, 0, 104, 56,
	58, 0, 0, 0, 40, 38, 35, 0, 0, 63,
	13, -2, 26, -2, 17, 0, 0, 23, 31, 32,
	0, 0, 29, 30, -2, 0, 0, 0, 72, 27,
	0, 0, 0, 0, -2, 0, 103, 0, -2, -2,
	0, 25, 15, -2, -2, 0, 103, -2, 65, 0,
	0, -2, 0, 0, 22, -2, 34, 66, 67, 0,
	18, 21, 0, 33, 68, 0, -2, -2, 0, 0,
	20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	71, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 62, 3, 3, 3, 57, 64, 3,
	66, 67, 55, 53, 50, 54, 65, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 61,
	52, 47, 51, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 68, 3, 69, 63, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 59, 70, 60,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 58,
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
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
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
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
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
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
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
			yychar = -1
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:63
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:68
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:75
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[1].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:82
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
				if l, ok := yylex.(*Lexer); ok {
					l.stmts = yyVAL.stmts
				}
			}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:93
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:98
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:106
		{
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:109
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:114
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:119
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:124
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:129
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:134
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.stmt = &ast.ForStmt{Var: yyDollar[2].tok.Lit, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:149
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:154
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:159
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:164
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:169
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:174
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:179
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:184
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:192
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:197
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
		//line parser.go.y:206
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:212
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:216
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:220
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:228
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
		//line parser.go.y:239
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:245
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:251
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:260
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:264
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[3].expr_pair)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:269
		{
			yyVAL.expr_idents = []string{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:273
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:277
		{
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[3].tok.Lit)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:282
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:288
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:292
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:296
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[3].tok.Lit})
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:301
		{
			yyVAL.exprs = nil
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:305
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:313
		{
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[3].tok.Lit})
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:324
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:329
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:334
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:354
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:364
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:374
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:384
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:389
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 65:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:394
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:399
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyDollar[3].tok.Lit}, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 67:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:404
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 68:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:409
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Args: []string{yyDollar[4].tok.Lit}, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:414
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[2].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:419
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[2].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:428
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:433
		{
			yyVAL.expr = &ast.NewExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:438
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:443
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:448
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:453
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:458
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:463
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:468
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:473
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:478
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:483
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:488
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:493
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:498
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:503
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:508
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:513
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:518
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:523
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:528
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:533
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:538
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:543
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:548
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:553
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:558
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:563
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:568
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:573
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 101:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:578
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:583
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:594
		{
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:597
		{
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:602
		{
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:605
		{
		}
	}
	goto yystack /* stack new state and value */
}
