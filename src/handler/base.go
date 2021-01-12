package handler

import (
	"net/http"
)

// BaseHandler here's where the incoming read request will be received
func (handler *handler) BaseHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hello := "hello, world"
		w.Write([]byte(hello))
	}
}

//    // decode request
//		json.NewDecoder(r.Body).Decode(&request)
//		ioutil.ReadAll(r.Body)
//		defer r.Body.Close()
//
//		// Return the response
//		// set the header
//		w.Header().Set("Content-Type", "application/json")
//
//    // encode response
//		json.NewEncoder(w).Encode(response)
