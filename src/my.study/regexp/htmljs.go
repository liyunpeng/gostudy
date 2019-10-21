package regexp1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var str  = `
<!DOCTYPE html>
<html>
   <head>
		<body> 
			网页正文
			<script src="jquery/jquery-3.3.1.min.js">
				$('test2').onclick = function(){console.log('test2 : click!');}
			</script>
 		</body>
   </body>
</html>
`
//去除所有尖括号内的HTML代码
func rmHtml()  {
	//pattern := "\\<[\\S\\s]+?\\>"
	pattern := `<[\S\s]+?>`
	re, _ := regexp.Compile(pattern)
	str1 := re.ReplaceAllString(str, "")
	fmt.Println("rmHtml pattern=", pattern, "str =", str1)
	/*
	运行结果:
	rmHtml pattern= <[\S\s]+?> str =




				网页正文

					$('test2').onclick = function(){console.log('test2 : click!');}
	 */
}

func rmScript() {

	pattenBackqute := `<script[\S\s]+?</script>`
	re, _ := regexp.Compile(pattenBackqute)
	str1 := re.ReplaceAllString(str, "")
	fmt.Println("pattenBackqute=", pattenBackqute, "  str=", str1)
	/*
	运行结果:
	pattenBackqute= <script[\S\s]+?</script>   str=
	<!DOCTYPE html>
	<html>
	   <head>
			<body>
				网页正文

	 		</body>
	   </body>
	</html>
	 */

	pattenDoublequte := "<script[\\S\\s]+?</script>"
	re, _ = regexp.Compile(pattenBackqute)
	str2 := re.ReplaceAllString(str, "")
	fmt.Println("pattenDoublequte=", pattenDoublequte, " str=", str2)
	/*
	运行结果:
	pattenDoublequte= <script[\S\s]+?</script>  str=
	<!DOCTYPE html>
	<html>
	   <head>
			<body>
				网页正文

	 		</body>
	   </body>
	</html>

	*/


	pattenDoublequte1 := "\\<script[\\S\\s]+?\\</script\\>"
	re, _ = regexp.Compile(pattenDoublequte1)
	str3 := re.ReplaceAllString(str, "")
	fmt.Println("pattenDoublequte1=", pattenDoublequte1, " str=", str3)
	/*
	运行结果:
	pattenDoublequte1= \<script[\S\s]+?\</script\>  str=
	<!DOCTYPE html>
	<html>
	   <head>
			<body>
				网页正文

	 		</body>
	   </body>
	</html>
	 */
}

func crawl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	//去除STYLE
	re, _ := regexp.Compile("<style[\\S\\s]+?</style>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("<script[\\S\\s]+?</script>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("<[\\S\\s]+?>")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println("crawl str=", strings.TrimSpace(src))
}

func HtmlJsMain()  {
	rmScript()
	rmHtml()
	/*
		crawl	比较耗时, 调试时注释
	 */
	//crawl("http://www.baidu.com")
}