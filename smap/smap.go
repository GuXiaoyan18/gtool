package smap

import "container/list"

type Smap struct {
	dataMap  map[string]interface{}
	dataList *list.List
}

func NewSmap() *Smap {
	return &Smap{
		dataMap: make(map[string]interface{}),
		dataList: list.New(),
	}
}

func (s *Smap) Add(key string, value interface{}) bool {
	_, exist := s.dataMap[key]
	if exist {
		return false
	} else {
		s.dataMap[key] = value
		s.dataList.PushBack(key)
		return true
	}
}

func (s *Smap) Remove(key string) {
	_, exist := s.dataMap[key]
	if exist {
		delete(s.dataMap, key) // 删除map的元素
		for e := s.dataList.Front(); e != nil; e = e.Next() {
			if e.Value == key {
				s.dataList.Remove(e)
				break
			}
		}
	}
}

func (s *Smap) Get(key string) interface{} {
	return s.dataMap[key]
}
