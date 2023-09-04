package template

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ParseMapper 解析Mapper
type ParseMapper struct {
	TableName string         // 表名
	ParamMap  map[string]any // 参数类型映射
	ResultMap map[string]any // 返回类型映射
}

func New(paramMap map[string]any, resultMap map[string]any) *ParseMapper {
	return &ParseMapper{
		ParamMap:  paramMap,
		ResultMap: resultMap,
	}
}

func (p ParseMapper) Parse(content string) (*MapperModel, error) {
	var (
		mapper  MapperItemModel
		decoder *xml.Decoder
		reader  io.Reader
		err     error
	)
	reader = strings.NewReader(content)
	decoder = xml.NewDecoder(reader)
	if err = decoder.Decode(&mapper); err != nil {
		return nil, errors.New("解析XML文件失败: " + err.Error())
	}
	fmt.Printf("读取结果：%v", mapper)
	return NewMapperModel(p.ParamMap, p.ResultMap, &mapper), nil
}
