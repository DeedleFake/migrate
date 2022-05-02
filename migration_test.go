package migration_test

import (
	"context"
	"testing"

	"deedles.dev/migration"
	"gotest.tools/v3/assert"
)

func TestPlan(t *testing.T) {
	funcs := map[string]migration.MigrationFunc{
		"Start":     func(m *migration.M) {},
		"Left":      func(m *migration.M) { m.Require("Start") },
		"LeftLeft":  func(m *migration.M) { m.Require("Left") },
		"LeftRight": func(m *migration.M) { m.Require("Left") },
		"Right":     func(m *migration.M) { m.Require("Start") },
		"End":       func(m *migration.M) { m.Require("LeftLeft", "LeftRight", "Right") },
	}

	plan, err := migration.Plan(context.TODO(), nil, funcs)
	assert.NilError(t, err, "generate plan")
	assert.DeepEqual(t, plan.Steps(), []string{"Start", "Left", "LeftLeft", "LeftRight", "Right", "End"})
}

func TestCyclicPlan(t *testing.T) {
	funcs := map[string]migration.MigrationFunc{
		"A": func(m *migration.M) { m.Require("C") },
		"B": func(m *migration.M) { m.Require("A") },
		"C": func(m *migration.M) { m.Require("B") },
	}

	_, err := migration.Plan(context.TODO(), nil, funcs)
	assert.ErrorContains(t, err, "dependency cycle detected at")
}

func TestNonExistentDependency(t *testing.T) {
	funcs := map[string]migration.MigrationFunc{
		"A": func(m *migration.M) { m.Require("Uh oh...") },
	}

	_, err := migration.Plan(context.TODO(), nil, funcs)
	assert.ErrorContains(t, err, "depends on non-existent migration")
}
