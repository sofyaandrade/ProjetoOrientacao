package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Nome       string       `json:"nome"`
	Descricao  string       `json:"descricao"`
	Telefone   string       `json:"telefone"`
	Email      string       `json:"email"`
	Genero     string       `json:"genero"`
	Avaliacoes []*Avaliacao `json:"avaliacoes"`
}

type ClienteRepository interface {
	CriarCliente(clientes *Cliente) error
	BuscarTodosClientes() (*[]Cliente, error)
	BuscarClientePorId(id uint) (*Cliente, error)
	BuscarClientePorEmail(email string) (*Cliente, error)
	EditarCliente(cliente *Cliente, id uint) error
	DeletarCliente(id uint) error
	ExisteClientePorId(id uint) error
}

type ClienteUsecase interface {
	CriarCliente(clientes *Cliente) error
	BuscarTodosClientes() (*[]Cliente, error)
	BuscarClientePorId(id uint) (*Cliente, error)
	BuscarClientePorEmail(email string) (*Cliente, error)
	EditarCliente(cliente *Cliente, id uint) error
	DeletarCliente(id uint) error
	ExisteClientePorId(id uint) error
}
