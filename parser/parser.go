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
const UNARY = 57369

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VAR",
	"VARARG",
	"FUNC",
	"RETURN",
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

//line parser.go.y:304

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 57
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 328

var yyAct = []int{

	1, 57, 7, 25, 26, 27, 28, 29, 60, 89,
	83, 90, 47, 119, 112, 89, 101, 93, 106, 56,
	89, 58, 102, 62, 101, 92, 121, 55, 116, 91,
	54, 95, 82, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 36,
	37, 39, 41, 43, 45, 85, 87, 58, 88, 50,
	31, 32, 33, 34, 35, 51, 52, 118, 113, 139,
	136, 134, 132, 30, 131, 38, 40, 42, 44, 130,
	128, 111, 98, 97, 137, 127, 123, 120, 117, 105,
	115, 51, 107, 99, 81, 80, 109, 110, 58, 86,
	108, 135, 103, 61, 114, 100, 84, 63, 53, 49,
	46, 59, 6, 5, 4, 3, 122, 2, 124, 125,
	0, 126, 0, 0, 129, 0, 0, 0, 133, 36,
	37, 39, 41, 43, 45, 0, 0, 0, 138, 0,
	31, 32, 33, 34, 35, 0, 0, 0, 0, 0,
	104, 0, 0, 30, 0, 38, 40, 42, 44, 36,
	37, 39, 41, 43, 45, 0, 0, 0, 0, 0,
	31, 32, 33, 34, 35, 0, 0, 0, 0, 0,
	0, 0, 0, 30, 96, 38, 40, 42, 44, 36,
	37, 39, 41, 43, 45, 0, 0, 0, 0, 0,
	31, 32, 33, 34, 35, 0, 0, 0, 0, 0,
	94, 0, 0, 30, 0, 38, 40, 42, 44, 36,
	37, 39, 41, 43, 45, 0, 0, 0, 0, 0,
	31, 32, 33, 34, 35, 0, 36, 37, 39, 41,
	43, 45, 0, 30, 0, 38, 40, 42, 44, 33,
	34, 35, 0, 36, 37, 39, 41, 43, 45, 0,
	30, 0, 38, 40, 42, 44, 15, 14, 17, 0,
	8, 0, 11, 9, 12, 0, 13, 30, 0, 38,
	40, 42, 44, 0, 21, 22, 23, 24, 10, 0,
	16, 15, 14, 17, 0, 0, 19, 48, 20, 0,
	0, 0, 18, 0, 0, 0, 0, 0, 0, 21,
	22, 23, 24, 0, 0, 16, 0, 0, 0, 0,
	0, 19, 0, 20, 0, 0, 0, 18,
}
var yyPact = []int{

	262, -1000, 262, 262, 262, 262, 262, 203, 106, 287,
	105, 55, 30, 104, -1000, -6, 287, -1000, 287, 97,
	287, 103, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	287, 287, 287, 287, 287, 287, 287, 287, 287, 287,
	287, 287, 287, 287, 287, 287, 62, 203, 29, 60,
	-4, 102, 287, 84, 287, 287, 237, -30, 203, -10,
	-1000, -21, 173, -5, 143, 220, 220, 237, 237, 237,
	203, 203, 203, 203, 203, 203, 203, 203, 203, 203,
	287, 262, 101, -15, 93, 113, 287, -19, 203, 287,
	-1000, 97, -1000, 287, -1000, 287, -1000, 203, 46, -23,
	59, 100, 56, -9, 54, 33, -1000, 203, -1000, 203,
	-24, -1000, 53, -11, -1000, 262, 52, 262, 262, -1000,
	262, 51, 45, 262, 44, 39, 37, 262, -1000, 36,
	88, -1000, -1000, 35, -1000, 50, -1000, 262, 34, -1000,
}
var yyPgo = []int{

	0, 0, 117, 115, 114, 113, 112, 2, 1, 8,
	111, 10,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	2, 3, 3, 5, 4, 6, 9, 10, 10, 10,
	11, 11, 11, 8, 8, 8, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 1, 4, 2,
	5, 8, 9, 11, 7, 7, 3, 0, 1, 3,
	0, 1, 3, 0, 1, 3, 1, 1, 2, 1,
	4, 3, 7, 8, 3, 4, 3, 5, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 8, 11,
	26, 10, 12, 14, 5, 4, 28, 6, 40, 34,
	36, 22, 23, 24, 25, -1, -1, -1, -1, -1,
	40, 27, 28, 29, 30, 31, 16, 17, 42, 18,
	43, 19, 44, 20, 45, 21, 4, -7, 10, 4,
	4, 36, 36, 4, 36, 33, -7, -8, -7, -10,
	-9, 6, -7, 4, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	33, 34, 36, -11, 4, -7, 15, -8, -7, 39,
	41, 39, 35, 38, 37, 36, 41, -7, -1, -11,
	4, 39, 37, 9, 37, -7, 37, -7, -9, -7,
	-8, 35, 37, 9, 4, 34, 37, 34, 34, 37,
	34, 37, -1, 34, -1, -1, -1, 34, 35, -1,
	35, 35, 35, -1, 35, 13, 35, 34, -1, 35,
}
var yyDef = []int{

	1, -2, 1, 1, 1, 1, 1, 7, 0, 0,
	0, 0, 0, 0, 26, 27, 0, 29, 23, 17,
	0, 0, 38, 39, 40, 2, 3, 4, 5, 6,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 9, 0, 0,
	0, 20, 0, 0, 23, 0, 28, 0, 24, 0,
	18, 0, 0, 0, 0, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 50, 51, 53, 54, 55, 56,
	0, 1, 20, 0, 21, 0, 0, 0, 52, 0,
	31, 0, 34, 0, 36, 23, 35, 8, 0, 0,
	21, 0, 0, 0, 0, 0, 30, 25, 19, 16,
	0, 10, 0, 0, 22, 1, 0, 1, 1, 37,
	1, 0, 0, 1, 0, 0, 0, 1, 32, 0,
	14, 15, 11, 0, 33, 0, 12, 1, 0, 13,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 31, 45, 3,
	36, 37, 29, 27, 39, 28, 3, 30, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 38, 3,
	43, 33, 42, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 40, 3, 41, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 34, 44, 35,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 32,
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
				yyS[yypt-1].stmt.Position(l.pos)
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:62
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:69
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:76
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_else}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:83
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:91
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-0].expr}
		}
	case 8:
		//line parser.go.y:95
		{
			yyVAL.stmt = &ast.VarStmt{Name: yyS[yypt-2].tok.lit, Expr: yyS[yypt-0].expr}
		}
	case 9:
		//line parser.go.y:99
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-0].expr}
		}
	case 10:
		//line parser.go.y:103
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyS[yypt-3].tok.lit, Stmts: yyS[yypt-1].stmts}
		}
	case 11:
		//line parser.go.y:108
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 12:
		//line parser.go.y:112
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-7].tok.lit, Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 13:
		//line parser.go.y:117
		{
			yyVAL.stmt_else = &ast.IfStmt{Expr: yyS[yypt-8].expr, ThenStmts: yyS[yypt-5].stmts, ElseStmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:122
		{
			yyVAL.stmt_if = &ast.IfStmt{Expr: yyS[yypt-4].expr, ThenStmts: yyS[yypt-1].stmts}
		}
	case 15:
		//line parser.go.y:127
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 16:
		//line parser.go.y:132
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 17:
		//line parser.go.y:137
		{
			yyVAL.pairs = []*ast.PairExpr{}
		}
	case 18:
		//line parser.go.y:141
		{
			yyVAL.pairs = []*ast.PairExpr{yyS[yypt-0].pair}
		}
	case 19:
		//line parser.go.y:145
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 20:
		//line parser.go.y:150
		{
			yyVAL.idents = []string{}
		}
	case 21:
		//line parser.go.y:154
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 22:
		//line parser.go.y:158
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 23:
		//line parser.go.y:163
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 24:
		//line parser.go.y:167
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 25:
		//line parser.go.y:171
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 26:
		//line parser.go.y:176
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 27:
		//line parser.go.y:180
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 28:
		//line parser.go.y:184
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 29:
		//line parser.go.y:188
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 30:
		//line parser.go.y:192
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 31:
		//line parser.go.y:196
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 32:
		//line parser.go.y:200
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 33:
		//line parser.go.y:204
		{
			yyVAL.expr = &ast.FuncExpr{Args: []string{yyS[yypt-5].tok.lit}, Stmts: yyS[yypt-1].stmts, VarArg: true}
		}
	case 34:
		//line parser.go.y:208
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.Key] = v.Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 35:
		//line parser.go.y:216
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 36:
		//line parser.go.y:220
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 37:
		//line parser.go.y:224
		{
			yyVAL.expr = &ast.NewExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 38:
		//line parser.go.y:228
		{
			yyVAL.expr = &ast.ConstExpr{Value: true}
		}
	case 39:
		//line parser.go.y:232
		{
			yyVAL.expr = &ast.ConstExpr{Value: false}
		}
	case 40:
		//line parser.go.y:236
		{
			yyVAL.expr = &ast.ConstExpr{Value: nil}
		}
	case 41:
		//line parser.go.y:240
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 42:
		//line parser.go.y:244
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 43:
		//line parser.go.y:248
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 44:
		//line parser.go.y:252
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 45:
		//line parser.go.y:256
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 46:
		//line parser.go.y:260
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 47:
		//line parser.go.y:264
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 48:
		//line parser.go.y:268
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 49:
		//line parser.go.y:272
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 50:
		//line parser.go.y:276
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 51:
		//line parser.go.y:280
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	case 52:
		//line parser.go.y:284
		{
			yyVAL.expr = &ast.LetExpr{Name: yyS[yypt-2].tok.lit, Expr: yyS[yypt-0].expr}
		}
	case 53:
		//line parser.go.y:288
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "|", Rhs: yyS[yypt-0].expr}
		}
	case 54:
		//line parser.go.y:292
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "||", Rhs: yyS[yypt-0].expr}
		}
	case 55:
		//line parser.go.y:296
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&", Rhs: yyS[yypt-0].expr}
		}
	case 56:
		//line parser.go.y:300
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "&&", Rhs: yyS[yypt-0].expr}
		}
	}
	goto yystack /* stack new state and value */
}
