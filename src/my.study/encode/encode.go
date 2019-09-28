package encode1

import (
	"fmt"
)



func Encode() {
	fmt.Println("<---------------------- Encode begin ---------------------->")

	Base64Encode()

	AesEncode()

	HashEncode()

	JsonEncode()

	fmt.Println("<---------------------- Encode end ---------------------->")
}
