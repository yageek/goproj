package internal

import (
	"errors"
	"fmt"
)

type EpsgIndex struct {
	nameMap map[string]interface{}
	epsgMap map[int]interface{}
}

func NewIndex() *EpsgIndex {
	index := &EpsgIndex{}
	index.nameMap = make(map[string]interface{})
	index.epsgMap = make(map[int]interface{})

	return index
}

func (index *EpsgIndex) ByName(name string) (interface{}, error) {

	val, ok := index.nameMap[name]
	if ok {
		return val, nil
	} else {
		return nil, errors.New(fmt.Sprintf("No existing element in index with name '%s'\n", name))
	}

}

func (index *EpsgIndex) ByEpsgCode(code int) (interface{}, error) {

	val, ok := index.epsgMap[code]
	if ok {
		return val, nil
	} else {
		return nil, errors.New(fmt.Sprintf("No existing element in index with code '%s'\n", code))
	}

}

func (index *EpsgIndex) AddEntry(name string, code int, i interface{}) error {

	_, err := index.ByName(name)
	if err != nil {
		return errors.New(fmt.Sprintf("An element prexists within the index with the name '%s'\n", name))
	}

	_, err = index.ByEpsgCode(code)
	if err != nil {
		return errors.New(fmt.Sprintf("An element prexist within the index with the code '%i'\n", code))
	}

	index.nameMap[name] = i
	index.epsgMap[code] = i

	return nil

}
