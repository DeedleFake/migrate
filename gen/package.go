package gen

import (
	"context"
	"fmt"
	"go/types"
	"path/filepath"
	"strings"

	"deedles.dev/migrate/internal/config"
	"golang.org/x/tools/go/packages"
)

const selfPkgPath = "deedles.dev/migrate"

type errSlice []packages.Error

func (err errSlice) Error() string {
	var buf strings.Builder

	var sep string
	for _, err := range err {
		fmt.Fprintf(&buf, "%v%v", sep, err)
		sep = "\n"
	}

	return buf.String()
}

type Package struct {
	Name  string
	Funcs []string
}

// Load loads migration information from the given package.
func Load(ctx context.Context, pkg string) (*Package, error) {
	path, err := filepath.Abs(pkg)
	if err != nil {
		return nil, fmt.Errorf("get path to package: %w", err)
	}

	pkgs, err := packages.Load(&packages.Config{
		Mode:    packages.NeedName | packages.NeedImports | packages.NeedDeps | packages.NeedTypes,
		Context: ctx,
	}, path, selfPkgPath)
	if err != nil {
		return nil, fmt.Errorf("load package: %w", err)
	}

	var target, self *packages.Package
	for _, pkg := range pkgs {
		if errs := pkg.Errors; len(errs) != 0 {
			return nil, errSlice(errs)
		}

		if pkg.PkgPath == selfPkgPath {
			self = pkg
			continue
		}
		target = pkg
	}

	sig := self.Types.Scope().Lookup("MigrationFunc").Type().Underlying()
	var funcs []string

	scope := target.Types.Scope()
	for _, n := range scope.Names() {
		if !strings.HasPrefix(n, config.MigrationPrefix) {
			continue
		}

		obj := scope.Lookup(n)
		if !types.Identical(obj.Type(), sig) {
			return nil, fmt.Errorf("declaration %q matches naming scheme but is wrong type", obj.Name())
		}

		funcs = append(funcs, obj.Name())
	}

	return &Package{
		Name:  target.Name,
		Funcs: funcs,
	}, nil
}
