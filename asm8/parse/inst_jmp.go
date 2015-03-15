package parse

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

var opJmpMap = map[string]uint32{
	"j":   2,
	"jal": 3,
}

func isValidSymbol(sym string) bool {
	return true
}

func parseInstJmp(p *parser, ops []*lex8.Token) (*ast.Inst, bool) {
	op0 := ops[0]
	opName := op0.Lit
	var op uint32

	// op sym
	switch opName {
	case "j":
		op = 2
	case "jal":
		op = 3
	default:
		return nil, false
	}

	var pack, sym string
	var fill int
	var symTok *lex8.Token

	if argCount(p, ops, 1) {
		symTok = ops[1]
		if parseLabel(p, ops[1]) {
			sym = ops[1].Lit
			fill = ast.FillLabel
		} else {
			pack, sym = parseSym(p, ops[1])
			fill = ast.FillLink
		}
	}

	ret := &ast.Inst{
		Inst:   (op & 0x3) << 30,
		Pkg:    pack,
		Sym:    sym,
		Fill:   fill,
		SymTok: symTok,
	}
	return ret, true
}