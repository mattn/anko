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
%type<stmt_for> stmt_for
%type<stmt_try_catch_finally> stmt_try_catch_finally
%type<expr> expr
%type<exprs> exprs
%type<pair> pair
%type<pairs> pairs
%type<idents> idents

%union{
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

%token<tok> IDENT NUMBER STRING ARRAY VARARG FUNC RETURN THROW IF ELSE FOR IN EQ NE GE LE OR AND NEW TRUE FALSE NIL MODULE TRY CATCH FINALLY

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
	| stmt_for stmts
	{
		$$ = append([]ast.Stmt{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}
	| stmt_try_catch_finally stmts
	{
		$$ = append([]ast.Stmt{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}

stmt : expr
	{
		$$ = &ast.ExprStmt{Expr: $1}
		if l, ok := yylex.(*Lexer); ok {
			$1.SetPos(l.pos)
		}
	}
	| idents '=' exprs
	{
		$$ = &ast.LetStmt{Names: $1, Exprs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| RETURN expr
	{
		$$ = &ast.ReturnStmt{Expr: $2}
		if l, ok := yylex.(*Lexer); ok {
			$2.SetPos(l.pos)
		}
	}
	| THROW expr
	{
		$$ = &ast.ThrowStmt{Expr: $2}
		if l, ok := yylex.(*Lexer); ok {
			$2.SetPos(l.pos)
		}
	}
	| MODULE IDENT '{' stmts '}'
	{
		$$ = &ast.ModuleStmt{Name: $2.lit, Stmts: $4}
	}

stmt_func : FUNC IDENT '(' idents ')' '{' stmts '}'
	{
		$$ = &ast.FuncStmt{Name: $2.lit, Args: $4, Stmts: $7}
	}
	| FUNC IDENT '(' IDENT VARARG ')' '{' stmts '}'
	{
		$$ = &ast.FuncStmt{Name: $2.lit, Args: []string{$4.lit}, Stmts: $8, VarArg: true}
	}

stmt_if : IF '(' expr ')' '{' stmts '}'
	{
		$$ = &ast.IfStmt{If: $3, Then: $6}
	}
	| stmt_if ELSE '{' stmts '}'
	{
		if $1.(*ast.IfStmt).Else != nil {
			yylex.Error("multiple else statement")
		} else {
			$1.(*ast.IfStmt).Else = $4
		}
	}
	| stmt_if ELSE IF '(' expr ')' '{' stmts '}'
	{
		$1.(*ast.IfStmt).ElseIf = append($1.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $5, Then: $8})
	}

stmt_for : FOR IDENT IN expr '{' stmts '}'
	{
		$$ = &ast.ForStmt{Var: $2.lit, Value: $4, Stmts: $6}
	}

stmt_try_catch_finally : TRY '{' stmts '}' CATCH '(' IDENT ')' '{' stmts '}' FINALLY '{' stmts '}'
	{
		$$ = &ast.TryStmt{Try: $3, Var: $7.lit, Catch: $10, Finally: $14}
	}
	| TRY '{' stmts '}' CATCH '{' stmts '}' FINALLY '{' stmts '}'
	{
		$$ = &ast.TryStmt{Try: $3, Catch: $7, Finally: $11}
	}
	| TRY '{' stmts '}' CATCH '(' IDENT ')' '{' stmts '}'
	{
		$$ = &ast.TryStmt{Try: $3, Var: $7.lit, Catch: $10}
	}
	| TRY '{' stmts '}' CATCH '{' stmts '}'
	{
		$$ = &ast.TryStmt{Try: $3, Catch: $7}
	}

pair : STRING ':' expr
	{
		$$ = &ast.PairExpr{Key: $1.lit, Value: $3}
	}

pairs :
	{
		$$ = []ast.Expr{}
	}
	| pair
	{
		$$ = []ast.Expr{$1}
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
	| FUNC '(' idents ')' '{' stmts '}'
	{
		$$ = &ast.FuncExpr{Args: $3, Stmts: $6}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '(' exprs  ')'
	{
		$$ = &ast.AnonCallExpr{Expr: $1, SubExprs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '.' IDENT
	{
		$$ = &ast.MemberExpr{Expr: $1, Method: $3.lit}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| IDENT '(' exprs ')'
	{
		$$ = &ast.CallExpr{Name: $1.lit, SubExprs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| '[' exprs ']'
	{
		$$ = &ast.ArrayExpr{Exprs: $2}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| FUNC '(' IDENT VARARG ')' '{' stmts '}'
	{
		$$ = &ast.FuncExpr{Args: []string{$3.lit}, Stmts: $7, VarArg: true}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| '{' pairs '}'
	{
		mapExpr := make(map[string]ast.Expr)
		for _, v := range $2 {
			mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
		}
		$$ = &ast.MapExpr{MapExpr: mapExpr}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '[' expr ']'
	{
		$$ = &ast.ItemExpr{Value: $1, Index: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| '(' expr ')'
	{
		$$ = &ast.ParenExpr{SubExpr: $2}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| NEW IDENT '(' exprs ')'
	{
		$$ = &ast.NewExpr{Name: $2.lit, SubExprs: $4}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| TRUE
	{
		$$ = &ast.ConstExpr{Value: true}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| FALSE
	{
		$$ = &ast.ConstExpr{Value: false}
	}
	| NIL
	{
		$$ = &ast.ConstExpr{Value: nil}
	}
	| expr '+' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "+", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '-' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "-", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '*' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "*", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '/' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "/", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '%' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "%", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr EQ expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "==", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr NE expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "!=", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '>' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: ">", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr GE expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: ">=", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '<' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "<", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr LE expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "<=", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| IDENT '=' expr
	{
		$$ = &ast.LetExpr{Name: $1.lit, Expr: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '|' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "|", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr OR expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "||", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr '&' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "&", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}
	| expr AND expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "&&", Rhs: $3}
		if l, ok := yylex.(*Lexer); ok {
			$$.SetPos(l.pos)
		}
	}

%%
