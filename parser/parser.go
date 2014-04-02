//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:23
type yySymType struct {
	yys                    int
	stmt_var               ast.Stmt
	stmt_func              ast.Stmt
	stmt_if                ast.Stmt
	stmt_for               ast.Stmt
	stmt_try_catch_finally ast.Stmt
	stmts                  []ast.Stmt
	stmt                   ast.Stmt
	teim                   ast.Expr
	expr                   ast.Expr
	tok                    Token
	idents                 []string
	exprs                  []ast.Expr
	pair                   ast.Expr
	pairs                  []ast.Expr
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
const UNARY = 57377

var yyToknames = []string{
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
	" =",
	" ?",
	" :",
	" >",
	" <",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:399

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 29,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	35, 29,
	51, 29,
	-2, 1,
	-1, 3,
	35, 29,
	51, 29,
	-2, 1,
	-1, 4,
	35, 29,
	51, 29,
	-2, 1,
	-1, 5,
	35, 29,
	51, 29,
	-2, 1,
	-1, 6,
	35, 29,
	51, 29,
	-2, 1,
	-1, 7,
	35, 29,
	51, 29,
	-2, 1,
	-1, 18,
	51, 30,
	-2, 66,
	-1, 30,
	35, 29,
	51, 29,
	-2, 1,
	-1, 69,
	48, 1,
	-2, 29,
	-1, 85,
	48, 1,
	-2, 29,
	-1, 95,
	17, 0,
	18, 0,
	-2, 50,
	-1, 96,
	17, 0,
	18, 0,
	-2, 51,
	-1, 108,
	48, 1,
	-2, 29,
	-1, 154,
	48, 1,
	-2, 29,
	-1, 156,
	48, 1,
	-2, 29,
	-1, 157,
	48, 1,
	-2, 29,
	-1, 161,
	48, 1,
	-2, 29,
	-1, 164,
	48, 1,
	-2, 29,
	-1, 168,
	48, 1,
	-2, 29,
	-1, 169,
	48, 1,
	-2, 29,
	-1, 171,
	48, 1,
	-2, 29,
	-1, 186,
	48, 1,
	-2, 29,
	-1, 189,
	48, 1,
	-2, 29,
	-1, 194,
	48, 1,
	-2, 29,
}

const yyNprod = 73
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 657

var yyAct = []int{

	1, 8, 79, 29, 31, 32, 33, 35, 36, 159,
	122, 9, 58, 61, 122, 152, 57, 123, 143, 122,
	122, 77, 182, 81, 76, 63, 83, 138, 57, 37,
	130, 84, 38, 39, 162, 51, 53, 131, 122, 77,
	155, 89, 90, 91, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 77, 65,
	125, 196, 87, 124, 72, 73, 74, 75, 71, 113,
	115, 128, 77, 117, 118, 119, 120, 121, 111, 110,
	66, 106, 70, 67, 193, 109, 129, 40, 41, 42,
	43, 44, 56, 190, 185, 116, 37, 184, 194, 38,
	39, 57, 51, 53, 66, 168, 86, 167, 57, 134,
	183, 77, 181, 179, 175, 174, 141, 172, 151, 148,
	142, 189, 136, 186, 144, 171, 169, 145, 146, 164,
	77, 161, 149, 156, 135, 150, 42, 43, 44, 154,
	85, 126, 108, 37, 69, 192, 38, 39, 187, 51,
	53, 158, 114, 147, 153, 163, 139, 165, 166, 80,
	176, 137, 170, 112, 107, 173, 88, 82, 68, 177,
	178, 64, 180, 62, 78, 45, 46, 48, 50, 52,
	54, 7, 6, 5, 60, 23, 25, 188, 4, 59,
	191, 3, 2, 0, 55, 195, 47, 49, 40, 41,
	42, 43, 44, 22, 26, 27, 28, 37, 160, 0,
	38, 39, 0, 51, 53, 45, 46, 48, 50, 52,
	54, 24, 0, 0, 0, 0, 0, 20, 0, 21,
	0, 0, 0, 19, 55, 0, 47, 49, 40, 41,
	42, 43, 44, 0, 0, 157, 0, 37, 0, 0,
	38, 39, 0, 51, 53, 45, 46, 48, 50, 52,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 0, 47, 49, 40, 41,
	42, 43, 44, 0, 0, 0, 0, 37, 140, 0,
	38, 39, 0, 51, 53, 45, 46, 48, 50, 52,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 133, 47, 49, 40, 41,
	42, 43, 44, 0, 0, 0, 0, 37, 0, 0,
	38, 39, 0, 51, 53, 45, 46, 48, 50, 52,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 0, 47, 49, 40, 41,
	42, 43, 44, 0, 0, 0, 0, 37, 0, 0,
	38, 39, 132, 51, 53, 45, 46, 48, 50, 52,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 0, 47, 49, 40, 41,
	42, 43, 44, 0, 0, 0, 0, 37, 127, 0,
	38, 39, 0, 51, 53, 45, 46, 48, 50, 52,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 0, 47, 49, 40, 41,
	42, 43, 44, 45, 46, 48, 50, 37, 54, 0,
	38, 39, 0, 51, 53, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 47, 49, 40, 41, 42, 43,
	44, 0, 0, 0, 0, 37, 0, 0, 38, 39,
	0, 51, 53, 18, 23, 25, 0, 0, 14, 10,
	13, 11, 15, 34, 16, 0, 0, 0, 0, 0,
	0, 0, 22, 26, 27, 28, 12, 17, 0, 0,
	0, 0, 0, 0, 0, 0, 18, 23, 25, 0,
	24, 14, 10, 13, 11, 15, 20, 16, 21, 0,
	0, 0, 19, 0, 0, 22, 26, 27, 28, 12,
	17, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 24, 0, 0, 0, 0, 30, 20,
	0, 21, 18, 23, 25, 19, 0, 14, 10, 13,
	11, 15, 0, 16, 0, 0, 0, 0, 0, 0,
	0, 22, 26, 27, 28, 12, 17, 0, 0, 0,
	0, 45, 46, 48, 50, 0, 0, 0, 0, 24,
	0, 0, 0, 0, 0, 20, 0, 21, 0, 0,
	0, 19, 47, 49, 40, 41, 42, 43, 44, 48,
	50, 0, 0, 37, 0, 0, 38, 39, 0, 51,
	53, 0, 0, 0, 0, 0, 0, 0, 47, 49,
	40, 41, 42, 43, 44, 0, 0, 0, 0, 37,
	0, 0, 38, 39, 0, 51, 53,
}
var yyPact = []int{

	558, -1000, 512, 558, 558, 479, 558, 558, 398, 57,
	180, 180, 169, 167, 55, 34, 164, 97, 33, 180,
	153, 180, 163, -1000, 180, -1000, -1000, -1000, -1000, -1000,
	558, -1000, -1000, -1000, 93, -1000, -1000, 180, 162, 180,
	180, 180, 180, 180, 180, 180, 180, 180, 180, 180,
	180, 180, 180, 180, 180, 180, 180, 160, 398, 31,
	33, 398, 95, 50, -1000, 30, 159, 180, 136, 558,
	180, 180, 180, 180, 180, 180, -37, 398, 12, -1000,
	104, 358, 22, -20, -1000, 558, -19, -13, -1000, 318,
	94, 94, -20, -20, -20, 600, 600, 47, 47, 47,
	47, 398, 426, 398, 574, 278, -31, -1000, 558, 180,
	157, -23, 148, 238, 180, 72, -32, 398, 398, 398,
	398, 398, 180, -1000, 153, -1000, 180, -1000, 180, 71,
	180, -1000, -1000, 180, 70, -31, -35, 146, 92, -10,
	86, 198, 122, -1000, 398, -1000, 398, -41, -1000, 158,
	398, -1000, 84, -16, 558, 82, 558, 558, 58, -1000,
	79, 558, 78, 69, 558, 67, 66, 156, 558, 558,
	65, 558, -1000, 64, -1000, -1000, -28, 62, 49, -1000,
	46, -1000, 76, 118, -1000, -1000, 558, 74, 45, 558,
	115, 36, 51, -1000, 558, 13, -1000,
}
var yyPgo = []int{

	0, 0, 192, 191, 188, 183, 182, 181, 1, 24,
	2, 174, 11,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 2, 2, 2, 3, 4, 4, 5, 5, 5,
	6, 7, 7, 7, 7, 10, 11, 11, 11, 12,
	12, 12, 9, 9, 9, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8,
}
var yyR2 = []int{

	0, 0, 2, 3, 2, 2, 2, 2, 2, 1,
	3, 2, 2, 5, 4, 8, 9, 7, 5, 9,
	7, 15, 12, 11, 8, 3, 0, 1, 3, 0,
	1, 3, 0, 1, 3, 4, 7, 3, 4, 3,
	8, 3, 4, 3, 5, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 1, 2, 1, 1,
	1, 1, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -12,
	10, 12, 27, 11, 9, 13, 15, 28, 4, 53,
	47, 49, 23, 5, 41, 6, 24, 25, 26, -1,
	46, -1, -1, -1, 14, -1, -1, 49, 52, 53,
	40, 41, 42, 43, 44, 17, 18, 38, 19, 39,
	20, 55, 21, 56, 22, 36, 35, 51, -8, 9,
	4, -8, 4, -12, 4, 4, 49, 49, 4, 47,
	49, 35, 31, 32, 33, 34, -9, -8, -11, -10,
	6, -8, 4, -8, -1, 47, 13, -9, 4, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -9, 4, 47, 35,
	49, -12, 4, -8, 16, -1, -9, -8, -8, -8,
	-8, -8, 51, 54, 51, 48, 37, 50, 49, -1,
	49, 50, 54, 37, -1, -9, -12, 4, 50, 8,
	50, -8, 48, 50, -8, -10, -8, -9, 48, -8,
	-8, 48, 50, 8, 47, 50, 47, 47, 29, 50,
	50, 47, 50, -1, 47, -1, -1, 49, 47, 47,
	-1, 47, 48, -1, 48, 48, 4, -1, -1, 48,
	-1, 48, 50, 48, 48, 48, 47, 30, -1, 47,
	48, -1, 30, 48, 47, -1, 48,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, -2, 9, 0,
	0, 0, 0, 29, 0, 0, 0, 0, -2, 32,
	26, 0, 0, 65, 0, 68, 69, 70, 71, 2,
	-2, 4, 5, 6, 0, 7, 8, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 32, 0, 11, 0,
	66, 12, 0, 0, 30, 0, 29, 0, 0, -2,
	32, 0, 0, 0, 0, 0, 0, 33, 0, 27,
	0, 0, 0, 67, 3, -2, 0, 0, 37, 0,
	45, 46, 47, 48, 49, -2, -2, 52, 53, 54,
	55, 61, 62, 63, 64, 0, 10, 31, -2, 32,
	29, 0, 30, 0, 0, 0, 0, 56, 57, 58,
	59, 60, 0, 39, 0, 41, 0, 43, 32, 0,
	0, 35, 42, 0, 0, 14, 0, 30, 0, 0,
	0, 0, 0, 38, 34, 28, 25, 0, 18, 0,
	72, 13, 0, 0, -2, 0, -2, -2, 0, 44,
	0, -2, 0, 0, -2, 0, 0, 0, -2, -2,
	0, -2, 36, 0, 17, 20, 0, 0, 0, 15,
	0, 40, 0, 24, 19, 16, -2, 0, 0, -2,
	23, 0, 0, 22, -2, 0, 21,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 44, 56, 3,
	49, 50, 42, 40, 51, 41, 52, 43, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 37, 46,
	39, 35, 38, 36, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 53, 3, 54, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 47, 55, 48,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 45,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
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
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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
		//line parser.go.y:56
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:63
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:70
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-2].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:77
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_var}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:84
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:91
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:98
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 8:
		//line parser.go.y:105
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 9:
		//line parser.go.y:113
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 10:
		//line parser.go.y:120
		{
			yyVAL.stmt = &ast.LetStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.stmt.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:127
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 12:
		//line parser.go.y:134
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 13:
		//line parser.go.y:141
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:146
		{
			yyVAL.stmt_var = &ast.LetStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
		}
	case 15:
		//line parser.go.y:151
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 16:
		//line parser.go.y:155
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 17:
		//line parser.go.y:160
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:164
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 19:
		//line parser.go.y:172
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 20:
		//line parser.go.y:177
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:182
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:186
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 23:
		//line parser.go.y:190
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 24:
		//line parser.go.y:194
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 25:
		//line parser.go.y:199
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:204
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 27:
		//line parser.go.y:208
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 28:
		//line parser.go.y:212
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 29:
		//line parser.go.y:217
		{
			yyVAL.idents = []string{}
		}
	case 30:
		//line parser.go.y:221
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 31:
		//line parser.go.y:225
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 32:
		//line parser.go.y:230
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 33:
		//line parser.go.y:234
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 34:
		//line parser.go.y:238
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 35:
		//line parser.go.y:243
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	case 36:
		//line parser.go.y:247
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 37:
		//line parser.go.y:251
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 38:
		//line parser.go.y:255
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 39:
		//line parser.go.y:259
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 40:
		//line parser.go.y:263
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 41:
		//line parser.go.y:267
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 42:
		//line parser.go.y:275
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 43:
		//line parser.go.y:279
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 44:
		//line parser.go.y:283
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 45:
		//line parser.go.y:287
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 46:
		//line parser.go.y:291
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 47:
		//line parser.go.y:295
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 48:
		//line parser.go.y:299
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 49:
		//line parser.go.y:303
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 50:
		//line parser.go.y:307
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 51:
		//line parser.go.y:311
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 52:
		//line parser.go.y:315
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 53:
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 54:
		//line parser.go.y:323
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "=", Expr: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:335
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "+", Expr: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "-", Expr: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:343
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "*", Expr: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:347
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "/", Expr: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:351
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:355
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 63:
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 64:
		//line parser.go.y:363
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 65:
		//line parser.go.y:367
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 66:
		//line parser.go.y:371
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 67:
		//line parser.go.y:375
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 68:
		//line parser.go.y:379
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 69:
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 70:
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 71:
		//line parser.go.y:391
		{
			yyVAL.expr = &ast.ConstExpr{Value: nil}
		}
	case 72:
		//line parser.go.y:395
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	}
	goto yystack /* stack new state and value */
}
