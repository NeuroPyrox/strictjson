package strictjson

func UnmarshalStruct(data []byte, v interface{}) error {
	su, err := NewStructUnmarshaler(v)
	if err != nil {
		return err
	}
	return su.UnmarshalJSON(data)
}
