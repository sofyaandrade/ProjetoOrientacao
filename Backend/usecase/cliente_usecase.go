package usecase

import "projeto-back/models"

type alunoUsecase struct {
	alunoRepository models.AlunoRepository
}

func NovoAlunoUsecase(alunoRepository models.AlunoRepository) models.AlunoUsecase {
	return &alunoUsecase{
		alunoRepository: alunoRepository,
	}
}

func (cs *alunoUsecase) CriarAluno(aluno *models.Aluno) error {
	return cs.alunoRepository.CriarAluno(aluno)
}

func (cs *alunoUsecase) BuscarTodosAlunos() (*[]models.Aluno, error) {
	return cs.alunoRepository.BuscarTodosAlunos()
}

func (cs *alunoUsecase) BuscarAlunoPorId(id uint) (*models.Aluno, error) {
	return cs.alunoRepository.BuscarAlunoPorId(id)
}

func (cs *alunoUsecase) BuscarAlunoPorEmail(email string) (*models.Aluno, error) {
	return cs.alunoRepository.BuscarAlunoPorEmail(email)
}

func (cs *alunoUsecase) EditarAluno(aluno *models.Aluno, id uint) error {
	return cs.alunoRepository.EditarAluno(aluno, id)
}

func (cs *alunoUsecase) DeletarAluno(id uint) error {
	return cs.alunoRepository.DeletarAluno(id)
}

func (cs *alunoUsecase) ExisteAlunoPorId(id uint) error {
	return cs.alunoRepository.ExisteAlunoPorId(id)
}
