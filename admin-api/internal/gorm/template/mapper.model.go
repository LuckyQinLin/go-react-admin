package template

// MapperModel mapper模型
type MapperModel struct {
	TableName string                // 表名
	Select    map[string]MapperItem // 查询
	Update    map[string]MapperItem // 更新
	Delete    map[string]MapperItem // 删除
	Insert    map[string]MapperItem // 插入
}

type ElementType interface {
	SelectElement | UpdateElement | DeleteElement | InsertElement
}

func Transform[T ElementType](list []T) map[string]MapperItem {
	result := make(map[string]MapperItem)
	if len(list) <= 0 {
		return result
	}
	for _, item := range list {
		switch v := any(item).(type) {
		case SelectElement:
			result[v.Id] = MapperItem{
				Id:         v.Id,
				Content:    v.Value,
				ParamName:  v.Param,
				ResultName: v.Result,
			}
		case InsertElement:
			result[v.Id] = MapperItem{
				Id:         v.Id,
				Content:    v.Value,
				ParamName:  v.Param,
				ResultName: v.Result,
			}
		case DeleteElement:
			result[v.Id] = MapperItem{
				Id:         v.Id,
				Content:    v.Value,
				ParamName:  v.Param,
				ResultName: v.Result,
			}
		case UpdateElement:
			result[v.Id] = MapperItem{
				Id:         v.Id,
				Content:    v.Value,
				ParamName:  v.Param,
				ResultName: v.Result,
			}
		}
	}
	return result
}

// NewMapperModel 创建
func NewMapperModel(data *MapperItemModel) *MapperModel {
	return &MapperModel{
		TableName: data.Namespace,
		Select:    Transform[SelectElement](data.Selects),
		Update:    Transform[UpdateElement](data.Updates),
		Delete:    Transform[DeleteElement](data.Deletes),
		Insert:    Transform[InsertElement](data.Inserts),
	}
}

// MapperItem SQL模型
type MapperItem struct {
	Id         string // SQL名称
	Content    string // SQL内容
	ParamName  string // 参数类型名称
	ResultName string // 返回类型名称
}
