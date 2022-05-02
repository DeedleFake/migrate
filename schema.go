package migration

import (
	"context"
	"database/sql"
	"fmt"

	"deedles.dev/migration/internal/util"
)

const schemaTable = "_migration_schema"

type schema struct {
	prev *util.Set[string]
}

func initSchema(ctx context.Context, db *sql.DB) error {
	// TODO: Dialect support.
	_, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS `+schemaTable+` (
		id BIGINT GENERATED ALWAYS AS IDENTITY,
		name TEXT NOT NULL
	);`)
	if err != nil {
		return fmt.Errorf("create table: %w", err)
	}

	return nil
}

func loadSchema(ctx context.Context, db *sql.DB) (*schema, error) {
	if db == nil {
		// Dirty hack to make testing work easier.
		// TODO: Find a better way to do this.
		return &schema{prev: &util.Set[string]{}}, nil
	}

	rows, err := db.QueryContext(ctx, `SELECT name FROM `+schemaTable+`;`)
	if err != nil {
		return nil, fmt.Errorf("query schema table: %w", err)
	}
	defer rows.Close()

	var prev util.Set[string]
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		prev.Add(name)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}

	return &schema{
		prev: &prev,
	}, nil
}

func (s *schema) addPrev(ctx context.Context, tx *sql.Tx, name string) error {
	// TODO: Dialect support.
	_, err := tx.ExecContext(ctx, `INSERT INTO `+schemaTable+` (name) VALUES ($1);`, name)
	if err != nil {
		return fmt.Errorf("insert into schema table: %w", err)
	}

	s.prev.Add(name)
	return nil
}
