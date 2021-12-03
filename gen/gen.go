package gen

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/DeedleFake/migrate/internal/tmpl"
)

func Runtime(w io.Writer, pkg *Package) error {
	migrations := make([]string, len(pkg.Funcs))
	for i := range pkg.Funcs {
		migrations[i] = strings.TrimPrefix(pkg.Funcs[i], "Migrate")
	}

	return tmpl.Templates.ExecuteTemplate(w, "runtime.tmpl", tmpl.RuntimeData{
		SelfPkgPath: selfPkgPath,
		SelfPkgName: filepath.Base(selfPkgPath),
		PkgName:     pkg.Name,
		Migrations:  migrations,
	})
}
