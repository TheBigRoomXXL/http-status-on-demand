package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var validStatusCodes = []int{
	100, 101, 102, 103,
	200, 201, 202, 203, 204, 205, 206, 207, 208, 226,
	300, 301, 302, 303, 304, 305, 307, 308,
	400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 425, 426, 428, 429, 431, 451,
	500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511,
}

type MyHandler struct {
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, " - ", req.URL.Path, " - ", req.RemoteAddr, " - ", req.UserAgent())
	path := req.URL.Path
	status, err := strconv.Atoi(path[1:])
	if err != nil {
		http.Error(w, "Invalid status code as input", http.StatusBadRequest)
		return
	}

	for _, s := range validStatusCodes {
		if s == status {
			w.WriteHeader(status)
			return
		}
	}
	http.Error(w, "Invalid status code as input", http.StatusBadRequest)
}

func main() {

	listen := ":8000"
	argsWithProg := os.Args
	if len(argsWithProg) > 1 {
		listen = argsWithProg[1]
	}

	fmt.Println("Starting server...")
	http.ListenAndServe(listen, &MyHandler{})
	fmt.Println("Stoping server.")

}
