package strictjson

import (
	"fmt"
)

type ErrNotAStructPointer struct {
	v interface{}
}

func (e ErrNotAStructPointer) Error() string {
	return fmt.Sprintf("Not a struct pointer: %v", e.v)
}

type ErrFieldNotFound struct {
	fieldName string
	data      []byte
}

func (e ErrFieldNotFound) Error() string {
	return fmt.Sprintf("Field %#v not found: %s", e.fieldName, e.data)
}

type ErrUnknownFields struct {
	fieldNames []string
}

func (e ErrUnknownFields) Error() string {
	return fmt.Sprintf("Unknown field(s): %v", e.fieldNames)
}
