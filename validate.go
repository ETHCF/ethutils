package ethutils

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	EthAddrRegex = `^0x[0-9|a-f|A-F]{40}$`
	EthHashRegex = `^0x[0-9|a-f|A-F]{64}$`
)

var (
	EthAddrExp        = regexp.MustCompile(EthAddrRegex)
	EthHashExp        = regexp.MustCompile(EthHashRegex)
	ErrInvalidEthAddr = fmt.Errorf("invalid ethereum address")
	ErrInvalidEthHash = fmt.Errorf("invalid ethereum hash")
)

// SanitizeEthAddr returns a valid Ethereum address or an error if the address is invalid.
// will add the 0x prefix if it's missing, and lowercase the input.
func SanitizeEthAddr(addr string) (string, error) {
	out := strings.ToLower(addr)
	if len(out) == 40 {
		out = "0x" + out // Add the 0x prefix if it's missing
	}
	if len(out) != 42 { // Ethereum addresses are 20 bytes
		return "", ErrInvalidEthAddr
	}
	if !EthAddrExp.MatchString(out) {
		return "", ErrInvalidEthAddr
	}
	return out, nil
}

// SanitizeEthHash returns a valid Ethereum hash or an error if the hash is invalid.
// will add the 0x prefix if it's missing, and lowercase the input.
func SanitizeEthHash(hash string) (string, error) {
	out := strings.ToLower(hash)
	if len(out) == 64 {
		out = "0x" + out // Add the 0x prefix if it's missing
	}
	if len(out) != 66 { // Ethereum hashes are 32 bytes
		return "", ErrInvalidEthHash
	}
	if !EthHashExp.MatchString(out) {
		return "", ErrInvalidEthHash
	}
	return out, nil
}
