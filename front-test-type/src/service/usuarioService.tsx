
import { createAsyncThunk } from "@reduxjs/toolkit";
import IUsuario from "../interface/IUsuario/IUsuario";
import authHeader, { api } from "./api";

export const getUsuarios = createAsyncThunk("usuarios/getUsuarios", async () => {
    try {
        const response = await api.get("usuarios/", {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})

export const addUsuario = createAsyncThunk("usuarios/addUsuarios", async (usuario: IUsuario) => {
    try {
        const response = await api.post("usuarios/", usuario, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})

export const updateUsuario = createAsyncThunk("usuarios/updateUsuarios",
    async (usuario: IUsuario) => {
        try {
            const response = await api.patch(`usuarios/${usuario.ID}/`, usuario, {headers: authHeader()});
            return response.data
        } catch (error) {
            console.log(error)
        }
    })

export const deleteUsuario = createAsyncThunk("usuarios/deleteUsuarios", async (usuarioId: number) => {
    try {
        const response = await api.delete(`usuarios/${usuarioId}/`, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})

export const getPerfilUsuarioLogado = createAsyncThunk("usuarios/getUsuarios", async (usuarioId: number) => {
    try {
        const response = await api.get(`/usuario-perfil/usuario/${usuarioId}/`, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})