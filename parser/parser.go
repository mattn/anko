//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:22
type yySymType struct {
	yys                    int
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
const THROW = 57353
const IF = 57354
const ELSE = 57355
const FOR = 57356
const IN = 57357
const EQEQ = 57358
const NEQ = 57359
const GE = 57360
const LE = 57361
const OR = 57362
const AND = 57363
const NEW = 57364
const TRUE = 57365
const FALSE = 57366
const NIL = 57367
const MODULE = 57368
const TRY = 57369
const CATCH = 57370
const FINALLY = 57371
const PLUSEQ = 57372
const MINUSEQ = 57373
const MULEQ = 57374
const DIVEQ = 57375
const UNARY = 57376

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"FUNC",
	"RETURN",
	"THROW",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"OR",
	"AND",
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
	" +",
	" -",
	" *",
	" /",
	" %",
	"UNARY",
	" ?",
	" :",
	" >",
	" <",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:376

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 26,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	44, 26,
	49, 26,
	-2, 1,
	-1, 3,
	44, 26,
	49, 26,
	-2, 1,
	-1, 4,
	44, 26,
	49, 26,
	-2, 1,
	-1, 5,
	44, 26,
	49, 26,
	-2, 1,
	-1, 6,
	44, 26,
	49, 26,
	-2, 1,
	-1, 16,
	49, 27,
	-2, 63,
	-1, 63,
	46, 1,
	-2, 26,
	-1, 78,
	46, 1,
	-2, 26,
	-1, 88,
	16, 0,
	17, 0,
	-2, 47,
	-1, 89,
	16, 0,
	17, 0,
	-2, 48,
	-1, 101,
	46, 1,
	-2, 26,
	-1, 145,
	46, 1,
	-2, 26,
	-1, 147,
	46, 1,
	-2, 26,
	-1, 148,
	46, 1,
	-2, 26,
	-1, 152,
	46, 1,
	-2, 26,
	-1, 155,
	46, 1,
	-2, 26,
	-1, 159,
	46, 1,
	-2, 26,
	-1, 160,
	46, 1,
	-2, 26,
	-1, 162,
	46, 1,
	-2, 26,
	-1, 177,
	46, 1,
	-2, 26,
	-1, 180,
	46, 1,
	-2, 26,
	-1, 185,
	46, 1,
	-2, 26,
}

const yyNprod = 70
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 593

