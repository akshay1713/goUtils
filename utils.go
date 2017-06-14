package goUtils

import "log"

func PadLeft(str, pad string, length int) string {
	if len(str) >= length {
		return str
	}
	for {
		str = pad + str
		if len(str) > length {
			return str[0:length]
		}
	}
}

func HandleErr(err error, prefix string) {
	if err != nil {
		log.Println(prefix, ": ", err)
	}
}

func PanicErr(err error){
	if err != nil {
		panic(err)
	}
}

func GetBytesFromUint16(source []byte, num uint16) {
	source[0] = byte(num >> 8)
	source[1] = byte(num)
}

func GetBytesFromUint32(source []byte, num uint32) {
	source[0] = byte(num >> 24)
	source[1] = byte(num >> 16)
	source[2] = byte(num >> 8)
	source[3] = byte(num)
}

func GetBytesFromUint64(source []byte, num uint64) {
	source[0] = byte(num >> 56)
	source[1] = byte(num >> 48)
	source[2] = byte(num >> 40)
	source[3] = byte(num >> 32)
	source[4] = byte(num >> 24)
	source[5] = byte(num >> 16)
	source[6] = byte(num >> 8)
	source[7] = byte(num)
}

func Pos(slice []string, value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
