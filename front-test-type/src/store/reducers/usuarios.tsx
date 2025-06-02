import { createSlice } from '@reduxjs/toolkit'
import { addUsuario, deleteUsuario, getUsuarios, updateUsuario } from '../../service/usuarioService';


const usuarioSlice = createSlice({
  name: 'usuario',
  initialState: {
    list: {
      isLoading: false,
      status: "",
      values: []
    },
    save: {
      isSaving: false,
      isDeleting: false
    }
  },
  reducers: {

  },
  extraReducers: builder => {
    builder.addCase(getUsuarios.pending, (state) => {
      state.list.status = "pending"
      state.list.isLoading = true
    })
    builder.addCase(getUsuarios.fulfilled, (state, { payload }) => {
      state.list.status = "success"
      state.list.values = payload
      state.list.isLoading = false
    })
    builder.addCase(getUsuarios.rejected, (state, action) => {
      state.list.status = "failed"
      state.list.isLoading = false
    })

    builder.addCase(addUsuario.pending, (state) => {
      state.save.isSaving = true
    })
    builder.addCase(addUsuario.fulfilled, (state, action) => {
      state.save.isSaving = false
    })
    builder.addCase(addUsuario.rejected, (state, action) => {
      state.save.isSaving = false
    })


    builder.addCase(updateUsuario.pending, (state) => {
      state.save.isSaving = true
    })
    builder.addCase(updateUsuario.fulfilled, (state, action) => {
      state.save.isSaving = false
    })
    builder.addCase(updateUsuario.rejected, (state, action) => {
      state.save.isSaving = false
    })

    builder.addCase(deleteUsuario.pending, (state) => {
      state.save.isSaving = true
    })
    builder.addCase(deleteUsuario.fulfilled, (state, action) => {
      state.save.isSaving = false
    })
    builder.addCase(deleteUsuario.rejected, (state, action) => {
      state.save.isSaving = false
    })

  }
})

export default usuarioSlice.reducer
