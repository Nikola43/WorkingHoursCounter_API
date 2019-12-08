package controllers

import (
	"github.com/nikola43/WorkingHoursCounterApi/models"
	"github.com/nikola43/WorkingHoursCounterApi/utils"
	"net/http"
	"time"
)

type UserResult struct {
	User chan models.User `json:"user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		_ = r.Body.Close()
	}()

	userResult := UserResult{
		User: make(chan models.User),
	}

	err := myPool.Submit(func() {
		user := models.User{FingerPrint: "hola"}
		userResult.User <- user
		time.Sleep(time.Duration(10) * time.Millisecond)
	})

	utils.HandleError(err)
	utils.RespondWithJSON(w, http.StatusOK, <-userResult.User)
}
