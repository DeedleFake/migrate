package migrate_test

import (
	"testing"

	"deedles.dev/migrate"
	"gotest.tools/v3/assert"
)

func TestPlan(t *testing.T) {
	funcs := map[string]migrate.MigrationFunc{
		"Start":     func(m *migrate.M) {},
		"Left":      func(m *migrate.M) { m.Require("Start") },
		"LeftLeft":  func(m *migrate.M) { m.Require("Left") },
		"LeftRight": func(m *migrate.M) { m.Require("Left") },
		"Right":     func(m *migrate.M) { m.Require("Start") },
		"End":       func(m *migrate.M) { m.Require("LeftLeft", "LeftRight", "Right") },
	}

	plan, err := migrate.Plan(funcs)
	assert.NilError(t, err, "generate plan")
	assert.DeepEqual(t, plan.Steps(), []string{"Start", "Left", "LeftLeft", "LeftRight", "Right", "End"})
}

func TestCyclicPlan(t *testing.T) {
	funcs := map[string]migrate.MigrationFunc{
		"A": func(m *migrate.M) { m.Require("C") },
		"B": func(m *migrate.M) { m.Require("A") },
		"C": func(m *migrate.M) { m.Require("B") },
	}

	_, err := migrate.Plan(funcs)
	assert.ErrorContains(t, err, "dependency cycle detected at")
}

func TestNonExistentDependency(t *testing.T) {
	funcs := map[string]migrate.MigrationFunc{
		"A": func(m *migrate.M) { m.Require("Uh oh...") },
	}

	_, err := migrate.Plan(funcs)
	assert.ErrorContains(t, err, "depends on non-existent migration")
}
