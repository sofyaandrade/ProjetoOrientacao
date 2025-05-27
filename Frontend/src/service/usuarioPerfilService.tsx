
import { createAsyncThunk } from "@reduxjs/toolkit";
import IPerfil from "../interface/IUsuario/IPerfil";
import authHeader, { api } from "./api";

export const getUsuarioPerfil = createAsyncThunk("usuario-perfil/getUsuarioPerfil", async () => {
    try {
        const response = await api.get("usuario-perfil/", {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})

export const addUsuarioPerfil = createAsyncThunk("usuario-perfil/addUsuarioPerfil", async (perfil: IPerfil) => {
    try {
        const response = await api.post("usuario-perfil/", perfil, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})

export const updateUsuarioPerfil = createAsyncThunk("usuario-perfil/updateUsuarioPerfil",
    async (perfil: IPerfil) => {
        try {
            const response = await api.patch(`usuario-perfil/${perfil.ID}/`, perfil, {headers: authHeader()});
            return response.data
        } catch (error) {
            console.log(error)
        }
    })

export const deleteUsuarioPerfil = createAsyncThunk("usuario-perfil/deleteUsuarioPerfil", async (perfilId: number) => {
    try {
        const response = await api.delete(`usuario-perfil/${perfilId}/`, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})