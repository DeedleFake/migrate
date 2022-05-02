package migration

import (
	"context"
	"database/sql"
)

type updown struct {
	up   MUp
	down MDown
}

func (ud updown) migrateUp(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	return ud.up.migrate(ctx, tx, dialect)
}

func (ud updown) migrateDown(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	return ud.down.migrate(ctx, tx, dialect)
}

type MUp struct{}

func (m MUp) migrate(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	panic("Not implemented.")
}

type MDown struct{}

func (m MDown) migrate(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	panic("Not implemented.")
}
