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

type MUp struct {
	steps []mstep
}

func (m MUp) migrate(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	for _, step := range m.steps {
		err := step.migrateUp(ctx, tx, dialect)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MUp) SQL(stmt string, args ...any) {
	m.steps = append(m.steps, sqlstep{stmt: stmt, args: args})
}

type MDown struct {
	steps []mstep
}

func (m MDown) migrate(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	for _, step := range m.steps {
		// Should this be migrateDown instead? If so, how will custom SQL
		// work?
		err := step.migrateUp(ctx, tx, dialect)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MDown) SQL(stmt string, args ...any) {
	m.steps = append(m.steps, sqlstep{stmt: stmt, args: args})
}
