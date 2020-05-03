package auth

import "net/http"

//a middleware returns a func that works like a "filter"
//if the func does not call next.ServeHTTP the chain will be interrupted
//this filter is applied at every incoming HTTP request
func CheckUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...

		//check http headers for auth (example)
		if r.Header.Get("X-AUTH-USERNAME") != "antares" {
			//this will interrupt the chain and will not pass execution to our handler in API code
			http.Error(w, "Missing auth header X-AUTH-USERNAME", http.StatusUnauthorized)
			return
		}

		//ok proceed to next middleware or final handler code
		next.ServeHTTP(w, r)
	})
}
