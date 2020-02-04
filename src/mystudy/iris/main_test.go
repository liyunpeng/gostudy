package iris1

import (
	"fmt"
	"github.com/kataras/iris/httptest"
	"gostudy/src/mystudy/iris/cookie"
	"testing"
)

func TestResponseWriterQuicktemplate(t *testing.T) {
	baseRawBody := `
<html>
    <head>
        <title>Quicktemplate integration with Iris</title>
    </head>
    <body>
        <div>
            Header contents here...
        </div>
        <div style="margin:10px;">
    <h1>%s</h1>
    <div>
        %s
    </div>
        </div>
    </body>
    <footer>
        Footer contents here...
    </footer>
</html>
`
	expectedIndexRawBody := fmt.Sprintf(baseRawBody, "Index Page", "This is our index page's body.")
	name := "yourname"
	expectedHelloRawBody := fmt.Sprintf(baseRawBody, "Hello World!", "Hello <b>"+name+"!</b>")
	app := cookie.newApp()
	e := httptest.New(t, app)
	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal(expectedIndexRawBody)
	e.GET("/" + name).Expect().Status(httptest.StatusOK).Body().Equal(expectedHelloRawBody)
}
