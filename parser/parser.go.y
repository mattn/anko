%{
package parser

import (
	"github.com/mattn/anko/ast"
)

%}

%type<stmts> stmts
%type<stmt> stmt
%type<stmt_if> stmt_if
%type<stmt_default> stmt_default
%type<stmt_case> stmt_case
%type<stmt_cases> stmt_cases
%type<expr> expr
%type<exprs> exprs
%type<expr_lhs> expr_lhs
%type<expr_lhss> expr_lhss
%type<expr_lets> expr_lets
%type<expr_pair> expr_pair
%type<expr_pairs> expr_pairs
%type<expr_idents> expr_idents

%union{
	stmt_if                ast.Stmt
	stmt_default           ast.Stmt
	stmt_case              ast.Stmt
	stmt_cases             []ast.Stmt
	stmts                  []ast.Stmt
	stmt                   ast.Stmt
	expr                   ast.Expr
	exprs                  []ast.Expr
	expr_lhs               ast.Expr
	expr_lhss              []ast.Expr
	expr_lets              ast.Expr
	expr_pair              ast.Expr
	expr_pairs             []ast.Expr
	expr_idents            []string
	tok                    ast.Token
}

%token<tok> IDENT NUMBER STRING ARRAY VARARG FUNC RETURN VAR THROW IF ELSE FOR IN EQEQ NEQ GE LE OROR ANDAND NEW TRUE FALSE NIL MODULE TRY CATCH FINALLY PLUSEQ MINUSEQ MULEQ DIVEQ ANDEQ OREQ BREAK CONTINUE PLUSPLUS MINUSMINUS POW SHIFTLEFT SHIFTRIGHT SWITCH CASE DEFAULT

%right '='
%right '?' ':'
%left OROR
%left ANDAND
%left IDENT
%nonassoc EQEQ NEQ ','
%left '>' GE '<' LE SHIFTLEFT SHIFTRIGHT

%left '+' '-' PLUSPLUS MINUSMINUS
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
			for _, s := range $2 {
				if $1.Position().Line == s.Position().Line {
					l.pos = $1.Position()
					yylex.Error("syntax error")
				}
			}
		}
	}
	| stmt ';' stmts
	{
		$$ = append([]ast.Stmt{$1}, $3...)
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
			$1.SetPosition(l.pos);
		}
	}

stmt : expr
	{
		$$ = &ast.ExprStmt{Expr: $1}
		$$.SetPosition($1.Position())
	}
	| BREAK
	{
		$$ = &ast.BreakStmt{}
		$$.SetPosition($1.Position())
	}
	| CONTINUE
	{
		$$ = &ast.ContinueStmt{}
		$$.SetPosition($1.Position())
	}
	| RETURN exprs
	{
		$$ = &ast.ReturnStmt{Exprs: $2}
		$$.SetPosition($1.Position())
	}
	| THROW expr
	{
		$$ = &ast.ThrowStmt{Expr: $2}
		$$.SetPosition($1.Position())
	}
	| MODULE IDENT '{' stmts '}'
	{
		$$ = &ast.ModuleStmt{Name: $2.Lit, Stmts: $4}
		$$.SetPosition($1.Position())
	}
	| VAR expr_idents '=' exprs
	{
		$$ = &ast.VarStmt{Names: $2, Exprs: $4}
		$$.SetPosition($1.Position())
	}
	| stmt_if
	{
		$$ = $1
		$$.SetPosition($1.Position())
	}
	| FOR IDENT IN expr '{' stmts '}'
	{
		$$ = &ast.ForStmt{Var: $2.Lit, Value: $4, Stmts: $6}
		$$.SetPosition($1.Position())
	}
	| FOR '{' stmts '}'
	{
		$$ = &ast.LoopStmt{Stmts: $3}
		$$.SetPosition($1.Position())
	}
	| FOR expr '{' stmts '}'
	{
		$$ = &ast.LoopStmt{Expr: $2, Stmts: $4}
		$$.SetPosition($1.Position())
	}
	| FOR expr_lets ';' expr ';' expr '{' stmts '}' 
	{
		$$ = &ast.CForStmt{Expr1: $2, Expr2: $4, Expr3: $6, Stmts: $8}
		$$.SetPosition($1.Position())
	}
	| TRY '{' stmts '}' CATCH IDENT '{' stmts '}' FINALLY '{' stmts '}'
	{
		$$ = &ast.TryStmt{Try: $3, Var: $6.Lit, Catch: $8, Finally: $12}
		$$.SetPosition($1.Position())
	}
	| TRY '{' stmts '}' CATCH '{' stmts '}' FINALLY '{' stmts '}'
	{
		$$ = &ast.TryStmt{Try: $3, Catch: $7, Finally: $11}
		$$.SetPosition($1.Position())
	}
	| TRY '{' stmts '}' CATCH IDENT '{' stmts '}'
	{
		$$ = &ast.TryStmt{Try: $3, Var: $6.Lit, Catch: $8}
		$$.SetPosition($1.Position())
	}
	| TRY '{' stmts '}' CATCH '{' stmts '}'
	{
		$$ = &ast.TryStmt{Try: $3, Catch: $7}
		$$.SetPosition($1.Position())
	}
	| SWITCH expr '{' stmt_cases '}'
	{
		$$ = &ast.SwitchStmt{Expr: $2, Cases: $4}
		$$.SetPosition($1.Position())
	}
	| expr_lets
	{
		$$ = &ast.ExprStmt{Expr: $1}
		$$.SetPosition($1.Position())
	}

