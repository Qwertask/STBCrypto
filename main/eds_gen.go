package main

import (
	"math/big"
)

func eds_gen(p string, a string, q string, G point, X string, d string) string {
	h := belt_hash(X)
	k, _ := RandomBigIntInRange(StrToBig("0"), StrToBig(q))
	k = StrToBig(rev("4C0E74B2CD5811AD21F23DE7E0FA742C3ED6EC483C461CE15C33A77AA308B7D2", len("4C0E74B2CD5811AD21F23DE7E0FA742C3ED6EC483C461CE15C33A77AA308B7D2")))
	R := mult_binary(G, k.Text(16), q, a, p)
	S0 := StrToBig(belt_hash(OIDh + rev(R.x, len(R.x)) + h)[:32])
	S0_out := S0.Text(16)
	S0 = StrToBig(rev(new(big.Int).Add(StrToBig(new(big.Int).Exp(StrToBig("2"), StrToBig("80"), nil).Text(16)), StrToBig(rev(S0_out, len(S0_out)))).Text(16), len(new(big.Int).Add(StrToBig(new(big.Int).Exp(StrToBig("2"), StrToBig("80"), nil).Text(16)), StrToBig(rev(S0_out, len(S0_out)))).Text(16))))
	k = StrToBig(rev(k.Text(16), len(k.Text(16))))
	S1 := rev(new(big.Int).Mod(new(big.Int).Sub(StrToBig(new(big.Int).Mod(new(big.Int).Sub(StrToBig(rev(k.Text(16), len(k.Text(16)))), StrToBig(rev(h, len(h)))), StrToBig(q)).Text(16)), StrToBig(new(big.Int).Mul(StrToBig(rev(S0.Text(16), len(S0.Text(16)))), StrToBig(rev(d, len(d)))).Text(16))), StrToBig(q)).Text(16), len(new(big.Int).Mod(new(big.Int).Sub(StrToBig(new(big.Int).Mod(new(big.Int).Sub(StrToBig(rev(k.Text(16), len(k.Text(16)))), StrToBig(rev(h, len(h)))), StrToBig(q)).Text(16)), StrToBig(new(big.Int).Mul(StrToBig(rev(S0.Text(16), len(S0.Text(16)))), StrToBig(rev(d, len(d)))).Text(16))), StrToBig(q)).Text(16)))
	return S0_out + S1
}
