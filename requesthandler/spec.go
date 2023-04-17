package requesthandler

import (
	"net/http"
)

type RequestHandler interface {
	HandleRequest(http.ResponseWriter, *http.Request) error
}
