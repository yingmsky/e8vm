package arch8

// InstSys exectues system instruction
type InstSys struct{}

// I executes the system instruction.
// Returns any exception encountered.
func (i *InstSys) I(cpu *CPU, in uint32) *Excep {
	op := (in >> 24) & 0xff // (32:24]
	src := (in >> 21) & 0x7 // (24:21]
	s := cpu.regs[src]

	switch op {
	case 64: // halt
		return errHalt
	case 65: // syscall
		return errSyscall
	case 66: // usermod
		cpu.ring = 1
	case 67: // vtable
		cpu.virtMem.SetTable(s)
	default:
		return errInvalidInst
	}

	return nil
}
