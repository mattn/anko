//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/ungo/ast"
)

//line parser.go.y:17
type yySymType struct {
	yys    int
	funct  ast.Stmt
	stmts  []ast.Stmt
	stmt   ast.Stmt
	expr   ast.Expr
	tok    Token
	idents []string
	exprs  []ast.Expr
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const VAR = 57349
const FUNC = 57350
const RETURN = 57351
const UNARY = 57352

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"VAR",
	"FUNC",
	"RETURN",
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

//line parser.go.y:141

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 24
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 80

var yyAct = []int{

	4, 1, 50, 44, 13, 14, 45, 22, 47, 42,
	34, 25, 41, 26, 24, 32, 48, 27, 28, 29,
	30, 31, 18, 19, 20, 36, 40, 9, 8, 11,
	5, 7, 6, 38, 10, 16, 17, 18, 19, 20,
	23, 12, 46, 21, 37, 9, 8, 11, 35, 49,
	39, 3, 10, 2, 16, 17, 18, 19, 20, 12,
	43, 16, 17, 18, 19, 20, 0, 33, 16, 17,
	18, 19, 20, 0, 15, 16, 17, 18, 19, 20,
}
var yyPact = []int{

	23, -1000, 23, 23, 58, 39, 41, 36, -1000, -4,
	41, -1000, 41, -1000, -1000, -1000, 41, 41, 41, 41,
	41, -2, 51, -8, 41, -1000, 25, 10, 10, -1000,
	-1000, -1000, 41, -1000, 22, -10, 65, -1000, 44, -16,
	-1000, 41, -1000, -1000, -12, 12, 65, 23, -1000, -19,
	-1000,
}
var yyPgo = []int{

	0, 1, 53, 51, 0, 50, 48,
}
var yyR1 = []int{

	0, 1, 1, 1, 2, 2, 2, 3, 5, 5,
	6, 6, 6, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 5, 3, 8, 1, 3,
	0, 1, 3, 1, 1, 2, 1, 4, 3, 3,
	3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 7, 9, 8, 5, 4,
	11, 6, 18, -1, -1, 16, 10, 11, 12, 13,
	14, 4, -4, 4, 18, -4, -4, -4, -4, -4,
	-4, -4, 17, 16, 18, -6, -4, 19, -4, -5,
	4, 22, 19, 16, 19, 22, -4, 20, 4, -1,
	21,
}
var yyDef = []int{

	1, -2, 1, 1, 0, 0, 0, 0, 13, 14,
	0, 16, 0, 2, 3, 4, 0, 0, 0, 0,
	0, 0, 0, 0, 10, 15, 0, 19, 20, 21,
	22, 23, 0, 6, 0, 0, 11, 18, 0, 0,
	8, 0, 17, 5, 0, 0, 12, 1, 9, 0,
	7,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 14, 3, 3,
	18, 19, 12, 10, 22, 11, 3, 13, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 16,
	3, 17, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 20, 3, 21,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 15,
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
		//line parser.go.y:36
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:43
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:50
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].funct}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:58
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-1].expr}
		}
	case 5:
		//line parser.go.y:62
		{
			yyVAL.stmt = &ast.VarStmt{Name: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expr}
		}
	case 6:
		//line parser.go.y:66
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-1].expr}
		}
	case 7:
		//line parser.go.y:71
		{
			yyVAL.funct = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 8:
		//line parser.go.y:76
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 9:
		//line parser.go.y:80
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 10:
		//line parser.go.y:85
		{
		}
	case 11:
		//line parser.go.y:88
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 12:
		//line parser.go.y:92
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 13:
		//line parser.go.y:97
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 14:
		//line parser.go.y:101
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 15:
		//line parser.go.y:105
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 16:
		//line parser.go.y:109
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 17:
		//line parser.go.y:113
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 18:
		//line parser.go.y:117
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 19:
		//line parser.go.y:121
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: int('+'), Rhs: yyS[yypt-0].expr}
		}
	case 20:
		//line parser.go.y:125
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: int('-'), Rhs: yyS[yypt-0].expr}
		}
	case 21:
		//line parser.go.y:129
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: int('*'), Rhs: yyS[yypt-0].expr}
		}
	case 22:
		//line parser.go.y:133
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: int('/'), Rhs: yyS[yypt-0].expr}
		}
	case 23:
		//line parser.go.y:137
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: int('%'), Rhs: yyS[yypt-0].expr}
		}
	}
	goto yystack /* stack new state and value */
}
