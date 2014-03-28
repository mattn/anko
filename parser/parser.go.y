%{
package parser

import (
	"github.com/mattn/ungo/ast"
)

%}

%type<stmts> stmts
%type<stmt> stmt
%type<funct> funct
%type<expr> expr
%type<idents> idents
%type<exprs> exprs

%union{
	funct  ast.Stmt
	stmts  []ast.Stmt
	stmt   ast.Stmt
	expr   ast.Expr
	tok    Token
	idents []string
	exprs  []ast.Expr
}

%token<tok> IDENT NUMBER STRING VAR FUNC RETURN

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
	| funct stmts
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

funct : FUNC IDENT '(' idents ')' '{' stmts '}'
	{
		$$ = &ast.FuncStmt{Name: $2.lit, Args: $4, Stmts: $7}
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
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: int('+'), Rhs: $3}
	}
	| expr '-' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: int('-'), Rhs: $3}
	}
	| expr '*' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: int('*'), Rhs: $3}
	}
	| expr '/' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: int('/'), Rhs: $3}
	}
	| expr '%' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: int('%'), Rhs: $3}
	}

%%
