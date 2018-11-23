package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gomodule/oauth1/oauth"
	"log"
	"net/http"
	"os"
)

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "http://www.fatsecret.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "http://www.fatsecret.com/oauth/authorize",
	TokenRequestURI:               "http://www.fatsecret.com/oauth/access_token",
	SignatureMethod:               oauth.HMACSHA1,
	Credentials: oauth.Credentials{
		Token: os.Getenv("FATSECRET_API_CONSUMER_KEY"),
		Secret: os.Getenv("FATSECRET_API_CONSUMER_SECRET"),
	},
	TemporaryCredentialsMethod: "GET",
	TokenCredentailsMethod: "GET",
}

// serveOAuthCallback handles callbacks from the OAuth server.
func callback(w http.ResponseWriter, r *http.Request) {

	/*s := session.Get(r)
	tempCred, _ := s[tempCredKey].(*oauth.Credentials)
	if tempCred == nil || tempCred.Token != r.FormValue("oauth_token") {
		http.Error(w, "Unknown oauth_token.", 500)
		return
	}
	tokenCred, _, err := oauthClient.RequestToken(nil, tempCred, r.FormValue("oauth_verifier"))
	if err != nil {
		http.Error(w, "Error getting request token, "+err.Error(), 500)
		return
	}
	delete(s, tempCredKey)
	s[tokenCredKey] = tokenCred
	if err := session.Save(w, r, s); err != nil {
		http.Error(w, "Error saving session , "+err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/", 302)*/
	panic(errors.New("callback"))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	callbackURL := "http://localhost:9080/callback"
	creds, err := oauthClient.RequestTemporaryCredentials(nil, callbackURL, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(creds)
}

var httpAddr = flag.String("addr", ":9080", "HTTP server address")

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/callback", callback)

	if err := http.ListenAndServe(*httpAddr, nil); err != nil {
		log.Fatalf("Error listening, %v", err)
	}
}
