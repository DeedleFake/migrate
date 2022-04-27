package migrate

type T struct {
	name    string
	cols    []sqler
	indices []*Index
}

func (t T) Name() string {
	return t.name
}

func addColumn[C columnType, V any](t *T, name string) *Column[V] {
	var ct C
	c := Column[V]{name: name, t: ct}
	t.cols = append(t.cols, &c)
	return &c
}

func (t *T) String(name string) *Column[string] {
	return addColumn[stringType, string](t, name)
}

func (t *T) Int(name string) *Column[int] {
	return addColumn[intType, int](t, name)
}

func (t *T) Index(names ...string) *Index {
	i := Index{names: names}
	t.indices = append(t.indices, &i)
	return &i
}

type Column[T any] struct {
	name string
	t    columnType
	d    *T
	null bool
}

func (c Column[T]) sql(d Dialect) string {
	panic("Not implemented.")
}

func (c *Column[T]) Default(d T) *Column[T] {
	c.d = &d
	return c
}

func (c *Column[T]) DefaultFunc() *Column[T] {
	panic("Not implemented.")
}

func (c *Column[T]) Null(allow bool) *Column[T] {
	c.null = allow
	return c
}

type Index struct {
	names  []string
	unique bool
}

func (i *Index) Unique(unique bool) *Index {
	i.unique = unique
	return i
}
