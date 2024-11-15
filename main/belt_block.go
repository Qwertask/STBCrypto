package main

import (
	"fmt"
	"math"
	"strconv"
)

func Belt_block_encrypt(X string, K string) string {
	if len(X) > 32 || len(K) > 64 {
		panic("Проверьте длину параметров")
	} else if len(X) < 32 {
		for i := 0; len(X) != 32; i++ {
			X = "0" + X
		}
	} else if len(K) < 64 {
		for i := 0; len(K) != 64; i++ {
			K = "0" + K
		}
	}
	X_parts := split(X, 4)
	K_parts := split(K, 8)
	X1, X2, X3, X4 := splitAndReverse(X_parts[0]),
		splitAndReverse(X_parts[1]),
		splitAndReverse(X_parts[2]),
		splitAndReverse(X_parts[3])

	k := make([]int, 56)
	for i := 0; i < 56; i++ {
		k[i] = splitAndReverse(K_parts[i%8])
	}
	t := 0
	a, b, c, d, e := X1, X2, X3, X4, 0
	for i := 0; i < 8; i++ {
		b = b ^ G(SquareSum(a, k[7*(i+1)-6-1]), 5)
		c = c ^ G(SquareSum(d, k[7*(i+1)-5-1]), 21)
		a = SquareMinus(a, G(SquareSum(b, k[7*(i+1)-4-1]), 13))
		e = G(SquareSum(b, SquareSum(c, k[7*(i+1)-3-1])), 21) ^ (i + 1)
		b = SquareSum(b, e)
		c = SquareMinus(c, e)
		d = SquareSum(d, G(SquareSum(c, k[7*(i+1)-2-1]), 13))
		b = b ^ G(SquareSum(a, k[7*(i+1)-1-1]), 21)
		c = c ^ G(SquareSum(d, k[7*(i+1)-1]), 5)
		//
		t = a
		a = b
		b = t
		//
		t = c
		c = d
		d = t
		//
		t = b
		b = c
		c = t
	}
	res := fmt.Sprintf("%08x", splitAndReverse(b))
	res += fmt.Sprintf("%08x", splitAndReverse(d))
	res += fmt.Sprintf("%08x", splitAndReverse(a))
	res += fmt.Sprintf("%08x", splitAndReverse(c))
	return res
}

func splitAndReverse(n int) int {
	hexStr := fmt.Sprintf("%08x", n)
	parts := []string{}
	for i := len(hexStr); i > 0; i -= 2 {
		parts = append(parts, hexStr[i-2:i])
	}
	resultStr := ""
	for _, part := range parts {
		resultStr += part
	}
	resultInt, _ := strconv.ParseInt(resultStr, 16, 64)
	return int(resultInt)
}

func split(x string, parts int) []int {
	result := make([]int, 0)
	for i := 0; i < parts; i++ {
		t, _ := strconv.ParseInt(x[8*i:8+8*i], 16, 64)
		result = append(result, int(t))
	}
	return result
}

func SquareMinus(a int, b int) int {
	a_sk := 0
	b_sk := 0
	a_bin := toBinary32(a)
	b_bin := toBinary32(b)
	a_mas := make([]int, 0)
	b_mas := make([]int, 0)
	for i := 0; i < 4; i++ {
		a_mas = append(a_mas, binaryToInt(a_bin[8*i:8+8*i]))
		b_mas = append(b_mas, binaryToInt(b_bin[8*i:8+8*i]))
	}
	for i := 0; i < 4; i++ {
		a_sk += a_mas[i] * int(math.Pow(256, float64(3-i)))
		b_sk += b_mas[i] * int(math.Pow(256, float64(3-i)))
	}
	return (a_sk - b_sk + int(math.Pow(2, 32))) % int(math.Pow(2, 32))
}

func SquareSum(a int, b int) int {
	a_sk := 0
	b_sk := 0
	a_bin := toBinary32(a)
	b_bin := toBinary32(b)
	a_mas := make([]int, 0)
	b_mas := make([]int, 0)
	for i := 0; i < 4; i++ {
		a_mas = append(a_mas, binaryToInt(a_bin[8*i:8+8*i]))
		b_mas = append(b_mas, binaryToInt(b_bin[8*i:8+8*i]))
	}
	for i := 0; i < 4; i++ {
		a_sk += a_mas[i] * int(math.Pow(256, float64(3-i)))
		b_sk += b_mas[i] * int(math.Pow(256, float64(3-i)))
	}
	resab := (a_sk + b_sk)
	return resab % (1 << 32)
}

func G(u, r int) int {
	newH := ""
	u_bin := toBinary32(u)
	u_mas := make([]string, 0)
	for i := 0; i < 4; i++ {
		u_mas = append(u_mas, u_bin[8*i:8+8*i])
	}
	for i := 0; i < 4; i++ {
		ind1, ind2, _ := binary16ToTwoDecimals(u_mas[i])
		newH += H[ind1][ind2]
	}
	decimal, _ := strconv.ParseInt(newH, 16, 64)
	return RotHi(int(decimal), r)
}

func binary16ToTwoDecimals(binary string) (int, int, error) {
	if len(binary) != 8 {
		return 0, 0, fmt.Errorf("длина бинарного числа должна быть 16 бит")
	}

	part1 := binary[:4]
	part2 := binary[4:]
	decimal1, err1 := strconv.ParseInt(part1, 2, 16)
	if err1 != nil {
		return 0, 0, err1
	}
	decimal2, err2 := strconv.ParseInt(part2, 2, 16)
	if err2 != nil {
		return 0, 0, err2
	}
	return int(decimal1), int(decimal2), nil
}

