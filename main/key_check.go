package main

import (
	"math/big"
)

func key_check(Q point, p string, a string, b string) bool {
	Qx, _ := new(big.Int).SetString(Q.x, 16)
	Qy, _ := new(big.Int).SetString(Q.y, 16)
	P_big := StrToBig(p)
	l_big := StrToBig(p)
	r_big := StrToBig("0")
	y_2 := new(big.Int).Exp(Qy, new(big.Int).SetInt64(int64(2)), P_big)

	x_3 := new(big.Int).Exp(Qx, new(big.Int).SetInt64(int64(3)), P_big)
	ax := new(big.Int).Mul(StrToBig(a), Qx)
	x_3_ax := new(big.Int).Add(x_3, ax)
	x3_ax_b := new(big.Int).Mod(new(big.Int).Add(x_3_ax, StrToBig(b)), P_big)
	if Qx.Cmp(r_big) != -1 && Qy.Cmp(r_big) != -1 && Qx.Cmp(l_big) == -1 && Qy.Cmp(l_big) == -1 && y_2.Cmp(x3_ax_b) == 0 {
		return true
	}
	return false
}

//87194355096605130634028764141169641745691786102469501246661297785830520686053
//87194355096605130634028764141169641745691786102469501246661297785830520686053
//
