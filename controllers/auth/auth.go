package auth

import (
	"encoding/json"
	"net/http"

	"github.com/jaysyanshar/godate-rest/models/restmodel"
	"github.com/jaysyanshar/godate-rest/services/auth"
)

type AuthController interface {
	SignUpHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	svc auth.AuthService
}

func NewController(svc auth.AuthService) AuthController {
	return &controller{
		svc: svc,
	}
}

func (c *controller) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req restmodel.SignUpRequest
	var res restmodel.SignUpResponse

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res = restmodel.SignUpResponse{
			Success: false,
			Message: err.Error(),
		}
		http.Error(w, res.Message, http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	res, err = c.svc.SignUp(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (c *controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req restmodel.LoginRequest
	var res restmodel.LoginResponse

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res = restmodel.LoginResponse{
			Success: false,
			Message: err.Error(),
		}
		http.Error(w, res.Message, http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	res, err = c.svc.Login(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(res)
}
