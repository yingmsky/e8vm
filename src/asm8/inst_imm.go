package asm8

import (
	"strconv"

	"lex8"
)

var (
	// op reg reg imm(signed)
	opImsMap = map[string]uint32{
		"addi": 1,
		"slti": 2,

		"lw":  6,
		"lb":  7,
		"lbu": 8,
		"sw":  9,
		"sb":  10,
	}

	// op reg reg imm(unsigned)
	opImuMap = map[string]uint32{
		"andi": 3,
		"ori":  4,
		"lui":  5,
	}
)

// parseImu parses an unsigned 16-bit immediate
func parseImu(p *Parser, op *lex8.Token) uint32 {
	ret, e := strconv.ParseUint(op.Lit, 0, 32)
	if e != nil {
		p.err(op.Pos, "invalid unsigned immediate %q: %s", op.Lit, e)
		return 0
	}

	if (ret & 0xffff) != ret {
		p.err(op.Pos, "immediate too large: %s", op.Lit)
		return 0
	}

	return uint32(ret)
}

// parseIms parses an unsigned 16-bit immediate
func parseIms(p *Parser, op *lex8.Token) uint32 {
	ret, e := strconv.ParseInt(op.Lit, 0, 32)
	if e != nil {
		p.err(op.Pos, "invalid signed immediate %q: %s", op.Lit, e)
		return 0
	}

	if ret > 0x7fff || ret < -0x8000 {
		p.err(op.Pos, "immediate out of 16-bit range: %s", op.Lit)
		return 0
	}

	return uint32(ret) & 0xffff
}

func makeInstImm(op, d, s, im uint32) *inst {
	ret := uint32(0)
	ret |= (op & 0xff) << 24
	ret |= (d & 0x7) << 21
	ret |= (s & 0x7) << 18
	ret |= (im & 0xffff)

	return &inst{inst: ret}
}

func parseInstImm(p *Parser, ops []*lex8.Token) (*inst, bool) {
	if len(ops) == 0 {
		panic("0 ops")
	}

	op0 := ops[0]
	opName := op0.Lit
	args := ops[1:]

	argCount := func(n int) bool { return argCount(p, ops, n) }

	var op, d, s, im uint32
	if len(args) >= 2 {
		d = parseReg(p, args[0])
		s = parseReg(p, args[1])
	}

	var found bool
	if op, found = opImsMap[opName]; found {
		// op reg reg imm(signed)
		if argCount(3) {
			im = parseIms(p, args[2])
		}
	} else if op, found = opImuMap[opName]; found {
		// op reg reg imm(unsigned)
		if argCount(3) {
			im = parseImu(p, args[2])
		}
	} else {
		return nil, false
	}

	return makeInstImm(op, d, s, im), true
}