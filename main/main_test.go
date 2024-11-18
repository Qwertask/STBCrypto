package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestBelt_block_encrypt(t *testing.T) {
	key_enc := "E9DEE72C8F0C0FA62DDB49F46F73964706075316ED247A3739CBA38303A98BF6"
	X_enc := "B194BAC80A08F53B366D008E584A5DE4"
	test := Belt_block_encrypt(X_enc, key_enc)
	res := strings.ToLower("69CCA1C93557C9E3D66BC3E0FA88FA6E")
	if test != res {
		t.Error("TestBelt_block_encrypt", test, res)
	}
}

func TestBelt_block_decrypt(t *testing.T) {
	key_enc := "92BD9B1CE5D141015445FBC95E4D0EF2682080AA227D642F2687F93490405511"
	X_enc := "E12BDC1AE28257EC703FCCF095EE8DF1"
	test := Belt_block_decrypt(X_enc, key_enc)
	res := strings.ToLower("0DC5300600CAB840B38448E5E993F421")
	if test != res {
		t.Error("TestBelt_block_decrypt", test, res)
	}
}

func TestBelt_hash(t *testing.T) {
	X := "B194BAC80A08F53B366D008E58"
	Y := belt_hash(X)
	res := strings.ToLower("ABEF9725D4C5A83597A367D14494CC2542F20F659DDFECC961A3EC550CBA8C75")
	if Y != res {
		t.Error("TestBelt_hash", res, Y)
	}
	X = "B194BAC80A08F53B366D008E584A5DE48504FA9D1BB6C7AC252E72C202FDCE0D5BE3D61217B96181FE6786AD716B890B"
	Y = belt_hash(X)
	res = strings.ToLower("9D02EE446FB6A29FE5C982D4B13AF9D3E90861BC4CEF27CF306BFB0B174A154A")
	check := Y != res
	if check {
		t.Error("TestBelt_hash", res, Y)
	}
}

func TestBelt_compress(t *testing.T) {
	X := "B194BAC80A08F53B366D008E584A5DE48504FA9D1BB6C7AC252E72C202FDCE0D5BE3D61217B96181FE6786AD716B890B5CB0C0FF33C356B835C405AED8E07F99"
	s_res, y_res := belt_compress(X)
	S := strings.ToLower("46FE7425C9B181EB41DFEE3E72163D5A")
	Y := strings.ToLower("ED2F5481D593F40D87FCE37D6BC1A2E1B7D1A2CC975C82D3C0497488C90D99D8")
	if s_res != S && y_res != Y {
		t.Error("TestBelt_compress", s_res, y_res)
	}
}

func TestMult(t *testing.T) {
	Q1 := mult_chain(point{
		x: fmt.Sprintf("%x", 11),
		y: fmt.Sprintf("%x", 17),
	}, fmt.Sprintf("%x", 7), fmt.Sprintf("%x", 11150), fmt.Sprintf("%x", 2), fmt.Sprintf("%x", 97))
	Q2 := mult_naf(point{
		x: fmt.Sprintf("%x", 11),
		y: fmt.Sprintf("%x", 17),
	}, fmt.Sprintf("%x", 7), fmt.Sprintf("%x", 11150), fmt.Sprintf("%x", 2), fmt.Sprintf("%x", 97))
	Q3 := mult_binary(point{
		x: fmt.Sprintf("%x", 11),
		y: fmt.Sprintf("%x", 17),
	}, fmt.Sprintf("%x", 7), fmt.Sprintf("%x", 11150), fmt.Sprintf("%x", 2), fmt.Sprintf("%x", 97))
	Q4 := mult_window(point{
		x: fmt.Sprintf("%x", 11),
		y: fmt.Sprintf("%x", 17),
	}, fmt.Sprintf("%x", 7), fmt.Sprintf("%x", 11150), fmt.Sprintf("%x", 2), fmt.Sprintf("%x", 97))

	Q5 := mult_sl_window(point{
		x: fmt.Sprintf("%x", 11),
		y: fmt.Sprintf("%x", 17),
	}, fmt.Sprintf("%x", 7), fmt.Sprintf("%x", 11150), fmt.Sprintf("%x", 2), fmt.Sprintf("%x", 97))

	if Q1 != Q2 && Q1 != Q3 && Q4 != Q1 && Q1 != Q5 {
		t.Error("TestMult", Q1, Q2, Q3, Q4, Q5)
	}
}

