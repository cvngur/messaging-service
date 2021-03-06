package handler

import (
	"cvngur/messaging-service/app/repositories"
	"cvngur/messaging-service/app/services"
	"encoding/json"
	"net/http"
)

var service = services.NewUserService(repositories.NewUserRepository())

func BlockUserHandler(w http.ResponseWriter, r *http.Request) {
	var b Block
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response := Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}
	err = service.BlockUser(b.Username, b.BlockedUser)
	if err != nil {
		return
	}
}
