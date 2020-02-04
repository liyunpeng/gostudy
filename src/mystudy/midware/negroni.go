package midware1

import (
	"github.com/urfave/negroni"
	"net/http"
)

func Negronimain() {

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewStatic(http.Dir("/tmp")))
	n.Run(":8080")
}
