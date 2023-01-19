package handler

import (
	"net/http"
	"travel-planner/util"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var mySigningKey []byte

func InitRouter(config *util.TokenInfo) http.Handler {
	mySigningKey = []byte(config.Secret)

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	router := mux.NewRouter()

	
	router.Handle("/signup", http.HandlerFunc(signupHandlerer)).Methods("POST")
	router.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST")
	router.Handle("/user/{id}",jwtMiddleware.Handler(http.HandlerFunc(UpdateUserHander))).Methods("POST")
	router.Handle("/getUser/{id}",jwtMiddleware.Handler(http.HandlerFunc(getUserHandler))).Methods("GET")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})

	return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}
