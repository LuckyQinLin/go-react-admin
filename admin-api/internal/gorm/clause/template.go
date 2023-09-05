package clause

import "github.com/flosch/pongo2/v6"

// Template raw expression
type Template struct {
	SQL                string
	Vars               pongo2.Context
	WithoutParentheses bool
}

func (t Template) Name() string {
	return "TEMPLATE"
}

func (t Template) MergeClause(clause *Clause) {
	clause.Name = ""
	clause.Expression = t
}

// Build raw expression
func (t Template) Build(builder Builder) {
	var (
		tp     *pongo2.Template
		result string
		err    error
	)
	if tp, err = pongo2.FromString(t.SQL); err != nil {

	}
	if result, err = tp.Execute(t.Vars); err != nil {

	}
	builder.WriteString(result)
}
