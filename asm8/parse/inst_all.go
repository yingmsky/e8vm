package parse

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

var insts = []instParse{
	parseInstReg,
	parseInstImm,
	parseInstBr,
	parseInstJmp,
	parseInstSys,
}

func parseInst(log lex8.Logger, ops []*lex8.Token) (i *ast.Inst) {
	return instParsers(insts).parse(log, ops)
}
