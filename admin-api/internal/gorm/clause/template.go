package clause

import "github.com/flosch/pongo2/v6"

// Template raw expression
type Template struct {
	SQL                string
	Vars               pongo2.Context
	WithoutParentheses bool
}

// Build build raw expression
func (expr Template) Build(builder Builder) {
	var (
		tp     *pongo2.Template
		result string
		err    error
	)
	if tp, err = pongo2.FromString(expr.SQL); err != nil {

	}
	if result, err = tp.Execute(expr.Vars); err != nil {

	}
	builder.WriteString(result)
}
