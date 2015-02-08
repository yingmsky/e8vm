package asm8

import (
	"bytes"
	"strconv"

	"lonnie.io/e8vm/lex8"
)

func parseDataHex(p *parser, args []*lex8.Token) ([]byte, uint32) {
	if !checkAllType(p, args, Operand) {
		return nil, 0
	}
	buf := new(bytes.Buffer)
	for _, arg := range args {
		b, e := strconv.ParseUint(arg.Lit, 16, 8)
		if e != nil {
			p.err(arg.Pos, "%s", e)
			return nil, 0
		}

		buf.WriteByte(byte(b))
	}

	return buf.Bytes(), 0
}