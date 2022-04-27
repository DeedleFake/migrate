package migrate

type T struct {
	name string
	cols []any
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

type Column[T any] struct {
	name string
	t    columnType
	d    *T
	null bool
}

func (c *Column[T]) Default(d T) *Column[T] {
	c.d = &d
	return c
}

func (c *Column[T]) Null(allow bool) *Column[T] {
	c.null = allow
	return c
}
