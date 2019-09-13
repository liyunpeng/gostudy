package auto

import (
	"fmt"
	"os"
)

func Genfile(){
	fmt.Println("<--------------- genfile -------------------> ")
	_ = os.MkdirAll("./dia1/auto/gen", os.ModePerm)
	
	f, _ := os.Create("./dia1/auto/gen/gen1.go")
	
//	f.Write([]byte("package gen\n"))

	//var s []byte
	
	const s = `
package gen  

func Testgen(){ 

	fmt.Println(1) 
	
|


`


	f.WriteString(s)

	/*

	f.Seek(0, os.SEEK_SET)

	p := make([]byte, 5)

	if _, err := f.Read(p); err != nil {
		log.Fatal("[File]", err)
	}
*/


	f.Close()	
}
