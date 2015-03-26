package ir

type Ref interface{}

// stackVar is a variable on stack
type stackVar struct {
	name   string
	id     int
	offset int32
	size   int32

	// reg is the register allocated
	// valid values are in range [1, 4] for normal values
	// and also ret register is 6
	viaReg uint32

	// regOnly stack vars does not take frame space on the stack
	regOnly bool
}

type heapVar struct{ pkg, sym int } // a variable symbol on heap
type funcSym struct{ pkg, sym int } // a function symbol
type number struct{ v uint32 }      // a constant number

func Num(v uint32) Ref { return &number{v} }
func Snum(v int32) Ref { return &number{uint32(v)} }
