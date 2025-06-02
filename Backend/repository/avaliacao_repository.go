package repository

import (
	"projeto-back/models"

	"gorm.io/gorm"
)

type avaliacaoRepository struct {
	database *gorm.DB
}

func NovaAvaliacaoRepository(db *gorm.DB) models.AvaliacaoRepository {
	return &avaliacaoRepository{
		database: db,
	}
}

func (cr *avaliacaoRepository) CriarAvaliacao(avaliacao *models.Avaliacao) error {
	err := cr.database.Create(avaliacao).Error
	return err
}

func (cr *avaliacaoRepository) BuscarTodasAvaliacoes() (*[]models.Avaliacao, error) {
	var avaliacaos *[]models.Avaliacao
	err := cr.database.Select("id", "aluno_id", "data_avaliacao", "torax_inspirado", "torax_normal", "torax_expirado", "cintura", "abdomen", "quadril", "ante_braco_direito", "ante_braco_esquerdo", "braco_contraido_direito", "braco_contraido_esquerdo", "braco_relaxado_direito", "braco_relaxado_esquerdo", "coxa_direita", "coxa_esquerda", "perna_direita", "perna_esquerda", "torax", "abdomem_masc", "coxa_masc", "supra_liaca", "triceps", "coxa_fem").Find(&avaliacaos).Error
	return avaliacaos, err
}

func (cr *avaliacaoRepository) BuscarAvaliacaoPorId(id uint) (*models.Avaliacao, error) {
	var avaliacao *models.Avaliacao
	err := cr.database.Select("id", "aluno_id", "data_avaliacao", "torax_inspirado", "torax_normal", "torax_expirado", "cintura", "abdomen", "quadril", "ante_braco_direito", "ante_braco_esquerdo", "braco_contraido_direito", "braco_contraido_esquerdo", "braco_relaxado_direito", "braco_relaxado_esquerdo", "coxa_direita", "coxa_esquerda", "perna_direita", "perna_esquerda", "torax", "abdomem_masc", "coxa_masc", "supra_liaca", "triceps", "coxa_fem").First(&avaliacao, id).Error
	return avaliacao, err
}

func (cr *avaliacaoRepository) BuscarAvaliacoesPorAluno(idAluno uint) ([]models.Avaliacao, error) {
	var avaliacoes []models.Avaliacao
	err := cr.database.Select("id", "aluno_id", "data_avaliacao", "torax_inspirado", "torax_normal", "torax_expirado", "cintura", "abdomen", "quadril", "ante_braco_direito", "ante_braco_esquerdo", "braco_contraido_direito", "braco_contraido_esquerdo", "braco_relaxado_direito", "braco_relaxado_esquerdo", "coxa_direita", "coxa_esquerda", "perna_direita", "perna_esquerda", "torax", "abdomem_masc", "coxa_masc", "supra_liaca", "triceps", "coxa_fem").Where("aluno_id = ?", idAluno).Find(&avaliacoes).Error
	return avaliacoes, err
}

func (cr *avaliacaoRepository) EditarAvaliacao(avaliacao *models.Avaliacao, id uint) error {
	err := cr.database.Model(&avaliacao).Where("id = ?", id).Updates(avaliacao).Error
	return err
}

func (cr *avaliacaoRepository) DeletarAvaliacao(id uint) error {
	var avaliacao *models.Avaliacao
	err := cr.database.Delete(&avaliacao, id).Error
	return err
}

func (cr *avaliacaoRepository) ExisteAvaliacaoPorId(id uint) error {
	var avaliacao *models.Avaliacao
	err := cr.database.First(&avaliacao, id).Error
	return err
}
