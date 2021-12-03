package tmpl

import (
	"embed"
	"text/template"
)

//go:embed *.tmpl
var fs embed.FS

var Templates = template.Must(template.ParseFS(fs, "*.tmpl"))

type RuntimeData struct {
	SelfPkgPath string
	SelfPkgName string
	PkgName     string
	Migrations  []string
}
