package repositorios

import (
	"database/sql"
	"fmt"

	"github.com/antony-raul/DevBook/src/model"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuario cria um repositorio de usuario
func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario model.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare("insert into usuarios (nome,nick,email,senha) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIDIserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDIserido), nil
}

// Buscar traz todos os usuarios que atendem um filtro de nome ouj nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]model.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, err := repositorio.db.Query(
		"select id,nome,nick,email,criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []model.Usuario

	for linhas.Next() {
		var usuario model.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorID(ID uint64) (model.Usuario, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if err != nil {
		return model.Usuario{}, err
	}
	defer linhas.Close()

	var usuario model.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return model.Usuario{}, err
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuario no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario model.Usuario) error {
	statement, err := repositorio.db.Prepare("update usuarios set nome=?, nick=?, email=? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

// Deletar exclui as informaçôes no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// BuscarPorEmail busca um usuario por e email e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (model.Usuario, error) {
	linhas, err := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return model.Usuario{}, err
	}
	defer linhas.Close()

	var usuario model.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Senha,
		); err != nil {
			return model.Usuario{}, err
		}

	}

	return usuario, nil

}
