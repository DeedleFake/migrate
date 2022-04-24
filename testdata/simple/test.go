package migrations

import "deedles.dev/migrate"

func Migrate1(m *migrate.M)                       {}
func Migrate2(m2 *migrate.M)                      { m2.Require("3") }
func Migrate3(somethingElseCompletely *migrate.M) {}
