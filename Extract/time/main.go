package main

import (
	mycrypto "blockwhisper/crypto"
	"blockwhisper/rpc"
	"blockwhisper/share"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"log"
	"net/http"
	"time"
)

func main() {
	// 与地址池大小有关，需要遍历地址池查找特殊交易
	// 地址池有100个地址，遍历地址池获取交易，交易数量根据嵌入交易数量确定

	start := time.Now()
	var inputAddr string
	for i := 0; i < 100; i++ {
		//	1 查找源地址和目的地址都属于addresspool的交易
		inputAddr = "SXXfUx9qdszdhEgFJMq5625co9JrqbeRBv"
		client := rpc.InitClient("localhost:28335", "simnet")
		FilterTransByInputaddrByAPI(client, inputAddr)

	}
	//	2 根据时间间隔区分交易
	inputAddr = "SXXfUx9qdszdhEgFJMq5625co9JrqbeRBv"

	//	3 解码交易，源地址异或，拼接，解密
	t, _ := mycrypto.Encrypt([]byte("123"), share.Key)
	var binaryCipher string
	p := 32
	for j := 0; j < 422; j++ {
		amount := "12247"
		M := DecodeAmountm(amount, p)
		M2 := share.Byte2binary([]byte(inputAddr))[:p]
		M1 := share.XorBinaryString(M, M2)
		binaryCipher += M1
	}
	//share.Binary2byte(binaryCipher)
	_, err := mycrypto.Decrypt(t, share.Key)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)

	fmt.Println(duration)

}

// filterTransByInputaddrByAPI 模拟主网查询请求，任意发送一个地址的请求，直接返回隐蔽交易的id（本地simnet网络无法调用第三方api）
func FilterTransByInputaddrByAPI(client *rpcclient.Client, addr string) (*chainhash.Hash, error) {
	url := fmt.Sprintf("https://api.3xpl.com/bitcoin/address/%s?token=3A0_t3st3xplor3rpub11cb3t4efcd21748a5e&data=events", "bc1qd4ysezhmypwty5dnw7c8nqy5h5nxg0xqsvaefd0qn5kq32vwnwqqgv4rzr")
	resp, err := http.Get(url)
	if resp.StatusCode != 200 {
		log.Fatalf("请求失败：%s", resp.Status)
		return nil, err
	}
	defer resp.Body.Close()
	return nil, nil
}
