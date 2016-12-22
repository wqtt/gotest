package function

import (
	"net/http"
	"testing"
)

func SayHelloHttps(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func TestHttps(t *testing.T) {
	/*
		http.HandleFunc("/", SayHelloHttps)
		err := http.ListenAndServeTLS(":8081", "cert/server.crt", "cert/server.key", nil)
		if err != nil {
			fmt.Printf("err:%s\n", err.Error())
		}
	*/

}
