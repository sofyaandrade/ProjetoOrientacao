package usecase

import "projeto-back/models"

type avaliacaoUsecase struct {
	avaliacaoRepository models.AvaliacaoRepository
}

func NovaAvaliacaoUsecase(avaliacaoRepository models.AvaliacaoRepository) models.AvaliacaoUsecase {
	return &avaliacaoUsecase{
		avaliacaoRepository: avaliacaoRepository,
	}
}

func (cs *avaliacaoUsecase) CriarAvaliacao(avaliacao *models.Avaliacao) error {
	return cs.avaliacaoRepository.CriarAvaliacao(avaliacao)
}

func (cs *avaliacaoUsecase) BuscarTodasAvaliacoes() (*[]models.Avaliacao, error) {
	return cs.avaliacaoRepository.BuscarTodasAvaliacoes()
}

func (cs *avaliacaoUsecase) BuscarAvaliacaoPorId(id uint) (*models.Avaliacao, error) {
	return cs.avaliacaoRepository.BuscarAvaliacaoPorId(id)
}

func (cs *avaliacaoUsecase) BuscarAvaliacoesPorAluno(id uint) ([]models.Avaliacao, error) {
	return cs.avaliacaoRepository.BuscarAvaliacoesPorAluno(id)
}

func (cs *avaliacaoUsecase) EditarAvaliacao(avaliacao *models.Avaliacao, id uint) error {
	return cs.avaliacaoRepository.EditarAvaliacao(avaliacao, id)
}

func (cs *avaliacaoUsecase) DeletarAvaliacao(id uint) error {
	return cs.avaliacaoRepository.DeletarAvaliacao(id)
}

func (cs *avaliacaoUsecase) ExisteAvaliacaoPorId(id uint) error {
	return cs.avaliacaoRepository.ExisteAvaliacaoPorId(id)
}
