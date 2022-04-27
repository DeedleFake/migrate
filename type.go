package migrate

type columnType interface {
	SQL(Dialect) string
}

type stringType struct{}

func (stringType) SQL(dialect Dialect) string {
	return "STRING"
}

type intType struct{}

func (intType) SQL(dialect Dialect) string {
	return "INT"
}
