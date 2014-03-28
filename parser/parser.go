//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:21
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
	pair         *ast.PairExpr
	pairs        []*ast.PairExpr
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VAR = 57350
const FUNC = 57351
const RETURN = 57352
const IF = 57353
const ELSE = 57354
const EQ = 57355
const NE = 57356
const GE = 57357
const LE = 57358
const UNARY = 57359

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
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

//line parser.go.y:228

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 40
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 207

var yyAct = []int{

	1, 6, 43, 18, 19, 20, 21, 66, 87, 62,
	35, 40, 63, 77, 89, 39, 78, 41, 72, 45,
	80, 62, 65, 64, 85, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 56, 84, 79, 59, 60,
	41, 28, 29, 31, 33, 23, 24, 25, 26, 27,
	61, 76, 38, 86, 37, 57, 44, 81, 70, 68,
	69, 30, 32, 36, 73, 34, 42, 74, 75, 28,
	29, 31, 33, 23, 24, 25, 26, 27, 5, 4,
	82, 83, 71, 3, 2, 0, 0, 0, 88, 30,
	32, 28, 29, 31, 33, 23, 24, 25, 26, 27,
	0, 0, 0, 0, 67, 28, 29, 31, 33, 0,
	0, 30, 32, 28, 29, 31, 33, 23, 24, 25,
	26, 27, 0, 58, 0, 30, 32, 0, 0, 0,
	0, 0, 0, 30, 32, 28, 29, 31, 33, 23,
	24, 25, 26, 27, 0, 22, 28, 29, 31, 33,
	23, 24, 25, 26, 27, 30, 32, 28, 29, 31,
	33, 0, 0, 25, 26, 27, 30, 32, 12, 11,
	14, 0, 7, 9, 8, 10, 0, 30, 32, 12,
	11, 14, 13, 0, 0, 0, 0, 0, 0, 17,
	0, 16, 0, 13, 0, 15, 0, 0, 0, 0,
	17, 0, 16, 0, 0, 0, 15,
}
var yyPact = []int{

	164, -1000, 164, 164, 164, 164, 122, 61, 175, 59,
	29, -1000, 27, 175, -1000, 175, 50, 175, -1000, -1000,
	-1000, -1000, -1000, 175, 175, 175, 175, 175, 175, 175,
	175, 175, 175, 175, 31, 100, 13, 175, 175, 92,
	-20, 133, -6, -1000, -23, 78, 144, 144, 92, 92,
	92, 133, 133, 133, 133, 133, 133, 175, -1000, 54,
	56, -8, 175, -1000, 50, -1000, 175, -1000, 28, -13,
	-1000, 10, -1000, 133, -1000, 133, -1000, -7, 53, 164,
	164, -1000, 8, -4, 41, -1000, -19, 164, -14, -1000,
}
var yyPgo = []int{

	0, 0, 84, 83, 79, 78, 1, 11, 2, 66,
	60,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 2, 2, 2, 3,
	5, 4, 10, 10, 8, 9, 9, 9, 7, 7,
	7, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 5, 3, 8,
	11, 7, 1, 3, 3, 0, 1, 3, 0, 1,
	3, 1, 1, 2, 1, 4, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, 8, 10, 9,
	11, 5, 4, 18, 6, 31, 27, 25, -1, -1,
	-1, -1, 23, 17, 18, 19, 20, 21, 13, 14,
	33, 15, 34, 16, 4, -6, 4, 25, 25, -6,
	-7, -6, -9, -8, 6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, 24, 23, 25,
	-6, -7, 29, 32, 29, 28, 30, 26, -6, -10,
	4, 26, 26, -6, -8, -6, 23, 26, 29, 27,
	27, 4, -1, -1, 28, 28, 12, 27, -1, 28,
}
var yyDef = []int{

	1, -2, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 21, 22, 0, 24, 18, 15, 0, 2, 3,
	4, 5, 6, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 18, 23,
	0, 19, 0, 16, 0, 0, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 0, 8, 0,
	0, 0, 0, 26, 0, 27, 0, 28, 0, 0,
	12, 0, 25, 20, 17, 14, 7, 0, 0, 1,
	1, 13, 0, 0, 11, 9, 0, 1, 0, 10,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 21, 3, 3,
	25, 26, 19, 17, 29, 18, 3, 20, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 30, 23,
	34, 24, 33, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 31, 3, 32, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 27, 3, 28,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 22,
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
		//line parser.go.y:44
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:51
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:58
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:65
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:72
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if_else}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:80
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-1].expr}
		}
	case 7:
		//line parser.go.y:84
		{
			yyVAL.stmt = &ast.VarStmt{Name: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expr}
		}
	case 8:
		//line parser.go.y:88
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-1].expr}
		}
	case 9:
		//line parser.go.y:93
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 10:
		//line parser.go.y:98
		{
			yyVAL.stmt_if_else = &ast.IfStmt{Expr: yyS[yypt-8].expr, ThenStmts: yyS[yypt-5].stmts, ElseStmts: yyS[yypt-1].stmts}
		}
	case 11:
		//line parser.go.y:103
		{
			yyVAL.stmt_if = &ast.IfStmt{Expr: yyS[yypt-4].expr, ThenStmts: yyS[yypt-1].stmts}
		}
	case 12:
		//line parser.go.y:108
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 13:
		//line parser.go.y:112
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 14:
		//line parser.go.y:117
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 15:
		//line parser.go.y:122
		{
			yyVAL.pairs = []*ast.PairExpr{}
		}
	case 16:
		//line parser.go.y:126
		{
			yyVAL.pairs = []*ast.PairExpr{yyS[yypt-0].pair}
		}
	case 17:
		//line parser.go.y:130
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 18:
		//line parser.go.y:135
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 19:
		//line parser.go.y:139
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 20:
		//line parser.go.y:143
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 21:
		//line parser.go.y:148
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 22:
		//line parser.go.y:152
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 23:
		//line parser.go.y:156
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 24:
		//line parser.go.y:160
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 25:
		//line parser.go.y:164
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 26:
		//line parser.go.y:168
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 27:
		//line parser.go.y:172
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.Key] = v.Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 28:
		//line parser.go.y:180
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 29:
		//line parser.go.y:184
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 30:
		//line parser.go.y:188
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 31:
		//line parser.go.y:192
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 32:
		//line parser.go.y:196
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 33:
		//line parser.go.y:200
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 34:
		//line parser.go.y:204
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 35:
		//line parser.go.y:208
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 36:
		//line parser.go.y:212
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 37:
		//line parser.go.y:216
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 38:
		//line parser.go.y:220
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 39:
		//line parser.go.y:224
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	}
	goto yystack /* stack new state and value */
}
