package models

import (
	"time"

	"gorm.io/gorm"
)

type Avaliacao struct {
	gorm.Model
	ClienteID              uint      `json:"cliente_id"`
	DataAvaliacao          time.Time `json:"data_avaliacao"`
	ToraxInspirado         float64   `json:"torax_inspirado"`
	ToraxNormal            float64   `json:"torax_normal"`
	ToraxExpirado          float64   `json:"torax_expirado"`
	Cintura                float64   `json:"cintura"`
	Abdomen                float64   `json:"abdomen"`
	Quadril                float64   `json:"quadril"`
	AnteBracoDireito       float64   `json:"ante_braco_direito"`
	AnteBracoEsquerdo      float64   `json:"ante_braco_esquerdo"`
	BracoContraidoDireito  float64   `json:"braco_contraido_direito"`
	BracoContraidoEsquerdo float64   `json:"braco_contraido_esquerdo"`
	BracoRelaxadoDireito   float64   `json:"braco_relaxado_direito"`
	BracoRelaxadoEsquerdo  float64   `json:"braco_relaxado_esquerdo"`
	CoxaDireita            float64   `json:"coxa_direita"`
	CoxaEsquerda           float64   `json:"coxa_esquerda"`
	PernaDireita           float64   `json:"perna_direita"`
	PernaEsquerda          float64   `json:"perna_esquerda"`

	//dobras cutaneas - homens
	Torax       float64 `json:"torax"`
	AbdomemMasc float64 `json:"abdomem_masc"`
	CoxaMasc    float64 `json:"coxa_masc"`

	//dobras cutaneas - mulheres
	SupraLiaca float64 `json:"supra_iliaca"`
	Triceps    float64 `json:"triceps"`
	CoxaFem    float64 `json:"coxa_fem"`
}

type AvaliacaoRepository interface {
	CriarAvaliacao(avaliacao *Avaliacao) error
	BuscarTodasAvaliacoes() (*[]Avaliacao, error)
	BuscarAvaliacaoPorId(id uint) (*Avaliacao, error)
	BuscarAvaliacoesPorCliente(idCliente uint) ([]Avaliacao, error)
	EditarAvaliacao(avaliacao *Avaliacao, id uint) error
	DeletarAvaliacao(id uint) error
	ExisteAvaliacaoPorId(id uint) error
}

type AvaliacaoUsecase interface {
	CriarAvaliacao(avaliacao *Avaliacao) error
	BuscarTodasAvaliacoes() (*[]Avaliacao, error)
	BuscarAvaliacaoPorId(id uint) (*Avaliacao, error)
	BuscarAvaliacoesPorCliente(idCliente uint) ([]Avaliacao, error)
	EditarAvaliacao(avaliacao *Avaliacao, id uint) error
	DeletarAvaliacao(id uint) error
	ExisteAvaliacaoPorId(id uint) error
}
