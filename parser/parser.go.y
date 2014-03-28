%{
package parser

import (
	"github.com/mattn/anko/ast"
)

%}

%type<stmts> stmts
%type<stmt> stmt
%type<stmt_func> stmt_func
%type<stmt_if> stmt_if
%type<stmt_else> stmt_else
%type<stmt_for> stmt_for
%type<expr> expr
%type<exprs> exprs
%type<pair> pair
%type<pairs> pairs
%type<idents> idents

%union{
	stmt_func    ast.Stmt
	stmt_if      ast.Stmt
	stmt_else    ast.Stmt
	stmt_for     ast.Stmt
	stmts        []ast.Stmt
	stmt         ast.Stmt
	teim         ast.Expr
	expr         ast.Expr
	tok          Token
	idents       []string
	exprs        []ast.Expr
	pair         *ast.PairExpr
	pairs        []*ast.PairExpr
}

%token<tok> IDENT NUMBER STRING ARRAY VAR FUNC RETURN IF ELSE FOR IN EQ NE GE LE

%left '+' '-'
%left '*' '/' '%'
%right UNARY

%%

stmts :
	{
		$$ = nil
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}
	| stmt stmts
	{
		$$ = append([]ast.Stmt{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}
	| stmt_func stmts
	{
		$$ = append([]ast.Stmt{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}
	| stmt_if stmts
	{
		$$ = append([]ast.Stmt{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}
	| stmt_else stmts
	{
		$$ = append([]ast.Stmt{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}
	| stmt_for stmts
	{
		$$ = append([]ast.Stmt{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}

stmt : expr ';'
	{
		$$ = &ast.ExprStmt{Expr: $1}
	}
	| VAR IDENT '=' expr ';'
	{
		$$ = &ast.VarStmt{Name: $2.lit, Expr: $4}
	}
	| RETURN expr ';'
	{
		$$ = &ast.ReturnStmt{Expr: $2}
	}

stmt_func : FUNC IDENT '(' idents ')' '{' stmts '}'
	{
		$$ = &ast.FuncStmt{Name: $2.lit, Args: $4, Stmts: $7}
	}

stmt_else : IF '(' expr ')' '{' stmts '}' ELSE '{' stmts '}'
	{
		$$ = &ast.IfStmt{Expr: $3, ThenStmts: $6, ElseStmts: $10}
	}

stmt_if : IF '(' expr ')' '{' stmts '}'
	{
		$$ = &ast.IfStmt{Expr: $3, ThenStmts: $6}
	}

stmt_for : FOR IDENT IN expr '{' stmts '}'
	{
		$$ = &ast.ForStmt{Var: $2.lit, Value: $4, Stmts: $6}
	}

pair : STRING ':' expr
	{
		$$ = &ast.PairExpr{Key: $1.lit, Value: $3}
	}

pairs :
	{
		$$ = []*ast.PairExpr{}
	}
	| pair
	{
		$$ = []*ast.PairExpr{$1}
	}
	| pairs ',' pair
	{
		$$ = append($1, $3)
	}

idents :
	{
		$$ = []string{}
	}
	| IDENT
	{
		$$ = []string{$1.lit}
	}
	| idents ',' IDENT
	{
		$$ = append($1, $3.lit)
	}

exprs :
	{
		$$ = []ast.Expr{}
	}
	| expr
	{
		$$ = []ast.Expr{$1}
	}
	| exprs ',' expr
	{
		$$ = append($1, $3)
	}

expr : NUMBER
	{
		$$ = &ast.NumberExpr{Lit: $1.lit}
	}
	| IDENT
	{
		$$ = &ast.IdentExpr{Lit: $1.lit}
	}
	| '-' expr %prec UNARY
	{
		$$ = &ast.UnaryMinusExpr{SubExpr: $2}
	}
	| STRING
	{
		$$ = &ast.StringExpr{Lit: $1.lit}
	}
	| IDENT '(' exprs ')'
	{
		$$ = &ast.CallExpr{Name: $1.lit, SubExprs: $3}
	}
	| '[' exprs ']'
	{
		$$ = &ast.ArrayExpr{Exprs: $2}
	}
	| FUNC '(' idents ')' '{' stmts '}'
	{
		$$ = &ast.FuncExpr{Args: $3, Stmts: $6}
	}

	| '{' pairs '}'
	{
		mapExpr := make(map[string]ast.Expr)
		for _, v := range $2 {
			mapExpr[v.Key] = v.Value
		}
		$$ = &ast.MapExpr{MapExpr: mapExpr}
	}
	| expr '[' expr ']'
	{
		$$ = &ast.ItemExpr{Value: $1, Index: $3}
	}
	| '(' expr ')'
	{
		$$ = &ast.ParenExpr{SubExpr: $2}
	}
	| expr '+' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "+", Rhs: $3}
	}
	| expr '-' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "-", Rhs: $3}
	}
	| expr '*' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "*", Rhs: $3}
	}
	| expr '/' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "/", Rhs: $3}
	}
	| expr '%' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "%", Rhs: $3}
	}
	| expr EQ expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "==", Rhs: $3}
	}
	| expr NE expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "!=", Rhs: $3}
	}
	| expr '>' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: ">", Rhs: $3}
	}
	| expr GE expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: ">=", Rhs: $3}
	}
	| expr '<' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "<", Rhs: $3}
	}
	| expr LE expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "<=", Rhs: $3}
	}

%%
