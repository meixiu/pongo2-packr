# pongo2-packr
pongo2 packr loader

```
import (
	"github.com/gobuffalo/packr/v2"
	"github.com/wbsifan/pongo2-packr"
)

// then
box := packr.New("template", "../template")
loader := pongo2packr.NewLoader(box)
t := pongo2.NewSet("", loader)
tpl, err = t.FromFile("index.html")
```