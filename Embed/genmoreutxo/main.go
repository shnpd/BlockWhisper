package main

import (
	"blockwhisper/blockchain"
	"blockwhisper/rpc"
	"log"
	"time"
)

func main() {
	client := rpc.InitClient("127.0.0.1:28335", "simnet")
	client.WalletPassphrase("ts0", 6000)
	address := "SXXfUx9qdszdhEgFJMq5625co9JrqbeRBv"
	address2 := "SjAUPA3G4LCeepgJDfGULEU3pdXyN4gCy4"
	//	生成address到address2的交易，再从address2转会address
	amount := 2510000000
	for i := 0; i < 150; i++ {
		_, err := blockchain.EntireSendTrans(client, address, address2, int64(amount), nil)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	//确认交易
	client.Generate(1)
	amount = 2400000000
	for i := 0; i < 150; i++ {
		_, err := blockchain.EntireSendTrans2(client, address2, address, int64(amount), nil)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	client.Generate(1)
}
