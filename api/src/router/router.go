package router

import (
	"github.com/antony-raul/DevBook/src/router/rotas"
	"github.com/gorilla/mux"
)

// Gerar vai retorna um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
