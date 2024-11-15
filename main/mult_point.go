package main

import (
	"math/big"
	"sort"
)

func Add(p point, q point, a string, m string) point {
	if p.x == "" && p.y == "" {
		return q
	} else if q.x == "" && q.y == "" {
		return p
	}
	l := new(big.Int)
	if p.x != q.x {
		l = new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Sub(StrToBig(q.y), StrToBig(p.y)), new(big.Int).Exp(new(big.Int).Sub(StrToBig(q.x), StrToBig(p.x)), new(big.Int).Sub(StrToBig(m), new(big.Int).SetInt64(int64(2))), StrToBig(m))), StrToBig(m))
	} else if p.Equal(q) && p.y != "0" {
		l = new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Add(StrToBig(a), new(big.Int).Mul(new(big.Int).SetInt64(int64(3)), new(big.Int).Exp(StrToBig(p.x), new(big.Int).SetInt64(int64(2)), StrToBig(m)))), new(big.Int).Exp(new(big.Int).Mul(StrToBig(p.y), new(big.Int).SetInt64(int64(2))), new(big.Int).SetInt64(int64(-1)), StrToBig(m))), StrToBig(m))
	} else {
		return point{
			x: "",
			y: "",
		}
	}
	l2 := new(big.Int).Exp(l, new(big.Int).SetInt64(int64(2)), StrToBig(m))
	x1_x2 := new(big.Int).Sub(new(big.Int).Neg(StrToBig(p.x)), StrToBig(q.x))

	x3 := new(big.Int).Mod(new(big.Int).Add(l2, x1_x2), StrToBig(m))
	x1_x3 := new(big.Int).Sub(StrToBig(p.x), x3)
	l_x1_x3 := new(big.Int).Mul(l, x1_x3)
	y3 := new(big.Int).Mod(new(big.Int).Sub(l_x1_x3, StrToBig(p.y)), StrToBig(m))
	return point{x: x3.Text(16), y: y3.Text(16)}
}

func (p *point) Equal(q point) bool {
	return p.y == q.y && p.x == q.x
}

func StrToBig(q string) *big.Int {
	bigstr, _ := new(big.Int).SetString(q, 16)
	return bigstr
}

func BinToBig(q string) *big.Int {
	bigstr, _ := new(big.Int).SetString(q, 2)
	return bigstr
}

func mult_binary(p point, d string, q string, a string, m string) point {
	var u point
	var v = p
	if new(big.Int).Mod(StrToBig(d), StrToBig(q)).Cmp(StrToBig("0")) == 0 {
		return point{
			x: "",
			y: "",
		}
	}
	d = new(big.Int).Mod(StrToBig(d), StrToBig(q)).Text(16)
	d_big, _ := new(big.Int).SetString(d, 16)
	d_bin_str := d_big.Text(2)
	for i := len(d_bin_str) - 1; i >= 0; i-- {
		if d_bin_str[i] != '0' {
			u = Add(u, v, a, m)
		}
		v = Add(v, v, a, m)
	}
	return u
}

func naf(d string) []int {
	nafSequence := []int{}
	d_big := StrToBig(d)
	for d_big.Cmp(new(big.Int).SetInt64(int64(1))) == 1 {
		if new(big.Int).Mod(d_big, new(big.Int).SetInt64(int64(2))).Cmp(new(big.Int).SetInt64(int64(1))) == 0 {
			nafDigit := 2 - int(new(big.Int).Mod(d_big, new(big.Int).SetInt64(int64(4))).Int64())
			nafSequence = append(nafSequence, nafDigit)
			d_big.Sub(d_big, new(big.Int).SetInt64(int64(nafDigit)))
		} else {
			nafSequence = append(nafSequence, 0)
		}
		d_big.Div(d_big, new(big.Int).SetInt64(int64(2)))
	}

	nafSequence = append(nafSequence, 1)
	for i, j := 0, len(nafSequence)-1; i < j; i, j = i+1, j-1 {
		nafSequence[i], nafSequence[j] = nafSequence[j], nafSequence[i]
	}

	return nafSequence
}

func mult_naf(v point, d, q, a, p string) point {
	var u point
	binaryD := naf(d)

	if new(big.Int).Mod(StrToBig(d), StrToBig(q)).Cmp(StrToBig("0")) == 0 {
		return point{
			x: "",
			y: "",
		}
	}
	d = new(big.Int).Mod(StrToBig(d), StrToBig(q)).Text(16)
	//биты уже заранее развернуты в naf
	for _, bit := range binaryD {
		u = Add(u, u, a, p)
		if bit == 1 {
			u = Add(u, v, a, p)
		}
		if bit == -1 {
			u = Add(u, point{
				x: new(big.Int).Mod(StrToBig(v.x), StrToBig(p)).Text(16),
				y: new(big.Int).Mod(new(big.Int).Neg(StrToBig(v.y)), StrToBig(p)).Text(16),
			}, a, p)
		}
	}

	return u
}
func GenerateAdditiveChain(d_str string) []*big.Int {
	d := StrToBig(d_str)
	chain := []*big.Int{big.NewInt(1)}
	for chain[len(chain)-1].Cmp(d) < 0 {
		var nextValue *big.Int
		if len(chain) == 1 {
			nextValue = big.NewInt(2)
		} else {
			nextValue = new(big.Int).Add(chain[len(chain)-1], chain[len(chain)-2])
		}

		if nextValue.Cmp(d) > 0 {
			temp := new(big.Int).Sub(d, chain[len(chain)-1])
			nextValue = new(big.Int).Set(d)

			for !contains(chain, temp) {
				insertSorted(chain, temp)
				index := indexOf(chain, temp)
				temp = new(big.Int).Sub(temp, chain[(index-1+len(chain))%len(chain)])
			}
		}

		chain = append(chain, nextValue)
	}
	return chain
}

