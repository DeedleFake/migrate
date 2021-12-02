package migrate

type Column struct {
	name string
	t    Type
}

func (c Column) Name() string {
	return c.name
}

func (c Column) Type() Type {
	return c.t
}
