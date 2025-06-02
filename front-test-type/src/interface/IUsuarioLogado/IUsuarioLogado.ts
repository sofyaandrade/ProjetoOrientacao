import IPerfil from "../IUsuario/IPerfil";

export default interface IUsuarioLogado {
    ID: number,
    Nome: string,
    Email: string,
    Password: string,
    UsuarioPerfilID: number
    UsuarioPerfil: IPerfil,
}