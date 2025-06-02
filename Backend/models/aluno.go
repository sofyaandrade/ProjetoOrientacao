package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome       string       `json:"Nome"`
	Descricao  string       `json:"Descricao"`
	Telefone   int          `json:"Telefone"`
	Email      string       `json:"Email"`
	Avaliacoes []*Avaliacao `json:"Avaliacoes"`
}

type AlunoRepository interface {
	CriarAluno(alunos *Aluno) error
	BuscarTodosAlunos() (*[]Aluno, error)
	BuscarAlunoPorId(id uint) (*Aluno, error)
	BuscarAlunoPorEmail(email string) (*Aluno, error)
	EditarAluno(aluno *Aluno, id uint) error
	DeletarAluno(id uint) error
	ExisteAlunoPorId(id uint) error
}

type AlunoUsecase interface {
	CriarAluno(alunos *Aluno) error
	BuscarTodosAlunos() (*[]Aluno, error)
	BuscarAlunoPorId(id uint) (*Aluno, error)
	BuscarAlunoPorEmail(email string) (*Aluno, error)
	EditarAluno(aluno *Aluno, id uint) error
	DeletarAluno(id uint) error
	ExisteAlunoPorId(id uint) error
}
