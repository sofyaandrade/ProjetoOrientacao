import { createAsyncThunk } from "@reduxjs/toolkit"
import authHeader, { api } from "./api"
import IAlunos from "../interface/IAluno/IAluno"


export const getAlunos = createAsyncThunk("aluno/getAlunos", async () => {
    try {
        const response = await api.get("aluno/", {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})

export const getAlunosId = createAsyncThunk("aluno/getAlunosId", async (aluno: IAlunos) => {
    try {
        const response = await api.get(`aluno/${aluno.ID}/`, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})


export const addAlunos = createAsyncThunk("aluno/addAlunos", async (aluno: IAlunos) => {
    try {
        const response = await api.post("aluno/", aluno, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})

export const updateAlunos = createAsyncThunk("aluno/updateAlunos",
    async (aluno: IAlunos) => {
        try {
            const response = await api.patch(`aluno/${aluno.ID}/`, aluno, {headers: authHeader()});
            return response.data
        } catch (error) {
            console.log(error)
        }
    })

export const deleteAlunos = createAsyncThunk("aluno/deleteAlunos", async (alunoId: number) => {
    try {
        const response = await api.delete(`aluno/${alunoId}/`, {headers: authHeader()})
        return response.data
    } catch (error) {
        console.log(error)
    }
})