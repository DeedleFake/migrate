package migrate

import (
	"context"
	"database/sql"
	"fmt"
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

	steps, err := flattenDAG(verts)
	if err != nil {
		return nil, fmt.Errorf("calculate migration order: %w", err)
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

func flattenDAG(verts map[string]*M) (steps []*M, err error) {
	steps = make([]*M, 0, len(verts))

	visited := make(map[*M]struct{}, len(verts))
	var inner func(*M)
	inner = func(m *M) {
		if err != nil {
			return
		}

		if _, ok := visited[m]; ok {
			return
		}
		visited[m] = struct{}{}

		for _, dep := range m.deps {
			d, ok := verts[dep]
			if !ok {
				err = fmt.Errorf("migration %v depends on non-existent migration %q", m.name, dep)
				return
			}
			if _, ok := visited[d]; ok {
				err = fmt.Errorf("dependency cycle detected: %v -> %v", m.name, dep)
				return
			}
			inner(d)
		}

		steps = append(steps, m)
	}

	for _, m := range verts {
		inner(m)
	}

	return steps, err
}
