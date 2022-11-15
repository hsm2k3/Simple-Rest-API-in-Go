package Routes

import (
	"github.com/hsm2k3/Simple-Rest-API-in-Go/App/Http/Controllers"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", Controllers.BaseController)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}
