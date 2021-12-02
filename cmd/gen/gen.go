package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DeedleFake/migrate/gen"
)

func main() {
	funcs, err := gen.Funcs(context.TODO(), os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load migrations:\n%v", err)
		os.Exit(1)
	}
	for _, f := range funcs {
		fmt.Println(f)
	}
}
