//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:22
type yySymType struct {
	yys       int
	stmt_func ast.Stmt
	stmt_if   ast.Stmt
	stmt_else ast.Stmt
	stmt_for  ast.Stmt
	stmts     []ast.Stmt
	stmt      ast.Stmt
	teim      ast.Expr
	expr      ast.Expr
	tok       Token
	idents    []string
	exprs     []ast.Expr
	pair      *ast.PairExpr
	pairs     []*ast.PairExpr
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VAR = 57350
const VARARG = 57351
const FUNC = 57352
const RETURN = 57353
const THROW = 57354
const IF = 57355
const ELSE = 57356
const FOR = 57357
const IN = 57358
const EQ = 57359
const NE = 57360
const GE = 57361
const LE = 57362
const OR = 57363
const AND = 57364
const NEW = 57365
const TRUE = 57366
const FALSE = 57367
const NIL = 57368
const MODULE = 57369
const TRY = 57370
const CATCH = 57371
const FINALLY = 57372
const UNARY = 57373

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VAR",
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

//line parser.go.y:408

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 60
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 379

var yyAct = []int{

	1, 61, 7, 26, 27, 28, 29, 30, 64, 99,
	89, 95, 50, 52, 96, 126, 119, 95, 108, 113,
	60, 95, 62, 109, 66, 108, 102, 98, 95, 54,
	128, 97, 123, 68, 62, 101, 70, 71, 72, 73,
	74, 75, 76, 77, 78, 79, 80, 81, 82, 83,
	84, 85, 39, 40, 42, 44, 46, 48, 88, 91,
	93, 62, 94, 55, 56, 55, 34, 35, 36, 37,
	38, 59, 146, 125, 58, 31, 143, 141, 139, 32,
	33, 138, 41, 43, 45, 47, 137, 135, 105, 104,
	118, 144, 134, 130, 127, 112, 124, 122, 114, 106,
	87, 86, 116, 117, 62, 92, 115, 142, 120, 110,
	65, 121, 107, 90, 69, 67, 57, 53, 49, 63,
	6, 5, 4, 129, 3, 131, 132, 2, 133, 0,
	0, 136, 0, 0, 0, 140, 39, 40, 42, 44,
	46, 48, 0, 0, 0, 145, 0, 0, 0, 0,
	34, 35, 36, 37, 38, 0, 0, 0, 0, 31,
	111, 0, 0, 32, 33, 0, 41, 43, 45, 47,
	39, 40, 42, 44, 46, 48, 0, 0, 0, 0,
	0, 0, 0, 0, 34, 35, 36, 37, 38, 0,
	0, 0, 0, 31, 0, 0, 0, 32, 33, 103,
	41, 43, 45, 47, 39, 40, 42, 44, 46, 48,
	0, 0, 0, 0, 0, 0, 0, 0, 34, 35,
	36, 37, 38, 0, 0, 0, 0, 31, 100, 0,
	0, 32, 33, 0, 41, 43, 45, 47, 39, 40,
	42, 44, 46, 48, 0, 0, 0, 0, 0, 0,
	0, 0, 34, 35, 36, 37, 38, 0, 0, 0,
	0, 31, 0, 0, 0, 32, 33, 0, 41, 43,
	45, 47, 39, 40, 42, 44, 46, 48, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 36, 37,
	38, 0, 0, 0, 0, 31, 0, 0, 0, 32,
	33, 0, 41, 43, 45, 47, 16, 15, 18, 0,
	8, 0, 12, 9, 10, 13, 0, 14, 39, 40,
	42, 44, 46, 48, 0, 22, 23, 24, 25, 11,
	0, 0, 0, 0, 17, 0, 0, 16, 15, 18,
	20, 31, 21, 51, 0, 32, 33, 19, 41, 43,
	45, 47, 0, 0, 0, 0, 22, 23, 24, 25,
	0, 0, 0, 0, 0, 17, 0, 0, 0, 0,
	0, 20, 0, 21, 0, 0, 0, 0, 19,
}
var yyPact = []int{

	302, -1000, 302, 302, 302, 302, 302, 221, 114, 333,
	333, 113, 25, 24, 112, -1000, 34, 333, -1000, 333,
	104, 333, 111, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 333, 110, 333, 333, 333, 333, 333, 333, 333,
	333, 333, 333, 333, 333, 333, 333, 333, 333, 64,
	221, 23, 221, 62, 18, 109, 333, 89, 333, 333,
	301, -32, 221, -12, -1000, -33, 187, -5, -15, -1000,
	153, 255, 255, 301, 301, 301, 221, 221, 221, 221,
	221, 221, 221, 221, 221, 221, 333, 302, 108, -18,
	100, 119, 333, -22, 221, 333, -1000, 104, -1000, 333,
	-1000, 333, -1000, -1000, 221, 51, -25, 99, 107, 59,
	-9, 58, 35, -1000, 221, -1000, 221, -26, -1000, 56,
	-11, -1000, 302, 55, 302, 302, -1000, 302, 54, 48,
	302, 47, 42, 39, 302, -1000, 38, 93, -1000, -1000,
	37, -1000, 53, -1000, 302, 33, -1000,
}
var yyPgo = []int{

	0, 0, 127, 124, 122, 121, 120, 2, 1, 8,
	119, 10,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	2, 2, 3, 3, 5, 4, 6, 9, 10, 10,
	10, 11, 11, 11, 8, 8, 8, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 1, 4, 2,
	2, 5, 8, 9, 11, 7, 7, 3, 0, 1,
	3, 0, 1, 3, 0, 1, 3, 1, 1, 2,
	1, 7, 4, 3, 4, 3, 8, 3, 4, 3,
	5, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 8, 11,
	12, 27, 10, 13, 15, 5, 4, 32, 6, 45,
	38, 40, 23, 24, 25, 26, -1, -1, -1, -1,
	-1, 40, 44, 45, 31, 32, 33, 34, 35, 17,
	18, 47, 19, 48, 20, 49, 21, 50, 22, 4,
	-7, 10, -7, 4, 4, 40, 40, 4, 40, 37,
	-7, -8, -7, -10, -9, 6, -7, 4, -8, 4,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, 37, 38, 40, -11,
	4, -7, 16, -8, -7, 43, 46, 43, 39, 42,
	41, 40, 41, 46, -7, -1, -11, 4, 43, 41,
	9, 41, -7, 41, -7, -9, -7, -8, 39, 41,
	9, 4, 38, 41, 38, 38, 41, 38, 41, -1,
	38, -1, -1, -1, 38, 39, -1, 39, 39, 39,
	-1, 39, 14, 39, 38, -1, 39,
}
var yyDef = []int{

	1, -2, 1, 1, 1, 1, 1, 7, 0, 0,
	0, 0, 0, 0, 0, 27, 28, 0, 30, 24,
	18, 0, 0, 41, 42, 43, 2, 3, 4, 5,
	6, 24, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 10, 0, 0, 21, 0, 0, 24, 0,
	29, 0, 25, 0, 19, 0, 0, 0, 0, 33,
	0, 44, 45, 46, 47, 48, 49, 50, 51, 52,
	53, 54, 56, 57, 58, 59, 0, 1, 21, 0,
	22, 0, 0, 0, 55, 0, 35, 0, 37, 0,
	39, 24, 32, 38, 8, 0, 0, 22, 0, 0,
	0, 0, 0, 34, 26, 20, 17, 0, 11, 0,
	0, 23, 1, 0, 1, 1, 40, 1, 0, 0,
	1, 0, 0, 0, 1, 31, 0, 15, 16, 12,
	0, 36, 0, 13, 1, 0, 14,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 35, 50, 3,
	40, 41, 33, 31, 43, 32, 44, 34, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 42, 3,
	48, 37, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 45, 3, 46, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 38, 49, 39,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 36,
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
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_else}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:82
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
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
			yyVAL.stmt = &ast.VarStmt{Name: yyS[yypt-2].tok.lit, Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyS[yypt-0].expr.SetPos(l.pos)
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
			yyVAL.stmt_else = &ast.IfStmt{Expr: yyS[yypt-8].expr, ThenStmts: yyS[yypt-5].stmts, ElseStmts: yyS[yypt-1].stmts}
		}
	case 15:
		//line parser.go.y:137
		{
			yyVAL.stmt_if = &ast.IfStmt{Expr: yyS[yypt-4].expr, ThenStmts: yyS[yypt-1].stmts}
		}
	case 16:
		//line parser.go.y:142
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 17:
		//line parser.go.y:147
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 18:
		//line parser.go.y:152
		{
			yyVAL.pairs = []*ast.PairExpr{}
		}
	case 19:
		//line parser.go.y:156
		{
			yyVAL.pairs = []*ast.PairExpr{yyS[yypt-0].pair}
		}
	case 20:
		//line parser.go.y:160
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 21:
		//line parser.go.y:165
		{
			yyVAL.idents = []string{}
		}
	case 22:
		//line parser.go.y:169
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 23:
		//line parser.go.y:173
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 24:
		//line parser.go.y:178
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 25:
		//line parser.go.y:182
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:186
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 27:
		//line parser.go.y:191
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 28:
		//line parser.go.y:195
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 29:
		//line parser.go.y:199
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 30:
		//line parser.go.y:203
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 31:
		//line parser.go.y:207
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 32:
		//line parser.go.y:214
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyS[yypt-3].expr, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 33:
		//line parser.go.y:221
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyS[yypt-2].expr, Method: yyS[yypt-0].tok.lit}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 34:
		//line parser.go.y:228
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 35:
		//line parser.go.y:235
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 36:
		//line parser.go.y:242
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 37:
		//line parser.go.y:249
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.Key] = v.Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 38:
		//line parser.go.y:260
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 39:
		//line parser.go.y:267
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 40:
		//line parser.go.y:274
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 41:
		//line parser.go.y:281
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 42:
		//line parser.go.y:288
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 43:
		//line parser.go.y:292
		{
			yyVAL.expr = &ast.ConstExpr{Value: nil}
		}
	case 44:
		//line parser.go.y:296
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 45:
		//line parser.go.y:303
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 46:
		//line parser.go.y:310
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 47:
		//line parser.go.y:317
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 48:
		//line parser.go.y:324
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 49:
		//line parser.go.y:331
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 50:
		//line parser.go.y:338
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 51:
		//line parser.go.y:345
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 52:
		//line parser.go.y:352
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 53:
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 54:
		//line parser.go.y:366
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 55:
		//line parser.go.y:373
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Expr: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 56:
		//line parser.go.y:380
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 57:
		//line parser.go.y:387
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 58:
		//line parser.go.y:394
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	case 59:
		//line parser.go.y:401
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPos(l.pos)
			}
		}
	}
	goto yystack /* stack new state and value */
}
