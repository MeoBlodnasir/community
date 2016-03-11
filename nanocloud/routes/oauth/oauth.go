package oauth

import (
	"net/http"

	"github.com/Nanocloud/community/nanocloud/oauth2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	oauth2.HandleRequest(w, r)
}
