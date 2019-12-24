package builder

import (
	"errors"
	"reflect"
	"strings"
)

type IndexValid interface {
	IdxValid(map[string]reflect.Type) bool
}

func IdxValid(model interface{}, cond Cond) (bool, error) {
	indexFields, err := parseIndexFields(model)
	if err != nil {
		return false, err
	}
	if len(indexFields) == 0 {
		return false, errors.New("no index found")
	}
	return cond.IdxValid(indexFields), nil
}

func parseIndexFields(model interface{}) (map[string]reflect.Type, error) {
	if model == nil {
		return nil, errors.New("param nil")
	}
	pType := reflect.TypeOf(model)
	if pType.Kind() != reflect.Ptr {
		return nil, errors.New("param not ptr")
	}
	v := reflect.ValueOf(model).Elem()
	t := v.Type()
	ret := map[string]reflect.Type{}
	for index := 0; index < v.NumField(); index++ {
		tField := t.Field(index)
		tag := tField.Tag.Get("validx")
		if isIdx, col := parseIdxTag(tag); isIdx {
			if col == "" {
				continue
			}
			ret[col] = tField.Type
		}
	}
	return ret, nil
}

// tag == "index,col:id"
func parseIdxTag(tag string) (isIdx bool, col string) {
	vals := strings.Split(tag, ",")
	for _, v := range vals {
		if strings.TrimSpace(v) == "index" {
			isIdx = true
			continue
		}
		kvs := strings.Split(v, ":")
		if len(kvs) >= 2 {
			if "col" == strings.TrimSpace(kvs[0]) {
				col = strings.TrimSpace(kvs[1])
			}
		}
	}
	return
}

func colIdxCheck(cols map[string]reflect.Type, colName string, colT reflect.Type, colVal interface{}) bool {
	vType, ok := cols[colName]
	if !ok {
		return false
	}
	return (colT == vType) && (colT == vType)
}
