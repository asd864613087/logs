package utils

import (
	"bytes"
	"encoding/binary"
)

// IntToBytes 整形转换成字节
func IntToBytes(n int) ([]byte, error) {
	x := int32(n) // 这里决定了len的位数

	bytesBuffer := bytes.NewBuffer([]byte{})

	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil {
		return []byte{}, err
	}

	return bytesBuffer.Bytes(), nil
}

// BytesToInt 字节转换成整形
func BytesToInt(b []byte) (int, error) {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	err := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		return -1, err
	}

	return int(x), nil
}
