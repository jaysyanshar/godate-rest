// this file is to test whether the authentication is working or not
package dashboard

import "net/http"

type DashboardController interface {
	HelloHandler(w http.ResponseWriter, r *http.Request)
}

type controller struct {
}

func NewController() DashboardController {
	return &controller{}
}

func (c *controller) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
