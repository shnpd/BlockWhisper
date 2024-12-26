package blockchain

import (
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
)

// EntireSendTrans 完整交易发送，包括交易生成、交易签名、交易广播，最终返回广播的交易id
func EntireSendTrans2(client *rpcclient.Client, sourceAddr, destAddr string, amount int64, embedMsg *[]byte) (*chainhash.Hash, error) {
	rawTx, err := GenerateTrans2(client, sourceAddr, destAddr, amount)
	if err != nil {
		return nil, err
	}
	signTx, err := SignTrans(client, rawTx, embedMsg)
	if err != nil {
		return nil, err
	}
	transId, err := BroadTrans(client, signTx)
	if err != nil {
		return nil, err
	}
	return transId, nil
}

func EntireSendTrans(client *rpcclient.Client, sourceAddr, destAddr string, amount int64, embedMsg *[]byte) (*chainhash.Hash, error) {
	rawTx, err := GenerateTrans(client, sourceAddr, destAddr, amount)
	if err != nil {
		return nil, err
	}
	signTx, err := SignTrans(client, rawTx, embedMsg)
	if err != nil {
		return nil, err
	}
	transId, err := BroadTrans(client, signTx)
	if err != nil {
		return nil, err
	}
	return transId, nil
}

// GenerateTrans 生成sourceAddr到destAddr的原始交易
func GenerateTrans(client *rpcclient.Client, sourceAddr, destAddr string, amount int64) (*wire.MsgTx, error) {
	// 筛选源地址的UTXO
	utxos, _ := client.ListUnspent()
	var sourceUTXO btcjson.ListUnspentResult
	for i, utxo := range utxos {
		if utxo.Address == sourceAddr {
			sourceUTXO = utxo
			break
		}
		if i == len(utxos)-1 {
			return nil, fmt.Errorf("UTXO not found")
		}
	}

	// 构造输入
	var inputs []btcjson.TransactionInput
	inputs = append(inputs, btcjson.TransactionInput{
		Txid: sourceUTXO.TxID,
		Vout: sourceUTXO.Vout,
	})
	//	构造输出
	outAddr, err := btcutil.DecodeAddress(destAddr, &chaincfg.SimNetParams)
	outAddrchange, err := btcutil.DecodeAddress(sourceAddr, &chaincfg.SimNetParams)
	if err != nil {
		return nil, err
	}
	outputs := map[btcutil.Address]btcutil.Amount{
		// 0.1BTC的手续费
		outAddr:       btcutil.Amount(amount),
		outAddrchange: btcutil.Amount(int64(sourceUTXO.Amount)*1e8 - amount - 1e6),
	}
	//	创建交易
	rawTx, err := client.CreateRawTransaction(inputs, outputs, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateRawTransaction error:%s", err)
	}
	return rawTx, nil
}
func GenerateTrans2(client *rpcclient.Client, sourceAddr, destAddr string, amount int64) (*wire.MsgTx, error) {
	// 筛选源地址的UTXO
	utxos, _ := client.ListUnspent()
	var sourceUTXO btcjson.ListUnspentResult
	for i, utxo := range utxos {
		if utxo.Address == sourceAddr {
			if sourceAddr == "SXXfUx9qdszdhEgFJMq5625co9JrqbeRBv" {
				if utxo.Amount == 50 {
					sourceUTXO = utxo
					break
				}
			} else {
				sourceUTXO = utxo
				break
			}

		}
		if i == len(utxos)-1 {
			return nil, fmt.Errorf("UTXO not found")
		}
	}

	// 构造输入
	var inputs []btcjson.TransactionInput
	inputs = append(inputs, btcjson.TransactionInput{
		Txid: sourceUTXO.TxID,
		Vout: sourceUTXO.Vout,
	})
	//	构造输出
	outAddr, err := btcutil.DecodeAddress(destAddr, &chaincfg.SimNetParams)
	outAddrchange, err := btcutil.DecodeAddress(sourceAddr, &chaincfg.SimNetParams)
	if err != nil {
		return nil, err
	}
	outputs := map[btcutil.Address]btcutil.Amount{
		// 0.1BTC的手续费
		outAddr:       btcutil.Amount(amount),
		outAddrchange: btcutil.Amount(int64(sourceUTXO.Amount)*1e8 - amount - 1e6),
	}
	//	创建交易
	rawTx, err := client.CreateRawTransaction(inputs, outputs, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateRawTransaction error:%s", err)
	}
	return rawTx, nil
}

// SignTrans 签名交易，嵌入秘密消息，并保存特殊q
func SignTrans(client *rpcclient.Client, rawTx *wire.MsgTx, embedMsg *[]byte) (*wire.MsgTx, error) {
	signedTx, complete, err, _ := client.SignRawTransaction(rawTx, embedMsg)
	if err != nil {
		return nil, fmt.Errorf("error signing transaction: %v", err)
	}
	if !complete {
		return nil, fmt.Errorf("transaction signing incomplete")
	}

	return signedTx, nil
}

// BroadTrans 广播交易
func BroadTrans(client *rpcclient.Client, signedTx *wire.MsgTx) (*chainhash.Hash, error) {
	txHash, err := client.SendRawTransaction(signedTx, false)
	if err != nil {
		return nil, fmt.Errorf("SendRawTransaction error: %v", err)
	}
	return txHash, nil
}
