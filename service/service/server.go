package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {

	mx.HandleFunc("/user/login/", getUserLoginHandler(formatter)).Methods("GET")
	mx.HandleFunc("/user/register", psotUserRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/user/verify/", getUserVerifyHandler(formatter)).Methods("GET")
	mx.HandleFunc("/user/logout/", getUserLogoutHandler(formatter)).Methods("GET")


}

