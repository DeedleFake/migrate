package migrations

import "deedles.dev/migration"

func Migrate1(m *migration.M)                       {}
func Migrate2(m2 *migration.M)                      { m2.Require("3") }
func Migrate3(somethingElseCompletely *migration.M) {}
