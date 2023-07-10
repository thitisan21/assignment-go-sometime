package app

import (
	"fmt"
	"net/http"
)

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}
