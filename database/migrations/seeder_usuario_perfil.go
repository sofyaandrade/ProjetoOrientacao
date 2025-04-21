package migrations

import (
	"fmt"

	"projeto-back/models"
	"projeto-back/utils"

	"gorm.io/gorm"
)

var listaUsuarioPerfil = []models.UsuarioPerfil{
	{
		Descricao: "ADMINISTRADOR",
	},
	{
		Descricao: "CLIENTE",
	},
}

func InicializaUsuarioPerfil(db *gorm.DB) {
	var usuarioPerfis []models.UsuarioPerfil

	err := db.Find(&usuarioPerfis).Error

	if err != nil {
		fmt.Println("não foi possível localizar a tabela usuário perfis: ", err)
		utils.ErroLog.Println("não foi possível localizar a tabela usuários perfis: ", err)
	}

	if len(usuarioPerfis) == 0 {
		for i := range listaUsuarioPerfil {
			err := db.Debug().Model(&models.UsuarioPerfil{}).Create(&listaUsuarioPerfil[i]).Error
			if err != nil {
				fmt.Println("não foi possível inserir perfil de usuários na tabela: ", err)
				utils.ErroLog.Println("não foi possível inserir perfil de usuários na tabela: ", err)
			}
		}
	}
}
