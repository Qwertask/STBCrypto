package main

import (
	"crypto/rand"
	"math/big"
)

func key_gen(q string, G point, a string, m string) (string, point) {
	min := new(big.Int).SetInt64(int64(1))
	maxi := new(big.Int).Sub(StrToBig(q), new(big.Int).SetInt64(int64(1)))
	d, _ := RandomBigIntInRange(min, maxi)
	Q := mult_binary(G, d.Text(16), q, a, m)
	return d.Text(16), Q
}

func RandomBigIntInRange(min, max *big.Int) (*big.Int, error) {
	rangeBigInt := new(big.Int).Sub(max, min)
	n, err := rand.Int(rand.Reader, rangeBigInt)
	if err != nil {
		return nil, err
	}
	n.Add(n, min)
	return n, nil
}
