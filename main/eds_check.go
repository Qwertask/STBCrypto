package main

import (
	_ "fmt"
	"math/big"
	"strings"
)

func eds_check(S string, G point, Q point, X string, q string, a string, m string) bool {
	if len(S) != 96 {
		return false
	}
	S0 := S[:32]
	S1 := S[32:]
	if StrToBig(rev(S1, len(S1))).Cmp(StrToBig(q)) == 1 {
		return false
	}
	h := belt_hash(X)
	s1_ := StrToBig(rev(S1, len(S1)))
	h_ := StrToBig(rev(h, len(h)))
	g_mul := new(big.Int).Mod(new(big.Int).Add(s1_, h_), StrToBig(q))
	s0_ := StrToBig(rev(new(big.Int).Add(StrToBig(new(big.Int).Exp(StrToBig("2"), StrToBig("80"), nil).Text(16)), StrToBig(rev(S0, len(S0)))).Text(16), len(new(big.Int).Add(StrToBig(new(big.Int).Exp(StrToBig("2"), StrToBig("80"), nil).Text(16)), StrToBig(rev(S0, len(S0)))).Text(16))))

	s0_ = StrToBig(rev(s0_.Text(16), len(s0_.Text(16))))
	R := Add(mult_naf(G, g_mul.Text(16), q, a, m), mult_naf(Q, s0_.Text(16), q, a, m), a, m)

	if R.x == "" && R.y == "" {
		return false
	}
	ch := belt_hash(OIDh + rev(R.x, len(R.x)) + h)[:32]
	if strings.ToLower(S0) != strings.ToLower(ch) {
		return false
	}
	return true
}
