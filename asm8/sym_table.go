package asm8

// SymTable save the symbol
type SymTable struct {
	m map[string]*Symbol
}

// NewSymTable creates an empty symbol table
func NewSymTable() *SymTable {
	ret := new(SymTable)
	ret.m = make(map[string]*Symbol)

	return ret
}

// Query searches for a symbol with a particular name.
func (tab *SymTable) Query(n string) *Symbol {
	s := tab.m[n]
	if s == nil {
		return nil
	}

	return s.clone()
}

// Declare adds a symbol into the table.
// It returns nil on successful, and returns the conflict symbol
// when it is already declared.
func (tab *SymTable) Declare(s *Symbol) *Symbol {
	n := s.Name
	p := tab.m[n]
	if p != nil {
		return p.clone()
	}

	tab.m[n] = s.clone()
	return nil
}