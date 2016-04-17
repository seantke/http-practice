package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		if src, hdr, err := req.FormFile(key); src != nil {
			panicError(err)
			defer src.Close()

			dst, err := os.Create(filepath.Join("./tmp/", hdr.Filename))
			panicError(err)
			defer dst.Close()

			io.Copy(dst, src)
		}

		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res,
			`<form method="POST" enctype="multipart/form-data">
                <input type="file" name="q">
                <input type="submit">
            </form>
			`)
	})
	http.ListenAndServe(":9000", nil)
}
