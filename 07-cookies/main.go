package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"

	"github.com/nu7hatch/gouuid"
)

func getCode(data string) string {
	secret := []byte("ourkey")
	h := hmac.New(sha256.New, secret)
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("session-id")
		if err != nil {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-id",
				Value: id.String(),
			}
			http.SetCookie(res, cookie)
		}
		fmt.Println(cookie)
	})
	http.ListenAndServe(":9000", nil)
}
