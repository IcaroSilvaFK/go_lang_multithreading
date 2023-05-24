package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {

	// m := sync.Mutex{}
	// var m = sync.Mutex
	// -race race conditional execution

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		w.Write([]byte(fmt.Sprintf("Hello user number %d\n", number)))
	})

	http.ListenAndServe(":8080", nil)

}
