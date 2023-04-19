package handler

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

type FilterType uint8

const (
	FilterTypeEqual FilterType = iota
	FilterTypeContains
	FilterTypeGreaterThan
	FilterTypeGreaterThanOrEqual
	FilterTypeLessThan
	FilterTypeLessThanOrEqual
	FilterTypeBetween
)

type PageQuery struct {
	PageNum  int `form:"pageNum" ` // 页码
	PageSize int `form:"pageSize"` // 每页数量
}

type PaginationData struct {
	Total    int         `json:"total"`
	List     interface{} `json:"list"`
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
}

type FilterQuery struct {
	FilterPropertyName string      `json:"propertyName"`
	FilterType         FilterType  `json:"filterType"`
	FilterValue        interface{} `json:"value"`
}

type SortQuery struct {
	PropertyName string `json:"propertyName"`
	Ascending    bool   `json:"ascending"`
}

/*
	{
		"resource": "pods",
		"page": {
			"pageNum": 1,
			"pageSize": 10
		},
		"filters": [
			{
				"propertyName": "name",
				"filterType": "equal",
				"value": "test"
			}

		],
		"sort": {
			"propertyName": "name",
			"ascending": true
		}

}
*/
type DataFilter struct {
	NameFilter  string        `json:"nameFilter"`
	PageQuery   PageQuery     `json:"page"`
	FilterQuery []FilterQuery `json:"filters"`
	SortQuery   SortQuery     `json:"sort"`
}

// 绑定数据
func (d *DataFilter) Bind(c *gin.Context) error {
	err := c.ShouldBindJSON(d)
	if err != nil {
		// 如果是EOF错误, 则说明没有传入数据, 则使用默认值
		if strings.Contains(err.Error(), "EOF") {
			return nil
		}
		return err
	}

	return nil
}

// 对数据进行过滤
func (d *DataFilter) Filter(data []interface{}) PaginationData {

	// 根据名称模糊查询
	if d.NameFilter != "" {
		var result []interface{}
		for _, item := range data {
			if strings.Contains(objectToMap(item)["name"].(string), d.NameFilter) {
				result = append(result, item)
			}
		}
		data = result
	}

	// 根据数据获取过滤器
	filters := d.getFilters()

	// 过滤数据
	for _, filter := range filters {
		data = filter.filter(data)
	}

	// 根据数据获取排序器
	sorter := d.getSorter(data)

	// 排序数据
	if sorter != nil {
		sorter.sort(data)
	}

	// 分页数据
	return d.page(data)
}

