package asm8

import (
	"io"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/lex8"
)

func listImport(f string, rc io.ReadCloser, pkg build8.Pkg) []*lex8.Error {
	astFile, es := parse.File(f, rc)
	if es != nil {
		return es
	}

	if astFile.Imports == nil {
		return nil
	}

	res := newResolver()
	imp := resolveImportDecl(res, astFile.Imports)
	if es := res.Errs(); es != nil {
		return es
	}

	for as, stmt := range imp.stmts {
		pkg.AddImport(as, stmt.path, stmt.Path.Pos)
	}

	return nil
}
