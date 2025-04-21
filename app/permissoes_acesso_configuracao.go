package app

import (
	"fmt"
	"projeto-back/database/migrations"
	"projeto-back/models"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var rbacModel string = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`

func PermissoesAcessoConfiguracao(db *gorm.DB) *casbin.Enforcer {

	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &models.PermissoesAcesso{}, "permissoes_acesso")
	if err != nil {
		fmt.Println("Falha ao iniciar casbin adapter" + err.Error())
	}

	model, err := model.NewModelFromString(rbacModel)
	if err != nil {
		fmt.Println(err.Error())
	}

	enforcer, err := casbin.NewEnforcer(model, adapter)
	if err != nil {
		fmt.Println("Erro ao iniciar casbim enforcer" + err.Error())
	}

	enforcer.LoadPolicy()

	migrations.InicializaPermissoesAcesso(enforcer)

	return enforcer
}
