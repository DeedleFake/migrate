package migrate

type Type interface {
	SQL(Dialect) string
}

type stringType struct{}

func (stringType) SQL(dialect Dialect) string {
	return "STRING"
}

func String() Type {
	return stringType{}
}

type intType struct{}

func (intType) SQL(dialect Dialect) string {
	return "INT"
}

func Int() Type {
	return intType{}
}
