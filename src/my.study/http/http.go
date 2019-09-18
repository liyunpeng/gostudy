package httpserver

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
	测试结果：
	http://127.0.0.1:9090/test?b=2&c=3
	r.Method= GET
	form= map[b:[2] c:[3]]
	path= /test
	scheme=
	r.Form["url_long"]= []
	key= b ,  val= 2
	key= c ,  val= 3
*/
func handlerReqeust(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                          //解析url传递的参数，对于POST则解析响应包的主体（request body）
	token := r.Form.Get("token")
	fmt.Fprintln(w, "token=", token)

	if token != "" {
		//验证token的合法性
	} else {
		//不存在token报错
	}

	fmt.Fprintln(w, "r.Method=", r.Method) //获取请求的方法
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Fprintln(w, "form=", r.Form)
	fmt.Fprintln(w, "path=", r.URL.Path)
	fmt.Fprintln(w, "scheme=", r.URL.Scheme)
	fmt.Fprintln(w, "r.Form[\"url_long\"]=", r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Fprintln(w, "key=", k, ",  val=", strings.Join(v, ""))
	}
}


var globalSessions *session.Manager

//然后在init函数中初始化
func inithttp() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法

	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {

		h := md5.New()
		crutime := time.Now().Unix()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")


		//t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")

		t.Execute(w, token)


		//t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
	/*
		if r.Method == "GET" {
			t, _ := template.ParseFiles("login.gtpl")
			log.Println(t.Execute(w, nil))
		} else {
			//请求的是登录数据，那么执行登录的逻辑判断
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
		}
	*/
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")

	timeLayout := "2006-01-02 15:04:05"
	s1 := "countnum=" + strconv.Itoa(sess.Get("countnum").(int)) +  ",  \n"
	//s3 := time.Unix(sess.Get("createtime").(int64), 0).Format(timeLayout)

	s2 := "session createtime="+ time.Unix(sess.Get("createtime").(int64), 0).Format(timeLayout) +", \n "
	t.Execute(w, s1 + s2 )

}

func HttpServer() {
	inithttp()
	http.HandleFunc("/", handlerReqeust) //设置访问的路由
	http.HandleFunc("/login", login)     //设置访问的路由
	http.HandleFunc("/count", count)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
