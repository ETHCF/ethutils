package ethutils

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func HexToByte32(hexString string) ([32]byte, error) {
	b, err := hex.DecodeString(strings.TrimLeft(hexString, "0x"))
	if err != nil {
		panic(err)
	}
	var out [32]byte
	if len(b) > 32 {
		return out, fmt.Errorf("input hex string is too long: %d", len(b))
	}
	if len(b) < 32 {
		b = append(b, make([]byte, 32-len(b))...)
	}
	copy(out[:], b)
	return out, nil
}
