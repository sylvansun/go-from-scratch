package advanced

import (
	"fmt"
	"net/http"
)

func SimpleHello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello from localhost\n")
}

func Headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func HttpServers() {

	http.HandleFunc("/hello", SimpleHello)
	http.HandleFunc("/headers", Headers)

	http.ListenAndServe(":8090", nil)
}
