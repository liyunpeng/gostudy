package encode1

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
)

func simple() {


	s1 := sha1.New()
	io.WriteString(s1, "123456")
	fmt.Printf("123456 sha1=% x  \n", s1.Sum(nil))

	s2 := sha256.New()
	io.WriteString(s2, "123456")
	fmt.Printf("123456 sha245=% x  \n", s2.Sum(nil))


	h := md5.New()
	io.WriteString(h, "123456")
	fmt.Printf("123456 md5=% x  \n", h.Sum(nil))
}

func salt()  {
	//import "crypto/md5"
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

	fmt.Printf("加盐方案(123456) = % x \n", last)

	fmt.Println("加盐方案(123456) = ", last)
}

func script()  {
	str := "123456"
	salt := []byte("@#$%")
	dk, _ := scrypt.Key([]byte(str), salt, 16384, 8, 1, 32)
	fmt.Printf("专家方案(123456) = % x ", dk)
}

