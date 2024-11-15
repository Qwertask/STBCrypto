package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const OIDh = "06092A7000020022651F51"

func oneshot_key_gen(q, d, H string) string {
	n := 2
	t := ""
	o := belt_hash(OIDh + d + t)
	r := H
	s := ""
	k := ""
	r_i := splitBinaryString(fmt.Sprintf("%b", StrToBig(r)), 128)
	for i := range r_i {
		r_i[i] = BinToBig(r_i[i]).Text(16)
	}
	for i := 1; ; i++ { //len[r_i] = 2
		if n == 2 {
			s = r_i[0]
		}
		r_i[0] = new(big.Int).Xor(StrToBig(Belt_block_encrypt(s, o)), new(big.Int).Xor(StrToBig(r_i[len(r_i)-1]), new(big.Int).Mod(StrToBig(rev(strconv.Itoa(i), 32)), new(big.Int).Exp(StrToBig("2"), StrToBig(fmt.Sprintf("%x", 128)), nil)))).Text(16)
		r_i[1] = s
		r = ""
		for j := range r_i {
			r += r_i[j]
		}
		if (i)%(2*n) == 0 && StrToBig(rev(r, len(r))).Cmp(StrToBig("1")) == 1 && StrToBig(rev(r, len(r))).Cmp(new(big.Int).Sub(StrToBig(q), StrToBig("1"))) == -1 {
			k = r
			break
		}
	}
	return k
}

func rev(d string, l int) string {
	for len(d)%2 != 0 || len(d) < l {
		d = "0" + d
	}
	temp := ""
	l2 := len(d)
	for len(temp) != l && len(temp) != l2 {
		temp += d[len(d)-2:]
		d = d[:len(d)-2]
	}
	return temp
}

func splitBinaryString(binaryStr string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(binaryStr); i += chunkSize {
		end := i + chunkSize
		if end > len(binaryStr) {
			end = len(binaryStr)
		}
		chunks = append(chunks, binaryStr[i:end])
	}
	return chunks
}
