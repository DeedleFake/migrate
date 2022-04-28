package migrate

import "deedles.dev/migrate/internal/util"

type M struct {
	name   string
	deps   util.Set[string]
	tables []*T
}

// Require marks other migrations as being dependencies of this one.
// In other words, the named migrations should be applied before this
// one is.
//
// Calling this function more than once is equivalent to calling it
// once with all of the same arguments.
//
// The provided migration names should be the name of the migration
// function minus the "Migrate" prefix. For example,
//
//    func MigrateFirst(m *migrate.M) {}
//
//    func MigrateSecond(m *migrate.M) {
//      // MigrateSecond depends on MigrateFirst.
//      m.Require("First")
//    }
func (m *M) Require(migrations ...string) {
	for _, mig := range migrations {
		m.deps.Add(mig)
	}
}

func (m *M) CreateTable(name string, f func(*T)) {
	t := T{name: name}
	m.tables = append(m.tables, &t)
	f(&t)
}
