package migrate

// MigrationFunc is the signature matched by functions that define
// migrations.
type MigrationFunc func(m *M)
