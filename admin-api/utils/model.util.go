package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// StructCopy 结构体复制
// source 当前有值的结构体
// target 接受值的结构体
// fields 需要的设置的属性
func StructCopy(source interface{}, target interface{}, fields ...string) (err error) {
	sourceKey := reflect.TypeOf(source)
	sourceVal := reflect.ValueOf(source)

	targetKey := reflect.TypeOf(target)
	targetVal := reflect.ValueOf(target)

	if targetKey.Kind() != reflect.Ptr {
		err = fmt.Errorf("被覆盖的数据必须是一个结构体指针")
		return
	}

	targetVal = reflect.ValueOf(targetVal.Interface())

	// 存放字段
	fieldItems := make([]string, 0)

	if len(fields) > 0 {
		fieldItems = fields
	} else {
		for i := 0; i < sourceVal.NumField(); i++ {
			fieldItems = append(fieldItems, sourceKey.Field(i).Name)
		}
	}

	for i := 0; i < len(fieldItems); i++ {
		field := targetVal.Elem().FieldByName(fieldItems[i])
		value := sourceVal.FieldByName(fieldItems[i])
		if field.IsValid() && field.Kind() == value.Kind() {
			field.Set(value)
		}

	}
	return

}

func StrToArray(str string) []int64 {
	var (
		strArray []string
		numArray = make([]int64, 0)
	)
	strArray = strings.Split(str, ",")
	for _, item := range strArray {
		newItem, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			continue
		}
		numArray = append(numArray, newItem)
	}
	return numArray
}
