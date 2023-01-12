// package custom

// import (
// 	"net/http"
// )

// func Hello(w http.ResponseWriter, req *http.Request) {
// 	w.Write([]byte("Hello, world!"))
// }

// func Display() {
// 	http.HandleFunc("/", Hello)
// 	http.ListenAndServe(":8090", nil)
// }