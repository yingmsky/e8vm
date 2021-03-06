package ir

import (
	"fmt"

	"lonnie.io/e8vm/link8"
)

func loadRetAddr(b *Block, v *stackVar) {
	if v.size != regSize {
		panic("ret must be regsize")
	}
	// using offset method before SP shift

	loadArg(b, _pc, v)
}

func saveRetAddr(b *Block, v *stackVar) {
	if v.size != regSize {
		panic("ret must be regsize")
	}
	saveArg(b, _ret, v)
}

func saveArg(b *Block, reg uint32, v *stackVar) {
	if v.size == regSize {
		b.inst(asm.sw(reg, _sp, -v.offset))
	} else if v.size == 1 {
		b.inst(asm.sb(reg, _sp, -v.offset))
	} else {
		panic("invalid size to save from a register")
	}
}

func loadArg(b *Block, reg uint32, v *stackVar) {
	if v.size == regSize {
		b.inst(asm.lw(reg, _sp, -v.offset))
	} else if v.size == 1 {
		if !v.u8 {
			b.inst(asm.lb(reg, _sp, -v.offset))
		} else {
			b.inst(asm.lbu(reg, _sp, -v.offset))
		}
	} else {
		panic("invalid size to save from a register")
	}
}

func saveVar(b *Block, reg uint32, v *stackVar) {
	if v.size == regSize {
		b.inst(asm.sw(reg, _sp, *b.frameSize-v.offset))
	} else if v.size == 1 {
		b.inst(asm.sb(reg, _sp, *b.frameSize-v.offset))
	} else {
		panic("invalid size to save from a register")
	}
}

func loadVar(b *Block, reg uint32, v *stackVar) {
	if v.size == regSize {
		b.inst(asm.lw(reg, _sp, *b.frameSize-v.offset))
	} else if v.size == 1 {
		if !v.u8 {
			b.inst(asm.lb(reg, _sp, *b.frameSize-v.offset))
		} else {
			b.inst(asm.lbu(reg, _sp, *b.frameSize-v.offset))
		}
	} else {
		panic("invalid size to load to a register")
	}
}

func saveRef(b *Block, reg uint32, r Ref) {
	switch r := r.(type) {
	case *stackVar:
		saveVar(b, reg, r)
	case *number:
		panic("numbers are read only")
	default:
		panic("not implemented")
	}
}

func loadSym(b *Block, reg uint32, pkg, sym uint32) {
	high := b.inst(asm.lui(reg, 0))
	high.sym = &linkSym{link8.FillHigh, pkg, sym}
	low := b.inst(asm.ori(reg, reg, 0))
	low.sym = &linkSym{link8.FillLow, pkg, sym}
}

func loadRef(b *Block, reg uint32, r Ref) {
	switch r := r.(type) {
	case *stackVar:
		loadVar(b, reg, r)
	case *number:
		high := r.v >> 16
		if high != 0 {
			b.inst(asm.lui(reg, high))
			b.inst(asm.ori(reg, reg, r.v))
		} else {
			b.inst(asm.ori(reg, _0, r.v))
		}
	case *Func:
		loadSym(b, reg, 0, r.index)
	case *funcSym:
		loadSym(b, reg, r.pkg, r.sym)
	default:
		panic(fmt.Errorf("not implemented: %T", r))
	}
}
