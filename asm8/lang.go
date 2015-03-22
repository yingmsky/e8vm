package asm8

import (
	"io"
	"path"
	"strings"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/pkg8"
)

type lang struct{}

func (lang) IsSrc(filename string) bool {
	return strings.HasSuffix(filename, ".s")
}

func (lang) ListImport(src pkg8.Files) ([]string, []*lex8.Error) {
	for f, rc := range src {
		if len(src) == 1 || path.Base(f) == "import.s" {
			return listImport(f, rc)
		}
	}
	return nil, nil
}

func (lang) Compile(
	p string,
	src pkg8.Files,
	importer pkg8.Importer,
) (
	pkg8.Linkable,
	[]*lex8.Error,
) {
	pkg := new(pkg)
	pkg.path = p

	errs := lex8.NewErrorList()

	for f, rc := range src {
		astFile, es := parse.File(f, rc)
		if es != nil {
			return nil, es
		}

		file := resolveFile(errs, astFile)
		pkg.files = append(pkg.files, file)

		if len(src) == 1 || path.Base(f) == "import.s" {
			if pkg.imports != nil {
				errs.Errorf(file.imports.Kw.Pos,
					"double valid import stmt, maybe two import.s?",
				)
			}
			pkg.imports = file.imports
		} else if file.imports != nil {
			errs.Errorf(file.imports.Kw.Pos,
				"invalid import outside import.s in a multi-file package",
			)
		}
	}

	if len(errs.Errs) != 0 {
		return nil, errs.Errs
	}

	if pkg.imports != nil {
		for _, stmt := range pkg.imports.stmts {
			stmt.linkable = importer.Import(stmt.path)
			if stmt.linkable == nil {
				errs.Errorf(stmt.Path.Pos,
					"import %s is missing by the importer",
					stmt.path,
				)
			} else {
				stmt.lib = stmt.linkable.Lib()
			}
		}

		if len(errs.Errs) != 0 {
			return nil, errs.Errs
		}
	}

	b := newBuilder()
	lib := buildLib(b, pkg)
	es := b.Errs()
	if len(es) != 0 {
		return nil, es
	}

	return lib, nil
}

func (lang) Load(r io.Reader) error {
	panic("todo")
}

// Lang is the assembly language, defined for the building system
var Lang pkg8.Lang = lang{}
