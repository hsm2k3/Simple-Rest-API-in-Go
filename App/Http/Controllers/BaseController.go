package Controllers

import (
	"fmt"
	"net/http"
)

// BaseController is the base controller for the application
func BaseController(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		return
	}
}
