package main

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	"github.com/DeedleFake/migrate/gen"
)

func main() {
	runtime := filepath.Join(os.Args[1], "runtime.go")
	err := os.Remove(runtime)
	if (err != nil) && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Failed to remove existing runtime.go: %v", err)
		os.Exit(1)
	}

	pkg, err := gen.Load(context.TODO(), os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load migrations:\n%v", err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	err = gen.Runtime(&buf, pkg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate code: %v", err)
		os.Exit(1)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to format output: %v", err)
		os.Exit(1)
	}

	file, err := os.Create(runtime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create runtime.go: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(formatted)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write runtime.go: %v", err)
		os.Exit(1)
	}
}
