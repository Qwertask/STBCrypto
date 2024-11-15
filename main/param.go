package main

import (
	"fmt"
	"math/big"
)

type point struct {
	x string
	y string
}

func base_point(b, p string) point {
	b_big, _ := new(big.Int).SetString(b, 16)
	p_big, _ := new(big.Int).SetString(p, 16)
	addResult := new(big.Int).Add(p_big, big.NewInt(1))
	divResult := new(big.Int).Div(addResult, big.NewInt(4))
	return point{
		x: fmt.Sprintf("%x", 0),
		y: new(big.Int).Exp(b_big, divResult, p_big).Text(16),
	}
}

func Legendre(u, n string) int {
	a := new(big.Int)
	p := new(big.Int)

	// Преобразование строк в big.Int
	a.SetString(u, 16)
	p.SetString(n, 16)

	// exp = (p - 1) / 2
	exp := new(big.Int).Sub(p, big.NewInt(1))
	exp.Div(exp, big.NewInt(2))

	// result = a^exp % p
	result := new(big.Int).Exp(a, exp, p)

	// Сравнение результата
	if result.Cmp(big.NewInt(1)) == 0 {
		return 1
	} else if result.Cmp(new(big.Int).Sub(p, big.NewInt(1))) == 0 {
		return -1
	} else {
		return 0
	}
}

type EC struct {
	p    string
	a    string
	b    string
	seed string
	q    string
	yG   point
}

//l = 128

var STBparam EC = EC{p: REV("43FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
	a:    REV("40FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
	b:    REV("F1039CD66B7D2EB253928B976950F54CBEFBD8E4AB3AC1D2EDA8F315156CCE77"),
	seed: REV("5E3801000000000016"),
	q:    REV("07663D2699BF5A7EFC4DFB0DD68E5CD9FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
	yG: point{
		x: "0",
		y: REV("936A510418CF291E52F608C4663991785D83D651A3C9E45C9FD616FB3CFCF76B"),
	},
}

func REV(s string) string {
	result := make([]byte, len(s))

	for i, j := len(s)-1, 0; i >= 1; i, j = i-2, j+2 {
		result[j] = s[i-1]
		result[j+1] = s[i]
	}

	return string(result)
}
