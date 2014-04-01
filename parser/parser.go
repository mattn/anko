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
const EQ = 57358
const NE = 57359
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
	"EQ",
	"NE",
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
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:460

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 26,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	40, 26,
	46, 26,
	-2, 1,
	-1, 3,
	40, 26,
	46, 26,
	-2, 1,
	-1, 4,
	40, 26,
	46, 26,
	-2, 1,
	-1, 5,
	40, 26,
	46, 26,
	-2, 1,
	-1, 6,
	40, 26,
	46, 26,
	-2, 1,
	-1, 17,
	46, 27,
	-2, 33,
	-1, 62,
	42, 1,
	-2, 26,
	-1, 77,
	42, 1,
	-2, 26,
	-1, 99,
	42, 1,
	-2, 26,
	-1, 141,
	42, 1,
	-2, 26,
	-1, 143,
	42, 1,
	-2, 26,
	-1, 144,
	42, 1,
	-2, 26,
	-1, 148,
	42, 1,
	-2, 26,
	-1, 151,
	42, 1,
	-2, 26,
	-1, 155,
	42, 1,
	-2, 26,
	-1, 156,
	42, 1,
	-2, 26,
	-1, 158,
	42, 1,
	-2, 26,
	-1, 173,
	42, 1,
	-2, 26,
	-1, 176,
	42, 1,
	-2, 26,
	-1, 181,
	42, 1,
	-2, 26,
}

const yyNprod = 69
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 531

