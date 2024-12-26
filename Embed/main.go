package main

import (
	"blockwhisper/blockchain"
	mycrypto "blockwhisper/crypto"
	"blockwhisper/rpc"
	"blockwhisper/share"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func init() {
}

func chooseLen() int {
	amountLength := []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	weights := []float64{0.00031, 0.08465, 0.15278, 0.47345, 0.13649, 0.08014, 0.03909, 0.02518, 0.0061, 0.00181} // 注意概率之和必须为 1
	rand.Seed(time.Now().UnixNano())                                                                              // 设置随机种子

	// 生成累积权重
	cumulativeWeights := make([]float64, len(weights))
	cumulativeWeights[0] = weights[0]
	for i := 1; i < len(weights); i++ {
		cumulativeWeights[i] = cumulativeWeights[i-1] + weights[i]
	}

	// 生成一个 0 到 1 的随机数
	r := rand.Float64()
	// 找到随机数对应的区间
	selectedIndex := 0
	for i, weight := range cumulativeWeights {
		if r <= weight {
			selectedIndex = i
			break
		}
	}

	selectedValue := amountLength[selectedIndex]
	return selectedValue
}
func main() {
	client := rpc.InitClient("127.0.0.1:28335", "simnet")
	client.WalletPassphrase("ts0", 6000)
	addresses, err := client.GetAddressesByAccount("default")
	if err != nil {
		log.Fatal(err)
	}
	share.AddressPool = addresses
	for temp := 0; temp < 5; temp++ {
		// 1. 选择输入输出地址
		inputAddr := "SXXfUx9qdszdhEgFJMq5625co9JrqbeRBv"
		outputAddr := "SXXfUx9qdszdhEgFJMq5625co9JrqbeRBv"
		//	2.
		msg := "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
		cipher, err := mycrypto.Encrypt([]byte(msg), share.Key)
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
			amount := EncodeAmountm(v)
			chlen := chooseLen()
			for len(amount) > chlen {
				amounts = append(amounts, amount[:chlen])
				amount = amount[chlen:]
			}
			if len(amount) >= 4 {
				amounts = append(amounts, amount)
			}
		}
		// 4 生成混淆交易
		// 5 发送交易
		var amountInt []int
		for _, v := range amounts {
			amountint, _ := strconv.Atoi(v)
			amountInt = append(amountInt, amountint)
		}
		// TODO：从大到小排序
		// 使用 sort.Slice 进行降序排序
		sort.Slice(amountInt, func(i, j int) bool {
			return amountInt[i] > amountInt[j] // 降序排序
		})

		start := time.Now()
		lenamountint := len(amountInt)
		fmt.Println(lenamountint)
		fmt.Println(lenamountint * 4 / 3)
		for _, amount := range amountInt {
			_, err = blockchain.EntireSendTrans(client, inputAddr, outputAddr, int64(amount), nil)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(100 * time.Millisecond)
		}
		duration := time.Since(start)

		fmt.Println()
		duration -= time.Duration(lenamountint) * (100 * time.Millisecond)
		//fmt.Println(duration)
		//	TODO：最终时间要乘4/3，因为包含混淆交易
		fmt.Println(duration)
		fmt.Println(duration * 4 / 3)
	}
}
