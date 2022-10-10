package rotas

import (
	"net/http"

	"github.com/antony-raul/DevBook/src/controllers"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
