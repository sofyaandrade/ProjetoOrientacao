package usecase

import "projeto-back/models"

type clienteUsecase struct {
	clienteRepository models.ClienteRepository
}

func NovoClienteUsecase(clienteRepository models.ClienteRepository) models.ClienteUsecase {
	return &clienteUsecase{
		clienteRepository: clienteRepository,
	}
}

func (cs *clienteUsecase) CriarCliente(cliente *models.Cliente) error {
	return cs.clienteRepository.CriarCliente(cliente)
}

func (cs *clienteUsecase) BuscarTodosClientes() (*[]models.Cliente, error) {
	return cs.clienteRepository.BuscarTodosClientes()
}

func (cs *clienteUsecase) BuscarClientePorId(id uint) (*models.Cliente, error) {
	return cs.clienteRepository.BuscarClientePorId(id)
}

func (cs *clienteUsecase) BuscarClientePorEmail(email string) (*models.Cliente, error) {
	return cs.clienteRepository.BuscarClientePorEmail(email)
}

func (cs *clienteUsecase) EditarCliente(cliente *models.Cliente, id uint) error {
	return cs.clienteRepository.EditarCliente(cliente, id)
}

func (cs *clienteUsecase) DeletarCliente(id uint) error {
	return cs.clienteRepository.DeletarCliente(id)
}

func (cs *clienteUsecase) ExisteClientePorId(id uint) error {
	return cs.clienteRepository.ExisteClientePorId(id)
}
