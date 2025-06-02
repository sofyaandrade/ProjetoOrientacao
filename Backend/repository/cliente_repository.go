package repository

import (
	"projeto-back/models"

	"gorm.io/gorm"
)

type alunoRepository struct {
	database *gorm.DB
}

func NovoAlunoRepository(db *gorm.DB) models.AlunoRepository {
	return &alunoRepository{
		database: db,
	}
}

func (cr *alunoRepository) CriarAluno(aluno *models.Aluno) error {
	err := cr.database.Create(aluno).Error
	return err
}

func (cr *alunoRepository) BuscarTodosAlunos() (*[]models.Aluno, error) {
	var alunos *[]models.Aluno
	err := cr.database.Select("id", "nome", "descricao", "telefone", "email").Preload("Avaliacoes").Find(&alunos).Error
	return alunos, err
}

func (cr *alunoRepository) BuscarAlunoPorId(id uint) (*models.Aluno, error) {
	var aluno *models.Aluno
	err := cr.database.Select("id", "nome", "descricao", "telefone", "email").Preload("Avaliacoes").First(&aluno, id).Error

	return aluno, err
}

func (cr *alunoRepository) BuscarAlunoPorEmail(email string) (*models.Aluno, error) {
	var aluno *models.Aluno
	err := cr.database.Where(models.Aluno{Email: email}).Preload("Avaliacoes").First(&aluno).Error
	return aluno, err
}

func (cr *alunoRepository) EditarAluno(aluno *models.Aluno, id uint) error {
	err := cr.database.Model(&aluno).Where("id = ?", id).Updates(aluno).Error
	return err
}

func (cr *alunoRepository) DeletarAluno(id uint) error {
	var aluno *models.Aluno
	err := cr.database.Preload("Avaliacoes").First(&aluno, id).Error
	if err != nil {
		return err
	}

	err = cr.database.Where("aluno_id = ?", aluno.ID).Delete(&aluno.Avaliacoes).Error
	if err != nil {
		return err
	}

	err = cr.database.Delete(&aluno).Error
	return err
}

func (cr *alunoRepository) ExisteAlunoPorId(id uint) error {
	var aluno *models.Aluno
	err := cr.database.First(&aluno, id).Error
	return err
}
