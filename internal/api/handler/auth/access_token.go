package auth

import (
	"fmt"
	"github.com/araailym/expense-tracker/internal/auth"
	dbAuth "github.com/araailym/expense-tracker/internal/db/auth"
	"github.com/araailym/expense-tracker/pkg/httputils/request"
	"github.com/araailym/expense-tracker/pkg/httputils/response"
	"net/http"
	"os"
)

type AccessTokenRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
type AccessTokenRequest struct {
	Data *AccessTokenRefreshTokenRequest `json:"data"`
}
type AccessTokenResponse struct {
	Data *AuthTokenPair `json:"data"`
}

func (h *Auth) AccessToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "AccessToken")
	// request parse
	requestBody := &AccessTokenRequest{}
	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}
	userData, err := auth.ParseToken(requestBody.Data.RefreshToken, os.Getenv("TOKEN_SECRET"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail authentication",
			"error", err,
		)
		http.Error(
			w,
			http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized,
		)
		return
	}
	// db request
	dbResp, err := h.db.AccessToken(ctx, &dbAuth.AccessTokenInput{UserID: userData.ID})
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, "failed to query from db", http.StatusInternalServerError)
		return
	}
	if dbResp == nil {
		log.ErrorContext(
			ctx,
			"row is empty",
		)
		//send critical ктото подделывает джвт или удаленный айди с жействующим рефреш токеном
		http.Error(w, "row is empty", http.StatusInternalServerError)
		return
	}
	tokenPair, err := auth.GenerateTokenPair(
		&auth.UserData{
			ID:    fmt.Sprint(dbResp.ID),
			Email: dbResp.Email,
			Role:  dbResp.Role,
		},
		os.Getenv("TOKEN_SECRET"),
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid request %w", err), http.StatusBadRequest)
		return
	}
	// response
	respBody := &AccessTokenResponse{
		Data: &AuthTokenPair{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
	}
	if err := response.JSON(
		w,
		http.StatusOK,
		respBody,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}
	log.InfoContext(
		ctx,
		"success generate access token",
	)
	return
}
