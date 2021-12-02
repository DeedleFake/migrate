package migrate

import "github.com/DeedleFake/migrate/internal/config"

type M struct {
	deps   []string
	tables []*Table
}

// Require marks another migration as being a dependency of this one.
// In other words, the named migration should be applied before this
// one is.
//
// The provided migration name should be the name of the migration
// function minus the "Migrate" prefix. For example,
//
//    func MigrateFirst(m *migrate.M) {}
//
//    func MigrateSecond(m *migrate.M) {
//      // MigrateSecond depends on MigrateFirst.
//      m.Require("First")
//    }
func (m *M) Require(migration string) {
	m.deps = append(m.deps, config.MigrationPrefix+migration)
}

// Table creates a new table with the given name.
func (m *M) Table(name string) *Table {
	t := &Table{name: name}
	m.tables = append(m.tables, t)
	return t
}
