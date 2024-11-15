package main

import (
	"math"
	"math/big"
)

func Skobka(a int) int {
	a_sk := 0
	a_bin := toBinary32(a)
	a_mas := make([]int, 0)
	for i := 0; i < len(a_bin)/8; i++ {
		a_mas = append(a_mas, binaryToInt(a_bin[8*i:8+8*i]))
	}
	for i := 0; i < len(a_mas); i++ {
		a_sk += a_mas[i] * int(math.Pow(256, float64(3-i)))
	}
	return a_sk % (1 << 32)
}
func belt_hash(X string) string {
	r := Skobka(splitAndReverse(len(X) * 4))
	r_big := new(big.Int).SetInt64(int64(r))
	for len(r_big.Text(16)) != 32 {
		r_big, _ = r_big.SetString(r_big.Text(16)+"0", 16)
	}
	s := new(big.Int).SetInt64(int64(0))
	h := "B194BAC80A08F53B366D008E584A5DE48504FA9D1BB6C7AC252E72C202FDCE0D"
	Xn := make([]string, 0)
	for len(X) > 0 {
		for len(X)%64 != 0 {
			X = X + "0"
		}
		Xn = append(Xn, (X[:64]))
		X = X[64:]
	}
	var t string
	for _, xi := range Xn {
		t, h = belt_compress(xi + h)
		temp, _ := new(big.Int).SetString(t, 16)
		s = new(big.Int).Xor(s, temp)
	}
	_, Y := belt_compress(r_big.Text(16) + s.Text(16) + h)
	return Y
}
func belt_compress(x string) (string, string) {
	for len(x) < 128 {
		x = "0" + x
	}

	X1, X2, X3, X4 := (x[0:32]), (x[32:64]), (x[64:96]), (x[96:128])

	x1, _ := new(big.Int).SetString(X1, 16)
	x2, _ := new(big.Int).SetString(X2, 16)
	x3, _ := new(big.Int).SetString(X3, 16)
	x4, _ := new(big.Int).SetString(X4, 16)
	S1, _ := new(big.Int).SetString(Belt_block_encrypt((new(big.Int).Xor(x3, x4)).Text(16), X1+X2), 16)
	S := new(big.Int).Xor(x3, new(big.Int).Xor(S1, x4))
	ed, _ := new(big.Int).SetString("11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111", 2)
	xor_n := new(big.Int).Xor(S, ed).Text(16)
	T, _ := new(big.Int).SetString(Belt_block_encrypt(X2, xor_n+X3), 16)
	Y2 := new(big.Int).Xor(x2, T)
	S1, _ = new(big.Int).SetString(Belt_block_encrypt(X1, S.Text(16)+X4), 16)
	Y1 := new(big.Int).Xor(S1, x1)
	Y1TXT := Y1.Text(16)
	Y2TXT := Y2.Text(16)
	for len(Y1TXT)%32 != 0 {
		Y1TXT = "0" + Y1TXT
	}
	for len(Y2TXT)%32 != 0 {
		Y2TXT = "0" + Y2TXT
	}
	return S.Text(16), (Y1TXT + Y2TXT)
}
