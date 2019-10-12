package encode1

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
)

/*
	本文件提供签名算法示例， 基本签名算法有MD5, SHA1, SHA256. SHA1据说已经破解
	以一个60M的文件为测试样本，经过1000次的测试平均值，三种算法的表现为：
	MD5算法运行1000次的平均时间为：226ms
	SHA1算法运行1000次的平均时间为：308ms
	SHA256算法运行1000次的平均时间为：473ms
	安全性方面，显然SHA256（又称SHA2）的安全性最高，但是耗时要比其他两种多很多。
	MD5相对较容易碰撞，因此，SHA1应该是这三种中性能最好的一款加密算法。
 */
func base() {
	s1 := sha1.New()
	io.WriteString(s1, "123456")
	fmt.Printf("被加密的字符串为123456 sha1=% x  \n", s1.Sum(nil))

	s2 := sha256.New()
	io.WriteString(s2, "123456")
	fmt.Printf("被加密的字符串为123456 sha245=% x  \n", s2.Sum(nil))

	h := md5.New()
	io.WriteString(h, "123456")
	fmt.Printf("被加密的字符串为123456 md5=% x  \n", h.Sum(nil))
}

/*
	加盐算法Salt是什么？盐就是一个随机生成的字符串。
	我们将盐与原始密码连接（concat）在一起（放在前面或后面都可以），
	然后将concat后的字符串加密。采用这种方式加密密码，
	查表法就不灵了
 */
func salt()  {
	//假设用户名abc，密码123456
	h := md5.New()
	io.WriteString(h, "123456")   //"需要加密的密码"

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 :=fmt.Sprintf("%x", h.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last :=fmt.Sprintf("%x", h.Sum(nil))

	fmt.Printf("被加密的字符串为123456 加盐方案 = % x \n", last)
}

/*
	专家方案
 */
func script()  {
	str := "123456"
	salt := []byte("@#$%")
	dk, _ := scrypt.Key([]byte(str), salt, 16384, 8, 1, 32)
	fmt.Printf("被加密的字符串为123456 专家方案 = % x \n ", dk)
}

func HashEncode()  {
	base()
	salt()
	script()
}
