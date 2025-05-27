import { createAsyncThunk } from "@reduxjs/toolkit";
import authHeader, { api } from "./api";
import IAvaliacao from "../interface/IAvaliacao/IAvaliacao";

export const getAvaliacao = createAsyncThunk("avaliacao/getAvaliacao", async () => {
    try {
        const response = await api.get("avaliacao/", {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})

export const getAvaliacaoId = createAsyncThunk("avaliacao/getAvaliacaoId", async (avaliacao: IAvaliacao) => {
    try {
        const response = await api.get(`avaliacao/${avaliacao.ID}/`, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})


export const addAvaliacao = createAsyncThunk("avaliacao/addAvaliacao", async (avaliacao: IAvaliacao) => {
    try {
        const response = await api.post("avaliacao/", avaliacao, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})

export const updateAvaliacao = createAsyncThunk("avaliacao/updateAvaliacao",
    async (avaliacao: IAvaliacao) => {
        try {
            const response = await api.patch(`avaliacao/${avaliacao.ID}/`, avaliacao, {headers: authHeader()});
            return response.data
        } catch (error) {
            console.log(error)
        }
    })

export const deleteAvaliacao = createAsyncThunk("avaliacao/deleteAvaliacao", async (avaliacaoId: number) => {
    try {
        const response = await api.delete(`avaliacao/${avaliacaoId}/`, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})
