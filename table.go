package migrate

type Table struct {
	name string
	cols []*Column
}

func (t Table) Name() string {
	return t.name
}

// Column adds a column to the table.
func (table *Table) Column(name string, t Type) *Column {
	c := &Column{name: name, t: t}
	table.cols = append(table.cols, c)
	return c
}
