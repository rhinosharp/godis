package godis

import (
	"errors"
	"fmt"
	"sort"
)

//DataMap 注意它不是并发安全的
type DataMap map[string]interface{}

//New 初始化slice实例
func New() DataMap {
	return make(DataMap, 0)
}

//List Key元素列表
func (dm DataMap) List() []string {
	var keys []string
	for k := range dm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

//Get Get Value by key
func (dm DataMap) Get(key string) interface{} {
	return dm[key]
}

//Add 添加元素
func (dm DataMap) Add(key string, elem interface{}) error {
	for k := range dm {
		if k == key {
			errMsg := fmt.Sprintf("[Error]Cannot add the same key: %v\n", key)
			return errors.New(errMsg)
		}
	}
	dm[key] = elem
	return nil
}

//Update 更新元素
func (dm DataMap) Update(key string, elem interface{}) error {
	for k := range dm {
		if k == key {
			dm[key] = elem
			return nil
		}
	}
	errMsg := fmt.Sprintf("[Error]No such key: %v\n", key)
	return errors.New(errMsg)
}

//Remove 删除元素
func (dm DataMap) Remove(key string) error {
	delete(dm, key)
	return nil
}