func toBinary32(n int) string {
	binary := strconv.FormatInt(int64(n), 2)
	return fmt.Sprintf("%032s", binary)
}

func binaryToInt(binary string) int {
	str, _ := strconv.ParseInt(binary, 2, 32)
	return int(str)
}

var H = [16][16]string{
	{"B1", "94", "BA", "C8", "0A", "08", "F5", "3B", "36", "6D", "00", "8E", "58", "4A", "5D", "E4"},
	{"85", "04", "FA", "9D", "1B", "B6", "C7", "AC", "25", "2E", "72", "C2", "02", "FD", "CE", "0D"},
	{"5B", "E3", "D6", "12", "17", "B9", "61", "81", "FE", "67", "86", "AD", "71", "6B", "89", "0B"},
	{"5C", "B0", "C0", "FF", "33", "C3", "56", "B8", "35", "C4", "05", "AE", "D8", "E0", "7F", "99"},
	{"E1", "2B", "DC", "1A", "E2", "82", "57", "EC", "70", "3F", "CC", "F0", "95", "EE", "8D", "F1"},
	{"C1", "AB", "76", "38", "9F", "E6", "78", "CA", "F7", "C6", "F8", "60", "D5", "BB", "9C", "4F"},
	{"F3", "3C", "65", "7B", "63", "7C", "30", "6A", "DD", "4E", "A7", "79", "9E", "B2", "3D", "31"},
	{"3E", "98", "B5", "6E", "27", "D3", "BC", "CF", "59", "1E", "18", "1F", "4C", "5A", "B7", "93"},
	{"E9", "DE", "E7", "2C", "8F", "0C", "0F", "A6", "2D", "DB", "49", "F4", "6F", "73", "96", "47"},
	{"06", "07", "53", "16", "ED", "24", "7A", "37", "39", "CB", "A3", "83", "03", "A9", "8B", "F6"},
	{"92", "BD", "9B", "1C", "E5", "D1", "41", "01", "54", "45", "FB", "C9", "5E", "4D", "0E", "F2"},
	{"68", "20", "80", "AA", "22", "7D", "64", "2F", "26", "87", "F9", "34", "90", "40", "55", "11"},
	{"BE", "32", "97", "13", "43", "FC", "9A", "48", "A0", "2A", "88", "5F", "19", "4B", "09", "A1"},
	{"7E", "CD", "A4", "D0", "15", "44", "AF", "8C", "A5", "84", "50", "BF", "66", "D2", "E8", "8A"},
	{"A2", "D7", "46", "52", "42", "A8", "DF", "B3", "69", "74", "C5", "51", "EB", "23", "29", "21"},
	{"D4", "EF", "D9", "B4", "3A", "62", "28", "75", "91", "14", "10", "EA", "77", "6C", "DA", "1D"},
}

func RotHi(u int, n int) int {
	res := u
	for i := 0; i < n; i++ {
		res = ShHi(res, 1) ^ ShLo(res, 31)
	}
	return res % int(math.Pow(2, 32))
}

func ShLo(u int, n int) int {
	for i := 0; i < 4; i++ {

	}
	for i := 0; i < n; i++ {
		u = u / 2
	}
	return u % int(math.Pow(2, 32))
}

func ShHi(u int, n int) int {
	for i := 0; i < n; i++ {
		u = 2 * u
	}
	return u % int(math.Pow(2, 32))
}

func Belt_block_decrypt(X string, K string) string {
	if len(X) > 32 || len(K) > 64 {
		panic("Проверьте длину параметров")
	} else if len(X) < 32 {
		for i := 0; len(X) != 32; i++ {
			X = "0" + X
		}
	} else if len(K) < 64 {
		for i := 0; len(K) != 64; i++ {
			K = "0" + K
		}
	}
	X_parts := split(X, 4)
	K_parts := split(K, 8)
	X1, X2, X3, X4 := splitAndReverse(X_parts[0]),
		splitAndReverse(X_parts[1]),
		splitAndReverse(X_parts[2]),
		splitAndReverse(X_parts[3])

	k := make([]int, 56)
	for i := 0; i < 56; i++ {
		k[i] = splitAndReverse(K_parts[i%8])
	}
	t := 0
	a, b, c, d, e := X1, X2, X3, X4, 0
	for i := 7; i >= 0; i-- {
		b = b ^ G(SquareSum(a, k[7*(i+1)-1]), 5)
		c = c ^ G(SquareSum(d, k[7*(i+1)-1-1]), 21)
		a = SquareMinus(a, G(SquareSum(b, k[7*(i+1)-2-1]), 13))
		e = G(SquareSum(b, SquareSum(c, k[7*(i+1)-3-1])), 21) ^ (i + 1)
		b = SquareSum(b, e)
		c = SquareMinus(c, e)
		d = SquareSum(d, G(SquareSum(c, k[7*(i+1)-4-1]), 13))
		b = b ^ G(SquareSum(a, k[7*(i+1)-5-1]), 21)
		c = c ^ G(SquareSum(d, k[7*(i+1)-6-1]), 5)
		//
		t = a
		a = b
		b = t
		//
		t = c
		c = d
		d = t
		//
		t = a
		a = d
		d = t
	}
	res := fmt.Sprintf("%08x", splitAndReverse(c))
	res += fmt.Sprintf("%08x", splitAndReverse(a))
	res += fmt.Sprintf("%08x", splitAndReverse(d))
	res += fmt.Sprintf("%08x", splitAndReverse(b))
	return res
}