var yyAct = []int{

	1, 7, 73, 27, 28, 29, 31, 32, 112, 169,
	112, 53, 56, 113, 8, 65, 66, 67, 68, 115,
	69, 70, 71, 114, 75, 64, 116, 146, 63, 112,
	139, 131, 52, 112, 126, 71, 52, 81, 82, 83,
	84, 85, 86, 87, 88, 89, 90, 91, 92, 93,
	94, 95, 96, 71, 149, 79, 41, 42, 44, 46,
	48, 50, 103, 105, 142, 71, 107, 108, 109, 110,
	111, 58, 120, 97, 101, 121, 51, 112, 119, 155,
	118, 154, 52, 33, 100, 106, 59, 34, 35, 60,
	43, 45, 47, 49, 183, 180, 177, 172, 171, 170,
	123, 168, 166, 162, 161, 159, 129, 138, 136, 78,
	59, 130, 181, 176, 132, 124, 173, 133, 134, 158,
	71, 156, 137, 151, 148, 143, 141, 99, 62, 179,
	174, 145, 104, 140, 127, 74, 163, 125, 77, 102,
	135, 98, 150, 80, 152, 153, 76, 61, 57, 157,
	72, 6, 160, 5, 4, 3, 164, 165, 2, 167,
	0, 0, 0, 0, 0, 41, 42, 44, 46, 48,
	50, 0, 0, 0, 175, 0, 0, 178, 0, 0,
	0, 0, 182, 36, 37, 38, 39, 40, 0, 0,
	0, 0, 33, 147, 0, 0, 34, 35, 0, 43,
	45, 47, 49, 41, 42, 44, 46, 48, 50, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 36, 37, 38, 39, 40, 0, 0, 144, 0,
	33, 0, 0, 0, 34, 35, 0, 43, 45, 47,
	49, 41, 42, 44, 46, 48, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 36,
	37, 38, 39, 40, 0, 0, 0, 0, 33, 128,
	0, 0, 34, 35, 0, 43, 45, 47, 49, 41,
	42, 44, 46, 48, 50, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 36, 37, 38,
	39, 40, 0, 0, 0, 0, 33, 0, 0, 0,
	34, 35, 122, 43, 45, 47, 49, 41, 42, 44,
	46, 48, 50, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 36, 37, 38, 39, 40,
	0, 0, 0, 0, 33, 117, 0, 0, 34, 35,
	0, 43, 45, 47, 49, 41, 42, 44, 46, 48,
	50, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 36, 37, 38, 39, 40, 0, 0,
	0, 0, 33, 0, 0, 0, 34, 35, 0, 43,
	45, 47, 49, 41, 42, 44, 46, 48, 50, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 38, 39, 40, 0, 0, 0, 0,
	33, 0, 0, 0, 34, 35, 0, 43, 45, 47,
	49, 17, 16, 19, 0, 0, 12, 9, 10, 13,
	30, 14, 0, 0, 0, 0, 0, 0, 0, 23,
	24, 25, 26, 11, 15, 0, 0, 0, 0, 17,
	16, 19, 18, 0, 12, 9, 10, 13, 21, 14,
	22, 0, 0, 0, 0, 20, 0, 23, 24, 25,
	26, 11, 15, 0, 0, 0, 54, 16, 19, 0,
	18, 55, 0, 0, 0, 0, 21, 0, 22, 0,
	0, 0, 0, 20, 23, 24, 25, 26, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 18, 0, 0,
	0, 0, 0, 21, 0, 22, 0, 0, 0, 0,
	20,
}
var yyPact = []int{

	455, -1000, 455, 455, 427, 455, 455, 339, 36, 482,
	482, 144, 67, 46, 143, 87, -1000, -15, 482, -1000,
	482, 129, 482, 142, -1000, -1000, -1000, -1000, -1000, -1000,
	97, -1000, -1000, 482, 139, 482, 482, 482, 482, 482,
	482, 482, 482, 482, 482, 482, 482, 482, 482, 482,
	482, 482, 137, 339, -15, 43, 339, 86, 41, 135,
	482, 117, 455, 482, 482, 482, 482, 482, 482, 40,
	-36, 339, -23, -1000, -19, 301, 37, 455, 29, 31,
	-1000, 263, 377, 377, 40, 40, 40, 339, 339, 339,
	339, 339, 339, 339, 339, 339, 339, -38, -1000, 455,
	133, -10, 126, 225, 482, 69, -13, 339, 339, 339,
	339, 339, 482, -1000, 129, -1000, 482, -1000, 482, 66,
	482, -1000, -1000, 65, -14, 125, 85, 20, 84, 187,
	103, -1000, 339, -1000, 339, -17, -1000, 149, -1000, 83,
	10, 455, 82, 455, 455, 38, -1000, 80, 455, 78,
	63, 455, 62, 61, 132, 455, 455, 60, 455, -1000,
	59, -1000, -1000, -35, 57, 56, -1000, 55, -1000, 75,
	101, -1000, -1000, 455, 72, 54, 455, 100, 53, 71,
	-1000, 455, 52, -1000,
}
var yyPgo = []int{

	0, 0, 158, 155, 154, 153, 151, 1, 21, 2,
	150, 14,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	2, 2, 3, 3, 4, 4, 4, 5, 6, 6,
	6, 6, 9, 10, 10, 10, 11, 11, 11, 8,
	8, 8, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 1, 3, 2,
	2, 5, 8, 9, 7, 5, 9, 7, 15, 12,
	11, 8, 3, 0, 1, 3, 0, 1, 3, 0,
	1, 3, 1, 1, 2, 1, 7, 4, 3, 4,
	3, 8, 3, 4, 3, 5, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -11, 10,
	11, 26, 9, 12, 14, 27, 5, 4, 35, 6,
	48, 41, 43, 22, 23, 24, 25, -1, -1, -1,
	13, -1, -1, 43, 47, 48, 34, 35, 36, 37,
	38, 16, 17, 50, 18, 51, 19, 52, 20, 53,
	21, 40, 46, -7, 4, 9, -7, 4, 4, 43,
	43, 4, 41, 43, 40, 30, 31, 32, 33, -7,
	-8, -7, -10, -9, 6, -7, 4, 41, 12, -8,
	4, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -8, 4, 41,
	43, -11, 4, -7, 15, -1, -8, -7, -7, -7,
	-7, -7, 46, 49, 46, 42, 45, 44, 43, -1,
	43, 44, 49, -1, -11, 4, 44, 8, 44, -7,
	42, 44, -7, -9, -7, -8, 42, -7, 42, 44,
	8, 41, 44, 41, 41, 28, 44, 44, 41, 44,
	-1, 41, -1, -1, 43, 41, 41, -1, 41, 42,
	-1, 42, 42, 4, -1, -1, 42, -1, 42, 44,
	42, 42, 42, 41, 29, -1, 41, 42, -1, 29,
	42, 41, -1, 42,
}
var yyDef = []int{

	-2, -2, -2, -2, -2, -2, -2, 7, 0, 0,
	0, 0, 0, 0, 0, 0, 32, -2, 0, 35,
	29, 23, 0, 0, 46, 47, 48, 2, 3, 4,
	0, 5, 6, 29, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 29, 0, 9, 33, 0, 10, 0, 0, 26,
	0, 0, -2, 29, 0, 0, 0, 0, 0, 34,
	0, 30, 0, 24, 0, 0, 0, -2, 0, 0,
	38, 0, 49, 50, 51, 52, 53, 54, 55, 56,
	57, 58, 59, 65, 66, 67, 68, 8, 28, -2,
	26, 0, 27, 0, 0, 0, 0, 60, 61, 62,
	63, 64, 0, 40, 0, 42, 0, 44, 29, 0,
	0, 37, 43, 0, 0, 27, 0, 0, 0, 0,
	0, 39, 31, 25, 22, 0, 15, 0, 11, 0,
	0, -2, 0, -2, -2, 0, 45, 0, -2, 0,
	0, -2, 0, 0, 0, -2, -2, 0, -2, 36,
	0, 14, 17, 0, 0, 0, 12, 0, 41, 0,
	21, 16, 13, -2, 0, 0, -2, 20, 0, 0,
	19, -2, 0, 18,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 38, 53, 3,
	43, 44, 36, 34, 46, 35, 47, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 45, 3,
	51, 40, 50, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 48, 3, 49, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 41, 52, 42,
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
		//line parser.go.y:47
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:54
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:61
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:68
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:75
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:82
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_try_catch_finally}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:90
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 8:
		//line parser.go.y:97
		{
			yyVAL.stmt = &ast.LetStmt{Names: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.stmt.SetPos(l.pos)
			}
		}
	case 9:
		//line parser.go.y:104
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 10:
		//line parser.go.y:111
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
			}
		}
	case 11:
		//line parser.go.y:118
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 12:
		//line parser.go.y:123
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 13:
		//line parser.go.y:127
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 14:
		//line parser.go.y:132
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts}
		}
	case 15:
		//line parser.go.y:136
		{
			if yyS[yypt-4].stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyS[yypt-4].stmt_if.(*ast.IfStmt).Else = yyS[yypt-1].stmts
			}
		}
	case 16:
		//line parser.go.y:144
		{
			yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf = append(yyS[yypt-8].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyS[yypt-4].expr, Then: yyS[yypt-1].stmts})
		}
	case 17:
		//line parser.go.y:149
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 18:
		//line parser.go.y:154
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-12].stmts, Var: yyS[yypt-8].tok.lit, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 19:
		//line parser.go.y:158
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-9].stmts, Catch: yyS[yypt-5].stmts, Finally: yyS[yypt-1].stmts}
		}
	case 20:
		//line parser.go.y:162
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-8].stmts, Var: yyS[yypt-4].tok.lit, Catch: yyS[yypt-1].stmts}
		}
	case 21:
		//line parser.go.y:166
		{
			yyVAL.stmt_try_catch_finally = &ast.TryStmt{Try: yyS[yypt-5].stmts, Catch: yyS[yypt-1].stmts}
		}
	case 22:
		//line parser.go.y:171
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 23:
		//line parser.go.y:176
		{
			yyVAL.pairs = []ast.Expr{}
		}
	case 24:
		//line parser.go.y:180
		{
			yyVAL.pairs = []ast.Expr{yyS[yypt-0].pair}
		}
	case 25:
		//line parser.go.y:184
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 26:
		//line parser.go.y:189
		{
			yyVAL.idents = []string{}
		}
	case 27:
		//line parser.go.y:193
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 28:
		//line parser.go.y:197
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 29:
		//line parser.go.y:202
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 30:
		//line parser.go.y:206
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 31:
		//line parser.go.y:210
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 32:
		//line parser.go.y:215
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 33:
		//line parser.go.y:219
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 34:
		//line parser.go.y:223
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 35:
		//line parser.go.y:227
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 36:
		//line parser.go.y:231
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 37:
		//line parser.go.y:238
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 38:
		//line parser.go.y:245
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 39:
		//line parser.go.y:252
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 40:
		//line parser.go.y:259
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 41:
		//line parser.go.y:266
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 42:
		//line parser.go.y:273
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 43:
		//line parser.go.y:284
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 44:
		//line parser.go.y:291
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 45:
		//line parser.go.y:298
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 46:
		//line parser.go.y:305
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 47:
		//line parser.go.y:312
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 48:
		//line parser.go.y:316
		{
			yyVAL.expr = &ast.ConstExpr{Value: nil}
		}
	case 49:
		//line parser.go.y:320
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 50:
		//line parser.go.y:327
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 51:
		//line parser.go.y:334
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 52:
		//line parser.go.y:341
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:348
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:355
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:362
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:376
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:383
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 59:
		//line parser.go.y:390
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 60:
		//line parser.go.y:397
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "=", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 61:
		//line parser.go.y:404
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "+", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 62:
		//line parser.go.y:411
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "-", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 63:
		//line parser.go.y:418
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "*", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 64:
		//line parser.go.y:425
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Operator: "/", Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 65:
		//line parser.go.y:432
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 66:
		//line parser.go.y:439
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 67:
		//line parser.go.y:446
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 68:
		//line parser.go.y:453
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
