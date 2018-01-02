package auth

import "fmt"
import "net/http"
import "github.com/skratchdot/open-golang/open"

type Authenticator interface {
	GetName() string
	GetURL() string
	GetToken(request *http.Request) string
}

var tokenChan chan string

func Authenticate(authenticator Authenticator) string {

	open.Run(authenticator.GetURL())

	http.HandleFunc("/auth/"+authenticator.GetName()+"/callback", generateAuthCompletionFunction(authenticator))
	go http.ListenAndServe(":6666", nil)

	token := <-tokenChan

	return token
}

func generateAuthCompletionFunction(authenticator Authenticator) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		token := authenticator.GetToken(r)

		fmt.Fprintf(w, "Login Completed!")
		tokenChan <- token
	}
}
