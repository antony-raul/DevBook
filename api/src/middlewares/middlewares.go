package middlewares

import (
	"log"
	"net/http"

	"github.com/antony-raul/DevBook/src/autenticacao"
	"github.com/antony-raul/DevBook/src/respostas"
)

func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// autenticar verifica se o usuario esta autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := autenticacao.ValidarToken(r); err != nil {
			respostas.Erro(w, http.StatusUnauthorized, err)
			return
		}

		proximaFuncao(w, r)
	}
}