func TestBase_point(t *testing.T) {
	G_test := base_point(ECEx.b, ECEx.p)

	if G_test.Equal(ECEx.yG) {
		t.Error("TestBase_point", G_test, ECEx.yG)
	}
}

func TestKeys(t *testing.T) {
	d, Q := key_gen(ECEx.q, ECEx.yG, ECEx.a, ECEx.p)
	check := key_check(Q, ECEx.p, ECEx.a, ECEx.b)
	if !check {
		t.Error("TestKeys", d, Q)
	}
	d, Q = key_gen(ECEx.q, ECEx.yG, ECEx.a, ECEx.p)
	check = key_check(Q, ECEx.p, ECEx.a, ECEx.b)
	if !check {
		t.Error("TestKeys", d, Q)
	}
	d, Q = key_gen(ECEx.q, ECEx.yG, ECEx.a, ECEx.p)
	check = key_check(Q, ECEx.p, ECEx.a, ECEx.b)
	if !check {
		t.Error("TestKeys", d, Q)
	}
	d, Q = key_gen(ECEx.q, ECEx.yG, ECEx.a, ECEx.p)
	check = key_check(Q, ECEx.p, ECEx.a, ECEx.b)
	if !check {
		t.Error("TestKeys", d, Q)
	}
}

func TestOneshot_key_gen(t *testing.T) {
	d := "1F66B5B84B7339674533F0329C74F21834281FED0732429E0C79235FC273E269"
	h := "ABEF9725D4C5A83597A367D14494CC2542F20F659DDFECC961A3EC550CBA8C75"
	k := oneshot_key_gen(ECEx.q, d, h)

	if strings.ToLower(k) != strings.ToLower("829614D8411DBBC4E1F2471A4004586440FD8C9553FAB6A1A45CE417AE97111E") {
		t.Error("TestOneshot_key_gen", k, h)
	}

}

// Check eds_gen.go
func TestEds_gen(t *testing.T) {
	d := "1F66B5B84B7339674533F0329C74F21834281FED0732429E0C79235FC273E269"
	X := "B194BAC80A08F53B366D008E58"
	S := eds_gen(ECEx.p, ECEx.a, ECEx.q, ECEx.yG, X, d)
	if strings.ToLower(S) != strings.ToLower("E36B7F0377AE4C524027C387FADF1B20CE72F1530B71F2B5FD3A8C584FE2E1AED20082E30C8AF65011F4FB54649DFD3D") {
		t.Error("TestEds_gen", S)
	}
}

func TestEds_check(t *testing.T) {
	Q := point{
		x: rev("BD1A5650179D79E03FCEE49D4C2BD5DDF54CE46D0CF11E4FF87BF7A890857FD0", 64),
		y: rev("7AC6A60361E8C8173491686D461B2826190C2EDA5909054A9AB84D2AB9D99A90", 64),
	}
	X := "B194BAC80A08F53B366D008E584A5DE48504FA9D1BB6C7AC252E72C202FDCE0D5BE3D61217B96181FE6786AD716B890B"
	S := "47A63C8B9C936E94B5FAB3D9CBD78366290F3210E163EEC8DB4E921E8479D4138F112CC23E6DCE65EC5FF21DF4231C28"

	if !eds_check(S, ECEx.yG, Q, X, ECEx.q, ECEx.a, ECEx.p) {
		t.Error("TestEds_check")
	}
}
