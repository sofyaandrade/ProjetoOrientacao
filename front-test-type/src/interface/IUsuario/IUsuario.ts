import IPerfil from "./IPerfil"

export default interface IUsuario {
    ID: number,
    Nome: string,
    Telefone: number,
    Email: string,
    Password: string,
    UsuarioPerfilID: number
    UsuarioPerfil: IPerfil,

}

export interface IUsuarioList extends IUsuario {
    dateCreated: any
}