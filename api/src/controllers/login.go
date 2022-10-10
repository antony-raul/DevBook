package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/antony-raul/DevBook/src/autenticacao"
	"github.com/antony-raul/DevBook/src/banco"
	"github.com/antony-raul/DevBook/src/model"
	"github.com/antony-raul/DevBook/src/repositorios"
	"github.com/antony-raul/DevBook/src/respostas"
	"github.com/antony-raul/DevBook/src/seguranca"
)

// Login responsavel por autenticar usuario na api
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisiacao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
	}

	var usuario model.Usuario
	if err := json.Unmarshal(corpoRequisiacao, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuarioSalvoNoBanco, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
