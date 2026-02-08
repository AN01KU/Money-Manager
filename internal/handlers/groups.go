package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/AN01KU/money-manager/internal/api"
	"github.com/AN01KU/money-manager/internal/middleware"
)

func (h *Handlers) createGroups(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.GetUser(r.Context())
	if !ok {
		api.InternalErrorHandler(w)
		return
	}

	var params api.CreateGroupParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	// validate params
	groupName := params.Name
	createdBy := user.Id

	if params.Name == "" {
		api.RequestErrorHandler(w, errors.New("group name is required"))
		return
	}

	database := h.DB
	group := database.CreateGroup(groupName, createdBy)
	if group == nil {
		api.InternalErrorHandler(w)
		return
	}

	response := api.GroupResponse{
		ID:        group.Id.String(),
		Name:      group.Name,
		CreatedBy: group.CreatedBy.String(),
		CreatedAt: group.CreatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
