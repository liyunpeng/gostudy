package iriscookie

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestCookiesBasic(t *testing.T) {
	app := newAppCookie()
	e := httptest.New(t, app, httptest.URL("http://localhost"))

	cookieName, cookieValue := "my_cookie_name", "my_cookie_value"

	// Test Set A Cookie.
	t1 := e.GET(fmt.Sprintf("/cookies/%s/%s", cookieName, cookieValue)).Expect().Status(httptest.StatusOK)
	t1.Cookie(cookieName).Value().Equal(cookieValue) // validate cookie's existence, it should be there now.
	t1.Body().Contains(cookieValue)

	// Test Retrieve A Cookie.
	t2 := e.GET(fmt.Sprintf("/cookies/%s", cookieName)).Expect().Status(httptest.StatusOK)
	t2.Body().Equal(cookieValue)

	/*
		凡是有testing的代码符合一定格式都可以做go test -v测试代码,
	本程序运行结果:
	Administrator@WIN-U9IV8COBU35 MINGW64 /d/gostudy/src/gostudy/src/mystudy/iris/cookie (master)
	$ go test -v
	=== RUN   TestCookiesBasic
	--- PASS: TestCookiesBasic (0.09s)
	PASS
	ok      gostudy/src/mystudy/iris/cookie    3.223s

	*/

	// Test Remove A Cookie.
	//t3 := e.DELETE(fmt.Sprintf("/cookies/%s", cookieName)).Expect().Status(httptest.StatusOK)
	//t3.Body().Contains(cookieName)
	//
	//t4 := e.GET(fmt.Sprintf("/cookies/%s", cookieName)).Expect().Status(httptest.StatusOK)
	//t4.Cookies().Empty()
	//t4.Body().Empty()
}