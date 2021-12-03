package gen

import (
	"context"
	"fmt"
	"go/types"
	"path/filepath"
	"strings"

	"github.com/DeedleFake/migrate/internal/config"
	"golang.org/x/tools/go/packages"
)

const selfPkgPath = "github.com/DeedleFake/migrate"

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

// Funcs returns a list of the names of functions in the given package
// that match the migration function signature. Functions must begin
// with the prefix "Migrate" have the signature
//
//    func(*migrate.M)
//
// The returned slice is sorted according to the standard string
// ordering.
func Funcs(ctx context.Context, pkg string) (funcs []string, err error) {
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

	return funcs, nil
}
