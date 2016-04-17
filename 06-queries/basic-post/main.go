package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		val := req.FormValue(key)
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res,
			`<form method="POST">
                <input type="text" name="q">
                <input type="submit">
            </form>
			<div>`+val+`</div`)
	})
	http.ListenAndServe(":9000", nil)
}
