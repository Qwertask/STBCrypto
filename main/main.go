package main

import (
	"fmt"
)

var ECEx = STBparam

func main() {
	fmt.Println(base_point(fmt.Sprintf("%x", 378999), fmt.Sprintf("%x", 7919)))
	d, Q := key_gen(ECEx.q, ECEx.yG, ECEx.a, ECEx.p)
	fmt.Println("KEY_CHECK:", key_check(Q, ECEx.p, ECEx.a, ECEx.b))
	fmt.Println(mult_naf(point{
		x: fmt.Sprintf("%x", 11),
		y: fmt.Sprintf("%x", 17),
	}, fmt.Sprintf("%x", 7), fmt.Sprintf("%x", 50), fmt.Sprintf("%x", 2), fmt.Sprintf("%x", 97)))
	fmt.Println(StrToBig(ECEx.q))
	fmt.Println(eds_gen(ECEx.p, ECEx.a, ECEx.q, ECEx.yG, "B194BAC80A08F53B366D008E58", "1F66B5B84B7339674533F0329C74F21834281FED0732429E0C79235FC273E269"))
	fmt.Println(Legendre("11", "7"))
	Q3 := mult_binary(point{
		x: fmt.Sprintf("%x", 11),
		y: fmt.Sprintf("%x", 17),
	}, fmt.Sprintf("%x", 7), fmt.Sprintf("%x", 50), fmt.Sprintf("%x", 2), fmt.Sprintf("%x", 97))
	fmt.Println("POINT:", Q3)
	Q3 = mult_sl_window(point{
		x: ECEx.yG.x,
		y: ECEx.yG.y,
	}, fmt.Sprintf("%x", 7), ECEx.q, ECEx.a, ECEx.p)
	fmt.Println("POINT:", Q3)
	d = "1F66B5B84B7339674533F0329C74F21834281FED0732429E0C79235FC273E269"
	fmt.Println(d)
	d = rev(d, len(d))
	fmt.Println(d)
	//fmt.Println(oneshot_key_gen(ECEx.q, d, "ABEF9725D4C5A83597A367D14494CC2542F20F659DDFECC961A3EC550CBA8C75"))
	fmt.Println(eds_gen(ECEx.p, ECEx.a, ECEx.q, ECEx.yG, "B194BAC80A08F53B366D008E58", "1F66B5B84B7339674533F0329C74F21834281FED0732429E0C79235FC273E269"))
	gmul := "c6112055501a9168c11815f335b4cde7781a8d7e8b5cf6341dcbed290b6e3172"
	qmul := "47a63c8b9c936e94b5fab3d9cbd7836601"
	Q = point{
		x: "bd1a5650179d79e03fcee49d4c2bd5ddf54ce46d0cf11e4ff87bf7a890857fd0",
		y: "7ac6a60361e8c8173491686d461b2826190c2eda5909054a9ab84d2ab9d99a90",
	}
	GP := point{
		x: "0",
		y: "936A510418CF291E52F608C4663991785D83D651A3C9E45C9FD616FB3CFCF76B",
	}
	Q.x = rev(Q.x, len(Q.x))
	Q.y = rev(Q.y, len(Q.y))
	GP.y = rev(GP.y, len(GP.y))
	fmt.Println(Q, qmul, GP, gmul)
	fmt.Println(Add(mult_naf(Q, rev(qmul, len(qmul)), ECEx.q, ECEx.a, ECEx.p), mult_window(GP, rev(gmul, len(gmul)), ECEx.q, ECEx.a, ECEx.p), ECEx.a, ECEx.p))

	fmt.Println(eds_check("47A63C8B9C936E94B5FAB3D9CBD78366290F3210E163EEC8DB4E921E8479D4138F112CC23E6DCE65EC5FF21DF4231C28", ECEx.yG, point{
		x: rev("BD1A5650179D79E03FCEE49D4C2BD5DDF54CE46D0CF11E4FF87BF7A890857FD0", 64),
		y: rev("7AC6A60361E8C8173491686D461B2826190C2EDA5909054A9AB84D2AB9D99A90", 64),
	}, "B194BAC80A08F53B366D008E584A5DE48504FA9D1BB6C7AC252E72C202FDCE0D5BE3D61217B96181FE6786AD716B890B", ECEx.q, ECEx.a, ECEx.p))
	d, Q = key_gen(ECEx.q, ECEx.yG, ECEx.a, ECEx.p)

	fmt.Println("KEY:", d, Q)
	fmt.Println(key_check(Q, ECEx.p, ECEx.a, ECEx.b))
}
