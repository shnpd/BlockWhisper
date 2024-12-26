package main

import (
	mycrypto "blockwhisper/crypto"
	"blockwhisper/share"
	"fmt"
	"log"
)

func main() {
	// 1. 选择输入输出地址
	inputAddr := share.AddressPool[0]
	//	2.
	msg := "hello"
	cipher, err := mycrypto.Encrypt([]byte(msg), share.Key)
	fmt.Println(cipher)
	if err != nil {
		log.Fatal("encrypto error:", err)
	}
	binaryCipher := share.Byte2binary(cipher)
	p := 32
	var Ms []string
	var M string
	for i := 0; i < len(binaryCipher); i += p {
		var M1 string
		if i+p >= len(binaryCipher) {
			M1 = binaryCipher[i:]
		} else {
			M1 = binaryCipher[i : i+p]
		}
		M2 := share.Byte2binary([]byte(inputAddr))[:p]
		M = share.XorBinaryString(M1, M2)
		Ms = append(Ms, M)
	}
	// 3 金额编码
	var amounts []string
	for _, v := range Ms {
		amounts = append(amounts, EncodeAmountm(v))
	}
	//——————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————
	//	1 查找源地址和目的地址都属于addresspool的交易
	//	2 根据时间间隔区分交易
	//	3 解码交易，源地址异或，拼接，解密
	var tbinaryCipher string
	for _, v := range amounts {
		tM := DecodeAmountm(v, p)
		tM2 := share.Byte2binary([]byte(inputAddr))[:p]
		tM1 := share.XorBinaryString(tM, tM2)
		tbinaryCipher += tM1
	}
	tcipher := share.Binary2byte(tbinaryCipher)
	plain, err := mycrypto.Decrypt(tcipher, share.Key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plain))
}
