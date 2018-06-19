package strictjson

import (
	"encoding/json"
	"reflect"
)

type structUnmarshaler struct {
	structVal reflect.Value
}

func NewStructUnmarshaler(v interface{}) (json.Unmarshaler, error) {
	ptrVal := reflect.ValueOf(v)
	if ptrVal.Kind() != reflect.Ptr {
		return nil, ErrNotAStructPointer{v}
	}
	structVal := ptrVal.Elem()
	if structVal.Kind() != reflect.Struct {
		return nil, ErrNotAStructPointer{v}
	}
	return &structUnmarshaler{structVal}, nil
}

func (su *structUnmarshaler) UnmarshalJSON(data []byte) error {
	var fieldsJSON map[string]json.RawMessage
	err := json.Unmarshal(data, &fieldsJSON)
	if err != nil {
		return err
	}
	structType := su.structVal.Type()
	newStructVal := reflect.New(structType).Elem()
	for i := 0; i < structType.NumField(); i++ {
		newFieldVal := newStructVal.Field(i)
		if !newFieldVal.CanSet() {
			continue
		}
		fieldName := structType.Field(i).Name
		fieldJSON, ok := fieldsJSON[fieldName]
		if !ok {
			return ErrFieldNotFound{fieldName, data}
		}
		err := json.Unmarshal(fieldJSON, newFieldVal.Addr().Interface())
		if err != nil {
			return err
		}
		delete(fieldsJSON, fieldName)
	}
	if len(fieldsJSON) != 0 {
		fieldNames := make([]string, 0, len(fieldsJSON))
		for fieldName, _ := range fieldsJSON {
			fieldNames = append(fieldNames, fieldName)
		}
		return ErrUnknownFields{fieldNames}
	}
	for i := 0; i < structType.NumField(); i++ {
		fieldVal := su.structVal.Field(i)
		if !fieldVal.CanSet() {
			continue
		}
		fieldVal.Set(newStructVal.Field(i))
	}
	return nil
}
