package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		session, _ := store.Get(req, "session")
		if req.FormValue("email") != "" {
			// check password
			session.Values["email"] = req.FormValue("email")
		}
		session.Save(req, res)

		io.WriteString(res, `
            <!DOCTYPE html>
            <html lang="en">
            <head>
                <meta charset="UTF-8">
                <title>Sessions</title>
            </head>
            <body>
                <form method="POST">
                    <div>`+fmt.Sprint(session.Values["email"])+`</div><br/>
                    <input type="email" name="email" placeholder="email">
                    <input type="password" name="password" placeholder="password">
                    <input type="submit" >
                </form>
            </body>
            </html>
            `)
	})
	http.ListenAndServe(":9000", nil)
}
