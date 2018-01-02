package spotify

import "fmt"

import "scionofbytes.me/shelper/auth"

func Authenticate() {

	authenticator := newAuthenticator()
	token = auth.Authenticate(authenticator)
	fmt.Println("Token: " + token)
}
