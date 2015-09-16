//https://github.com/kjk/kjkpub/blob/master/go/github_sample/sample_1.go

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2/google"
	"os"
	"overseaagent/vendor/golang.org/x/oauth2"
)

const (
	CLIENT_SECRET = "SgL8YOaB_MsdVNaV4drQ9Qxg"
	REDIRECT_URL  = "http://127.0.0.1:8080/oauth2/callback"
	CLIENT_ID     = "730583430805-f2a2e2sg0dqp74m5gps1l6el1g6emb49.apps.googleusercontent.com"

	// random string for oauth2 API calls to protect against CSRF
	OAuth2State = "oversea_agent_state"

	// View and manage your Google Play Developer account
	AndroidpublisherScope = "https://www.googleapis.com/auth/androidpublisher"
)

const htmlIndex = `<html><body>
Logged in with <a href="/login">Google</a>
</body></html>
`

var (
	oauth2Conf = &oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectURL:  REDIRECT_URL,
		Scopes:       []string{AndroidpublisherScope},
		Endpoint:     google.Endpoint,
	}

	oauth2Client *http.Client
)

// /index
func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlIndex))
}

// /login
func handleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Conf.AuthCodeURL(OAuth2State, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// /github_oauth_cb. Called by github after authorization is granted
func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {

	state := r.FormValue("state")
	if state != OAuth2State {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", OAuth2State, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	os.Setenv("HTTP_PROXY", "http://192.168.6.72:1080")

	code := r.FormValue("code")
	token, err := oauth2Conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println(token.AccessToken)

	/*
			oauthClient := oauthConf.Client(oauth2.NoContext, token)
		client := github.NewClient(oauthClient)
		user, _, err := client.Users.Get("")
		if err != nil {
			fmt.Printf("client.Users.Get() faled with '%s'\n", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	*/

}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGitHubLogin)
	http.HandleFunc("/oauth2/callback", handleGoogleCallback)
	fmt.Print("Started running on http://127.0.0.1:8080\n")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
