package util

import (
	"bytes"
	"encoding/binary"
	"math/rand"
)

//自行设置随机种子
//一般服务器启动时，设置
//rand.Seed(time.Now().UnixNano())
func RandBetween(min int, max int) int {

	if min == max {
		return min
	} else if min > max {
		return max + rand.Intn(min-max+1)
	} else {
		return min + rand.Intn(max-min+1)
	}
}

func Int2Byte(data int) (ret []byte) {

	value := int32(data)
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, value)
	if err != nil {
		return nil
	}

	return buffer.Bytes()

}

func Byte2Int(data []byte) int {

	var buffer = bytes.NewBuffer(data)
	value := int32(0)
	binary.Read(buffer, binary.BigEndian, &value)

	return int(value)
}
