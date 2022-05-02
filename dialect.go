package migration

import "strings"

// Dialect represents an SQL dialect. Dialects are comparable, and the
// instances of this struct returned by the various functions in this
// package are guaranteed to always be equal to the result of another
// call to the same function.
type Dialect struct {
	name    string
	quoteID rune
}

func (d Dialect) Name() string {
	return d.name
}

var (
	postgres = Dialect{
		name:    "postgres",
		quoteID: '"',
	}

	sqlite3 = Dialect{
		name:    "sqlite3",
		quoteID: '`',
	}
)

func Postgres() Dialect { return postgres }
func SQLite3() Dialect  { return sqlite3 }

func (d Dialect) id(name string) string {
	var str strings.Builder
	str.WriteRune(d.quoteID)
	str.WriteString(name)
	str.WriteRune(d.quoteID)
	return str.String()
}