var yyAct = []int{

	1, 7, 73, 27, 28, 29, 31, 32, 150, 114,
	114, 54, 57, 115, 114, 8, 143, 53, 173, 71,
	117, 75, 70, 116, 77, 134, 114, 129, 53, 123,
	114, 52, 153, 146, 59, 71, 53, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 97, 98, 71, 122, 80, 41, 42, 44,
	46, 48, 50, 105, 107, 120, 71, 109, 110, 111,
	112, 113, 159, 102, 158, 99, 103, 60, 60, 121,
	61, 51, 187, 43, 45, 184, 181, 108, 33, 176,
	175, 34, 35, 174, 47, 49, 66, 67, 68, 69,
	172, 170, 126, 166, 44, 46, 48, 50, 132, 165,
	65, 48, 50, 64, 163, 142, 135, 139, 127, 136,
	137, 133, 71, 185, 140, 79, 180, 141, 43, 45,
	177, 149, 162, 33, 160, 155, 34, 35, 33, 47,
	49, 34, 35, 138, 47, 49, 154, 152, 156, 157,
	147, 145, 101, 161, 63, 118, 164, 183, 78, 178,
	168, 169, 106, 171, 144, 130, 74, 167, 128, 41,
	42, 44, 46, 48, 50, 104, 100, 81, 179, 76,
	62, 182, 58, 72, 6, 5, 186, 36, 37, 38,
	39, 40, 4, 51, 3, 43, 45, 2, 0, 0,
	33, 151, 0, 34, 35, 0, 47, 49, 41, 42,
	44, 46, 48, 50, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 36, 37, 38, 39,
	40, 0, 51, 0, 43, 45, 0, 148, 0, 33,
	0, 0, 34, 35, 0, 47, 49, 41, 42, 44,
	46, 48, 50, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 36, 37, 38, 39, 40,
	0, 51, 0, 43, 45, 0, 0, 0, 33, 131,
	0, 34, 35, 0, 47, 49, 41, 42, 44, 46,
	48, 50, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 36, 37, 38, 39, 40, 0,
	51, 125, 43, 45, 0, 0, 0, 33, 0, 0,
	34, 35, 0, 47, 49, 41, 42, 44, 46, 48,
	50, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 36, 37, 38, 39, 40, 0, 51,
	0, 43, 45, 0, 0, 0, 33, 0, 0, 34,
	35, 124, 47, 49, 41, 42, 44, 46, 48, 50,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 36, 37, 38, 39, 40, 0, 51, 0,
	43, 45, 0, 0, 0, 33, 119, 0, 34, 35,
	0, 47, 49, 41, 42, 44, 46, 48, 50, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 36, 37, 38, 39, 40, 0, 51, 0, 43,
	45, 0, 0, 0, 33, 0, 0, 34, 35, 0,
	47, 49, 41, 42, 44, 46, 48, 50, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 38, 39, 40, 0, 51, 0, 43, 45,
	0, 0, 0, 33, 0, 0, 34, 35, 0, 47,
	49, 16, 21, 23, 0, 0, 12, 9, 10, 13,
	30, 14, 0, 0, 0, 0, 0, 0, 0, 20,
	24, 25, 26, 11, 15, 0, 0, 0, 0, 0,
	0, 0, 22, 16, 21, 23, 0, 0, 12, 9,
	10, 13, 18, 14, 19, 0, 0, 0, 17, 0,
	0, 20, 24, 25, 26, 11, 15, 0, 0, 0,
	0, 0, 0, 0, 22, 56, 21, 23, 0, 0,
	55, 0, 0, 0, 18, 0, 19, 0, 0, 0,
	17, 0, 0, 20, 24, 25, 26, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 22, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 18, 0, 19, 0,
	0, 0, 17,
}
var yyPact = []int{

	509, -1000, 509, 509, 477, 509, 509, 387, -13, 541,
	541, 178, 30, 33, 176, 109, 66, 541, 160, 541,
	175, -1000, 541, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	113, -1000, -1000, 541, 173, 541, 541, 541, 541, 541,
	541, 541, 541, 541, 541, 541, 541, 541, 541, 541,
	541, 541, 541, 172, 387, 31, 66, 387, 107, 26,
	171, 541, 147, 509, 541, 541, 541, 541, 541, 541,
	-39, 387, -26, -1000, 114, 348, 18, 41, 509, 8,
	-19, -1000, 309, 426, 426, 41, 41, 41, 86, 86,
	91, 91, 91, 91, 387, 387, 387, 387, 270, -35,
	-1000, 509, 164, -21, 157, 231, 541, 75, -23, 387,
	387, 387, 387, 387, 541, -1000, 160, -1000, 541, -1000,
	541, 71, 541, -1000, -1000, 541, 69, -32, 156, 106,
	-15, 105, 192, 103, -1000, 387, -1000, 387, -40, -1000,
	153, 41, -1000, 102, -16, 509, 90, 509, 509, 27,
	-1000, 89, 509, 87, 68, 509, 63, 57, 163, 509,
	509, 55, 509, -1000, 54, -1000, -1000, -30, 47, 44,
	-1000, 43, -1000, 85, 130, -1000, -1000, 509, 81, 40,
	509, 128, 39, 78, -1000, 509, 36, -1000,
}
var yyPgo = []int{

	0, 0, 197, 194, 192, 185, 184, 1, 22, 2,
	183, 15,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	2, 2, 3, 3, 4, 4, 4, 5, 6, 6,
	6, 6, 9, 10, 10, 10, 11, 11, 11, 8,
	8, 8, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 1, 3, 2,
	2, 5, 8, 9, 7, 5, 9, 7, 15, 12,
	11, 8, 3, 0, 1, 3, 0, 1, 3, 0,
	1, 3, 4, 7, 3, 4, 3, 8, 3, 4,
	3, 5, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1, 1, 2, 1, 1, 1, 1, 5,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -11, 10,
	11, 26, 9, 12, 14, 27, 4, 51, 45, 47,
	22, 5, 35, 6, 23, 24, 25, -1, -1, -1,
	13, -1, -1, 47, 50, 51, 34, 35, 36, 37,
	38, 16, 17, 42, 18, 43, 19, 53, 20, 54,
	21, 40, 44, 49, -7, 9, 4, -7, 4, 4,
	47, 47, 4, 45, 47, 44, 30, 31, 32, 33,
	-8, -7, -10, -9, 6, -7, 4, -7, 45, 12,
	-8, 4, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -8,
	4, 45, 47, -11, 4, -7, 15, -1, -8, -7,
	-7, -7, -7, -7, 49, 52, 49, 46, 41, 48,
	47, -1, 47, 48, 52, 41, -1, -11, 4, 48,
	8, 48, -7, 46, 48, -7, -9, -7, -8, 46,
	-7, -7, 46, 48, 8, 45, 48, 45, 45, 28,
	48, 48, 45, 48, -1, 45, -1, -1, 47, 45,
	45, -1, 45, 46, -1, 46, 46, 4, -1, -1,
	46, -1, 46, 48, 46, 46, 46, 45, 29, -1,
	45, 46, -1, 29, 46, 45, -1, 46,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, 7, 0, 0,
	0, 0, 0, 0, 0, 0, -2, 29, 23, 0,
	0, 62, 0, 65, 66, 67, 68, 2, 3, 4,
	0, 5, 6, 29, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 29, 0, 9, 0, 63, 10, 0, 0,
	26, 0, 0, -2, 29, 0, 0, 0, 0, 0,
	0, 30, 0, 24, 0, 0, 0, 64, -2, 0,
	0, 34, 0, 42, 43, 44, 45, 46, -2, -2,
	49, 50, 51, 52, 58, 59, 60, 61, 0, 8,
	28, -2, 26, 0, 27, 0, 0, 0, 0, 53,
	54, 55, 56, 57, 0, 36, 0, 38, 0, 40,
	29, 0, 0, 32, 39, 0, 0, 0, 27, 0,
	0, 0, 0, 0, 35, 31, 25, 22, 0, 15,
	0, 69, 11, 0, 0, -2, 0, -2, -2, 0,
	41, 0, -2, 0, 0, -2, 0, 0, 0, -2,
	-2, 0, -2, 33, 0, 14, 17, 0, 0, 0,
	12, 0, 37, 0, 21, 16, 13, -2, 0, 0,
	-2, 20, 0, 0, 19, -2, 0, 18,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 38, 54, 3,
	47, 48, 36, 34, 49, 35, 50, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 41, 3,
	43, 44, 42, 40, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 51, 3, 52, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 45, 53, 46,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 39,
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
		//line parser.go.y:52
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:59
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:66
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:73
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:80
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:87
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:95
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 8:
		//line parser.go.y:102
		{
			yyVAL.stmt = &ast.LetStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.stmt.SetPos(l.pos)
			}
		}
	case 9:
		//line parser.go.y:109
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 10:
		//line parser.go.y:116
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:123
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 12:
		//line parser.go.y:128
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 13:
		//line parser.go.y:132
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 14:
		//line parser.go.y:137
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 15:
		//line parser.go.y:141
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 16:
		//line parser.go.y:149
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 17:
		//line parser.go.y:154
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:159
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:163
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:167
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:171
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:176
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 23:
		//line parser.go.y:181
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 24:
		//line parser.go.y:185
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 25:
		//line parser.go.y:189
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 26:
		//line parser.go.y:194
		{
			yyVAL.idents = []string{}
		}
	case 27:
		//line parser.go.y:198
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 28:
		//line parser.go.y:202
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 29:
		//line parser.go.y:207
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 30:
		//line parser.go.y:211
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 31:
		//line parser.go.y:215
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 32:
		//line parser.go.y:220
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
		}
	case 33:
		//line parser.go.y:224
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 34:
		//line parser.go.y:228
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
		}
	case 35:
		//line parser.go.y:232
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 36:
		//line parser.go.y:236
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 37:
		//line parser.go.y:240
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 38:
		//line parser.go.y:244
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 39:
		//line parser.go.y:252
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 40:
		//line parser.go.y:256
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 41:
		//line parser.go.y:260
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 42:
		//line parser.go.y:264
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 43:
		//line parser.go.y:268
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 44:
		//line parser.go.y:272
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 45:
		//line parser.go.y:276
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 46:
		//line parser.go.y:280
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 47:
		//line parser.go.y:284
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 48:
		//line parser.go.y:288
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 49:
		//line parser.go.y:292
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 50:
		//line parser.go.y:296
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 51:
		//line parser.go.y:300
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 52:
		//line parser.go.y:304
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 53:
		//line parser.go.y:308
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "=", Expr: yyS[yypt-0].expr}
		}
	case 54:
		//line parser.go.y:312
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "+", Expr: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:316
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "-", Expr: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:320
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "*", Expr: yyS[yypt-0].expr}
		}
	case 57:
		//line parser.go.y:324
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "/", Expr: yyS[yypt-0].expr}
		}
	case 58:
		//line parser.go.y:328
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 59:
		//line parser.go.y:332
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 60:
		//line parser.go.y:336
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 61:
		//line parser.go.y:340
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	case 62:
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 63:
		//line parser.go.y:348
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 64:
		//line parser.go.y:352
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 65:
		//line parser.go.y:356
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 66:
		//line parser.go.y:360
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 67:
		//line parser.go.y:364
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 68:
		//line parser.go.y:368
		{
			yyVAL.expr = &ast.ConstExpr{Value: nil}
		}
	case 69:
		//line parser.go.y:372
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyS[yypt-4].expr, Lhs: yyS[yypt-2].expr, Rhs: yyS[yypt-0].expr}
		}
	}
	goto yystack /* stack new state and value */
}
