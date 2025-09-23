package ethutils

import (
	"math/big"
)

// SolidityIntSize is the size of a Solidity int in bits.
const SolidityIntSize uint = 256

func ToDecimal(balance *big.Int, decimals int) *big.Float {
	pow := big.NewInt(10)
	pow = pow.Exp(pow, big.NewInt(int64(decimals)), nil)

	normalizedBalance := big.NewFloat(0).SetPrec(SolidityIntSize).SetInt(balance)
	return normalizedBalance.Quo(normalizedBalance, big.NewFloat(0).SetPrec(SolidityIntSize).SetInt(pow))
}

func FromDecimal(balance *big.Float, decimals int) *big.Int {
	pow := big.NewInt(10)
	pow = pow.Exp(pow, big.NewInt(int64(decimals)), nil)

	balance = balance.Mul(balance, big.NewFloat(0).SetPrec(SolidityIntSize).SetInt(pow))

	res, _ := balance.Int(nil)
	return res
}
