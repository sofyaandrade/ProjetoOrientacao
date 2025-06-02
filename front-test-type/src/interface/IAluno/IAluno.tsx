import IAvaliacao from "../IAvaliacao/IAvaliacao";

export default interface IAluno{
    ID: number,
    Nome: string,
    Descricao: string,
    Telefone: number,
    Email: string,
    Avaliacao: IAvaliacao
}