package clause

import "github.com/flosch/pongo2/v6"

// PageCount raw expression
type PageCount struct {
	CountSQL           string         // 统计SQL
	CountVars          pongo2.Context // 统计SQL参数
	WithoutParentheses bool
}

func (p PageCount) Name() string {
	return "PAGE COUNT"
}

func (p PageCount) MergeClause(clause *Clause) {
	clause.Name = ""
	clause.Expression = p
}

// Build raw expression
func (p PageCount) Build(builder Builder) {
	var (
		tp     *pongo2.Template
		result string
		err    error
	)
	if tp, err = pongo2.FromString(p.CountSQL); err != nil {

	}
	if result, err = tp.Execute(p.CountVars); err != nil {

	}
	builder.WriteString(result)
}

// PageQuery raw expression
type PageQuery struct {
	SQL                string         // 分页SQL
	Vars               pongo2.Context // 分页SQL参数
	WithoutParentheses bool
}

func (p PageQuery) Name() string {
	return "PAGE QUERY"
}

func (p PageQuery) MergeClause(clause *Clause) {
	clause.Name = ""
	clause.Expression = p
}

// Build raw expression
func (p PageQuery) Build(builder Builder) {
	var (
		tp     *pongo2.Template
		result string
		err    error
	)
	if tp, err = pongo2.FromString(p.SQL); err != nil {

	}
	if result, err = tp.Execute(p.Vars); err != nil {

	}
	builder.WriteString(result)
}
