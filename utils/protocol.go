package utils

import (
	"fmt"
	"strings"
)

const (
	protocolHeader = "StreamDataHeader"
	defaultDataSizeLen = 4
)

var (
	ptcHeader = ProtocolHeader{
		Header:      protocolHeader,
		HeaderLen:   len(protocolHeader),
		DataSizeLen: defaultDataSizeLen,
	}
)

type ProtocolHeader struct {
	Header      string
	HeaderLen   int
	DataSizeLen int // 存储数据长度的位数
}

func PackStreamData(b []byte) []byte {
	// header
	baseBuf := []byte(ptcHeader.Header)

	// dataSizeLen
	dlBytes, err := IntToBytes(len(b))
	if err != nil {
		fmt.Printf("[PackStreamData] IntToBytes HeaderLen Failed, err=%s", err)
	}
	baseBuf = append(baseBuf, dlBytes...)

	// Msg - StreamData
	baseBuf = append(baseBuf, b...)

	return baseBuf
}

func UnPackStreamData(b []byte, c chan []byte) ([]byte, error) {

	hl := ptcHeader.HeaderLen
	dsl := ptcHeader.DataSizeLen

	// Index返回子串第一个字母出现的下标，而非下标前一位
	for {
		lens := len(b)
		s := string(b)

		if i := strings.Index(s, ptcHeader.Header); i != - 1 {
			DataSize, err := BytesToInt(b[i + hl : i + hl + dsl])
			if err != nil {
				fmt.Printf("[UnPackStreamData] Failed, err:=%s", err)
			}

			// b[totalHL]指向数据段第一位
			totalHL := i + hl + dsl

			// 边界：i + hl + dsl + DataSize - 1 == lens - 1
			if  totalHL + DataSize > lens {
				return b[i:], nil
			}

			// 截断头部写进channel
			c <- b[totalHL : totalHL + DataSize]

			// buf重置
			if totalHL + DataSize == lens {
				return []byte{}, nil
			}

			b = b[totalHL + DataSize:]
		}
	}


}
