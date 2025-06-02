export default interface IAvaliacao {
	ID: number,
	AlunoID: number,
	ToraxInspirado: number,
	ToraxNormal: number,
	ToraxExpirado: number,
	Cintura: number,
	Abdomen: number,
	Quadril: number,
	AnteBracoDireito: number,
	AnteBracoEsquerdo: number,
	BracoContraidoDireito: number,
	BracoContraidoEsquerdo: number,
	BracoRelaxadoDireito: number,
	BracoRelaxadoEsquerdo: number,
	CoxaDireita: number,
	CoxaEsquerda: number,
	PernaDireita: number,
	PernaEsquerda: number,
	//dobras cutaneas
	//homens
	Torax: number,
	AbdomenMasc: number,
	//Coxa       float64 `json:"Coxa"`

	//dobras cutaneas
	//mulheres
	SupraIliaca: number,
	Triceps: number,
	Coxa: number,
}