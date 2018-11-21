package packrloader

import (
	"bytes"
	"io"

	"github.com/gobuffalo/packr/v2"
)

// Fileb0xLoader provides a Pongo2 loader for fileb0x template files.
type (
	Loader struct {
		Box *packr.Box
	}
)

// Abs returns the absolute path to a template file.
func (this *Loader) Abs(base, name string) string {
	return name
}

// Get retrieves a reader for the specified path.
func (this *Loader) Get(path string) (io.Reader, error) {
	b, err := this.Box.Find(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