stmt_if : stmt_if ELSE IF expr '{' stmts '}'
	{
		$1.(*ast.IfStmt).ElseIf = append($1.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $4, Then: $6})
	}
	| stmt_if ELSE '{' stmts '}'
	{
		if $$.(*ast.IfStmt).Else != nil {
			yylex.Error("multiple else statement")
		} else {
			$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
		}
	}
	| IF expr '{' stmts '}'
	{
		$$ = &ast.IfStmt{If: $2, Then: $4, Else: nil}
	}

stmt_cases :
	{
		$$ = []ast.Stmt{}
	}
	| stmt_case
	{
		$$ = []ast.Stmt{$1}
	}
	| stmt_cases stmt_case
	{
		$$ = append($1, $2)
	}
	| stmt_default
	{
		$$ = []ast.Stmt{$1}
	}
	| stmt_cases stmt_default
	{
		for _, stmt := range $1 {
			if _, ok := stmt.(*ast.DefaultStmt); ok {
				yylex.Error("multiple default statement")
			}
		}
		$$ = append($1, $2)
	}

stmt_case : CASE expr ':' stmts
	{
		$$ = &ast.CaseStmt{Expr: $2, Stmts: $4}
	}

stmt_default : DEFAULT ':' stmts
	{
		$$ = &ast.DefaultStmt{Stmts: $3}
	}

expr_pair : STRING ':' expr
	{
		$$ = &ast.PairExpr{Key: $1.Lit, Value: $3}
	}

expr_pairs :
	{
		$$ = []ast.Expr{}
	}
	| expr_pair
	{
		$$ = []ast.Expr{$1}
	}
	| expr_pairs ',' expr_pair
	{
		$$ = append($1, $3)
	}

expr_idents :
	{
		$$ = []string{}
	}
	| IDENT
	{
		$$ = []string{$1.Lit}
	}
	| expr_idents ',' IDENT
	{
		$$ = append($1, $3.Lit)
	}

