package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/ppirch/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})

	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedFollowToAPIFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollowsByUserId(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollowByUserID(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseFeedFollowToAPIFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "FeedFollowID")
	FeedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing UUID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     FeedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}

	responseWithJSON(w, 200, nil)
}
