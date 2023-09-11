package template

import (
	"encoding/xml"
	"io"
)

type TType int

const (
	Query TType = iota
	Update
	Delete
	Insert
)

// ParseMapper 解析Mapper
type ParseMapper struct {
	mapper *MapperItemModel
}

func New(file io.Reader) *ParseMapper {
	var mapper MapperItemModel
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&mapper); err != nil {
		return &ParseMapper{mapper: nil}
	}
	return &ParseMapper{mapper: &mapper}
}

func (p ParseMapper) Parse() (*MapperModel, error) {
	return NewMapperModel(p.mapper), nil
}
