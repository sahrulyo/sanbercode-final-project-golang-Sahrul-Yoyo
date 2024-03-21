package hello

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Tulis pesan "hello world" sebagai respons
	fmt.Fprintf(w, "hello world to the world")
}
