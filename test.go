//go:build ignore
// +build ignore

package main

import (
	"context"

	migrations "github.com/DeedleFake/migrate/testdata/simple"
)

func main() {
	err := migrations.Apply(context.TODO(), nil, migrations.Names()...)
	if err != nil {
		panic(err)
	}
}
