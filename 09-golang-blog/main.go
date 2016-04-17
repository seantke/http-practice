package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

// HandleRoute is route handle
func HandleRoute(res http.ResponseWriter, req *http.Request) {
	// Step 1: Parse template
	tpl, err := template.ParseFiles("assets/templates/index.gohtml")
	if err != nil {
		fmt.Println(err)
		http.Error(res, err.Error(), 500)
		return
	}

	type MyModel struct {
		SomeField int
		Values    []int
		Data      string
	}

	// Step 2: Execute template
	tpl.ExecuteTemplate(res, "assets/templates/index.gohtml", MyModel{
		SomeField: 123,
		Values:    []int{1, 2, 3, 4, 5},
		Data:      "Some value",
	})
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func loginPage(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")
		if email == "whatever" && password == "some-password" {
			session.Values["logged_in"] = "YES"
		} else {
			http.Error(res, "invalid credentials", 401)
			return
		}
		// save session
		session.Save(req, res)
		// redirect to main page
		http.Redirect(res, req, "/", 302)
		return
	}

	// RENDER LOGIN FORM
}

func logoutPage(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	delete(session.Values, "logged_in")
	// save session
	session.Save(req, res)
	// redirect to main page
	http.Redirect(res, req, "/", 302)
	return
}

func main() {
	// http.HandleFunc("/session", HandleSessionRoute)
	http.HandleFunc("/", HandleRoute)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/logout", logoutPage)

	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}

// // HandleSessionRoute is session route
// func HandleSessionRoute(res http.ResponseWriter, req *http.Request) {
// 	session, _ := store.Get(req, "session")
// 	// Set session.Values
// 	session.Values["whatever"] = 42
// 	// Get session.Values; is an interface{}
// 	value := session.Values["whatever"]
// 	// Get session.Values assertion
// 	str, _ := session.Values["whatever"].(string) // _ is err
// 	fmt.Println(value, str)
// 	// Delete session.Values
// 	delete(session.Values, "whatever")
//
// 	// Save the session
// 	session.Save(req, res)
// }
