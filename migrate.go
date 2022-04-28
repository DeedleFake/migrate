package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"deedles.dev/migrate/internal/util"
)

// MigrationFunc is the signature matched by functions that define
// migrations.
type MigrationFunc func(m *M)

type Migration struct {
	steps []*M
}

// Plan produces a migration plan for a given set of migration
// functions. It is intended for internal use.
func Plan(funcs map[string]MigrationFunc) (*Migration, error) {
	verts := make(map[string]*M, len(funcs))
	for n, f := range funcs {
		m := M{name: n}
		f(&m)
		verts[n] = &m
	}

	steps := make([]*M, 0, len(verts))
	for _, m := range verts {
		m.fillDeps(verts)
		steps = util.SortedInsertFunc(steps, m, func(v1, v2 *M) int {
			if v2.deps.Contains(v1.name) {
				return -1
			}
			if v1.deps.Contains(v2.name) {
				return 1
			}
			return strings.Compare(v1.name, v2.name)
		})
	}

	return &Migration{
		steps: steps,
	}, nil
}

func (m *Migration) Run(ctx context.Context, db *sql.DB) error {
	for _, s := range m.steps {
		fmt.Printf("%+v\n", s)
	}

	panic("Not implemented.")
}

func (m *Migration) Steps() []string {
	steps := make([]string, 0, len(m.steps))
	for _, s := range m.steps {
		steps = append(steps, s.name)
	}
	return steps
}

type sqler interface {
	sql(Dialect) string
}
