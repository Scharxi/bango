package util

import (
	"crypto/rand"
	"math/big"
)

func GenerateAccountNumber() (int64, error) {
	min := big.NewInt(1000000000) // kleinste 10-stellige Zahl
	max := big.NewInt(9999999999) // größte 10-stellige Zahl
	randomNumber, err := rand.Int(rand.Reader, max.Sub(max, min).Add(max, big.NewInt(1)))
	if err != nil {
		return 0, err
	}
	return randomNumber.Int64(), nil
}
