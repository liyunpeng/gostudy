package encodemain

import (
	"fmt"
)



func Encode() {
	fmt.Println("<---------------------- Encode begin ---------------------->")

	Base64Encode()

	AesEncode()

	HashEncode()

	JsonEncode()

	JsonEasyEncode()

	fmt.Println("<---------------------- Encode end ---------------------->")
}
