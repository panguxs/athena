package staticfile

import (
	"net/http"
)

//FileNameSpaceHandler file namespace
type FileNameSpaceHandler struct {
	Name    string
	Hosts   []string
	Handler http.Handler
}

//FileNameSpaceHandlerMap namespace handlers
type FileNameSpaceHandlerMap map[string]*FileNameSpaceHandler

const (
	requiredUser     = "user"
	requiredPassword = "pwd"
)

// Implement the ServeHTTP method on our new type
func (hs FileNameSpaceHandlerMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Check if a http.Handler is registered for the given host.
	// If yes, use it to handle the request.
	if handler := hs[r.Host]; handler != nil {

		//verify token
		//basic auth
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			handler.Handler.ServeHTTP(w, r)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
		// handler.Handler.ServeHTTP(w, r)
	} else {
		// Handle host names for which no handler is registered
		http.Error(w, "Forbidden", 403) // Or Redirect?
	}

}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
