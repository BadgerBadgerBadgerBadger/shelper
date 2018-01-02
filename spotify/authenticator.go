package spotify

import "log"
import "net/http"
import "github.com/zmb3/spotify"

import "scionofbytes.me/shelper/util"
import "scionofbytes.me/shelper/auth"

const name = "spotify"

type spotifyAuthenticator struct {
	name  string
	state string
	auth  spotify.Authenticator
}

func (sa spotifyAuthenticator) GetName() string {
	return sa.name
}

func (sa spotifyAuthenticator) GetURL() string {
	return "http://localhost:6666/auth/" + name
}

func (sa spotifyAuthenticator) GetToken(request *http.Request) string {

	token, err := sa.auth.Token(sa.state, request)

	if err != nil {
		log.Fatal(err)
	}

	return token.AccessToken
}

func newAuthenticator() auth.Authenticator {

	callbackURL := "http://localhost:6666/auth/" + name + "/callback"

	return spotifyAuthenticator{
		name:  "spotify",
		state: util.RandStringBytesRmndr(12),
		auth:  spotify.NewAuthenticator(callbackURL, spotify.ScopeUserReadPrivate),
	}
}
