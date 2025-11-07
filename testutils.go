package ethutils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	// #nosec G404
	random = rand.New(rand.NewSource(time.Now().Unix()))
}

func GenRandHexString(n int) string {
	bytes := make([]byte, n)
	random.Read(bytes)
	return hex.EncodeToString(bytes)
}

// GenRandEVMAddr generates a random Ethereum address using non-cryptographic randomness
func GenRandEVMAddr() string {
	return "0x" + GenRandHexString(20)
}

// GenRandEVMHash generates a random Ethereum hash using non-cryptographic randomness
func GenRandEVMHash() string {
	return "0x" + GenRandHexString(32)
}
