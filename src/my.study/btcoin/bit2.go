package btcoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"golang.org/x/crypto/ripemd160"
)

func GenerateBTC() (string, string, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}

	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, false)
	if err != nil {
		return "", "", err
	}
	pubKeySerial := privKey.PubKey().SerializeUncompressed()

	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", err
	}

	return privKeyWif.String(), pubKeyAddress.EncodeAddress(), nil
}

func GenerateBTCTest() (string, string, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}

	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.TestNet3Params, false)
	if err != nil {
		return "", "", err
	}
	pubKeySerial := privKey.PubKey().SerializeUncompressed()

	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.TestNet3Params)
	if err != nil {
		return "", "", err
	}

	return privKeyWif.String(), pubKeyAddress.EncodeAddress(), nil
}

func Test2() {
	wifKey, address, _ := GenerateBTCTest() // 测试地址
	// wifKey, address, _ := GenerateBTC() // 正式地址
	fmt.Println(address, wifKey)
}

// 256加密算法  返回256位比特位，即64个字节
func TestSha256() {
	hasher := sha256.New()
	hasher.Write([]byte("The quick brown fox jumps over the lazy dog"))
	hashBytes := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	fmt.Println(hashString)
}

// ripemd160 加密算法 ， 返回160位比特位， 即40个字节
func TestRipe() {
	hasher := ripemd160.New()
	hasher.Write([]byte("The quick brown fox jumps over the lazy dog"))
	hashBytes := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	fmt.Println(hashString)
}

/*
第五步: 输出交易原始信息, 广播到网络上
// 4. 输出Hex
buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
if err := tx.Serialize(buf); err != nil {
}
txHex := hex.EncodeToString(buf.Bytes())
fmt.Println("hex", txHex)
将输出的hex广播到网络上, https://tbtc.blockdozer.com/insight/tx/send
下面给出完整源码:
*/
func Transaction() {
	address := "mt4p3rZpJE5fXEqvGzNBk9HxYXcWKpPJSd"
	var balance int64 = 65000000 // 余额
	var fee int64 = 0.001 * 1e8  // 交易费
	var leftToMe = balance - fee // 余额-交易费就是剩下再给我的

	// 1. 构造输出
	outputs := []*wire.TxOut{}

	// 1.1 输出1, 给自己转剩下的钱
	addr, _ := btcutil.DecodeAddress(address, &chaincfg.SimNetParams)
	pkScript, _ := txscript.PayToAddrScript(addr)
	outputs = append(outputs, wire.NewTxOut(leftToMe, pkScript))

	// 1.2 输出2, 添加文字
	comment := "这是一个留言, 哈哈"
	pkScript, _ = txscript.NullDataScript([]byte(comment))
	outputs = append(outputs, wire.NewTxOut(int64(0), pkScript))

	// 2. 构造输入
	prevTxHash := "48eea09764713f3dadcfed29490ab5e288299e01e571e1f7a1396a75ce38e067"
	prevPkScriptHex := "76a91489a7f0117eaf47d8b4af740c66116e35ffe1bea988ac"
	prevTxOutputN := uint32(0)

	hash, _ := chainhash.NewHashFromStr(prevTxHash)   // tx hash
	outPoint := wire.NewOutPoint(hash, prevTxOutputN) // 第几个输出
	txIn := wire.NewTxIn(outPoint, nil, nil)
	inputs := []*wire.TxIn{txIn}

	prevPkScript, _ := hex.DecodeString(prevPkScriptHex)
	prevPkScripts := make([][]byte, 1)
	prevPkScripts[0] = prevPkScript

	tx := &wire.MsgTx{
		Version:  wire.TxVersion,
		TxIn:     inputs,
		TxOut:    outputs,
		LockTime: 0,
	}

	// 3. 签名
	privKey := "cV4HmdzGF3gG7NdEtVV7sjq22yoBmZBe5MEGKUqvQTXXXXX" // 私钥
	sign(tx, privKey, prevPkScripts)

	// 4. 输出Hex
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	if err := tx.Serialize(buf); err != nil {
	}
	txHex := hex.EncodeToString(buf.Bytes())
	fmt.Println("hex", txHex)
}

// 签名
func sign(tx *wire.MsgTx, privKeyStr string, prevPkScripts [][]byte) {
	inputs := tx.TxIn
	wif, err := btcutil.DecodeWIF(privKeyStr)

	fmt.Println("wif err", err)
	privKey := wif.PrivKey

	for i := range inputs {
		pkScript := prevPkScripts[i]
		var script []byte
		script, err = txscript.SignatureScript(tx, i, pkScript, txscript.SigHashAll,
			privKey, false)
		inputs[i].SignatureScript = script
	}
}