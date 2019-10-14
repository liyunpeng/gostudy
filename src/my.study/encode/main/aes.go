package encodemain

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

/*
	高级加密标准(AES,Advanced Encryption Standard)为最常见的对称加密算法
	(微信小程序加密传输就是用这个加密算法的)。
	对称加密算法也就是加密和解密用相同的密钥，
	加密和解密用到的密钥是相同的，这种加密方式加密速度非常快，
	适合经常发送数据的场合。缺点是密钥的传输比较麻烦。
	1:最安全--AES,
	2:最方便--MD5---但是无法反向,
	3:最麻烦--AES,
*/
func aesDemo() {
	/*
		设置秘钥
		在对称加密算法中，加密与解密的密钥是相同的。
		密钥为接收方与发送方协商产生，但不可以直接在网络上传输，否则会导致密钥泄漏，
		通常是通过非对称加密算法加密密钥，然后再通过网络传输给对方，或者直接面对面商量密钥。
		密钥是绝对不可以泄漏的，否则会被攻击者还原密文，窃取机密数据。
		这个秘钥key必须是16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法
	 */
	key := []byte("1234567890123456")

	//明文
	origData := []byte("hello world")

	//加密
	en := AESEncrypt(origData, key)

	//解密
	de := AESDecrypt(en, key)
	fmt.Println(string(de))
}

//加密
func AESEncrypt(origData, key []byte) []byte {
	//获取block块
	block, _ := aes.NewCipher(key)

	//补码
	origData = PKCS7Padding(origData, block.BlockSize())

	//加密模式，
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])

	//创建明文长度的数组
	crypted := make([]byte, len(origData))

	//加密明文
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}

//补码
func PKCS7Padding(origData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padtext...)
}

//解密
func AESDecrypt(crypted, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData
}

//去补码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}

func AesEncode() {
	aesDemo()
}