expr_lhs : IDENT
	{
		$$ = &ast.IdentExpr{Lit: $1.Lit}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| expr '.' IDENT
	{
		$$ = &ast.MemberExpr{Expr: $1, Name: $3.Lit}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| expr '[' expr ']'
	{
		$$ = &ast.ItemExpr{Value: $1, Index: $3}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}

expr_lets : expr_lhss '=' exprs
	{
		$$ = &ast.LetsExpr{Lhss: $1, Operator: "=", Rhss: $3}
		$$.SetPosition($1[0].Position())
	}

expr_lhss :
		  /*
	{
		$$ = []ast.Expr{}
	}
	*/
	expr_lhs
	{
		$$ = []ast.Expr{$1}
	}
	| expr_lhs ',' expr_lhss
	{
		$$ = append([]ast.Expr{$1}, $3...)
	}

exprs :
	{
		$$ = []ast.Expr{}
	}
	| expr
	{
		$$ = []ast.Expr{$1}
	}
	| expr ',' exprs
	{
		$$ = append([]ast.Expr{$1}, $3...)
	}
	| IDENT ',' exprs
	{
		$$ = append([]ast.Expr{&ast.IdentExpr{Lit: $1.Lit}}, $3...)
	}

expr :
	IDENT
	{
		$$ = &ast.IdentExpr{Lit: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| NUMBER
	{
		$$ = &ast.NumberExpr{Lit: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| '-' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "-", Expr: $2}
		$$.SetPosition($2.Position())
	}
	| '!' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "!", Expr: $2}
		$$.SetPosition($2.Position())
	}
	| '^' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "^", Expr: $2}
		$$.SetPosition($2.Position())
	}
	| '&' IDENT %prec UNARY
	{
		$$ = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: $2.Lit}}
		$$.SetPosition($2.Position())
	}
	| '&' expr '.' IDENT %prec UNARY
	{
		$$ = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: $2, Name: $4.Lit}}
		$$.SetPosition($2.Position())
	}
	| '*' IDENT %prec UNARY
	{
		$$ = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: $2.Lit}}
		$$.SetPosition($2.Position())
	}
	| '*' expr '.' IDENT %prec UNARY
	{
		$$ = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: $2, Name: $4.Lit}}
		$$.SetPosition($2.Position())
	}
	| STRING
	{
		$$ = &ast.StringExpr{Lit: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| TRUE
	{
		$$ = &ast.ConstExpr{Value: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| FALSE
	{
		$$ = &ast.ConstExpr{Value: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| NIL
	{
		$$ = &ast.ConstExpr{Value: $1.Lit}
		$$.SetPosition($1.Position())
	}
	| expr '?' expr ':' expr
	{
		$$ = &ast.TernaryOpExpr{Expr: $1, Lhs: $3, Rhs: $5}
		$$.SetPosition($1.Position())
	}
	| expr '.' IDENT
	{
		$$ = &ast.MemberExpr{Expr: $1, Name: $3.Lit}
		$$.SetPosition($1.Position())
	}
	| FUNC '(' expr_idents ')' '{' stmts '}'
	{
		$$ = &ast.FuncExpr{Args: $3, Stmts: $6}
		$$.SetPosition($1.Position())
	}
	| FUNC '(' IDENT VARARG ')' '{' stmts '}'
	{
		$$ = &ast.FuncExpr{Args: []string{$3.Lit}, Stmts: $7, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| FUNC IDENT '(' expr_idents ')' '{' stmts '}'
	{
		$$ = &ast.FuncExpr{Name: $2.Lit, Args: $4, Stmts: $7}
		$$.SetPosition($1.Position())
	}
	| FUNC IDENT '(' IDENT VARARG ')' '{' stmts '}'
	{
		$$ = &ast.FuncExpr{Name: $2.Lit, Args: []string{$4.Lit}, Stmts: $8, VarArg: true}
		$$.SetPosition($1.Position())
	}
	| '[' exprs ']'
	{
		$$ = &ast.ArrayExpr{Exprs: $2}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '{' expr_pairs '}'
	{
		mapExpr := make(map[string]ast.Expr)
		for _, v := range $2 {
			mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
		}
		$$ = &ast.MapExpr{MapExpr: mapExpr}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| '(' expr ')'
	{
		$$ = &ast.ParenExpr{SubExpr: $2}
		if l, ok := yylex.(*Lexer); ok { $$.SetPosition(l.pos) }
	}
	| NEW IDENT '(' exprs ')'
	{
		$$ = &ast.NewExpr{Name: $2.Lit, SubExprs: $4}
		$$.SetPosition($1.Position())
	}
	| expr '+' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "+", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '-' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "-", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '*' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "*", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '/' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "/", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '%' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "%", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr POW expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "**", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr SHIFTLEFT expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "<<", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr SHIFTRIGHT expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: ">>", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr EQEQ expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "==", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr NEQ expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "!=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '>' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: ">", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr GE expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: ">=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '<' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "<", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr LE expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "<=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr PLUSEQ expr
	{
		$$ = &ast.AssocExpr{Lhs: $1, Operator: "+=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr MINUSEQ expr
	{
		$$ = &ast.AssocExpr{Lhs: $1, Operator: "-=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr MULEQ expr
	{
		$$ = &ast.AssocExpr{Lhs: $1, Operator: "*=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr DIVEQ expr
	{
		$$ = &ast.AssocExpr{Lhs: $1, Operator: "/=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr ANDEQ expr
	{
		$$ = &ast.AssocExpr{Lhs: $1, Operator: "&=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr OREQ expr
	{
		$$ = &ast.AssocExpr{Lhs: $1, Operator: "|=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '=' expr
	{
		$$ = &ast.AssocExpr{Lhs: $1, Operator: "=", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr PLUSPLUS
	{
		$$ = &ast.AssocExpr{Lhs: $1, Operator: "++"}
		$$.SetPosition($1.Position())
	}
	| expr MINUSMINUS
	{
		$$ = &ast.AssocExpr{Lhs: $1, Operator: "--"}
		$$.SetPosition($1.Position())
	}
	| expr '|' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "|", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr OROR expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "||", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr '&' expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "&", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| expr ANDAND expr
	{
		$$ = &ast.BinOpExpr{Lhs: $1, Operator: "&&", Rhs: $3}
		$$.SetPosition($1.Position())
	}
	| IDENT '(' exprs ')'
	{
		$$ = &ast.CallExpr{Name: $1.Lit, SubExprs: $3}
		$$.SetPosition($1.Position())
	}
	| IDENT '[' expr ']'
	{
		$$ = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: $1.Lit}, Index: $3}
		$$.SetPosition($1.Position())
	}
	| expr '[' expr ']'
	{
		$$ = &ast.ItemExpr{Value: $1, Index: $3}
		$$.SetPosition($1.Position())
	}
	| expr '(' exprs  ')'
	{
		$$ = &ast.AnonCallExpr{Expr: $1, SubExprs: $3}
		$$.SetPosition($1.Position())
	}

%%
