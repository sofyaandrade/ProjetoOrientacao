package migrations

import (
	"fmt"

	"projeto-back/models"
	"projeto-back/utils"

	"gorm.io/gorm"
)

var listaUsuarios = []models.Usuario{
	{
		Nome:            "admin",
		Email:           "admin",
		Telefone:        0,
		Password:        "$2a$10$9RL6Ktc0tE3eSPiIA7x9g.AD9A6uW.CT4LcmQPU5jHL6HU5GR23UW",
		UsuarioPerfilID: 1,
	},
}

func InicializaUsuariosBasicos(db *gorm.DB) {
	var usuarios []models.Usuario

	err := db.Find(&usuarios).Error

	if err != nil {
		fmt.Println("não foi possível localizar a tabela usuários: ", err)
		utils.ErroLog.Println("não foi possível localizar a tabela usuários: ", err)

	}

	if len(usuarios) == 0 {
		for i := range listaUsuarios {
			err := db.Debug().Model(&models.Usuario{}).Create(&listaUsuarios[i]).Error
			if err != nil {
				fmt.Println("não foi possível inserir usuários na tabela: ", err)
				utils.ErroLog.Println("não foi possível inserir usuários na tabela: ", err)
			}
		}
	}
}
