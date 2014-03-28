%{
package parser

import (
	"github.com/mattn/ungo/ast"
)

%}

%type<stmts> stmts
%type<stmt> stmt
%type<stmt_func> stmt_func
%type<stmt_if> stmt_if
%type<stmt_if_else> stmt_if_else
%type<expr> expr
%type<idents> idents
%type<exprs> exprs

%union{
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

%token<tok> IDENT NUMBER STRING VAR FUNC RETURN IF ELSE EQ NE GE LE

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
	| stmt_if_else stmts
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

stmt_if_else : IF '(' expr ')' '{' stmts '}' ELSE '{' stmts '}'
	{
		$$ = &ast.IfStmt{Expr: $3, ThenStmts: $6, ElseStmts: $10}
	}

stmt_if : IF '(' expr ')' '{' stmts '}'
	{
		$$ = &ast.IfStmt{Expr: $3, ThenStmts: $6}
	}

idents : IDENT
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
