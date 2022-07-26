package handlers

import (
	"net/http"

	"github.com/Kenny2397/visual-programming/handlers/drawflow"
	"github.com/Kenny2397/visual-programming/server"
)

func ReadAll(s server.Server) http.HandlerFunc {
	return drawflow.GetAllDrawflows(s)
}

func ReadByIdg(s server.Server) http.HandlerFunc {
	return drawflow.GetDrawflowByIdg(s)
}

func CreateDrawflow(s server.Server) http.HandlerFunc {
	return drawflow.InsertDrawflow(s)
}
