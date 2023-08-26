package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"net/mail"

	"github.com/Bhimmo/golang-simple-api/pkg/auth"
)

type BodyInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
type BodyOutput struct {
	AccessToken string `json:"access-token"`
}

func AccessToken(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.Write([]byte("Erro body"))
		return
	}

	var bodyInput BodyInput
	var BodyOutput BodyOutput
	err = json.Unmarshal(body, &bodyInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body error"))
		return
	}

	if bodyInput.Username == "" || bodyInput.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("credential invalid: username and email"))
		return
	}
	_, err = mail.ParseAddress(bodyInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("credential invalid: email invalid"))
		return
	}

	token := auth.GenerateJWT(bodyInput.Username, bodyInput.Email)
	BodyOutput.AccessToken = token

	json.NewEncoder(w).Encode(&BodyOutput)
}
