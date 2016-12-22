package function

import (
	"fmt"
	"net/http"
	"testing"
)

func SayHelloHttp(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func TestHttp(t *testing.T) {

	http.HandleFunc("/", SayHelloHttp)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}

}
