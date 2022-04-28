package migrate

import (
	"context"
	"database/sql"
	"fmt"

	"deedles.dev/migrate/internal/util"
	"golang.org/x/exp/slices"
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

func (m *Migration) Steps() []string {
	steps := make([]string, 0, len(m.steps))
	for _, s := range m.steps {
		steps = append(steps, s.name)
	}
	return steps
}

func flattenDAG(verts map[string]*M) (steps []*M, err error) {
	defer func() {
		switch r := recover().(type) {
		case nil:
		case error:
			err = r
		default:
			panic(r)
		}
	}()

	rem := util.SortedKeys(verts) // To ensure determinsitic behavior.
	steps = make([]*M, 0, len(verts))

	var perm util.Set[string]
	var tmp util.Set[string]
	var inner func(*M)
	inner = func(m *M) {
		if perm.Contains(m.name) {
			return
		}
		defer perm.Add(m.name)

		if tmp.Contains(m.name) {
			panic(fmt.Errorf("dependency cycle detected at %q", m.name))
		}
		tmp.Add(m.name)
		defer tmp.Remove(m.name)

		deps := m.deps.Slice()
		slices.Sort(deps) // To ensure deterministic behavior.
		for _, dep := range deps {
			d, ok := verts[dep]
			if !ok {
				panic(fmt.Errorf("migration %q depends on non-existent migration %q", m.name, dep))
			}
			inner(d)
		}

		steps = append(steps, m)
	}

	for perm.Len() < len(rem) {
		for _, name := range rem {
			inner(verts[name])
		}
	}
	return steps, nil
}

type sqler interface {
	sql(Dialect) string
}
