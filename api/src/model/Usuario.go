package model

import (
	"errors"
	"strings"
	"time"

	"github.com/antony-raul/DevBook/src/seguranca"
	"github.com/badoux/checkmail"
)

// Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	if err := usuario.formatar(etapa); err != nil {
		return nil
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome e obrigatorio não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("o email e obrigatorio não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("o email inserido é invalido")
	}

	if usuario.Nick == "" {
		return errors.New("o nick e obrigatorio não pode estar em branco")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("a senha e obrigatorio não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, err := seguranca.Hash(usuario.Senha)
		if err != nil {
			return err
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
