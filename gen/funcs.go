package gen

import (
	"context"
	"fmt"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

const selfPkgPath = "github.com/DeedleFake/migrate"

type errSlice []packages.Error

func (err errSlice) Error() string {
	var buf strings.Builder

	var sep string
	for _, err := range err {
		fmt.Fprintf(&buf, "%v%v", sep, err.Msg)
		sep = "\n"
	}

	return buf.String()
}

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

	m := self.Types.Scope().Lookup("M").Type()
	sig := types.NewSignature(
		nil,
		types.NewTuple(
			types.NewParam(token.NoPos, nil, "m", types.NewPointer(m)),
		),
		nil,
		false,
	)

	scope := target.Types.Scope()
	for _, n := range scope.Names() {
		if !strings.HasPrefix(n, "Migrate") {
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
