package migrations

import (
	"log"

	"projeto-back/utils"

	"github.com/casbin/casbin/v2"
)

var administradorPolicy = [][]string{
	{"ADMINISTRADOR", "/*", "(GET)|(POST)|(PATCH)|(DELETE)|(PUT)|(OPTIONS)"},
}

var clientePolicy = [][]string{
	{"CLIENTE", "/*", "(GET)|(POST)|(PATCH)|(DELETE)|(PUT)|(OPTIONS)"},
}

func InicializaPermissoesAcesso(enforcer *casbin.Enforcer) {
	adicionarPoliticas(administradorPolicy, enforcer)
	adicionarPoliticas(clientePolicy, enforcer)
}

func adicionarPoliticas(politicas [][]string, enforcer *casbin.Enforcer) {
	for _, politica := range politicas {
		if existePolitica := enforcer.HasPolicy(politica); !existePolitica {
			ok, err := enforcer.AddPolicy(politica)
			if err != nil {
				utils.ErroLog.Printf("Erro ao adicionar política: %v", err)
			} else if ok {
				log.Printf("Política adicionada: %v", politica)
			}
		}
	}
}
