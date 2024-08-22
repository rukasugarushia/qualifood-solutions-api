/*package main

import (
	"log"
	"qualifood-solutions-api/internal/infrastructure"
)

func main() {
	router := infrastructure.SetupRouter()
	log.Println("Server is running at http://localhost:8080")
	router.Run(":8080")
}*/

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

// Initialize the Google OAuth configuration
func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "130205075627-ejesblbq90hen8t0vej75k5e7vp7nv2a.apps.googleusercontent.com", /*os.Getenv("GOOGLE_CLIENT_ID")*/     // Set this via environment variable
		ClientSecret: "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",                                      /*os.Getenv("GOOGLE_CLIENT_SECRET")*/ // Set this via environment variable
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `<html><body><a href="/login">Google Login</a></body></html>`
	fmt.Fprint(w, html)
}

// Redirect to Google's OAuth 2.0 consent page
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("randomStateToken")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Handle callback from Google after user consents
func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != "randomStateToken" {
		fmt.Println("State is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Could not get token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?alt=json")
	if err != nil {
		fmt.Println("Error getting user info:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer userInfo.Body.Close()
	content, err := io.ReadAll(userInfo.Body)
	if err != nil {
		fmt.Println("Error reading user info:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "User Info: %s", content)
}
