package ir

type linkSym struct {
	fill int
	pkg  uint32 // package index, 0 for the same package
	sym  uint32 // symbol index
}

type inst struct {
	inst uint32
	sym  *linkSym
}
