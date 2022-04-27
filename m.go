package migrate

type M struct {
	name   string
	deps   []string
	tables []*T
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
	m.deps = append(m.deps, migration)
}

func (m *M) CreateTable(name string, f func(*T)) {
	t := T{name: name}
	m.tables = append(m.tables, &t)
	f(&t)
}
