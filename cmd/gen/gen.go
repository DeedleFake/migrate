package main

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"os"

	"github.com/DeedleFake/migrate/gen"
)

func main() {
	funcs, err := gen.Funcs(context.TODO(), os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load migrations:\n%v", err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	err = gen.Generator{
		PkgName: "test",
	}.Runtime(&buf, funcs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate code: %v", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to format output: %v", err)
		os.Exit(1)
	}

	os.Stdout.Write(formatted)
}
