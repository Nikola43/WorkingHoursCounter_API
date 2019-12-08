package controllers

import (
	"github.com/nikola43/WorkingHoursCounterApi/models"
	"github.com/nikola43/WorkingHoursCounterApi/utils"
	"net/http"
)

type UserResult struct {
	User chan models.User `json:"user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	userResult := UserResult{
		User: make(chan models.User),
	}

	err := myPool.Submit(func() {
		user := models.User{FingerPrint: "hola"}
		userResult.User <- user
	})
	utils.HandleError(err)
	utils.RespondWithJSON(w, http.StatusOK, <-userResult.User)
}
