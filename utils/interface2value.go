package utils

import (
	"fmt"
	"strconv"
)

func CtxValueToBytes(value interface{}) ([]byte, error) {
	switch value.(type) {
	case string:
		if s, ok := value.(string); ok {
			return []byte(s), nil
		}
		return []byte{}, fmt.Errorf("[CtxValueToBytes] Convert string Failed: value=%+v", value)
	case []byte:
		if b, ok := value.([]byte); ok {
			return b, nil
		}
		return []byte{}, fmt.Errorf("[CtxValueToBytes] Convert []byte Failed: value=%+v", value)
	case []rune:
		if r, ok := value.([]rune); ok {
			return []byte(string(r)), nil
		}
		return []byte{}, fmt.Errorf("[CtxValueToBytes] Convert []rune Failed: value=%+v", value)
	case int:
		if vInt, ok := value.(int); ok {
			s := strconv.Itoa(vInt)
			return []byte(s), nil
		}
		return []byte{}, fmt.Errorf("[CtxValueToBytes] Convert Int Failed: value=%+v", value)
	case int32:
		if vInt32, ok := value.(int32); ok {
			s := strconv.Itoa(int(vInt32))
			return []byte(s), nil
		}
		return []byte{}, fmt.Errorf("[CtxValueToBytes] Convert int32 Failed: value=%+v", value)
	case int64:
		if vInt64, ok := value.(int64); ok {
			s := strconv.FormatInt(vInt64, 10)
			return []byte(s), nil
		}
		return []byte{}, fmt.Errorf("[CtxValueToBytes] Convert int64 Failed: value=%+v", value)
	case uint:
		if vUint, ok := value.(uint); ok {
			s := strconv.FormatUint(uint64(vUint), 10)
			return []byte(s), nil
		}
		return []byte{}, fmt.Errorf("[CtxValueToBytes] Convert uint Failed: value=%+v", value)
	case uint32:
		if vUint32, ok := value.(uint32); ok {
			s := strconv.FormatUint(uint64(vUint32), 10)
			return []byte(s), nil
		}
		return []byte{}, fmt.Errorf("[CtxValueToBytes] Convert uint32 Failed: value=%+v", value)
	case uint64:
		if vUint64, ok := value.(uint64); ok {
			s := strconv.FormatUint(vUint64, 10)
			return []byte(s), nil
		}
		return []byte{}, fmt.Errorf("[CtxValueToBytes] Convert uint64 Failed: value=%+v", value)

	default:
		// TODO: 编解码问题
		//buffer := bytes.NewBuffer([]byte{})
		//enc := gob.NewEncoder(buffer)
		//if err := enc.Encode(value); err == nil {
		//	return buffer.Bytes(), nil
		//}
		//return []byte{},  fmt.Errorf("[CtxValueToBytes] Convert Struct Failed: value=%+v", value)
		return []byte(fmt.Sprintf("%+v", value)), nil
	}
}
