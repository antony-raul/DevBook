package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antony-raul/DevBook/src/config"
	"github.com/antony-raul/DevBook/src/router"
)

func main() {
	config.Carregar()

	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
