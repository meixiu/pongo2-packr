package pongo2packr

import (
	"bytes"
	"io"
	"path/filepath"

	"github.com/gobuffalo/packr/v2"
)

// Fileb0xLoader provides a Pongo2 loader for fileb0x template files.
type (
	Loader struct {
		Box *packr.Box
	}
)

func NewLoader(box *packr.Box) *Loader {
	return &Loader{box}
}

// Abs returns the absolute path to a template file.
func (this *Loader) Abs(base, name string) string {
	if filepath.IsAbs(name) {
		return name
	}
	p := filepath.Join(filepath.Dir(base), name)
	return p
}

// Get retrieves a reader for the specified path.
func (this *Loader) Get(path string) (io.Reader, error) {
	b, err := this.Box.Find(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
