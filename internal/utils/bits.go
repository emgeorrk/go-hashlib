package utils

import (
	"fmt"
	"strings"
)

// BoolXor returns a xor of multiple bool values
func BoolXor(a ...bool) bool {
	result := false
	for _, v := range a {
		result = result != v
	}
	return result
}

// BitsToHexString converts an array of bits to a hex string
func BitsToHexString(bits []bool) string {
	var hexStr strings.Builder

	// Process bits in 4-bit chunks
	for i := 0; i < len(bits); i += 4 {
		val := 0
		for j := 0; j < 4 && i+j < len(bits); j++ {
			if bits[i+j] {
				val |= 1 << (3 - j)
			}
		}
		hexStr.WriteString(fmt.Sprintf("%x", val))
	}

	return hexStr.String()
}

// BitsToBytes converts an array of bits to an array of bytes
func BitsToBytes(bits []bool) []byte {
	byteLen := (len(bits) + 7) / 8
	bytes := make([]byte, byteLen)

	for i, bit := range bits {
		if bit {
			byteIndex := i / 8
			bitIndex := 7 - (i % 8)
			bytes[byteIndex] |= 1 << bitIndex
		}
	}

	return bytes
}

// BytesToBits converts a byte slice to a bit array
func BytesToBits(data []byte) []bool {
	bits := make([]bool, len(data)*8)
	for i, b := range data {
		for j := 0; j < 8; j++ {
			bits[i*8+j] = ((b >> (7 - j)) & 1) == 1
		}
	}
	return bits
}
