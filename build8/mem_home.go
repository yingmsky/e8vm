package build8

import (
	"io"
	"sort"
	"strings"
)

// MemHome is a memory based building home.
type MemHome struct {
	pkgs  map[string]*MemPkg
	langs *langPicker
}

// NewMemHome creates a new memory-based home
func NewMemHome(lang Lang) *MemHome {
	ret := new(MemHome)
	ret.pkgs = make(map[string]*MemPkg)
	ret.langs = newLangPicker(lang)
	return ret
}

// NewPkg creates (or replaces) a package of a particular path in this home.
func (h *MemHome) NewPkg(path string) *MemPkg {
	ret := newMemPkg(path)
	h.pkgs[path] = ret
	return ret
}

// Pkgs lists all the packages in this home
func (h *MemHome) Pkgs(prefix string) []string {
	var ret []string
	for p := range h.pkgs {
		if strings.HasPrefix(p, prefix) {
			ret = append(ret, p)
		}
	}
	sort.Strings(ret)
	return ret
}

// Src lists all the source files in this home
func (h *MemHome) Src(p string) map[string]*File {
	pkg := h.pkgs[p]
	if pkg == nil {
		return nil
	}

	if len(pkg.files) == 0 {
		return nil
	}

	ret := make(map[string]*File)
	for name, f := range pkg.files {
		ret[name] = &File{
			Path:       "$" + p + "/" + name,
			Name:       name,
			ReadCloser: f.Reader(),
		}
	}

	return ret
}

// CreateLib opens the library file for writing
func (h *MemHome) CreateLib(p string) io.WriteCloser {
	pkg := h.pkgs[p]
	if pkg == nil {
		panic("pkg not exists")
	}
	if pkg.lib == nil {
		pkg.lib = newMemFile()
	} else {
		pkg.lib.Reset()
	}
	return pkg.lib
}

// CreateBin opens the library binary for writing
func (h *MemHome) CreateBin(p string) io.WriteCloser {
	pkg := h.pkgs[p]
	if pkg == nil {
		panic("pkg not exists")
	}
	if pkg.bin == nil {
		pkg.bin = newMemFile()
	} else {
		pkg.bin.Reset()
	}
	return pkg.bin
}

// Bin returns the binary for math maitn
func (h *MemHome) Bin(p string) []byte {
	pkg := h.pkgs[p]
	if pkg == nil {
		panic("pkg not exists")
	}
	if pkg.bin == nil {
		return nil
	}
	return pkg.bin.Bytes()
}

// CreateLog creates a log file for writing
func (h *MemHome) CreateLog(p, name string) io.WriteCloser {
	pkg := h.pkgs[p]
	if pkg == nil {
		panic("pkg not exists")
	}
	ret := newMemFile()
	pkg.logs[name] = ret
	return ret
}

// Lang returns the language for path
func (h *MemHome) Lang(path string) Lang { return h.langs.lang(path) }

// AddLang adds a language to a prefix
func (h *MemHome) AddLang(prefix string, lang Lang) {
	h.langs.addLang(prefix, lang)
}

var _ Home = new(MemHome) // satisfying the interface
