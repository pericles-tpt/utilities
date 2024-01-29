package utilities

import (
	"encoding/binary"
	"math"
	"unsafe"
)

var (
	// Source: https://stackoverflow.com/a/59220735
	bigEndian = (*(*[2]uint8)(unsafe.Pointer(&[]uint16{1}[0])))[0] == 0
)

func GetCorrectEndianTimestamp(leastSignificantBytes []byte) uint64 {
	lenRemainingBytes := 8 - len(leastSignificantBytes)
	if lenRemainingBytes < 0 {
		return math.MaxUint64
	}

	remainingBytes := make([]byte, 8-len(leastSignificantBytes))
	if bigEndian {
		return binary.BigEndian.Uint64(append(leastSignificantBytes, remainingBytes...))
	}
	return binary.LittleEndian.Uint64(append(remainingBytes, leastSignificantBytes...))
}

func GetCorrectEndianUint32(bs []byte) uint32 {
	if bigEndian {
		return binary.BigEndian.Uint32(bs)
	}
	return binary.LittleEndian.Uint32(bs)
}

func GetCorrectEndianUint64(bs []byte) uint64 {
	if bigEndian {
		return binary.BigEndian.Uint64(bs)
	}
	return binary.LittleEndian.Uint64(bs)
}

func GetCorrectEndianBytes(data interface{}) []byte {
	if bigEndian {
		return GetBigEndianBytes(data)
	}
	return GetLittleEndianBytes(data)
}

func GetBigEndianBytes(data interface{}) []byte {
	var ret []byte
	switch v := data.(type) {
	case uint32:
		binary.BigEndian.PutUint32(ret, v)
	case uint64:
		binary.BigEndian.PutUint64(ret, v)
	}
	return ret
}

func GetLittleEndianBytes(data interface{}) []byte {
	var ret []byte
	switch v := data.(type) {
	case uint32:
		binary.LittleEndian.PutUint32(ret, v)
	case uint64:
		binary.LittleEndian.PutUint64(ret, v)
	}
	return ret
}
