package main

import "math/big"

type EdEC struct {
	d, p *big.Int
}

func ConvertWeierstrassToEdwards(ec EC) EdEC {
	// Примерные преобразования для полей
	p := StrToBig(ec.p)
	a := StrToBig(ec.a)
	b := StrToBig(ec.b)
	// Вычисление коэффициента d для кривой Эдвардса (здесь предположим, что у нас подходящая форма)
	d := new(big.Int).Mod(new(big.Int).Sub(new(big.Int).Mul(big.NewInt(3), a), new(big.Int).Mul(big.NewInt(2), b)), p)
	return EdEC{d: d, p: p}
}

func (ec EdEC) AddEdwardsPoints(P, Q point) point {
	x1, y1 := StrToBig(P.x), StrToBig(P.y)
	x2, y2 := StrToBig(Q.x), StrToBig(Q.y)
	p := ec.p
	d := ec.d

	one := big.NewInt(1)

	// x3 = (x1*y2 + y1*x2) / (1 + d*x1*x2*y1*y2)
	numX := new(big.Int).Add(new(big.Int).Mul(x1, y2), new(big.Int).Mul(y1, x2))
	denX := new(big.Int).Add(one, new(big.Int).Mul(d, new(big.Int).Mul(x1, new(big.Int).Mul(x2, new(big.Int).Mul(y1, y2)))))
	x3 := new(big.Int).Mod(new(big.Int).Mul(numX, modInverse(denX, p)), p)

	// y3 = (y1*y2 - x1*x2) / (1 - d*x1*x2*y1*y2)
	numY := new(big.Int).Sub(new(big.Int).Mul(y1, y2), new(big.Int).Mul(x1, x2))
	denY := new(big.Int).Sub(one, new(big.Int).Mul(d, new(big.Int).Mul(x1, new(big.Int).Mul(x2, new(big.Int).Mul(y1, y2)))))
	y3 := new(big.Int).Mod(new(big.Int).Mul(numY, modInverse(denY, p)), p)

	return point{x: x3.Text(16), y: y3.Text(16)}
}

func (ec EdEC) mult_edwards(P point, k *big.Int) point {

	R := point{"0", "1"} // Нулевая точка на кривой Эдвардса
	N := P

	for k.Sign() != 0 {
		if k.Bit(0) == 1 {
			R = ec.AddEdwardsPoints(R, N)
		}
		N = ec.AddEdwardsPoints(N, N)
		k.Rsh(k, 1)
	}
	return R
}
func ConvertEdwardsToWeierstrass(edEC EdEC, P point) point {
	x := StrToBig(P.x)
	y := StrToBig(P.y)
	p := edEC.p

	weierstrassX := new(big.Int).Mod(new(big.Int).Mul(x, modInverse(new(big.Int).Sub(big.NewInt(1), y), p)), p)
	weierstrassY := new(big.Int).Mod(new(big.Int).Mul(y, modInverse(new(big.Int).Sub(big.NewInt(1), y), p)), p)

	return point{x: weierstrassX.Text(16), y: weierstrassY.Text(16)}
}

func modInverse(a, m *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, m)
}
