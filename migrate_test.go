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
