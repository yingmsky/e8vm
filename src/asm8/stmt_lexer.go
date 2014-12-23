package asm8

import (
	"io"
)

// StmtLexer replaces end-lines with semicolons
type StmtLexer struct {
	x          *Lexer
	save       *Token
	insertSemi bool
}

// NewStmtLexer creates a new statement lexer.
func NewStmtLexer(file string, r io.ReadCloser) *StmtLexer {
	ret := new(StmtLexer)
	ret.x = NewLexer(file, r)

	return ret
}

// Token returns the next token of lexing
func (sx *StmtLexer) Token() *Token {
	if sx.save != nil {
		ret := sx.save
		sx.save = nil
		return ret
	}

	for {
		t := sx.x.Token()
		switch t.Type {
		case Lbrace, Semi:
			sx.insertSemi = false
		case EOF:
			if sx.insertSemi {
				sx.insertSemi = false
				sx.save = t
				return &Token{Semi, t.Lit, t.Pos}
			}
		case Rbrace:
			if sx.insertSemi {
				sx.save = t
				return &Token{Semi, t.Lit, t.Pos}
			}
			sx.insertSemi = true
		case Endl:
			if sx.insertSemi {
				sx.insertSemi = false
				return &Token{Semi, "\n", t.Pos}
			}
			continue // ignore this end line
		case Comment:
			// do nothing
		default:
			sx.insertSemi = true
		}

		return t
	}
}

// Errs returns the list of lexing errors.
func (sx *StmtLexer) Errs() []*Error {
	return sx.x.Errs()
}
