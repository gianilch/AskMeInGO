package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AskMeInGO/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func sendJSON(w http.ResponseWriter, rawData any) {
	data, _ := json.Marshal(rawData)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func (h apiHandler) findRoomId(rawRoomId string, w http.ResponseWriter, r *http.Request) pgstore.Room {
	roomId, err := uuid.Parse(rawRoomId)
	if err != nil {
		http.Error(w, "invalid room id", http.StatusBadRequest)
		room := pgstore.Room{}
		return room
	}

	room, err := h.q.GetRoom(r.Context(), roomId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "room not found", http.StatusBadRequest)
			room := pgstore.Room{}
			return room
		}
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		room := pgstore.Room{}
		return room
	}

	return room
}
