package gen

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/DeedleFake/migrate/internal/tmpl"
)

type Generator struct {
	PkgName string
}

func (g Generator) Runtime(w io.Writer, funcs []string) error {
	migrations := make([]string, len(funcs))
	for i := range funcs {
		migrations[i] = strings.TrimPrefix(funcs[i], "Migrate")
	}

	return tmpl.Templates.ExecuteTemplate(w, "runtime.tmpl", tmpl.RuntimeData{
		SelfPkgPath: selfPkgPath,
		SelfPkgName: filepath.Base(selfPkgPath),
		PkgName:     g.PkgName,
		Migrations:  migrations,
	})
}
