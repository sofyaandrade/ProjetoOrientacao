import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit'
import authHeader, { api } from '../../service/api';
import IUsuarioLogado from '../../interface/IUsuarioLogado/IUsuarioLogado';

type InitialState = {
	loading: boolean
	usuarioLogado: IUsuarioLogado
	error: string
}

const initialState: InitialState = {
	loading: false,
	usuarioLogado: {
		ID: 0,
		Nome: '',
		Email: '',
		Password: '',
		UsuarioPerfilID: 0,
		UsuarioPerfil: {
			ID: 0,
			Descricao: "",
			Usuarios: [],
		},
	},
	error: '',
}

export const fetchUsuarioLogado = createAsyncThunk('usuarios/getUsuarioLogado', async () => {
	let idUsuarioLogado = sessionStorage.getItem('Id');
	const response = await api.get(`/usuarios/${idUsuarioLogado}/`, {headers: authHeader()});

	sessionStorage.setItem("User", response.data.Nome);

	return response.data;
})

const usuarioLogadoSlice = createSlice({
	name: 'usuarioLogado',
	initialState,
	reducers: {
		setId(state, action: PayloadAction<string>) {
			let usuarioLogadoId = state.usuarioLogado.ID.toString()
			usuarioLogadoId = action.payload
		}
	},
	extraReducers: builder => {
		builder.addCase(fetchUsuarioLogado.pending, (state) => {
			state.loading = true
		})
		builder.addCase(fetchUsuarioLogado.fulfilled, (state, action: PayloadAction<IUsuarioLogado>) => {
			state.loading = false
			state.usuarioLogado = action.payload
			state.error = ''
		})
		builder.addCase(fetchUsuarioLogado.rejected, (state, action) => {
			state.loading = false
			state.usuarioLogado = {
				ID: 0,
				Nome: '',
				Email: '',
				Password: '',
				UsuarioPerfilID: 0,
				UsuarioPerfil: {
					ID: 0,
					Descricao: "",
					Usuarios: [],
				},
			}
			state.error = action.error.message || 'Alguma coisa deu errado'
		})
	}
})

export default usuarioLogadoSlice.reducer
export const { setId } = usuarioLogadoSlice.actions
