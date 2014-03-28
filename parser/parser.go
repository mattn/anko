//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:19
type yySymType struct {
	yys          int
	stmt_func    ast.Stmt
	stmt_if      ast.Stmt
	stmt_if_else ast.Stmt
	stmts        []ast.Stmt
	stmt         ast.Stmt
	expr         ast.Expr
	tok          Token
	idents       []string
	exprs        []ast.Expr
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const VAR = 57349
const FUNC = 57350
const RETURN = 57351
const IF = 57352
const ELSE = 57353
const EQ = 57354
const NE = 57355
const GE = 57356
const LE = 57357
const UNARY = 57358

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"VAR",
	"FUNC",
	"RETURN",
	"IF",
	"ELSE",
	"EQ",
	"NE",
	"GE",
	"LE",
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

//line parser.go.y:194

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 34
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 196

var yyAct = []int{

	1, 6, 77, 16, 17, 18, 19, 73, 64, 62,
	33, 65, 61, 72, 75, 37, 68, 38, 26, 27,
	29, 31, 66, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 52, 28, 30, 53, 55, 26,
	27, 29, 31, 21, 22, 23, 24, 25, 36, 63,
	35, 50, 57, 12, 11, 14, 28, 30, 74, 69,
	59, 34, 32, 67, 54, 58, 13, 70, 5, 71,
	4, 3, 2, 15, 0, 0, 76, 26, 27, 29,
	31, 21, 22, 23, 24, 25, 0, 0, 0, 0,
	60, 0, 0, 0, 28, 30, 26, 27, 29, 31,
	21, 22, 23, 24, 25, 0, 0, 0, 0, 56,
	0, 0, 0, 28, 30, 26, 27, 29, 31, 21,
	22, 23, 24, 25, 0, 51, 0, 0, 0, 0,
	0, 0, 28, 30, 26, 27, 29, 31, 21, 22,
	23, 24, 25, 0, 20, 0, 0, 0, 0, 0,
	0, 28, 30, 26, 27, 29, 31, 21, 22, 23,
	24, 25, 0, 0, 0, 0, 26, 27, 29, 31,
	28, 30, 23, 24, 25, 12, 11, 14, 7, 9,
	8, 10, 0, 28, 30, 0, 0, 0, 13, 0,
	0, 0, 0, 0, 0, 15,
}
var yyPact = []int{

	171, -1000, 171, 171, 171, 171, 122, 58, 49, 57,
	26, -1000, 24, 49, -1000, 49, -1000, -1000, -1000, -1000,
	-1000, 49, 49, 49, 49, 49, 49, 49, 49, 49,
	49, 49, 28, 103, 10, 49, 49, 6, 84, 154,
	154, 6, 6, 6, 141, 141, 141, 141, 141, 141,
	49, -1000, 56, 65, -16, 141, -1000, 27, -17, -1000,
	-4, 49, -1000, -1000, -10, 55, 171, 141, 171, -1000,
	-14, -20, 47, -1000, -12, 171, -25, -1000,
}
var yyPgo = []int{

	0, 0, 72, 71, 70, 68, 1, 65, 64,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 2, 2, 2, 3,
	5, 4, 7, 7, 8, 8, 8, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 5, 3, 8,
	11, 7, 1, 3, 0, 1, 3, 1, 1, 2,
	1, 4, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, 7, 9, 8,
	10, 5, 4, 17, 6, 24, -1, -1, -1, -1,
	22, 16, 17, 18, 19, 20, 12, 13, 29, 14,
	30, 15, 4, -6, 4, 24, 24, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	23, 22, 24, -6, -8, -6, 25, -6, -7, 4,
	25, 28, 25, 22, 25, 28, 26, -6, 26, 4,
	-1, -1, 27, 27, 11, 26, -1, 27,
}
var yyDef = []int{

	1, -2, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 17, 18, 0, 20, 0, 2, 3, 4, 5,
	6, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 14, 19, 0, 23,
	24, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	0, 8, 0, 0, 0, 15, 22, 0, 0, 12,
	0, 0, 21, 7, 0, 0, 1, 16, 1, 13,
	0, 0, 11, 9, 0, 1, 0, 10,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 20, 3, 3,
	24, 25, 18, 16, 28, 17, 3, 19, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 22,
	30, 23, 29, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 26, 3, 27,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 21,
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
		//line parser.go.y:40
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:47
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:54
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:61
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:68
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if_else}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:76
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-1].expr}
		}
	case 7:
		//line parser.go.y:80
		{
			yyVAL.stmt = &ast.VarStmt{Name: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expr}
		}
	case 8:
		//line parser.go.y:84
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-1].expr}
		}
	case 9:
		//line parser.go.y:89
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 10:
		//line parser.go.y:94
		{
			yyVAL.stmt_if_else = &ast.IfStmt{Expr: yyS[yypt-8].expr, ThenStmts: yyS[yypt-5].stmts, ElseStmts: yyS[yypt-1].stmts}
		}
	case 11:
		//line parser.go.y:99
		{
			yyVAL.stmt_if = &ast.IfStmt{Expr: yyS[yypt-4].expr, ThenStmts: yyS[yypt-1].stmts}
		}
	case 12:
		//line parser.go.y:104
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 13:
		//line parser.go.y:108
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 14:
		//line parser.go.y:113
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 15:
		//line parser.go.y:117
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 16:
		//line parser.go.y:121
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 17:
		//line parser.go.y:126
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 18:
		//line parser.go.y:130
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 19:
		//line parser.go.y:134
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 20:
		//line parser.go.y:138
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 21:
		//line parser.go.y:142
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 22:
		//line parser.go.y:146
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 23:
		//line parser.go.y:150
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 24:
		//line parser.go.y:154
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 25:
		//line parser.go.y:158
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:162
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 27:
		//line parser.go.y:166
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 28:
		//line parser.go.y:170
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 29:
		//line parser.go.y:174
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 30:
		//line parser.go.y:178
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 31:
		//line parser.go.y:182
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 32:
		//line parser.go.y:186
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 33:
		//line parser.go.y:190
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	}
	goto yystack /* stack new state and value */
}