// 分页数据, 默认第一页, 每页10条
func (d *DataFilter) page(data []interface{}) PaginationData {

	if d.PageQuery.PageNum == 0 {
		d.PageQuery.PageNum = 1
	}

	if d.PageQuery.PageSize == 0 {
		d.PageQuery.PageSize = 10
	}

	// 获取分页数据
	pageNum := d.PageQuery.PageNum
	pageSize := d.PageQuery.PageSize

	// 计算总数
	total := len(data)

	// 计算分页数据
	start := (pageNum - 1) * pageSize
	end := pageNum * pageSize

	if start > total {
		start = total
	}

	if end > total {
		end = total
	}

	// 获取分页数据
	list := data[start:end]

	return PaginationData{
		Total:    total,
		List:     list,
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

// 根据数据获取排序器
func (d *DataFilter) getSorter(data []interface{}) sorter {
	// 如果没有排序字段, 则不排序
	if d.SortQuery.PropertyName == "" {
		return nil
	}

	if len(data) == 0 {
		return nil
	}

	// 获取第一个元素

	// 获取属性值
	value := objectToMap(data[0])[d.SortQuery.PropertyName]

	// 获取属性类型
	valueType := reflect.ValueOf(value).Type()

	// 根据类型获取排序器
	switch valueType.Kind() {
	case reflect.String:
		return &stringSorter{
			propertyName: d.SortQuery.PropertyName,
			asc:          d.SortQuery.Ascending,
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &numSorter{
			propertyName: d.SortQuery.PropertyName,
			asc:          d.SortQuery.Ascending,
		}
	case reflect.Float32, reflect.Float64:
		return &numSorter{
			propertyName: d.SortQuery.PropertyName,
			asc:          d.SortQuery.Ascending,
		}
	}

	return nil
}

// 根据数据获取过滤器
func (d *DataFilter) getFilters() []filter {
	var filters []filter
	for _, item := range d.FilterQuery {
		switch item.FilterType {
		case FilterTypeEqual:
			filters = append(filters, &equalFilter{
				propertyName: item.FilterPropertyName,
				value:        item.FilterValue,
			})
		case FilterTypeContains:
			filters = append(filters, &containsFilter{
				propertyName: item.FilterPropertyName,
				value:        item.FilterValue,
			})
		case FilterTypeGreaterThan:
			filters = append(filters, &greaterThanFilter{
				propertyName: item.FilterPropertyName,
				value:        item.FilterValue,
			})
		case FilterTypeGreaterThanOrEqual:
			filters = append(filters, &greaterThanOrEqualFilter{
				propertyName: item.FilterPropertyName,
				value:        item.FilterValue,
			})
		case FilterTypeLessThan:
			filters = append(filters, &lessThanFilter{
				propertyName: item.FilterPropertyName,
				value:        item.FilterValue,
			})
		case FilterTypeLessThanOrEqual:
			filters = append(filters, &lessThanOrEqualFilter{
				propertyName: item.FilterPropertyName,
				value:        item.FilterValue,
			})
		case FilterTypeBetween:
			filters = append(filters, &betweenFilter{
				propertyName: item.FilterPropertyName,
				value:        item.FilterValue,
			})
		}
	}
	return filters
}

// 定义过滤器
type filter interface {
	filter(data []interface{}) []interface{}
}

// 定义等于过滤器
type equalFilter struct {
	propertyName string
	value        interface{}
}

func (e *equalFilter) filter(data []interface{}) []interface{} {
	// 判断属性是否存在
	if _, ok := objectToMap(data[0])[e.propertyName]; !ok {
		return data
	}

	var result []interface{}
	for _, item := range data {
		if objectToMap(item)[e.propertyName] == e.value {
			result = append(result, item)
		}
	}
	return result
}

// 定义包含过滤器
type containsFilter struct {
	propertyName string
	value        interface{}
}

func (e *containsFilter) filter(data []interface{}) []interface{} {
	// 判断属性是否存在
	if _, ok := objectToMap(data[0])[e.propertyName]; !ok {
		return data
	}

	var result []interface{}
	for _, item := range data {
		if strings.Contains(objectToMap(item)[e.propertyName].(string), e.value.(string)) {
			result = append(result, item)
		}
	}
	return result
}

// 定义大于过滤器
type greaterThanFilter struct {
	propertyName string
	value        interface{}
}

func (e *greaterThanFilter) filter(data []interface{}) []interface{} {
	// 判断属性是否存在
	if _, ok := objectToMap(data[0])[e.propertyName]; !ok {
		return data
	}

	var result []interface{}
	for _, item := range data {
		if objectToMap(item)[e.propertyName].(float64) > e.value.(float64) {
			result = append(result, item)
		}
	}
	return result
}

// 定义大于等于过滤器
type greaterThanOrEqualFilter struct {
	propertyName string
	value        interface{}
}

func (e *greaterThanOrEqualFilter) filter(data []interface{}) []interface{} {
	// 判断属性是否存在
	if _, ok := objectToMap(data[0])[e.propertyName]; !ok {
		return data
	}

	var result []interface{}
	for _, item := range data {
		if objectToMap(item)[e.propertyName].(float64) >= e.value.(float64) {
			result = append(result, item)
		}
	}
	return result
}

// 定义小于过滤器
type lessThanFilter struct {
	propertyName string
	value        interface{}
}

func (e *lessThanFilter) filter(data []interface{}) []interface{} {
	// 判断属性是否存在
	if _, ok := objectToMap(data[0])[e.propertyName]; !ok {
		return data
	}

	var result []interface{}
	for _, item := range data {
		if objectToMap(item)[e.propertyName].(float64) < e.value.(float64) {
			result = append(result, item)
		}
	}
	return result
}

// 定义小于等于过滤器
type lessThanOrEqualFilter struct {
	propertyName string
	value        interface{}
}

func (e *lessThanOrEqualFilter) filter(data []interface{}) []interface{} {

	// 判断属性是否存在
	if _, ok := objectToMap(data[0])[e.propertyName]; !ok {
		return data
	}

	var result []interface{}
	for _, item := range data {
		if objectToMap(item)[e.propertyName].(float64) <= e.value.(float64) {
			result = append(result, item)
		}
	}
	return result
}

// 定义between过滤器
type betweenFilter struct {
	propertyName string
	value        interface{}
}

func (e *betweenFilter) filter(data []interface{}) []interface{} {

	// 判断属性是否存在
	if _, ok := objectToMap(data[0])[e.propertyName]; !ok {
		return data
	}

	// 判断value是否是数组
	if reflect.TypeOf(e.value).Kind() != reflect.Slice {
		return data
	}

	var result []interface{}
	for _, item := range data {
		if objectToMap(item)[e.propertyName].(float64) >= e.value.([]interface{})[0].(float64) &&
			objectToMap(item)[e.propertyName].(float64) <= e.value.([]interface{})[1].(float64) {
			result = append(result, item)
		}
	}
	return result
}

// 定义排序器
type sorter interface {
	sort(data []interface{}) []interface{}
}

// 定义整形排序器
type numSorter struct {
	propertyName string
	asc          bool
}

func (sorter *numSorter) sort(data []interface{}) []interface{} {

	// 判断是否为升序
	if sorter.asc {
		sort.Slice(data, func(i, j int) bool {
			return objectToMap(data[i])[sorter.propertyName].(float64) < objectToMap(data[j])[sorter.propertyName].(float64)
		})
	} else {
		sort.Slice(data, func(i, j int) bool {
			return objectToMap(data[i])[sorter.propertyName].(float64) > objectToMap(data[j])[sorter.propertyName].(float64)
		})
	}
	return data
}

// 定义字符串排序器
type stringSorter struct {
	propertyName string
	asc          bool
}

func (sorter *stringSorter) sort(data []interface{}) []interface{} {

	// 判断属性是否存在
	if _, ok := objectToMap(data[0])[sorter.propertyName]; !ok {
		return data
	}

	// 判断属性是否为字符串
	if reflect.TypeOf(objectToMap(data[0])[sorter.propertyName]).Kind() != reflect.String {
		return data
	}

	// 判断是否为升序
	if sorter.asc {
		sort.Slice(data, func(i, j int) bool {
			return objectToMap(data[i])[sorter.propertyName].(string) < objectToMap(data[j])[sorter.propertyName].(string)
		})
	} else {
		sort.Slice(data, func(i, j int) bool {
			return objectToMap(data[i])[sorter.propertyName].(string) > objectToMap(data[j])[sorter.propertyName].(string)
		})
	}
	return data
}

// 将对象转化为map[string]interface{}
func objectToMap(obj interface{}) map[string]interface{} {
	var result map[string]interface{}
	jsonData, _ := json.Marshal(obj)
	json.Unmarshal(jsonData, &result)
	return result
}
