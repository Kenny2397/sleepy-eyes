package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Kenny2397/visual-programming/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	// Dg      dgo.Dgraph `json:"dg"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	// dg := database.NewClient()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// status 200 ok!
		w.WriteHeader(http.StatusOK)
		// fmt.Println(dg)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to my server 🚀",
			Status:  true,
			// Dg:      *dg,
		})
	}
}
