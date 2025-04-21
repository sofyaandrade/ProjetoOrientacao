package repository

import (
	"projeto-back/models"

	"gorm.io/gorm"
)

type clienteRepository struct {
	database *gorm.DB
}

func NovoClienteRepository(db *gorm.DB) models.ClienteRepository {
	return &clienteRepository{
		database: db,
	}
}

func (cr *clienteRepository) CriarCliente(cliente *models.Cliente) error {
	err := cr.database.Create(cliente).Error
	return err
}

func (cr *clienteRepository) BuscarTodosClientes() (*[]models.Cliente, error) {
	var clientes *[]models.Cliente
	err := cr.database.Select("id", "nome", "telefone", "email").Preload("Avaliacoes").Find(&clientes).Error
	return clientes, err
}

func (cr *clienteRepository) BuscarClientePorId(id uint) (*models.Cliente, error) {
	var cliente *models.Cliente
	err := cr.database.Select("id", "nome", "telefone", "email").Preload("Avaliacoes").First(&cliente, id).Error

	return cliente, err
}

func (cr *clienteRepository) BuscarClientePorEmail(email string) (*models.Cliente, error) {
	var cliente *models.Cliente
	err := cr.database.Where(models.Cliente{Email: email}).Preload("Avaliacoes").First(&cliente).Error
	return cliente, err
}

func (cr *clienteRepository) EditarCliente(cliente *models.Cliente, id uint) error {
	err := cr.database.Model(&cliente).Where("id = ?", id).Updates(cliente).Error
	return err
}

func (cr *clienteRepository) DeletarCliente(id uint) error {
	var cliente *models.Cliente
	err := cr.database.Preload("Avaliacoes").First(&cliente, id).Error
	if err != nil {
		return err
	}

	err = cr.database.Where("cliente_id = ?", cliente.ID).Delete(&cliente.Avaliacoes).Error
	if err != nil {
		return err
	}

	err = cr.database.Delete(&cliente).Error
	return err
}

func (cr *clienteRepository) ExisteClientePorId(id uint) error {
	var cliente *models.Cliente
	err := cr.database.First(&cliente, id).Error
	return err
}
