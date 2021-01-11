package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // 未知のtoken文字列
	EOF = "EOF"

	// 識別子 + リテラル
	INDENT = "INDENT"
	INT = "INT"

	// 演算子
	ASSIGN = "="
	PLUS = "+"

	// デミリタ
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET = "LET"


)