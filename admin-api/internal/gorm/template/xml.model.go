package template

import "encoding/xml"

// MapperItemModel 对应一个实体的查询模型
type MapperItemModel struct {
	Name      xml.Name        `xml:"mapper"`
	Namespace string          `xml:"namespace,attr"`
	Selects   []SelectElement `xml:"select"`
	Updates   []UpdateElement `xml:"update"`
	Inserts   []InsertElement `xml:"insert"`
	Deletes   []DeleteElement `xml:"delete"`
}

// SelectElement 查询
type SelectElement struct {
	Name   xml.Name `xml:"select"`
	Id     string   `xml:"id,attr"`
	Param  string   `xml:"param,attr"`
	Result string   `xml:"result,attr"`
	Value  string   `xml:",innerxml"`
}

// UpdateElement 更新
type UpdateElement struct {
	Name   xml.Name `xml:"update"`
	Id     string   `xml:"id,attr"`
	Param  string   `xml:"param,attr"`
	Result string   `xml:"result,attr"`
	Value  string   `xml:",innerxml"`
}

// DeleteElement 删除
type DeleteElement struct {
	Name   xml.Name `xml:"delete"`
	Id     string   `xml:"id,attr"`
	Param  string   `xml:"param,attr"`
	Result string   `xml:"result,attr"`
	Value  string   `xml:",innerxml"`
}

// InsertElement 插入
type InsertElement struct {
	Name   xml.Name `xml:"insert"`
	Id     string   `xml:"id,attr"`
	Param  string   `xml:"param,attr"`
	Result string   `xml:"result,attr"`
	Value  string   `xml:",innerxml"`
}
