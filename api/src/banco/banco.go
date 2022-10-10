package banco

import (
	"database/sql"

	"github.com/antony-raul/DevBook/src/config"
	_ "github.com/go-sql-driver/mysql" //Driver
)

// Conectar abre a conexão com banco de daods e a retorna
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
