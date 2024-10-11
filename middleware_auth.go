package main

import (
	"fmt"
	"net/http"

	"github.com/ppirch/rssagg/internal/auth"
	"github.com/ppirch/rssagg/internal/database"
)

type authenHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) authMiddleware(next authenHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r)
		if err != nil {
			responseWithError(w, 401, fmt.Sprintf("Authentication error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 404, fmt.Sprintf("User not found: %v", err))
			return
		}

		next(w, r, user)
	}
}
