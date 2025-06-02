import { configureStore } from '@reduxjs/toolkit';
import usuarioLogado from './reducers/usuarioLogado';
import usuarios from './reducers/usuarios';


const store = configureStore({
    reducer: {
        usuarioLogado: usuarioLogado,
        usuario: usuarios,
    },
});

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export default store
