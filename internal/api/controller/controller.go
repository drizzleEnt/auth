package controller

import (
	"encoding/json"
	"net/http"

	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/service"
)

type LoginController struct {
	s service.LoginService
}

func NewController(s service.LoginService) *LoginController {
	return &LoginController{
		s: s,
	}
}

func (c *LoginController) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]interface{}, 0)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&vars)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	refToken, ok := vars["refresh_token"]
	if !ok {
		RespondWithError(w, http.StatusBadRequest, "field refresh not found")
		return
	}

	accessToken, err := c.s.GetAccessToken(r.Context(), refToken.(string))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, accessToken)
}

func (c *LoginController) GetRefreshToken(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]interface{}, 0)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&vars)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	accessTokenReq, ok := vars["refresh_token"]
	if !ok {
		RespondWithError(w, http.StatusBadRequest, "field refresh not found")
		return
	}

	accessTokenResp, err := c.s.GetRefreshToken(r.Context(), accessTokenReq.(string))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, accessTokenResp)
}

func (c *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]interface{}, 0)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&vars)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	username, ok := vars["username"]
	if !ok {
		RespondWithError(w, http.StatusBadRequest, "field username not found")
		return
	}
	pswd, ok := vars["password"]
	if !ok {
		RespondWithError(w, http.StatusBadRequest, "field password not found")
		return
	}

	refToken, err := c.s.Login(r.Context(), &model.UserClaims{
		Username: username.(string),
		Password: pswd.(string),
	})

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{"refresh_token": refToken, "success": true})
}