// FindIndexes finds the indexes for the additive chain.
func FindIndexes(nums []*big.Int) [][2]int {
	indexMap := map[string]int{"1": 0}
	result := [][2]int{}

	for i := 1; i < len(nums); i++ {
		target := nums[i]
		for j := 0; j < i; j++ {
			k := new(big.Int).Sub(target, nums[j])
			if idx, found := indexMap[k.String()]; found {
				result = append(result, [2]int{idx, j})
				break
			}
		}
		indexMap[nums[i].String()] = i
	}
	return result
}

func mult_chain(v point, d, q, a, p string) point {
	if new(big.Int).Mod(StrToBig(d), StrToBig(q)).Cmp(StrToBig("0")) == 0 {
		return point{
			x: "",
			y: "",
		}
	}
	d = new(big.Int).Mod(StrToBig(d), StrToBig(q)).Text(16)
	p_c := make([]point, 0)
	p_c = append(p_c, v)
	chain := GenerateAdditiveChain(d)
	ind := FindIndexes(chain)

	for i, _ := range ind {
		p_c = append(p_c, Add(p_c[ind[i][0]], p_c[ind[i][1]], a, p))
	}
	return p_c[len(p_c)-1]
}
func contains(slice []*big.Int, value *big.Int) bool {
	for _, v := range slice {
		if v.Cmp(value) == 0 {
			return true
		}
	}
	return false
}

func insertSorted(slice []*big.Int, value *big.Int) {
	slice = append(slice, value)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Cmp(slice[j]) < 0
	})
}

func indexOf(slice []*big.Int, value *big.Int) int {
	for i, v := range slice {
		if v.Cmp(value) == 0 {
			return i
		}
	}
	return -1
}

func mult_window(v point, d, q, a, p string) point {
	w := 8
	dBig := new(big.Int)
	dBig.SetString(d, 16)

	qBig := new(big.Int)
	qBig.SetString(q, 16)

	zero := big.NewInt(0)
	if new(big.Int).Mod(dBig, qBig).Cmp(zero) == 0 {
		return point{x: "", y: ""}
	}

	dBig = new(big.Int).Mod(dBig, qBig)
	l := dBig.BitLen()
	s := (l + w - 1) / w // количество шагов

	dMas := convertToBase(dBig, 1<<w)

	pMas := make([]point, 1<<w)
	pMas[0] = point{x: "", y: ""} // P0 = O (нулевая точка)
	pMas[1] = v                   // P1 = v

	for i := 2; i < len(pMas); i++ {
		pMas[i] = Add(pMas[i-1], v, a, p)
	}

	u := point{x: "", y: ""}
	for i := s - 1; i >= 0; i-- {
		for j := 0; j < w; j++ {
			u = Add(u, u, a, p)
		}
		u = Add(u, pMas[dMas[i]], a, p)
	}

	return u
}
func mult_sl_window(v point, d, q, a, p string) point {
	w := 8
	dBig := new(big.Int)
	dBig.SetString(d, 16)

	qBig := new(big.Int)
	qBig.SetString(q, 16)

	zero := big.NewInt(0)
	if new(big.Int).Mod(dBig, qBig).Cmp(zero) == 0 {
		return point{x: "", y: ""}
	}

	dBig = new(big.Int).Mod(dBig, qBig)
	l := dBig.BitLen()

	// Инициализация массива точек
	pMas := make([]point, 1<<(w-1))
	pMas[0] = point{x: "", y: ""} // P0 = O (нулевая точка)
	pMas[1] = v                   // P1 = v
	pMas[2] = Add(v, v, a, p)     // P2 = 2v
	for i := 3; i < len(pMas); i++ {
		pMas[i] = Add(pMas[i-1], v, a, p)
	}

	u := point{x: "", y: ""}
	for i := l - 1; i >= 0; {
		if dBig.Bit(i) == 0 {
			u = Add(u, u, a, p)
			i--
		} else {
			j := max(0, i-w+1)
			for dBig.Bit(j) == 0 {
				j++
			}
			k := 0
			for m := j; m <= i; m++ {
				k = (k << 1) | int(dBig.Bit(m))
			}
			for m := 0; m <= i-j; m++ {
				u = Add(u, u, a, p)
			}
			u = Add(u, pMas[k>>1], a, p)
			i = j - 1
		}
	}

	return u
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func convertToBase(n *big.Int, base int) []int64 {
	if n.Cmp(big.NewInt(0)) == 0 {
		return []int64{0}
	}

	baseBig := big.NewInt(int64(base))
	zero := big.NewInt(0)
	remainder := new(big.Int)
	quotient := new(big.Int).Set(n)
	var remainders []int64

	for quotient.Cmp(zero) > 0 {
		quotient.DivMod(quotient, baseBig, remainder)
		remainders = append(remainders, remainder.Int64())
	}
	for len(remainders) < 16 {
		remainders = append(remainders, 0)
	}

	return remainders
}
